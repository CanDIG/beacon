# Ga4ghReadGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The read group ID. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the dataset this read group belongs to. | [optional] 
**Name** | Pointer to **string** | The read group name. | [optional] 
**Description** | Pointer to **string** | The read group description. | [optional] 
**SampleName** | Pointer to **string** | A name for the sample this read group&#39;s data were generated from. This field contains an arbitrary string, typically corresponding to the SM tag in a BAM file. | [optional] 
**BiosampleId** | Pointer to **string** | The Biosample this read group&#39;s data was generated from. | [optional] 
**Experiment** | Pointer to [**Ga4ghExperiment**](ga4ghExperiment.md) |  | [optional] 
**PredictedInsertSize** | Pointer to **int32** | The predicted insert size of this read group. | [optional] 
**Created** | Pointer to **string** | The time at which this read group was created in milliseconds from the epoch. | [optional] 
**Updated** | Pointer to **string** | The time at which this read group was last updated in milliseconds from the epoch. | [optional] 
**Stats** | Pointer to [**Ga4ghReadStats**](ga4ghReadStats.md) |  | [optional] 
**Programs** | Pointer to [**[]Ga4ghProgram**](ga4ghProgram.md) | Program can be used to track the provenance of how read data was generated. | [optional] 
**ReferenceSetId** | Pointer to **string** | The ID of the reference set to which the reads in this read group are aligned. Required if there are any read alignments. | [optional] 

## Methods

### GetId

`func (o *Ga4ghReadGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghReadGroup) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghReadGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghReadGroup) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghReadGroup) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghReadGroup) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghReadGroup) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghReadGroup) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghReadGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghReadGroup) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghReadGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghReadGroup) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghReadGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghReadGroup) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghReadGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghReadGroup) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSampleName

`func (o *Ga4ghReadGroup) GetSampleName() string`

GetSampleName returns the SampleName field if non-nil, zero value otherwise.

### GetSampleNameOk

`func (o *Ga4ghReadGroup) GetSampleNameOk() (string, bool)`

GetSampleNameOk returns a tuple with the SampleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleName

`func (o *Ga4ghReadGroup) HasSampleName() bool`

HasSampleName returns a boolean if a field has been set.

### SetSampleName

`func (o *Ga4ghReadGroup) SetSampleName(v string)`

SetSampleName gets a reference to the given string and assigns it to the SampleName field.

### GetBiosampleId

`func (o *Ga4ghReadGroup) GetBiosampleId() string`

GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.

### GetBiosampleIdOk

`func (o *Ga4ghReadGroup) GetBiosampleIdOk() (string, bool)`

GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiosampleId

`func (o *Ga4ghReadGroup) HasBiosampleId() bool`

HasBiosampleId returns a boolean if a field has been set.

### SetBiosampleId

`func (o *Ga4ghReadGroup) SetBiosampleId(v string)`

SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.

### GetExperiment

`func (o *Ga4ghReadGroup) GetExperiment() Ga4ghExperiment`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *Ga4ghReadGroup) GetExperimentOk() (Ga4ghExperiment, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperiment

`func (o *Ga4ghReadGroup) HasExperiment() bool`

HasExperiment returns a boolean if a field has been set.

### SetExperiment

`func (o *Ga4ghReadGroup) SetExperiment(v Ga4ghExperiment)`

SetExperiment gets a reference to the given Ga4ghExperiment and assigns it to the Experiment field.

### GetPredictedInsertSize

`func (o *Ga4ghReadGroup) GetPredictedInsertSize() int32`

GetPredictedInsertSize returns the PredictedInsertSize field if non-nil, zero value otherwise.

### GetPredictedInsertSizeOk

`func (o *Ga4ghReadGroup) GetPredictedInsertSizeOk() (int32, bool)`

GetPredictedInsertSizeOk returns a tuple with the PredictedInsertSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPredictedInsertSize

`func (o *Ga4ghReadGroup) HasPredictedInsertSize() bool`

HasPredictedInsertSize returns a boolean if a field has been set.

### SetPredictedInsertSize

`func (o *Ga4ghReadGroup) SetPredictedInsertSize(v int32)`

SetPredictedInsertSize gets a reference to the given int32 and assigns it to the PredictedInsertSize field.

### GetCreated

`func (o *Ga4ghReadGroup) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghReadGroup) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghReadGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghReadGroup) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghReadGroup) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghReadGroup) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghReadGroup) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghReadGroup) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetStats

`func (o *Ga4ghReadGroup) GetStats() Ga4ghReadStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Ga4ghReadGroup) GetStatsOk() (Ga4ghReadStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStats

`func (o *Ga4ghReadGroup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStats

`func (o *Ga4ghReadGroup) SetStats(v Ga4ghReadStats)`

SetStats gets a reference to the given Ga4ghReadStats and assigns it to the Stats field.

### GetPrograms

`func (o *Ga4ghReadGroup) GetPrograms() []Ga4ghProgram`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *Ga4ghReadGroup) GetProgramsOk() ([]Ga4ghProgram, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrograms

`func (o *Ga4ghReadGroup) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### SetPrograms

`func (o *Ga4ghReadGroup) SetPrograms(v []Ga4ghProgram)`

SetPrograms gets a reference to the given []Ga4ghProgram and assigns it to the Programs field.

### GetReferenceSetId

`func (o *Ga4ghReadGroup) GetReferenceSetId() string`

GetReferenceSetId returns the ReferenceSetId field if non-nil, zero value otherwise.

### GetReferenceSetIdOk

`func (o *Ga4ghReadGroup) GetReferenceSetIdOk() (string, bool)`

GetReferenceSetIdOk returns a tuple with the ReferenceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceSetId

`func (o *Ga4ghReadGroup) HasReferenceSetId() bool`

HasReferenceSetId returns a boolean if a field has been set.

### SetReferenceSetId

`func (o *Ga4ghReadGroup) SetReferenceSetId(v string)`

SetReferenceSetId gets a reference to the given string and assigns it to the ReferenceSetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


