/*
 * GA4GH Beacon API Specification
 *
 * A Beacon is a web service for genetic data sharing that can be queried for information about specific alleles.
 *
 * API version: 1.0.1
 * Contact: beacon@ga4gh.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

// ConsentCodeDataUseCondition - Data use condition.
type ConsentCodeDataUseCondition struct {

	// Consent code abbreviation.
	Code string `json:"code"`

	// Description of the condition.
	Description string `json:"description,omitempty"`
}
