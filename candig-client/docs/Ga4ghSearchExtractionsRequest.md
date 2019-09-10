# Ga4ghSearchExtractionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** | Optionally specify the dataset to search within. | [optional] 
**Name** | Pointer to **string** | Returns Extractions with the given name found by case-sensitive string matching. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page.If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets.To get the next page of results, set this parameter to the value of&#x60;nextPageToken&#x60; from the previous response. | [optional] 
**SampleId** | Pointer to **string** | Optionally specify the sample to search within. | [optional] 
**Filters** | Pointer to [**[]Ga4ghFilter**](ga4ghFilter.md) |  | [optional] 

## Methods

### GetDatasetId

`func (o *Ga4ghSearchExtractionsRequest) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghSearchExtractionsRequest) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghSearchExtractionsRequest) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghSearchExtractionsRequest) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghSearchExtractionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghSearchExtractionsRequest) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghSearchExtractionsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghSearchExtractionsRequest) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPageSize

`func (o *Ga4ghSearchExtractionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchExtractionsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchExtractionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchExtractionsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchExtractionsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchExtractionsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchExtractionsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchExtractionsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.

### GetSampleId

`func (o *Ga4ghSearchExtractionsRequest) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghSearchExtractionsRequest) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghSearchExtractionsRequest) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghSearchExtractionsRequest) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetFilters

`func (o *Ga4ghSearchExtractionsRequest) GetFilters() []Ga4ghFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Ga4ghSearchExtractionsRequest) GetFiltersOk() ([]Ga4ghFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *Ga4ghSearchExtractionsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *Ga4ghSearchExtractionsRequest) SetFilters(v []Ga4ghFilter)`

SetFilters gets a reference to the given []Ga4ghFilter and assigns it to the Filters field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


