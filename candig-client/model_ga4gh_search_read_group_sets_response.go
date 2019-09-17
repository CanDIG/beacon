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

// This is the response from `POST /readgroupsets/search` expressed as JSON.
type Ga4ghSearchReadGroupSetsResponse struct {
	// The list of matching read group sets.
	ReadGroupSets *[]Ga4ghReadGroupSet `json:"read_group_sets,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetReadGroupSets returns the ReadGroupSets field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsResponse) GetReadGroupSets() []Ga4ghReadGroupSet {
	if o == nil || o.ReadGroupSets == nil {
		var ret []Ga4ghReadGroupSet
		return ret
	}
	return *o.ReadGroupSets
}

// GetReadGroupSetsOk returns a tuple with the ReadGroupSets field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsResponse) GetReadGroupSetsOk() ([]Ga4ghReadGroupSet, bool) {
	if o == nil || o.ReadGroupSets == nil {
		var ret []Ga4ghReadGroupSet
		return ret, false
	}
	return *o.ReadGroupSets, true
}

// HasReadGroupSets returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsResponse) HasReadGroupSets() bool {
	if o != nil && o.ReadGroupSets != nil {
		return true
	}

	return false
}

// SetReadGroupSets gets a reference to the given []Ga4ghReadGroupSet and assigns it to the ReadGroupSets field.
func (o *Ga4ghSearchReadGroupSetsResponse) SetReadGroupSets(v []Ga4ghReadGroupSet) {
	o.ReadGroupSets = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchReadGroupSetsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchReadGroupSetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReadGroupSets != nil {
		toSerialize["read_group_sets"] = o.ReadGroupSets
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

