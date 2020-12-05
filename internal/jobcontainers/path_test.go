package jobcontainers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFindExecutable(t *testing.T) {
	var (
		testExecutable = "test.exe"
		executablePath = `/path/to/binary`
	)

	if _, err := findExecutable("ping", ""); err != nil {
		t.Fatalf("failed to find executable: %s", err)
	}

	if _, err := findExecutable("ping.exe", ""); err != nil {
		t.Fatalf("failed to find executable: %s", err)
	}

	if _, err := findExecutable("C:\\windows\\system32\\ping", ""); err != nil {
		t.Fatalf("failed to find executable: %s", err)
	}

	if _, err := findExecutable("C:\\windows\\system32\\ping.exe", ""); err != nil {
		t.Fatalf("failed to find executable: %s", err)
	}

	// Create nested directory structure with blank test executables.
	path, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("failed to make temporary directory: %s", err)
	}
	defer os.RemoveAll(path)

	_, err = os.Create(filepath.Join(path, testExecutable))
	if err != nil {
		t.Fatalf("failed to create test executable: %s", err)
	}

	nestedPath := filepath.Join(path, executablePath)
	if err := os.MkdirAll(nestedPath, 0700); err != nil {
		t.Fatalf("failed to create nested directory structure: %s", err)
	}

	nestedExe := filepath.Join(nestedPath, testExecutable)
	_, err = os.Create(nestedExe)
	if err != nil {
		t.Fatalf("failed to create test executable: %s", err)
	}

	if testPath, err := findExecutable(testExecutable, path); err != nil {
		t.Fatalf("failed to find executable: %s", err)
	} else {
		if testPath != filepath.Join(path, testExecutable) {
			t.Fatalf("test executable location does not match, expected `%s` and received `%s`", filepath.Join(path, testExecutable), testPath)
		}
	}

	if testPath, err := findExecutable(nestedExe, path); err != nil {
		t.Fatalf("failed to find executable: %s", err)
	} else {
		if testPath != nestedExe {
			t.Fatalf("test executable location does not match, expected `%s` and received `%s`", nestedExe, testPath)
		}
	}
}
