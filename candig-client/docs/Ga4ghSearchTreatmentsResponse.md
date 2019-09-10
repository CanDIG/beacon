# Ga4ghSearchTreatmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Treatments** | Pointer to [**[]Ga4ghTreatment**](ga4ghTreatment.md) | The list of treatments. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetTreatments

`func (o *Ga4ghSearchTreatmentsResponse) GetTreatments() []Ga4ghTreatment`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *Ga4ghSearchTreatmentsResponse) GetTreatmentsOk() ([]Ga4ghTreatment, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatments

`func (o *Ga4ghSearchTreatmentsResponse) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### SetTreatments

`func (o *Ga4ghSearchTreatmentsResponse) SetTreatments(v []Ga4ghTreatment)`

SetTreatments gets a reference to the given []Ga4ghTreatment and assigns it to the Treatments field.

### GetNextPageToken

`func (o *Ga4ghSearchTreatmentsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchTreatmentsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchTreatmentsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchTreatmentsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


