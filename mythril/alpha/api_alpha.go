package alpha

import (
	"context"
	"github.com/JoranHonig/shard/common"
	"github.com/JoranHonig/shard/mythril/generic"
	"github.com/JoranHonig/shard/openapi/out/go"
	"github.com/google/uuid"
)

func BuildMythrilServiceALPHA(apiKey string) generic.MythrilService {
	c := openapi.NewConfiguration()
	client := openapi.NewAPIClient(c)

	service := MythrilServiceALPHA{client, apiKey}
	return &service
}

type MythrilServiceALPHA struct {
	openApiClient *openapi.APIClient
	apiKey        string
}

func (api *MythrilServiceALPHA) Submit(bytecode string) (*uuid.UUID, error) {
	ctx := api.getAuthenticatedContext()

	options := openapi.SubmitAnalysisOpts{
		"bytecode",
		bytecode,
		nil,
	}
	analysisApi := api.openApiClient.AnalysisApi
	response, _, err := analysisApi.SubmitAnalysis(ctx, options)

	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(response.Uuid)

	return &uuid, err
}

func (api *MythrilServiceALPHA) CheckStatus(_uuid uuid.UUID) (*generic.AnalysisJobStatus, error) {
	ctx := api.getAuthenticatedContext()

	analysisApi := api.openApiClient.AnalysisApi
	response, _, err := analysisApi.GetAnalysis(ctx, _uuid.String())

	if err != nil {
		return nil, err
	}

	return &generic.AnalysisJobStatus{Uuid: _uuid, Status: response.Status}, nil
}

func (api *MythrilServiceALPHA) GetIssueResult(_uuid uuid.UUID) ([]common.Issue, error) {
	ctx := api.getAuthenticatedContext()

	analysisApi := api.openApiClient.AnalysisApi
	response, _, err := analysisApi.GetAnalysisIssues(ctx, _uuid.String())

	if err != nil {
		return nil, err
	}

	result := make([]common.Issue, 0)
	for _, issue := range response {
		result = append(result,
			common.Issue{
				Title:       issue.Title,
				Description: issue.Description,
				Debug:       issue.Debug,
				Address:     issue.Address,
				Type:        issue.Type,
				Function:    issue.Function,
			})
	}

	return result, nil
}

func (api *MythrilServiceALPHA) getAuthenticatedContext() context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, openapi.ContextAPIKey, openapi.APIKey{api.apiKey, "Bearer"})
}
