/*
 * CanDIG Services
 *
 * Below is a list of APIs that CanDIG currently supports.<br/><br/>For /search and /count endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/search_count_services_usage.pdf) for instructions and sample queries.<br/>For all metadata and variant services endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/metadata_variants_services_sample_queries.pdf) for sample queries.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// CigarUnitOperation : Describes the different types of CIGAR alignment operations that exist. Used wherever CIGAR alignments are used.   - OPERATION_UNSPECIFIED: The null operation.  - ALIGNMENT_MATCH: An alignment match indicates that a sequence can be aligned to the reference without evidence of an INDEL. Unlike the `SEQUENCE_MATCH` and `SEQUENCE_MISMATCH` operators, the `ALIGNMENT_MATCH` operator does not indicate whether the reference and read sequences are an exact match. This operator is equivalent to SAM's `M`.  - INSERT: The insert operator indicates that the read contains evidence of bases being inserted into the reference. This operator is equivalent to SAM's `I`.  - DELETE: The delete operator indicates that the read contains evidence of bases being deleted from the reference. This operator is equivalent to SAM's `D`.  - SKIP: The skip operator indicates that this read skips a long segment of the reference, but the bases have not been deleted. This operator is commonly used when working with RNA-seq data, where reads may skip long segments of the reference between exons. This operator is equivalent to SAM's `N`.  - CLIP_SOFT: The soft clip operator indicates that bases at the start/end of a read have not been considered during alignment. This may occur if the majority of a read maps, except for low quality bases at the start/end of a read. This operator is equivalent to SAM's `S`. Bases that are soft clipped will still be stored in the read.  - CLIP_HARD: The hard clip operator indicates that bases at the start/end of a read have been omitted from this alignment. This may occur if this linear alignment is part of a chimeric alignment, or if the read has been trimmed (for example, during error correction or to trim poly-A tails for RNA-seq). This operator is equivalent to SAM's `H`.  - PAD: The pad operator indicates that there is padding in an alignment. This operator is equivalent to SAM's `P`.  - SEQUENCE_MATCH: This operator indicates that this portion of the aligned sequence exactly matches the reference. This operator is equivalent to SAM's `=`.  - SEQUENCE_MISMATCH: This operator indicates that this portion of the aligned sequence is an alignment match to the reference, but a sequence mismatch. This can indicate a SNP or a read error. This operator is equivalent to SAM's `X`.
type CigarUnitOperation string

// List of CigarUnitOperation
const (
	OPERATION_UNSPECIFIED CigarUnitOperation = "OPERATION_UNSPECIFIED"
	ALIGNMENT_MATCH CigarUnitOperation = "ALIGNMENT_MATCH"
	INSERT CigarUnitOperation = "INSERT"
	DELETE CigarUnitOperation = "DELETE"
	SKIP CigarUnitOperation = "SKIP"
	CLIP_SOFT CigarUnitOperation = "CLIP_SOFT"
	CLIP_HARD CigarUnitOperation = "CLIP_HARD"
	PAD CigarUnitOperation = "PAD"
	SEQUENCE_MATCH CigarUnitOperation = "SEQUENCE_MATCH"
	SEQUENCE_MISMATCH CigarUnitOperation = "SEQUENCE_MISMATCH"
)

