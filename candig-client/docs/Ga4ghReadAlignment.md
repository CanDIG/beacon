# Ga4ghReadAlignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The read alignment ID. This ID is unique within the read group this alignment belongs to.  For performance reasons, this field may be omitted by a backend. If provided, its intended use is to make caching and UI display easier for genome browsers and other lightweight clients. | [optional] 
**ReadGroupId** | Pointer to **string** |  | [optional] 
**FragmentName** | Pointer to **string** | The fragment name. Equivalent to QNAME (query template name) in SAM. | [optional] 
**ImproperPlacement** | Pointer to **bool** | The orientation and the distance between reads from the fragment are inconsistent with the sequencing protocol (inverse of SAM flag 0x2). | [optional] 
**DuplicateFragment** | Pointer to **bool** | The fragment is a PCR or optical duplicate (SAM flag 0x400). | [optional] 
**NumberReads** | Pointer to **int32** | The number of reads in the fragment (extension to SAM flag 0x1). | [optional] 
**FragmentLength** | Pointer to **int32** | The observed length of the fragment, equivalent to TLEN in SAM. | [optional] 
**ReadNumber** | Pointer to **int32** | The read ordinal in the fragment, 0-based and less than numberReads. This field replaces SAM flag 0x40 and 0x80 and is intended to more cleanly represent multiple reads per fragment. | [optional] 
**FailedVendorQualityChecks** | Pointer to **bool** | The read fails platform or vendor quality checks (SAM flag 0x200). | [optional] 
**Alignment** | Pointer to [**Ga4ghLinearAlignment**](ga4ghLinearAlignment.md) |  | [optional] 
**SecondaryAlignment** | Pointer to **bool** | Whether this alignment is secondary. Equivalent to SAM flag 0x100. A secondary alignment represents an alternative to the primary alignment for this read. Aligners may return secondary alignments if a read can map ambiguously to multiple coordinates in the genome.  By convention, each read has one and only one alignment where both secondaryAlignment and supplementaryAlignment are false. | [optional] 
**SupplementaryAlignment** | Pointer to **bool** | Whether this alignment is supplementary. Equivalent to SAM flag 0x800. Supplementary alignments are used in the representation of a chimeric alignment. In a chimeric alignment, a read is split into multiple linear alignments that map to different reference contigs. The first linear alignment in the read will be designated as the representative alignment; the remaining linear alignments will be designated as supplementary alignments. These alignments may have different mapping quality scores.  In each linear alignment in a chimeric alignment, the read will be hard clipped. The &#x60;alignedSequence&#x60; and &#x60;alignedQuality&#x60; fields in the alignment message will only represent the bases for its respective linear alignment. | [optional] 
**AlignedSequence** | Pointer to **string** | The bases of the read sequence contained in this alignment record (equivalent to SEQ in SAM).  &#x60;alignedSequence&#x60; and &#x60;alignedQuality&#x60; may be shorter than the full read sequence and quality. This will occur if the alignment is part of a chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR for this read will begin/end with a hard clip operator that will indicate the length of the excised sequence. | [optional] 
**AlignedQuality** | Pointer to **[]int32** | The quality of the read sequence contained in this alignment message (equivalent to QUAL in SAM).  &#x60;alignedSequence&#x60; and &#x60;alignedQuality&#x60; may be shorter than the full read sequence and quality. This will occur if the alignment is part of a chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR for this read will begin/end with a hard clip operator that will indicate the length of the excised sequence. | [optional] 
**NextMatePosition** | Pointer to [**Ga4ghPosition**](ga4ghPosition.md) |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghReadAlignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghReadAlignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghReadAlignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghReadAlignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetReadGroupId

`func (o *Ga4ghReadAlignment) GetReadGroupId() string`

GetReadGroupId returns the ReadGroupId field if non-nil, zero value otherwise.

### GetReadGroupIdOk

`func (o *Ga4ghReadAlignment) GetReadGroupIdOk() (string, bool)`

GetReadGroupIdOk returns a tuple with the ReadGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadGroupId

`func (o *Ga4ghReadAlignment) HasReadGroupId() bool`

HasReadGroupId returns a boolean if a field has been set.

### SetReadGroupId

`func (o *Ga4ghReadAlignment) SetReadGroupId(v string)`

SetReadGroupId gets a reference to the given string and assigns it to the ReadGroupId field.

### GetFragmentName

`func (o *Ga4ghReadAlignment) GetFragmentName() string`

GetFragmentName returns the FragmentName field if non-nil, zero value otherwise.

### GetFragmentNameOk

`func (o *Ga4ghReadAlignment) GetFragmentNameOk() (string, bool)`

GetFragmentNameOk returns a tuple with the FragmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFragmentName

`func (o *Ga4ghReadAlignment) HasFragmentName() bool`

HasFragmentName returns a boolean if a field has been set.

### SetFragmentName

`func (o *Ga4ghReadAlignment) SetFragmentName(v string)`

SetFragmentName gets a reference to the given string and assigns it to the FragmentName field.

### GetImproperPlacement

`func (o *Ga4ghReadAlignment) GetImproperPlacement() bool`

GetImproperPlacement returns the ImproperPlacement field if non-nil, zero value otherwise.

### GetImproperPlacementOk

`func (o *Ga4ghReadAlignment) GetImproperPlacementOk() (bool, bool)`

GetImproperPlacementOk returns a tuple with the ImproperPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImproperPlacement

`func (o *Ga4ghReadAlignment) HasImproperPlacement() bool`

HasImproperPlacement returns a boolean if a field has been set.

### SetImproperPlacement

`func (o *Ga4ghReadAlignment) SetImproperPlacement(v bool)`

SetImproperPlacement gets a reference to the given bool and assigns it to the ImproperPlacement field.

### GetDuplicateFragment

`func (o *Ga4ghReadAlignment) GetDuplicateFragment() bool`

GetDuplicateFragment returns the DuplicateFragment field if non-nil, zero value otherwise.

### GetDuplicateFragmentOk

`func (o *Ga4ghReadAlignment) GetDuplicateFragmentOk() (bool, bool)`

GetDuplicateFragmentOk returns a tuple with the DuplicateFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDuplicateFragment

`func (o *Ga4ghReadAlignment) HasDuplicateFragment() bool`

HasDuplicateFragment returns a boolean if a field has been set.

### SetDuplicateFragment

`func (o *Ga4ghReadAlignment) SetDuplicateFragment(v bool)`

SetDuplicateFragment gets a reference to the given bool and assigns it to the DuplicateFragment field.

### GetNumberReads

`func (o *Ga4ghReadAlignment) GetNumberReads() int32`

GetNumberReads returns the NumberReads field if non-nil, zero value otherwise.

### GetNumberReadsOk

`func (o *Ga4ghReadAlignment) GetNumberReadsOk() (int32, bool)`

GetNumberReadsOk returns a tuple with the NumberReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberReads

`func (o *Ga4ghReadAlignment) HasNumberReads() bool`

HasNumberReads returns a boolean if a field has been set.

### SetNumberReads

`func (o *Ga4ghReadAlignment) SetNumberReads(v int32)`

SetNumberReads gets a reference to the given int32 and assigns it to the NumberReads field.

### GetFragmentLength

`func (o *Ga4ghReadAlignment) GetFragmentLength() int32`

GetFragmentLength returns the FragmentLength field if non-nil, zero value otherwise.

### GetFragmentLengthOk

`func (o *Ga4ghReadAlignment) GetFragmentLengthOk() (int32, bool)`

GetFragmentLengthOk returns a tuple with the FragmentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFragmentLength

`func (o *Ga4ghReadAlignment) HasFragmentLength() bool`

HasFragmentLength returns a boolean if a field has been set.

### SetFragmentLength

`func (o *Ga4ghReadAlignment) SetFragmentLength(v int32)`

SetFragmentLength gets a reference to the given int32 and assigns it to the FragmentLength field.

### GetReadNumber

`func (o *Ga4ghReadAlignment) GetReadNumber() int32`

GetReadNumber returns the ReadNumber field if non-nil, zero value otherwise.

### GetReadNumberOk

`func (o *Ga4ghReadAlignment) GetReadNumberOk() (int32, bool)`

GetReadNumberOk returns a tuple with the ReadNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadNumber

`func (o *Ga4ghReadAlignment) HasReadNumber() bool`

HasReadNumber returns a boolean if a field has been set.

### SetReadNumber

`func (o *Ga4ghReadAlignment) SetReadNumber(v int32)`

SetReadNumber gets a reference to the given int32 and assigns it to the ReadNumber field.

### GetFailedVendorQualityChecks

`func (o *Ga4ghReadAlignment) GetFailedVendorQualityChecks() bool`

GetFailedVendorQualityChecks returns the FailedVendorQualityChecks field if non-nil, zero value otherwise.

### GetFailedVendorQualityChecksOk

`func (o *Ga4ghReadAlignment) GetFailedVendorQualityChecksOk() (bool, bool)`

GetFailedVendorQualityChecksOk returns a tuple with the FailedVendorQualityChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailedVendorQualityChecks

`func (o *Ga4ghReadAlignment) HasFailedVendorQualityChecks() bool`

HasFailedVendorQualityChecks returns a boolean if a field has been set.

### SetFailedVendorQualityChecks

`func (o *Ga4ghReadAlignment) SetFailedVendorQualityChecks(v bool)`

SetFailedVendorQualityChecks gets a reference to the given bool and assigns it to the FailedVendorQualityChecks field.

### GetAlignment

`func (o *Ga4ghReadAlignment) GetAlignment() Ga4ghLinearAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *Ga4ghReadAlignment) GetAlignmentOk() (Ga4ghLinearAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignment

`func (o *Ga4ghReadAlignment) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### SetAlignment

`func (o *Ga4ghReadAlignment) SetAlignment(v Ga4ghLinearAlignment)`

SetAlignment gets a reference to the given Ga4ghLinearAlignment and assigns it to the Alignment field.

### GetSecondaryAlignment

`func (o *Ga4ghReadAlignment) GetSecondaryAlignment() bool`

GetSecondaryAlignment returns the SecondaryAlignment field if non-nil, zero value otherwise.

### GetSecondaryAlignmentOk

`func (o *Ga4ghReadAlignment) GetSecondaryAlignmentOk() (bool, bool)`

GetSecondaryAlignmentOk returns a tuple with the SecondaryAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecondaryAlignment

`func (o *Ga4ghReadAlignment) HasSecondaryAlignment() bool`

HasSecondaryAlignment returns a boolean if a field has been set.

### SetSecondaryAlignment

`func (o *Ga4ghReadAlignment) SetSecondaryAlignment(v bool)`

SetSecondaryAlignment gets a reference to the given bool and assigns it to the SecondaryAlignment field.

### GetSupplementaryAlignment

`func (o *Ga4ghReadAlignment) GetSupplementaryAlignment() bool`

GetSupplementaryAlignment returns the SupplementaryAlignment field if non-nil, zero value otherwise.

### GetSupplementaryAlignmentOk

`func (o *Ga4ghReadAlignment) GetSupplementaryAlignmentOk() (bool, bool)`

GetSupplementaryAlignmentOk returns a tuple with the SupplementaryAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupplementaryAlignment

`func (o *Ga4ghReadAlignment) HasSupplementaryAlignment() bool`

HasSupplementaryAlignment returns a boolean if a field has been set.

### SetSupplementaryAlignment

`func (o *Ga4ghReadAlignment) SetSupplementaryAlignment(v bool)`

SetSupplementaryAlignment gets a reference to the given bool and assigns it to the SupplementaryAlignment field.

### GetAlignedSequence

`func (o *Ga4ghReadAlignment) GetAlignedSequence() string`

GetAlignedSequence returns the AlignedSequence field if non-nil, zero value otherwise.

### GetAlignedSequenceOk

`func (o *Ga4ghReadAlignment) GetAlignedSequenceOk() (string, bool)`

GetAlignedSequenceOk returns a tuple with the AlignedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignedSequence

`func (o *Ga4ghReadAlignment) HasAlignedSequence() bool`

HasAlignedSequence returns a boolean if a field has been set.

### SetAlignedSequence

`func (o *Ga4ghReadAlignment) SetAlignedSequence(v string)`

SetAlignedSequence gets a reference to the given string and assigns it to the AlignedSequence field.

### GetAlignedQuality

`func (o *Ga4ghReadAlignment) GetAlignedQuality() []int32`

GetAlignedQuality returns the AlignedQuality field if non-nil, zero value otherwise.

### GetAlignedQualityOk

`func (o *Ga4ghReadAlignment) GetAlignedQualityOk() ([]int32, bool)`

GetAlignedQualityOk returns a tuple with the AlignedQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignedQuality

`func (o *Ga4ghReadAlignment) HasAlignedQuality() bool`

HasAlignedQuality returns a boolean if a field has been set.

### SetAlignedQuality

`func (o *Ga4ghReadAlignment) SetAlignedQuality(v []int32)`

SetAlignedQuality gets a reference to the given []int32 and assigns it to the AlignedQuality field.

### GetNextMatePosition

`func (o *Ga4ghReadAlignment) GetNextMatePosition() Ga4ghPosition`

GetNextMatePosition returns the NextMatePosition field if non-nil, zero value otherwise.

### GetNextMatePositionOk

`func (o *Ga4ghReadAlignment) GetNextMatePositionOk() (Ga4ghPosition, bool)`

GetNextMatePositionOk returns a tuple with the NextMatePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextMatePosition

`func (o *Ga4ghReadAlignment) HasNextMatePosition() bool`

HasNextMatePosition returns a boolean if a field has been set.

### SetNextMatePosition

`func (o *Ga4ghReadAlignment) SetNextMatePosition(v Ga4ghPosition)`

SetNextMatePosition gets a reference to the given Ga4ghPosition and assigns it to the NextMatePosition field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


