# Ga4ghSearchLabtestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labtests** | Pointer to [**[]Ga4ghLabtest**](ga4ghLabtest.md) | The list of labtests. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetLabtests

`func (o *Ga4ghSearchLabtestsResponse) GetLabtests() []Ga4ghLabtest`

GetLabtests returns the Labtests field if non-nil, zero value otherwise.

### GetLabtestsOk

`func (o *Ga4ghSearchLabtestsResponse) GetLabtestsOk() ([]Ga4ghLabtest, bool)`

GetLabtestsOk returns a tuple with the Labtests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabtests

`func (o *Ga4ghSearchLabtestsResponse) HasLabtests() bool`

HasLabtests returns a boolean if a field has been set.

### SetLabtests

`func (o *Ga4ghSearchLabtestsResponse) SetLabtests(v []Ga4ghLabtest)`

SetLabtests gets a reference to the given []Ga4ghLabtest and assigns it to the Labtests field.

### GetNextPageToken

`func (o *Ga4ghSearchLabtestsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchLabtestsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchLabtestsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchLabtestsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


