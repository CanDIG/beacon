# Ga4ghSearchConsentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consents** | Pointer to [**[]Ga4ghConsent**](ga4ghConsent.md) | The list of consents. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetConsents

`func (o *Ga4ghSearchConsentsResponse) GetConsents() []Ga4ghConsent`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *Ga4ghSearchConsentsResponse) GetConsentsOk() ([]Ga4ghConsent, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsents

`func (o *Ga4ghSearchConsentsResponse) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### SetConsents

`func (o *Ga4ghSearchConsentsResponse) SetConsents(v []Ga4ghConsent)`

SetConsents gets a reference to the given []Ga4ghConsent and assigns it to the Consents field.

### GetNextPageToken

`func (o *Ga4ghSearchConsentsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchConsentsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchConsentsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchConsentsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


