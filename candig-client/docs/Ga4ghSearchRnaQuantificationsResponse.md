# Ga4ghSearchRnaQuantificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RnaQuantifications** | Pointer to [**[]Ga4ghRnaQuantification**](ga4ghRnaQuantification.md) | The list of matching quantifications. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#39;nextPageToken&#39; from the previous response. | [optional] 

## Methods

### GetRnaQuantifications

`func (o *Ga4ghSearchRnaQuantificationsResponse) GetRnaQuantifications() []Ga4ghRnaQuantification`

GetRnaQuantifications returns the RnaQuantifications field if non-nil, zero value otherwise.

### GetRnaQuantificationsOk

`func (o *Ga4ghSearchRnaQuantificationsResponse) GetRnaQuantificationsOk() ([]Ga4ghRnaQuantification, bool)`

GetRnaQuantificationsOk returns a tuple with the RnaQuantifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaQuantifications

`func (o *Ga4ghSearchRnaQuantificationsResponse) HasRnaQuantifications() bool`

HasRnaQuantifications returns a boolean if a field has been set.

### SetRnaQuantifications

`func (o *Ga4ghSearchRnaQuantificationsResponse) SetRnaQuantifications(v []Ga4ghRnaQuantification)`

SetRnaQuantifications gets a reference to the given []Ga4ghRnaQuantification and assigns it to the RnaQuantifications field.

### GetNextPageToken

`func (o *Ga4ghSearchRnaQuantificationsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchRnaQuantificationsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchRnaQuantificationsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchRnaQuantificationsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


