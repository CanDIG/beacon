# Ga4ghSearchReferenceSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceSets** | Pointer to [**[]Ga4ghReferenceSet**](ga4ghReferenceSet.md) | The list of matching reference sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetReferenceSets

`func (o *Ga4ghSearchReferenceSetsResponse) GetReferenceSets() []Ga4ghReferenceSet`

GetReferenceSets returns the ReferenceSets field if non-nil, zero value otherwise.

### GetReferenceSetsOk

`func (o *Ga4ghSearchReferenceSetsResponse) GetReferenceSetsOk() ([]Ga4ghReferenceSet, bool)`

GetReferenceSetsOk returns a tuple with the ReferenceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceSets

`func (o *Ga4ghSearchReferenceSetsResponse) HasReferenceSets() bool`

HasReferenceSets returns a boolean if a field has been set.

### SetReferenceSets

`func (o *Ga4ghSearchReferenceSetsResponse) SetReferenceSets(v []Ga4ghReferenceSet)`

SetReferenceSets gets a reference to the given []Ga4ghReferenceSet and assigns it to the ReferenceSets field.

### GetNextPageToken

`func (o *Ga4ghSearchReferenceSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchReferenceSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchReferenceSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchReferenceSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


