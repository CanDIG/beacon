# Ga4ghSearchContinuousResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Continuous** | Pointer to [**[]Ga4ghContinuous**](ga4ghContinuous.md) | The list of matching continuous values, sorted by start position. All sampled values within the query range are returned. Unsampled values are assigned &#39;NaN&#39; value. The values returned do not necessarily cover the same range as the query as all unsampled values might not be returned or if the query range extends beyond the reference range. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetContinuous

`func (o *Ga4ghSearchContinuousResponse) GetContinuous() []Ga4ghContinuous`

GetContinuous returns the Continuous field if non-nil, zero value otherwise.

### GetContinuousOk

`func (o *Ga4ghSearchContinuousResponse) GetContinuousOk() ([]Ga4ghContinuous, bool)`

GetContinuousOk returns a tuple with the Continuous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContinuous

`func (o *Ga4ghSearchContinuousResponse) HasContinuous() bool`

HasContinuous returns a boolean if a field has been set.

### SetContinuous

`func (o *Ga4ghSearchContinuousResponse) SetContinuous(v []Ga4ghContinuous)`

SetContinuous gets a reference to the given []Ga4ghContinuous and assigns it to the Continuous field.

### GetNextPageToken

`func (o *Ga4ghSearchContinuousResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchContinuousResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchContinuousResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchContinuousResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


