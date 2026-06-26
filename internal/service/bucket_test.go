package service

import (
	"context"
	"testing"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/repository"
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

func TestBucketServiceCreateBucketUsesDryRunRepository(t *testing.T) {
	planner := NewBucketPlanner()
	plan, err := planner.PlanCreateBucket(domain.BucketRequest{
		Name:   "test-s3-bucket-name",
		Region: "ca-central-1",
	})
	if err != nil {
		t.Fatalf("PlanCreateBucket() error = %v", err)
	}

	service := NewBucketService(repository.NewDryRunBucketRepository())
	result, err := service.CreateBucket(context.Background(), plan)
	if err != nil {
		t.Fatalf("CreateBucket() error = %v", err)
	}

	if result.Created {
		t.Fatal("Created = true, want false")
	}
}

func TestBucketServiceRejectsNonDryRunPlan(t *testing.T) {
	service := NewBucketService(repository.NewDryRunBucketRepository())

	_, err := service.CreateBucket(context.Background(), domain.BucketPlan{
		Name:   "test-s3-bucket-name",
		Region: "ca-central-1",
		DryRun: false,
	})
	if err == nil {
		t.Fatal("CreateBucket() error = nil, want error")
	}
}
