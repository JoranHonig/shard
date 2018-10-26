package mythril

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
)

type statusResponse struct {
	apiVersion     string
	mythrilVersion string
	queueTime      int
	runTime        int
	status         string
	submittedAt    string
	submittedBy    string
	uuid           string
}

type analysisRequest struct {
	analysisType string `json:"type"`
	contractBody string `json:"contract"`
}

type ApiV1Client struct {
	httpClient *http.Client
	BaseURL    *url.URL
	apiKey     string
}

func BuildApiV1Client(apiKey string, baseUrl string) ApiV1Client {
	client := http.Client{}
	url, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return ApiV1Client{&client, url , apiKey}
}

func (api *ApiV1Client) Submit(bytecode string) (*uuid.UUID, error) {
	request := analysisRequest{analysisType: "full", contractBody: bytecode}
	req, err := api.newRequest("POST", "/v1/analyses", request)

	if err != nil {
		return nil, err
	}

	resp := statusResponse{}
	_, err = api.do(req, &resp)
	if err != nil {
		return nil, err
	}

	_uuid, err := uuid.Parse(resp.uuid)

	if err != nil {
		return nil, err
	}

	return &_uuid, nil
}

func (api *ApiV1Client) CheckStatus(_uuid uuid.UUID) (*analysisJobStatus, error) {
	req, err := api.newRequest("GET", fmt.Sprintf("/v1/analyses/%s/issues", _uuid.String()), nil)
	if err != nil {
		return nil, err
	}

	resp := statusResponse{}
	_, err = api.do(req, &resp)
	if err != nil {
		return nil, err
	}

	result := &analysisJobStatus{uuid: _uuid, status: resp.status}
	return result, nil
}

func (c *ApiV1Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	relativePath := &url.URL{Path: path}
	targetUrl := c.BaseURL.ResolveReference(relativePath)
	fmt.Println("UUUH")
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, targetUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	return req, nil
}

func (c *ApiV1Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf(string(b))
		err = json.NewDecoder(resp.Body).Decode(v)
	} else {
		return nil, errors.New("Unexpected Status Code")
	}

	return resp, err
}