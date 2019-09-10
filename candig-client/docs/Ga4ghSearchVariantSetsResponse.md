# Ga4ghSearchVariantSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantSets** | Pointer to [**[]Ga4ghVariantSet**](ga4ghVariantSet.md) | The list of matching variant sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetVariantSets

`func (o *Ga4ghSearchVariantSetsResponse) GetVariantSets() []Ga4ghVariantSet`

GetVariantSets returns the VariantSets field if non-nil, zero value otherwise.

### GetVariantSetsOk

`func (o *Ga4ghSearchVariantSetsResponse) GetVariantSetsOk() ([]Ga4ghVariantSet, bool)`

GetVariantSetsOk returns a tuple with the VariantSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantSets

`func (o *Ga4ghSearchVariantSetsResponse) HasVariantSets() bool`

HasVariantSets returns a boolean if a field has been set.

### SetVariantSets

`func (o *Ga4ghSearchVariantSetsResponse) SetVariantSets(v []Ga4ghVariantSet)`

SetVariantSets gets a reference to the given []Ga4ghVariantSet and assigns it to the VariantSets field.

### GetNextPageToken

`func (o *Ga4ghSearchVariantSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchVariantSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchVariantSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchVariantSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


