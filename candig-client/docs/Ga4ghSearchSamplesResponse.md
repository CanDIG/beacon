# Ga4ghSearchSamplesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]Ga4ghSample**](ga4ghSample.md) | The list of samples. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetSamples

`func (o *Ga4ghSearchSamplesResponse) GetSamples() []Ga4ghSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *Ga4ghSearchSamplesResponse) GetSamplesOk() ([]Ga4ghSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamples

`func (o *Ga4ghSearchSamplesResponse) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### SetSamples

`func (o *Ga4ghSearchSamplesResponse) SetSamples(v []Ga4ghSample)`

SetSamples gets a reference to the given []Ga4ghSample and assigns it to the Samples field.

### GetNextPageToken

`func (o *Ga4ghSearchSamplesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchSamplesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchSamplesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchSamplesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


