package generic

import (
	"github.com/google/uuid"
)

type MythrilService interface {
	Submit(bytecode string) (*uuid.UUID, error)
	CheckStatus(_uuid uuid.UUID) (*AnalysisJobStatus, error)
	GetIssueResult(_uuid uuid.UUID) ([]Issue, error)
}

type Issue struct {
	title string
	description string
}

type AnalysisJobStatus struct {
	Uuid uuid.UUID
	Status string
}



