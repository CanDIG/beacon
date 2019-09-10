# Ga4ghSearchSequencingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequencing** | Pointer to [**[]Ga4ghSequencing**](ga4ghSequencing.md) | The list of sequencing metadata. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetSequencing

`func (o *Ga4ghSearchSequencingResponse) GetSequencing() []Ga4ghSequencing`

GetSequencing returns the Sequencing field if non-nil, zero value otherwise.

### GetSequencingOk

`func (o *Ga4ghSearchSequencingResponse) GetSequencingOk() ([]Ga4ghSequencing, bool)`

GetSequencingOk returns a tuple with the Sequencing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSequencing

`func (o *Ga4ghSearchSequencingResponse) HasSequencing() bool`

HasSequencing returns a boolean if a field has been set.

### SetSequencing

`func (o *Ga4ghSearchSequencingResponse) SetSequencing(v []Ga4ghSequencing)`

SetSequencing gets a reference to the given []Ga4ghSequencing and assigns it to the Sequencing field.

### GetNextPageToken

`func (o *Ga4ghSearchSequencingResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchSequencingResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchSequencingResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchSequencingResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


