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

// ****************  /expressionlevels/search  ******************* This request maps to the body of 'POST /expressionlevels/search' as JSON.
type Ga4ghSearchExpressionLevelsRequest struct {
	// The rnaQuantification to restrict search to.
	RnaQuantificationId *string `json:"rna_quantification_id,omitempty"`

	// Only return expressions with any of the names (strict string matching).
	Names *[]string `json:"names,omitempty"`

	Threshold *float32 `json:"threshold,omitempty"`

	// Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used.
	PageSize *int32 `json:"page_size,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of 'nextPageToken' from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetRnaQuantificationId returns the RnaQuantificationId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsRequest) GetRnaQuantificationId() string {
	if o == nil || o.RnaQuantificationId == nil {
		var ret string
		return ret
	}
	return *o.RnaQuantificationId
}

// GetRnaQuantificationIdOk returns a tuple with the RnaQuantificationId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) GetRnaQuantificationIdOk() (string, bool) {
	if o == nil || o.RnaQuantificationId == nil {
		var ret string
		return ret, false
	}
	return *o.RnaQuantificationId, true
}

// HasRnaQuantificationId returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) HasRnaQuantificationId() bool {
	if o != nil && o.RnaQuantificationId != nil {
		return true
	}

	return false
}

// SetRnaQuantificationId gets a reference to the given string and assigns it to the RnaQuantificationId field.
func (o *Ga4ghSearchExpressionLevelsRequest) SetRnaQuantificationId(v string) {
	o.RnaQuantificationId = &v
}

// GetNames returns the Names field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsRequest) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) GetNamesOk() ([]string, bool) {
	if o == nil || o.Names == nil {
		var ret []string
		return ret, false
	}
	return *o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *Ga4ghSearchExpressionLevelsRequest) SetNames(v []string) {
	o.Names = &v
}

// GetThreshold returns the Threshold field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsRequest) GetThreshold() float32 {
	if o == nil || o.Threshold == nil {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) GetThresholdOk() (float32, bool) {
	if o == nil || o.Threshold == nil {
		var ret float32
		return ret, false
	}
	return *o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *Ga4ghSearchExpressionLevelsRequest) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetPageSize returns the PageSize field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) GetPageSizeOk() (int32, bool) {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret, false
	}
	return *o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Ga4ghSearchExpressionLevelsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchExpressionLevelsRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchExpressionLevelsRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghSearchExpressionLevelsRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghSearchExpressionLevelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RnaQuantificationId != nil {
		toSerialize["rna_quantification_id"] = o.RnaQuantificationId
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

