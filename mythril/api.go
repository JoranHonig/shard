package mythril

import "github.com/google/uuid"

type api interface {
	Submit(bytecode string) (*uuid.UUID, error)
	CheckStatus(uuid uuid.UUID) (*analysisJobStatus, error)
}

type analysisJobStatus struct {
	uuid uuid.UUID
	status string
}

func buildAPI(version int) api {
	return nil
}