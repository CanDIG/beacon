# \VariantServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallSet**](VariantServiceApi.md#GetCallSet) | **Get** /callsets/{call_set_id} | Gets a &#x60;CallSet&#x60; by ID.
[**GetVariant**](VariantServiceApi.md#GetVariant) | **Post** /variants/{variant_id} | Gets a &#x60;Variant&#x60; by ID.
[**GetVariantSet**](VariantServiceApi.md#GetVariantSet) | **Get** /variantsets/{variant_set_id} | Gets a &#x60;VariantSet&#x60; by ID.
[**SearchCallSets**](VariantServiceApi.md#SearchCallSets) | **Post** /callsets/search | Gets a list of call sets matching the search criteria.
[**SearchVariantSets**](VariantServiceApi.md#SearchVariantSets) | **Post** /variantsets/search | Gets a list of &#x60;VariantSet&#x60; matching the search criteria.
[**SearchVariants**](VariantServiceApi.md#SearchVariants) | **Post** /variants/search | Gets a list of &#x60;Variant&#x60; matching the search criteria.
[**SearchVariantsByGeneName**](VariantServiceApi.md#SearchVariantsByGeneName) | **Post** /variantsbygenesearch | Gets a list of &#x60;Variants&#x60; matching the search criteria.



## GetCallSet

> Ga4ghCallSet GetCallSet(ctx, callSetId)
Gets a `CallSet` by ID.

`GET /callsets/{id}` will return a JSON version of `CallSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSetId** | **string**| The ID of the &#x60;CallSet&#x60; to be retrieved. | 

### Return type

[**Ga4ghCallSet**](ga4ghCallSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariant

> Ga4ghVariant GetVariant(ctx, variantId)
Gets a `Variant` by ID.

`GET /variants/{id}` will return a JSON version of `Variant`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantId** | **string**| The ID of the &#x60;Variant&#x60; to be retrieved. | 

### Return type

[**Ga4ghVariant**](ga4ghVariant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantSet

> Ga4ghVariantSet GetVariantSet(ctx, variantSetId)
Gets a `VariantSet` by ID.

`GET /variantsets/{variant_set_id}` will return a JSON version of `VariantSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantSetId** | **string**| The ID of the &#x60;VariantSet&#x60; to be retrieved. | 

### Return type

[**Ga4ghVariantSet**](ga4ghVariantSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCallSets

> Ga4ghSearchCallSetsResponse SearchCallSets(ctx, body)
Gets a list of call sets matching the search criteria.

`POST /callsets/search` must accept a JSON version of `SearchCallSetsRequest` as the post body and will return a JSON version of `SearchCallSetsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchCallSetsRequest**](Ga4ghSearchCallSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchCallSetsResponse**](ga4ghSearchCallSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVariantSets

> Ga4ghSearchVariantSetsResponse SearchVariantSets(ctx, body)
Gets a list of `VariantSet` matching the search criteria.

`POST /variantsets/search` must accept a JSON version of `SearchVariantSetsRequest` as the post body and will return a JSON version of `SearchVariantSetsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchVariantSetsRequest**](Ga4ghSearchVariantSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchVariantSetsResponse**](ga4ghSearchVariantSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVariants

> Ga4ghSearchVariantsResponse SearchVariants(ctx, body)
Gets a list of `Variant` matching the search criteria.

`POST /variants/search` must accept a JSON version of `SearchVariantsRequest` as the post body and will return a JSON version of `SearchVariantsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchVariantsRequest**](Ga4ghSearchVariantsRequest.md)|  | 

### Return type

[**Ga4ghSearchVariantsResponse**](ga4ghSearchVariantsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVariantsByGeneName

> Ga4ghSearchVariantsByGeneNameResponse SearchVariantsByGeneName(ctx, body)
Gets a list of `Variants` matching the search criteria.

`POST /variantsetsByGeneName/search` must accept a JSON version of `SearchVariantsByGeneNameRequest` as the post body and will return a JSON version of `SearchVariantsByGeneNameResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchVariantsByGeneNameRequest**](Ga4ghSearchVariantsByGeneNameRequest.md)|  | 

### Return type

[**Ga4ghSearchVariantsByGeneNameResponse**](ga4ghSearchVariantsByGeneNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

