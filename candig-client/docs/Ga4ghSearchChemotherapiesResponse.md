# Ga4ghSearchChemotherapiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chemotherapies** | Pointer to [**[]Ga4ghChemotherapy**](ga4ghChemotherapy.md) | The list of chemotherapies. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetChemotherapies

`func (o *Ga4ghSearchChemotherapiesResponse) GetChemotherapies() []Ga4ghChemotherapy`

GetChemotherapies returns the Chemotherapies field if non-nil, zero value otherwise.

### GetChemotherapiesOk

`func (o *Ga4ghSearchChemotherapiesResponse) GetChemotherapiesOk() ([]Ga4ghChemotherapy, bool)`

GetChemotherapiesOk returns a tuple with the Chemotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChemotherapies

`func (o *Ga4ghSearchChemotherapiesResponse) HasChemotherapies() bool`

HasChemotherapies returns a boolean if a field has been set.

### SetChemotherapies

`func (o *Ga4ghSearchChemotherapiesResponse) SetChemotherapies(v []Ga4ghChemotherapy)`

SetChemotherapies gets a reference to the given []Ga4ghChemotherapy and assigns it to the Chemotherapies field.

### GetNextPageToken

`func (o *Ga4ghSearchChemotherapiesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchChemotherapiesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchChemotherapiesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchChemotherapiesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


