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

// This request maps to the body of `POST /features/search` as JSON.
type Ga4ghSearchFeaturesRequest struct {
	// The annotation set to search within. Either `feature_set_id` or `parent_id` must be non-empty.
	FeatureSetId *string `json:"feature_set_id,omitempty"`

	// Only returns features with this name (case-sensitive, exact match).
	Name *string `json:"name,omitempty"`

	// Only return features with matching the provided gene symbol (case-sensitive, exact match). This field may be replaced with a more generic representation in a future version.
	GeneSymbol *string `json:"gene_symbol,omitempty"`

	// Restricts the search to direct children of the given parent `feature` ID. Either `feature_set_id` or `parent_id` must be non-empty.
	ParentId *string `json:"parent_id,omitempty"`

	// Only return features on the reference with this name (matched to literal reference name as imported from the GFF3).
	ReferenceName *string `json:"reference_name,omitempty"`

	// Required, if name or symbol not provided. The beginning of the window (0-based, inclusive) for which overlapping features should be returned.  Genomic positions are non-negative integers less than reference length.  Requests spanning the join of circular genomes are represented as two requests one on each side of the join (position 0).
	Start *string `json:"start,omitempty"`

	// Required, if name or symbol not provided. The end of the window (0-based, exclusive) for which overlapping features should be returned.
	End *string `json:"end,omitempty"`

	// TODO: To be replaced with a fully featured ontology search once the Metadata definitions are rounded out. If specified, this query matches only annotations whose `feature_type` matches one of the provided ontology terms.
	FeatureTypes *[]string `json:"feature_types,omitempty"`

	// Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used.
	PageSize *int32 `json:"page_size,omitempty"`

	// The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `next_page_token` from the previous response.
	PageToken *string `json:"page_token,omitempty"`

}

// GetFeatureSetId returns the FeatureSetId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetFeatureSetId() string {
	if o == nil || o.FeatureSetId == nil {
		var ret string
		return ret
	}
	return *o.FeatureSetId
}

// GetFeatureSetIdOk returns a tuple with the FeatureSetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetFeatureSetIdOk() (string, bool) {
	if o == nil || o.FeatureSetId == nil {
		var ret string
		return ret, false
	}
	return *o.FeatureSetId, true
}

// HasFeatureSetId returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasFeatureSetId() bool {
	if o != nil && o.FeatureSetId != nil {
		return true
	}

	return false
}

// SetFeatureSetId gets a reference to the given string and assigns it to the FeatureSetId field.
func (o *Ga4ghSearchFeaturesRequest) SetFeatureSetId(v string) {
	o.FeatureSetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghSearchFeaturesRequest) SetName(v string) {
	o.Name = &v
}

// GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetGeneSymbol() string {
	if o == nil || o.GeneSymbol == nil {
		var ret string
		return ret
	}
	return *o.GeneSymbol
}

// GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetGeneSymbolOk() (string, bool) {
	if o == nil || o.GeneSymbol == nil {
		var ret string
		return ret, false
	}
	return *o.GeneSymbol, true
}

// HasGeneSymbol returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasGeneSymbol() bool {
	if o != nil && o.GeneSymbol != nil {
		return true
	}

	return false
}

// SetGeneSymbol gets a reference to the given string and assigns it to the GeneSymbol field.
func (o *Ga4ghSearchFeaturesRequest) SetGeneSymbol(v string) {
	o.GeneSymbol = &v
}

// GetParentId returns the ParentId field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetParentIdOk() (string, bool) {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret, false
	}
	return *o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Ga4ghSearchFeaturesRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetReferenceName() string {
	if o == nil || o.ReferenceName == nil {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetReferenceNameOk() (string, bool) {
	if o == nil || o.ReferenceName == nil {
		var ret string
		return ret, false
	}
	return *o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasReferenceName() bool {
	if o != nil && o.ReferenceName != nil {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *Ga4ghSearchFeaturesRequest) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetStart returns the Start field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetStartOk() (string, bool) {
	if o == nil || o.Start == nil {
		var ret string
		return ret, false
	}
	return *o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *Ga4ghSearchFeaturesRequest) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetEndOk() (string, bool) {
	if o == nil || o.End == nil {
		var ret string
		return ret, false
	}
	return *o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *Ga4ghSearchFeaturesRequest) SetEnd(v string) {
	o.End = &v
}

// GetFeatureTypes returns the FeatureTypes field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetFeatureTypes() []string {
	if o == nil || o.FeatureTypes == nil {
		var ret []string
		return ret
	}
	return *o.FeatureTypes
}

// GetFeatureTypesOk returns a tuple with the FeatureTypes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetFeatureTypesOk() ([]string, bool) {
	if o == nil || o.FeatureTypes == nil {
		var ret []string
		return ret, false
	}
	return *o.FeatureTypes, true
}

// HasFeatureTypes returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasFeatureTypes() bool {
	if o != nil && o.FeatureTypes != nil {
		return true
	}

	return false
}

// SetFeatureTypes gets a reference to the given []string and assigns it to the FeatureTypes field.
func (o *Ga4ghSearchFeaturesRequest) SetFeatureTypes(v []string) {
	o.FeatureTypes = &v
}

// GetPageSize returns the PageSize field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetPageSizeOk() (int32, bool) {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret, false
	}
	return *o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Ga4ghSearchFeaturesRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchFeaturesRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchFeaturesRequest) GetPageTokenOk() (string, bool) {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret, false
	}
	return *o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchFeaturesRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *Ga4ghSearchFeaturesRequest) SetPageToken(v string) {
	o.PageToken = &v
}


func (o Ga4ghSearchFeaturesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureSetId != nil {
		toSerialize["feature_set_id"] = o.FeatureSetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.GeneSymbol != nil {
		toSerialize["gene_symbol"] = o.GeneSymbol
	}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.ReferenceName != nil {
		toSerialize["reference_name"] = o.ReferenceName
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.FeatureTypes != nil {
		toSerialize["feature_types"] = o.FeatureTypes
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}


