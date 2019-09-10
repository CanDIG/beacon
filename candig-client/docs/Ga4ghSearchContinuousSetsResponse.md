# Ga4ghSearchContinuousSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousSets** | Pointer to [**[]Ga4ghContinuousSet**](ga4ghContinuousSet.md) | The list of matching feature sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetContinuousSets

`func (o *Ga4ghSearchContinuousSetsResponse) GetContinuousSets() []Ga4ghContinuousSet`

GetContinuousSets returns the ContinuousSets field if non-nil, zero value otherwise.

### GetContinuousSetsOk

`func (o *Ga4ghSearchContinuousSetsResponse) GetContinuousSetsOk() ([]Ga4ghContinuousSet, bool)`

GetContinuousSetsOk returns a tuple with the ContinuousSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContinuousSets

`func (o *Ga4ghSearchContinuousSetsResponse) HasContinuousSets() bool`

HasContinuousSets returns a boolean if a field has been set.

### SetContinuousSets

`func (o *Ga4ghSearchContinuousSetsResponse) SetContinuousSets(v []Ga4ghContinuousSet)`

SetContinuousSets gets a reference to the given []Ga4ghContinuousSet and assigns it to the ContinuousSets field.

### GetNextPageToken

`func (o *Ga4ghSearchContinuousSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchContinuousSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchContinuousSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchContinuousSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


