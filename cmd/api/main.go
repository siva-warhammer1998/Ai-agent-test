package main

import (
	"fmt"
	"io"
	"os"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/config"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/service"
)

func main() {
	run(os.Stdout)
}

func run(writer io.Writer) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(writer, "configuration error: %v\n", err)
		return
	}

	planner := service.NewBucketPlanner()
	plan, err := planner.PlanCreateBucket(domain.BucketRequest{
		Name:   cfg.BucketName,
		Region: cfg.AWSRegion,
	})
	if err != nil {
		fmt.Fprintf(writer, "planning error: %v\n", err)
		return
	}

	fmt.Fprintln(writer, "AWS S3 bucket agent")
	fmt.Fprintf(writer, "operation: %s\n", plan.Operation)
	fmt.Fprintf(writer, "target bucket: %s\n", plan.Name)
	fmt.Fprintf(writer, "target region: %s\n", plan.Region)
	fmt.Fprintf(writer, "dry run: %t\n", plan.DryRun)
	fmt.Fprintln(writer, "status: dry run only; no AWS resources were created or modified")
}
