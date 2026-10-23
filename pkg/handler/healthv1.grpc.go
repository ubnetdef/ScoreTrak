package handler

import (
	healthv1 "github.com/ubnetdef/scoretrakapis/gen/proto/go/grpc/health/v1"
	"context"
)

type HealthV1Controller struct {
	healthv1.UnimplementedHealthServer
}

func NewHealthV1Controller() *HealthV1Controller {
	return &HealthV1Controller{}
}

func (h HealthV1Controller) Check(ctx context.Context, request *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{Status: healthv1.HealthCheckResponse_SERVING}, nil
}
