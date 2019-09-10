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

type Ga4ghCelltransplant struct {
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

	CellSource *string `json:"cellSource,omitempty"`

	DonorType *string `json:"donorType,omitempty"`

	TreatmentPlanId *string `json:"treatmentPlanId,omitempty"`

	CourseNumber *string `json:"courseNumber,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghCelltransplant) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghCelltransplant) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghCelltransplant) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghCelltransplant) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghCelltransplant) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghCelltransplant) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghCelltransplant) SetPatientId(v string) {
	o.PatientId = &v
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetStartDateOk() (string, bool) {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Ga4ghCelltransplant) SetStartDate(v string) {
	o.StartDate = &v
}

// GetCellSource returns the CellSource field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetCellSource() string {
	if o == nil || o.CellSource == nil {
		var ret string
		return ret
	}
	return *o.CellSource
}

// GetCellSourceOk returns a tuple with the CellSource field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetCellSourceOk() (string, bool) {
	if o == nil || o.CellSource == nil {
		var ret string
		return ret, false
	}
	return *o.CellSource, true
}

// HasCellSource returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasCellSource() bool {
	if o != nil && o.CellSource != nil {
		return true
	}

	return false
}

// SetCellSource gets a reference to the given string and assigns it to the CellSource field.
func (o *Ga4ghCelltransplant) SetCellSource(v string) {
	o.CellSource = &v
}

// GetDonorType returns the DonorType field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetDonorType() string {
	if o == nil || o.DonorType == nil {
		var ret string
		return ret
	}
	return *o.DonorType
}

// GetDonorTypeOk returns a tuple with the DonorType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetDonorTypeOk() (string, bool) {
	if o == nil || o.DonorType == nil {
		var ret string
		return ret, false
	}
	return *o.DonorType, true
}

// HasDonorType returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasDonorType() bool {
	if o != nil && o.DonorType != nil {
		return true
	}

	return false
}

// SetDonorType gets a reference to the given string and assigns it to the DonorType field.
func (o *Ga4ghCelltransplant) SetDonorType(v string) {
	o.DonorType = &v
}

// GetTreatmentPlanId returns the TreatmentPlanId field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetTreatmentPlanId() string {
	if o == nil || o.TreatmentPlanId == nil {
		var ret string
		return ret
	}
	return *o.TreatmentPlanId
}

// GetTreatmentPlanIdOk returns a tuple with the TreatmentPlanId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetTreatmentPlanIdOk() (string, bool) {
	if o == nil || o.TreatmentPlanId == nil {
		var ret string
		return ret, false
	}
	return *o.TreatmentPlanId, true
}

// HasTreatmentPlanId returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasTreatmentPlanId() bool {
	if o != nil && o.TreatmentPlanId != nil {
		return true
	}

	return false
}

// SetTreatmentPlanId gets a reference to the given string and assigns it to the TreatmentPlanId field.
func (o *Ga4ghCelltransplant) SetTreatmentPlanId(v string) {
	o.TreatmentPlanId = &v
}

// GetCourseNumber returns the CourseNumber field if non-nil, zero value otherwise.
func (o *Ga4ghCelltransplant) GetCourseNumber() string {
	if o == nil || o.CourseNumber == nil {
		var ret string
		return ret
	}
	return *o.CourseNumber
}

// GetCourseNumberOk returns a tuple with the CourseNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCelltransplant) GetCourseNumberOk() (string, bool) {
	if o == nil || o.CourseNumber == nil {
		var ret string
		return ret, false
	}
	return *o.CourseNumber, true
}

// HasCourseNumber returns a boolean if a field has been set.
func (o *Ga4ghCelltransplant) HasCourseNumber() bool {
	if o != nil && o.CourseNumber != nil {
		return true
	}

	return false
}

// SetCourseNumber gets a reference to the given string and assigns it to the CourseNumber field.
func (o *Ga4ghCelltransplant) SetCourseNumber(v string) {
	o.CourseNumber = &v
}


func (o Ga4ghCelltransplant) MarshalJSON() ([]byte, error) {
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
	if o.CellSource != nil {
		toSerialize["cellSource"] = o.CellSource
	}
	if o.DonorType != nil {
		toSerialize["donorType"] = o.DonorType
	}
	if o.TreatmentPlanId != nil {
		toSerialize["treatmentPlanId"] = o.TreatmentPlanId
	}
	if o.CourseNumber != nil {
		toSerialize["courseNumber"] = o.CourseNumber
	}
	return json.Marshal(toSerialize)
}


