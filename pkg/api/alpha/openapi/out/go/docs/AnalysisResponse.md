# AnalysisResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Mythril API version at the moment of analysis execution; or at the moment of submission, if this analysis is still queued.  | [optional] 
**MythrilVersion** | **string** | Mythril core version at the moment of analysis execution; or at the moment of submission, if this analysis is still queued.  | [optional] 
**QueueTime** | **int64** | The time [ms] from analysis submission to its execution start, or to the present moment, if this analysis is still in the queue.  | [optional] 
**RunTime** | **int64** | The time [ms] from the start of analysis execution till its end. Equals zero, if this analysis is still in the queue.  | [optional] 
**Status** | **string** | Current status of the analysis. | [optional] 
**SubmittedAt** | [**time.Time**](time.Time.md) | Timestamp of the analysis submission to the API. | [optional] 
**SubmittedBy** | [**time.Time**](time.Time.md) | ID of the submitter. | [optional] 
**Uuid** | **string** | Unique identifier of the analysis. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


