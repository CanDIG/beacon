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

// BeaconAlleleRequest - Allele request as interpreted by the beacon.
type BeaconAlleleRequest struct {

	ReferenceName Chromosome `json:"referenceName"`

	// Precise start coordinate position, allele locus (0-based, inclusive). * start only:   - for single positions, e.g. the start of a specified sequence alteration where the size is given through the specified alternateBases   - typical use are queries for SNV and small InDels   - the use of \"start\" without an \"end\" parameter requires the use of \"referenceBases\" * start and end:   - special use case for exactly determined structural changes 
	Start int64 `json:"start,omitempty"`

	// Precise end coordinate (0-based, exclusive). See start.
	End int32 `json:"end,omitempty"`

	// Minimum start coordinate * startMin + startMax + endMin + endMax   - for querying imprecise positions (e.g. identifying all structural variants starting anywhere between startMin <-> startMax, and ending anywhere between endMin <-> endMax)   - single or double sided precise matches can be achieved by setting startMin = startMax XOR endMin = endMax 
	StartMin int32 `json:"startMin,omitempty"`

	// Maximum start coordinate. See startMin.
	StartMax int32 `json:"startMax,omitempty"`

	// Minimum end coordinate. See startMin.
	EndMin int32 `json:"endMin,omitempty"`

	// Maximum end coordinate. See startMin.
	EndMax int32 `json:"endMax,omitempty"`

	// Reference bases for this variant (starting from `start`). Accepted values: [ACGT]*   When querying for variants without specific base alterations (e.g. imprecise structural variants with separate variant_type as well as start_min & end_min ... parameters), the use of a single \"N\" value is required. 
	ReferenceBases string `json:"referenceBases"`

	// The bases that appear instead of the reference bases. Accepted values: [ACGT]* or N. Symbolic ALT alleles (DEL, INS, DUP, INV, CNV, DUP:TANDEM, DEL:ME, INS:ME) will be represented in `variantType`. Optional: either `alternateBases` or `variantType` is required. 
	AlternateBases string `json:"alternateBases,omitempty"`

	// The `variantType` is used to denote e.g. structural variants. Examples: * DUP: duplication of sequence following `start`; not necessarily in situ * DEL: deletion of sequence following `start`  Optional: either `alternateBases` or `variantType` is required. 
	VariantType string `json:"variantType,omitempty"`

	// Assembly identifier (GRC notation, e.g. `GRCh37`).
	AssemblyId string `json:"assemblyId"`

	// Identifiers of datasets, as defined in `BeaconDataset`. If this field is null/not specified, all datasets should be queried.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// Indicator of whether responses for individual datasets (datasetAlleleResponses) should be included in the response (BeaconAlleleResponse) to this request or not. If null (not specified), the default value of NONE is assumed.
	IncludeDatasetResponses string `json:"includeDatasetResponses,omitempty"`
}
