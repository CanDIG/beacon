# Ga4ghSearchReadGroupSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadGroupSets** | Pointer to [**[]Ga4ghReadGroupSet**](ga4ghReadGroupSet.md) | The list of matching read group sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetReadGroupSets

`func (o *Ga4ghSearchReadGroupSetsResponse) GetReadGroupSets() []Ga4ghReadGroupSet`

GetReadGroupSets returns the ReadGroupSets field if non-nil, zero value otherwise.

### GetReadGroupSetsOk

`func (o *Ga4ghSearchReadGroupSetsResponse) GetReadGroupSetsOk() ([]Ga4ghReadGroupSet, bool)`

GetReadGroupSetsOk returns a tuple with the ReadGroupSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadGroupSets

`func (o *Ga4ghSearchReadGroupSetsResponse) HasReadGroupSets() bool`

HasReadGroupSets returns a boolean if a field has been set.

### SetReadGroupSets

`func (o *Ga4ghSearchReadGroupSetsResponse) SetReadGroupSets(v []Ga4ghReadGroupSet)`

SetReadGroupSets gets a reference to the given []Ga4ghReadGroupSet and assigns it to the ReadGroupSets field.

### GetNextPageToken

`func (o *Ga4ghSearchReadGroupSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchReadGroupSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchReadGroupSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchReadGroupSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


