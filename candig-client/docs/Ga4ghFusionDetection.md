# Ga4ghFusionDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is unique in the context of the server instance. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the dataset this object belongs to. | [optional] 
**Name** | Pointer to **string** | This is a label or symbolic identifier for this object. | [optional] 
**Description** | Pointer to **string** | This attribute contains human readable text. | [optional] 
**Created** | Pointer to **string** | The :ref:&#x60;ISO 8601&lt;metadata_date_time&gt;&#x60; time at which this record was created. | [optional] 
**Updated** | Pointer to **string** | The :ref:&#x60;ISO 8601&lt;metadata_date_time&gt;&#x60; time at which this record was updated. | [optional] 
**SampleId** | Pointer to **string** |  | [optional] 
**InHousePipeline** | Pointer to **string** |  | [optional] 
**SvDetection** | Pointer to **string** |  | [optional] 
**FusionDetection** | Pointer to **string** |  | [optional] 
**Realignment** | Pointer to **string** |  | [optional] 
**Annotation** | Pointer to **string** |  | [optional] 
**GenomeReference** | Pointer to **string** |  | [optional] 
**GeneModels** | Pointer to **string** |  | [optional] 
**FusionDetectionId** | Pointer to **string** |  | [optional] 
**AlignmentId** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghFusionDetection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghFusionDetection) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghFusionDetection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghFusionDetection) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghFusionDetection) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghFusionDetection) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghFusionDetection) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghFusionDetection) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghFusionDetection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghFusionDetection) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghFusionDetection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghFusionDetection) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghFusionDetection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghFusionDetection) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghFusionDetection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghFusionDetection) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghFusionDetection) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghFusionDetection) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghFusionDetection) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghFusionDetection) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghFusionDetection) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghFusionDetection) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghFusionDetection) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghFusionDetection) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetSampleId

`func (o *Ga4ghFusionDetection) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghFusionDetection) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghFusionDetection) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghFusionDetection) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetInHousePipeline

`func (o *Ga4ghFusionDetection) GetInHousePipeline() string`

GetInHousePipeline returns the InHousePipeline field if non-nil, zero value otherwise.

### GetInHousePipelineOk

`func (o *Ga4ghFusionDetection) GetInHousePipelineOk() (string, bool)`

GetInHousePipelineOk returns a tuple with the InHousePipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInHousePipeline

`func (o *Ga4ghFusionDetection) HasInHousePipeline() bool`

HasInHousePipeline returns a boolean if a field has been set.

### SetInHousePipeline

`func (o *Ga4ghFusionDetection) SetInHousePipeline(v string)`

SetInHousePipeline gets a reference to the given string and assigns it to the InHousePipeline field.

### GetSvDetection

`func (o *Ga4ghFusionDetection) GetSvDetection() string`

GetSvDetection returns the SvDetection field if non-nil, zero value otherwise.

### GetSvDetectionOk

`func (o *Ga4ghFusionDetection) GetSvDetectionOk() (string, bool)`

GetSvDetectionOk returns a tuple with the SvDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSvDetection

`func (o *Ga4ghFusionDetection) HasSvDetection() bool`

HasSvDetection returns a boolean if a field has been set.

### SetSvDetection

`func (o *Ga4ghFusionDetection) SetSvDetection(v string)`

SetSvDetection gets a reference to the given string and assigns it to the SvDetection field.

### GetFusionDetection

`func (o *Ga4ghFusionDetection) GetFusionDetection() string`

GetFusionDetection returns the FusionDetection field if non-nil, zero value otherwise.

### GetFusionDetectionOk

`func (o *Ga4ghFusionDetection) GetFusionDetectionOk() (string, bool)`

GetFusionDetectionOk returns a tuple with the FusionDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFusionDetection

`func (o *Ga4ghFusionDetection) HasFusionDetection() bool`

HasFusionDetection returns a boolean if a field has been set.

### SetFusionDetection

`func (o *Ga4ghFusionDetection) SetFusionDetection(v string)`

SetFusionDetection gets a reference to the given string and assigns it to the FusionDetection field.

### GetRealignment

`func (o *Ga4ghFusionDetection) GetRealignment() string`

GetRealignment returns the Realignment field if non-nil, zero value otherwise.

### GetRealignmentOk

`func (o *Ga4ghFusionDetection) GetRealignmentOk() (string, bool)`

GetRealignmentOk returns a tuple with the Realignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRealignment

`func (o *Ga4ghFusionDetection) HasRealignment() bool`

HasRealignment returns a boolean if a field has been set.

### SetRealignment

`func (o *Ga4ghFusionDetection) SetRealignment(v string)`

SetRealignment gets a reference to the given string and assigns it to the Realignment field.

### GetAnnotation

`func (o *Ga4ghFusionDetection) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *Ga4ghFusionDetection) GetAnnotationOk() (string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnnotation

`func (o *Ga4ghFusionDetection) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotation

`func (o *Ga4ghFusionDetection) SetAnnotation(v string)`

SetAnnotation gets a reference to the given string and assigns it to the Annotation field.

### GetGenomeReference

`func (o *Ga4ghFusionDetection) GetGenomeReference() string`

GetGenomeReference returns the GenomeReference field if non-nil, zero value otherwise.

### GetGenomeReferenceOk

`func (o *Ga4ghFusionDetection) GetGenomeReferenceOk() (string, bool)`

GetGenomeReferenceOk returns a tuple with the GenomeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGenomeReference

`func (o *Ga4ghFusionDetection) HasGenomeReference() bool`

HasGenomeReference returns a boolean if a field has been set.

### SetGenomeReference

`func (o *Ga4ghFusionDetection) SetGenomeReference(v string)`

SetGenomeReference gets a reference to the given string and assigns it to the GenomeReference field.

### GetGeneModels

`func (o *Ga4ghFusionDetection) GetGeneModels() string`

GetGeneModels returns the GeneModels field if non-nil, zero value otherwise.

### GetGeneModelsOk

`func (o *Ga4ghFusionDetection) GetGeneModelsOk() (string, bool)`

GetGeneModelsOk returns a tuple with the GeneModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeneModels

`func (o *Ga4ghFusionDetection) HasGeneModels() bool`

HasGeneModels returns a boolean if a field has been set.

### SetGeneModels

`func (o *Ga4ghFusionDetection) SetGeneModels(v string)`

SetGeneModels gets a reference to the given string and assigns it to the GeneModels field.

### GetFusionDetectionId

`func (o *Ga4ghFusionDetection) GetFusionDetectionId() string`

GetFusionDetectionId returns the FusionDetectionId field if non-nil, zero value otherwise.

### GetFusionDetectionIdOk

`func (o *Ga4ghFusionDetection) GetFusionDetectionIdOk() (string, bool)`

GetFusionDetectionIdOk returns a tuple with the FusionDetectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFusionDetectionId

`func (o *Ga4ghFusionDetection) HasFusionDetectionId() bool`

HasFusionDetectionId returns a boolean if a field has been set.

### SetFusionDetectionId

`func (o *Ga4ghFusionDetection) SetFusionDetectionId(v string)`

SetFusionDetectionId gets a reference to the given string and assigns it to the FusionDetectionId field.

### GetAlignmentId

`func (o *Ga4ghFusionDetection) GetAlignmentId() string`

GetAlignmentId returns the AlignmentId field if non-nil, zero value otherwise.

### GetAlignmentIdOk

`func (o *Ga4ghFusionDetection) GetAlignmentIdOk() (string, bool)`

GetAlignmentIdOk returns a tuple with the AlignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignmentId

`func (o *Ga4ghFusionDetection) HasAlignmentId() bool`

HasAlignmentId returns a boolean if a field has been set.

### SetAlignmentId

`func (o *Ga4ghFusionDetection) SetAlignmentId(v string)`

SetAlignmentId gets a reference to the given string and assigns it to the AlignmentId field.

### GetSite

`func (o *Ga4ghFusionDetection) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Ga4ghFusionDetection) GetSiteOk() (string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSite

`func (o *Ga4ghFusionDetection) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSite

`func (o *Ga4ghFusionDetection) SetSite(v string)`

SetSite gets a reference to the given string and assigns it to the Site field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


