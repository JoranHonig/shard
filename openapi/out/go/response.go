/*
 * Mythril API
 *
 * The Mythril service is currently in closed alpha. API keys are made available to selected community members (join the Mythril Discord to request one).  **Note that the interface will still undergo significant changes.**  Mythril is a security analysis tool for Ethereum smart contracts. It uses concolic analysis, taint analysis, and control flow checking to detect a variety of security vulnerabilities.  Mythril API exposes this functionality over the network.  In the future this API will integrate a number of other security analyzer tools seemlessly. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}