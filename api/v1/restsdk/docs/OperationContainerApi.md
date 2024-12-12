# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationContainerCreateOperationStatus**](OperationContainerApi.md#OperationContainerCreateOperationStatus) | **Post** /v1/operation/{operationId} | Create a new Operation Status
[**OperationContainerGetOperationStatus**](OperationContainerApi.md#OperationContainerGetOperationStatus) | **Get** /v1/operation/{operationId} | Get Operation Status
[**OperationContainerUpdateOperationStatus**](OperationContainerApi.md#OperationContainerUpdateOperationStatus) | **Put** /v1/operation/{operationId} | Update Operation Status

# **OperationContainerCreateOperationStatus**
> interface{} OperationContainerCreateOperationStatus(ctx, body, operationId)
Create a new Operation Status

This operation creates a new operation status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OperationContainerCreateOperationStatusBody**](OperationContainerCreateOperationStatusBody.md)|  | 
  **operationId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationContainerGetOperationStatus**
> OperationContainerGetOperationStatusResponse OperationContainerGetOperationStatus(ctx, operationId)
Get Operation Status

This operation gets the operation status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operationId** | **string**|  | 

### Return type

[**OperationContainerGetOperationStatusResponse**](OperationContainerGetOperationStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationContainerUpdateOperationStatus**
> interface{} OperationContainerUpdateOperationStatus(ctx, body, operationId)
Update Operation Status

This operation updates the operation status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OperationContainerUpdateOperationStatusBody**](OperationContainerUpdateOperationStatusBody.md)|  | 
  **operationId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

