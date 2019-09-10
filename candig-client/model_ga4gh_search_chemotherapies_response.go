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

// This is the response from `POST /chemotherapies/search` expressed as JSON.
type Ga4ghSearchChemotherapiesResponse struct {
	// The list of chemotherapies.
	Chemotherapies *[]Ga4ghChemotherapy `json:"chemotherapies,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetChemotherapies returns the Chemotherapies field if non-nil, zero value otherwise.
func (o *Ga4ghSearchChemotherapiesResponse) GetChemotherapies() []Ga4ghChemotherapy {
	if o == nil || o.Chemotherapies == nil {
		var ret []Ga4ghChemotherapy
		return ret
	}
	return *o.Chemotherapies
}

// GetChemotherapiesOk returns a tuple with the Chemotherapies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchChemotherapiesResponse) GetChemotherapiesOk() ([]Ga4ghChemotherapy, bool) {
	if o == nil || o.Chemotherapies == nil {
		var ret []Ga4ghChemotherapy
		return ret, false
	}
	return *o.Chemotherapies, true
}

// HasChemotherapies returns a boolean if a field has been set.
func (o *Ga4ghSearchChemotherapiesResponse) HasChemotherapies() bool {
	if o != nil && o.Chemotherapies != nil {
		return true
	}

	return false
}

// SetChemotherapies gets a reference to the given []Ga4ghChemotherapy and assigns it to the Chemotherapies field.
func (o *Ga4ghSearchChemotherapiesResponse) SetChemotherapies(v []Ga4ghChemotherapy) {
	o.Chemotherapies = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchChemotherapiesResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchChemotherapiesResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchChemotherapiesResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchChemotherapiesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchChemotherapiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Chemotherapies != nil {
		toSerialize["chemotherapies"] = o.Chemotherapies
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


