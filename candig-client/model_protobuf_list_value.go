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

// `ListValue` is a wrapper around a repeated field of values.  The JSON representation for `ListValue` is JSON array.
type ProtobufListValue struct {
	// Repeated field of dynamically typed values.
	Values *[]ProtobufValue `json:"values,omitempty"`

}

// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *ProtobufListValue) GetValues() []ProtobufValue {
	if o == nil || o.Values == nil {
		var ret []ProtobufValue
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ProtobufListValue) GetValuesOk() ([]ProtobufValue, bool) {
	if o == nil || o.Values == nil {
		var ret []ProtobufValue
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ProtobufListValue) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []ProtobufValue and assigns it to the Values field.
func (o *ProtobufListValue) SetValues(v []ProtobufValue) {
	o.Values = &v
}


func (o ProtobufListValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}


