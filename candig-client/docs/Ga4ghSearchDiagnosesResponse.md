# Ga4ghSearchDiagnosesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diagnoses** | Pointer to [**[]Ga4ghDiagnosis**](ga4ghDiagnosis.md) | The list of diagnoses. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetDiagnoses

`func (o *Ga4ghSearchDiagnosesResponse) GetDiagnoses() []Ga4ghDiagnosis`

GetDiagnoses returns the Diagnoses field if non-nil, zero value otherwise.

### GetDiagnosesOk

`func (o *Ga4ghSearchDiagnosesResponse) GetDiagnosesOk() ([]Ga4ghDiagnosis, bool)`

GetDiagnosesOk returns a tuple with the Diagnoses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnoses

`func (o *Ga4ghSearchDiagnosesResponse) HasDiagnoses() bool`

HasDiagnoses returns a boolean if a field has been set.

### SetDiagnoses

`func (o *Ga4ghSearchDiagnosesResponse) SetDiagnoses(v []Ga4ghDiagnosis)`

SetDiagnoses gets a reference to the given []Ga4ghDiagnosis and assigns it to the Diagnoses field.

### GetNextPageToken

`func (o *Ga4ghSearchDiagnosesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchDiagnosesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchDiagnosesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchDiagnosesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


