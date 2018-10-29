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
	Title string
	Description string
	Function string
	Type string
	Address string
	Debug string
}

type AnalysisJobStatus struct {
	Uuid uuid.UUID
	Status string
}



