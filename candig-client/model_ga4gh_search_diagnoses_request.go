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

//  Diagnosis  This request maps to the body of `POST /diagnoses/search` as JSON.
type Ga4ghSearchDiagnosesRequest struct {
	// Optionally specify the dataset to search within.
	DatasetId *string `json:"dataset_id,omitempty"`

	// Returns Diagnoses with the given name found by case-sensitive string matching.
	Name *string `json:"name,omitempty"`

	Filters *[]Ga4ghFilter `json:"filters,omitempty"`

	// Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used.
	PageSize *int32 `json:"page_size,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `nextPageToken` from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDiagnosesRequest) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDiagnosesRequest) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghSearchDiagnosesRequest) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghSearchDiagnosesRequest) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDiagnosesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDiagnosesRequest) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghSearchDiagnosesRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghSearchDiagnosesRequest) SetName(v string) {
	o.Name = &v
}

// GetFilters returns the Filters field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDiagnosesRequest) GetFilters() []Ga4ghFilter {
	if o == nil || o.Filters == nil {
		var ret []Ga4ghFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDiagnosesRequest) GetFiltersOk() ([]Ga4ghFilter, bool) {
	if o == nil || o.Filters == nil {
		var ret []Ga4ghFilter
		return ret, false
	}
	return *o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *Ga4ghSearchDiagnosesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []Ga4ghFilter and assigns it to the Filters field.
func (o *Ga4ghSearchDiagnosesRequest) SetFilters(v []Ga4ghFilter) {
	o.Filters = &v
}

// GetPageSize returns the PageSize field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDiagnosesRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDiagnosesRequest) GetPageSizeOk() (int32, bool) {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret, false
	}
	return *o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Ga4ghSearchDiagnosesRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Ga4ghSearchDiagnosesRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDiagnosesRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDiagnosesRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchDiagnosesRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghSearchDiagnosesRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghSearchDiagnosesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

