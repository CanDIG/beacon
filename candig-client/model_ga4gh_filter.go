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

// Ga4ghFilter struct for Ga4ghFilter
type Ga4ghFilter struct {
	Field *string `json:"field,omitempty"`

	Operator *string `json:"operator,omitempty"`

	Value *string `json:"value,omitempty"`

	Values *[]string `json:"values,omitempty"`
}

// GetField returns the Field field if non-nil, zero value otherwise.
func (o *Ga4ghFilter) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghFilter) GetFieldOk() (string, bool) {
	if o == nil || o.Field == nil {
		var ret string
		return ret, false
	}
	return *o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *Ga4ghFilter) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *Ga4ghFilter) SetField(v string) {
	o.Field = &v
}

// GetOperator returns the Operator field if non-nil, zero value otherwise.
func (o *Ga4ghFilter) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghFilter) GetOperatorOk() (string, bool) {
	if o == nil || o.Operator == nil {
		var ret string
		return ret, false
	}
	return *o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Ga4ghFilter) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *Ga4ghFilter) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *Ga4ghFilter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghFilter) GetValueOk() (string, bool) {
	if o == nil || o.Value == nil {
		var ret string
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Ga4ghFilter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Ga4ghFilter) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *Ga4ghFilter) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghFilter) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		var ret []string
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Ga4ghFilter) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *Ga4ghFilter) SetValues(v []string) {
	o.Values = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Ga4ghFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}
