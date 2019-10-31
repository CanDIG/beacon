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

// Ga4ghOntologyTerm struct for Ga4ghOntologyTerm
type Ga4ghOntologyTerm struct {
	// Ontology term identifier - the CURIE for an ontology term. It differs from the standard GA4GH schema's :ref:`id <apidesign_object_ids>` in that it is a CURIE pointing to an information resource outside of the scope of the schema or its resource implementation.
	TermId *string `json:"term_id,omitempty"`

	// Ontology term - the label of the ontology term the termId is pointing to.
	Term *string `json:"term,omitempty"`
}

// GetTermId returns the TermId field if non-nil, zero value otherwise.
func (o *Ga4ghOntologyTerm) GetTermId() string {
	if o == nil || o.TermId == nil {
		var ret string
		return ret
	}
	return *o.TermId
}

// GetTermIdOk returns a tuple with the TermId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOntologyTerm) GetTermIdOk() (string, bool) {
	if o == nil || o.TermId == nil {
		var ret string
		return ret, false
	}
	return *o.TermId, true
}

// HasTermId returns a boolean if a field has been set.
func (o *Ga4ghOntologyTerm) HasTermId() bool {
	if o != nil && o.TermId != nil {
		return true
	}

	return false
}

// SetTermId gets a reference to the given string and assigns it to the TermId field.
func (o *Ga4ghOntologyTerm) SetTermId(v string) {
	o.TermId = &v
}

// GetTerm returns the Term field if non-nil, zero value otherwise.
func (o *Ga4ghOntologyTerm) GetTerm() string {
	if o == nil || o.Term == nil {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghOntologyTerm) GetTermOk() (string, bool) {
	if o == nil || o.Term == nil {
		var ret string
		return ret, false
	}
	return *o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *Ga4ghOntologyTerm) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *Ga4ghOntologyTerm) SetTerm(v string) {
	o.Term = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Ga4ghOntologyTerm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TermId != nil {
		toSerialize["term_id"] = o.TermId
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	return json.Marshal(toSerialize)
}
