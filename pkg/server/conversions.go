package server

import (
	"github.com/opengovern/og-agent/pkg/database"
	"github.com/opengovern/og-agent/pkg/proto/src/golang"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dbOptimizationJobToApiOptimizationJob(job *database.OptimizationJob) *golang.OptimizationJob {
	return &golang.OptimizationJob{
		Id:           uint64(job.ID),
		Command:      job.Command,
		Status:       string(job.Status),
		ErrorMessage: job.ErrorMessage,
		CreatedAt:    timestamppb.New(job.CreatedAt),
		UpdatedAt:    timestamppb.New(job.UpdatedAt),
	}
}
