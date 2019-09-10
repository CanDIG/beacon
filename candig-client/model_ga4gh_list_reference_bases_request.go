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

// This request retrieves a region of a reference genome when sent to  the `/listreferencebases` endpoint.
type Ga4ghListReferenceBasesRequest struct {
	// The ID of the `Reference` to be retrieved.
	ReferenceId *string `json:"reference_id,omitempty"`

	// The start position (0-based) of this query. Defaults to 0. Genomic positions are non-negative integers less than reference length. Requests spanning the join of circular genomes are represented as two requests one on each side of the join (position 0).
	Start *string `json:"start,omitempty"`

	// The end position (0-based, exclusive) of this query. Defaults to the length of this `Reference`.
	End *string `json:"end,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `next_page_token` from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.
func (o *Ga4ghListReferenceBasesRequest) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghListReferenceBasesRequest) GetReferenceIdOk() (string, bool) {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret, false
	}
	return *o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *Ga4ghListReferenceBasesRequest) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *Ga4ghListReferenceBasesRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetStart returns the Start field if non-nil, zero value otherwise.
func (o *Ga4ghListReferenceBasesRequest) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghListReferenceBasesRequest) GetStartOk() (string, bool) {
	if o == nil || o.Start == nil {
		var ret string
		return ret, false
	}
	return *o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Ga4ghListReferenceBasesRequest) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *Ga4ghListReferenceBasesRequest) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field if non-nil, zero value otherwise.
func (o *Ga4ghListReferenceBasesRequest) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghListReferenceBasesRequest) GetEndOk() (string, bool) {
	if o == nil || o.End == nil {
		var ret string
		return ret, false
	}
	return *o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Ga4ghListReferenceBasesRequest) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *Ga4ghListReferenceBasesRequest) SetEnd(v string) {
	o.End = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghListReferenceBasesRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghListReferenceBasesRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghListReferenceBasesRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghListReferenceBasesRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghListReferenceBasesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceId != nil {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}


