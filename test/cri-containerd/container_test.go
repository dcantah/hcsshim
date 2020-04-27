// +build functional

package cri_containerd

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	testutilities "github.com/Microsoft/hcsshim/test/functional/utilities"
	"github.com/sirupsen/logrus"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func runLogRotationContainer(t *testing.T, sandboxRequest *runtime.RunPodSandboxRequest, request *runtime.CreateContainerRequest, log string, logArchive string) {
	client := newTestRuntimeClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	podID := runPodSandbox(t, client, ctx, sandboxRequest)
	defer removePodSandbox(t, client, ctx, podID)
	defer stopPodSandbox(t, client, ctx, podID)

	request.PodSandboxId = podID
	request.SandboxConfig = sandboxRequest.Config

	containerID := createContainer(t, client, ctx, request)
	defer removeContainer(t, client, ctx, containerID)

	startContainer(t, client, ctx, containerID)
	defer stopContainer(t, client, ctx, containerID)

	// Give some time for log output to accumulate.
	time.Sleep(3 * time.Second)

	// Rotate the logs. This is done by first renaming the existing log file,
	// then calling ReopenContainerLog to cause containerd to start writing to
	// a new log file.

	if err := os.Rename(log, logArchive); err != nil {
		t.Fatalf("failed to rename log: %v", err)
	}

	if _, err := client.ReopenContainerLog(ctx, &runtime.ReopenContainerLogRequest{ContainerId: containerID}); err != nil {
		t.Fatalf("failed to reopen log: %v", err)
	}

	// Give some time for log output to accumulate.
	time.Sleep(3 * time.Second)
}

func runContainerLifetime(t *testing.T, client runtime.RuntimeServiceClient, ctx context.Context, containerID string) {
	defer removeContainer(t, client, ctx, containerID)
	startContainer(t, client, ctx, containerID)
	stopContainer(t, client, ctx, containerID)
}

func Test_RotateLogs_LCOW(t *testing.T) {
	requireFeatures(t, featureLCOW)

	image := "alpine:latest"
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("failed creating temp dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("failed deleting temp dir: %v", err)
		}
	}()
	log := filepath.Join(dir, "log.txt")
	logArchive := filepath.Join(dir, "log-archive.txt")

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, image})
	logrus.SetLevel(logrus.DebugLevel)

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name() + "-Sandbox",
				Namespace: testNamespace,
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}

	request := &runtime.CreateContainerRequest{
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: image,
			},
			Command: []string{
				"ash",
				"-c",
				"i=0; while true; do echo $i; i=$(expr $i + 1); sleep .1; done",
			},
			LogPath: log,
			Linux:   &runtime.LinuxContainerConfig{},
		},
	}

	runLogRotationContainer(t, sandboxRequest, request, log, logArchive)

	// Make sure we didn't lose any values while rotating. First set of output
	// should be in logArchive, followed by the output in log.

	logArchiveFile, err := os.Open(logArchive)
	if err != nil {
		t.Fatal(err)
	}
	defer logArchiveFile.Close()

	logFile, err := os.Open(log)
	if err != nil {
		t.Fatal(err)
	}
	defer logFile.Close()

	s := bufio.NewScanner(io.MultiReader(logArchiveFile, logFile))
	expected := 0
	for s.Scan() {
		v := strings.Fields(s.Text())
		n, err := strconv.Atoi(v[len(v)-1])
		if err != nil {
			t.Fatalf("failed to parse log value as integer: %v", err)
		}
		if n != expected {
			t.Fatalf("missing expected output value: %v (got %v)", expected, n)
		}
		expected++
	}
}

func Test_RunContainer_Events_LCOW(t *testing.T) {
	requireFeatures(t, featureLCOW)

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, imageLcowAlpine})
	client := newTestRuntimeClient(t)

	podctx, podcancel := context.WithCancel(context.Background())
	defer podcancel()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	targetNamespace := "k8s.io"

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name(),
				Uid:       "0",
				Namespace: testNamespace,
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}

	podID := runPodSandbox(t, client, podctx, sandboxRequest)
	defer removePodSandbox(t, client, podctx, podID)
	defer stopPodSandbox(t, client, podctx, podID)

	request := &runtime.CreateContainerRequest{
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageLcowAlpine,
			},
			Command: []string{
				"top",
			},
			Linux: &runtime.LinuxContainerConfig{},
		},
		PodSandboxId:  podID,
		SandboxConfig: sandboxRequest.Config,
	}

	topicNames, filters := getTargetRunTopics()
	eventService := newTestEventService(t)
	stream, errs := eventService.Subscribe(ctx, filters...)

	containerID := createContainer(t, client, podctx, request)
	runContainerLifetime(t, client, podctx, containerID)

	for _, topic := range topicNames {
		select {
		case env := <-stream:
			if topic != env.Topic {
				t.Fatalf("event topic %v does not match expected topic %v", env.Topic, topic)
			}
			if targetNamespace != env.Namespace {
				t.Fatalf("event namespace %v does not match expected namespace %v", env.Namespace, targetNamespace)
			}
			t.Logf("event topic seen: %v", env.Topic)

			id, _, err := convertEvent(env.Event)
			if err != nil {
				t.Fatalf("topic %v event: %v", env.Topic, err)
			}
			if id != containerID {
				t.Fatalf("event topic %v belongs to container %v, not targeted container %v", env.Topic, id, containerID)
			}
		case e := <-errs:
			t.Fatalf("event subscription err %v", e)
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				t.Fatalf("event %v deadline exceeded", topic)
			}
		}
	}
}

func Test_RunContainer_ForksThenExits_ShowsAsExited_LCOW(t *testing.T) {
	requireFeatures(t, featureLCOW)

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, imageLcowAlpine})
	client := newTestRuntimeClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	podRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name(),
				Namespace: testNamespace,
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}
	podID := runPodSandbox(t, client, ctx, podRequest)
	defer removePodSandbox(t, client, ctx, podID)
	defer stopPodSandbox(t, client, ctx, podID)

	containerRequest := &runtime.CreateContainerRequest{
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageLcowAlpine,
			},
			Command: []string{
				// Fork a background process (that runs forever), then exit.
				"ash",
				"-c",
				"ash -c 'while true; do echo foo; sleep 1; done' &",
			},
			Linux: &runtime.LinuxContainerConfig{},
		},
		PodSandboxId:  podID,
		SandboxConfig: podRequest.Config,
	}
	containerID := createContainer(t, client, ctx, containerRequest)
	defer removeContainer(t, client, ctx, containerID)
	startContainer(t, client, ctx, containerID)
	defer stopContainer(t, client, ctx, containerID)

	// Give the container init time to exit.
	time.Sleep(5 * time.Second)

	// Validate that the container shows as exited. Once the container init
	// dies, the forked background process should be killed off.
	statusResponse, err := client.ContainerStatus(ctx, &runtime.ContainerStatusRequest{ContainerId: containerID})
	if err != nil {
		t.Fatalf("failed to get container status: %v", err)
	}
	if statusResponse.Status.State != runtime.ContainerState_CONTAINER_EXITED {
		t.Fatalf("container expected to be exited but is in state %s", statusResponse.Status.State)
	}
}

func Test_RunContainer_ZeroVPMEM_LCOW(t *testing.T) {
	requireFeatures(t, featureLCOW)

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, imageLcowAlpine})

	client := newTestRuntimeClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name() + "-Sandbox",
				Namespace: testNamespace,
			},
			Annotations: map[string]string{
				"io.microsoft.virtualmachine.lcow.preferredrootfstype":         "initrd",
				"io.microsoft.virtualmachine.devices.virtualpmem.maximumcount": "0",
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}

	podID := runPodSandbox(t, client, ctx, sandboxRequest)
	defer removePodSandbox(t, client, ctx, podID)
	defer stopPodSandbox(t, client, ctx, podID)

	request := &runtime.CreateContainerRequest{
		PodSandboxId: podID,
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageLcowAlpine,
			},
			Command: []string{
				"top",
			},
		},
		SandboxConfig: sandboxRequest.Config,
	}

	containerID := createContainer(t, client, ctx, request)
	runContainerLifetime(t, client, ctx, containerID)
}

func Test_RunContainer_ZeroVPMEM_Multiple_LCOW(t *testing.T) {
	requireFeatures(t, featureLCOW)

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, imageLcowAlpine})

	client := newTestRuntimeClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name() + "-Sandbox",
				Namespace: testNamespace,
			},
			Annotations: map[string]string{
				"io.microsoft.virtualmachine.lcow.preferredrootfstype":         "initrd",
				"io.microsoft.virtualmachine.devices.virtualpmem.maximumcount": "0",
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}

	podID := runPodSandbox(t, client, ctx, sandboxRequest)
	defer removePodSandbox(t, client, ctx, podID)
	defer stopPodSandbox(t, client, ctx, podID)

	request := &runtime.CreateContainerRequest{
		PodSandboxId: podID,
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: "",
			},
			Image: &runtime.ImageSpec{
				Image: imageLcowAlpine,
			},
			Command: []string{
				"top",
			},
		},
		SandboxConfig: sandboxRequest.Config,
	}

	request.Config.Metadata.Name = "Container-1"
	containerIDOne := createContainer(t, client, ctx, request)
	defer removeContainer(t, client, ctx, containerIDOne)
	startContainer(t, client, ctx, containerIDOne)
	defer stopContainer(t, client, ctx, containerIDOne)

	request.Config.Metadata.Name = "Container-2"
	containerIDTwo := createContainer(t, client, ctx, request)
	defer removeContainer(t, client, ctx, containerIDTwo)
	startContainer(t, client, ctx, containerIDTwo)
	defer stopContainer(t, client, ctx, containerIDTwo)
}

func Test_RunContainer_GMSA_WCOW_Process(t *testing.T) {
	requireFeatures(t, featureWCOWProcess, featureGMSA)

	credSpec := gmsaSetup(t)
	pullRequiredImages(t, []string{imageWindowsNanoserver})
	client := newTestRuntimeClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name() + "-Sandbox",
				Namespace: testNamespace,
			},
		},
		RuntimeHandler: wcowProcessRuntimeHandler,
	}

	podID := runPodSandbox(t, client, ctx, sandboxRequest)
	defer removePodSandbox(t, client, ctx, podID)
	defer stopPodSandbox(t, client, ctx, podID)

	request := &runtime.CreateContainerRequest{
		PodSandboxId: podID,
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageWindowsNanoserver,
			},
			Command: []string{
				"cmd",
				"/c",
				"ping",
				"-t",
				"127.0.0.1",
			},
			Windows: &runtime.WindowsContainerConfig{
				SecurityContext: &runtime.WindowsContainerSecurityContext{
					CredentialSpec: credSpec,
				},
			},
		},
		SandboxConfig: sandboxRequest.Config,
	}

	containerID := createContainer(t, client, ctx, request)
	defer removeContainer(t, client, ctx, containerID)
	startContainer(t, client, ctx, containerID)
	defer stopContainer(t, client, ctx, containerID)

	// No klist and no powershell available
	cmd := []string{"cmd", "/c", "set", "USERDNSDOMAIN"}
	containerExecReq := &runtime.ExecSyncRequest{
		ContainerId: containerID,
		Cmd:         cmd,
		Timeout:     20,
	}
	r := execSync(t, client, ctx, containerExecReq)
	if r.ExitCode != 0 {
		t.Fatalf("failed with exit code %d running 'set USERDNSDOMAIN': %s", r.ExitCode, string(r.Stderr))
	}
	// Check for USERDNSDOMAIN environment variable. This acts as a way tell if a
	// user is joined to an Active Directory Domain and is successfully
	// authenticated as a domain identity.
	if !strings.Contains(string(r.Stdout), "USERDNSDOMAIN") {
		t.Fatalf("expected to see USERDNSDOMAIN entry")
	}
}

func Test_RunContainer_GMSA_WCOW_Hypervisor(t *testing.T) {
	t.Skip("GMSA is not supported for Hyper-V isolated containers")
	requireFeatures(t, featureWCOWHypervisor, featureGMSA)

	credSpec := gmsaSetup(t)
	pullRequiredImages(t, []string{imageWindowsNanoserver})
	client := newTestRuntimeClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name() + "-Sandbox",
				Namespace: testNamespace,
			},
		},
		RuntimeHandler: wcowHypervisorRuntimeHandler,
	}

	podID := runPodSandbox(t, client, ctx, sandboxRequest)
	defer removePodSandbox(t, client, ctx, podID)
	defer stopPodSandbox(t, client, ctx, podID)

	request := &runtime.CreateContainerRequest{
		PodSandboxId: podID,
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageWindowsNanoserver,
			},
			Command: []string{
				"cmd",
				"/c",
				"ping",
				"-t",
				"127.0.0.1",
			},
			Windows: &runtime.WindowsContainerConfig{
				SecurityContext: &runtime.WindowsContainerSecurityContext{
					CredentialSpec: credSpec,
				},
			},
		},
		SandboxConfig: sandboxRequest.Config,
	}

	containerID := createContainer(t, client, ctx, request)
	defer removeContainer(t, client, ctx, containerID)
	startContainer(t, client, ctx, containerID)
	defer stopContainer(t, client, ctx, containerID)

	// No klist and no powershell available
	cmd := []string{"cmd", "/c", "set", "USERDNSDOMAIN"}
	containerExecReq := &runtime.ExecSyncRequest{
		ContainerId: containerID,
		Cmd:         cmd,
		Timeout:     20,
	}
	r := execSync(t, client, ctx, containerExecReq)
	if r.ExitCode != 0 {
		t.Fatalf("failed with exit code %d running 'set USERDNSDOMAIN': %s", r.ExitCode, string(r.Stderr))
	}
	// Check for USERDNSDOMAIN environment variable. This acts as a way tell if a
	// user is joined to an Active Directory Domain and is successfully
	// authenticated as a domain identity.
	if !strings.Contains(string(r.Stdout), "USERDNSDOMAIN") {
		t.Fatalf("expected to see USERDNSDOMAIN entry")
	}
}

func Test_RunContainer_NUMA_Nodes_Default_LCOW(t *testing.T) {
	t.Skip("NUMA is not supported for LCOW")
	requireFeatures(t, featureLCOW)

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, imageLcowAlpine})
	client := newTestRuntimeClient(t)

	podctx := context.Background()
	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name(),
				Namespace: testNamespace,
			},
			Annotations: map[string]string{
				"io.microsoft.virtualmachine.computetopology.numa.default": "true",
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}

	podID := runPodSandbox(t, client, podctx, sandboxRequest)
	defer removePodSandbox(t, client, podctx, podID)
	defer stopPodSandbox(t, client, podctx, podID)

	containerRequest := &runtime.CreateContainerRequest{
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageLcowAlpine,
			},
			Command: []string{
				"top",
			},
			Linux: &runtime.LinuxContainerConfig{},
		},
		PodSandboxId:  podID,
		SandboxConfig: sandboxRequest.Config,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	containerID := createContainer(t, client, ctx, containerRequest)
	runContainerLifetime(t, client, ctx, containerID)
}

func Test_RunContainer_NUMA_Nodes_Default_WCOW_Hypervisor(t *testing.T) {
	requireFeatures(t, featureWCOWHypervisor)

	pullRequiredImages(t, []string{imageWindowsNanoserver})
	client := newTestRuntimeClient(t)
	podctx := context.Background()

	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name(),
				Namespace: testNamespace,
			},
			Annotations: map[string]string{
				"io.microsoft.virtualmachine.computetopology.numa.default": "true",
			},
		},
		RuntimeHandler: wcowHypervisorRuntimeHandler,
	}

	podID := runPodSandbox(t, client, podctx, sandboxRequest)
	defer removePodSandbox(t, client, podctx, podID)
	defer stopPodSandbox(t, client, podctx, podID)

	request := &runtime.CreateContainerRequest{
		PodSandboxId: podID,
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageWindowsNanoserver,
			},
			Command: []string{
				"cmd",
				"/c",
				"ping",
				"-t",
				"127.0.0.1",
			},
		},
		SandboxConfig: sandboxRequest.Config,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	containerID := createContainer(t, client, ctx, request)
	runContainerLifetime(t, client, ctx, containerID)
}

func Test_RunContainer_NUMA_Nodes_Custom_LCOW(t *testing.T) {
	t.Skip("NUMA is not supported on LCOW")
	requireFeatures(t, featureLCOW)
	testutilities.RequiresBuild(t, 18943)

	pullRequiredLcowImages(t, []string{imageLcowK8sPause, imageLcowAlpine})
	client := newTestRuntimeClient(t)

	podctx := context.Background()
	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name(),
				Namespace: testNamespace,
			},
			Annotations: map[string]string{
				"io.microsoft.virtualmachine.computetopology.processor.count":                    "2",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodecount":              "2",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.virtualnode":    "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.physicalnode":   "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.virtualsocket":  "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.processorcount": "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.memoryamount":   "512",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.virtualnode":    "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.physicalnode":   "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.virtualsocket":  "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.processorcount": "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.memoryamount":   "512",
			},
		},
		RuntimeHandler: lcowRuntimeHandler,
	}

	podID := runPodSandbox(t, client, podctx, sandboxRequest)
	defer removePodSandbox(t, client, podctx, podID)
	defer stopPodSandbox(t, client, podctx, podID)

	request := &runtime.CreateContainerRequest{
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageLcowAlpine,
			},
			Command: []string{
				"top",
			},
			Linux: &runtime.LinuxContainerConfig{},
		},
		PodSandboxId:  podID,
		SandboxConfig: sandboxRequest.Config,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	containerID := createContainer(t, client, ctx, request)
	runContainerLifetime(t, client, ctx, containerID)
}

func Test_RunContainer_NUMA_Nodes_Custom_WCOW_Hypervisor(t *testing.T) {
	requireFeatures(t, featureWCOWHypervisor)
	testutilities.RequiresBuild(t, 18943)

	pullRequiredImages(t, []string{imageWindowsNanoserver})
	client := newTestRuntimeClient(t)

	podctx := context.Background()
	sandboxRequest := &runtime.RunPodSandboxRequest{
		Config: &runtime.PodSandboxConfig{
			Metadata: &runtime.PodSandboxMetadata{
				Name:      t.Name(),
				Namespace: testNamespace,
			},
			Annotations: map[string]string{
				"io.microsoft.virtualmachine.computetopology.processor.count":                    "2",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodecount":              "2",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.virtualnode":    "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.physicalnode":   "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.virtualsocket":  "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.processorcount": "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.0.memoryamount":   "512",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.virtualnode":    "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.physicalnode":   "0",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.virtualsocket":  "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.processorcount": "1",
				"io.microsoft.virtualmachine.computetopology.numa.virtualnodes.1.memoryamount":   "512",
			},
		},
		RuntimeHandler: wcowHypervisorRuntimeHandler,
	}

	podID := runPodSandbox(t, client, podctx, sandboxRequest)
	defer removePodSandbox(t, client, podctx, podID)
	defer stopPodSandbox(t, client, podctx, podID)

	request := &runtime.CreateContainerRequest{
		PodSandboxId: podID,
		Config: &runtime.ContainerConfig{
			Metadata: &runtime.ContainerMetadata{
				Name: t.Name() + "-Container",
			},
			Image: &runtime.ImageSpec{
				Image: imageWindowsNanoserver,
			},
			Command: []string{
				"cmd",
				"/c",
				"ping",
				"-t",
				"127.0.0.1",
			},
		},
		SandboxConfig: sandboxRequest.Config,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	containerID := createContainer(t, client, ctx, request)
	runContainerLifetime(t, client, ctx, containerID)
}
