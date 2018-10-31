# \AnalysisApi

All URIs are relative to *https://api.mythril.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalysis**](AnalysisApi.md#GetAnalysis) | **Get** /analyses/{uuid} | Analysis by UUID
[**GetAnalysisIssues**](AnalysisApi.md#GetAnalysisIssues) | **Get** /analyses/{uuid}/issues | Detected issues
[**ListAnalyses**](AnalysisApi.md#ListAnalyses) | **Get** /analyses | List of analyses
[**SubmitAnalysis**](AnalysisApi.md#SubmitAnalysis) | **Post** /analyses | New analysis


# **GetAnalysis**
> AnalysisResponse GetAnalysis(ctx, uuid)
Analysis by UUID

Gets status and metadata of the analysis specified by UUID. When analysis **status** is _Finished_, [`GET /analyses/{uuid}/issues`](#operation/getAnalysisIssues) may be used to list detected vulnerabilities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **int32**| analysis id | 

### Return type

[**AnalysisResponse**](AnalysisResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalysisIssues**
> AnalysisIssuesResponse GetAnalysisIssues(ctx, uuid)
Detected issues

Lists issues detected during the analysis specified by UUID. Request will fail for unfinished analyses, use [`GET /analyses/{uuid}](#operation/getAnalysis) to verify the current analysis status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **int32**| analysis id | 

### Return type

[**AnalysisIssuesResponse**](AnalysisIssuesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAnalyses**
> []AnalysisResponse ListAnalyses(ctx, optional)
List of analyses

Lists analyses visible to the user, at most 20 records a time, sorted by submission time from the most recent to the oldest ones. Use `offset` parameter for results pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAnalysesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAnalysesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| Pagination offset. Number of records to skip. | 
 **dateFrom** | **optional.Time**| Submission time filter. Restricts results to analyses submitted after this time. | 
 **dateTo** | **optional.Time**| Submission time filter. Restricts results to analyses submitted before this time. | 

### Return type

[**[]AnalysisResponse**](AnalysisResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitAnalysis**
> AnalysisResponse SubmitAnalysis(ctx, uNKNOWNBASETYPE)
New analysis

Submits Ethereum contract(s) for vulnerability analysis, and returns created analysis record. **uuid** field of the response should be used in subsequent calls to [`GET /analysis/{uuid}`](#operation/getAnalysis) and [`GET /analysis/{uuid}/issues`](#operation/getAnalysisIssues) to check analysis status, metadata, and the list of detected issues.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**AnalysisResponse**](AnalysisResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

