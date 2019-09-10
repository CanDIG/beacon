# Ga4ghSearchQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** |  | [optional] 
**Logic** | Pointer to [**Ga4ghLogic**](ga4ghLogic.md) |  | [optional] 
**Components** | Pointer to [**[]Ga4ghComponent**](ga4ghComponent.md) |  | [optional] 
**Results** | Pointer to [**[]Ga4ghResult**](ga4ghResult.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;nextPageToken&#x60; from the previous response. | [optional] 

## Methods

### GetDatasetId

`func (o *Ga4ghSearchQueryRequest) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghSearchQueryRequest) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghSearchQueryRequest) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghSearchQueryRequest) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetLogic

`func (o *Ga4ghSearchQueryRequest) GetLogic() Ga4ghLogic`

GetLogic returns the Logic field if non-nil, zero value otherwise.

### GetLogicOk

`func (o *Ga4ghSearchQueryRequest) GetLogicOk() (Ga4ghLogic, bool)`

GetLogicOk returns a tuple with the Logic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogic

`func (o *Ga4ghSearchQueryRequest) HasLogic() bool`

HasLogic returns a boolean if a field has been set.

### SetLogic

`func (o *Ga4ghSearchQueryRequest) SetLogic(v Ga4ghLogic)`

SetLogic gets a reference to the given Ga4ghLogic and assigns it to the Logic field.

### GetComponents

`func (o *Ga4ghSearchQueryRequest) GetComponents() []Ga4ghComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Ga4ghSearchQueryRequest) GetComponentsOk() ([]Ga4ghComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComponents

`func (o *Ga4ghSearchQueryRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponents

`func (o *Ga4ghSearchQueryRequest) SetComponents(v []Ga4ghComponent)`

SetComponents gets a reference to the given []Ga4ghComponent and assigns it to the Components field.

### GetResults

`func (o *Ga4ghSearchQueryRequest) GetResults() []Ga4ghResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Ga4ghSearchQueryRequest) GetResultsOk() ([]Ga4ghResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResults

`func (o *Ga4ghSearchQueryRequest) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResults

`func (o *Ga4ghSearchQueryRequest) SetResults(v []Ga4ghResult)`

SetResults gets a reference to the given []Ga4ghResult and assigns it to the Results field.

### GetLimit

`func (o *Ga4ghSearchQueryRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Ga4ghSearchQueryRequest) GetLimitOk() (int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *Ga4ghSearchQueryRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *Ga4ghSearchQueryRequest) SetLimit(v int32)`

SetLimit gets a reference to the given int32 and assigns it to the Limit field.

### GetPageSize

`func (o *Ga4ghSearchQueryRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchQueryRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchQueryRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchQueryRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchQueryRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchQueryRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchQueryRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchQueryRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


