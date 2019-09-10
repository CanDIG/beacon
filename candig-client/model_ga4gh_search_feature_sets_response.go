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

// This is the response from `POST /featuresets/search` expressed as JSON.
type Ga4ghSearchFeatureSetsResponse struct {
	// The list of matching feature sets.
	FeatureSets *[]Ga4ghFeatureSet `json:"feature_sets,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetFeatureSets returns the FeatureSets field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeatureSetsResponse) GetFeatureSets() []Ga4ghFeatureSet {
	if o == nil || o.FeatureSets == nil {
		var ret []Ga4ghFeatureSet
		return ret
	}
	return *o.FeatureSets
}

// GetFeatureSetsOk returns a tuple with the FeatureSets field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeatureSetsResponse) GetFeatureSetsOk() ([]Ga4ghFeatureSet, bool) {
	if o == nil || o.FeatureSets == nil {
		var ret []Ga4ghFeatureSet
		return ret, false
	}
	return *o.FeatureSets, true
}

// HasFeatureSets returns a boolean if a field has been set.
func (o *Ga4ghSearchFeatureSetsResponse) HasFeatureSets() bool {
	if o != nil && o.FeatureSets != nil {
		return true
	}

	return false
}

// SetFeatureSets gets a reference to the given []Ga4ghFeatureSet and assigns it to the FeatureSets field.
func (o *Ga4ghSearchFeatureSetsResponse) SetFeatureSets(v []Ga4ghFeatureSet) {
	o.FeatureSets = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeatureSetsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeatureSetsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchFeatureSetsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchFeatureSetsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchFeatureSetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureSets != nil {
		toSerialize["feature_sets"] = o.FeatureSets
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


