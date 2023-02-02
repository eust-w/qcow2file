package src

import (
	"errors"
	"os"
	"testing"
)

//func TestCreateQcowFromBase(t *testing.T) {
//	base := "base.qcow2"
//	out := "out.qcow2"
//	expectedPath, _ := filepath.Abs(out)
//
//	// Test with valid base file
//	path, err := createQcowFromBase(base, out)
//	if err != nil {
//		t.Fatalf("Unexpected error: %s", err)
//	}
//	if path != expectedPath {
//		t.Fatalf("Unexpected path. Got %s, expected %s", path, expectedPath)
//	}
//
//	// Test with non-existent base file
//	path, err = createQcowFromBase("nonexistent.qcow2", out)
//	if err == nil {
//		t.Fatalf("Expected error, got nil")
//	}
//
//	// Clean up created file
//	err = os.Remove(out)
//	if err != nil {
//		t.Fatalf("Failed to remove created file: %s", out)
//	}
//}

func TestCheckPath(t *testing.T) {
	file := "file.txt"
	dir := "dir"

	// Create test file
	f, err := os.Create(file)
	if err != nil {
		t.Fatalf("Failed to create test file: %s", file)
	}
	f.Close()

	// Test with valid file
	err = checkPath(file)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// Test with valid directory
	err = os.Mkdir(dir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %s", dir)
	}
	err = checkPath(dir)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	expectedErr := errors.New("path is a dir")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Unexpected error message. Got %s, expected %s", err, expectedErr)
	}

	// Test with non-existent path
	err = checkPath("nonexistent")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	expectedErr = errors.New("file not exist")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Unexpected error message. Got %s, expected %s", err, expectedErr)
	}

	// Clean up created files/directories
	err = os.Remove(file)
	if err != nil {
		t.Fatalf("Failed to remove created file: %s", file)
	}
	err = os.Remove(dir)
	if err != nil {
		t.Fatalf("Failed to remove created directory: %s", dir)
	}
}
