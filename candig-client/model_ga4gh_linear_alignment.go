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

// A linear alignment describes the alignment of a read to a Reference, using a position and CIGAR array.
type Ga4ghLinearAlignment struct {
	Position *Ga4ghPosition `json:"position,omitempty"`

	// The mapping quality of this alignment, meaning the likelihood that the read maps to this position.  Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to the nearest integer.
	MappingQuality *int32 `json:"mapping_quality,omitempty"`

	// Represents the local alignment of this sequence (alignment matches, indels, etc) versus the reference.
	Cigar *[]Ga4ghCigarUnit `json:"cigar,omitempty"`

}

// GetPosition returns the Position field if non-nil, zero value otherwise.
func (o *Ga4ghLinearAlignment) GetPosition() Ga4ghPosition {
	if o == nil || o.Position == nil {
		var ret Ga4ghPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLinearAlignment) GetPositionOk() (Ga4ghPosition, bool) {
	if o == nil || o.Position == nil {
		var ret Ga4ghPosition
		return ret, false
	}
	return *o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Ga4ghLinearAlignment) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given Ga4ghPosition and assigns it to the Position field.
func (o *Ga4ghLinearAlignment) SetPosition(v Ga4ghPosition) {
	o.Position = &v
}

// GetMappingQuality returns the MappingQuality field if non-nil, zero value otherwise.
func (o *Ga4ghLinearAlignment) GetMappingQuality() int32 {
	if o == nil || o.MappingQuality == nil {
		var ret int32
		return ret
	}
	return *o.MappingQuality
}

// GetMappingQualityOk returns a tuple with the MappingQuality field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLinearAlignment) GetMappingQualityOk() (int32, bool) {
	if o == nil || o.MappingQuality == nil {
		var ret int32
		return ret, false
	}
	return *o.MappingQuality, true
}

// HasMappingQuality returns a boolean if a field has been set.
func (o *Ga4ghLinearAlignment) HasMappingQuality() bool {
	if o != nil && o.MappingQuality != nil {
		return true
	}

	return false
}

// SetMappingQuality gets a reference to the given int32 and assigns it to the MappingQuality field.
func (o *Ga4ghLinearAlignment) SetMappingQuality(v int32) {
	o.MappingQuality = &v
}

// GetCigar returns the Cigar field if non-nil, zero value otherwise.
func (o *Ga4ghLinearAlignment) GetCigar() []Ga4ghCigarUnit {
	if o == nil || o.Cigar == nil {
		var ret []Ga4ghCigarUnit
		return ret
	}
	return *o.Cigar
}

// GetCigarOk returns a tuple with the Cigar field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLinearAlignment) GetCigarOk() ([]Ga4ghCigarUnit, bool) {
	if o == nil || o.Cigar == nil {
		var ret []Ga4ghCigarUnit
		return ret, false
	}
	return *o.Cigar, true
}

// HasCigar returns a boolean if a field has been set.
func (o *Ga4ghLinearAlignment) HasCigar() bool {
	if o != nil && o.Cigar != nil {
		return true
	}

	return false
}

// SetCigar gets a reference to the given []Ga4ghCigarUnit and assigns it to the Cigar field.
func (o *Ga4ghLinearAlignment) SetCigar(v []Ga4ghCigarUnit) {
	o.Cigar = &v
}


func (o Ga4ghLinearAlignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.MappingQuality != nil {
		toSerialize["mapping_quality"] = o.MappingQuality
	}
	if o.Cigar != nil {
		toSerialize["cigar"] = o.Cigar
	}
	return json.Marshal(toSerialize)
}


