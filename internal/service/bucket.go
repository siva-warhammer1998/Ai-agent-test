package service

import (
	"fmt"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
)

const createBucketOperation = "create_s3_bucket"

type BucketPlanner struct{}

func NewBucketPlanner() BucketPlanner {
	return BucketPlanner{}
}

func (planner BucketPlanner) PlanCreateBucket(request domain.BucketRequest) (domain.BucketPlan, error) {
	if request.Name == "" {
		return domain.BucketPlan{}, fmt.Errorf("bucket name is required")
	}

	if request.Region == "" {
		return domain.BucketPlan{}, fmt.Errorf("aws region is required")
	}

	return domain.BucketPlan{
		Name:      request.Name,
		Region:    request.Region,
		Operation: createBucketOperation,
		DryRun:    true,
	}, nil
}
