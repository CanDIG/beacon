# Ga4ghSearchCelltransplantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Celltransplants** | Pointer to [**[]Ga4ghCelltransplant**](ga4ghCelltransplant.md) | The list of celltransplants. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetCelltransplants

`func (o *Ga4ghSearchCelltransplantsResponse) GetCelltransplants() []Ga4ghCelltransplant`

GetCelltransplants returns the Celltransplants field if non-nil, zero value otherwise.

### GetCelltransplantsOk

`func (o *Ga4ghSearchCelltransplantsResponse) GetCelltransplantsOk() ([]Ga4ghCelltransplant, bool)`

GetCelltransplantsOk returns a tuple with the Celltransplants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCelltransplants

`func (o *Ga4ghSearchCelltransplantsResponse) HasCelltransplants() bool`

HasCelltransplants returns a boolean if a field has been set.

### SetCelltransplants

`func (o *Ga4ghSearchCelltransplantsResponse) SetCelltransplants(v []Ga4ghCelltransplant)`

SetCelltransplants gets a reference to the given []Ga4ghCelltransplant and assigns it to the Celltransplants field.

### GetNextPageToken

`func (o *Ga4ghSearchCelltransplantsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchCelltransplantsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchCelltransplantsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchCelltransplantsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


