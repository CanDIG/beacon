# Ga4ghReadGroupSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The read group set ID. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the dataset this read group set belongs to. | [optional] 
**Name** | Pointer to **string** | The read group set name. | [optional] 
**Stats** | Pointer to [**Ga4ghReadStats**](ga4ghReadStats.md) |  | [optional] 
**ReadGroups** | Pointer to [**[]Ga4ghReadGroup**](ga4ghReadGroup.md) | The read groups in this set. | [optional] 
**PatientId** | Pointer to **string** |  | [optional] 
**SampleId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghReadGroupSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghReadGroupSet) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghReadGroupSet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghReadGroupSet) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghReadGroupSet) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghReadGroupSet) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghReadGroupSet) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghReadGroupSet) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghReadGroupSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghReadGroupSet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghReadGroupSet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghReadGroupSet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStats

`func (o *Ga4ghReadGroupSet) GetStats() Ga4ghReadStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Ga4ghReadGroupSet) GetStatsOk() (Ga4ghReadStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStats

`func (o *Ga4ghReadGroupSet) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStats

`func (o *Ga4ghReadGroupSet) SetStats(v Ga4ghReadStats)`

SetStats gets a reference to the given Ga4ghReadStats and assigns it to the Stats field.

### GetReadGroups

`func (o *Ga4ghReadGroupSet) GetReadGroups() []Ga4ghReadGroup`

GetReadGroups returns the ReadGroups field if non-nil, zero value otherwise.

### GetReadGroupsOk

`func (o *Ga4ghReadGroupSet) GetReadGroupsOk() ([]Ga4ghReadGroup, bool)`

GetReadGroupsOk returns a tuple with the ReadGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadGroups

`func (o *Ga4ghReadGroupSet) HasReadGroups() bool`

HasReadGroups returns a boolean if a field has been set.

### SetReadGroups

`func (o *Ga4ghReadGroupSet) SetReadGroups(v []Ga4ghReadGroup)`

SetReadGroups gets a reference to the given []Ga4ghReadGroup and assigns it to the ReadGroups field.

### GetPatientId

`func (o *Ga4ghReadGroupSet) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghReadGroupSet) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghReadGroupSet) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghReadGroupSet) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetSampleId

`func (o *Ga4ghReadGroupSet) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghReadGroupSet) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghReadGroupSet) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghReadGroupSet) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


