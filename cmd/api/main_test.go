package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunPrintsDryRunPlan(t *testing.T) {
	t.Setenv("S3_BUCKET_NAME", "test-s3-bucket-name")
	t.Setenv("AWS_REGION", "ca-central-1")

	var output bytes.Buffer
	run(&output)

	wantLines := []string{
		"AWS S3 bucket agent",
		"operation: create_s3_bucket",
		"target bucket: test-s3-bucket-name",
		"target region: ca-central-1",
		"dry run: true",
		"status: dry run only; no AWS resources were created or modified",
	}

	for _, line := range wantLines {
		if !strings.Contains(output.String(), line) {
			t.Fatalf("output missing %q:\n%s", line, output.String())
		}
	}
}

func TestRunPrintsConfigurationError(t *testing.T) {
	t.Setenv("S3_BUCKET_NAME", "Invalid_Bucket")
	t.Setenv("AWS_REGION", "ca-central-1")

	var output bytes.Buffer
	run(&output)

	if !strings.Contains(output.String(), "configuration error:") {
		t.Fatalf("output = %q, want configuration error", output.String())
	}
}
