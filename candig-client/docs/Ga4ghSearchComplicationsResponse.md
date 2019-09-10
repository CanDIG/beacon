# Ga4ghSearchComplicationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complications** | Pointer to [**[]Ga4ghComplication**](ga4ghComplication.md) | The list of complications. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetComplications

`func (o *Ga4ghSearchComplicationsResponse) GetComplications() []Ga4ghComplication`

GetComplications returns the Complications field if non-nil, zero value otherwise.

### GetComplicationsOk

`func (o *Ga4ghSearchComplicationsResponse) GetComplicationsOk() ([]Ga4ghComplication, bool)`

GetComplicationsOk returns a tuple with the Complications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplications

`func (o *Ga4ghSearchComplicationsResponse) HasComplications() bool`

HasComplications returns a boolean if a field has been set.

### SetComplications

`func (o *Ga4ghSearchComplicationsResponse) SetComplications(v []Ga4ghComplication)`

SetComplications gets a reference to the given []Ga4ghComplication and assigns it to the Complications field.

### GetNextPageToken

`func (o *Ga4ghSearchComplicationsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchComplicationsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchComplicationsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchComplicationsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


