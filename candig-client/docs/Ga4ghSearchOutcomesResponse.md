# Ga4ghSearchOutcomesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcomes** | Pointer to [**[]Ga4ghOutcome**](ga4ghOutcome.md) | The list of outcomes. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetOutcomes

`func (o *Ga4ghSearchOutcomesResponse) GetOutcomes() []Ga4ghOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *Ga4ghSearchOutcomesResponse) GetOutcomesOk() ([]Ga4ghOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutcomes

`func (o *Ga4ghSearchOutcomesResponse) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomes

`func (o *Ga4ghSearchOutcomesResponse) SetOutcomes(v []Ga4ghOutcome)`

SetOutcomes gets a reference to the given []Ga4ghOutcome and assigns it to the Outcomes field.

### GetNextPageToken

`func (o *Ga4ghSearchOutcomesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchOutcomesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchOutcomesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchOutcomesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


