package repository

import (
	"context"
	"fmt"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
)

type BucketCreator interface {
	CreateBucket(ctx context.Context, plan domain.BucketPlan) (domain.BucketCreationResult, error)
}

type DryRunBucketRepository struct{}

func NewDryRunBucketRepository() DryRunBucketRepository {
	return DryRunBucketRepository{}
}

func (repository DryRunBucketRepository) CreateBucket(ctx context.Context, plan domain.BucketPlan) (domain.BucketCreationResult, error) {
	if err := ctx.Err(); err != nil {
		return domain.BucketCreationResult{}, fmt.Errorf("create bucket canceled: %w", err)
	}

	if plan.Name == "" {
		return domain.BucketCreationResult{}, fmt.Errorf("bucket name is required")
	}

	if plan.Region == "" {
		return domain.BucketCreationResult{}, fmt.Errorf("aws region is required")
	}

	return domain.BucketCreationResult{
		Name:    plan.Name,
		Region:  plan.Region,
		Created: false,
		Message: "dry run only; no AWS resources were created or modified",
	}, nil
}
