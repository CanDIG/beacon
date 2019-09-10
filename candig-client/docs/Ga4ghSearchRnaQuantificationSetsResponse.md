# Ga4ghSearchRnaQuantificationSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RnaQuantificationSets** | Pointer to [**[]Ga4ghRnaQuantificationSet**](ga4ghRnaQuantificationSet.md) | The list of matching quantification sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#39;nextPageToken&#39; from the previous response. | [optional] 

## Methods

### GetRnaQuantificationSets

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) GetRnaQuantificationSets() []Ga4ghRnaQuantificationSet`

GetRnaQuantificationSets returns the RnaQuantificationSets field if non-nil, zero value otherwise.

### GetRnaQuantificationSetsOk

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) GetRnaQuantificationSetsOk() ([]Ga4ghRnaQuantificationSet, bool)`

GetRnaQuantificationSetsOk returns a tuple with the RnaQuantificationSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaQuantificationSets

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) HasRnaQuantificationSets() bool`

HasRnaQuantificationSets returns a boolean if a field has been set.

### SetRnaQuantificationSets

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) SetRnaQuantificationSets(v []Ga4ghRnaQuantificationSet)`

SetRnaQuantificationSets gets a reference to the given []Ga4ghRnaQuantificationSet and assigns it to the RnaQuantificationSets field.

### GetNextPageToken

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchRnaQuantificationSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


