# Ga4ghSearchStudiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Studies** | Pointer to [**[]Ga4ghStudy**](ga4ghStudy.md) | The list of studies. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetStudies

`func (o *Ga4ghSearchStudiesResponse) GetStudies() []Ga4ghStudy`

GetStudies returns the Studies field if non-nil, zero value otherwise.

### GetStudiesOk

`func (o *Ga4ghSearchStudiesResponse) GetStudiesOk() ([]Ga4ghStudy, bool)`

GetStudiesOk returns a tuple with the Studies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStudies

`func (o *Ga4ghSearchStudiesResponse) HasStudies() bool`

HasStudies returns a boolean if a field has been set.

### SetStudies

`func (o *Ga4ghSearchStudiesResponse) SetStudies(v []Ga4ghStudy)`

SetStudies gets a reference to the given []Ga4ghStudy and assigns it to the Studies field.

### GetNextPageToken

`func (o *Ga4ghSearchStudiesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchStudiesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchStudiesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchStudiesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


