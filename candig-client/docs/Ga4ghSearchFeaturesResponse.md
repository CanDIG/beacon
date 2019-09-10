# Ga4ghSearchFeaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]Ga4ghFeature**](ga4ghFeature.md) | The list of matching annotations, sorted by start position. Annotations which share a start position are returned in a deterministic order. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetFeatures

`func (o *Ga4ghSearchFeaturesResponse) GetFeatures() []Ga4ghFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Ga4ghSearchFeaturesResponse) GetFeaturesOk() ([]Ga4ghFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *Ga4ghSearchFeaturesResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *Ga4ghSearchFeaturesResponse) SetFeatures(v []Ga4ghFeature)`

SetFeatures gets a reference to the given []Ga4ghFeature and assigns it to the Features field.

### GetNextPageToken

`func (o *Ga4ghSearchFeaturesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchFeaturesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchFeaturesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchFeaturesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


