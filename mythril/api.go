package mythril

import (
	"shard/mythril/alpha"
	"github.com/google/uuid"
	"errors"
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
	uuid uuid.UUID
	status string
}

type MythrilServiceType int

const (
	ALPHA MythrilServiceType = 1 << iota
	V1
	V2
)

func BuildMythrilService(version MythrilServiceType, apiKey string) (MythrilService, error){
	switch version {
	case ALPHA:
		return alpha.BuildMythrilServiceALPHA(apiKey), nil
	default:
		errors.New("Invalid MythrilServiceType")
	}
}