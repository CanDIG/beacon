# Ga4ghSearchFeatureSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureSets** | Pointer to [**[]Ga4ghFeatureSet**](ga4ghFeatureSet.md) | The list of matching feature sets. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetFeatureSets

`func (o *Ga4ghSearchFeatureSetsResponse) GetFeatureSets() []Ga4ghFeatureSet`

GetFeatureSets returns the FeatureSets field if non-nil, zero value otherwise.

### GetFeatureSetsOk

`func (o *Ga4ghSearchFeatureSetsResponse) GetFeatureSetsOk() ([]Ga4ghFeatureSet, bool)`

GetFeatureSetsOk returns a tuple with the FeatureSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureSets

`func (o *Ga4ghSearchFeatureSetsResponse) HasFeatureSets() bool`

HasFeatureSets returns a boolean if a field has been set.

### SetFeatureSets

`func (o *Ga4ghSearchFeatureSetsResponse) SetFeatureSets(v []Ga4ghFeatureSet)`

SetFeatureSets gets a reference to the given []Ga4ghFeatureSet and assigns it to the FeatureSets field.

### GetNextPageToken

`func (o *Ga4ghSearchFeatureSetsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchFeatureSetsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchFeatureSetsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchFeatureSetsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


