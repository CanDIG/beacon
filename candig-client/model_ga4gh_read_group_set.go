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

// A ReadGroupSet is a logical collection of ReadGroups. Typically one ReadGroupSet represents all the reads from one experimental sample.
type Ga4ghReadGroupSet struct {
	// The read group set ID.
	Id *string `json:"id,omitempty"`

	// The ID of the dataset this read group set belongs to.
	DatasetId *string `json:"dataset_id,omitempty"`

	// The read group set name.
	Name *string `json:"name,omitempty"`

	Stats *Ga4ghReadStats `json:"stats,omitempty"`

	// The read groups in this set.
	ReadGroups *[]Ga4ghReadGroup `json:"read_groups,omitempty"`

	PatientId *string `json:"patient_id,omitempty"`

	SampleId *string `json:"sample_id,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghReadGroupSet) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghReadGroupSet) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghReadGroupSet) SetName(v string) {
	o.Name = &v
}

// GetStats returns the Stats field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetStats() Ga4ghReadStats {
	if o == nil || o.Stats == nil {
		var ret Ga4ghReadStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetStatsOk() (Ga4ghReadStats, bool) {
	if o == nil || o.Stats == nil {
		var ret Ga4ghReadStats
		return ret, false
	}
	return *o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given Ga4ghReadStats and assigns it to the Stats field.
func (o *Ga4ghReadGroupSet) SetStats(v Ga4ghReadStats) {
	o.Stats = &v
}

// GetReadGroups returns the ReadGroups field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetReadGroups() []Ga4ghReadGroup {
	if o == nil || o.ReadGroups == nil {
		var ret []Ga4ghReadGroup
		return ret
	}
	return *o.ReadGroups
}

// GetReadGroupsOk returns a tuple with the ReadGroups field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetReadGroupsOk() ([]Ga4ghReadGroup, bool) {
	if o == nil || o.ReadGroups == nil {
		var ret []Ga4ghReadGroup
		return ret, false
	}
	return *o.ReadGroups, true
}

// HasReadGroups returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasReadGroups() bool {
	if o != nil && o.ReadGroups != nil {
		return true
	}

	return false
}

// SetReadGroups gets a reference to the given []Ga4ghReadGroup and assigns it to the ReadGroups field.
func (o *Ga4ghReadGroupSet) SetReadGroups(v []Ga4ghReadGroup) {
	o.ReadGroups = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghReadGroupSet) SetPatientId(v string) {
	o.PatientId = &v
}

// GetSampleId returns the SampleId field if non-nil, zero value otherwise.
func (o *Ga4ghReadGroupSet) GetSampleId() string {
	if o == nil || o.SampleId == nil {
		var ret string
		return ret
	}
	return *o.SampleId
}

// GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadGroupSet) GetSampleIdOk() (string, bool) {
	if o == nil || o.SampleId == nil {
		var ret string
		return ret, false
	}
	return *o.SampleId, true
}

// HasSampleId returns a boolean if a field has been set.
func (o *Ga4ghReadGroupSet) HasSampleId() bool {
	if o != nil && o.SampleId != nil {
		return true
	}

	return false
}

// SetSampleId gets a reference to the given string and assigns it to the SampleId field.
func (o *Ga4ghReadGroupSet) SetSampleId(v string) {
	o.SampleId = &v
}


func (o Ga4ghReadGroupSet) MarshalJSON() ([]byte, error) {
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
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	if o.ReadGroups != nil {
		toSerialize["read_groups"] = o.ReadGroups
	}
	if o.PatientId != nil {
		toSerialize["patient_id"] = o.PatientId
	}
	if o.SampleId != nil {
		toSerialize["sample_id"] = o.SampleId
	}
	return json.Marshal(toSerialize)
}


