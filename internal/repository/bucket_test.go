package repository

import (
	"context"
	"testing"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
)

func TestDryRunBucketRepositoryCreateBucket(t *testing.T) {
	repository := NewDryRunBucketRepository()

	result, err := repository.CreateBucket(context.Background(), domain.BucketPlan{
		Name:      "test-s3-bucket-name",
		Region:    "ca-central-1",
		Operation: "create_s3_bucket",
		DryRun:    true,
	})
	if err != nil {
		t.Fatalf("CreateBucket() error = %v", err)
	}

	if result.Name != "test-s3-bucket-name" {
		t.Fatalf("Name = %q", result.Name)
	}

	if result.Region != "ca-central-1" {
		t.Fatalf("Region = %q", result.Region)
	}

	if result.Created {
		t.Fatal("Created = true, want false")
	}
}

func TestDryRunBucketRepositoryHonorsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	repository := NewDryRunBucketRepository()

	_, err := repository.CreateBucket(ctx, domain.BucketPlan{
		Name:   "test-s3-bucket-name",
		Region: "ca-central-1",
	})
	if err == nil {
		t.Fatal("CreateBucket() error = nil, want error")
	}
}
