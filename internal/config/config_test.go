package config

import "testing"

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("S3_BUCKET_NAME", "")
	t.Setenv("AWS_REGION", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.BucketName != defaultBucketName {
		t.Fatalf("BucketName = %q, want %q", cfg.BucketName, defaultBucketName)
	}

	if cfg.AWSRegion != defaultAWSRegion {
		t.Fatalf("AWSRegion = %q, want %q", cfg.AWSRegion, defaultAWSRegion)
	}
}

func TestLoadUsesEnvironment(t *testing.T) {
	t.Setenv("S3_BUCKET_NAME", "agent-test-bucket")
	t.Setenv("AWS_REGION", "us-east-1")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.BucketName != "agent-test-bucket" {
		t.Fatalf("BucketName = %q", cfg.BucketName)
	}

	if cfg.AWSRegion != "us-east-1" {
		t.Fatalf("AWSRegion = %q", cfg.AWSRegion)
	}
}

func TestConfigValidateRejectsInvalidBucket(t *testing.T) {
	cfg := Config{
		BucketName: "Invalid_Bucket",
		AWSRegion:  defaultAWSRegion,
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}

func TestConfigValidateRejectsInvalidRegion(t *testing.T) {
	cfg := Config{
		BucketName: defaultBucketName,
		AWSRegion:  "canada",
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}
