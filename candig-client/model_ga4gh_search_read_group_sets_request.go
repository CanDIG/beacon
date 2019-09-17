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

// This request maps to the body of `POST /readgroupsets/search` as JSON.  TODO: Factor this out to a common API patterns section. - If searching by a resource ID, and that resource is not found, the method   will return a `404` HTTP status code (`NOT_FOUND`). - If searching by other attributes, e.g. `name`, and no matches are found, the   method will return a `200` HTTP status code (`OK`) with an empty result list.
type Ga4ghSearchReadGroupSetsRequest struct {
	// The dataset to search.
	DatasetId *string `json:"dataset_id,omitempty"`

	// Only return read group sets with this name (case-sensitive, exact match).
	Name *string `json:"name,omitempty"`

	// Specifying the id of a Biosample record will return only readgroups  with the given biosampleId.
	BiosampleId *string `json:"biosample_id,omitempty"`

	// Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used.
	PageSize *int32 `json:"page_size,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `next_page_token` from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsRequest) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghSearchReadGroupSetsRequest) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghSearchReadGroupSetsRequest) SetName(v string) {
	o.Name = &v
}

// GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsRequest) GetBiosampleId() string {
	if o == nil || o.BiosampleId == nil {
		var ret string
		return ret
	}
	return *o.BiosampleId
}

// GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) GetBiosampleIdOk() (string, bool) {
	if o == nil || o.BiosampleId == nil {
		var ret string
		return ret, false
	}
	return *o.BiosampleId, true
}

// HasBiosampleId returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) HasBiosampleId() bool {
	if o != nil && o.BiosampleId != nil {
		return true
	}

	return false
}

// SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.
func (o *Ga4ghSearchReadGroupSetsRequest) SetBiosampleId(v string) {
	o.BiosampleId = &v
}

// GetPageSize returns the PageSize field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) GetPageSizeOk() (int32, bool) {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret, false
	}
	return *o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Ga4ghSearchReadGroupSetsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReadGroupSetsRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchReadGroupSetsRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghSearchReadGroupSetsRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghSearchReadGroupSetsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BiosampleId != nil {
		toSerialize["biosample_id"] = o.BiosampleId
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

