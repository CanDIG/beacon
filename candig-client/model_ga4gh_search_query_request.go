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

type Ga4ghSearchQueryRequest struct {
	DatasetId *string `json:"dataset_id,omitempty"`

	Logic *Ga4ghLogic `json:"logic,omitempty"`

	Components *[]Ga4ghComponent `json:"components,omitempty"`

	Results *[]Ga4ghResult `json:"results,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	// Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used.
	PageSize *int32 `json:"page_size,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `nextPageToken` from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghSearchQueryRequest) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetLogic returns the Logic field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetLogic() Ga4ghLogic {
	if o == nil || o.Logic == nil {
		var ret Ga4ghLogic
		return ret
	}
	return *o.Logic
}

// GetLogicOk returns a tuple with the Logic field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetLogicOk() (Ga4ghLogic, bool) {
	if o == nil || o.Logic == nil {
		var ret Ga4ghLogic
		return ret, false
	}
	return *o.Logic, true
}

// HasLogic returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasLogic() bool {
	if o != nil && o.Logic != nil {
		return true
	}

	return false
}

// SetLogic gets a reference to the given Ga4ghLogic and assigns it to the Logic field.
func (o *Ga4ghSearchQueryRequest) SetLogic(v Ga4ghLogic) {
	o.Logic = &v
}

// GetComponents returns the Components field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetComponents() []Ga4ghComponent {
	if o == nil || o.Components == nil {
		var ret []Ga4ghComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetComponentsOk() ([]Ga4ghComponent, bool) {
	if o == nil || o.Components == nil {
		var ret []Ga4ghComponent
		return ret, false
	}
	return *o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []Ga4ghComponent and assigns it to the Components field.
func (o *Ga4ghSearchQueryRequest) SetComponents(v []Ga4ghComponent) {
	o.Components = &v
}

// GetResults returns the Results field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetResults() []Ga4ghResult {
	if o == nil || o.Results == nil {
		var ret []Ga4ghResult
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetResultsOk() ([]Ga4ghResult, bool) {
	if o == nil || o.Results == nil {
		var ret []Ga4ghResult
		return ret, false
	}
	return *o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Ga4ghResult and assigns it to the Results field.
func (o *Ga4ghSearchQueryRequest) SetResults(v []Ga4ghResult) {
	o.Results = &v
}

// GetLimit returns the Limit field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetLimitOk() (int32, bool) {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret, false
	}
	return *o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *Ga4ghSearchQueryRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetPageSize returns the PageSize field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetPageSizeOk() (int32, bool) {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret, false
	}
	return *o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Ga4ghSearchQueryRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghSearchQueryRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghSearchQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if o.Logic != nil {
		toSerialize["logic"] = o.Logic
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}


