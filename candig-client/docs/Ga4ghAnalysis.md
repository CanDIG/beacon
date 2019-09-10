# Ga4ghAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Formats of id | name | description | accessions are described in the documentation on general attributes and formats. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | The time at which this record was created, in ISO 8601 format. | [optional] 
**Updated** | Pointer to **string** | The time at which this record was last updated, in ISO 8601 format. | [optional] 
**Type** | Pointer to **string** | The type of analysis. | [optional] 
**Software** | Pointer to **[]string** | The software run to generate this analysis. | [optional] 
**DatasetId** | Pointer to **string** | ### &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; ### # PROFYLE MODIFICATION BEGIN ### &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; ### The ID of the dataset this Analysis belongs to. | [optional] 
**ExperimentId** | Pointer to **string** |  | [optional] 
**OtherAnalysisDescriptor** | Pointer to **string** |  | [optional] 
**OtherAnalysisCompletitionDate** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghAnalysis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghAnalysis) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghAnalysis) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghAnalysis) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghAnalysis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghAnalysis) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghAnalysis) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghAnalysis) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghAnalysis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghAnalysis) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghAnalysis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghAnalysis) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghAnalysis) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghAnalysis) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghAnalysis) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghAnalysis) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghAnalysis) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghAnalysis) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghAnalysis) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghAnalysis) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetType

`func (o *Ga4ghAnalysis) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ga4ghAnalysis) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Ga4ghAnalysis) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Ga4ghAnalysis) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetSoftware

`func (o *Ga4ghAnalysis) GetSoftware() []string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *Ga4ghAnalysis) GetSoftwareOk() ([]string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSoftware

`func (o *Ga4ghAnalysis) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftware

`func (o *Ga4ghAnalysis) SetSoftware(v []string)`

SetSoftware gets a reference to the given []string and assigns it to the Software field.

### GetDatasetId

`func (o *Ga4ghAnalysis) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghAnalysis) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghAnalysis) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghAnalysis) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetExperimentId

`func (o *Ga4ghAnalysis) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *Ga4ghAnalysis) GetExperimentIdOk() (string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExperimentId

`func (o *Ga4ghAnalysis) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentId

`func (o *Ga4ghAnalysis) SetExperimentId(v string)`

SetExperimentId gets a reference to the given string and assigns it to the ExperimentId field.

### GetOtherAnalysisDescriptor

`func (o *Ga4ghAnalysis) GetOtherAnalysisDescriptor() string`

GetOtherAnalysisDescriptor returns the OtherAnalysisDescriptor field if non-nil, zero value otherwise.

### GetOtherAnalysisDescriptorOk

`func (o *Ga4ghAnalysis) GetOtherAnalysisDescriptorOk() (string, bool)`

GetOtherAnalysisDescriptorOk returns a tuple with the OtherAnalysisDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherAnalysisDescriptor

`func (o *Ga4ghAnalysis) HasOtherAnalysisDescriptor() bool`

HasOtherAnalysisDescriptor returns a boolean if a field has been set.

### SetOtherAnalysisDescriptor

`func (o *Ga4ghAnalysis) SetOtherAnalysisDescriptor(v string)`

SetOtherAnalysisDescriptor gets a reference to the given string and assigns it to the OtherAnalysisDescriptor field.

### GetOtherAnalysisCompletitionDate

`func (o *Ga4ghAnalysis) GetOtherAnalysisCompletitionDate() string`

GetOtherAnalysisCompletitionDate returns the OtherAnalysisCompletitionDate field if non-nil, zero value otherwise.

### GetOtherAnalysisCompletitionDateOk

`func (o *Ga4ghAnalysis) GetOtherAnalysisCompletitionDateOk() (string, bool)`

GetOtherAnalysisCompletitionDateOk returns a tuple with the OtherAnalysisCompletitionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherAnalysisCompletitionDate

`func (o *Ga4ghAnalysis) HasOtherAnalysisCompletitionDate() bool`

HasOtherAnalysisCompletitionDate returns a boolean if a field has been set.

### SetOtherAnalysisCompletitionDate

`func (o *Ga4ghAnalysis) SetOtherAnalysisCompletitionDate(v string)`

SetOtherAnalysisCompletitionDate gets a reference to the given string and assigns it to the OtherAnalysisCompletitionDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


