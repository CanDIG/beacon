# Ga4ghSearchEnrollmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enrollments** | Pointer to [**[]Ga4ghEnrollment**](ga4ghEnrollment.md) | The list of enrollments. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetEnrollments

`func (o *Ga4ghSearchEnrollmentsResponse) GetEnrollments() []Ga4ghEnrollment`

GetEnrollments returns the Enrollments field if non-nil, zero value otherwise.

### GetEnrollmentsOk

`func (o *Ga4ghSearchEnrollmentsResponse) GetEnrollmentsOk() ([]Ga4ghEnrollment, bool)`

GetEnrollmentsOk returns a tuple with the Enrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollments

`func (o *Ga4ghSearchEnrollmentsResponse) HasEnrollments() bool`

HasEnrollments returns a boolean if a field has been set.

### SetEnrollments

`func (o *Ga4ghSearchEnrollmentsResponse) SetEnrollments(v []Ga4ghEnrollment)`

SetEnrollments gets a reference to the given []Ga4ghEnrollment and assigns it to the Enrollments field.

### GetNextPageToken

`func (o *Ga4ghSearchEnrollmentsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchEnrollmentsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchEnrollmentsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchEnrollmentsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


