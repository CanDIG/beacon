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

// This is the response from 'POST /expressionlevels/search' expressed as JSON.
type Ga4ghSearchExpressionLevelsResponse struct {
	// The list of matching quantifications.
	ExpressionLevels *[]Ga4ghExpressionLevel `json:"expression_levels,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of 'nextPageToken' from the previous response.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetExpressionLevels returns the ExpressionLevels field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsResponse) GetExpressionLevels() []Ga4ghExpressionLevel {
	if o == nil || o.ExpressionLevels == nil {
		var ret []Ga4ghExpressionLevel
		return ret
	}
	return *o.ExpressionLevels
}

// GetExpressionLevelsOk returns a tuple with the ExpressionLevels field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsResponse) GetExpressionLevelsOk() ([]Ga4ghExpressionLevel, bool) {
	if o == nil || o.ExpressionLevels == nil {
		var ret []Ga4ghExpressionLevel
		return ret, false
	}
	return *o.ExpressionLevels, true
}

// HasExpressionLevels returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsResponse) HasExpressionLevels() bool {
	if o != nil && o.ExpressionLevels != nil {
		return true
	}

	return false
}

// SetExpressionLevels gets a reference to the given []Ga4ghExpressionLevel and assigns it to the ExpressionLevels field.
func (o *Ga4ghSearchExpressionLevelsResponse) SetExpressionLevels(v []Ga4ghExpressionLevel) {
	o.ExpressionLevels = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchExpressionLevelsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchExpressionLevelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpressionLevels != nil {
		toSerialize["expression_levels"] = o.ExpressionLevels
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

