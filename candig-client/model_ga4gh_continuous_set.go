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

// A set of Continous messages. Continuous values can be sent as a single Continuous message containing all values or a series of Continuous messages to either limit the size of the values array or to skip NaN values.
type Ga4ghContinuousSet struct {
	// The ID of this annotation set.
	Id *string `json:"id,omitempty"`

	// The ID of the dataset this annotation set belongs to.
	DatasetId *string `json:"dataset_id,omitempty"`

	// The ID of the reference set which defines the coordinate-space for this set of annotations.
	ReferenceSetId *string `json:"reference_set_id,omitempty"`

	// The display name for this annotation set.
	Name *string `json:"name,omitempty"`

	// The source URI describing the file from which this annotation set was generated, if any.
	SourceUri *string `json:"source_uri,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghContinuousSet) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghContinuousSet) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghContinuousSet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghContinuousSet) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghContinuousSet) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghContinuousSet) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghContinuousSet) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghContinuousSet) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetReferenceSetId returns the ReferenceSetId field if non-nil, zero value otherwise.
func (o *Ga4ghContinuousSet) GetReferenceSetId() string {
	if o == nil || o.ReferenceSetId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceSetId
}

// GetReferenceSetIdOk returns a tuple with the ReferenceSetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghContinuousSet) GetReferenceSetIdOk() (string, bool) {
	if o == nil || o.ReferenceSetId == nil {
		var ret string
		return ret, false
	}
	return *o.ReferenceSetId, true
}

// HasReferenceSetId returns a boolean if a field has been set.
func (o *Ga4ghContinuousSet) HasReferenceSetId() bool {
	if o != nil && o.ReferenceSetId != nil {
		return true
	}

	return false
}

// SetReferenceSetId gets a reference to the given string and assigns it to the ReferenceSetId field.
func (o *Ga4ghContinuousSet) SetReferenceSetId(v string) {
	o.ReferenceSetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghContinuousSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghContinuousSet) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghContinuousSet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghContinuousSet) SetName(v string) {
	o.Name = &v
}

// GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.
func (o *Ga4ghContinuousSet) GetSourceUri() string {
	if o == nil || o.SourceUri == nil {
		var ret string
		return ret
	}
	return *o.SourceUri
}

// GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghContinuousSet) GetSourceUriOk() (string, bool) {
	if o == nil || o.SourceUri == nil {
		var ret string
		return ret, false
	}
	return *o.SourceUri, true
}

// HasSourceUri returns a boolean if a field has been set.
func (o *Ga4ghContinuousSet) HasSourceUri() bool {
	if o != nil && o.SourceUri != nil {
		return true
	}

	return false
}

// SetSourceUri gets a reference to the given string and assigns it to the SourceUri field.
func (o *Ga4ghContinuousSet) SetSourceUri(v string) {
	o.SourceUri = &v
}


func (o Ga4ghContinuousSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if o.ReferenceSetId != nil {
		toSerialize["reference_set_id"] = o.ReferenceSetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SourceUri != nil {
		toSerialize["source_uri"] = o.SourceUri
	}
	return json.Marshal(toSerialize)
}


