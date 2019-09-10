# Ga4ghSearchDatasetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasets** | Pointer to [**[]Ga4ghDataset**](ga4ghDataset.md) | The list of datasets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetDatasets

`func (o *Ga4ghSearchDatasetsResponse) GetDatasets() []Ga4ghDataset`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *Ga4ghSearchDatasetsResponse) GetDatasetsOk() ([]Ga4ghDataset, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasets

`func (o *Ga4ghSearchDatasetsResponse) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### SetDatasets

`func (o *Ga4ghSearchDatasetsResponse) SetDatasets(v []Ga4ghDataset)`

SetDatasets gets a reference to the given []Ga4ghDataset and assigns it to the Datasets field.

### GetNextPageToken

`func (o *Ga4ghSearchDatasetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchDatasetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchDatasetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchDatasetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


