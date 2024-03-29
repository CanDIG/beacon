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

// ****************  /references  ******************* This request maps to the body of `POST /references/search` as JSON.
type Ga4ghSearchReferencesRequest struct {
	// The `ReferenceSet` to search.
	ReferenceSetId *string `json:"reference_set_id,omitempty"`

	// If specified, return the references for which the `md5checksum` matches this string (case-sensitive, exact match). See `ReferenceSet::md5checksum` for details.
	Md5checksum *string `json:"md5checksum,omitempty"`

	// If specified, return the references for which the `accession` matches this string (case-sensitive, exact match).
	Accession *string `json:"accession,omitempty"`

	// Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used.
	PageSize *int32 `json:"page_size,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `next_page_token` from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetReferenceSetId returns the ReferenceSetId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferencesRequest) GetReferenceSetId() string {
	if o == nil || o.ReferenceSetId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceSetId
}

// GetReferenceSetIdOk returns a tuple with the ReferenceSetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferencesRequest) GetReferenceSetIdOk() (string, bool) {
	if o == nil || o.ReferenceSetId == nil {
		var ret string
		return ret, false
	}
	return *o.ReferenceSetId, true
}

// HasReferenceSetId returns a boolean if a field has been set.
func (o *Ga4ghSearchReferencesRequest) HasReferenceSetId() bool {
	if o != nil && o.ReferenceSetId != nil {
		return true
	}

	return false
}

// SetReferenceSetId gets a reference to the given string and assigns it to the ReferenceSetId field.
func (o *Ga4ghSearchReferencesRequest) SetReferenceSetId(v string) {
	o.ReferenceSetId = &v
}

// GetMd5checksum returns the Md5checksum field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferencesRequest) GetMd5checksum() string {
	if o == nil || o.Md5checksum == nil {
		var ret string
		return ret
	}
	return *o.Md5checksum
}

// GetMd5checksumOk returns a tuple with the Md5checksum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferencesRequest) GetMd5checksumOk() (string, bool) {
	if o == nil || o.Md5checksum == nil {
		var ret string
		return ret, false
	}
	return *o.Md5checksum, true
}

// HasMd5checksum returns a boolean if a field has been set.
func (o *Ga4ghSearchReferencesRequest) HasMd5checksum() bool {
	if o != nil && o.Md5checksum != nil {
		return true
	}

	return false
}

// SetMd5checksum gets a reference to the given string and assigns it to the Md5checksum field.
func (o *Ga4ghSearchReferencesRequest) SetMd5checksum(v string) {
	o.Md5checksum = &v
}

// GetAccession returns the Accession field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferencesRequest) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferencesRequest) GetAccessionOk() (string, bool) {
	if o == nil || o.Accession == nil {
		var ret string
		return ret, false
	}
	return *o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *Ga4ghSearchReferencesRequest) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *Ga4ghSearchReferencesRequest) SetAccession(v string) {
	o.Accession = &v
}

// GetPageSize returns the PageSize field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferencesRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferencesRequest) GetPageSizeOk() (int32, bool) {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret, false
	}
	return *o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Ga4ghSearchReferencesRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Ga4ghSearchReferencesRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchReferencesRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchReferencesRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchReferencesRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghSearchReferencesRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghSearchReferencesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceSetId != nil {
		toSerialize["reference_set_id"] = o.ReferenceSetId
	}
	if o.Md5checksum != nil {
		toSerialize["md5checksum"] = o.Md5checksum
	}
	if o.Accession != nil {
		toSerialize["accession"] = o.Accession
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}


