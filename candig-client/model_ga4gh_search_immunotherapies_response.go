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

// This is the response from `POST /immunotherapies/search` expressed as JSON.
type Ga4ghSearchImmunotherapiesResponse struct {
	// The list of immunotherapies.
	Immunotherapies *[]Ga4ghImmunotherapy `json:"immunotherapies,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetImmunotherapies returns the Immunotherapies field if non-nil, zero value otherwise.
func (o *Ga4ghSearchImmunotherapiesResponse) GetImmunotherapies() []Ga4ghImmunotherapy {
	if o == nil || o.Immunotherapies == nil {
		var ret []Ga4ghImmunotherapy
		return ret
	}
	return *o.Immunotherapies
}

// GetImmunotherapiesOk returns a tuple with the Immunotherapies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchImmunotherapiesResponse) GetImmunotherapiesOk() ([]Ga4ghImmunotherapy, bool) {
	if o == nil || o.Immunotherapies == nil {
		var ret []Ga4ghImmunotherapy
		return ret, false
	}
	return *o.Immunotherapies, true
}

// HasImmunotherapies returns a boolean if a field has been set.
func (o *Ga4ghSearchImmunotherapiesResponse) HasImmunotherapies() bool {
	if o != nil && o.Immunotherapies != nil {
		return true
	}

	return false
}

// SetImmunotherapies gets a reference to the given []Ga4ghImmunotherapy and assigns it to the Immunotherapies field.
func (o *Ga4ghSearchImmunotherapiesResponse) SetImmunotherapies(v []Ga4ghImmunotherapy) {
	o.Immunotherapies = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchImmunotherapiesResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchImmunotherapiesResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchImmunotherapiesResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchImmunotherapiesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchImmunotherapiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Immunotherapies != nil {
		toSerialize["immunotherapies"] = o.Immunotherapies
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

