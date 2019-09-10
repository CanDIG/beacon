# Ga4ghSearchSurgeriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Surgeries** | Pointer to [**[]Ga4ghSurgery**](ga4ghSurgery.md) | The list of surgeries. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetSurgeries

`func (o *Ga4ghSearchSurgeriesResponse) GetSurgeries() []Ga4ghSurgery`

GetSurgeries returns the Surgeries field if non-nil, zero value otherwise.

### GetSurgeriesOk

`func (o *Ga4ghSearchSurgeriesResponse) GetSurgeriesOk() ([]Ga4ghSurgery, bool)`

GetSurgeriesOk returns a tuple with the Surgeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurgeries

`func (o *Ga4ghSearchSurgeriesResponse) HasSurgeries() bool`

HasSurgeries returns a boolean if a field has been set.

### SetSurgeries

`func (o *Ga4ghSearchSurgeriesResponse) SetSurgeries(v []Ga4ghSurgery)`

SetSurgeries gets a reference to the given []Ga4ghSurgery and assigns it to the Surgeries field.

### GetNextPageToken

`func (o *Ga4ghSearchSurgeriesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchSurgeriesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchSurgeriesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchSurgeriesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


