# Ga4ghSearchRnaQuantificationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RnaQuantificationSetId** | Pointer to **string** | Return only Rna Quantifications which belong to this set. Must be specified. | [optional] 
**BiosampleId** | Pointer to **string** | Return only RNA quantifications regarding the specified biosample. Optional. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#39;nextPageToken&#39; from the previous response. | [optional] 
**SampleId** | Pointer to **string** | Return only RNA quantifications related to the specified sampleId. | [optional] 
**PatientId** | Pointer to **string** | Return only RNA quantifications related to the specified patientId. | [optional] 

## Methods

### GetRnaQuantificationSetId

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetRnaQuantificationSetId() string`

GetRnaQuantificationSetId returns the RnaQuantificationSetId field if non-nil, zero value otherwise.

### GetRnaQuantificationSetIdOk

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetRnaQuantificationSetIdOk() (string, bool)`

GetRnaQuantificationSetIdOk returns a tuple with the RnaQuantificationSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaQuantificationSetId

`func (o *Ga4ghSearchRnaQuantificationsRequest) HasRnaQuantificationSetId() bool`

HasRnaQuantificationSetId returns a boolean if a field has been set.

### SetRnaQuantificationSetId

`func (o *Ga4ghSearchRnaQuantificationsRequest) SetRnaQuantificationSetId(v string)`

SetRnaQuantificationSetId gets a reference to the given string and assigns it to the RnaQuantificationSetId field.

### GetBiosampleId

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetBiosampleId() string`

GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.

### GetBiosampleIdOk

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetBiosampleIdOk() (string, bool)`

GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiosampleId

`func (o *Ga4ghSearchRnaQuantificationsRequest) HasBiosampleId() bool`

HasBiosampleId returns a boolean if a field has been set.

### SetBiosampleId

`func (o *Ga4ghSearchRnaQuantificationsRequest) SetBiosampleId(v string)`

SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.

### GetPageSize

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchRnaQuantificationsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchRnaQuantificationsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchRnaQuantificationsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchRnaQuantificationsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.

### GetSampleId

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghSearchRnaQuantificationsRequest) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghSearchRnaQuantificationsRequest) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetPatientId

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghSearchRnaQuantificationsRequest) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghSearchRnaQuantificationsRequest) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghSearchRnaQuantificationsRequest) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


