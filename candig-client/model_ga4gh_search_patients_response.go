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

// This is the response from `POST /patients/search` expressed as JSON.
type Ga4ghSearchPatientsResponse struct {
	// The list of patients.
	Patients *[]Ga4ghPatient `json:"patients,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetPatients returns the Patients field if non-nil, zero value otherwise.
func (o *Ga4ghSearchPatientsResponse) GetPatients() []Ga4ghPatient {
	if o == nil || o.Patients == nil {
		var ret []Ga4ghPatient
		return ret
	}
	return *o.Patients
}

// GetPatientsOk returns a tuple with the Patients field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchPatientsResponse) GetPatientsOk() ([]Ga4ghPatient, bool) {
	if o == nil || o.Patients == nil {
		var ret []Ga4ghPatient
		return ret, false
	}
	return *o.Patients, true
}

// HasPatients returns a boolean if a field has been set.
func (o *Ga4ghSearchPatientsResponse) HasPatients() bool {
	if o != nil && o.Patients != nil {
		return true
	}

	return false
}

// SetPatients gets a reference to the given []Ga4ghPatient and assigns it to the Patients field.
func (o *Ga4ghSearchPatientsResponse) SetPatients(v []Ga4ghPatient) {
	o.Patients = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchPatientsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchPatientsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchPatientsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchPatientsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchPatientsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Patients != nil {
		toSerialize["patients"] = o.Patients
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


