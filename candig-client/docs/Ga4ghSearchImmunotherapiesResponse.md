# Ga4ghSearchImmunotherapiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immunotherapies** | Pointer to [**[]Ga4ghImmunotherapy**](ga4ghImmunotherapy.md) | The list of immunotherapies. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetImmunotherapies

`func (o *Ga4ghSearchImmunotherapiesResponse) GetImmunotherapies() []Ga4ghImmunotherapy`

GetImmunotherapies returns the Immunotherapies field if non-nil, zero value otherwise.

### GetImmunotherapiesOk

`func (o *Ga4ghSearchImmunotherapiesResponse) GetImmunotherapiesOk() ([]Ga4ghImmunotherapy, bool)`

GetImmunotherapiesOk returns a tuple with the Immunotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImmunotherapies

`func (o *Ga4ghSearchImmunotherapiesResponse) HasImmunotherapies() bool`

HasImmunotherapies returns a boolean if a field has been set.

### SetImmunotherapies

`func (o *Ga4ghSearchImmunotherapiesResponse) SetImmunotherapies(v []Ga4ghImmunotherapy)`

SetImmunotherapies gets a reference to the given []Ga4ghImmunotherapy and assigns it to the Immunotherapies field.

### GetNextPageToken

`func (o *Ga4ghSearchImmunotherapiesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchImmunotherapiesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchImmunotherapiesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchImmunotherapiesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


