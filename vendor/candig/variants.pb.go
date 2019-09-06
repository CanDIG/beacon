// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candig/variants.proto

package candig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// This metadata represents VCF header information.
type VariantSetMetadata struct {
	// The top-level key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value field for simple metadata.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// User-provided ID field, not enforced by this API.
	// Two or more pieces of structured metadata with identical
	// id and key fields are considered equivalent.
	// FIXME: If it's not enforced, then why can't it be null?
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The type of data.
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// The number of values that can be included in a field described by this
	// metadata.
	Number string `protobuf:"bytes,5,opt,name=number,proto3" json:"number,omitempty"`
	// A textual description of this metadata.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// A map of additional information about the metadata record.
	Attributes           *Attributes `protobuf:"bytes,8,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VariantSetMetadata) Reset()         { *m = VariantSetMetadata{} }
func (m *VariantSetMetadata) String() string { return proto.CompactTextString(m) }
func (*VariantSetMetadata) ProtoMessage()    {}
func (*VariantSetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c8f728f42a356bd, []int{0}
}

func (m *VariantSetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariantSetMetadata.Unmarshal(m, b)
}
func (m *VariantSetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariantSetMetadata.Marshal(b, m, deterministic)
}
func (m *VariantSetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariantSetMetadata.Merge(m, src)
}
func (m *VariantSetMetadata) XXX_Size() int {
	return xxx_messageInfo_VariantSetMetadata.Size(m)
}
func (m *VariantSetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_VariantSetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_VariantSetMetadata proto.InternalMessageInfo

func (m *VariantSetMetadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *VariantSetMetadata) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *VariantSetMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VariantSetMetadata) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VariantSetMetadata) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *VariantSetMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *VariantSetMetadata) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A VariantSet is a collection of variants and variant calls intended to be
// analyzed together.
type VariantSet struct {
	// The variant set ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The variant set name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The ID of the dataset this variant set belongs to.
	DatasetId string `protobuf:"bytes,3,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The ID of the reference set that describes the sequences used by the
	// variants in this set.
	ReferenceSetId string `protobuf:"bytes,4,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// For variant data identification and search join purposes
	PatientId string `protobuf:"bytes,5,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	// For variant data identification and search join purposes
	SampleId string `protobuf:"bytes,6,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	// Optional metadata associated with this variant set.
	// This array can be used to store information about the variant set, such as
	// information found in VCF header fields, that isn't already available in
	// first class fields such as "name".
	Metadata             []*VariantSetMetadata `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *VariantSet) Reset()         { *m = VariantSet{} }
func (m *VariantSet) String() string { return proto.CompactTextString(m) }
func (*VariantSet) ProtoMessage()    {}
func (*VariantSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c8f728f42a356bd, []int{1}
}

func (m *VariantSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariantSet.Unmarshal(m, b)
}
func (m *VariantSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariantSet.Marshal(b, m, deterministic)
}
func (m *VariantSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariantSet.Merge(m, src)
}
func (m *VariantSet) XXX_Size() int {
	return xxx_messageInfo_VariantSet.Size(m)
}
func (m *VariantSet) XXX_DiscardUnknown() {
	xxx_messageInfo_VariantSet.DiscardUnknown(m)
}

var xxx_messageInfo_VariantSet proto.InternalMessageInfo

func (m *VariantSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VariantSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VariantSet) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *VariantSet) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *VariantSet) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func (m *VariantSet) GetSampleId() string {
	if m != nil {
		return m.SampleId
	}
	return ""
}

func (m *VariantSet) GetMetadata() []*VariantSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// A CallSet is a collection of calls that were generated by the same analysis
// of the same sample.
type CallSet struct {
	// The call set ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The call set name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The Biosample the call set data was generated from.
	BiosampleId string `protobuf:"bytes,3,opt,name=biosample_id,json=biosampleId,proto3" json:"biosample_id,omitempty"`
	// The IDs of the variant sets this call set has calls in.
	VariantSetIds []string `protobuf:"bytes,4,rep,name=variant_set_ids,json=variantSetIds,proto3" json:"variant_set_ids,omitempty"`
	// The date this call set was created in milliseconds from the epoch.
	Created int64 `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	// The time at which this call set was last updated in
	// milliseconds from the epoch.
	Updated int64 `protobuf:"varint,6,opt,name=updated,proto3" json:"updated,omitempty"`
	// A map of additional information about the Call Set.
	Attributes           *Attributes `protobuf:"bytes,8,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CallSet) Reset()         { *m = CallSet{} }
func (m *CallSet) String() string { return proto.CompactTextString(m) }
func (*CallSet) ProtoMessage()    {}
func (*CallSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c8f728f42a356bd, []int{2}
}

func (m *CallSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallSet.Unmarshal(m, b)
}
func (m *CallSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallSet.Marshal(b, m, deterministic)
}
func (m *CallSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallSet.Merge(m, src)
}
func (m *CallSet) XXX_Size() int {
	return xxx_messageInfo_CallSet.Size(m)
}
func (m *CallSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CallSet.DiscardUnknown(m)
}

var xxx_messageInfo_CallSet proto.InternalMessageInfo

func (m *CallSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CallSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CallSet) GetBiosampleId() string {
	if m != nil {
		return m.BiosampleId
	}
	return ""
}

func (m *CallSet) GetVariantSetIds() []string {
	if m != nil {
		return m.VariantSetIds
	}
	return nil
}

func (m *CallSet) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *CallSet) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *CallSet) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A `Call` represents the determination of genotype with respect to a
// particular `Variant`.
//
// It may include associated information such as quality
// and phasing. For example, a call might assign a probability of 0.32 to
// the occurrence of a SNP named rs1234 in a call set with the name NA12345.
type Call struct {
	// The name of the call set this variant call belongs to.
	// If this field is not present, the ordering of the call sets from a
	// `SearchCallSetsRequest` over this `VariantSet` is guaranteed to match
	// the ordering of the calls on this `Variant`.
	// The number of results will also be the same.
	CallSetName string `protobuf:"bytes,1,opt,name=call_set_name,json=callSetName,proto3" json:"call_set_name,omitempty"`
	// The ID of the call set this variant call belongs to.
	//
	// If this field is not present, the ordering of the call sets from a
	// `SearchCallSetsRequest` over this `VariantSet` is guaranteed to match
	// the ordering of the calls on this `Variant`.
	// The number of results will also be the same.
	CallSetId string `protobuf:"bytes,2,opt,name=call_set_id,json=callSetId,proto3" json:"call_set_id,omitempty"`
	// The genotype of this variant call.
	//
	// A 0 value represents the reference allele of the associated `Variant`. Any
	// other value is a 1-based index into the alternate alleles of the associated
	// `Variant`.
	//
	// If a variant had a referenceBases field of "T", an alternateBases
	// value of ["A", "C"], and the genotype was [2, 1], that would mean the call
	// represented the heterozygous value "CA" for this variant. If the genotype
	// was instead [0, 1] the represented value would be "TA". Ordering of the
	// genotype values is important if the phaseset field is present.
	// Missing genotype genotypes may be indicated using the "dot annotation" [".", "."],
	// as specified in VCF4.2; this is e.g. used for types of structural variants.
	Genotype *_struct.ListValue `protobuf:"bytes,7,opt,name=genotype,proto3" json:"genotype,omitempty"`
	// If this field is populated, this variant call's genotype ordering implies
	// the phase of the bases and is consistent with any other variant calls on
	// the same contig which have the same phaseset string.
	Phaseset string `protobuf:"bytes,4,opt,name=phaseset,proto3" json:"phaseset,omitempty"`
	// The genotype likelihoods for this variant call. Each array entry
	// represents how likely a specific genotype is for this call as
	// log10(P(data | genotype)), analogous to the GL tag in the VCF spec. The
	// value ordering is defined by the GL tag in the VCF spec.
	GenotypeLikelihood []float64 `protobuf:"fixed64,5,rep,packed,name=genotype_likelihood,json=genotypeLikelihood,proto3" json:"genotype_likelihood,omitempty"`
	// A map of additional information about the Call.
	Attributes           *Attributes `protobuf:"bytes,8,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c8f728f42a356bd, []int{3}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetCallSetName() string {
	if m != nil {
		return m.CallSetName
	}
	return ""
}

func (m *Call) GetCallSetId() string {
	if m != nil {
		return m.CallSetId
	}
	return ""
}

func (m *Call) GetGenotype() *_struct.ListValue {
	if m != nil {
		return m.Genotype
	}
	return nil
}

func (m *Call) GetPhaseset() string {
	if m != nil {
		return m.Phaseset
	}
	return ""
}

func (m *Call) GetGenotypeLikelihood() []float64 {
	if m != nil {
		return m.GenotypeLikelihood
	}
	return nil
}

func (m *Call) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A `Variant` represents a change in DNA sequence relative to some reference.
// For example, a variant could represent a SNP or an insertion.
// Variants belong to a `VariantSet`.
// This is equivalent to a row in VCF.
type Variant struct {
	// The variant ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the `VariantSet` this variant belongs to. This transitively
	// defines
	// the `ReferenceSet` against which the `Variant` is to be interpreted.
	VariantSetId string `protobuf:"bytes,2,opt,name=variant_set_id,json=variantSetId,proto3" json:"variant_set_id,omitempty"`
	// Names for the variant, for example a RefSNP ID.
	Names []string `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	// The date this variant was created in milliseconds from the epoch.
	Created int64 `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	// The time at which this variant was last updated in
	// milliseconds from the epoch.
	Updated int64 `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	// The reference on which this variant occurs.
	// (e.g. `chr20` or `X`)
	ReferenceName string `protobuf:"bytes,6,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The start position at which this variant occurs (0-based).
	// This corresponds to the first base of the string of reference bases.
	// Genomic positions are non-negative integers less than reference length.
	// Variants spanning the join of circular genomes are represented as
	// two variants one on each side of the join (position 0).
	Start int64 `protobuf:"varint,7,opt,name=start,proto3" json:"start,omitempty"`
	// The end position (exclusive), resulting in [start, end) closed-open
	// interval.
	// This is typically calculated by `start + referenceBases.length`.
	End int64 `protobuf:"varint,8,opt,name=end,proto3" json:"end,omitempty"`
	// The reference bases for this variant. They start at the given start
	// position.
	ReferenceBases string `protobuf:"bytes,9,opt,name=reference_bases,json=referenceBases,proto3" json:"reference_bases,omitempty"`
	// The bases that appear instead of the reference bases. Multiple alternate
	// alleles are possible.
	AlternateBases []string `protobuf:"bytes,10,rep,name=alternate_bases,json=alternateBases,proto3" json:"alternate_bases,omitempty"`
	// A map of additional information about the Variant.
	Attributes *Attributes `protobuf:"bytes,13,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The variant calls for this particular variant. Each one represents the
	// determination of genotype with respect to this variant. `Call`s in this
	// array are implicitly associated with this `Variant`.
	Calls []*Call `protobuf:"bytes,12,rep,name=calls,proto3" json:"calls,omitempty"`
	// The "variant_type" is used to denote e.g. structural variants.
	// Examples:
	//   DUP  : duplication of sequence following "start"; not necessarily in situ
	//   DEL  : deletion of sequence following "start"
	VariantType string `protobuf:"bytes,17,opt,name=variant_type,json=variantType,proto3" json:"variant_type,omitempty"`
	// Length of the - if labeled as such in variant_type - structural variation.
	// Based on the use in VCFv4.2
	Svlen int64 `protobuf:"varint,18,opt,name=svlen,proto3" json:"svlen,omitempty"`
	// In the case of structural variants, start and end of the variant may not
	// be known with an exact base position. "cipos" provides an interval with
	// high confidence for the start position. The interval is provided by 0 or
	// 2 signed integers which are added to the start position.
	// Based on the use in VCFv4.2
	// Example:
	//   [ -12000, 1000 ]
	Cipos []int32 `protobuf:"zigzag32,19,rep,packed,name=cipos,proto3" json:"cipos,omitempty"`
	// Similar to "cipos", but for the variant's end position (which is derived
	// from start + svlen).
	// Example:
	//   [ -1000, 0 ]
	Ciend []int32 `protobuf:"zigzag32,20,rep,packed,name=ciend,proto3" json:"ciend,omitempty"`
	// True if filters were applied for this variant. VCF column 7 "FILTER"
	// any value other than the missing value.
	FiltersApplied bool `protobuf:"varint,14,opt,name=filters_applied,json=filtersApplied,proto3" json:"filters_applied,omitempty"`
	// True if all filters for this variant passed. VCF column 7 "FILTER"
	// value PASS.
	FiltersPassed bool `protobuf:"varint,15,opt,name=filters_passed,json=filtersPassed,proto3" json:"filters_passed,omitempty"`
	// Zero or more filters that failed for this variant. VCF column 7 "FILTER"
	// shared across all alleles in the same VCF record.
	FiltersFailed []string `protobuf:"bytes,16,rep,name=filters_failed,json=filtersFailed,proto3" json:"filters_failed,omitempty"`
	// Patient_Id, this field is not populated in the database, but generated on-the-fly
	// when variants are being returned.
	PatientId            string   `protobuf:"bytes,21,opt,name=patientId,proto3" json:"patientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Variant) Reset()         { *m = Variant{} }
func (m *Variant) String() string { return proto.CompactTextString(m) }
func (*Variant) ProtoMessage()    {}
func (*Variant) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c8f728f42a356bd, []int{4}
}

func (m *Variant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variant.Unmarshal(m, b)
}
func (m *Variant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variant.Marshal(b, m, deterministic)
}
func (m *Variant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variant.Merge(m, src)
}
func (m *Variant) XXX_Size() int {
	return xxx_messageInfo_Variant.Size(m)
}
func (m *Variant) XXX_DiscardUnknown() {
	xxx_messageInfo_Variant.DiscardUnknown(m)
}

var xxx_messageInfo_Variant proto.InternalMessageInfo

func (m *Variant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Variant) GetVariantSetId() string {
	if m != nil {
		return m.VariantSetId
	}
	return ""
}

func (m *Variant) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Variant) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Variant) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Variant) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *Variant) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Variant) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Variant) GetReferenceBases() string {
	if m != nil {
		return m.ReferenceBases
	}
	return ""
}

func (m *Variant) GetAlternateBases() []string {
	if m != nil {
		return m.AlternateBases
	}
	return nil
}

func (m *Variant) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Variant) GetCalls() []*Call {
	if m != nil {
		return m.Calls
	}
	return nil
}

func (m *Variant) GetVariantType() string {
	if m != nil {
		return m.VariantType
	}
	return ""
}

func (m *Variant) GetSvlen() int64 {
	if m != nil {
		return m.Svlen
	}
	return 0
}

func (m *Variant) GetCipos() []int32 {
	if m != nil {
		return m.Cipos
	}
	return nil
}

func (m *Variant) GetCiend() []int32 {
	if m != nil {
		return m.Ciend
	}
	return nil
}

func (m *Variant) GetFiltersApplied() bool {
	if m != nil {
		return m.FiltersApplied
	}
	return false
}

func (m *Variant) GetFiltersPassed() bool {
	if m != nil {
		return m.FiltersPassed
	}
	return false
}

func (m *Variant) GetFiltersFailed() []string {
	if m != nil {
		return m.FiltersFailed
	}
	return nil
}

func (m *Variant) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func init() {
	proto.RegisterType((*VariantSetMetadata)(nil), "candig.VariantSetMetadata")
	proto.RegisterType((*VariantSet)(nil), "candig.VariantSet")
	proto.RegisterType((*CallSet)(nil), "candig.CallSet")
	proto.RegisterType((*Call)(nil), "candig.Call")
	proto.RegisterType((*Variant)(nil), "candig.Variant")
}

func init() { proto.RegisterFile("candig/variants.proto", fileDescriptor_9c8f728f42a356bd) }

var fileDescriptor_9c8f728f42a356bd = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdb, 0x6a, 0xdb, 0x4a,
	0x14, 0x86, 0x91, 0xe5, 0x93, 0x96, 0x0f, 0x49, 0x26, 0xc9, 0x66, 0xf0, 0x4e, 0x36, 0xda, 0xa6,
	0x07, 0x5d, 0xd9, 0x90, 0x42, 0xef, 0xd3, 0x42, 0xc1, 0x90, 0x96, 0xa2, 0x94, 0xdc, 0x9a, 0xb1,
	0x66, 0xec, 0x0c, 0x91, 0x25, 0xa1, 0x19, 0x1b, 0x72, 0xdd, 0xf7, 0xeb, 0x33, 0xf4, 0x05, 0xfa,
	0x06, 0xbd, 0x28, 0x33, 0x6b, 0x24, 0x3b, 0x09, 0x85, 0x90, 0xbb, 0x59, 0xff, 0xfa, 0x91, 0xe6,
	0x5b, 0x87, 0x81, 0xd3, 0x84, 0x65, 0x5c, 0xae, 0xa6, 0x5b, 0x56, 0x4a, 0x96, 0x69, 0x35, 0x29,
	0xca, 0x5c, 0xe7, 0xa4, 0x8d, 0xf2, 0xe8, 0xd8, 0xa5, 0x93, 0x7c, 0xbd, 0xce, 0x33, 0x4c, 0x8e,
	0xce, 0x56, 0x79, 0xbe, 0x4a, 0xc5, 0xd4, 0x46, 0x8b, 0xcd, 0x72, 0xaa, 0x74, 0xb9, 0x49, 0x34,
	0x66, 0xc7, 0x3f, 0x3c, 0x20, 0x37, 0xf8, 0xb5, 0x6b, 0xa1, 0x3f, 0x0b, 0xcd, 0x38, 0xd3, 0x8c,
	0x1c, 0x82, 0x7f, 0x27, 0xee, 0xa9, 0x17, 0x7a, 0x51, 0x10, 0x9b, 0x23, 0x39, 0x81, 0xd6, 0x96,
	0xa5, 0x1b, 0x41, 0x1b, 0x56, 0xc3, 0x80, 0x0c, 0xa1, 0x21, 0x39, 0xf5, 0xad, 0xd4, 0x90, 0x9c,
	0x10, 0x68, 0xea, 0xfb, 0x42, 0xd0, 0xa6, 0x55, 0xec, 0x99, 0xfc, 0x03, 0xed, 0x6c, 0xb3, 0x5e,
	0x88, 0x92, 0xb6, 0xac, 0xea, 0x22, 0x12, 0x42, 0x8f, 0x0b, 0x95, 0x94, 0xb2, 0xd0, 0x32, 0xcf,
	0x68, 0xdb, 0x26, 0xf7, 0x25, 0x72, 0x01, 0xc0, 0xb4, 0x2e, 0xe5, 0x62, 0xa3, 0x85, 0xa2, 0xdd,
	0xd0, 0x8b, 0x7a, 0x17, 0x64, 0x82, 0x90, 0x93, 0xcb, 0x3a, 0x13, 0xef, 0xb9, 0xc6, 0xbf, 0x3c,
	0x80, 0x1d, 0x90, 0xbb, 0xa0, 0xb7, 0x7f, 0xc1, 0x8c, 0xad, 0x2b, 0x0a, 0x7b, 0x26, 0xe7, 0x00,
	0x06, 0x5a, 0x09, 0x3d, 0xaf, 0x61, 0x02, 0xa7, 0xcc, 0x38, 0x89, 0xe0, 0xb0, 0x14, 0x4b, 0x51,
	0x8a, 0x2c, 0x11, 0x73, 0x67, 0x42, 0xbe, 0x61, 0xad, 0x5f, 0x5b, 0xe7, 0x39, 0x40, 0xc1, 0xb4,
	0x14, 0x99, 0xf5, 0x20, 0x6d, 0xe0, 0x94, 0x19, 0x27, 0xff, 0x42, 0xa0, 0xd8, 0xba, 0x48, 0x85,
	0xc9, 0x22, 0x6e, 0x17, 0x85, 0x19, 0x27, 0xef, 0xa1, 0xbb, 0x76, 0xd5, 0xa7, 0x9d, 0xd0, 0x8f,
	0x7a, 0x17, 0xa3, 0x8a, 0xf4, 0x69, 0x7f, 0xe2, 0xda, 0x3b, 0xfe, 0xe9, 0x41, 0xe7, 0x23, 0x4b,
	0xd3, 0xe7, 0xc2, 0xfe, 0x0f, 0xfd, 0x85, 0xcc, 0x77, 0xf7, 0x40, 0xdc, 0x5e, 0xad, 0xcd, 0x38,
	0x79, 0x03, 0x07, 0x6e, 0xc0, 0x1c, 0xae, 0xa2, 0xcd, 0xd0, 0x8f, 0x82, 0x78, 0xb0, 0xad, 0x6f,
	0x32, 0xe3, 0x8a, 0x50, 0xe8, 0x24, 0xa5, 0x60, 0x5a, 0x20, 0xab, 0x1f, 0x57, 0xa1, 0xc9, 0x6c,
	0x0a, 0x6e, 0x33, 0x6d, 0xcc, 0xb8, 0xf0, 0x45, 0x2d, 0xfd, 0xde, 0x80, 0xa6, 0x41, 0x24, 0x63,
	0x18, 0x24, 0x2c, 0x4d, 0xed, 0xad, 0x2c, 0x18, 0xa2, 0xf6, 0x12, 0xe4, 0xff, 0x62, 0xf8, 0xfe,
	0x83, 0x5e, 0xed, 0x91, 0xdc, 0xa1, 0x07, 0xce, 0x81, 0x75, 0x5e, 0x89, 0x2c, 0xb7, 0x53, 0xda,
	0xb1, 0xbf, 0x1f, 0x4d, 0x70, 0x43, 0x26, 0xd5, 0x86, 0x4c, 0xae, 0xa4, 0xd2, 0x37, 0x66, 0xbe,
	0xe3, 0xda, 0x4b, 0x46, 0xd0, 0x2d, 0x6e, 0x99, 0x12, 0x4a, 0x68, 0xd7, 0xfd, 0x3a, 0x26, 0x53,
	0x38, 0xae, 0x7c, 0xf3, 0x54, 0xde, 0x89, 0x54, 0xde, 0xe6, 0xb9, 0x29, 0x8a, 0x1f, 0x79, 0x31,
	0xa9, 0x52, 0x57, 0x75, 0xe6, 0x45, 0x55, 0xf8, 0xdd, 0x84, 0x8e, 0x9b, 0x84, 0x27, 0x8d, 0x7e,
	0x05, 0xc3, 0x87, 0x1d, 0x73, 0xdc, 0xfd, 0xfd, 0x86, 0x99, 0x15, 0x36, 0x55, 0x53, 0xd4, 0xb7,
	0xdd, 0xc4, 0x60, 0xbf, 0x8b, 0xcd, 0xbf, 0x76, 0xb1, 0xf5, 0xb0, 0x8b, 0xaf, 0x61, 0x37, 0xfa,
	0xd8, 0x09, 0x1c, 0xe7, 0x41, 0xad, 0xda, 0x5e, 0x9c, 0x40, 0x4b, 0x69, 0x56, 0x6a, 0x5b, 0x68,
	0x3f, 0xc6, 0xc0, 0xbc, 0x2d, 0x22, 0xe3, 0x96, 0xda, 0x8f, 0xcd, 0x91, 0xbc, 0x85, 0x83, 0xdd,
	0xe7, 0x16, 0xa6, 0xa8, 0x34, 0x78, 0xb4, 0x60, 0x1f, 0x8c, 0x6a, 0x8c, 0x2c, 0xd5, 0xa2, 0xcc,
	0x98, 0xae, 0x8c, 0x60, 0x59, 0x86, 0xb5, 0x8c, 0xc6, 0x87, 0x05, 0x1e, 0x3c, 0xa7, 0xc0, 0x64,
	0x0c, 0x2d, 0x33, 0x26, 0x8a, 0xf6, 0xed, 0xfa, 0xf5, 0x2b, 0xbb, 0x19, 0xbd, 0x18, 0x53, 0x66,
	0x7b, 0xaa, 0x42, 0xdb, 0x09, 0x3a, 0xc2, 0x01, 0x74, 0xda, 0x37, 0x33, 0x28, 0x06, 0x7a, 0x9b,
	0x8a, 0x8c, 0x12, 0x07, 0x6d, 0x02, 0xa3, 0x26, 0xb2, 0xc8, 0x15, 0x3d, 0x0e, 0xfd, 0xe8, 0x28,
	0xc6, 0x00, 0x55, 0x53, 0x8c, 0x93, 0x4a, 0x75, 0xe5, 0x58, 0x4a, 0xc3, 0xa3, 0xe6, 0xac, 0x28,
	0x52, 0x29, 0x38, 0x1d, 0x86, 0x5e, 0xd4, 0x8d, 0x87, 0x4e, 0xbe, 0x44, 0xd5, 0xb4, 0xa1, 0x32,
	0x16, 0x4c, 0x29, 0xc1, 0xe9, 0x81, 0xf5, 0x0d, 0x9c, 0xfa, 0xd5, 0x8a, 0xfb, 0xb6, 0x25, 0x93,
	0xa9, 0xe0, 0xf4, 0x10, 0xd7, 0xd9, 0xa9, 0x9f, 0xac, 0x48, 0xce, 0x60, 0xf7, 0x56, 0xd1, 0xd3,
	0x47, 0x8f, 0xd7, 0xa2, 0x6d, 0xb7, 0xe3, 0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x37,
	0x4f, 0x8b, 0x83, 0x06, 0x00, 0x00,
}