# Ga4ghSearchReadsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadGroupIds** | Pointer to **[]string** | The ReadGroups to search. At least one id must be specified. | [optional] 
**ReferenceId** | Pointer to **string** | The reference to query. Leaving blank returns results from all references, including unmapped reads - this could be very large. | [optional] 
**Start** | Pointer to **string** | The start position (0-based) of this query. If a reference is specified, this defaults to 0. Genomic positions are non-negative integers less than reference length. Requests spanning the join of circular genomes are represented as two requests one on each side of the join (position 0). | [optional] 
**End** | Pointer to **string** | The end position (0-based, exclusive) of this query. If a reference is specified, this defaults to the reference&#39;s length. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 

## Methods

### GetReadGroupIds

`func (o *Ga4ghSearchReadsRequest) GetReadGroupIds() []string`

GetReadGroupIds returns the ReadGroupIds field if non-nil, zero value otherwise.

### GetReadGroupIdsOk

`func (o *Ga4ghSearchReadsRequest) GetReadGroupIdsOk() ([]string, bool)`

GetReadGroupIdsOk returns a tuple with the ReadGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadGroupIds

`func (o *Ga4ghSearchReadsRequest) HasReadGroupIds() bool`

HasReadGroupIds returns a boolean if a field has been set.

### SetReadGroupIds

`func (o *Ga4ghSearchReadsRequest) SetReadGroupIds(v []string)`

SetReadGroupIds gets a reference to the given []string and assigns it to the ReadGroupIds field.

### GetReferenceId

`func (o *Ga4ghSearchReadsRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Ga4ghSearchReadsRequest) GetReferenceIdOk() (string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceId

`func (o *Ga4ghSearchReadsRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceId

`func (o *Ga4ghSearchReadsRequest) SetReferenceId(v string)`

SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.

### GetStart

`func (o *Ga4ghSearchReadsRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghSearchReadsRequest) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghSearchReadsRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghSearchReadsRequest) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghSearchReadsRequest) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghSearchReadsRequest) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghSearchReadsRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghSearchReadsRequest) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetPageSize

`func (o *Ga4ghSearchReadsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchReadsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchReadsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchReadsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchReadsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchReadsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchReadsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchReadsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


