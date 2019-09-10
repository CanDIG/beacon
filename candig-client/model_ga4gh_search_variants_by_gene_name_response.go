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

// This is the response from `POST /variantsbygenename` expressed as JSON.
type Ga4ghSearchVariantsByGeneNameResponse struct {
	// The list of matching variants. If the `callSetId` field on the returned calls is not present, the ordering of the call sets from a `SearchCallSetsRequest` over the parent `VariantSet` is guaranteed to match the ordering of the calls on each `Variant`. The number of results will also be the same.
	Variants *[]Ga4ghVariant `json:"variants,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetVariants returns the Variants field if non-nil, zero value otherwise.
func (o *Ga4ghSearchVariantsByGeneNameResponse) GetVariants() []Ga4ghVariant {
	if o == nil || o.Variants == nil {
		var ret []Ga4ghVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchVariantsByGeneNameResponse) GetVariantsOk() ([]Ga4ghVariant, bool) {
	if o == nil || o.Variants == nil {
		var ret []Ga4ghVariant
		return ret, false
	}
	return *o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *Ga4ghSearchVariantsByGeneNameResponse) HasVariants() bool {
	if o != nil && o.Variants != nil {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []Ga4ghVariant and assigns it to the Variants field.
func (o *Ga4ghSearchVariantsByGeneNameResponse) SetVariants(v []Ga4ghVariant) {
	o.Variants = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchVariantsByGeneNameResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchVariantsByGeneNameResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchVariantsByGeneNameResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchVariantsByGeneNameResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchVariantsByGeneNameResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


