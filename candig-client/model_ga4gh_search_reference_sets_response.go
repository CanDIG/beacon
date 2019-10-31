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

// Ga4ghSearchReferenceSetsResponse This is the response from `POST /referencesets/search` expressed as JSON.
type Ga4ghSearchReferenceSetsResponse struct {
	// The list of matching reference sets.
	ReferenceSets *[]Ga4ghReferenceSet `json:"reference_sets,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// GetReferenceSets returns the ReferenceSets field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferenceSetsResponse) GetReferenceSets() []Ga4ghReferenceSet {
	if o == nil || o.ReferenceSets == nil {
		var ret []Ga4ghReferenceSet
		return ret
	}
	return *o.ReferenceSets
}

// GetReferenceSetsOk returns a tuple with the ReferenceSets field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferenceSetsResponse) GetReferenceSetsOk() ([]Ga4ghReferenceSet, bool) {
	if o == nil || o.ReferenceSets == nil {
		var ret []Ga4ghReferenceSet
		return ret, false
	}
	return *o.ReferenceSets, true
}

// HasReferenceSets returns a boolean if a field has been set.
func (o *Ga4ghSearchReferenceSetsResponse) HasReferenceSets() bool {
	if o != nil && o.ReferenceSets != nil {
		return true
	}

	return false
}

// SetReferenceSets gets a reference to the given []Ga4ghReferenceSet and assigns it to the ReferenceSets field.
func (o *Ga4ghSearchReferenceSetsResponse) SetReferenceSets(v []Ga4ghReferenceSet) {
	o.ReferenceSets = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferenceSetsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferenceSetsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchReferenceSetsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchReferenceSetsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Ga4ghSearchReferenceSetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceSets != nil {
		toSerialize["reference_sets"] = o.ReferenceSets
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}
