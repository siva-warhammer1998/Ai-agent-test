package service

import (
	"context"
	"fmt"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/repository"
)

const createBucketOperation = "create_s3_bucket"

type BucketPlanner struct{}

func NewBucketPlanner() BucketPlanner {
	return BucketPlanner{}
}

type BucketService struct {
	creator repository.BucketCreator
}

func NewBucketService(creator repository.BucketCreator) BucketService {
	return BucketService{
		creator: creator,
	}
}

func (service BucketService) CreateBucket(ctx context.Context, plan domain.BucketPlan) (domain.BucketCreationResult, error) {
	if !plan.DryRun {
		return domain.BucketCreationResult{}, fmt.Errorf("non-dry-run bucket creation requires explicit approval")
	}

	result, err := service.creator.CreateBucket(ctx, plan)
	if err != nil {
		return domain.BucketCreationResult{}, fmt.Errorf("create bucket: %w", err)
	}

	return result, nil
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
