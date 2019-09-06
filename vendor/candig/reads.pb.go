// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candig/reads.proto

package candig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Describes the different types of CIGAR alignment operations that exist.
// Used wherever CIGAR alignments are used.
type CigarUnit_Operation int32

const (
	// The null operation.
	CigarUnit_OPERATION_UNSPECIFIED CigarUnit_Operation = 0
	// An alignment match indicates that a sequence can be aligned to the
	// reference without evidence of an INDEL. Unlike the
	// `SEQUENCE_MATCH` and `SEQUENCE_MISMATCH` operators,
	// the `ALIGNMENT_MATCH` operator does not indicate whether the
	// reference and read sequences are an exact match. This operator is
	// equivalent to SAM's `M`.
	CigarUnit_ALIGNMENT_MATCH CigarUnit_Operation = 1
	// The insert operator indicates that the read contains evidence of bases
	// being inserted into the reference. This operator is equivalent to SAM's
	// `I`.
	CigarUnit_INSERT CigarUnit_Operation = 2
	// The delete operator indicates that the read contains evidence of bases
	// being deleted from the reference. This operator is equivalent to SAM's
	// `D`.
	CigarUnit_DELETE CigarUnit_Operation = 3
	// The skip operator indicates that this read skips a long segment of the
	// reference, but the bases have not been deleted. This operator is commonly
	// used when working with RNA-seq data, where reads may skip long segments
	// of the reference between exons. This operator is equivalent to SAM's
	// `N`.
	CigarUnit_SKIP CigarUnit_Operation = 4
	// The soft clip operator indicates that bases at the start/end of a read
	// have not been considered during alignment. This may occur if the majority
	// of a read maps, except for low quality bases at the start/end of a read.
	// This operator is equivalent to SAM's `S`. Bases that are soft
	// clipped will still be stored in the read.
	CigarUnit_CLIP_SOFT CigarUnit_Operation = 5
	// The hard clip operator indicates that bases at the start/end of a read
	// have been omitted from this alignment. This may occur if this linear
	// alignment is part of a chimeric alignment, or if the read has been
	// trimmed (for example, during error correction or to trim poly-A tails for
	// RNA-seq). This operator is equivalent to SAM's `H`.
	CigarUnit_CLIP_HARD CigarUnit_Operation = 6
	// The pad operator indicates that there is padding in an alignment. This
	// operator is equivalent to SAM's `P`.
	CigarUnit_PAD CigarUnit_Operation = 7
	// This operator indicates that this portion of the aligned sequence exactly
	// matches the reference. This operator is equivalent to SAM's `=`.
	CigarUnit_SEQUENCE_MATCH CigarUnit_Operation = 8
	// This operator indicates that this portion of the aligned sequence is an
	// alignment match to the reference, but a sequence mismatch. This can
	// indicate a SNP or a read error. This operator is equivalent to SAM's
	// `X`.
	CigarUnit_SEQUENCE_MISMATCH CigarUnit_Operation = 9
)

var CigarUnit_Operation_name = map[int32]string{
	0: "OPERATION_UNSPECIFIED",
	1: "ALIGNMENT_MATCH",
	2: "INSERT",
	3: "DELETE",
	4: "SKIP",
	5: "CLIP_SOFT",
	6: "CLIP_HARD",
	7: "PAD",
	8: "SEQUENCE_MATCH",
	9: "SEQUENCE_MISMATCH",
}

var CigarUnit_Operation_value = map[string]int32{
	"OPERATION_UNSPECIFIED": 0,
	"ALIGNMENT_MATCH":       1,
	"INSERT":                2,
	"DELETE":                3,
	"SKIP":                  4,
	"CLIP_SOFT":             5,
	"CLIP_HARD":             6,
	"PAD":                   7,
	"SEQUENCE_MATCH":        8,
	"SEQUENCE_MISMATCH":     9,
}

func (x CigarUnit_Operation) String() string {
	return proto.EnumName(CigarUnit_Operation_name, int32(x))
}

func (CigarUnit_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{5, 0}
}

// ReadStats can be used to provide summary statistics about read data.
type ReadStats struct {
	// The number of aligned reads.
	AlignedReadCount int64 `protobuf:"varint,1,opt,name=aligned_read_count,json=alignedReadCount,proto3" json:"aligned_read_count,omitempty"`
	// The number of unaligned reads.
	UnalignedReadCount int64 `protobuf:"varint,2,opt,name=unaligned_read_count,json=unalignedReadCount,proto3" json:"unaligned_read_count,omitempty"`
	// The total number of bases.
	// This is equivalent to the sum of `alignedSequence.length` for all reads.
	BaseCount            int64    `protobuf:"varint,3,opt,name=base_count,json=baseCount,proto3" json:"base_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStats) Reset()         { *m = ReadStats{} }
func (m *ReadStats) String() string { return proto.CompactTextString(m) }
func (*ReadStats) ProtoMessage()    {}
func (*ReadStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{0}
}

func (m *ReadStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStats.Unmarshal(m, b)
}
func (m *ReadStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStats.Marshal(b, m, deterministic)
}
func (m *ReadStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStats.Merge(m, src)
}
func (m *ReadStats) XXX_Size() int {
	return xxx_messageInfo_ReadStats.Size(m)
}
func (m *ReadStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStats.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStats proto.InternalMessageInfo

func (m *ReadStats) GetAlignedReadCount() int64 {
	if m != nil {
		return m.AlignedReadCount
	}
	return 0
}

func (m *ReadStats) GetUnalignedReadCount() int64 {
	if m != nil {
		return m.UnalignedReadCount
	}
	return 0
}

func (m *ReadStats) GetBaseCount() int64 {
	if m != nil {
		return m.BaseCount
	}
	return 0
}

// A ReadGroup is a set of reads derived from one physical sequencing process.
type ReadGroup struct {
	// The read group ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the dataset this read group belongs to.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The read group name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The read group description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// A name for the sample this read group's data were generated from.
	// This field contains an arbitrary string, typically
	// corresponding to the SM tag in a BAM file.
	SampleName string `protobuf:"bytes,5,opt,name=sample_name,json=sampleName,proto3" json:"sample_name,omitempty"`
	// The Biosample this read group's data was generated from.
	BiosampleId string `protobuf:"bytes,6,opt,name=biosample_id,json=biosampleId,proto3" json:"biosample_id,omitempty"`
	// The experiment used to generate this read group.
	Experiment *Experiment `protobuf:"bytes,7,opt,name=experiment,proto3" json:"experiment,omitempty"`
	// The predicted insert size of this read group.
	PredictedInsertSize int32 `protobuf:"varint,8,opt,name=predicted_insert_size,json=predictedInsertSize,proto3" json:"predicted_insert_size,omitempty"`
	// The time at which this read group was created in milliseconds from the
	// epoch.
	Created int64 `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	// The time at which this read group was last updated in milliseconds
	// from the epoch.
	Updated int64 `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
	// Statistical data on reads in this read group.
	Stats *ReadStats `protobuf:"bytes,11,opt,name=stats,proto3" json:"stats,omitempty"`
	// Program can be used to track the provenance of how read data was generated.
	Programs []*Program `protobuf:"bytes,12,rep,name=programs,proto3" json:"programs,omitempty"`
	// The ID of the reference set to which the reads in this read group are
	// aligned. Required if there are any read alignments.
	ReferenceSetId string `protobuf:"bytes,13,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// A map of additional information about the Read Group.
	Attributes           *Attributes `protobuf:"bytes,15,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadGroup) Reset()         { *m = ReadGroup{} }
func (m *ReadGroup) String() string { return proto.CompactTextString(m) }
func (*ReadGroup) ProtoMessage()    {}
func (*ReadGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{1}
}

func (m *ReadGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroup.Unmarshal(m, b)
}
func (m *ReadGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroup.Marshal(b, m, deterministic)
}
func (m *ReadGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroup.Merge(m, src)
}
func (m *ReadGroup) XXX_Size() int {
	return xxx_messageInfo_ReadGroup.Size(m)
}
func (m *ReadGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroup proto.InternalMessageInfo

func (m *ReadGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroup) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ReadGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReadGroup) GetSampleName() string {
	if m != nil {
		return m.SampleName
	}
	return ""
}

func (m *ReadGroup) GetBiosampleId() string {
	if m != nil {
		return m.BiosampleId
	}
	return ""
}

func (m *ReadGroup) GetExperiment() *Experiment {
	if m != nil {
		return m.Experiment
	}
	return nil
}

func (m *ReadGroup) GetPredictedInsertSize() int32 {
	if m != nil {
		return m.PredictedInsertSize
	}
	return 0
}

func (m *ReadGroup) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ReadGroup) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ReadGroup) GetStats() *ReadStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *ReadGroup) GetPrograms() []*Program {
	if m != nil {
		return m.Programs
	}
	return nil
}

func (m *ReadGroup) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *ReadGroup) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A ReadGroupSet is a logical collection of ReadGroups. Typically one
// ReadGroupSet represents all the reads from one experimental sample.
type ReadGroupSet struct {
	// The read group set ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the dataset this read group set belongs to.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The read group set name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Statistical data on reads in this read group set.
	Stats *ReadStats `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
	// The read groups in this set.
	ReadGroups []*ReadGroup `protobuf:"bytes,5,rep,name=read_groups,json=readGroups,proto3" json:"read_groups,omitempty"`
	// For read group data identification and search join purposes
	PatientId string `protobuf:"bytes,6,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	// For read group data identification and search join purposes
	SampleId string `protobuf:"bytes,7,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	// A map of additional information about the Read Group Set.
	Attributes           *Attributes `protobuf:"bytes,8,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadGroupSet) Reset()         { *m = ReadGroupSet{} }
func (m *ReadGroupSet) String() string { return proto.CompactTextString(m) }
func (*ReadGroupSet) ProtoMessage()    {}
func (*ReadGroupSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{2}
}

func (m *ReadGroupSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroupSet.Unmarshal(m, b)
}
func (m *ReadGroupSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroupSet.Marshal(b, m, deterministic)
}
func (m *ReadGroupSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroupSet.Merge(m, src)
}
func (m *ReadGroupSet) XXX_Size() int {
	return xxx_messageInfo_ReadGroupSet.Size(m)
}
func (m *ReadGroupSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroupSet.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroupSet proto.InternalMessageInfo

func (m *ReadGroupSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroupSet) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ReadGroupSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroupSet) GetStats() *ReadStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *ReadGroupSet) GetReadGroups() []*ReadGroup {
	if m != nil {
		return m.ReadGroups
	}
	return nil
}

func (m *ReadGroupSet) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func (m *ReadGroupSet) GetSampleId() string {
	if m != nil {
		return m.SampleId
	}
	return ""
}

func (m *ReadGroupSet) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A linear alignment describes the alignment of a read to a Reference, using a
// position and CIGAR array.
type LinearAlignment struct {
	// The position of this alignment.
	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// The mapping quality of this alignment, meaning the likelihood that the read
	// maps to this position.
	//
	// Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to
	// the nearest integer.
	MappingQuality int32 `protobuf:"varint,2,opt,name=mapping_quality,json=mappingQuality,proto3" json:"mapping_quality,omitempty"`
	// Represents the local alignment of this sequence (alignment matches, indels,
	// etc)
	// versus the reference.
	Cigar                []*CigarUnit `protobuf:"bytes,3,rep,name=cigar,proto3" json:"cigar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LinearAlignment) Reset()         { *m = LinearAlignment{} }
func (m *LinearAlignment) String() string { return proto.CompactTextString(m) }
func (*LinearAlignment) ProtoMessage()    {}
func (*LinearAlignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{3}
}

func (m *LinearAlignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinearAlignment.Unmarshal(m, b)
}
func (m *LinearAlignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinearAlignment.Marshal(b, m, deterministic)
}
func (m *LinearAlignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinearAlignment.Merge(m, src)
}
func (m *LinearAlignment) XXX_Size() int {
	return xxx_messageInfo_LinearAlignment.Size(m)
}
func (m *LinearAlignment) XXX_DiscardUnknown() {
	xxx_messageInfo_LinearAlignment.DiscardUnknown(m)
}

var xxx_messageInfo_LinearAlignment proto.InternalMessageInfo

func (m *LinearAlignment) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *LinearAlignment) GetMappingQuality() int32 {
	if m != nil {
		return m.MappingQuality
	}
	return 0
}

func (m *LinearAlignment) GetCigar() []*CigarUnit {
	if m != nil {
		return m.Cigar
	}
	return nil
}

// Each read alignment describes an alignment with additional information
// about the fragment and the read. A read alignment object is equivalent to a
// line in a SAM file.
type ReadAlignment struct {
	// The read alignment ID. This ID is unique within the read group this
	// alignment belongs to.
	//
	// For performance reasons, this field may be omitted by a backend.
	// If provided, its intended use is to make caching and UI display easier for
	// genome browsers and other lightweight clients.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the read group this read belongs to.
	// (Every read must belong to exactly one read group.)
	ReadGroupId string `protobuf:"bytes,2,opt,name=read_group_id,json=readGroupId,proto3" json:"read_group_id,omitempty"`
	// The fragment name. Equivalent to QNAME (query template name) in SAM.
	FragmentName string `protobuf:"bytes,3,opt,name=fragment_name,json=fragmentName,proto3" json:"fragment_name,omitempty"`
	// The orientation and the distance between reads from the fragment are
	// inconsistent with the sequencing protocol (inverse of SAM flag 0x2).
	ImproperPlacement bool `protobuf:"varint,4,opt,name=improper_placement,json=improperPlacement,proto3" json:"improper_placement,omitempty"`
	// The fragment is a PCR or optical duplicate (SAM flag 0x400).
	DuplicateFragment bool `protobuf:"varint,5,opt,name=duplicate_fragment,json=duplicateFragment,proto3" json:"duplicate_fragment,omitempty"`
	// The number of reads in the fragment (extension to SAM flag 0x1).
	NumberReads int32 `protobuf:"varint,6,opt,name=number_reads,json=numberReads,proto3" json:"number_reads,omitempty"`
	// The observed length of the fragment, equivalent to TLEN in SAM.
	FragmentLength int32 `protobuf:"varint,7,opt,name=fragment_length,json=fragmentLength,proto3" json:"fragment_length,omitempty"`
	// The read ordinal in the fragment, 0-based and less than numberReads. This
	// field replaces SAM flag 0x40 and 0x80 and is intended to more cleanly
	// represent multiple reads per fragment.
	ReadNumber int32 `protobuf:"varint,8,opt,name=read_number,json=readNumber,proto3" json:"read_number,omitempty"`
	// The read fails platform or vendor quality checks (SAM flag 0x200).
	FailedVendorQualityChecks bool `protobuf:"varint,9,opt,name=failed_vendor_quality_checks,json=failedVendorQualityChecks,proto3" json:"failed_vendor_quality_checks,omitempty"`
	// The alignment for this alignment message. This field will be null if the
	// read is unmapped.
	Alignment *LinearAlignment `protobuf:"bytes,10,opt,name=alignment,proto3" json:"alignment,omitempty"`
	// Whether this alignment is secondary. Equivalent to SAM flag 0x100.
	// A secondary alignment represents an alternative to the primary alignment
	// for this read. Aligners may return secondary alignments if a read can map
	// ambiguously to multiple coordinates in the genome.
	//
	// By convention, each read has one and only one alignment where both
	// secondaryAlignment and supplementaryAlignment are false.
	SecondaryAlignment bool `protobuf:"varint,11,opt,name=secondary_alignment,json=secondaryAlignment,proto3" json:"secondary_alignment,omitempty"`
	// Whether this alignment is supplementary. Equivalent to SAM flag 0x800.
	// Supplementary alignments are used in the representation of a chimeric
	// alignment. In a chimeric alignment, a read is split into multiple
	// linear alignments that map to different reference contigs. The first
	// linear alignment in the read will be designated as the representative
	// alignment; the remaining linear alignments will be designated as
	// supplementary alignments. These alignments may have different mapping
	// quality scores.
	//
	// In each linear alignment in a chimeric alignment, the read will be hard
	// clipped. The `alignedSequence` and `alignedQuality` fields in the alignment
	// message will only represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `protobuf:"varint,12,opt,name=supplementary_alignment,json=supplementaryAlignment,proto3" json:"supplementary_alignment,omitempty"`
	// The bases of the read sequence contained in this alignment record
	// (equivalent to SEQ in SAM).
	//
	// `alignedSequence` and `alignedQuality` may be shorter than the full read
	// sequence and quality. This will occur if the alignment is part of a
	// chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will indicate
	// the length of the excised sequence.
	AlignedSequence string `protobuf:"bytes,13,opt,name=aligned_sequence,json=alignedSequence,proto3" json:"aligned_sequence,omitempty"`
	// The quality of the read sequence contained in this alignment message
	// (equivalent to QUAL in SAM).
	//
	// `alignedSequence` and `alignedQuality` may be shorter than the full read
	// sequence and quality. This will occur if the alignment is part of a
	// chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will indicate
	// the length of the excised sequence.
	AlignedQuality []int32 `protobuf:"varint,14,rep,packed,name=aligned_quality,json=alignedQuality,proto3" json:"aligned_quality,omitempty"`
	// The mapping of the primary alignment of the `(readNumber+1)%numberReads`
	// read in the fragment. It replaces mate position and mate strand in SAM.
	NextMatePosition *Position `protobuf:"bytes,15,opt,name=next_mate_position,json=nextMatePosition,proto3" json:"next_mate_position,omitempty"`
	// A map of additional information about the Alignment.
	Attributes           *Attributes `protobuf:"bytes,17,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadAlignment) Reset()         { *m = ReadAlignment{} }
func (m *ReadAlignment) String() string { return proto.CompactTextString(m) }
func (*ReadAlignment) ProtoMessage()    {}
func (*ReadAlignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{4}
}

func (m *ReadAlignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAlignment.Unmarshal(m, b)
}
func (m *ReadAlignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAlignment.Marshal(b, m, deterministic)
}
func (m *ReadAlignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAlignment.Merge(m, src)
}
func (m *ReadAlignment) XXX_Size() int {
	return xxx_messageInfo_ReadAlignment.Size(m)
}
func (m *ReadAlignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAlignment.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAlignment proto.InternalMessageInfo

func (m *ReadAlignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadAlignment) GetReadGroupId() string {
	if m != nil {
		return m.ReadGroupId
	}
	return ""
}

func (m *ReadAlignment) GetFragmentName() string {
	if m != nil {
		return m.FragmentName
	}
	return ""
}

func (m *ReadAlignment) GetImproperPlacement() bool {
	if m != nil {
		return m.ImproperPlacement
	}
	return false
}

func (m *ReadAlignment) GetDuplicateFragment() bool {
	if m != nil {
		return m.DuplicateFragment
	}
	return false
}

func (m *ReadAlignment) GetNumberReads() int32 {
	if m != nil {
		return m.NumberReads
	}
	return 0
}

func (m *ReadAlignment) GetFragmentLength() int32 {
	if m != nil {
		return m.FragmentLength
	}
	return 0
}

func (m *ReadAlignment) GetReadNumber() int32 {
	if m != nil {
		return m.ReadNumber
	}
	return 0
}

func (m *ReadAlignment) GetFailedVendorQualityChecks() bool {
	if m != nil {
		return m.FailedVendorQualityChecks
	}
	return false
}

func (m *ReadAlignment) GetAlignment() *LinearAlignment {
	if m != nil {
		return m.Alignment
	}
	return nil
}

func (m *ReadAlignment) GetSecondaryAlignment() bool {
	if m != nil {
		return m.SecondaryAlignment
	}
	return false
}

func (m *ReadAlignment) GetSupplementaryAlignment() bool {
	if m != nil {
		return m.SupplementaryAlignment
	}
	return false
}

func (m *ReadAlignment) GetAlignedSequence() string {
	if m != nil {
		return m.AlignedSequence
	}
	return ""
}

func (m *ReadAlignment) GetAlignedQuality() []int32 {
	if m != nil {
		return m.AlignedQuality
	}
	return nil
}

func (m *ReadAlignment) GetNextMatePosition() *Position {
	if m != nil {
		return m.NextMatePosition
	}
	return nil
}

func (m *ReadAlignment) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A single CIGAR operation.
type CigarUnit struct {
	Operation CigarUnit_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=candig.CigarUnit_Operation" json:"operation,omitempty"`
	// The number of genomic bases that the operation runs for. Required.
	OperationLength int64 `protobuf:"varint,2,opt,name=operation_length,json=operationLength,proto3" json:"operation_length,omitempty"`
	// `referenceSequence` is only used at mismatches
	// (`SEQUENCE_MISMATCH`) and deletions (`DELETE`).
	// Filling this field replaces SAM's MD tag. If the relevant information is
	// not available, this field is unset.
	ReferenceSequence    string   `protobuf:"bytes,3,opt,name=reference_sequence,json=referenceSequence,proto3" json:"reference_sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CigarUnit) Reset()         { *m = CigarUnit{} }
func (m *CigarUnit) String() string { return proto.CompactTextString(m) }
func (*CigarUnit) ProtoMessage()    {}
func (*CigarUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63ed78fb4692717, []int{5}
}

func (m *CigarUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CigarUnit.Unmarshal(m, b)
}
func (m *CigarUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CigarUnit.Marshal(b, m, deterministic)
}
func (m *CigarUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CigarUnit.Merge(m, src)
}
func (m *CigarUnit) XXX_Size() int {
	return xxx_messageInfo_CigarUnit.Size(m)
}
func (m *CigarUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_CigarUnit.DiscardUnknown(m)
}

var xxx_messageInfo_CigarUnit proto.InternalMessageInfo

func (m *CigarUnit) GetOperation() CigarUnit_Operation {
	if m != nil {
		return m.Operation
	}
	return CigarUnit_OPERATION_UNSPECIFIED
}

func (m *CigarUnit) GetOperationLength() int64 {
	if m != nil {
		return m.OperationLength
	}
	return 0
}

func (m *CigarUnit) GetReferenceSequence() string {
	if m != nil {
		return m.ReferenceSequence
	}
	return ""
}

func init() {
	proto.RegisterEnum("candig.CigarUnit_Operation", CigarUnit_Operation_name, CigarUnit_Operation_value)
	proto.RegisterType((*ReadStats)(nil), "candig.ReadStats")
	proto.RegisterType((*ReadGroup)(nil), "candig.ReadGroup")
	proto.RegisterType((*ReadGroupSet)(nil), "candig.ReadGroupSet")
	proto.RegisterType((*LinearAlignment)(nil), "candig.LinearAlignment")
	proto.RegisterType((*ReadAlignment)(nil), "candig.ReadAlignment")
	proto.RegisterType((*CigarUnit)(nil), "candig.CigarUnit")
}

func init() { proto.RegisterFile("candig/reads.proto", fileDescriptor_f63ed78fb4692717) }

var fileDescriptor_f63ed78fb4692717 = []byte{
	// 1029 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x73, 0xdb, 0xc4,
	0x17, 0xfd, 0xf9, 0x5f, 0x62, 0x5d, 0x27, 0xb6, 0xbc, 0xf9, 0x95, 0xaa, 0x14, 0x06, 0x63, 0x1e,
	0x1a, 0x86, 0x36, 0x65, 0xcc, 0x30, 0x0c, 0x2f, 0x30, 0x1e, 0x47, 0x69, 0x35, 0x24, 0x8e, 0xbb,
	0x72, 0x78, 0xd5, 0x6c, 0xb4, 0x1b, 0x77, 0x07, 0x5b, 0x52, 0x57, 0x6b, 0xa6, 0xed, 0x33, 0x4f,
	0xbc, 0xf1, 0x05, 0xf8, 0x0a, 0x7c, 0x17, 0xde, 0xf9, 0x2e, 0xcc, 0x5e, 0x49, 0x2b, 0x27, 0xc0,
	0xd0, 0x07, 0xde, 0xa4, 0x73, 0xce, 0xde, 0x5d, 0x9f, 0x7b, 0xee, 0xca, 0x40, 0x62, 0x96, 0x70,
	0xb9, 0x7a, 0xaa, 0x04, 0xe3, 0xf9, 0x49, 0xa6, 0x52, 0x9d, 0x92, 0xbd, 0x02, 0x7b, 0xff, 0xa8,
	0xe4, 0xe2, 0x74, 0xb3, 0x49, 0x93, 0x82, 0x1c, 0xff, 0xdc, 0x00, 0x87, 0x0a, 0xc6, 0x43, 0xcd,
	0x74, 0x4e, 0x1e, 0x03, 0x61, 0x6b, 0xb9, 0x4a, 0x04, 0x8f, 0x4c, 0x85, 0x28, 0x4e, 0xb7, 0x89,
	0xf6, 0x1a, 0xa3, 0xc6, 0x71, 0x8b, 0xba, 0x25, 0x63, 0xd4, 0x33, 0x83, 0x93, 0xcf, 0xe1, 0xff,
	0xdb, 0xe4, 0x6f, 0xf4, 0x4d, 0xd4, 0x13, 0xcb, 0xd5, 0x2b, 0x3e, 0x04, 0xb8, 0x66, 0xb9, 0x28,
	0x75, 0x2d, 0xd4, 0x39, 0x06, 0x41, 0x7a, 0xfc, 0x53, 0xbb, 0x38, 0xcc, 0x33, 0x95, 0x6e, 0x33,
	0xd2, 0x87, 0xa6, 0xe4, 0xb8, 0xb9, 0x43, 0x9b, 0x92, 0x9b, 0xc5, 0x9c, 0x69, 0x96, 0x0b, 0x1d,
	0x49, 0x8e, 0x9b, 0x38, 0xd4, 0x29, 0x91, 0x80, 0x13, 0x02, 0xed, 0x84, 0x6d, 0x04, 0x56, 0x75,
	0x28, 0x3e, 0x93, 0x11, 0xf4, 0xb8, 0xc8, 0x63, 0x25, 0x33, 0x2d, 0xd3, 0xc4, 0x6b, 0x23, 0xb5,
	0x0b, 0x91, 0x8f, 0xa0, 0x97, 0xb3, 0x4d, 0xb6, 0x16, 0x11, 0x2e, 0xee, 0xa0, 0x02, 0x0a, 0x68,
	0x6e, 0x4a, 0x7c, 0x0c, 0x07, 0xd7, 0x32, 0x2d, 0x35, 0x92, 0x7b, 0x7b, 0x45, 0x0d, 0x8b, 0x05,
	0x9c, 0x4c, 0x00, 0xc4, 0xeb, 0x4c, 0x28, 0xb9, 0x11, 0x89, 0xf6, 0xf6, 0x47, 0x8d, 0xe3, 0xde,
	0x84, 0x9c, 0x14, 0x6e, 0x9f, 0xf8, 0x96, 0xa1, 0x3b, 0x2a, 0x32, 0x81, 0x7b, 0x99, 0x12, 0x5c,
	0xc6, 0x5a, 0xf0, 0x48, 0x26, 0xb9, 0x50, 0x3a, 0xca, 0xe5, 0x5b, 0xe1, 0x75, 0x47, 0x8d, 0xe3,
	0x0e, 0x3d, 0xb2, 0x64, 0x80, 0x5c, 0x28, 0xdf, 0x0a, 0xe2, 0xc1, 0x7e, 0xac, 0x04, 0xd3, 0x82,
	0x7b, 0x0e, 0x5a, 0x57, 0xbd, 0x1a, 0x66, 0x9b, 0x71, 0x64, 0xa0, 0x60, 0xca, 0x57, 0xf2, 0x08,
	0x3a, 0xb9, 0x69, 0xad, 0xd7, 0xc3, 0x63, 0x0d, 0xab, 0x63, 0xd9, 0x9e, 0xd3, 0x82, 0x27, 0x9f,
	0x41, 0x37, 0x53, 0xe9, 0x4a, 0xb1, 0x4d, 0xee, 0x1d, 0x8c, 0x5a, 0xc7, 0xbd, 0xc9, 0xa0, 0xd2,
	0x2e, 0x0a, 0x9c, 0x5a, 0x01, 0x39, 0x06, 0x57, 0x89, 0x1b, 0xa1, 0x44, 0x12, 0x8b, 0xa8, 0x6c,
	0xc8, 0x21, 0x1a, 0xd3, 0xb7, 0x78, 0x88, 0x5d, 0x99, 0x00, 0x30, 0xad, 0x95, 0xbc, 0xde, 0x6a,
	0x91, 0x7b, 0x83, 0xdb, 0xde, 0x4c, 0x2d, 0x43, 0x77, 0x54, 0xe3, 0x5f, 0x9b, 0x70, 0x60, 0x63,
	0x10, 0x0a, 0xfd, 0x5f, 0x24, 0xc1, 0xfa, 0xd0, 0xfe, 0x17, 0x1f, 0x26, 0xd0, 0xc3, 0x28, 0xaf,
	0xcc, 0xe6, 0xb9, 0xd7, 0x41, 0x2b, 0x6e, 0xc9, 0xf1, 0x58, 0x14, 0x54, 0xf5, 0x98, 0x9b, 0xf3,
	0x64, 0x4c, 0x4b, 0x91, 0xe8, 0x3a, 0x21, 0x4e, 0x89, 0x04, 0x9c, 0x3c, 0x04, 0xa7, 0xce, 0xcf,
	0x3e, 0xb2, 0xdd, 0xdd, 0xf0, 0xec, 0x18, 0xd4, 0x7d, 0x27, 0x83, 0x7e, 0x69, 0xc0, 0xe0, 0x5c,
	0x26, 0x82, 0xa9, 0xa9, 0x99, 0x30, 0x0c, 0xd4, 0x63, 0xe8, 0x66, 0x69, 0x2e, 0x31, 0xe7, 0x0d,
	0xac, 0xe2, 0xda, 0xfe, 0x95, 0x38, 0xb5, 0x0a, 0xf2, 0x08, 0x06, 0x1b, 0x96, 0x65, 0x32, 0x59,
	0x45, 0xaf, 0xb6, 0x6c, 0x2d, 0xf5, 0x1b, 0xb4, 0xb1, 0x43, 0xfb, 0x25, 0xfc, 0xa2, 0x40, 0x8d,
	0x6f, 0xb1, 0x5c, 0x31, 0xe5, 0xb5, 0x6e, 0x1b, 0x31, 0x33, 0xe0, 0x55, 0x22, 0x35, 0x2d, 0xf8,
	0xf1, 0x1f, 0x1d, 0x38, 0x34, 0xee, 0xd4, 0x27, 0xba, 0xdb, 0xb5, 0x31, 0x1c, 0xd6, 0xce, 0xd6,
	0x8d, 0xeb, 0x59, 0x23, 0x03, 0x4e, 0x3e, 0x81, 0xc3, 0x1b, 0xc5, 0x56, 0x66, 0x7d, 0xb4, 0xd3,
	0xc3, 0x83, 0x0a, 0xc4, 0x91, 0x7c, 0x02, 0x44, 0x6e, 0x32, 0x95, 0x66, 0x42, 0x45, 0xd9, 0x9a,
	0xc5, 0x02, 0xe7, 0xce, 0x34, 0xb6, 0x4b, 0x87, 0x15, 0xb3, 0xa8, 0x08, 0x23, 0xe7, 0xdb, 0x6c,
	0x2d, 0x63, 0xa6, 0x45, 0x54, 0x15, 0xc2, 0x49, 0xef, 0xd2, 0xa1, 0x65, 0xce, 0x4a, 0xc2, 0x0c,
	0x7c, 0xb2, 0xdd, 0x5c, 0x0b, 0x85, 0x57, 0x5a, 0x8e, 0xed, 0xec, 0xd0, 0x5e, 0x81, 0x99, 0x5f,
	0x98, 0x1b, 0xf7, 0xec, 0x29, 0xd7, 0x22, 0x59, 0xe9, 0x97, 0xd8, 0xd6, 0x0e, 0xed, 0x57, 0xf0,
	0x39, 0xa2, 0xe6, 0x76, 0xc1, 0x9f, 0x5c, 0x2c, 0x2e, 0x67, 0x1b, 0x93, 0x33, 0x47, 0x84, 0x7c,
	0x0b, 0x1f, 0xdc, 0x30, 0xb9, 0x16, 0x3c, 0xfa, 0x51, 0x24, 0x3c, 0x55, 0x55, 0x37, 0xa2, 0xf8,
	0xa5, 0x88, 0x7f, 0xc8, 0x71, 0xce, 0xbb, 0xf4, 0x41, 0xa1, 0xf9, 0x1e, 0x25, 0x65, 0x67, 0x66,
	0x28, 0x20, 0x5f, 0x82, 0xc3, 0x2a, 0xc7, 0x71, 0xf6, 0x7b, 0x93, 0xfb, 0x55, 0x8f, 0xee, 0x44,
	0x84, 0xd6, 0x4a, 0xf2, 0x14, 0x8e, 0x72, 0x11, 0xa7, 0x09, 0x67, 0xea, 0x4d, 0x54, 0x17, 0xe8,
	0xe1, 0x76, 0xc4, 0x52, 0x75, 0x33, 0xbf, 0x82, 0xfb, 0xf9, 0x36, 0xcb, 0xd6, 0x68, 0xe9, 0xed,
	0x45, 0x07, 0xb8, 0xe8, 0xbd, 0x5b, 0x74, 0xbd, 0xf0, 0x53, 0xa8, 0x3e, 0x1c, 0x51, 0x2e, 0x5e,
	0x6d, 0xcd, 0xcd, 0x50, 0x5e, 0x15, 0x83, 0x12, 0x0f, 0x4b, 0xd8, 0xd8, 0x5a, 0x49, 0xab, 0x50,
	0xf6, 0x47, 0x2d, 0x63, 0x6b, 0x09, 0x57, 0xa1, 0xfc, 0x06, 0x48, 0x22, 0x5e, 0xeb, 0x68, 0x63,
	0x3a, 0x6a, 0x53, 0x3f, 0xf8, 0x87, 0xd4, 0xbb, 0x46, 0x7b, 0xc1, 0xb4, 0xa8, 0x90, 0x3b, 0x33,
	0x37, 0x7c, 0xa7, 0x99, 0xfb, 0xbd, 0x09, 0x8e, 0x0d, 0x3d, 0xf9, 0x1a, 0x1c, 0x13, 0x32, 0x66,
	0xc7, 0xad, 0x3f, 0x79, 0xf8, 0x97, 0xd1, 0x38, 0xb9, 0xac, 0x24, 0xb4, 0x56, 0x1b, 0x43, 0xec,
	0x4b, 0x95, 0x9e, 0xe2, 0x8b, 0x39, 0xb0, 0x78, 0x19, 0x9f, 0x27, 0x40, 0x76, 0xaf, 0xd9, 0xd2,
	0xbd, 0x62, 0x24, 0x86, 0x3b, 0x17, 0x6d, 0x41, 0x8c, 0x7f, 0x6b, 0x80, 0x63, 0xb7, 0x24, 0x0f,
	0xe0, 0xde, 0xe5, 0xc2, 0xa7, 0xd3, 0x65, 0x70, 0x39, 0x8f, 0xae, 0xe6, 0xe1, 0xc2, 0x9f, 0x05,
	0x67, 0x81, 0x7f, 0xea, 0xfe, 0x8f, 0x1c, 0xc1, 0x60, 0x7a, 0x1e, 0x3c, 0x9b, 0x5f, 0xf8, 0xf3,
	0x65, 0x74, 0x31, 0x5d, 0xce, 0x9e, 0xbb, 0x0d, 0x02, 0xb0, 0x17, 0xcc, 0x43, 0x9f, 0x2e, 0xdd,
	0xa6, 0x79, 0x3e, 0xf5, 0xcf, 0xfd, 0xa5, 0xef, 0xb6, 0x48, 0x17, 0xda, 0xe1, 0x77, 0xc1, 0xc2,
	0x6d, 0x93, 0x43, 0x70, 0x66, 0xe7, 0xc1, 0x22, 0x0a, 0x2f, 0xcf, 0x96, 0x6e, 0xc7, 0xbe, 0x3e,
	0x9f, 0xd2, 0x53, 0x77, 0x8f, 0xec, 0x43, 0x6b, 0x31, 0x3d, 0x75, 0xf7, 0x09, 0x81, 0x7e, 0xe8,
	0xbf, 0xb8, 0xf2, 0xe7, 0x33, 0xbf, 0x2c, 0xde, 0x25, 0xf7, 0x60, 0x58, 0x63, 0x41, 0x58, 0xc0,
	0xce, 0xf5, 0x1e, 0xfe, 0x09, 0xf9, 0xe2, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x26, 0x15,
	0xa1, 0xb7, 0x08, 0x00, 0x00,
}