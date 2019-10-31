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

// Ga4ghSearchExtractionsResponse This is the response from `POST /extractions/search` expressed as JSON.
type Ga4ghSearchExtractionsResponse struct {
	// The list of extractions.
	Extractions *[]Ga4ghExtraction `json:"extractions,omitempty"`

	// The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// GetExtractions returns the Extractions field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExtractionsResponse) GetExtractions() []Ga4ghExtraction {
	if o == nil || o.Extractions == nil {
		var ret []Ga4ghExtraction
		return ret
	}
	return *o.Extractions
}

// GetExtractionsOk returns a tuple with the Extractions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExtractionsResponse) GetExtractionsOk() ([]Ga4ghExtraction, bool) {
	if o == nil || o.Extractions == nil {
		var ret []Ga4ghExtraction
		return ret, false
	}
	return *o.Extractions, true
}

// HasExtractions returns a boolean if a field has been set.
func (o *Ga4ghSearchExtractionsResponse) HasExtractions() bool {
	if o != nil && o.Extractions != nil {
		return true
	}

	return false
}

// SetExtractions gets a reference to the given []Ga4ghExtraction and assigns it to the Extractions field.
func (o *Ga4ghSearchExtractionsResponse) SetExtractions(v []Ga4ghExtraction) {
	o.Extractions = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExtractionsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExtractionsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchExtractionsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchExtractionsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Ga4ghSearchExtractionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extractions != nil {
		toSerialize["extractions"] = o.Extractions
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}
