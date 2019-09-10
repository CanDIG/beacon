# Ga4ghSearchTumourboardsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tumourboards** | Pointer to [**[]Ga4ghTumourboard**](ga4ghTumourboard.md) | The list of tumourboards. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetTumourboards

`func (o *Ga4ghSearchTumourboardsResponse) GetTumourboards() []Ga4ghTumourboard`

GetTumourboards returns the Tumourboards field if non-nil, zero value otherwise.

### GetTumourboardsOk

`func (o *Ga4ghSearchTumourboardsResponse) GetTumourboardsOk() ([]Ga4ghTumourboard, bool)`

GetTumourboardsOk returns a tuple with the Tumourboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTumourboards

`func (o *Ga4ghSearchTumourboardsResponse) HasTumourboards() bool`

HasTumourboards returns a boolean if a field has been set.

### SetTumourboards

`func (o *Ga4ghSearchTumourboardsResponse) SetTumourboards(v []Ga4ghTumourboard)`

SetTumourboards gets a reference to the given []Ga4ghTumourboard and assigns it to the Tumourboards field.

### GetNextPageToken

`func (o *Ga4ghSearchTumourboardsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchTumourboardsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchTumourboardsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchTumourboardsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


