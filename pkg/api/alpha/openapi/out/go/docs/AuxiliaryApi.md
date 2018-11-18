# \AuxiliaryApi

All URIs are relative to *https://api.mythril.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenApiYaml**](AuxiliaryApi.md#GetOpenApiYaml) | **Get** /openapi.yaml | OpenAPI specification
[**GetVersion**](AuxiliaryApi.md#GetVersion) | **Get** /version | API version


# **GetOpenApiYaml**
> Yaml GetOpenApiYaml(ctx, )
OpenAPI specification

Gets OpenAPI specification of Mythril API.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Yaml**](YAML.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersion**
> InlineResponse200 GetVersion(ctx, )
API version

Gets current versions of Mythril API and its core sub-modules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

