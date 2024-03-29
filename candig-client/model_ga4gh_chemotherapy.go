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

type Ga4ghChemotherapy struct {
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

	CourseNumber *string `json:"courseNumber,omitempty"`

	StartDate *string `json:"startDate,omitempty"`

	StopDate *string `json:"stopDate,omitempty"`

	SystematicTherapyAgentName *string `json:"systematicTherapyAgentName,omitempty"`

	Route *string `json:"route,omitempty"`

	Dose *string `json:"dose,omitempty"`

	DoseFrequency *string `json:"doseFrequency,omitempty"`

	DoseUnit *string `json:"doseUnit,omitempty"`

	DaysPerCycle *string `json:"daysPerCycle,omitempty"`

	NumberOfCycle *string `json:"numberOfCycle,omitempty"`

	TreatmentIntent *string `json:"treatmentIntent,omitempty"`

	TreatingCentreName *string `json:"treatingCentreName,omitempty"`

	Type *string `json:"type,omitempty"`

	ProtocolCode *string `json:"protocolCode,omitempty"`

	RecordingDate *string `json:"recordingDate,omitempty"`

	TreatmentPlanId *string `json:"treatmentPlanId,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghChemotherapy) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghChemotherapy) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghChemotherapy) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghChemotherapy) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghChemotherapy) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghChemotherapy) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghChemotherapy) SetPatientId(v string) {
	o.PatientId = &v
}

// GetCourseNumber returns the CourseNumber field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetCourseNumber() string {
	if o == nil || o.CourseNumber == nil {
		var ret string
		return ret
	}
	return *o.CourseNumber
}

// GetCourseNumberOk returns a tuple with the CourseNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetCourseNumberOk() (string, bool) {
	if o == nil || o.CourseNumber == nil {
		var ret string
		return ret, false
	}
	return *o.CourseNumber, true
}

// HasCourseNumber returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasCourseNumber() bool {
	if o != nil && o.CourseNumber != nil {
		return true
	}

	return false
}

// SetCourseNumber gets a reference to the given string and assigns it to the CourseNumber field.
func (o *Ga4ghChemotherapy) SetCourseNumber(v string) {
	o.CourseNumber = &v
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetStartDateOk() (string, bool) {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Ga4ghChemotherapy) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStopDate returns the StopDate field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetStopDate() string {
	if o == nil || o.StopDate == nil {
		var ret string
		return ret
	}
	return *o.StopDate
}

// GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetStopDateOk() (string, bool) {
	if o == nil || o.StopDate == nil {
		var ret string
		return ret, false
	}
	return *o.StopDate, true
}

// HasStopDate returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasStopDate() bool {
	if o != nil && o.StopDate != nil {
		return true
	}

	return false
}

// SetStopDate gets a reference to the given string and assigns it to the StopDate field.
func (o *Ga4ghChemotherapy) SetStopDate(v string) {
	o.StopDate = &v
}

// GetSystematicTherapyAgentName returns the SystematicTherapyAgentName field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetSystematicTherapyAgentName() string {
	if o == nil || o.SystematicTherapyAgentName == nil {
		var ret string
		return ret
	}
	return *o.SystematicTherapyAgentName
}

// GetSystematicTherapyAgentNameOk returns a tuple with the SystematicTherapyAgentName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetSystematicTherapyAgentNameOk() (string, bool) {
	if o == nil || o.SystematicTherapyAgentName == nil {
		var ret string
		return ret, false
	}
	return *o.SystematicTherapyAgentName, true
}

// HasSystematicTherapyAgentName returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasSystematicTherapyAgentName() bool {
	if o != nil && o.SystematicTherapyAgentName != nil {
		return true
	}

	return false
}

// SetSystematicTherapyAgentName gets a reference to the given string and assigns it to the SystematicTherapyAgentName field.
func (o *Ga4ghChemotherapy) SetSystematicTherapyAgentName(v string) {
	o.SystematicTherapyAgentName = &v
}

// GetRoute returns the Route field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetRoute() string {
	if o == nil || o.Route == nil {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetRouteOk() (string, bool) {
	if o == nil || o.Route == nil {
		var ret string
		return ret, false
	}
	return *o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasRoute() bool {
	if o != nil && o.Route != nil {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *Ga4ghChemotherapy) SetRoute(v string) {
	o.Route = &v
}

// GetDose returns the Dose field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetDose() string {
	if o == nil || o.Dose == nil {
		var ret string
		return ret
	}
	return *o.Dose
}

// GetDoseOk returns a tuple with the Dose field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetDoseOk() (string, bool) {
	if o == nil || o.Dose == nil {
		var ret string
		return ret, false
	}
	return *o.Dose, true
}

// HasDose returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasDose() bool {
	if o != nil && o.Dose != nil {
		return true
	}

	return false
}

// SetDose gets a reference to the given string and assigns it to the Dose field.
func (o *Ga4ghChemotherapy) SetDose(v string) {
	o.Dose = &v
}

// GetDoseFrequency returns the DoseFrequency field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetDoseFrequency() string {
	if o == nil || o.DoseFrequency == nil {
		var ret string
		return ret
	}
	return *o.DoseFrequency
}

// GetDoseFrequencyOk returns a tuple with the DoseFrequency field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetDoseFrequencyOk() (string, bool) {
	if o == nil || o.DoseFrequency == nil {
		var ret string
		return ret, false
	}
	return *o.DoseFrequency, true
}

// HasDoseFrequency returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasDoseFrequency() bool {
	if o != nil && o.DoseFrequency != nil {
		return true
	}

	return false
}

// SetDoseFrequency gets a reference to the given string and assigns it to the DoseFrequency field.
func (o *Ga4ghChemotherapy) SetDoseFrequency(v string) {
	o.DoseFrequency = &v
}

// GetDoseUnit returns the DoseUnit field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetDoseUnit() string {
	if o == nil || o.DoseUnit == nil {
		var ret string
		return ret
	}
	return *o.DoseUnit
}

// GetDoseUnitOk returns a tuple with the DoseUnit field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetDoseUnitOk() (string, bool) {
	if o == nil || o.DoseUnit == nil {
		var ret string
		return ret, false
	}
	return *o.DoseUnit, true
}

// HasDoseUnit returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasDoseUnit() bool {
	if o != nil && o.DoseUnit != nil {
		return true
	}

	return false
}

// SetDoseUnit gets a reference to the given string and assigns it to the DoseUnit field.
func (o *Ga4ghChemotherapy) SetDoseUnit(v string) {
	o.DoseUnit = &v
}

// GetDaysPerCycle returns the DaysPerCycle field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetDaysPerCycle() string {
	if o == nil || o.DaysPerCycle == nil {
		var ret string
		return ret
	}
	return *o.DaysPerCycle
}

// GetDaysPerCycleOk returns a tuple with the DaysPerCycle field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetDaysPerCycleOk() (string, bool) {
	if o == nil || o.DaysPerCycle == nil {
		var ret string
		return ret, false
	}
	return *o.DaysPerCycle, true
}

// HasDaysPerCycle returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasDaysPerCycle() bool {
	if o != nil && o.DaysPerCycle != nil {
		return true
	}

	return false
}

// SetDaysPerCycle gets a reference to the given string and assigns it to the DaysPerCycle field.
func (o *Ga4ghChemotherapy) SetDaysPerCycle(v string) {
	o.DaysPerCycle = &v
}

// GetNumberOfCycle returns the NumberOfCycle field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetNumberOfCycle() string {
	if o == nil || o.NumberOfCycle == nil {
		var ret string
		return ret
	}
	return *o.NumberOfCycle
}

// GetNumberOfCycleOk returns a tuple with the NumberOfCycle field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetNumberOfCycleOk() (string, bool) {
	if o == nil || o.NumberOfCycle == nil {
		var ret string
		return ret, false
	}
	return *o.NumberOfCycle, true
}

// HasNumberOfCycle returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasNumberOfCycle() bool {
	if o != nil && o.NumberOfCycle != nil {
		return true
	}

	return false
}

// SetNumberOfCycle gets a reference to the given string and assigns it to the NumberOfCycle field.
func (o *Ga4ghChemotherapy) SetNumberOfCycle(v string) {
	o.NumberOfCycle = &v
}

// GetTreatmentIntent returns the TreatmentIntent field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetTreatmentIntent() string {
	if o == nil || o.TreatmentIntent == nil {
		var ret string
		return ret
	}
	return *o.TreatmentIntent
}

// GetTreatmentIntentOk returns a tuple with the TreatmentIntent field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetTreatmentIntentOk() (string, bool) {
	if o == nil || o.TreatmentIntent == nil {
		var ret string
		return ret, false
	}
	return *o.TreatmentIntent, true
}

// HasTreatmentIntent returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasTreatmentIntent() bool {
	if o != nil && o.TreatmentIntent != nil {
		return true
	}

	return false
}

// SetTreatmentIntent gets a reference to the given string and assigns it to the TreatmentIntent field.
func (o *Ga4ghChemotherapy) SetTreatmentIntent(v string) {
	o.TreatmentIntent = &v
}

// GetTreatingCentreName returns the TreatingCentreName field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetTreatingCentreName() string {
	if o == nil || o.TreatingCentreName == nil {
		var ret string
		return ret
	}
	return *o.TreatingCentreName
}

// GetTreatingCentreNameOk returns a tuple with the TreatingCentreName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetTreatingCentreNameOk() (string, bool) {
	if o == nil || o.TreatingCentreName == nil {
		var ret string
		return ret, false
	}
	return *o.TreatingCentreName, true
}

// HasTreatingCentreName returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasTreatingCentreName() bool {
	if o != nil && o.TreatingCentreName != nil {
		return true
	}

	return false
}

// SetTreatingCentreName gets a reference to the given string and assigns it to the TreatingCentreName field.
func (o *Ga4ghChemotherapy) SetTreatingCentreName(v string) {
	o.TreatingCentreName = &v
}

// GetType returns the Type field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Ga4ghChemotherapy) SetType(v string) {
	o.Type = &v
}

// GetProtocolCode returns the ProtocolCode field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetProtocolCode() string {
	if o == nil || o.ProtocolCode == nil {
		var ret string
		return ret
	}
	return *o.ProtocolCode
}

// GetProtocolCodeOk returns a tuple with the ProtocolCode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetProtocolCodeOk() (string, bool) {
	if o == nil || o.ProtocolCode == nil {
		var ret string
		return ret, false
	}
	return *o.ProtocolCode, true
}

// HasProtocolCode returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasProtocolCode() bool {
	if o != nil && o.ProtocolCode != nil {
		return true
	}

	return false
}

// SetProtocolCode gets a reference to the given string and assigns it to the ProtocolCode field.
func (o *Ga4ghChemotherapy) SetProtocolCode(v string) {
	o.ProtocolCode = &v
}

// GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetRecordingDate() string {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret
	}
	return *o.RecordingDate
}

// GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetRecordingDateOk() (string, bool) {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret, false
	}
	return *o.RecordingDate, true
}

// HasRecordingDate returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasRecordingDate() bool {
	if o != nil && o.RecordingDate != nil {
		return true
	}

	return false
}

// SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.
func (o *Ga4ghChemotherapy) SetRecordingDate(v string) {
	o.RecordingDate = &v
}

// GetTreatmentPlanId returns the TreatmentPlanId field if non-nil, zero value otherwise.
func (o *Ga4ghChemotherapy) GetTreatmentPlanId() string {
	if o == nil || o.TreatmentPlanId == nil {
		var ret string
		return ret
	}
	return *o.TreatmentPlanId
}

// GetTreatmentPlanIdOk returns a tuple with the TreatmentPlanId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghChemotherapy) GetTreatmentPlanIdOk() (string, bool) {
	if o == nil || o.TreatmentPlanId == nil {
		var ret string
		return ret, false
	}
	return *o.TreatmentPlanId, true
}

// HasTreatmentPlanId returns a boolean if a field has been set.
func (o *Ga4ghChemotherapy) HasTreatmentPlanId() bool {
	if o != nil && o.TreatmentPlanId != nil {
		return true
	}

	return false
}

// SetTreatmentPlanId gets a reference to the given string and assigns it to the TreatmentPlanId field.
func (o *Ga4ghChemotherapy) SetTreatmentPlanId(v string) {
	o.TreatmentPlanId = &v
}


func (o Ga4ghChemotherapy) MarshalJSON() ([]byte, error) {
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
	if o.CourseNumber != nil {
		toSerialize["courseNumber"] = o.CourseNumber
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.StopDate != nil {
		toSerialize["stopDate"] = o.StopDate
	}
	if o.SystematicTherapyAgentName != nil {
		toSerialize["systematicTherapyAgentName"] = o.SystematicTherapyAgentName
	}
	if o.Route != nil {
		toSerialize["route"] = o.Route
	}
	if o.Dose != nil {
		toSerialize["dose"] = o.Dose
	}
	if o.DoseFrequency != nil {
		toSerialize["doseFrequency"] = o.DoseFrequency
	}
	if o.DoseUnit != nil {
		toSerialize["doseUnit"] = o.DoseUnit
	}
	if o.DaysPerCycle != nil {
		toSerialize["daysPerCycle"] = o.DaysPerCycle
	}
	if o.NumberOfCycle != nil {
		toSerialize["numberOfCycle"] = o.NumberOfCycle
	}
	if o.TreatmentIntent != nil {
		toSerialize["treatmentIntent"] = o.TreatmentIntent
	}
	if o.TreatingCentreName != nil {
		toSerialize["treatingCentreName"] = o.TreatingCentreName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ProtocolCode != nil {
		toSerialize["protocolCode"] = o.ProtocolCode
	}
	if o.RecordingDate != nil {
		toSerialize["recordingDate"] = o.RecordingDate
	}
	if o.TreatmentPlanId != nil {
		toSerialize["treatmentPlanId"] = o.TreatmentPlanId
	}
	return json.Marshal(toSerialize)
}


