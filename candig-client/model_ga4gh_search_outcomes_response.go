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

// This is the response from `POST /outcomes/search` expressed as JSON.
type Ga4ghSearchOutcomesResponse struct {
	// The list of outcomes.
	Outcomes *[]Ga4ghOutcome `json:"outcomes,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.
func (o *Ga4ghSearchOutcomesResponse) GetOutcomes() []Ga4ghOutcome {
	if o == nil || o.Outcomes == nil {
		var ret []Ga4ghOutcome
		return ret
	}
	return *o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchOutcomesResponse) GetOutcomesOk() ([]Ga4ghOutcome, bool) {
	if o == nil || o.Outcomes == nil {
		var ret []Ga4ghOutcome
		return ret, false
	}
	return *o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *Ga4ghSearchOutcomesResponse) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given []Ga4ghOutcome and assigns it to the Outcomes field.
func (o *Ga4ghSearchOutcomesResponse) SetOutcomes(v []Ga4ghOutcome) {
	o.Outcomes = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchOutcomesResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchOutcomesResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchOutcomesResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchOutcomesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchOutcomesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Outcomes != nil {
		toSerialize["outcomes"] = o.Outcomes
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


