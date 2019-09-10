# Ga4ghListReferenceBasesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** | The ID of the &#x60;Reference&#x60; to be retrieved. | [optional] 
**Start** | Pointer to **string** | The start position (0-based) of this query. Defaults to 0. Genomic positions are non-negative integers less than reference length. Requests spanning the join of circular genomes are represented as two requests one on each side of the join (position 0). | [optional] 
**End** | Pointer to **string** | The end position (0-based, exclusive) of this query. Defaults to the length of this &#x60;Reference&#x60;. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 

## Methods

### GetReferenceId

`func (o *Ga4ghListReferenceBasesRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Ga4ghListReferenceBasesRequest) GetReferenceIdOk() (string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceId

`func (o *Ga4ghListReferenceBasesRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceId

`func (o *Ga4ghListReferenceBasesRequest) SetReferenceId(v string)`

SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.

### GetStart

`func (o *Ga4ghListReferenceBasesRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghListReferenceBasesRequest) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghListReferenceBasesRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghListReferenceBasesRequest) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghListReferenceBasesRequest) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghListReferenceBasesRequest) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghListReferenceBasesRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghListReferenceBasesRequest) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetPageToken

`func (o *Ga4ghListReferenceBasesRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghListReferenceBasesRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghListReferenceBasesRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghListReferenceBasesRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


