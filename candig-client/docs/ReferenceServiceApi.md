# \ReferenceServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReference**](ReferenceServiceApi.md#GetReference) | **Get** /references/{reference_id} | Gets a &#x60;Reference&#x60; by ID.
[**GetReferenceSet**](ReferenceServiceApi.md#GetReferenceSet) | **Get** /referencesets/{reference_set_id} | Gets a &#x60;ReferenceSet&#x60; by ID.
[**ListReferenceBases**](ReferenceServiceApi.md#ListReferenceBases) | **Post** /references/{reference_id}/bases | Lists &#x60;Reference&#x60; bases by ID and optional range.
[**SearchReferenceSets**](ReferenceServiceApi.md#SearchReferenceSets) | **Post** /referencesets/search | Gets a list of &#x60;ReferenceSet&#x60; matching the search criteria.
[**SearchReferences**](ReferenceServiceApi.md#SearchReferences) | **Post** /references/search | Gets a list of &#x60;Reference&#x60; matching the search criteria.



## GetReference

> Ga4ghReference GetReference(ctx, referenceId)
Gets a `Reference` by ID.

`GET /references/{reference_id}` will return a JSON version of `Reference`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceId** | **string**| The ID of the &#x60;Reference&#x60; to be retrieved. | 

### Return type

[**Ga4ghReference**](ga4ghReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferenceSet

> Ga4ghReferenceSet GetReferenceSet(ctx, referenceSetId)
Gets a `ReferenceSet` by ID.

`GET /referencesets/{reference_set_id}` will return a JSON version of `ReferenceSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceSetId** | **string**| The ID of the &#x60;ReferenceSet&#x60; to be retrieved. | 

### Return type

[**Ga4ghReferenceSet**](ga4ghReferenceSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReferenceBases

> Ga4ghListReferenceBasesResponse ListReferenceBases(ctx, referenceId, body)
Lists `Reference` bases by ID and optional range.

`POST /listreferencebases` will return a JSON version of `ListReferenceBasesResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceId** | **string**| The ID of the &#x60;Reference&#x60; to be retrieved. | 
**body** | [**Ga4ghListReferenceBasesRequest**](Ga4ghListReferenceBasesRequest.md)|  | 

### Return type

[**Ga4ghListReferenceBasesResponse**](ga4ghListReferenceBasesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReferenceSets

> Ga4ghSearchReferenceSetsResponse SearchReferenceSets(ctx, body)
Gets a list of `ReferenceSet` matching the search criteria.

`POST /referencesets/search` must accept a JSON version of `SearchReferenceSetsRequest` as the post body and will return a JSON version of `SearchReferenceSetsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchReferenceSetsRequest**](Ga4ghSearchReferenceSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchReferenceSetsResponse**](ga4ghSearchReferenceSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReferences

> Ga4ghSearchReferencesResponse SearchReferences(ctx, body)
Gets a list of `Reference` matching the search criteria.

`POST /references/search` must accept a JSON version of `SearchReferencesRequest` as the post body and will return a JSON version of `SearchReferencesResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchReferencesRequest**](Ga4ghSearchReferencesRequest.md)|  | 

### Return type

[**Ga4ghSearchReferencesResponse**](ga4ghSearchReferencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

