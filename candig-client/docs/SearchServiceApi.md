# \SearchServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCount**](SearchServiceApi.md#GetCount) | **Post** /count | 
[**GetItem**](SearchServiceApi.md#GetItem) | **Post** /search | 



## GetCount

> Ga4ghSearchQueryResponse GetCount(ctx, body)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchQueryRequest**](Ga4ghSearchQueryRequest.md)|  | 

### Return type

[**Ga4ghSearchQueryResponse**](ga4ghSearchQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> Ga4ghSearchQueryResponse GetItem(ctx, body)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchQueryRequest**](Ga4ghSearchQueryRequest.md)|  | 

### Return type

[**Ga4ghSearchQueryResponse**](ga4ghSearchQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

