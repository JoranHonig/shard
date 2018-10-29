package alpha

import (
	"shard/openapi/out/go"
	"context"
	"github.com/google/uuid"
	"shard/mythril/generic"
	log "github.com/sirupsen/logrus"
)

func BuildMythrilServiceALPHA(apiKey string) generic.MythrilService {
	log.Debug("Building API cl")

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
	ctx := context.Background()

	ctx = context.WithValue(ctx, openapi.ContextAPIKey, openapi.APIKey{api.apiKey, "Bearer"})

	options := openapi.SubmitAnalysisOpts{
		"bytecode",
		bytecode,
		nil,
	}
	analysisApi := api.openApiClient.AnalysisApi
	response, _, error := analysisApi.SubmitAnalysis(ctx, options)

	if error != nil {
		log.Fatal(error)
	}

	uuid, err := uuid.Parse(response.Uuid)
	if err != nil {
		log.Fatal(err)
	}

	return &uuid, nil
}

func (api *MythrilServiceALPHA) CheckStatus(_uuid uuid.UUID) (*generic.AnalysisJobStatus, error) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, openapi.ContextAPIKey, openapi.APIKey{api.apiKey, "Bearer"})


	analysisApi := api.openApiClient.AnalysisApi
	response, _, error := analysisApi.GetAnalysis(ctx, _uuid.String())

	//log.Info(response)
	//log.Info(http_response)

	if error != nil {
		log.Fatal(error)
	}

	return &generic.AnalysisJobStatus{Uuid: _uuid, Status: response.Status}, nil

}

func (api *MythrilServiceALPHA) GetIssueResult(_uuid uuid.UUID) ([]generic.Issue, error) {
	return nil, nil
}
