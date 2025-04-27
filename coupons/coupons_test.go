package coupons

import (
	"compress/gzip"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func createTempGzipFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "test_coupon_*.gz")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	gw := gzip.NewWriter(tmpFile)
	_, err = gw.Write([]byte(content))
	if err != nil {
		t.Fatalf("Failed to write gzip content: %v", err)
	}
	gw.Close()
	tmpFile.Close()

	return tmpFile.Name()
}

func TestLoadCoupons_Success(t *testing.T) {
	expectedContent := "This is a test coupon file."

	tempFile := createTempGzipFile(t, expectedContent)
	defer os.Remove(tempFile)

	loadedContent := LoadCoupons(tempFile)

	if strings.TrimSpace(loadedContent) != expectedContent {
		t.Errorf("Expected content %q, but got %q", expectedContent, loadedContent)
	}
}

func TestLoadCoupons_FileNotFound(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		LoadCoupons("non_existing_file.gz")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLoadCoupons_FileNotFound")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	} else {
		t.Fatalf("Process should have exited with error for missing file!")
	}
}

func TestLoadCoupons_InvalidGzip(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "2" {
		tmpFile, _ := os.CreateTemp("", "invalid_gzip_*.gz")
		defer os.Remove(tmpFile.Name())

		tmpFile.WriteString("this is not a valid gzip content")
		tmpFile.Close()

		LoadCoupons(tmpFile.Name())
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLoadCoupons_InvalidGzip")
	cmd.Env = append(os.Environ(), "BE_CRASHER=2")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	} else {
		t.Fatalf("Process should have exited with error for invalid gzip!")
	}
}
