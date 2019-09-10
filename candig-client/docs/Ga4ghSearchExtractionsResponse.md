# Ga4ghSearchExtractionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extractions** | Pointer to [**[]Ga4ghExtraction**](ga4ghExtraction.md) | The list of extractions. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetExtractions

`func (o *Ga4ghSearchExtractionsResponse) GetExtractions() []Ga4ghExtraction`

GetExtractions returns the Extractions field if non-nil, zero value otherwise.

### GetExtractionsOk

`func (o *Ga4ghSearchExtractionsResponse) GetExtractionsOk() ([]Ga4ghExtraction, bool)`

GetExtractionsOk returns a tuple with the Extractions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtractions

`func (o *Ga4ghSearchExtractionsResponse) HasExtractions() bool`

HasExtractions returns a boolean if a field has been set.

### SetExtractions

`func (o *Ga4ghSearchExtractionsResponse) SetExtractions(v []Ga4ghExtraction)`

SetExtractions gets a reference to the given []Ga4ghExtraction and assigns it to the Extractions field.

### GetNextPageToken

`func (o *Ga4ghSearchExtractionsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchExtractionsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchExtractionsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchExtractionsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


