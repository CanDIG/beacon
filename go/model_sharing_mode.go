/*
 * GA4GH Beacon API Specification
 *
 * A Beacon is a web service for genetic data sharing that can be queried for information about specific alleles.
 *
 * API version: 1.0.1
 * Contact: beacon@ga4gh.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SharingMode string

// List of SharingMode
const (
	UNKNOWNs             SharingMode = "UNKNOWN"
	DISCOVERY            SharingMode = "DISCOVERY"
	ACCESS               SharingMode = "ACCESS"
	DISCOVERY_AND_ACCESS SharingMode = "DISCOVERY_AND_ACCESS"
)
