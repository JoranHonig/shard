package generic

import (
	"github.com/JoranHonig/shard/common"
	"github.com/google/uuid"
)

type MythrilService interface {
	Submit(bytecode string) (*uuid.UUID, error)
	CheckStatus(_uuid uuid.UUID) (*AnalysisJobStatus, error)
	GetIssueResult(_uuid uuid.UUID) ([]common.Issue, error)
}

type AnalysisJobStatus struct {
	Uuid   uuid.UUID
	Status string
}
