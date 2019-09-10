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

type Ga4ghOutcome struct {
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

	PatientId *string `json:"patientId,omitempty"`

	PhysicalExamId *string `json:"physicalExamId,omitempty"`

	DateOfAssessment *string `json:"dateOfAssessment,omitempty"`

	DiseaseResponseOrStatus *string `json:"diseaseResponseOrStatus,omitempty"`

	OtherResponseClassification *string `json:"otherResponseClassification,omitempty"`

	MinimalResidualDiseaseAssessment *string `json:"minimalResidualDiseaseAssessment,omitempty"`

	MethodOfResponseEvaluation *string `json:"methodOfResponseEvaluation,omitempty"`

	ResponseCriteriaUsed *string `json:"responseCriteriaUsed,omitempty"`

	SummaryStage *string `json:"summaryStage,omitempty"`

	SitesOfAnyProgressionOrRecurrence *string `json:"sitesOfAnyProgressionOrRecurrence,omitempty"`

	VitalStatus *string `json:"vitalStatus,omitempty"`

	Height *string `json:"height,omitempty"`

	Weight *string `json:"weight,omitempty"`

	HeightUnits *string `json:"heightUnits,omitempty"`

	WeightUnits *string `json:"weightUnits,omitempty"`

	PerformanceStatus *string `json:"performanceStatus,omitempty"`

	OverallSurvivalInMonths *string `json:"overallSurvivalInMonths,omitempty"`

	DiseaseFreeSurvivalInMonths *string `json:"diseaseFreeSurvivalInMonths,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghOutcome) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghOutcome) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghOutcome) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghOutcome) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghOutcome) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghOutcome) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghOutcome) SetPatientId(v string) {
	o.PatientId = &v
}

// GetPhysicalExamId returns the PhysicalExamId field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetPhysicalExamId() string {
	if o == nil || o.PhysicalExamId == nil {
		var ret string
		return ret
	}
	return *o.PhysicalExamId
}

// GetPhysicalExamIdOk returns a tuple with the PhysicalExamId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetPhysicalExamIdOk() (string, bool) {
	if o == nil || o.PhysicalExamId == nil {
		var ret string
		return ret, false
	}
	return *o.PhysicalExamId, true
}

// HasPhysicalExamId returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasPhysicalExamId() bool {
	if o != nil && o.PhysicalExamId != nil {
		return true
	}

	return false
}

// SetPhysicalExamId gets a reference to the given string and assigns it to the PhysicalExamId field.
func (o *Ga4ghOutcome) SetPhysicalExamId(v string) {
	o.PhysicalExamId = &v
}

// GetDateOfAssessment returns the DateOfAssessment field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetDateOfAssessment() string {
	if o == nil || o.DateOfAssessment == nil {
		var ret string
		return ret
	}
	return *o.DateOfAssessment
}

// GetDateOfAssessmentOk returns a tuple with the DateOfAssessment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetDateOfAssessmentOk() (string, bool) {
	if o == nil || o.DateOfAssessment == nil {
		var ret string
		return ret, false
	}
	return *o.DateOfAssessment, true
}

// HasDateOfAssessment returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasDateOfAssessment() bool {
	if o != nil && o.DateOfAssessment != nil {
		return true
	}

	return false
}

// SetDateOfAssessment gets a reference to the given string and assigns it to the DateOfAssessment field.
func (o *Ga4ghOutcome) SetDateOfAssessment(v string) {
	o.DateOfAssessment = &v
}

// GetDiseaseResponseOrStatus returns the DiseaseResponseOrStatus field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetDiseaseResponseOrStatus() string {
	if o == nil || o.DiseaseResponseOrStatus == nil {
		var ret string
		return ret
	}
	return *o.DiseaseResponseOrStatus
}

// GetDiseaseResponseOrStatusOk returns a tuple with the DiseaseResponseOrStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetDiseaseResponseOrStatusOk() (string, bool) {
	if o == nil || o.DiseaseResponseOrStatus == nil {
		var ret string
		return ret, false
	}
	return *o.DiseaseResponseOrStatus, true
}

// HasDiseaseResponseOrStatus returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasDiseaseResponseOrStatus() bool {
	if o != nil && o.DiseaseResponseOrStatus != nil {
		return true
	}

	return false
}

// SetDiseaseResponseOrStatus gets a reference to the given string and assigns it to the DiseaseResponseOrStatus field.
func (o *Ga4ghOutcome) SetDiseaseResponseOrStatus(v string) {
	o.DiseaseResponseOrStatus = &v
}

// GetOtherResponseClassification returns the OtherResponseClassification field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetOtherResponseClassification() string {
	if o == nil || o.OtherResponseClassification == nil {
		var ret string
		return ret
	}
	return *o.OtherResponseClassification
}

// GetOtherResponseClassificationOk returns a tuple with the OtherResponseClassification field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetOtherResponseClassificationOk() (string, bool) {
	if o == nil || o.OtherResponseClassification == nil {
		var ret string
		return ret, false
	}
	return *o.OtherResponseClassification, true
}

// HasOtherResponseClassification returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasOtherResponseClassification() bool {
	if o != nil && o.OtherResponseClassification != nil {
		return true
	}

	return false
}

// SetOtherResponseClassification gets a reference to the given string and assigns it to the OtherResponseClassification field.
func (o *Ga4ghOutcome) SetOtherResponseClassification(v string) {
	o.OtherResponseClassification = &v
}

// GetMinimalResidualDiseaseAssessment returns the MinimalResidualDiseaseAssessment field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetMinimalResidualDiseaseAssessment() string {
	if o == nil || o.MinimalResidualDiseaseAssessment == nil {
		var ret string
		return ret
	}
	return *o.MinimalResidualDiseaseAssessment
}

// GetMinimalResidualDiseaseAssessmentOk returns a tuple with the MinimalResidualDiseaseAssessment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetMinimalResidualDiseaseAssessmentOk() (string, bool) {
	if o == nil || o.MinimalResidualDiseaseAssessment == nil {
		var ret string
		return ret, false
	}
	return *o.MinimalResidualDiseaseAssessment, true
}

// HasMinimalResidualDiseaseAssessment returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasMinimalResidualDiseaseAssessment() bool {
	if o != nil && o.MinimalResidualDiseaseAssessment != nil {
		return true
	}

	return false
}

// SetMinimalResidualDiseaseAssessment gets a reference to the given string and assigns it to the MinimalResidualDiseaseAssessment field.
func (o *Ga4ghOutcome) SetMinimalResidualDiseaseAssessment(v string) {
	o.MinimalResidualDiseaseAssessment = &v
}

// GetMethodOfResponseEvaluation returns the MethodOfResponseEvaluation field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetMethodOfResponseEvaluation() string {
	if o == nil || o.MethodOfResponseEvaluation == nil {
		var ret string
		return ret
	}
	return *o.MethodOfResponseEvaluation
}

// GetMethodOfResponseEvaluationOk returns a tuple with the MethodOfResponseEvaluation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetMethodOfResponseEvaluationOk() (string, bool) {
	if o == nil || o.MethodOfResponseEvaluation == nil {
		var ret string
		return ret, false
	}
	return *o.MethodOfResponseEvaluation, true
}

// HasMethodOfResponseEvaluation returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasMethodOfResponseEvaluation() bool {
	if o != nil && o.MethodOfResponseEvaluation != nil {
		return true
	}

	return false
}

// SetMethodOfResponseEvaluation gets a reference to the given string and assigns it to the MethodOfResponseEvaluation field.
func (o *Ga4ghOutcome) SetMethodOfResponseEvaluation(v string) {
	o.MethodOfResponseEvaluation = &v
}

// GetResponseCriteriaUsed returns the ResponseCriteriaUsed field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetResponseCriteriaUsed() string {
	if o == nil || o.ResponseCriteriaUsed == nil {
		var ret string
		return ret
	}
	return *o.ResponseCriteriaUsed
}

// GetResponseCriteriaUsedOk returns a tuple with the ResponseCriteriaUsed field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetResponseCriteriaUsedOk() (string, bool) {
	if o == nil || o.ResponseCriteriaUsed == nil {
		var ret string
		return ret, false
	}
	return *o.ResponseCriteriaUsed, true
}

// HasResponseCriteriaUsed returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasResponseCriteriaUsed() bool {
	if o != nil && o.ResponseCriteriaUsed != nil {
		return true
	}

	return false
}

// SetResponseCriteriaUsed gets a reference to the given string and assigns it to the ResponseCriteriaUsed field.
func (o *Ga4ghOutcome) SetResponseCriteriaUsed(v string) {
	o.ResponseCriteriaUsed = &v
}

// GetSummaryStage returns the SummaryStage field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetSummaryStage() string {
	if o == nil || o.SummaryStage == nil {
		var ret string
		return ret
	}
	return *o.SummaryStage
}

// GetSummaryStageOk returns a tuple with the SummaryStage field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetSummaryStageOk() (string, bool) {
	if o == nil || o.SummaryStage == nil {
		var ret string
		return ret, false
	}
	return *o.SummaryStage, true
}

// HasSummaryStage returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasSummaryStage() bool {
	if o != nil && o.SummaryStage != nil {
		return true
	}

	return false
}

// SetSummaryStage gets a reference to the given string and assigns it to the SummaryStage field.
func (o *Ga4ghOutcome) SetSummaryStage(v string) {
	o.SummaryStage = &v
}

// GetSitesOfAnyProgressionOrRecurrence returns the SitesOfAnyProgressionOrRecurrence field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetSitesOfAnyProgressionOrRecurrence() string {
	if o == nil || o.SitesOfAnyProgressionOrRecurrence == nil {
		var ret string
		return ret
	}
	return *o.SitesOfAnyProgressionOrRecurrence
}

// GetSitesOfAnyProgressionOrRecurrenceOk returns a tuple with the SitesOfAnyProgressionOrRecurrence field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetSitesOfAnyProgressionOrRecurrenceOk() (string, bool) {
	if o == nil || o.SitesOfAnyProgressionOrRecurrence == nil {
		var ret string
		return ret, false
	}
	return *o.SitesOfAnyProgressionOrRecurrence, true
}

// HasSitesOfAnyProgressionOrRecurrence returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasSitesOfAnyProgressionOrRecurrence() bool {
	if o != nil && o.SitesOfAnyProgressionOrRecurrence != nil {
		return true
	}

	return false
}

// SetSitesOfAnyProgressionOrRecurrence gets a reference to the given string and assigns it to the SitesOfAnyProgressionOrRecurrence field.
func (o *Ga4ghOutcome) SetSitesOfAnyProgressionOrRecurrence(v string) {
	o.SitesOfAnyProgressionOrRecurrence = &v
}

// GetVitalStatus returns the VitalStatus field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetVitalStatus() string {
	if o == nil || o.VitalStatus == nil {
		var ret string
		return ret
	}
	return *o.VitalStatus
}

// GetVitalStatusOk returns a tuple with the VitalStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetVitalStatusOk() (string, bool) {
	if o == nil || o.VitalStatus == nil {
		var ret string
		return ret, false
	}
	return *o.VitalStatus, true
}

// HasVitalStatus returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasVitalStatus() bool {
	if o != nil && o.VitalStatus != nil {
		return true
	}

	return false
}

// SetVitalStatus gets a reference to the given string and assigns it to the VitalStatus field.
func (o *Ga4ghOutcome) SetVitalStatus(v string) {
	o.VitalStatus = &v
}

// GetHeight returns the Height field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetHeight() string {
	if o == nil || o.Height == nil {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetHeightOk() (string, bool) {
	if o == nil || o.Height == nil {
		var ret string
		return ret, false
	}
	return *o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *Ga4ghOutcome) SetHeight(v string) {
	o.Height = &v
}

// GetWeight returns the Weight field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetWeight() string {
	if o == nil || o.Weight == nil {
		var ret string
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetWeightOk() (string, bool) {
	if o == nil || o.Weight == nil {
		var ret string
		return ret, false
	}
	return *o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given string and assigns it to the Weight field.
func (o *Ga4ghOutcome) SetWeight(v string) {
	o.Weight = &v
}

// GetHeightUnits returns the HeightUnits field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetHeightUnits() string {
	if o == nil || o.HeightUnits == nil {
		var ret string
		return ret
	}
	return *o.HeightUnits
}

// GetHeightUnitsOk returns a tuple with the HeightUnits field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetHeightUnitsOk() (string, bool) {
	if o == nil || o.HeightUnits == nil {
		var ret string
		return ret, false
	}
	return *o.HeightUnits, true
}

// HasHeightUnits returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasHeightUnits() bool {
	if o != nil && o.HeightUnits != nil {
		return true
	}

	return false
}

// SetHeightUnits gets a reference to the given string and assigns it to the HeightUnits field.
func (o *Ga4ghOutcome) SetHeightUnits(v string) {
	o.HeightUnits = &v
}

// GetWeightUnits returns the WeightUnits field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetWeightUnits() string {
	if o == nil || o.WeightUnits == nil {
		var ret string
		return ret
	}
	return *o.WeightUnits
}

// GetWeightUnitsOk returns a tuple with the WeightUnits field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetWeightUnitsOk() (string, bool) {
	if o == nil || o.WeightUnits == nil {
		var ret string
		return ret, false
	}
	return *o.WeightUnits, true
}

// HasWeightUnits returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasWeightUnits() bool {
	if o != nil && o.WeightUnits != nil {
		return true
	}

	return false
}

// SetWeightUnits gets a reference to the given string and assigns it to the WeightUnits field.
func (o *Ga4ghOutcome) SetWeightUnits(v string) {
	o.WeightUnits = &v
}

// GetPerformanceStatus returns the PerformanceStatus field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetPerformanceStatus() string {
	if o == nil || o.PerformanceStatus == nil {
		var ret string
		return ret
	}
	return *o.PerformanceStatus
}

// GetPerformanceStatusOk returns a tuple with the PerformanceStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetPerformanceStatusOk() (string, bool) {
	if o == nil || o.PerformanceStatus == nil {
		var ret string
		return ret, false
	}
	return *o.PerformanceStatus, true
}

// HasPerformanceStatus returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasPerformanceStatus() bool {
	if o != nil && o.PerformanceStatus != nil {
		return true
	}

	return false
}

// SetPerformanceStatus gets a reference to the given string and assigns it to the PerformanceStatus field.
func (o *Ga4ghOutcome) SetPerformanceStatus(v string) {
	o.PerformanceStatus = &v
}

// GetOverallSurvivalInMonths returns the OverallSurvivalInMonths field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetOverallSurvivalInMonths() string {
	if o == nil || o.OverallSurvivalInMonths == nil {
		var ret string
		return ret
	}
	return *o.OverallSurvivalInMonths
}

// GetOverallSurvivalInMonthsOk returns a tuple with the OverallSurvivalInMonths field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetOverallSurvivalInMonthsOk() (string, bool) {
	if o == nil || o.OverallSurvivalInMonths == nil {
		var ret string
		return ret, false
	}
	return *o.OverallSurvivalInMonths, true
}

// HasOverallSurvivalInMonths returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasOverallSurvivalInMonths() bool {
	if o != nil && o.OverallSurvivalInMonths != nil {
		return true
	}

	return false
}

// SetOverallSurvivalInMonths gets a reference to the given string and assigns it to the OverallSurvivalInMonths field.
func (o *Ga4ghOutcome) SetOverallSurvivalInMonths(v string) {
	o.OverallSurvivalInMonths = &v
}

// GetDiseaseFreeSurvivalInMonths returns the DiseaseFreeSurvivalInMonths field if non-nil, zero value otherwise.
func (o *Ga4ghOutcome) GetDiseaseFreeSurvivalInMonths() string {
	if o == nil || o.DiseaseFreeSurvivalInMonths == nil {
		var ret string
		return ret
	}
	return *o.DiseaseFreeSurvivalInMonths
}

// GetDiseaseFreeSurvivalInMonthsOk returns a tuple with the DiseaseFreeSurvivalInMonths field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOutcome) GetDiseaseFreeSurvivalInMonthsOk() (string, bool) {
	if o == nil || o.DiseaseFreeSurvivalInMonths == nil {
		var ret string
		return ret, false
	}
	return *o.DiseaseFreeSurvivalInMonths, true
}

// HasDiseaseFreeSurvivalInMonths returns a boolean if a field has been set.
func (o *Ga4ghOutcome) HasDiseaseFreeSurvivalInMonths() bool {
	if o != nil && o.DiseaseFreeSurvivalInMonths != nil {
		return true
	}

	return false
}

// SetDiseaseFreeSurvivalInMonths gets a reference to the given string and assigns it to the DiseaseFreeSurvivalInMonths field.
func (o *Ga4ghOutcome) SetDiseaseFreeSurvivalInMonths(v string) {
	o.DiseaseFreeSurvivalInMonths = &v
}


func (o Ga4ghOutcome) MarshalJSON() ([]byte, error) {
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
	if o.PatientId != nil {
		toSerialize["patientId"] = o.PatientId
	}
	if o.PhysicalExamId != nil {
		toSerialize["physicalExamId"] = o.PhysicalExamId
	}
	if o.DateOfAssessment != nil {
		toSerialize["dateOfAssessment"] = o.DateOfAssessment
	}
	if o.DiseaseResponseOrStatus != nil {
		toSerialize["diseaseResponseOrStatus"] = o.DiseaseResponseOrStatus
	}
	if o.OtherResponseClassification != nil {
		toSerialize["otherResponseClassification"] = o.OtherResponseClassification
	}
	if o.MinimalResidualDiseaseAssessment != nil {
		toSerialize["minimalResidualDiseaseAssessment"] = o.MinimalResidualDiseaseAssessment
	}
	if o.MethodOfResponseEvaluation != nil {
		toSerialize["methodOfResponseEvaluation"] = o.MethodOfResponseEvaluation
	}
	if o.ResponseCriteriaUsed != nil {
		toSerialize["responseCriteriaUsed"] = o.ResponseCriteriaUsed
	}
	if o.SummaryStage != nil {
		toSerialize["summaryStage"] = o.SummaryStage
	}
	if o.SitesOfAnyProgressionOrRecurrence != nil {
		toSerialize["sitesOfAnyProgressionOrRecurrence"] = o.SitesOfAnyProgressionOrRecurrence
	}
	if o.VitalStatus != nil {
		toSerialize["vitalStatus"] = o.VitalStatus
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.HeightUnits != nil {
		toSerialize["heightUnits"] = o.HeightUnits
	}
	if o.WeightUnits != nil {
		toSerialize["weightUnits"] = o.WeightUnits
	}
	if o.PerformanceStatus != nil {
		toSerialize["performanceStatus"] = o.PerformanceStatus
	}
	if o.OverallSurvivalInMonths != nil {
		toSerialize["overallSurvivalInMonths"] = o.OverallSurvivalInMonths
	}
	if o.DiseaseFreeSurvivalInMonths != nil {
		toSerialize["diseaseFreeSurvivalInMonths"] = o.DiseaseFreeSurvivalInMonths
	}
	return json.Marshal(toSerialize)
}


