# Ga4ghSearchReferencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to [**[]Ga4ghReference**](ga4ghReference.md) | The list of matching references. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetReferences

`func (o *Ga4ghSearchReferencesResponse) GetReferences() []Ga4ghReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Ga4ghSearchReferencesResponse) GetReferencesOk() ([]Ga4ghReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferences

`func (o *Ga4ghSearchReferencesResponse) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferences

`func (o *Ga4ghSearchReferencesResponse) SetReferences(v []Ga4ghReference)`

SetReferences gets a reference to the given []Ga4ghReference and assigns it to the References field.

### GetNextPageToken

`func (o *Ga4ghSearchReferencesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchReferencesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchReferencesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchReferencesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


