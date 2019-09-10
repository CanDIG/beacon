# Ga4ghSearchPatientsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patients** | Pointer to [**[]Ga4ghPatient**](ga4ghPatient.md) | The list of patients. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetPatients

`func (o *Ga4ghSearchPatientsResponse) GetPatients() []Ga4ghPatient`

GetPatients returns the Patients field if non-nil, zero value otherwise.

### GetPatientsOk

`func (o *Ga4ghSearchPatientsResponse) GetPatientsOk() ([]Ga4ghPatient, bool)`

GetPatientsOk returns a tuple with the Patients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatients

`func (o *Ga4ghSearchPatientsResponse) HasPatients() bool`

HasPatients returns a boolean if a field has been set.

### SetPatients

`func (o *Ga4ghSearchPatientsResponse) SetPatients(v []Ga4ghPatient)`

SetPatients gets a reference to the given []Ga4ghPatient and assigns it to the Patients field.

### GetNextPageToken

`func (o *Ga4ghSearchPatientsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchPatientsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchPatientsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchPatientsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


