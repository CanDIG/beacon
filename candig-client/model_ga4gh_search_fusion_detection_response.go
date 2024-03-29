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

// This is the response from `POST /fusiondetection/search` expressed as JSON.
type Ga4ghSearchFusionDetectionResponse struct {
	// The list of fusion/sv detections.
	Fusiondetection *[]Ga4ghFusionDetection `json:"fusiondetection,omitempty"`

	// The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetFusiondetection returns the Fusiondetection field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFusionDetectionResponse) GetFusiondetection() []Ga4ghFusionDetection {
	if o == nil || o.Fusiondetection == nil {
		var ret []Ga4ghFusionDetection
		return ret
	}
	return *o.Fusiondetection
}

// GetFusiondetectionOk returns a tuple with the Fusiondetection field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFusionDetectionResponse) GetFusiondetectionOk() ([]Ga4ghFusionDetection, bool) {
	if o == nil || o.Fusiondetection == nil {
		var ret []Ga4ghFusionDetection
		return ret, false
	}
	return *o.Fusiondetection, true
}

// HasFusiondetection returns a boolean if a field has been set.
func (o *Ga4ghSearchFusionDetectionResponse) HasFusiondetection() bool {
	if o != nil && o.Fusiondetection != nil {
		return true
	}

	return false
}

// SetFusiondetection gets a reference to the given []Ga4ghFusionDetection and assigns it to the Fusiondetection field.
func (o *Ga4ghSearchFusionDetectionResponse) SetFusiondetection(v []Ga4ghFusionDetection) {
	o.Fusiondetection = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFusionDetectionResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFusionDetectionResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchFusionDetectionResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchFusionDetectionResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchFusionDetectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fusiondetection != nil {
		toSerialize["fusiondetection"] = o.Fusiondetection
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


