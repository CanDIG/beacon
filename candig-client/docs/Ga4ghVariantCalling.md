# Ga4ghVariantCalling

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
**VariantCaller** | Pointer to **string** |  | [optional] 
**Tabulate** | Pointer to **string** |  | [optional] 
**Annotation** | Pointer to **string** |  | [optional] 
**MergeTool** | Pointer to **string** |  | [optional] 
**RdaToTab** | Pointer to **string** |  | [optional] 
**Delly** | Pointer to **string** |  | [optional] 
**PostFilter** | Pointer to **string** |  | [optional] 
**ClipFilter** | Pointer to **string** |  | [optional] 
**Cosmic** | Pointer to **string** |  | [optional] 
**DbSnp** | Pointer to **string** |  | [optional] 
**VariantCallingId** | Pointer to **string** |  | [optional] 
**AlignmentId** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghVariantCalling) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghVariantCalling) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghVariantCalling) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghVariantCalling) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghVariantCalling) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghVariantCalling) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghVariantCalling) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghVariantCalling) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghVariantCalling) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghVariantCalling) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghVariantCalling) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghVariantCalling) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghVariantCalling) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghVariantCalling) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghVariantCalling) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghVariantCalling) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghVariantCalling) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghVariantCalling) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghVariantCalling) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghVariantCalling) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghVariantCalling) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghVariantCalling) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghVariantCalling) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghVariantCalling) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetSampleId

`func (o *Ga4ghVariantCalling) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghVariantCalling) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghVariantCalling) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghVariantCalling) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetInHousePipeline

`func (o *Ga4ghVariantCalling) GetInHousePipeline() string`

GetInHousePipeline returns the InHousePipeline field if non-nil, zero value otherwise.

### GetInHousePipelineOk

`func (o *Ga4ghVariantCalling) GetInHousePipelineOk() (string, bool)`

GetInHousePipelineOk returns a tuple with the InHousePipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInHousePipeline

`func (o *Ga4ghVariantCalling) HasInHousePipeline() bool`

HasInHousePipeline returns a boolean if a field has been set.

### SetInHousePipeline

`func (o *Ga4ghVariantCalling) SetInHousePipeline(v string)`

SetInHousePipeline gets a reference to the given string and assigns it to the InHousePipeline field.

### GetVariantCaller

`func (o *Ga4ghVariantCalling) GetVariantCaller() string`

GetVariantCaller returns the VariantCaller field if non-nil, zero value otherwise.

### GetVariantCallerOk

`func (o *Ga4ghVariantCalling) GetVariantCallerOk() (string, bool)`

GetVariantCallerOk returns a tuple with the VariantCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantCaller

`func (o *Ga4ghVariantCalling) HasVariantCaller() bool`

HasVariantCaller returns a boolean if a field has been set.

### SetVariantCaller

`func (o *Ga4ghVariantCalling) SetVariantCaller(v string)`

SetVariantCaller gets a reference to the given string and assigns it to the VariantCaller field.

### GetTabulate

`func (o *Ga4ghVariantCalling) GetTabulate() string`

GetTabulate returns the Tabulate field if non-nil, zero value otherwise.

### GetTabulateOk

`func (o *Ga4ghVariantCalling) GetTabulateOk() (string, bool)`

GetTabulateOk returns a tuple with the Tabulate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTabulate

`func (o *Ga4ghVariantCalling) HasTabulate() bool`

HasTabulate returns a boolean if a field has been set.

### SetTabulate

`func (o *Ga4ghVariantCalling) SetTabulate(v string)`

SetTabulate gets a reference to the given string and assigns it to the Tabulate field.

### GetAnnotation

`func (o *Ga4ghVariantCalling) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *Ga4ghVariantCalling) GetAnnotationOk() (string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnnotation

`func (o *Ga4ghVariantCalling) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotation

`func (o *Ga4ghVariantCalling) SetAnnotation(v string)`

SetAnnotation gets a reference to the given string and assigns it to the Annotation field.

### GetMergeTool

`func (o *Ga4ghVariantCalling) GetMergeTool() string`

GetMergeTool returns the MergeTool field if non-nil, zero value otherwise.

### GetMergeToolOk

`func (o *Ga4ghVariantCalling) GetMergeToolOk() (string, bool)`

GetMergeToolOk returns a tuple with the MergeTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMergeTool

`func (o *Ga4ghVariantCalling) HasMergeTool() bool`

HasMergeTool returns a boolean if a field has been set.

### SetMergeTool

`func (o *Ga4ghVariantCalling) SetMergeTool(v string)`

SetMergeTool gets a reference to the given string and assigns it to the MergeTool field.

### GetRdaToTab

`func (o *Ga4ghVariantCalling) GetRdaToTab() string`

GetRdaToTab returns the RdaToTab field if non-nil, zero value otherwise.

### GetRdaToTabOk

`func (o *Ga4ghVariantCalling) GetRdaToTabOk() (string, bool)`

GetRdaToTabOk returns a tuple with the RdaToTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRdaToTab

`func (o *Ga4ghVariantCalling) HasRdaToTab() bool`

HasRdaToTab returns a boolean if a field has been set.

### SetRdaToTab

`func (o *Ga4ghVariantCalling) SetRdaToTab(v string)`

SetRdaToTab gets a reference to the given string and assigns it to the RdaToTab field.

### GetDelly

`func (o *Ga4ghVariantCalling) GetDelly() string`

GetDelly returns the Delly field if non-nil, zero value otherwise.

### GetDellyOk

`func (o *Ga4ghVariantCalling) GetDellyOk() (string, bool)`

GetDellyOk returns a tuple with the Delly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDelly

`func (o *Ga4ghVariantCalling) HasDelly() bool`

HasDelly returns a boolean if a field has been set.

### SetDelly

`func (o *Ga4ghVariantCalling) SetDelly(v string)`

SetDelly gets a reference to the given string and assigns it to the Delly field.

### GetPostFilter

`func (o *Ga4ghVariantCalling) GetPostFilter() string`

GetPostFilter returns the PostFilter field if non-nil, zero value otherwise.

### GetPostFilterOk

`func (o *Ga4ghVariantCalling) GetPostFilterOk() (string, bool)`

GetPostFilterOk returns a tuple with the PostFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPostFilter

`func (o *Ga4ghVariantCalling) HasPostFilter() bool`

HasPostFilter returns a boolean if a field has been set.

### SetPostFilter

`func (o *Ga4ghVariantCalling) SetPostFilter(v string)`

SetPostFilter gets a reference to the given string and assigns it to the PostFilter field.

### GetClipFilter

`func (o *Ga4ghVariantCalling) GetClipFilter() string`

GetClipFilter returns the ClipFilter field if non-nil, zero value otherwise.

### GetClipFilterOk

`func (o *Ga4ghVariantCalling) GetClipFilterOk() (string, bool)`

GetClipFilterOk returns a tuple with the ClipFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClipFilter

`func (o *Ga4ghVariantCalling) HasClipFilter() bool`

HasClipFilter returns a boolean if a field has been set.

### SetClipFilter

`func (o *Ga4ghVariantCalling) SetClipFilter(v string)`

SetClipFilter gets a reference to the given string and assigns it to the ClipFilter field.

### GetCosmic

`func (o *Ga4ghVariantCalling) GetCosmic() string`

GetCosmic returns the Cosmic field if non-nil, zero value otherwise.

### GetCosmicOk

`func (o *Ga4ghVariantCalling) GetCosmicOk() (string, bool)`

GetCosmicOk returns a tuple with the Cosmic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCosmic

`func (o *Ga4ghVariantCalling) HasCosmic() bool`

HasCosmic returns a boolean if a field has been set.

### SetCosmic

`func (o *Ga4ghVariantCalling) SetCosmic(v string)`

SetCosmic gets a reference to the given string and assigns it to the Cosmic field.

### GetDbSnp

`func (o *Ga4ghVariantCalling) GetDbSnp() string`

GetDbSnp returns the DbSnp field if non-nil, zero value otherwise.

### GetDbSnpOk

`func (o *Ga4ghVariantCalling) GetDbSnpOk() (string, bool)`

GetDbSnpOk returns a tuple with the DbSnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDbSnp

`func (o *Ga4ghVariantCalling) HasDbSnp() bool`

HasDbSnp returns a boolean if a field has been set.

### SetDbSnp

`func (o *Ga4ghVariantCalling) SetDbSnp(v string)`

SetDbSnp gets a reference to the given string and assigns it to the DbSnp field.

### GetVariantCallingId

`func (o *Ga4ghVariantCalling) GetVariantCallingId() string`

GetVariantCallingId returns the VariantCallingId field if non-nil, zero value otherwise.

### GetVariantCallingIdOk

`func (o *Ga4ghVariantCalling) GetVariantCallingIdOk() (string, bool)`

GetVariantCallingIdOk returns a tuple with the VariantCallingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantCallingId

`func (o *Ga4ghVariantCalling) HasVariantCallingId() bool`

HasVariantCallingId returns a boolean if a field has been set.

### SetVariantCallingId

`func (o *Ga4ghVariantCalling) SetVariantCallingId(v string)`

SetVariantCallingId gets a reference to the given string and assigns it to the VariantCallingId field.

### GetAlignmentId

`func (o *Ga4ghVariantCalling) GetAlignmentId() string`

GetAlignmentId returns the AlignmentId field if non-nil, zero value otherwise.

### GetAlignmentIdOk

`func (o *Ga4ghVariantCalling) GetAlignmentIdOk() (string, bool)`

GetAlignmentIdOk returns a tuple with the AlignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignmentId

`func (o *Ga4ghVariantCalling) HasAlignmentId() bool`

HasAlignmentId returns a boolean if a field has been set.

### SetAlignmentId

`func (o *Ga4ghVariantCalling) SetAlignmentId(v string)`

SetAlignmentId gets a reference to the given string and assigns it to the AlignmentId field.

### GetSite

`func (o *Ga4ghVariantCalling) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Ga4ghVariantCalling) GetSiteOk() (string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSite

`func (o *Ga4ghVariantCalling) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSite

`func (o *Ga4ghVariantCalling) SetSite(v string)`

SetSite gets a reference to the given string and assigns it to the Site field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


