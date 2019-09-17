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

// Each read alignment describes an alignment with additional information about the fragment and the read. A read alignment object is equivalent to a line in a SAM file.
type Ga4ghReadAlignment struct {
	// The read alignment ID. This ID is unique within the read group this alignment belongs to.  For performance reasons, this field may be omitted by a backend. If provided, its intended use is to make caching and UI display easier for genome browsers and other lightweight clients.
	Id *string `json:"id,omitempty"`

	ReadGroupId *string `json:"read_group_id,omitempty"`

	// The fragment name. Equivalent to QNAME (query template name) in SAM.
	FragmentName *string `json:"fragment_name,omitempty"`

	// The orientation and the distance between reads from the fragment are inconsistent with the sequencing protocol (inverse of SAM flag 0x2).
	ImproperPlacement *bool `json:"improper_placement,omitempty"`

	// The fragment is a PCR or optical duplicate (SAM flag 0x400).
	DuplicateFragment *bool `json:"duplicate_fragment,omitempty"`

	// The number of reads in the fragment (extension to SAM flag 0x1).
	NumberReads *int32 `json:"number_reads,omitempty"`

	// The observed length of the fragment, equivalent to TLEN in SAM.
	FragmentLength *int32 `json:"fragment_length,omitempty"`

	// The read ordinal in the fragment, 0-based and less than numberReads. This field replaces SAM flag 0x40 and 0x80 and is intended to more cleanly represent multiple reads per fragment.
	ReadNumber *int32 `json:"read_number,omitempty"`

	// The read fails platform or vendor quality checks (SAM flag 0x200).
	FailedVendorQualityChecks *bool `json:"failed_vendor_quality_checks,omitempty"`

	Alignment *Ga4ghLinearAlignment `json:"alignment,omitempty"`

	// Whether this alignment is secondary. Equivalent to SAM flag 0x100. A secondary alignment represents an alternative to the primary alignment for this read. Aligners may return secondary alignments if a read can map ambiguously to multiple coordinates in the genome.  By convention, each read has one and only one alignment where both secondaryAlignment and supplementaryAlignment are false.
	SecondaryAlignment *bool `json:"secondary_alignment,omitempty"`

	// Whether this alignment is supplementary. Equivalent to SAM flag 0x800. Supplementary alignments are used in the representation of a chimeric alignment. In a chimeric alignment, a read is split into multiple linear alignments that map to different reference contigs. The first linear alignment in the read will be designated as the representative alignment; the remaining linear alignments will be designated as supplementary alignments. These alignments may have different mapping quality scores.  In each linear alignment in a chimeric alignment, the read will be hard clipped. The `alignedSequence` and `alignedQuality` fields in the alignment message will only represent the bases for its respective linear alignment.
	SupplementaryAlignment *bool `json:"supplementary_alignment,omitempty"`

	// The bases of the read sequence contained in this alignment record (equivalent to SEQ in SAM).  `alignedSequence` and `alignedQuality` may be shorter than the full read sequence and quality. This will occur if the alignment is part of a chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR for this read will begin/end with a hard clip operator that will indicate the length of the excised sequence.
	AlignedSequence *string `json:"aligned_sequence,omitempty"`

	// The quality of the read sequence contained in this alignment message (equivalent to QUAL in SAM).  `alignedSequence` and `alignedQuality` may be shorter than the full read sequence and quality. This will occur if the alignment is part of a chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR for this read will begin/end with a hard clip operator that will indicate the length of the excised sequence.
	AlignedQuality *[]int32 `json:"aligned_quality,omitempty"`

	NextMatePosition *Ga4ghPosition `json:"next_mate_position,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghReadAlignment) SetId(v string) {
	o.Id = &v
}

// GetReadGroupId returns the ReadGroupId field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetReadGroupId() string {
	if o == nil || o.ReadGroupId == nil {
		var ret string
		return ret
	}
	return *o.ReadGroupId
}

// GetReadGroupIdOk returns a tuple with the ReadGroupId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetReadGroupIdOk() (string, bool) {
	if o == nil || o.ReadGroupId == nil {
		var ret string
		return ret, false
	}
	return *o.ReadGroupId, true
}

// HasReadGroupId returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasReadGroupId() bool {
	if o != nil && o.ReadGroupId != nil {
		return true
	}

	return false
}

// SetReadGroupId gets a reference to the given string and assigns it to the ReadGroupId field.
func (o *Ga4ghReadAlignment) SetReadGroupId(v string) {
	o.ReadGroupId = &v
}

// GetFragmentName returns the FragmentName field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetFragmentName() string {
	if o == nil || o.FragmentName == nil {
		var ret string
		return ret
	}
	return *o.FragmentName
}

// GetFragmentNameOk returns a tuple with the FragmentName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetFragmentNameOk() (string, bool) {
	if o == nil || o.FragmentName == nil {
		var ret string
		return ret, false
	}
	return *o.FragmentName, true
}

// HasFragmentName returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasFragmentName() bool {
	if o != nil && o.FragmentName != nil {
		return true
	}

	return false
}

// SetFragmentName gets a reference to the given string and assigns it to the FragmentName field.
func (o *Ga4ghReadAlignment) SetFragmentName(v string) {
	o.FragmentName = &v
}

// GetImproperPlacement returns the ImproperPlacement field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetImproperPlacement() bool {
	if o == nil || o.ImproperPlacement == nil {
		var ret bool
		return ret
	}
	return *o.ImproperPlacement
}

// GetImproperPlacementOk returns a tuple with the ImproperPlacement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetImproperPlacementOk() (bool, bool) {
	if o == nil || o.ImproperPlacement == nil {
		var ret bool
		return ret, false
	}
	return *o.ImproperPlacement, true
}

// HasImproperPlacement returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasImproperPlacement() bool {
	if o != nil && o.ImproperPlacement != nil {
		return true
	}

	return false
}

// SetImproperPlacement gets a reference to the given bool and assigns it to the ImproperPlacement field.
func (o *Ga4ghReadAlignment) SetImproperPlacement(v bool) {
	o.ImproperPlacement = &v
}

// GetDuplicateFragment returns the DuplicateFragment field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetDuplicateFragment() bool {
	if o == nil || o.DuplicateFragment == nil {
		var ret bool
		return ret
	}
	return *o.DuplicateFragment
}

// GetDuplicateFragmentOk returns a tuple with the DuplicateFragment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetDuplicateFragmentOk() (bool, bool) {
	if o == nil || o.DuplicateFragment == nil {
		var ret bool
		return ret, false
	}
	return *o.DuplicateFragment, true
}

// HasDuplicateFragment returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasDuplicateFragment() bool {
	if o != nil && o.DuplicateFragment != nil {
		return true
	}

	return false
}

// SetDuplicateFragment gets a reference to the given bool and assigns it to the DuplicateFragment field.
func (o *Ga4ghReadAlignment) SetDuplicateFragment(v bool) {
	o.DuplicateFragment = &v
}

// GetNumberReads returns the NumberReads field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetNumberReads() int32 {
	if o == nil || o.NumberReads == nil {
		var ret int32
		return ret
	}
	return *o.NumberReads
}

// GetNumberReadsOk returns a tuple with the NumberReads field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetNumberReadsOk() (int32, bool) {
	if o == nil || o.NumberReads == nil {
		var ret int32
		return ret, false
	}
	return *o.NumberReads, true
}

// HasNumberReads returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasNumberReads() bool {
	if o != nil && o.NumberReads != nil {
		return true
	}

	return false
}

// SetNumberReads gets a reference to the given int32 and assigns it to the NumberReads field.
func (o *Ga4ghReadAlignment) SetNumberReads(v int32) {
	o.NumberReads = &v
}

// GetFragmentLength returns the FragmentLength field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetFragmentLength() int32 {
	if o == nil || o.FragmentLength == nil {
		var ret int32
		return ret
	}
	return *o.FragmentLength
}

// GetFragmentLengthOk returns a tuple with the FragmentLength field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetFragmentLengthOk() (int32, bool) {
	if o == nil || o.FragmentLength == nil {
		var ret int32
		return ret, false
	}
	return *o.FragmentLength, true
}

// HasFragmentLength returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasFragmentLength() bool {
	if o != nil && o.FragmentLength != nil {
		return true
	}

	return false
}

// SetFragmentLength gets a reference to the given int32 and assigns it to the FragmentLength field.
func (o *Ga4ghReadAlignment) SetFragmentLength(v int32) {
	o.FragmentLength = &v
}

// GetReadNumber returns the ReadNumber field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetReadNumber() int32 {
	if o == nil || o.ReadNumber == nil {
		var ret int32
		return ret
	}
	return *o.ReadNumber
}

// GetReadNumberOk returns a tuple with the ReadNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetReadNumberOk() (int32, bool) {
	if o == nil || o.ReadNumber == nil {
		var ret int32
		return ret, false
	}
	return *o.ReadNumber, true
}

// HasReadNumber returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasReadNumber() bool {
	if o != nil && o.ReadNumber != nil {
		return true
	}

	return false
}

// SetReadNumber gets a reference to the given int32 and assigns it to the ReadNumber field.
func (o *Ga4ghReadAlignment) SetReadNumber(v int32) {
	o.ReadNumber = &v
}

// GetFailedVendorQualityChecks returns the FailedVendorQualityChecks field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetFailedVendorQualityChecks() bool {
	if o == nil || o.FailedVendorQualityChecks == nil {
		var ret bool
		return ret
	}
	return *o.FailedVendorQualityChecks
}

// GetFailedVendorQualityChecksOk returns a tuple with the FailedVendorQualityChecks field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetFailedVendorQualityChecksOk() (bool, bool) {
	if o == nil || o.FailedVendorQualityChecks == nil {
		var ret bool
		return ret, false
	}
	return *o.FailedVendorQualityChecks, true
}

// HasFailedVendorQualityChecks returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasFailedVendorQualityChecks() bool {
	if o != nil && o.FailedVendorQualityChecks != nil {
		return true
	}

	return false
}

// SetFailedVendorQualityChecks gets a reference to the given bool and assigns it to the FailedVendorQualityChecks field.
func (o *Ga4ghReadAlignment) SetFailedVendorQualityChecks(v bool) {
	o.FailedVendorQualityChecks = &v
}

// GetAlignment returns the Alignment field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetAlignment() Ga4ghLinearAlignment {
	if o == nil || o.Alignment == nil {
		var ret Ga4ghLinearAlignment
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetAlignmentOk() (Ga4ghLinearAlignment, bool) {
	if o == nil || o.Alignment == nil {
		var ret Ga4ghLinearAlignment
		return ret, false
	}
	return *o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasAlignment() bool {
	if o != nil && o.Alignment != nil {
		return true
	}

	return false
}

// SetAlignment gets a reference to the given Ga4ghLinearAlignment and assigns it to the Alignment field.
func (o *Ga4ghReadAlignment) SetAlignment(v Ga4ghLinearAlignment) {
	o.Alignment = &v
}

// GetSecondaryAlignment returns the SecondaryAlignment field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetSecondaryAlignment() bool {
	if o == nil || o.SecondaryAlignment == nil {
		var ret bool
		return ret
	}
	return *o.SecondaryAlignment
}

// GetSecondaryAlignmentOk returns a tuple with the SecondaryAlignment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetSecondaryAlignmentOk() (bool, bool) {
	if o == nil || o.SecondaryAlignment == nil {
		var ret bool
		return ret, false
	}
	return *o.SecondaryAlignment, true
}

// HasSecondaryAlignment returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasSecondaryAlignment() bool {
	if o != nil && o.SecondaryAlignment != nil {
		return true
	}

	return false
}

// SetSecondaryAlignment gets a reference to the given bool and assigns it to the SecondaryAlignment field.
func (o *Ga4ghReadAlignment) SetSecondaryAlignment(v bool) {
	o.SecondaryAlignment = &v
}

// GetSupplementaryAlignment returns the SupplementaryAlignment field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetSupplementaryAlignment() bool {
	if o == nil || o.SupplementaryAlignment == nil {
		var ret bool
		return ret
	}
	return *o.SupplementaryAlignment
}

// GetSupplementaryAlignmentOk returns a tuple with the SupplementaryAlignment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetSupplementaryAlignmentOk() (bool, bool) {
	if o == nil || o.SupplementaryAlignment == nil {
		var ret bool
		return ret, false
	}
	return *o.SupplementaryAlignment, true
}

// HasSupplementaryAlignment returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasSupplementaryAlignment() bool {
	if o != nil && o.SupplementaryAlignment != nil {
		return true
	}

	return false
}

// SetSupplementaryAlignment gets a reference to the given bool and assigns it to the SupplementaryAlignment field.
func (o *Ga4ghReadAlignment) SetSupplementaryAlignment(v bool) {
	o.SupplementaryAlignment = &v
}

// GetAlignedSequence returns the AlignedSequence field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetAlignedSequence() string {
	if o == nil || o.AlignedSequence == nil {
		var ret string
		return ret
	}
	return *o.AlignedSequence
}

// GetAlignedSequenceOk returns a tuple with the AlignedSequence field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetAlignedSequenceOk() (string, bool) {
	if o == nil || o.AlignedSequence == nil {
		var ret string
		return ret, false
	}
	return *o.AlignedSequence, true
}

// HasAlignedSequence returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasAlignedSequence() bool {
	if o != nil && o.AlignedSequence != nil {
		return true
	}

	return false
}

// SetAlignedSequence gets a reference to the given string and assigns it to the AlignedSequence field.
func (o *Ga4ghReadAlignment) SetAlignedSequence(v string) {
	o.AlignedSequence = &v
}

// GetAlignedQuality returns the AlignedQuality field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetAlignedQuality() []int32 {
	if o == nil || o.AlignedQuality == nil {
		var ret []int32
		return ret
	}
	return *o.AlignedQuality
}

// GetAlignedQualityOk returns a tuple with the AlignedQuality field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetAlignedQualityOk() ([]int32, bool) {
	if o == nil || o.AlignedQuality == nil {
		var ret []int32
		return ret, false
	}
	return *o.AlignedQuality, true
}

// HasAlignedQuality returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasAlignedQuality() bool {
	if o != nil && o.AlignedQuality != nil {
		return true
	}

	return false
}

// SetAlignedQuality gets a reference to the given []int32 and assigns it to the AlignedQuality field.
func (o *Ga4ghReadAlignment) SetAlignedQuality(v []int32) {
	o.AlignedQuality = &v
}

// GetNextMatePosition returns the NextMatePosition field if non-nil, zero value otherwise.
func (o *Ga4ghReadAlignment) GetNextMatePosition() Ga4ghPosition {
	if o == nil || o.NextMatePosition == nil {
		var ret Ga4ghPosition
		return ret
	}
	return *o.NextMatePosition
}

// GetNextMatePositionOk returns a tuple with the NextMatePosition field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReadAlignment) GetNextMatePositionOk() (Ga4ghPosition, bool) {
	if o == nil || o.NextMatePosition == nil {
		var ret Ga4ghPosition
		return ret, false
	}
	return *o.NextMatePosition, true
}

// HasNextMatePosition returns a boolean if a field has been set.
func (o *Ga4ghReadAlignment) HasNextMatePosition() bool {
	if o != nil && o.NextMatePosition != nil {
		return true
	}

	return false
}

// SetNextMatePosition gets a reference to the given Ga4ghPosition and assigns it to the NextMatePosition field.
func (o *Ga4ghReadAlignment) SetNextMatePosition(v Ga4ghPosition) {
	o.NextMatePosition = &v
}


func (o Ga4ghReadAlignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ReadGroupId != nil {
		toSerialize["read_group_id"] = o.ReadGroupId
	}
	if o.FragmentName != nil {
		toSerialize["fragment_name"] = o.FragmentName
	}
	if o.ImproperPlacement != nil {
		toSerialize["improper_placement"] = o.ImproperPlacement
	}
	if o.DuplicateFragment != nil {
		toSerialize["duplicate_fragment"] = o.DuplicateFragment
	}
	if o.NumberReads != nil {
		toSerialize["number_reads"] = o.NumberReads
	}
	if o.FragmentLength != nil {
		toSerialize["fragment_length"] = o.FragmentLength
	}
	if o.ReadNumber != nil {
		toSerialize["read_number"] = o.ReadNumber
	}
	if o.FailedVendorQualityChecks != nil {
		toSerialize["failed_vendor_quality_checks"] = o.FailedVendorQualityChecks
	}
	if o.Alignment != nil {
		toSerialize["alignment"] = o.Alignment
	}
	if o.SecondaryAlignment != nil {
		toSerialize["secondary_alignment"] = o.SecondaryAlignment
	}
	if o.SupplementaryAlignment != nil {
		toSerialize["supplementary_alignment"] = o.SupplementaryAlignment
	}
	if o.AlignedSequence != nil {
		toSerialize["aligned_sequence"] = o.AlignedSequence
	}
	if o.AlignedQuality != nil {
		toSerialize["aligned_quality"] = o.AlignedQuality
	}
	if o.NextMatePosition != nil {
		toSerialize["next_mate_position"] = o.NextMatePosition
	}
	return json.Marshal(toSerialize)
}

