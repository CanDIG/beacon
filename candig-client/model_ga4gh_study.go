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

type Ga4ghStudy struct {
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

	EndDate *string `json:"endDate,omitempty"`

	Status *string `json:"status,omitempty"`

	RecordingDate *string `json:"recordingDate,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghStudy) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghStudy) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghStudy) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghStudy) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghStudy) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghStudy) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghStudy) SetPatientId(v string) {
	o.PatientId = &v
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetStartDateOk() (string, bool) {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Ga4ghStudy) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetEndDateOk() (string, bool) {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret, false
	}
	return *o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Ga4ghStudy) SetEndDate(v string) {
	o.EndDate = &v
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Ga4ghStudy) SetStatus(v string) {
	o.Status = &v
}

// GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.
func (o *Ga4ghStudy) GetRecordingDate() string {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret
	}
	return *o.RecordingDate
}

// GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghStudy) GetRecordingDateOk() (string, bool) {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret, false
	}
	return *o.RecordingDate, true
}

// HasRecordingDate returns a boolean if a field has been set.
func (o *Ga4ghStudy) HasRecordingDate() bool {
	if o != nil && o.RecordingDate != nil {
		return true
	}

	return false
}

// SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.
func (o *Ga4ghStudy) SetRecordingDate(v string) {
	o.RecordingDate = &v
}


func (o Ga4ghStudy) MarshalJSON() ([]byte, error) {
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
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.RecordingDate != nil {
		toSerialize["recordingDate"] = o.RecordingDate
	}
	return json.Marshal(toSerialize)
}


