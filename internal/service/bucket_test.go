package service

import (
	"testing"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
)

func TestPlanCreateBucket(t *testing.T) {
	planner := NewBucketPlanner()

	plan, err := planner.PlanCreateBucket(domain.BucketRequest{
		Name:   "test-s3-bucket-name",
		Region: "ca-central-1",
	})
	if err != nil {
		t.Fatalf("PlanCreateBucket() error = %v", err)
	}

	if plan.Name != "test-s3-bucket-name" {
		t.Fatalf("Name = %q", plan.Name)
	}

	if plan.Region != "ca-central-1" {
		t.Fatalf("Region = %q", plan.Region)
	}

	if plan.Operation != createBucketOperation {
		t.Fatalf("Operation = %q", plan.Operation)
	}

	if !plan.DryRun {
		t.Fatal("DryRun = false, want true")
	}
}

func TestPlanCreateBucketRequiresName(t *testing.T) {
	planner := NewBucketPlanner()

	_, err := planner.PlanCreateBucket(domain.BucketRequest{
		Region: "ca-central-1",
	})
	if err == nil {
		t.Fatal("PlanCreateBucket() error = nil, want error")
	}
}

func TestPlanCreateBucketRequiresRegion(t *testing.T) {
	planner := NewBucketPlanner()

	_, err := planner.PlanCreateBucket(domain.BucketRequest{
		Name: "test-s3-bucket-name",
	})
	if err == nil {
		t.Fatal("PlanCreateBucket() error = nil, want error")
	}
}
