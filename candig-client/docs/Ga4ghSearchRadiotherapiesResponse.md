# Ga4ghSearchRadiotherapiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Radiotherapies** | Pointer to [**[]Ga4ghRadiotherapy**](ga4ghRadiotherapy.md) | The list of radiotherapies. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetRadiotherapies

`func (o *Ga4ghSearchRadiotherapiesResponse) GetRadiotherapies() []Ga4ghRadiotherapy`

GetRadiotherapies returns the Radiotherapies field if non-nil, zero value otherwise.

### GetRadiotherapiesOk

`func (o *Ga4ghSearchRadiotherapiesResponse) GetRadiotherapiesOk() ([]Ga4ghRadiotherapy, bool)`

GetRadiotherapiesOk returns a tuple with the Radiotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRadiotherapies

`func (o *Ga4ghSearchRadiotherapiesResponse) HasRadiotherapies() bool`

HasRadiotherapies returns a boolean if a field has been set.

### SetRadiotherapies

`func (o *Ga4ghSearchRadiotherapiesResponse) SetRadiotherapies(v []Ga4ghRadiotherapy)`

SetRadiotherapies gets a reference to the given []Ga4ghRadiotherapy and assigns it to the Radiotherapies field.

### GetNextPageToken

`func (o *Ga4ghSearchRadiotherapiesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchRadiotherapiesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchRadiotherapiesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchRadiotherapiesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


