# \SequenceAnnotationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContinuousSet**](SequenceAnnotationServiceApi.md#GetContinuousSet) | **Get** /continuoussets/{continuous_set_id} | Gets a &#x60;ContinuousSet&#x60; by ID.
[**GetFeature**](SequenceAnnotationServiceApi.md#GetFeature) | **Get** /features/{feature_id} | Gets a &#x60;Feature&#x60; by ID.
[**GetFeatureSet**](SequenceAnnotationServiceApi.md#GetFeatureSet) | **Get** /featuresets/{feature_set_id} | Gets a &#x60;FeatureSet&#x60; by ID.
[**SearchContinuous**](SequenceAnnotationServiceApi.md#SearchContinuous) | **Post** /continuous/search | Gets continuous values matching the search criteria.
[**SearchContinuousSets**](SequenceAnnotationServiceApi.md#SearchContinuousSets) | **Post** /continuoussets/search | Gets a list of &#x60;ContinuousSet&#x60; matching the search criteria.
[**SearchFeatureSets**](SequenceAnnotationServiceApi.md#SearchFeatureSets) | **Post** /featuresets/search | Gets a list of &#x60;FeatureSet&#x60; matching the search criteria.
[**SearchFeatures**](SequenceAnnotationServiceApi.md#SearchFeatures) | **Post** /features/search | Gets a list of &#x60;Feature&#x60; matching the search criteria.



## GetContinuousSet

> Ga4ghContinuousSet GetContinuousSet(ctx, continuousSetId)
Gets a `ContinuousSet` by ID.

`GET /continuoussets/{id}` will return a JSON version of `ContinuousSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**continuousSetId** | **string**| The ID of the &#x60;ContinuousSet&#x60; to be retrieved. | 

### Return type

[**Ga4ghContinuousSet**](ga4ghContinuousSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> Ga4ghFeature GetFeature(ctx, featureId)
Gets a `Feature` by ID.

`GET /features/{id}` will return a JSON version of `Feature`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string**| The ID of the &#x60;Feature&#x60; to be retrieved. | 

### Return type

[**Ga4ghFeature**](ga4ghFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureSet

> Ga4ghFeatureSet GetFeatureSet(ctx, featureSetId)
Gets a `FeatureSet` by ID.

`GET /featuresets/{id}` will return a JSON version of `FeatureSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureSetId** | **string**| The ID of the &#x60;FeatureSet&#x60; to be retrieved. | 

### Return type

[**Ga4ghFeatureSet**](ga4ghFeatureSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchContinuous

> Ga4ghSearchContinuousResponse SearchContinuous(ctx, body)
Gets continuous values matching the search criteria.

`POST /continuous/search` must accept a JSON version of `SearchContinuousRequest` as the post body and will return a JSON version of `SearchContinuousResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchContinuousRequest**](Ga4ghSearchContinuousRequest.md)|  | 

### Return type

[**Ga4ghSearchContinuousResponse**](ga4ghSearchContinuousResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchContinuousSets

> Ga4ghSearchContinuousSetsResponse SearchContinuousSets(ctx, body)
Gets a list of `ContinuousSet` matching the search criteria.

`POST /continuoussets/search` must accept a JSON version of `SearchContinuousSetsRequest` as the post body and will return a JSON version of `SearchContinuousSetsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchContinuousSetsRequest**](Ga4ghSearchContinuousSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchContinuousSetsResponse**](ga4ghSearchContinuousSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFeatureSets

> Ga4ghSearchFeatureSetsResponse SearchFeatureSets(ctx, body)
Gets a list of `FeatureSet` matching the search criteria.

`POST /featuresets/search` must accept a JSON version of `SearchFeatureSetsRequest` as the post body and will return a JSON version of `SearchFeatureSetsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchFeatureSetsRequest**](Ga4ghSearchFeatureSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchFeatureSetsResponse**](ga4ghSearchFeatureSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFeatures

> Ga4ghSearchFeaturesResponse SearchFeatures(ctx, body)
Gets a list of `Feature` matching the search criteria.

`POST /features/search` must accept a JSON version of `SearchFeaturesRequest` as the post body and will return a JSON version of `SearchFeaturesResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchFeaturesRequest**](Ga4ghSearchFeaturesRequest.md)|  | 

### Return type

[**Ga4ghSearchFeaturesResponse**](ga4ghSearchFeaturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

