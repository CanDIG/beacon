# \ReadServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReadGroupSet**](ReadServiceApi.md#GetReadGroupSet) | **Get** /readgroupsets/{read_group_set_id} | Gets a &#x60;ReadGroupSet&#x60; by ID.
[**SearchReadGroupSets**](ReadServiceApi.md#SearchReadGroupSets) | **Post** /readgroupsets/search | Gets a list of &#x60;ReadGroupSet&#x60; matching the search criteria.
[**SearchReads**](ReadServiceApi.md#SearchReads) | **Post** /reads/search | Gets a list of &#x60;ReadAlignment&#x60; s for one or more &#x60;ReadGroup&#x60; s.



## GetReadGroupSet

> Ga4ghReadGroupSet GetReadGroupSet(ctx, readGroupSetId)
Gets a `ReadGroupSet` by ID.

`GET /readgroupsets/{read_group_set_id}` will return a JSON version of `ReadGroupSet`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readGroupSetId** | **string**| The ID of the &#x60;ReadGroupSet&#x60; to be retrieved. | 

### Return type

[**Ga4ghReadGroupSet**](ga4ghReadGroupSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReadGroupSets

> Ga4ghSearchReadGroupSetsResponse SearchReadGroupSets(ctx, body)
Gets a list of `ReadGroupSet` matching the search criteria.

`POST /readgroupsets/search` must accept a JSON version of `SearchReadGroupSetsRequest` as the post body and will return a JSON version of `SearchReadGroupSetsResponse` . Only readgroups that match an optionally supplied biosampleId will be included in the response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchReadGroupSetsRequest**](Ga4ghSearchReadGroupSetsRequest.md)|  | 

### Return type

[**Ga4ghSearchReadGroupSetsResponse**](ga4ghSearchReadGroupSetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchReads

> Ga4ghSearchReadsResponse SearchReads(ctx, body)
Gets a list of `ReadAlignment` s for one or more `ReadGroup` s.

`searchReads` operates over a genomic coordinate space of reference sequence and position defined by the `Reference` s to which the requested `ReadGroup` s are aligned.  If a target positional range is specified, search returns all reads whose alignment to the reference genome *overlap* the range. A query which specifies only read group IDs yields all reads in those read groups, including unmapped reads.  All reads returned (including reads on subsequent pages) are ordered by genomic coordinate (by reference sequence, then position). Reads with equivalent genomic coordinates are returned in an unspecified order. This order must be consistent for a given repository, such that two queries for the same content (regardless of page size) yield reads in the same order across their respective streams of paginated responses.  `POST /reads/search` must accept a JSON version of `SearchReadsRequest` as the post body and will return a JSON version of `SearchReadsResponse`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchReadsRequest**](Ga4ghSearchReadsRequest.md)|  | 

### Return type

[**Ga4ghSearchReadsResponse**](ga4ghSearchReadsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

