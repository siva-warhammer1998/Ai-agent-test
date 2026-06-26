package config

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

const (
	defaultBucketName = "test-s3-bucket-name"
	defaultAWSRegion  = "ca-central-1"
)

var (
	bucketNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`)
	awsRegionPattern  = regexp.MustCompile(`^[a-z]{2}-[a-z]+-[0-9]$`)
)

type Config struct {
	BucketName string
	AWSRegion  string
}

func Load() (Config, error) {
	cfg := Config{
		BucketName: getenvDefault("S3_BUCKET_NAME", defaultBucketName),
		AWSRegion:  getenvDefault("AWS_REGION", defaultAWSRegion),
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg Config) Validate() error {
	if err := validateBucketName(cfg.BucketName); err != nil {
		return fmt.Errorf("bucket name: %w", err)
	}

	if err := validateAWSRegion(cfg.AWSRegion); err != nil {
		return fmt.Errorf("aws region: %w", err)
	}

	return nil
}

func validateBucketName(value string) error {
	if value == "" {
		return errors.New("required")
	}

	if len(value) < 3 || len(value) > 63 {
		return errors.New("must be between 3 and 63 characters")
	}

	if !bucketNamePattern.MatchString(value) {
		return errors.New("must contain only lowercase letters, numbers, dots, and hyphens, and must start and end with a letter or number")
	}

	if regexp.MustCompile(`\.\.`).MatchString(value) {
		return errors.New("must not contain adjacent periods")
	}

	if regexp.MustCompile(`\.-|-\.`).MatchString(value) {
		return errors.New("must not contain periods next to hyphens")
	}

	return nil
}

func validateAWSRegion(value string) error {
	if value == "" {
		return errors.New("required")
	}

	if !awsRegionPattern.MatchString(value) {
		return errors.New("must look like a valid AWS region, for example ca-central-1")
	}

	return nil
}

func getenvDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
