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

type Terms struct {

	NoAuthorizationTerms bool `json:"noAuthorizationTerms,omitempty"`

	WhichAuthorizationTerms []string `json:"whichAuthorizationTerms,omitempty"`

	NoPublicationTerms bool `json:"noPublicationTerms,omitempty"`

	WhichPublicationTerms []string `json:"whichPublicationTerms,omitempty"`

	NoTimelineTerms bool `json:"noTimelineTerms,omitempty"`

	WhichTimelineTerms []string `json:"whichTimelineTerms,omitempty"`

	NoSecurityTerms bool `json:"noSecurityTerms,omitempty"`

	WhichSecurityTerms []string `json:"whichSecurityTerms,omitempty"`

	NoExpungingTerms bool `json:"noExpungingTerms,omitempty"`

	WhichExpungingTerms []string `json:"whichExpungingTerms,omitempty"`

	NoLinkingTerms bool `json:"noLinkingTerms,omitempty"`

	WhichLinkingTerms []string `json:"whichLinkingTerms,omitempty"`

	NoRecontactTerms bool `json:"noRecontactTerms,omitempty"`

	AllowedRecontactTerms []string `json:"allowedRecontactTerms,omitempty"`

	CompulsoryRecontactTerms []string `json:"compulsoryRecontactTerms,omitempty"`

	NoIPClaimTerms bool `json:"noIPClaimTerms,omitempty"`

	WhichIPClaimTerms []string `json:"whichIPClaimTerms,omitempty"`

	NoReportingTerms bool `json:"noReportingTerms,omitempty"`

	WhichReportingTerms []string `json:"whichReportingTerms,omitempty"`

	NoCollaborationTerms bool `json:"noCollaborationTerms,omitempty"`

	WhichCollaborationTerms []string `json:"whichCollaborationTerms,omitempty"`

	NoPaymentTerms bool `json:"noPaymentTerms,omitempty"`

	WhichPaymentTerms []string `json:"whichPaymentTerms,omitempty"`
}