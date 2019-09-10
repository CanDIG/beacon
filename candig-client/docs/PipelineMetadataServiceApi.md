# \PipelineMetadataServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlignment**](PipelineMetadataServiceApi.md#GetAlignment) | **Get** /alignments/{alignment_id} | Gets &#x60;Alignment&#x60; by ID.
[**GetExpressionAnalysis**](PipelineMetadataServiceApi.md#GetExpressionAnalysis) | **Get** /expressionanalysis/{expressionanalysis_id} | Gets &#x60;ExpressionAnalysis&#x60; by ID.
[**GetExtraction**](PipelineMetadataServiceApi.md#GetExtraction) | **Get** /extractions/{extraction_id} | Gets &#x60;Extraction&#x60; by ID.
[**GetFusionDetection**](PipelineMetadataServiceApi.md#GetFusionDetection) | **Get** /fusiondetection/{fusiondetection_id} | Gets &#x60;FusionDetection&#x60; by ID.
[**GetSequencing**](PipelineMetadataServiceApi.md#GetSequencing) | **Get** /sequencing/{sequencing_id} | Gets &#x60;Sequencing&#x60; by ID.
[**GetVariantCalling**](PipelineMetadataServiceApi.md#GetVariantCalling) | **Get** /variantcalling/{variantcalling_id} | Gets &#x60;VariantCalling&#x60; by ID.
[**SearchAlignments**](PipelineMetadataServiceApi.md#SearchAlignments) | **Post** /alignments/search | Gets a list of Alignments accessible through the API.
[**SearchExpressionAnalysis**](PipelineMetadataServiceApi.md#SearchExpressionAnalysis) | **Post** /expressionanalysis/search | Gets a list of ExpressionAnalysis metadata accessible through the API.
[**SearchExtractions**](PipelineMetadataServiceApi.md#SearchExtractions) | **Post** /extractions/search | Gets a list of Extractions accessible through the API.
[**SearchFusionDetections**](PipelineMetadataServiceApi.md#SearchFusionDetections) | **Post** /fusiondetections/search | Gets a list of FusionDetections accessible through the API.
[**SearchSequencing**](PipelineMetadataServiceApi.md#SearchSequencing) | **Post** /sequencing/search | Gets a list of Sequencing metadata accessible through the API.
[**SearchVariantCalling**](PipelineMetadataServiceApi.md#SearchVariantCalling) | **Post** /variantcalling/search | Gets a list of VariantCalling metadata accessible through the API.



## GetAlignment

> Ga4ghAlignment GetAlignment(ctx, alignmentId)
Gets `Alignment` by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alignmentId** | **string**| The ID of the alignment requested | 

### Return type

[**Ga4ghAlignment**](ga4ghAlignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpressionAnalysis

> Ga4ghExpressionAnalysis GetExpressionAnalysis(ctx, expressionanalysisId)
Gets `ExpressionAnalysis` by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expressionanalysisId** | **string**| The ID of the expressionanalysis requested | 

### Return type

[**Ga4ghExpressionAnalysis**](ga4ghExpressionAnalysis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtraction

> Ga4ghExtraction GetExtraction(ctx, extractionId)
Gets `Extraction` by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extractionId** | **string**| The ID of the patient requested | 

### Return type

[**Ga4ghExtraction**](ga4ghExtraction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFusionDetection

> Ga4ghFusionDetection GetFusionDetection(ctx, fusiondetectionId)
Gets `FusionDetection` by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fusiondetectionId** | **string**| The ID of the fusion/sv detection requested | 

### Return type

[**Ga4ghFusionDetection**](ga4ghFusionDetection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSequencing

> Ga4ghSequencing GetSequencing(ctx, sequencingId)
Gets `Sequencing` by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequencingId** | **string**| The ID of the sequencing requested | 

### Return type

[**Ga4ghSequencing**](ga4ghSequencing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantCalling

> Ga4ghVariantCalling GetVariantCalling(ctx, variantcallingId)
Gets `VariantCalling` by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantcallingId** | **string**| The ID of the variant calling requested | 

### Return type

[**Ga4ghVariantCalling**](ga4ghVariantCalling.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAlignments

> Ga4ghSearchAlignmentsResponse SearchAlignments(ctx, body)
Gets a list of Alignments accessible through the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchAlignmentsRequest**](Ga4ghSearchAlignmentsRequest.md)|  | 

### Return type

[**Ga4ghSearchAlignmentsResponse**](ga4ghSearchAlignmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchExpressionAnalysis

> Ga4ghSearchExpressionAnalysisResponse SearchExpressionAnalysis(ctx, body)
Gets a list of ExpressionAnalysis metadata accessible through the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchExpressionAnalysisRequest**](Ga4ghSearchExpressionAnalysisRequest.md)|  | 

### Return type

[**Ga4ghSearchExpressionAnalysisResponse**](ga4ghSearchExpressionAnalysisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchExtractions

> Ga4ghSearchExtractionsResponse SearchExtractions(ctx, body)
Gets a list of Extractions accessible through the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchExtractionsRequest**](Ga4ghSearchExtractionsRequest.md)|  | 

### Return type

[**Ga4ghSearchExtractionsResponse**](ga4ghSearchExtractionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFusionDetections

> Ga4ghSearchFusionDetectionResponse SearchFusionDetections(ctx, body)
Gets a list of FusionDetections accessible through the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchFusionDetectionRequest**](Ga4ghSearchFusionDetectionRequest.md)|  | 

### Return type

[**Ga4ghSearchFusionDetectionResponse**](ga4ghSearchFusionDetectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSequencing

> Ga4ghSearchSequencingResponse SearchSequencing(ctx, body)
Gets a list of Sequencing metadata accessible through the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchSequencingRequest**](Ga4ghSearchSequencingRequest.md)|  | 

### Return type

[**Ga4ghSearchSequencingResponse**](ga4ghSearchSequencingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVariantCalling

> Ga4ghSearchVariantCallingResponse SearchVariantCalling(ctx, body)
Gets a list of VariantCalling metadata accessible through the API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchVariantCallingRequest**](Ga4ghSearchVariantCallingRequest.md)|  | 

### Return type

[**Ga4ghSearchVariantCallingResponse**](ga4ghSearchVariantCallingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

