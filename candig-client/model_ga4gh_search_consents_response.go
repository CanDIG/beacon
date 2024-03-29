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

// This is the response from `POST /consents/search` expressed as JSON.
type Ga4ghSearchConsentsResponse struct {
	// The list of consents.
	Consents *[]Ga4ghConsent `json:"consents,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetConsents returns the Consents field if non-nil, zero value otherwise.
func (o *Ga4ghSearchConsentsResponse) GetConsents() []Ga4ghConsent {
	if o == nil || o.Consents == nil {
		var ret []Ga4ghConsent
		return ret
	}
	return *o.Consents
}

// GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchConsentsResponse) GetConsentsOk() ([]Ga4ghConsent, bool) {
	if o == nil || o.Consents == nil {
		var ret []Ga4ghConsent
		return ret, false
	}
	return *o.Consents, true
}

// HasConsents returns a boolean if a field has been set.
func (o *Ga4ghSearchConsentsResponse) HasConsents() bool {
	if o != nil && o.Consents != nil {
		return true
	}

	return false
}

// SetConsents gets a reference to the given []Ga4ghConsent and assigns it to the Consents field.
func (o *Ga4ghSearchConsentsResponse) SetConsents(v []Ga4ghConsent) {
	o.Consents = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchConsentsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchConsentsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchConsentsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchConsentsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchConsentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Consents != nil {
		toSerialize["consents"] = o.Consents
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


