/*
 * CanDIG Services
 *
 * Below is a list of APIs that CanDIG currently supports.<br/><br/>For /search and /count endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/search_count_services_usage.pdf) for instructions and sample queries.<br/>For all metadata and variant services endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/metadata_variants_services_sample_queries.pdf) for sample queries.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"encoding/json"
)

type Ga4ghAlignment struct {
	// This is unique in the context of the server instance.
	Id *string `json:"id,omitempty"`

	// The ID of the dataset this object belongs to.
	DatasetId *string `json:"dataset_id,omitempty"`

	// This is a label or symbolic identifier for this object.
	Name *string `json:"name,omitempty"`

	// This attribute contains human readable text.
	Description *string `json:"description,omitempty"`

	// The :ref:`ISO 8601<metadata_date_time>` time at which this record was created.
	Created *string `json:"created,omitempty"`

	// The :ref:`ISO 8601<metadata_date_time>` time at which this record was updated.
	Updated *string `json:"updated,omitempty"`

	SampleId *string `json:"sampleId,omitempty"`

	InHousePipeline *string `json:"inHousePipeline,omitempty"`

	AlignmentTool *string `json:"alignmentTool,omitempty"`

	MergeTool *string `json:"mergeTool,omitempty"`

	MarkDuplicates *string `json:"markDuplicates,omitempty"`

	RealignerTarget *string `json:"realignerTarget,omitempty"`

	IndelRealigner *string `json:"indelRealigner,omitempty"`

	BaseRecalibrator *string `json:"baseRecalibrator,omitempty"`

	PrintReads *string `json:"printReads,omitempty"`

	IdxStats *string `json:"idxStats,omitempty"`

	FlagStat *string `json:"flagStat,omitempty"`

	Coverage *string `json:"coverage,omitempty"`

	InsertSizeMetrics *string `json:"insertSizeMetrics,omitempty"`

	Fastqc *string `json:"fastqc,omitempty"`

	Reference *string `json:"reference,omitempty"`

	AlignmentId *string `json:"alignmentId,omitempty"`

	SequencingId *string `json:"sequencingId,omitempty"`

	Site *string `json:"site,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghAlignment) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghAlignment) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghAlignment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghAlignment) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghAlignment) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghAlignment) SetUpdated(v string) {
	o.Updated = &v
}

// GetSampleId returns the SampleId field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetSampleId() string {
	if o == nil || o.SampleId == nil {
		var ret string
		return ret
	}
	return *o.SampleId
}

// GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetSampleIdOk() (string, bool) {
	if o == nil || o.SampleId == nil {
		var ret string
		return ret, false
	}
	return *o.SampleId, true
}

// HasSampleId returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasSampleId() bool {
	if o != nil && o.SampleId != nil {
		return true
	}

	return false
}

// SetSampleId gets a reference to the given string and assigns it to the SampleId field.
func (o *Ga4ghAlignment) SetSampleId(v string) {
	o.SampleId = &v
}

// GetInHousePipeline returns the InHousePipeline field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetInHousePipeline() string {
	if o == nil || o.InHousePipeline == nil {
		var ret string
		return ret
	}
	return *o.InHousePipeline
}

// GetInHousePipelineOk returns a tuple with the InHousePipeline field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetInHousePipelineOk() (string, bool) {
	if o == nil || o.InHousePipeline == nil {
		var ret string
		return ret, false
	}
	return *o.InHousePipeline, true
}

// HasInHousePipeline returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasInHousePipeline() bool {
	if o != nil && o.InHousePipeline != nil {
		return true
	}

	return false
}

// SetInHousePipeline gets a reference to the given string and assigns it to the InHousePipeline field.
func (o *Ga4ghAlignment) SetInHousePipeline(v string) {
	o.InHousePipeline = &v
}

// GetAlignmentTool returns the AlignmentTool field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetAlignmentTool() string {
	if o == nil || o.AlignmentTool == nil {
		var ret string
		return ret
	}
	return *o.AlignmentTool
}

// GetAlignmentToolOk returns a tuple with the AlignmentTool field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetAlignmentToolOk() (string, bool) {
	if o == nil || o.AlignmentTool == nil {
		var ret string
		return ret, false
	}
	return *o.AlignmentTool, true
}

// HasAlignmentTool returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasAlignmentTool() bool {
	if o != nil && o.AlignmentTool != nil {
		return true
	}

	return false
}

// SetAlignmentTool gets a reference to the given string and assigns it to the AlignmentTool field.
func (o *Ga4ghAlignment) SetAlignmentTool(v string) {
	o.AlignmentTool = &v
}

// GetMergeTool returns the MergeTool field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetMergeTool() string {
	if o == nil || o.MergeTool == nil {
		var ret string
		return ret
	}
	return *o.MergeTool
}

// GetMergeToolOk returns a tuple with the MergeTool field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetMergeToolOk() (string, bool) {
	if o == nil || o.MergeTool == nil {
		var ret string
		return ret, false
	}
	return *o.MergeTool, true
}

// HasMergeTool returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasMergeTool() bool {
	if o != nil && o.MergeTool != nil {
		return true
	}

	return false
}

// SetMergeTool gets a reference to the given string and assigns it to the MergeTool field.
func (o *Ga4ghAlignment) SetMergeTool(v string) {
	o.MergeTool = &v
}

// GetMarkDuplicates returns the MarkDuplicates field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetMarkDuplicates() string {
	if o == nil || o.MarkDuplicates == nil {
		var ret string
		return ret
	}
	return *o.MarkDuplicates
}

// GetMarkDuplicatesOk returns a tuple with the MarkDuplicates field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetMarkDuplicatesOk() (string, bool) {
	if o == nil || o.MarkDuplicates == nil {
		var ret string
		return ret, false
	}
	return *o.MarkDuplicates, true
}

// HasMarkDuplicates returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasMarkDuplicates() bool {
	if o != nil && o.MarkDuplicates != nil {
		return true
	}

	return false
}

// SetMarkDuplicates gets a reference to the given string and assigns it to the MarkDuplicates field.
func (o *Ga4ghAlignment) SetMarkDuplicates(v string) {
	o.MarkDuplicates = &v
}

// GetRealignerTarget returns the RealignerTarget field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetRealignerTarget() string {
	if o == nil || o.RealignerTarget == nil {
		var ret string
		return ret
	}
	return *o.RealignerTarget
}

// GetRealignerTargetOk returns a tuple with the RealignerTarget field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetRealignerTargetOk() (string, bool) {
	if o == nil || o.RealignerTarget == nil {
		var ret string
		return ret, false
	}
	return *o.RealignerTarget, true
}

// HasRealignerTarget returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasRealignerTarget() bool {
	if o != nil && o.RealignerTarget != nil {
		return true
	}

	return false
}

// SetRealignerTarget gets a reference to the given string and assigns it to the RealignerTarget field.
func (o *Ga4ghAlignment) SetRealignerTarget(v string) {
	o.RealignerTarget = &v
}

// GetIndelRealigner returns the IndelRealigner field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetIndelRealigner() string {
	if o == nil || o.IndelRealigner == nil {
		var ret string
		return ret
	}
	return *o.IndelRealigner
}

// GetIndelRealignerOk returns a tuple with the IndelRealigner field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetIndelRealignerOk() (string, bool) {
	if o == nil || o.IndelRealigner == nil {
		var ret string
		return ret, false
	}
	return *o.IndelRealigner, true
}

// HasIndelRealigner returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasIndelRealigner() bool {
	if o != nil && o.IndelRealigner != nil {
		return true
	}

	return false
}

// SetIndelRealigner gets a reference to the given string and assigns it to the IndelRealigner field.
func (o *Ga4ghAlignment) SetIndelRealigner(v string) {
	o.IndelRealigner = &v
}

// GetBaseRecalibrator returns the BaseRecalibrator field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetBaseRecalibrator() string {
	if o == nil || o.BaseRecalibrator == nil {
		var ret string
		return ret
	}
	return *o.BaseRecalibrator
}

// GetBaseRecalibratorOk returns a tuple with the BaseRecalibrator field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetBaseRecalibratorOk() (string, bool) {
	if o == nil || o.BaseRecalibrator == nil {
		var ret string
		return ret, false
	}
	return *o.BaseRecalibrator, true
}

// HasBaseRecalibrator returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasBaseRecalibrator() bool {
	if o != nil && o.BaseRecalibrator != nil {
		return true
	}

	return false
}

// SetBaseRecalibrator gets a reference to the given string and assigns it to the BaseRecalibrator field.
func (o *Ga4ghAlignment) SetBaseRecalibrator(v string) {
	o.BaseRecalibrator = &v
}

// GetPrintReads returns the PrintReads field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetPrintReads() string {
	if o == nil || o.PrintReads == nil {
		var ret string
		return ret
	}
	return *o.PrintReads
}

// GetPrintReadsOk returns a tuple with the PrintReads field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetPrintReadsOk() (string, bool) {
	if o == nil || o.PrintReads == nil {
		var ret string
		return ret, false
	}
	return *o.PrintReads, true
}

// HasPrintReads returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasPrintReads() bool {
	if o != nil && o.PrintReads != nil {
		return true
	}

	return false
}

// SetPrintReads gets a reference to the given string and assigns it to the PrintReads field.
func (o *Ga4ghAlignment) SetPrintReads(v string) {
	o.PrintReads = &v
}

// GetIdxStats returns the IdxStats field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetIdxStats() string {
	if o == nil || o.IdxStats == nil {
		var ret string
		return ret
	}
	return *o.IdxStats
}

// GetIdxStatsOk returns a tuple with the IdxStats field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetIdxStatsOk() (string, bool) {
	if o == nil || o.IdxStats == nil {
		var ret string
		return ret, false
	}
	return *o.IdxStats, true
}

// HasIdxStats returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasIdxStats() bool {
	if o != nil && o.IdxStats != nil {
		return true
	}

	return false
}

// SetIdxStats gets a reference to the given string and assigns it to the IdxStats field.
func (o *Ga4ghAlignment) SetIdxStats(v string) {
	o.IdxStats = &v
}

// GetFlagStat returns the FlagStat field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetFlagStat() string {
	if o == nil || o.FlagStat == nil {
		var ret string
		return ret
	}
	return *o.FlagStat
}

// GetFlagStatOk returns a tuple with the FlagStat field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetFlagStatOk() (string, bool) {
	if o == nil || o.FlagStat == nil {
		var ret string
		return ret, false
	}
	return *o.FlagStat, true
}

// HasFlagStat returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasFlagStat() bool {
	if o != nil && o.FlagStat != nil {
		return true
	}

	return false
}

// SetFlagStat gets a reference to the given string and assigns it to the FlagStat field.
func (o *Ga4ghAlignment) SetFlagStat(v string) {
	o.FlagStat = &v
}

// GetCoverage returns the Coverage field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetCoverage() string {
	if o == nil || o.Coverage == nil {
		var ret string
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetCoverageOk() (string, bool) {
	if o == nil || o.Coverage == nil {
		var ret string
		return ret, false
	}
	return *o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasCoverage() bool {
	if o != nil && o.Coverage != nil {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given string and assigns it to the Coverage field.
func (o *Ga4ghAlignment) SetCoverage(v string) {
	o.Coverage = &v
}

// GetInsertSizeMetrics returns the InsertSizeMetrics field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetInsertSizeMetrics() string {
	if o == nil || o.InsertSizeMetrics == nil {
		var ret string
		return ret
	}
	return *o.InsertSizeMetrics
}

// GetInsertSizeMetricsOk returns a tuple with the InsertSizeMetrics field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetInsertSizeMetricsOk() (string, bool) {
	if o == nil || o.InsertSizeMetrics == nil {
		var ret string
		return ret, false
	}
	return *o.InsertSizeMetrics, true
}

// HasInsertSizeMetrics returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasInsertSizeMetrics() bool {
	if o != nil && o.InsertSizeMetrics != nil {
		return true
	}

	return false
}

// SetInsertSizeMetrics gets a reference to the given string and assigns it to the InsertSizeMetrics field.
func (o *Ga4ghAlignment) SetInsertSizeMetrics(v string) {
	o.InsertSizeMetrics = &v
}

// GetFastqc returns the Fastqc field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetFastqc() string {
	if o == nil || o.Fastqc == nil {
		var ret string
		return ret
	}
	return *o.Fastqc
}

// GetFastqcOk returns a tuple with the Fastqc field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetFastqcOk() (string, bool) {
	if o == nil || o.Fastqc == nil {
		var ret string
		return ret, false
	}
	return *o.Fastqc, true
}

// HasFastqc returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasFastqc() bool {
	if o != nil && o.Fastqc != nil {
		return true
	}

	return false
}

// SetFastqc gets a reference to the given string and assigns it to the Fastqc field.
func (o *Ga4ghAlignment) SetFastqc(v string) {
	o.Fastqc = &v
}

// GetReference returns the Reference field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetReferenceOk() (string, bool) {
	if o == nil || o.Reference == nil {
		var ret string
		return ret, false
	}
	return *o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Ga4ghAlignment) SetReference(v string) {
	o.Reference = &v
}

// GetAlignmentId returns the AlignmentId field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetAlignmentId() string {
	if o == nil || o.AlignmentId == nil {
		var ret string
		return ret
	}
	return *o.AlignmentId
}

// GetAlignmentIdOk returns a tuple with the AlignmentId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetAlignmentIdOk() (string, bool) {
	if o == nil || o.AlignmentId == nil {
		var ret string
		return ret, false
	}
	return *o.AlignmentId, true
}

// HasAlignmentId returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasAlignmentId() bool {
	if o != nil && o.AlignmentId != nil {
		return true
	}

	return false
}

// SetAlignmentId gets a reference to the given string and assigns it to the AlignmentId field.
func (o *Ga4ghAlignment) SetAlignmentId(v string) {
	o.AlignmentId = &v
}

// GetSequencingId returns the SequencingId field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetSequencingId() string {
	if o == nil || o.SequencingId == nil {
		var ret string
		return ret
	}
	return *o.SequencingId
}

// GetSequencingIdOk returns a tuple with the SequencingId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetSequencingIdOk() (string, bool) {
	if o == nil || o.SequencingId == nil {
		var ret string
		return ret, false
	}
	return *o.SequencingId, true
}

// HasSequencingId returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasSequencingId() bool {
	if o != nil && o.SequencingId != nil {
		return true
	}

	return false
}

// SetSequencingId gets a reference to the given string and assigns it to the SequencingId field.
func (o *Ga4ghAlignment) SetSequencingId(v string) {
	o.SequencingId = &v
}

// GetSite returns the Site field if non-nil, zero value otherwise.
func (o *Ga4ghAlignment) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghAlignment) GetSiteOk() (string, bool) {
	if o == nil || o.Site == nil {
		var ret string
		return ret, false
	}
	return *o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Ga4ghAlignment) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *Ga4ghAlignment) SetSite(v string) {
	o.Site = &v
}


func (o Ga4ghAlignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.SampleId != nil {
		toSerialize["sampleId"] = o.SampleId
	}
	if o.InHousePipeline != nil {
		toSerialize["inHousePipeline"] = o.InHousePipeline
	}
	if o.AlignmentTool != nil {
		toSerialize["alignmentTool"] = o.AlignmentTool
	}
	if o.MergeTool != nil {
		toSerialize["mergeTool"] = o.MergeTool
	}
	if o.MarkDuplicates != nil {
		toSerialize["markDuplicates"] = o.MarkDuplicates
	}
	if o.RealignerTarget != nil {
		toSerialize["realignerTarget"] = o.RealignerTarget
	}
	if o.IndelRealigner != nil {
		toSerialize["indelRealigner"] = o.IndelRealigner
	}
	if o.BaseRecalibrator != nil {
		toSerialize["baseRecalibrator"] = o.BaseRecalibrator
	}
	if o.PrintReads != nil {
		toSerialize["printReads"] = o.PrintReads
	}
	if o.IdxStats != nil {
		toSerialize["idxStats"] = o.IdxStats
	}
	if o.FlagStat != nil {
		toSerialize["flagStat"] = o.FlagStat
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.InsertSizeMetrics != nil {
		toSerialize["insertSizeMetrics"] = o.InsertSizeMetrics
	}
	if o.Fastqc != nil {
		toSerialize["fastqc"] = o.Fastqc
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.AlignmentId != nil {
		toSerialize["alignmentId"] = o.AlignmentId
	}
	if o.SequencingId != nil {
		toSerialize["sequencingId"] = o.SequencingId
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	return json.Marshal(toSerialize)
}

