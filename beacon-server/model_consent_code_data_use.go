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

// ConsentCodeDataUse - Data use of a resource based on consent codes.
type ConsentCodeDataUse struct {

	PrimaryCategory ConsentCodeDataUseCondition `json:"primaryCategory"`

	SecondaryCategories []ConsentCodeDataUseCondition `json:"secondaryCategories,omitempty"`

	Requirements []ConsentCodeDataUseCondition `json:"requirements,omitempty"`

	// Version of the data use specification.
	Version string `json:"version"`
}