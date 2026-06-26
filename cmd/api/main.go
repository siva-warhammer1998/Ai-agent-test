package main

import (
	"fmt"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/config"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("configuration error: %v\n", err)
		return
	}

	planner := service.NewBucketPlanner()
	plan, err := planner.PlanCreateBucket(domain.BucketRequest{
		Name:   cfg.BucketName,
		Region: cfg.AWSRegion,
	})
	if err != nil {
		fmt.Printf("planning error: %v\n", err)
		return
	}

	fmt.Println("AWS S3 bucket agent")
	fmt.Printf("operation: %s\n", plan.Operation)
	fmt.Printf("target bucket: %s\n", plan.Name)
	fmt.Printf("target region: %s\n", plan.Region)
	fmt.Printf("dry run: %t\n", plan.DryRun)
	fmt.Println("status: dry run only; no AWS resources were created or modified")
}
