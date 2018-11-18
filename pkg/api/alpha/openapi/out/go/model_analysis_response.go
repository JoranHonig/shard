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
	"time"
)

type AnalysisResponse struct {
	// Mythril API version at the moment of analysis execution; or at the moment of submission, if this analysis is still queued.
	ApiVersion string `json:"apiVersion,omitempty"`
	// Mythril core version at the moment of analysis execution; or at the moment of submission, if this analysis is still queued.
	MythrilVersion string `json:"mythrilVersion,omitempty"`
	// The time [ms] from analysis submission to its execution start, or to the present moment, if this analysis is still in the queue.
	QueueTime int64 `json:"queueTime,omitempty"`
	// The time [ms] from the start of analysis execution till its end. Equals zero, if this analysis is still in the queue.
	RunTime int64 `json:"runTime,omitempty"`
	// Current status of the analysis.
	Status string `json:"status,omitempty"`
	// Timestamp of the analysis submission to the API.
	SubmittedAt time.Time `json:"submittedAt,omitempty"`
	// ID of the submitter.
	SubmittedBy string `json:"submittedBy,omitempty"`
	// Unique identifier of the analysis.
	Uuid string `json:"uuid,omitempty"`
}