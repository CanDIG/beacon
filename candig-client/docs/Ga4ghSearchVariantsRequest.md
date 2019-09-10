# Ga4ghSearchVariantsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** | The &#x60;dataset&#x60; to search. | [optional] 
**CallSetIds** | Pointer to **[]string** | Only return variant calls which belong to call sets with these IDs. If unspecified, return all variants and no variant call objects. | [optional] 
**ReferenceName** | Pointer to **string** | Required. Only return variants on this reference. | [optional] 
**Start** | Pointer to **string** | Required. The beginning of the window (0-based, inclusive) for which overlapping variants should be returned. Genomic positions are non-negative integers less than reference length. Requests spanning the join of circular genomes are represented as two requests one on each side of the join (position 0). | [optional] 
**End** | Pointer to **string** | Required. The end of the window (0-based, exclusive) for which overlapping variants should be returned. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 
**VariantSetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### GetDatasetId

`func (o *Ga4ghSearchVariantsRequest) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghSearchVariantsRequest) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghSearchVariantsRequest) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghSearchVariantsRequest) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetCallSetIds

`func (o *Ga4ghSearchVariantsRequest) GetCallSetIds() []string`

GetCallSetIds returns the CallSetIds field if non-nil, zero value otherwise.

### GetCallSetIdsOk

`func (o *Ga4ghSearchVariantsRequest) GetCallSetIdsOk() ([]string, bool)`

GetCallSetIdsOk returns a tuple with the CallSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallSetIds

`func (o *Ga4ghSearchVariantsRequest) HasCallSetIds() bool`

HasCallSetIds returns a boolean if a field has been set.

### SetCallSetIds

`func (o *Ga4ghSearchVariantsRequest) SetCallSetIds(v []string)`

SetCallSetIds gets a reference to the given []string and assigns it to the CallSetIds field.

### GetReferenceName

`func (o *Ga4ghSearchVariantsRequest) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghSearchVariantsRequest) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghSearchVariantsRequest) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghSearchVariantsRequest) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetStart

`func (o *Ga4ghSearchVariantsRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghSearchVariantsRequest) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghSearchVariantsRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghSearchVariantsRequest) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghSearchVariantsRequest) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghSearchVariantsRequest) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghSearchVariantsRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghSearchVariantsRequest) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetPageSize

`func (o *Ga4ghSearchVariantsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchVariantsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchVariantsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchVariantsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchVariantsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchVariantsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchVariantsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchVariantsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.

### GetVariantSetIds

`func (o *Ga4ghSearchVariantsRequest) GetVariantSetIds() []string`

GetVariantSetIds returns the VariantSetIds field if non-nil, zero value otherwise.

### GetVariantSetIdsOk

`func (o *Ga4ghSearchVariantsRequest) GetVariantSetIdsOk() ([]string, bool)`

GetVariantSetIdsOk returns a tuple with the VariantSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantSetIds

`func (o *Ga4ghSearchVariantsRequest) HasVariantSetIds() bool`

HasVariantSetIds returns a boolean if a field has been set.

### SetVariantSetIds

`func (o *Ga4ghSearchVariantsRequest) SetVariantSetIds(v []string)`

SetVariantSetIds gets a reference to the given []string and assigns it to the VariantSetIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


