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

// A single CIGAR operation.
type Ga4ghCigarUnit struct {
	Operation *CigarUnitOperation `json:"operation,omitempty"`

	// The number of genomic bases that the operation runs for. Required.
	OperationLength *string `json:"operation_length,omitempty"`

	// `referenceSequence` is only used at mismatches (`SEQUENCE_MISMATCH`) and deletions (`DELETE`). Filling this field replaces SAM's MD tag. If the relevant information is not available, this field is unset.
	ReferenceSequence *string `json:"reference_sequence,omitempty"`

}

// GetOperation returns the Operation field if non-nil, zero value otherwise.
func (o *Ga4ghCigarUnit) GetOperation() CigarUnitOperation {
	if o == nil || o.Operation == nil {
		var ret CigarUnitOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCigarUnit) GetOperationOk() (CigarUnitOperation, bool) {
	if o == nil || o.Operation == nil {
		var ret CigarUnitOperation
		return ret, false
	}
	return *o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Ga4ghCigarUnit) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given CigarUnitOperation and assigns it to the Operation field.
func (o *Ga4ghCigarUnit) SetOperation(v CigarUnitOperation) {
	o.Operation = &v
}

// GetOperationLength returns the OperationLength field if non-nil, zero value otherwise.
func (o *Ga4ghCigarUnit) GetOperationLength() string {
	if o == nil || o.OperationLength == nil {
		var ret string
		return ret
	}
	return *o.OperationLength
}

// GetOperationLengthOk returns a tuple with the OperationLength field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCigarUnit) GetOperationLengthOk() (string, bool) {
	if o == nil || o.OperationLength == nil {
		var ret string
		return ret, false
	}
	return *o.OperationLength, true
}

// HasOperationLength returns a boolean if a field has been set.
func (o *Ga4ghCigarUnit) HasOperationLength() bool {
	if o != nil && o.OperationLength != nil {
		return true
	}

	return false
}

// SetOperationLength gets a reference to the given string and assigns it to the OperationLength field.
func (o *Ga4ghCigarUnit) SetOperationLength(v string) {
	o.OperationLength = &v
}

// GetReferenceSequence returns the ReferenceSequence field if non-nil, zero value otherwise.
func (o *Ga4ghCigarUnit) GetReferenceSequence() string {
	if o == nil || o.ReferenceSequence == nil {
		var ret string
		return ret
	}
	return *o.ReferenceSequence
}

// GetReferenceSequenceOk returns a tuple with the ReferenceSequence field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghCigarUnit) GetReferenceSequenceOk() (string, bool) {
	if o == nil || o.ReferenceSequence == nil {
		var ret string
		return ret, false
	}
	return *o.ReferenceSequence, true
}

// HasReferenceSequence returns a boolean if a field has been set.
func (o *Ga4ghCigarUnit) HasReferenceSequence() bool {
	if o != nil && o.ReferenceSequence != nil {
		return true
	}

	return false
}

// SetReferenceSequence gets a reference to the given string and assigns it to the ReferenceSequence field.
func (o *Ga4ghCigarUnit) SetReferenceSequence(v string) {
	o.ReferenceSequence = &v
}


func (o Ga4ghCigarUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.OperationLength != nil {
		toSerialize["operation_length"] = o.OperationLength
	}
	if o.ReferenceSequence != nil {
		toSerialize["reference_sequence"] = o.ReferenceSequence
	}
	return json.Marshal(toSerialize)
}


