package handler

import (
	"encoding/json"
	"net/http"

	"github.com/siva-warhammer1998/Ai-agent-test/internal/config"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/domain"
	"github.com/siva-warhammer1998/Ai-agent-test/internal/service"
)

type BucketHandler struct {
	cfg     config.Config
	planner service.BucketPlanner
}

func NewBucketHandler(cfg config.Config, planner service.BucketPlanner) BucketHandler {
	return BucketHandler{
		cfg:     cfg,
		planner: planner,
	}
}

func (handler BucketHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /bucket-plan", handler.BucketPlan)
}

func (handler BucketHandler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (handler BucketHandler) BucketPlan(w http.ResponseWriter, _ *http.Request) {
	plan, err := handler.planner.PlanCreateBucket(domain.BucketRequest{
		Name:   handler.cfg.BucketName,
		Region: handler.cfg.AWSRegion,
	})
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, plan)
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
