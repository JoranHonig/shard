package mythril

import (
	"fmt"
	"github.com/google/uuid"
	"armlet-go/openapi/out/go"
	"context"
)

import (
	log "github.com/sirupsen/logrus"
)

type MythrilService struct {
	openApiClient *openapi.APIClient
	apiKey string
}

func BuildMythrilService(apiKey string) *MythrilService {
	log.Debug("Building API cl")

	c := openapi.NewConfiguration()
	client := openapi.NewAPIClient(c)

	return &MythrilService{client, apiKey}
}

func (api *MythrilService) Submit(bytecode string) (*uuid.UUID, error) {
	ctx := context.Background();

	ctx = context.WithValue(ctx, "ApiKey", api.apiKey)

	options := openapi.SubmitAnalysisOpts{
		"full",
		bytecode,
		nil,
	}
	analysisApi := api.openApiClient.AnalysisApi
	response, http_response, error := analysisApi.SubmitAnalysis(ctx, options);

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(response)
	fmt.Println(http_response)
	uuid, err := uuid.Parse(response.Uuid)
	if err != nil {
		log.Fatal(err)
	}

	return &uuid, nil
}

func (api *MythrilService) CheckStatus(_uuid *uuid.UUID) (*analysisJobStatus, error) {
	return nil, nil
}

func (api *MythrilService) GetIssues(_uuid *uuid.UUID) () {
	ctx := context.Background();

	ctx = context.WithValue(ctx, "ApiKey", api.apiKey)

	analysisApi := api.openApiClient.AnalysisApi
	response, http_response, error := analysisApi.GetAnalysisIssues(ctx, _uuid.String());

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(response)
	fmt.Println(http_response)
	uuid, err := uuid.Parse(response.Uuid)
	if err != nil {
		log.Fatal(err)
	}

	return &uuid, nil
}