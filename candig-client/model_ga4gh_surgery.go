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

type Ga4ghSurgery struct {
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

	StartDate *string `json:"startDate,omitempty"`

	StopDate *string `json:"stopDate,omitempty"`

	SampleId *string `json:"sampleId,omitempty"`

	CollectionTimePoint *string `json:"collectionTimePoint,omitempty"`

	DiagnosisDate *string `json:"diagnosisDate,omitempty"`

	Site *string `json:"site,omitempty"`

	Type *string `json:"type,omitempty"`

	RecordingDate *string `json:"recordingDate,omitempty"`

	TreatmentPlanId *string `json:"treatmentPlanId,omitempty"`

	CourseNumber *string `json:"courseNumber,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghSurgery) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghSurgery) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghSurgery) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghSurgery) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghSurgery) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghSurgery) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghSurgery) SetPatientId(v string) {
	o.PatientId = &v
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetStartDateOk() (string, bool) {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Ga4ghSurgery) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStopDate returns the StopDate field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetStopDate() string {
	if o == nil || o.StopDate == nil {
		var ret string
		return ret
	}
	return *o.StopDate
}

// GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetStopDateOk() (string, bool) {
	if o == nil || o.StopDate == nil {
		var ret string
		return ret, false
	}
	return *o.StopDate, true
}

// HasStopDate returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasStopDate() bool {
	if o != nil && o.StopDate != nil {
		return true
	}

	return false
}

// SetStopDate gets a reference to the given string and assigns it to the StopDate field.
func (o *Ga4ghSurgery) SetStopDate(v string) {
	o.StopDate = &v
}

// GetSampleId returns the SampleId field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetSampleId() string {
	if o == nil || o.SampleId == nil {
		var ret string
		return ret
	}
	return *o.SampleId
}

// GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetSampleIdOk() (string, bool) {
	if o == nil || o.SampleId == nil {
		var ret string
		return ret, false
	}
	return *o.SampleId, true
}

// HasSampleId returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasSampleId() bool {
	if o != nil && o.SampleId != nil {
		return true
	}

	return false
}

// SetSampleId gets a reference to the given string and assigns it to the SampleId field.
func (o *Ga4ghSurgery) SetSampleId(v string) {
	o.SampleId = &v
}

// GetCollectionTimePoint returns the CollectionTimePoint field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetCollectionTimePoint() string {
	if o == nil || o.CollectionTimePoint == nil {
		var ret string
		return ret
	}
	return *o.CollectionTimePoint
}

// GetCollectionTimePointOk returns a tuple with the CollectionTimePoint field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetCollectionTimePointOk() (string, bool) {
	if o == nil || o.CollectionTimePoint == nil {
		var ret string
		return ret, false
	}
	return *o.CollectionTimePoint, true
}

// HasCollectionTimePoint returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasCollectionTimePoint() bool {
	if o != nil && o.CollectionTimePoint != nil {
		return true
	}

	return false
}

// SetCollectionTimePoint gets a reference to the given string and assigns it to the CollectionTimePoint field.
func (o *Ga4ghSurgery) SetCollectionTimePoint(v string) {
	o.CollectionTimePoint = &v
}

// GetDiagnosisDate returns the DiagnosisDate field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetDiagnosisDate() string {
	if o == nil || o.DiagnosisDate == nil {
		var ret string
		return ret
	}
	return *o.DiagnosisDate
}

// GetDiagnosisDateOk returns a tuple with the DiagnosisDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetDiagnosisDateOk() (string, bool) {
	if o == nil || o.DiagnosisDate == nil {
		var ret string
		return ret, false
	}
	return *o.DiagnosisDate, true
}

// HasDiagnosisDate returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasDiagnosisDate() bool {
	if o != nil && o.DiagnosisDate != nil {
		return true
	}

	return false
}

// SetDiagnosisDate gets a reference to the given string and assigns it to the DiagnosisDate field.
func (o *Ga4ghSurgery) SetDiagnosisDate(v string) {
	o.DiagnosisDate = &v
}

// GetSite returns the Site field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetSiteOk() (string, bool) {
	if o == nil || o.Site == nil {
		var ret string
		return ret, false
	}
	return *o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *Ga4ghSurgery) SetSite(v string) {
	o.Site = &v
}

// GetType returns the Type field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Ga4ghSurgery) SetType(v string) {
	o.Type = &v
}

// GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetRecordingDate() string {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret
	}
	return *o.RecordingDate
}

// GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetRecordingDateOk() (string, bool) {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret, false
	}
	return *o.RecordingDate, true
}

// HasRecordingDate returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasRecordingDate() bool {
	if o != nil && o.RecordingDate != nil {
		return true
	}

	return false
}

// SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.
func (o *Ga4ghSurgery) SetRecordingDate(v string) {
	o.RecordingDate = &v
}

// GetTreatmentPlanId returns the TreatmentPlanId field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetTreatmentPlanId() string {
	if o == nil || o.TreatmentPlanId == nil {
		var ret string
		return ret
	}
	return *o.TreatmentPlanId
}

// GetTreatmentPlanIdOk returns a tuple with the TreatmentPlanId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetTreatmentPlanIdOk() (string, bool) {
	if o == nil || o.TreatmentPlanId == nil {
		var ret string
		return ret, false
	}
	return *o.TreatmentPlanId, true
}

// HasTreatmentPlanId returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasTreatmentPlanId() bool {
	if o != nil && o.TreatmentPlanId != nil {
		return true
	}

	return false
}

// SetTreatmentPlanId gets a reference to the given string and assigns it to the TreatmentPlanId field.
func (o *Ga4ghSurgery) SetTreatmentPlanId(v string) {
	o.TreatmentPlanId = &v
}

// GetCourseNumber returns the CourseNumber field if non-nil, zero value otherwise.
func (o *Ga4ghSurgery) GetCourseNumber() string {
	if o == nil || o.CourseNumber == nil {
		var ret string
		return ret
	}
	return *o.CourseNumber
}

// GetCourseNumberOk returns a tuple with the CourseNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSurgery) GetCourseNumberOk() (string, bool) {
	if o == nil || o.CourseNumber == nil {
		var ret string
		return ret, false
	}
	return *o.CourseNumber, true
}

// HasCourseNumber returns a boolean if a field has been set.
func (o *Ga4ghSurgery) HasCourseNumber() bool {
	if o != nil && o.CourseNumber != nil {
		return true
	}

	return false
}

// SetCourseNumber gets a reference to the given string and assigns it to the CourseNumber field.
func (o *Ga4ghSurgery) SetCourseNumber(v string) {
	o.CourseNumber = &v
}


func (o Ga4ghSurgery) MarshalJSON() ([]byte, error) {
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
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.StopDate != nil {
		toSerialize["stopDate"] = o.StopDate
	}
	if o.SampleId != nil {
		toSerialize["sampleId"] = o.SampleId
	}
	if o.CollectionTimePoint != nil {
		toSerialize["collectionTimePoint"] = o.CollectionTimePoint
	}
	if o.DiagnosisDate != nil {
		toSerialize["diagnosisDate"] = o.DiagnosisDate
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.RecordingDate != nil {
		toSerialize["recordingDate"] = o.RecordingDate
	}
	if o.TreatmentPlanId != nil {
		toSerialize["treatmentPlanId"] = o.TreatmentPlanId
	}
	if o.CourseNumber != nil {
		toSerialize["courseNumber"] = o.CourseNumber
	}
	return json.Marshal(toSerialize)
}


