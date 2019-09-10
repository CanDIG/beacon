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

type Ga4ghLabtest struct {
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

	CollectionDate *string `json:"collectionDate,omitempty"`

	EndDate *string `json:"endDate,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	TestResults *string `json:"testResults,omitempty"`

	TimePoint *string `json:"timePoint,omitempty"`

	RecordingDate *string `json:"recordingDate,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghLabtest) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghLabtest) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghLabtest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghLabtest) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghLabtest) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghLabtest) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghLabtest) SetPatientId(v string) {
	o.PatientId = &v
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetStartDateOk() (string, bool) {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Ga4ghLabtest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetCollectionDate() string {
	if o == nil || o.CollectionDate == nil {
		var ret string
		return ret
	}
	return *o.CollectionDate
}

// GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetCollectionDateOk() (string, bool) {
	if o == nil || o.CollectionDate == nil {
		var ret string
		return ret, false
	}
	return *o.CollectionDate, true
}

// HasCollectionDate returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasCollectionDate() bool {
	if o != nil && o.CollectionDate != nil {
		return true
	}

	return false
}

// SetCollectionDate gets a reference to the given string and assigns it to the CollectionDate field.
func (o *Ga4ghLabtest) SetCollectionDate(v string) {
	o.CollectionDate = &v
}

// GetEndDate returns the EndDate field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetEndDateOk() (string, bool) {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret, false
	}
	return *o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Ga4ghLabtest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEventType returns the EventType field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetEventTypeOk() (string, bool) {
	if o == nil || o.EventType == nil {
		var ret string
		return ret, false
	}
	return *o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *Ga4ghLabtest) SetEventType(v string) {
	o.EventType = &v
}

// GetTestResults returns the TestResults field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetTestResults() string {
	if o == nil || o.TestResults == nil {
		var ret string
		return ret
	}
	return *o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetTestResultsOk() (string, bool) {
	if o == nil || o.TestResults == nil {
		var ret string
		return ret, false
	}
	return *o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasTestResults() bool {
	if o != nil && o.TestResults != nil {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given string and assigns it to the TestResults field.
func (o *Ga4ghLabtest) SetTestResults(v string) {
	o.TestResults = &v
}

// GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetTimePoint() string {
	if o == nil || o.TimePoint == nil {
		var ret string
		return ret
	}
	return *o.TimePoint
}

// GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetTimePointOk() (string, bool) {
	if o == nil || o.TimePoint == nil {
		var ret string
		return ret, false
	}
	return *o.TimePoint, true
}

// HasTimePoint returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasTimePoint() bool {
	if o != nil && o.TimePoint != nil {
		return true
	}

	return false
}

// SetTimePoint gets a reference to the given string and assigns it to the TimePoint field.
func (o *Ga4ghLabtest) SetTimePoint(v string) {
	o.TimePoint = &v
}

// GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.
func (o *Ga4ghLabtest) GetRecordingDate() string {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret
	}
	return *o.RecordingDate
}

// GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLabtest) GetRecordingDateOk() (string, bool) {
	if o == nil || o.RecordingDate == nil {
		var ret string
		return ret, false
	}
	return *o.RecordingDate, true
}

// HasRecordingDate returns a boolean if a field has been set.
func (o *Ga4ghLabtest) HasRecordingDate() bool {
	if o != nil && o.RecordingDate != nil {
		return true
	}

	return false
}

// SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.
func (o *Ga4ghLabtest) SetRecordingDate(v string) {
	o.RecordingDate = &v
}


func (o Ga4ghLabtest) MarshalJSON() ([]byte, error) {
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
	if o.CollectionDate != nil {
		toSerialize["collectionDate"] = o.CollectionDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.TestResults != nil {
		toSerialize["testResults"] = o.TestResults
	}
	if o.TimePoint != nil {
		toSerialize["timePoint"] = o.TimePoint
	}
	if o.RecordingDate != nil {
		toSerialize["recordingDate"] = o.RecordingDate
	}
	return json.Marshal(toSerialize)
}


