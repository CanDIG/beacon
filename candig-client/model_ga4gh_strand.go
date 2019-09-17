/*
 * CanDIG Services
 *
 * Below is a list of APIs that CanDIG currently supports.<br/><br/>For /search and /count endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/search_count_services_usage.pdf) for instructions and sample queries.<br/>For all metadata and variant services endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/metadata_variants_services_sample_queries.pdf) for sample queries.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Ga4ghStrand : Indicates the associated DNA strand for some data item.   - STRAND_UNSPECIFIED: If no strand data is available.  - NEG_STRAND: The negative (-) strand.  - POS_STRAND: The postive (+) strand.
type Ga4ghStrand string

// List of ga4ghStrand
const (
	STRAND_UNSPECIFIED Ga4ghStrand = "STRAND_UNSPECIFIED"
	NEG_STRAND Ga4ghStrand = "NEG_STRAND"
	POS_STRAND Ga4ghStrand = "POS_STRAND"
)
