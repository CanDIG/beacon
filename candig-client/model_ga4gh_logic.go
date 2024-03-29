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

type Ga4ghLogic struct {
	And *[]Ga4ghLogic `json:"and,omitempty"`

	Or *[]Ga4ghLogic `json:"or,omitempty"`

	Id *string `json:"id,omitempty"`

	Negate *bool `json:"negate,omitempty"`

}

// GetAnd returns the And field if non-nil, zero value otherwise.
func (o *Ga4ghLogic) GetAnd() []Ga4ghLogic {
	if o == nil || o.And == nil {
		var ret []Ga4ghLogic
		return ret
	}
	return *o.And
}

// GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLogic) GetAndOk() ([]Ga4ghLogic, bool) {
	if o == nil || o.And == nil {
		var ret []Ga4ghLogic
		return ret, false
	}
	return *o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *Ga4ghLogic) HasAnd() bool {
	if o != nil && o.And != nil {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []Ga4ghLogic and assigns it to the And field.
func (o *Ga4ghLogic) SetAnd(v []Ga4ghLogic) {
	o.And = &v
}

// GetOr returns the Or field if non-nil, zero value otherwise.
func (o *Ga4ghLogic) GetOr() []Ga4ghLogic {
	if o == nil || o.Or == nil {
		var ret []Ga4ghLogic
		return ret
	}
	return *o.Or
}

// GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLogic) GetOrOk() ([]Ga4ghLogic, bool) {
	if o == nil || o.Or == nil {
		var ret []Ga4ghLogic
		return ret, false
	}
	return *o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *Ga4ghLogic) HasOr() bool {
	if o != nil && o.Or != nil {
		return true
	}

	return false
}

// SetOr gets a reference to the given []Ga4ghLogic and assigns it to the Or field.
func (o *Ga4ghLogic) SetOr(v []Ga4ghLogic) {
	o.Or = &v
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghLogic) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLogic) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghLogic) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghLogic) SetId(v string) {
	o.Id = &v
}

// GetNegate returns the Negate field if non-nil, zero value otherwise.
func (o *Ga4ghLogic) GetNegate() bool {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghLogic) GetNegateOk() (bool, bool) {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret, false
	}
	return *o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *Ga4ghLogic) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *Ga4ghLogic) SetNegate(v bool) {
	o.Negate = &v
}


func (o Ga4ghLogic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.And != nil {
		toSerialize["and"] = o.And
	}
	if o.Or != nil {
		toSerialize["or"] = o.Or
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}
	return json.Marshal(toSerialize)
}


