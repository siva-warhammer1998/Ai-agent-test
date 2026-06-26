package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/config"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/service"
)

func TestHealth(t *testing.T) {
	handler := NewBucketHandler(config.Config{}, service.NewBucketPlanner())
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	handler.Health(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}

	var body map[string]string
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("status body = %q", body["status"])
	}
}

func TestBucketPlan(t *testing.T) {
	handler := NewBucketHandler(config.Config{
		BucketName: "test-s3-bucket-name",
		AWSRegion:  "ca-central-1",
	}, service.NewBucketPlanner())
	request := httptest.NewRequest(http.MethodGet, "/bucket-plan", nil)
	response := httptest.NewRecorder()

	handler.BucketPlan(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}

	var plan domain.BucketPlan
	if err := json.NewDecoder(response.Body).Decode(&plan); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if plan.Name != "test-s3-bucket-name" {
		t.Fatalf("Name = %q", plan.Name)
	}

	if plan.Region != "ca-central-1" {
		t.Fatalf("Region = %q", plan.Region)
	}

	if !plan.DryRun {
		t.Fatal("DryRun = false, want true")
	}
}
