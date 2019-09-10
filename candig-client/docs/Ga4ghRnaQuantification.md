# Ga4ghRnaQuantification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID assigned to the results of running the described programs on the specified reads and assignment to the listed annotation. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BiosampleId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ReadGroupIds** | Pointer to **[]string** | ID(s) of the ReadGroup(s) providing the reads for the analysis. | [optional] 
**Programs** | Pointer to [**[]Ga4ghProgram**](ga4ghProgram.md) | Programs can be used to track the provenance of how read data was quantified. | [optional] 
**FeatureSetIds** | Pointer to **[]string** | List of annotation sets used. | [optional] 
**RnaQuantificationSetId** | Pointer to **string** | ID of the containing RnaQuantificationSet. | [optional] 
**SampleId** | Pointer to **string** |  | [optional] 
**PatientId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghRnaQuantification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghRnaQuantification) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghRnaQuantification) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghRnaQuantification) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghRnaQuantification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghRnaQuantification) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghRnaQuantification) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghRnaQuantification) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetBiosampleId

`func (o *Ga4ghRnaQuantification) GetBiosampleId() string`

GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.

### GetBiosampleIdOk

`func (o *Ga4ghRnaQuantification) GetBiosampleIdOk() (string, bool)`

GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiosampleId

`func (o *Ga4ghRnaQuantification) HasBiosampleId() bool`

HasBiosampleId returns a boolean if a field has been set.

### SetBiosampleId

`func (o *Ga4ghRnaQuantification) SetBiosampleId(v string)`

SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.

### GetDescription

`func (o *Ga4ghRnaQuantification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghRnaQuantification) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghRnaQuantification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghRnaQuantification) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetReadGroupIds

`func (o *Ga4ghRnaQuantification) GetReadGroupIds() []string`

GetReadGroupIds returns the ReadGroupIds field if non-nil, zero value otherwise.

### GetReadGroupIdsOk

`func (o *Ga4ghRnaQuantification) GetReadGroupIdsOk() ([]string, bool)`

GetReadGroupIdsOk returns a tuple with the ReadGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadGroupIds

`func (o *Ga4ghRnaQuantification) HasReadGroupIds() bool`

HasReadGroupIds returns a boolean if a field has been set.

### SetReadGroupIds

`func (o *Ga4ghRnaQuantification) SetReadGroupIds(v []string)`

SetReadGroupIds gets a reference to the given []string and assigns it to the ReadGroupIds field.

### GetPrograms

`func (o *Ga4ghRnaQuantification) GetPrograms() []Ga4ghProgram`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *Ga4ghRnaQuantification) GetProgramsOk() ([]Ga4ghProgram, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrograms

`func (o *Ga4ghRnaQuantification) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### SetPrograms

`func (o *Ga4ghRnaQuantification) SetPrograms(v []Ga4ghProgram)`

SetPrograms gets a reference to the given []Ga4ghProgram and assigns it to the Programs field.

### GetFeatureSetIds

`func (o *Ga4ghRnaQuantification) GetFeatureSetIds() []string`

GetFeatureSetIds returns the FeatureSetIds field if non-nil, zero value otherwise.

### GetFeatureSetIdsOk

`func (o *Ga4ghRnaQuantification) GetFeatureSetIdsOk() ([]string, bool)`

GetFeatureSetIdsOk returns a tuple with the FeatureSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureSetIds

`func (o *Ga4ghRnaQuantification) HasFeatureSetIds() bool`

HasFeatureSetIds returns a boolean if a field has been set.

### SetFeatureSetIds

`func (o *Ga4ghRnaQuantification) SetFeatureSetIds(v []string)`

SetFeatureSetIds gets a reference to the given []string and assigns it to the FeatureSetIds field.

### GetRnaQuantificationSetId

`func (o *Ga4ghRnaQuantification) GetRnaQuantificationSetId() string`

GetRnaQuantificationSetId returns the RnaQuantificationSetId field if non-nil, zero value otherwise.

### GetRnaQuantificationSetIdOk

`func (o *Ga4ghRnaQuantification) GetRnaQuantificationSetIdOk() (string, bool)`

GetRnaQuantificationSetIdOk returns a tuple with the RnaQuantificationSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaQuantificationSetId

`func (o *Ga4ghRnaQuantification) HasRnaQuantificationSetId() bool`

HasRnaQuantificationSetId returns a boolean if a field has been set.

### SetRnaQuantificationSetId

`func (o *Ga4ghRnaQuantification) SetRnaQuantificationSetId(v string)`

SetRnaQuantificationSetId gets a reference to the given string and assigns it to the RnaQuantificationSetId field.

### GetSampleId

`func (o *Ga4ghRnaQuantification) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghRnaQuantification) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghRnaQuantification) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghRnaQuantification) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetPatientId

`func (o *Ga4ghRnaQuantification) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghRnaQuantification) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghRnaQuantification) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghRnaQuantification) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


