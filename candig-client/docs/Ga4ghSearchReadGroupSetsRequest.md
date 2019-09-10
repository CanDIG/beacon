# Ga4ghSearchReadGroupSetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** | The dataset to search. | [optional] 
**Name** | Pointer to **string** | Only return read group sets with this name (case-sensitive, exact match). | [optional] 
**BiosampleId** | Pointer to **string** | Specifying the id of a Biosample record will return only readgroups  with the given biosampleId. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 

## Methods

### GetDatasetId

`func (o *Ga4ghSearchReadGroupSetsRequest) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghSearchReadGroupSetsRequest) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghSearchReadGroupSetsRequest) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghSearchReadGroupSetsRequest) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghSearchReadGroupSetsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghSearchReadGroupSetsRequest) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghSearchReadGroupSetsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghSearchReadGroupSetsRequest) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetBiosampleId

`func (o *Ga4ghSearchReadGroupSetsRequest) GetBiosampleId() string`

GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.

### GetBiosampleIdOk

`func (o *Ga4ghSearchReadGroupSetsRequest) GetBiosampleIdOk() (string, bool)`

GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiosampleId

`func (o *Ga4ghSearchReadGroupSetsRequest) HasBiosampleId() bool`

HasBiosampleId returns a boolean if a field has been set.

### SetBiosampleId

`func (o *Ga4ghSearchReadGroupSetsRequest) SetBiosampleId(v string)`

SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.

### GetPageSize

`func (o *Ga4ghSearchReadGroupSetsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchReadGroupSetsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchReadGroupSetsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchReadGroupSetsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchReadGroupSetsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchReadGroupSetsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchReadGroupSetsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchReadGroupSetsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


