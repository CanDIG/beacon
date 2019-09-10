# Ga4ghSearchCallSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallSets** | Pointer to [**[]Ga4ghCallSet**](ga4ghCallSet.md) | The list of matching call sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetCallSets

`func (o *Ga4ghSearchCallSetsResponse) GetCallSets() []Ga4ghCallSet`

GetCallSets returns the CallSets field if non-nil, zero value otherwise.

### GetCallSetsOk

`func (o *Ga4ghSearchCallSetsResponse) GetCallSetsOk() ([]Ga4ghCallSet, bool)`

GetCallSetsOk returns a tuple with the CallSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallSets

`func (o *Ga4ghSearchCallSetsResponse) HasCallSets() bool`

HasCallSets returns a boolean if a field has been set.

### SetCallSets

`func (o *Ga4ghSearchCallSetsResponse) SetCallSets(v []Ga4ghCallSet)`

SetCallSets gets a reference to the given []Ga4ghCallSet and assigns it to the CallSets field.

### GetNextPageToken

`func (o *Ga4ghSearchCallSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchCallSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchCallSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchCallSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


