# \RnaQuantificationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpressionLevel**](RnaQuantificationServiceApi.md#GetExpressionLevel) | **Get** /expressionlevels/{expression_level_id} | Gets a &#x60;ExpressionLevel&#x60; by ID. &#x60;GET /expressionlevels/{id}&#x60; will return a JSON version of &#x60;ExpressionLevel&#x60;.
[**GetRnaQuantification**](RnaQuantificationServiceApi.md#GetRnaQuantification) | **Get** /rnaquantifications/{rna_quantification_id} | Gets a &#x60;RnaQuantification&#x60; by ID. &#x60;GET /rnaquantifications/{id}&#x60; will return a JSON version of &#x60;RnaQuantification&#x60;.
[**GetRnaQuantificationSet**](RnaQuantificationServiceApi.md#GetRnaQuantificationSet) | **Get** /rnaquantificationsets/{rna_quantification_set_id} | Gets a &#x60;RnaQuantificationSet&#x60; by ID. &#x60;GET /rnaquantificationsets/{id}&#x60; will return a JSON version of &#x60;RnaQuantificationSet&#x60;.
[**SearchExpressionLevels**](RnaQuantificationServiceApi.md#SearchExpressionLevels) | **Post** /expressionlevels/search | &#39;POST /expressionlevels/search&#39; must accept JSON version of &#39;SearchExpressionLevelsRequest&#39; as the post body and will return a JSON version of &#39;SearchExpressionLevelsResponse&#39;.
[**SearchRnaQuantificationSets**](RnaQuantificationServiceApi.md#SearchRnaQuantificationSets) | **Post** /rnaquantificationsets/search | Gets a list of &#39;RnaQuantificationSet&#39; matching the search criteria. &#39;POST /rnaquantificationsets/search&#39; must accept JSON version of &#39;SearchRnaQuantificationSetRequest&#39; as the post body and will return a JSON version of &#39;SearchRnaQuantificationSetResponse&#39;.
[**SearchRnaQuantifications**](RnaQuantificationServiceApi.md#SearchRnaQuantifications) | **Post** /rnaquantifications/search | Gets a list of &#39;RnaQuantification&#39; matching the search criteria. &#39;POST /rnaquantifications/search&#39; must accept JSON version of &#39;SearchRnaQuantificationsRequest&#39; as the post body and will return a JSON version of &#39;SearchRnaQuantificationResponse&#39;.



## GetExpressionLevel

> Ga4ghExpressionLevel GetExpressionLevel(ctx, expressionLevelId)
Gets a `ExpressionLevel` by ID. `GET /expressionlevels/{id}` will return a JSON version of `ExpressionLevel`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expressionLevelId** | **string**| The ID of the &#x60;ExpressionLevel&#x60;. | 

### Return type

[**Ga4ghExpressionLevel**](ga4ghExpressionLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRnaQuantification

> Ga4ghRnaQuantification GetRnaQuantification(ctx, rnaQuantificationId)
Gets a `RnaQuantification` by ID. `GET /rnaquantifications/{id}` will return a JSON version of `RnaQuantification`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rnaQuantificationId** | **string**| The ID of the &#x60;RnaQuantification&#x60;. | 

### Return type

[**Ga4ghRnaQuantification**](ga4ghRnaQuantification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRnaQuantificationSet

> Ga4ghRnaQuantificationSet GetRnaQuantificationSet(ctx, rnaQuantificationSetId)
Gets a `RnaQuantificationSet` by ID. `GET /rnaquantificationsets/{id}` will return a JSON version of `RnaQuantificationSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rnaQuantificationSetId** | **string**| The ID of the &#x60;RnaQuantificationSet&#x60;. | 

### Return type

[**Ga4ghRnaQuantificationSet**](ga4ghRnaQuantificationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchExpressionLevels

> Ga4ghSearchExpressionLevelsResponse SearchExpressionLevels(ctx, body)
'POST /expressionlevels/search' must accept JSON version of 'SearchExpressionLevelsRequest' as the post body and will return a JSON version of 'SearchExpressionLevelsResponse'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchExpressionLevelsRequest**](Ga4ghSearchExpressionLevelsRequest.md)|  | 

### Return type

[**Ga4ghSearchExpressionLevelsResponse**](ga4ghSearchExpressionLevelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRnaQuantificationSets

> Ga4ghSearchRnaQuantificationSetsResponse SearchRnaQuantificationSets(ctx, body)
Gets a list of 'RnaQuantificationSet' matching the search criteria. 'POST /rnaquantificationsets/search' must accept JSON version of 'SearchRnaQuantificationSetRequest' as the post body and will return a JSON version of 'SearchRnaQuantificationSetResponse'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchRnaQuantificationSetsRequest**](Ga4ghSearchRnaQuantificationSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchRnaQuantificationSetsResponse**](ga4ghSearchRnaQuantificationSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRnaQuantifications

> Ga4ghSearchRnaQuantificationsResponse SearchRnaQuantifications(ctx, body)
Gets a list of 'RnaQuantification' matching the search criteria. 'POST /rnaquantifications/search' must accept JSON version of 'SearchRnaQuantificationsRequest' as the post body and will return a JSON version of 'SearchRnaQuantificationResponse'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchRnaQuantificationsRequest**](Ga4ghSearchRnaQuantificationsRequest.md)|  | 

### Return type

[**Ga4ghSearchRnaQuantificationsResponse**](ga4ghSearchRnaQuantificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

