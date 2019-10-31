# \MetadataServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataset**](MetadataServiceApi.md#GetDataset) | **Get** /datasets/{datasetId} | Gets a &#x60;Dataset&#x60; by ID.
[**SearchDatasets**](MetadataServiceApi.md#SearchDatasets) | **Post** /datasets/search | Gets a list of &#x60;Dataset&#x60; matching the search criteria.



## GetDataset

> Ga4ghDataset GetDataset(ctx, datasetId)
Gets a `Dataset` by ID.

`GET /datasets/{datasetId}` will return a JSON version of `Dataset`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasetId** | **string**| The ID of the &#x60;Dataset&#x60; to be retrieved. | 

### Return type

[**Ga4ghDataset**](ga4ghDataset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDatasets

> Ga4ghSearchDatasetsResponse SearchDatasets(ctx, body)
Gets a list of `Dataset` matching the search criteria.

`POST /datasets/search` must accept a JSON version of `SearchDatasetsRequest` as the post body and will return a JSON version of `SearchDatasetsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchDatasetsRequest**](Ga4ghSearchDatasetsRequest.md)|  | 

### Return type

[**Ga4ghSearchDatasetsResponse**](ga4ghSearchDatasetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

