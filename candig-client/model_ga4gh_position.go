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

// A `Position` is an unoriented base in some `Reference`. A `Position` is represented by a `Reference` name, and a base number on that `Reference` (0-based).
type Ga4ghPosition struct {
	// The name of the `Reference` on which the `Position` is located.
	ReferenceName *string `json:"reference_name,omitempty"`

	// The 0-based offset from the start of the forward strand for that `Reference`. Genomic positions are non-negative integers less than `Reference` length.
	Position *string `json:"position,omitempty"`

	Strand *Ga4ghStrand `json:"strand,omitempty"`

}

// GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.
func (o *Ga4ghPosition) GetReferenceName() string {
	if o == nil || o.ReferenceName == nil {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghPosition) GetReferenceNameOk() (string, bool) {
	if o == nil || o.ReferenceName == nil {
		var ret string
		return ret, false
	}
	return *o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *Ga4ghPosition) HasReferenceName() bool {
	if o != nil && o.ReferenceName != nil {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *Ga4ghPosition) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetPosition returns the Position field if non-nil, zero value otherwise.
func (o *Ga4ghPosition) GetPosition() string {
	if o == nil || o.Position == nil {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghPosition) GetPositionOk() (string, bool) {
	if o == nil || o.Position == nil {
		var ret string
		return ret, false
	}
	return *o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Ga4ghPosition) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *Ga4ghPosition) SetPosition(v string) {
	o.Position = &v
}

// GetStrand returns the Strand field if non-nil, zero value otherwise.
func (o *Ga4ghPosition) GetStrand() Ga4ghStrand {
	if o == nil || o.Strand == nil {
		var ret Ga4ghStrand
		return ret
	}
	return *o.Strand
}

// GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghPosition) GetStrandOk() (Ga4ghStrand, bool) {
	if o == nil || o.Strand == nil {
		var ret Ga4ghStrand
		return ret, false
	}
	return *o.Strand, true
}

// HasStrand returns a boolean if a field has been set.
func (o *Ga4ghPosition) HasStrand() bool {
	if o != nil && o.Strand != nil {
		return true
	}

	return false
}

// SetStrand gets a reference to the given Ga4ghStrand and assigns it to the Strand field.
func (o *Ga4ghPosition) SetStrand(v Ga4ghStrand) {
	o.Strand = &v
}


func (o Ga4ghPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceName != nil {
		toSerialize["reference_name"] = o.ReferenceName
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Strand != nil {
		toSerialize["strand"] = o.Strand
	}
	return json.Marshal(toSerialize)
}


