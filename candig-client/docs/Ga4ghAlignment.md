# Ga4ghAlignment

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
**AlignmentTool** | Pointer to **string** |  | [optional] 
**MergeTool** | Pointer to **string** |  | [optional] 
**MarkDuplicates** | Pointer to **string** |  | [optional] 
**RealignerTarget** | Pointer to **string** |  | [optional] 
**IndelRealigner** | Pointer to **string** |  | [optional] 
**BaseRecalibrator** | Pointer to **string** |  | [optional] 
**PrintReads** | Pointer to **string** |  | [optional] 
**IdxStats** | Pointer to **string** |  | [optional] 
**FlagStat** | Pointer to **string** |  | [optional] 
**Coverage** | Pointer to **string** |  | [optional] 
**InsertSizeMetrics** | Pointer to **string** |  | [optional] 
**Fastqc** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**AlignmentId** | Pointer to **string** |  | [optional] 
**SequencingId** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghAlignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghAlignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghAlignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghAlignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghAlignment) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghAlignment) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghAlignment) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghAlignment) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghAlignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghAlignment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghAlignment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghAlignment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghAlignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghAlignment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghAlignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghAlignment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghAlignment) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghAlignment) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghAlignment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghAlignment) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghAlignment) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghAlignment) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghAlignment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghAlignment) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetSampleId

`func (o *Ga4ghAlignment) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghAlignment) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghAlignment) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghAlignment) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetInHousePipeline

`func (o *Ga4ghAlignment) GetInHousePipeline() string`

GetInHousePipeline returns the InHousePipeline field if non-nil, zero value otherwise.

### GetInHousePipelineOk

`func (o *Ga4ghAlignment) GetInHousePipelineOk() (string, bool)`

GetInHousePipelineOk returns a tuple with the InHousePipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInHousePipeline

`func (o *Ga4ghAlignment) HasInHousePipeline() bool`

HasInHousePipeline returns a boolean if a field has been set.

### SetInHousePipeline

`func (o *Ga4ghAlignment) SetInHousePipeline(v string)`

SetInHousePipeline gets a reference to the given string and assigns it to the InHousePipeline field.

### GetAlignmentTool

`func (o *Ga4ghAlignment) GetAlignmentTool() string`

GetAlignmentTool returns the AlignmentTool field if non-nil, zero value otherwise.

### GetAlignmentToolOk

`func (o *Ga4ghAlignment) GetAlignmentToolOk() (string, bool)`

GetAlignmentToolOk returns a tuple with the AlignmentTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignmentTool

`func (o *Ga4ghAlignment) HasAlignmentTool() bool`

HasAlignmentTool returns a boolean if a field has been set.

### SetAlignmentTool

`func (o *Ga4ghAlignment) SetAlignmentTool(v string)`

SetAlignmentTool gets a reference to the given string and assigns it to the AlignmentTool field.

### GetMergeTool

`func (o *Ga4ghAlignment) GetMergeTool() string`

GetMergeTool returns the MergeTool field if non-nil, zero value otherwise.

### GetMergeToolOk

`func (o *Ga4ghAlignment) GetMergeToolOk() (string, bool)`

GetMergeToolOk returns a tuple with the MergeTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMergeTool

`func (o *Ga4ghAlignment) HasMergeTool() bool`

HasMergeTool returns a boolean if a field has been set.

### SetMergeTool

`func (o *Ga4ghAlignment) SetMergeTool(v string)`

SetMergeTool gets a reference to the given string and assigns it to the MergeTool field.

### GetMarkDuplicates

`func (o *Ga4ghAlignment) GetMarkDuplicates() string`

GetMarkDuplicates returns the MarkDuplicates field if non-nil, zero value otherwise.

### GetMarkDuplicatesOk

`func (o *Ga4ghAlignment) GetMarkDuplicatesOk() (string, bool)`

GetMarkDuplicatesOk returns a tuple with the MarkDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMarkDuplicates

`func (o *Ga4ghAlignment) HasMarkDuplicates() bool`

HasMarkDuplicates returns a boolean if a field has been set.

### SetMarkDuplicates

`func (o *Ga4ghAlignment) SetMarkDuplicates(v string)`

SetMarkDuplicates gets a reference to the given string and assigns it to the MarkDuplicates field.

### GetRealignerTarget

`func (o *Ga4ghAlignment) GetRealignerTarget() string`

GetRealignerTarget returns the RealignerTarget field if non-nil, zero value otherwise.

### GetRealignerTargetOk

`func (o *Ga4ghAlignment) GetRealignerTargetOk() (string, bool)`

GetRealignerTargetOk returns a tuple with the RealignerTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRealignerTarget

`func (o *Ga4ghAlignment) HasRealignerTarget() bool`

HasRealignerTarget returns a boolean if a field has been set.

### SetRealignerTarget

`func (o *Ga4ghAlignment) SetRealignerTarget(v string)`

SetRealignerTarget gets a reference to the given string and assigns it to the RealignerTarget field.

### GetIndelRealigner

`func (o *Ga4ghAlignment) GetIndelRealigner() string`

GetIndelRealigner returns the IndelRealigner field if non-nil, zero value otherwise.

### GetIndelRealignerOk

`func (o *Ga4ghAlignment) GetIndelRealignerOk() (string, bool)`

GetIndelRealignerOk returns a tuple with the IndelRealigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndelRealigner

`func (o *Ga4ghAlignment) HasIndelRealigner() bool`

HasIndelRealigner returns a boolean if a field has been set.

### SetIndelRealigner

`func (o *Ga4ghAlignment) SetIndelRealigner(v string)`

SetIndelRealigner gets a reference to the given string and assigns it to the IndelRealigner field.

### GetBaseRecalibrator

`func (o *Ga4ghAlignment) GetBaseRecalibrator() string`

GetBaseRecalibrator returns the BaseRecalibrator field if non-nil, zero value otherwise.

### GetBaseRecalibratorOk

`func (o *Ga4ghAlignment) GetBaseRecalibratorOk() (string, bool)`

GetBaseRecalibratorOk returns a tuple with the BaseRecalibrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBaseRecalibrator

`func (o *Ga4ghAlignment) HasBaseRecalibrator() bool`

HasBaseRecalibrator returns a boolean if a field has been set.

### SetBaseRecalibrator

`func (o *Ga4ghAlignment) SetBaseRecalibrator(v string)`

SetBaseRecalibrator gets a reference to the given string and assigns it to the BaseRecalibrator field.

### GetPrintReads

`func (o *Ga4ghAlignment) GetPrintReads() string`

GetPrintReads returns the PrintReads field if non-nil, zero value otherwise.

### GetPrintReadsOk

`func (o *Ga4ghAlignment) GetPrintReadsOk() (string, bool)`

GetPrintReadsOk returns a tuple with the PrintReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrintReads

`func (o *Ga4ghAlignment) HasPrintReads() bool`

HasPrintReads returns a boolean if a field has been set.

### SetPrintReads

`func (o *Ga4ghAlignment) SetPrintReads(v string)`

SetPrintReads gets a reference to the given string and assigns it to the PrintReads field.

### GetIdxStats

`func (o *Ga4ghAlignment) GetIdxStats() string`

GetIdxStats returns the IdxStats field if non-nil, zero value otherwise.

### GetIdxStatsOk

`func (o *Ga4ghAlignment) GetIdxStatsOk() (string, bool)`

GetIdxStatsOk returns a tuple with the IdxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdxStats

`func (o *Ga4ghAlignment) HasIdxStats() bool`

HasIdxStats returns a boolean if a field has been set.

### SetIdxStats

`func (o *Ga4ghAlignment) SetIdxStats(v string)`

SetIdxStats gets a reference to the given string and assigns it to the IdxStats field.

### GetFlagStat

`func (o *Ga4ghAlignment) GetFlagStat() string`

GetFlagStat returns the FlagStat field if non-nil, zero value otherwise.

### GetFlagStatOk

`func (o *Ga4ghAlignment) GetFlagStatOk() (string, bool)`

GetFlagStatOk returns a tuple with the FlagStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlagStat

`func (o *Ga4ghAlignment) HasFlagStat() bool`

HasFlagStat returns a boolean if a field has been set.

### SetFlagStat

`func (o *Ga4ghAlignment) SetFlagStat(v string)`

SetFlagStat gets a reference to the given string and assigns it to the FlagStat field.

### GetCoverage

`func (o *Ga4ghAlignment) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *Ga4ghAlignment) GetCoverageOk() (string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoverage

`func (o *Ga4ghAlignment) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### SetCoverage

`func (o *Ga4ghAlignment) SetCoverage(v string)`

SetCoverage gets a reference to the given string and assigns it to the Coverage field.

### GetInsertSizeMetrics

`func (o *Ga4ghAlignment) GetInsertSizeMetrics() string`

GetInsertSizeMetrics returns the InsertSizeMetrics field if non-nil, zero value otherwise.

### GetInsertSizeMetricsOk

`func (o *Ga4ghAlignment) GetInsertSizeMetricsOk() (string, bool)`

GetInsertSizeMetricsOk returns a tuple with the InsertSizeMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInsertSizeMetrics

`func (o *Ga4ghAlignment) HasInsertSizeMetrics() bool`

HasInsertSizeMetrics returns a boolean if a field has been set.

### SetInsertSizeMetrics

`func (o *Ga4ghAlignment) SetInsertSizeMetrics(v string)`

SetInsertSizeMetrics gets a reference to the given string and assigns it to the InsertSizeMetrics field.

### GetFastqc

`func (o *Ga4ghAlignment) GetFastqc() string`

GetFastqc returns the Fastqc field if non-nil, zero value otherwise.

### GetFastqcOk

`func (o *Ga4ghAlignment) GetFastqcOk() (string, bool)`

GetFastqcOk returns a tuple with the Fastqc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFastqc

`func (o *Ga4ghAlignment) HasFastqc() bool`

HasFastqc returns a boolean if a field has been set.

### SetFastqc

`func (o *Ga4ghAlignment) SetFastqc(v string)`

SetFastqc gets a reference to the given string and assigns it to the Fastqc field.

### GetReference

`func (o *Ga4ghAlignment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Ga4ghAlignment) GetReferenceOk() (string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReference

`func (o *Ga4ghAlignment) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReference

`func (o *Ga4ghAlignment) SetReference(v string)`

SetReference gets a reference to the given string and assigns it to the Reference field.

### GetAlignmentId

`func (o *Ga4ghAlignment) GetAlignmentId() string`

GetAlignmentId returns the AlignmentId field if non-nil, zero value otherwise.

### GetAlignmentIdOk

`func (o *Ga4ghAlignment) GetAlignmentIdOk() (string, bool)`

GetAlignmentIdOk returns a tuple with the AlignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignmentId

`func (o *Ga4ghAlignment) HasAlignmentId() bool`

HasAlignmentId returns a boolean if a field has been set.

### SetAlignmentId

`func (o *Ga4ghAlignment) SetAlignmentId(v string)`

SetAlignmentId gets a reference to the given string and assigns it to the AlignmentId field.

### GetSequencingId

`func (o *Ga4ghAlignment) GetSequencingId() string`

GetSequencingId returns the SequencingId field if non-nil, zero value otherwise.

### GetSequencingIdOk

`func (o *Ga4ghAlignment) GetSequencingIdOk() (string, bool)`

GetSequencingIdOk returns a tuple with the SequencingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSequencingId

`func (o *Ga4ghAlignment) HasSequencingId() bool`

HasSequencingId returns a boolean if a field has been set.

### SetSequencingId

`func (o *Ga4ghAlignment) SetSequencingId(v string)`

SetSequencingId gets a reference to the given string and assigns it to the SequencingId field.

### GetSite

`func (o *Ga4ghAlignment) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Ga4ghAlignment) GetSiteOk() (string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSite

`func (o *Ga4ghAlignment) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSite

`func (o *Ga4ghAlignment) SetSite(v string)`

SetSite gets a reference to the given string and assigns it to the Site field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


