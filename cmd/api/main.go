package main

import (
	"fmt"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("configuration error: %v\n", err)
		return
	}

	fmt.Println("AWS S3 bucket agent")
	fmt.Printf("target bucket: %s\n", cfg.BucketName)
	fmt.Printf("target region: %s\n", cfg.AWSRegion)
	fmt.Println("status: dry run only; no AWS resources were created or modified")
}
