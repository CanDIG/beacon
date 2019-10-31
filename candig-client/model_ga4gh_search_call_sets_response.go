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

// Ga4ghSearchCallSetsResponse This is the response from `POST /callsets/search` expressed as JSON.
type Ga4ghSearchCallSetsResponse struct {
	// The list of matching call sets.
	CallSets *[]Ga4ghCallSet `json:"call_sets,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// GetCallSets returns the CallSets field if non-nil, zero value otherwise.
func (o *Ga4ghSearchCallSetsResponse) GetCallSets() []Ga4ghCallSet {
	if o == nil || o.CallSets == nil {
		var ret []Ga4ghCallSet
		return ret
	}
	return *o.CallSets
}

// GetCallSetsOk returns a tuple with the CallSets field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchCallSetsResponse) GetCallSetsOk() ([]Ga4ghCallSet, bool) {
	if o == nil || o.CallSets == nil {
		var ret []Ga4ghCallSet
		return ret, false
	}
	return *o.CallSets, true
}

// HasCallSets returns a boolean if a field has been set.
func (o *Ga4ghSearchCallSetsResponse) HasCallSets() bool {
	if o != nil && o.CallSets != nil {
		return true
	}

	return false
}

// SetCallSets gets a reference to the given []Ga4ghCallSet and assigns it to the CallSets field.
func (o *Ga4ghSearchCallSetsResponse) SetCallSets(v []Ga4ghCallSet) {
	o.CallSets = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchCallSetsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchCallSetsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchCallSetsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchCallSetsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Ga4ghSearchCallSetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallSets != nil {
		toSerialize["call_sets"] = o.CallSets
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}
