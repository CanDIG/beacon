// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candig/reference_service.proto

package candig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google/api"
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

// ****************  /referencesets  *******************
// This request maps to the body of `POST /referencesets/search` as JSON.
type SearchReferenceSetsRequest struct {
	// If unset, return the reference sets for which the
	// `md5checksum` matches this string (case-sensitive, exact match).
	// See `ReferenceSet::md5checksum` for details.
	Md5Checksum string `protobuf:"bytes,1,opt,name=md5checksum,proto3" json:"md5checksum,omitempty"`
	// If unset, return the reference sets for which the `accession`
	// matches this string (case-sensitive, exact match).
	Accession string `protobuf:"bytes,2,opt,name=accession,proto3" json:"accession,omitempty"`
	// If unset, return the reference sets for which the `assemblyId`
	// matches this string (case-sensitive, exact match).
	AssemblyId string `protobuf:"bytes,3,opt,name=assembly_id,json=assemblyId,proto3" json:"assembly_id,omitempty"`
	// Specifies the maximum number of results to return in a single page.
	// If unspecified, a system default will be used.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `next_page_token` from the previous response.
	PageToken            string   `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReferenceSetsRequest) Reset()         { *m = SearchReferenceSetsRequest{} }
func (m *SearchReferenceSetsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchReferenceSetsRequest) ProtoMessage()    {}
func (*SearchReferenceSetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{0}
}

func (m *SearchReferenceSetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReferenceSetsRequest.Unmarshal(m, b)
}
func (m *SearchReferenceSetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReferenceSetsRequest.Marshal(b, m, deterministic)
}
func (m *SearchReferenceSetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReferenceSetsRequest.Merge(m, src)
}
func (m *SearchReferenceSetsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchReferenceSetsRequest.Size(m)
}
func (m *SearchReferenceSetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReferenceSetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReferenceSetsRequest proto.InternalMessageInfo

func (m *SearchReferenceSetsRequest) GetMd5Checksum() string {
	if m != nil {
		return m.Md5Checksum
	}
	return ""
}

func (m *SearchReferenceSetsRequest) GetAccession() string {
	if m != nil {
		return m.Accession
	}
	return ""
}

func (m *SearchReferenceSetsRequest) GetAssemblyId() string {
	if m != nil {
		return m.AssemblyId
	}
	return ""
}

func (m *SearchReferenceSetsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchReferenceSetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// This is the response from `POST /referencesets/search` expressed as JSON.
type SearchReferenceSetsResponse struct {
	// The list of matching reference sets.
	ReferenceSets []*ReferenceSet `protobuf:"bytes,1,rep,name=reference_sets,json=referenceSets,proto3" json:"reference_sets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReferenceSetsResponse) Reset()         { *m = SearchReferenceSetsResponse{} }
func (m *SearchReferenceSetsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchReferenceSetsResponse) ProtoMessage()    {}
func (*SearchReferenceSetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{1}
}

func (m *SearchReferenceSetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReferenceSetsResponse.Unmarshal(m, b)
}
func (m *SearchReferenceSetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReferenceSetsResponse.Marshal(b, m, deterministic)
}
func (m *SearchReferenceSetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReferenceSetsResponse.Merge(m, src)
}
func (m *SearchReferenceSetsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchReferenceSetsResponse.Size(m)
}
func (m *SearchReferenceSetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReferenceSetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReferenceSetsResponse proto.InternalMessageInfo

func (m *SearchReferenceSetsResponse) GetReferenceSets() []*ReferenceSet {
	if m != nil {
		return m.ReferenceSets
	}
	return nil
}

func (m *SearchReferenceSetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// This request maps to the URL `GET /referencesets/{reference_set_id}`.
type GetReferenceSetRequest struct {
	// The ID of the `ReferenceSet` to be retrieved.
	ReferenceSetId       string   `protobuf:"bytes,1,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReferenceSetRequest) Reset()         { *m = GetReferenceSetRequest{} }
func (m *GetReferenceSetRequest) String() string { return proto.CompactTextString(m) }
func (*GetReferenceSetRequest) ProtoMessage()    {}
func (*GetReferenceSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{2}
}

func (m *GetReferenceSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReferenceSetRequest.Unmarshal(m, b)
}
func (m *GetReferenceSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReferenceSetRequest.Marshal(b, m, deterministic)
}
func (m *GetReferenceSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReferenceSetRequest.Merge(m, src)
}
func (m *GetReferenceSetRequest) XXX_Size() int {
	return xxx_messageInfo_GetReferenceSetRequest.Size(m)
}
func (m *GetReferenceSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReferenceSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReferenceSetRequest proto.InternalMessageInfo

func (m *GetReferenceSetRequest) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

// ****************  /references  *******************
// This request maps to the body of `POST /references/search` as JSON.
type SearchReferencesRequest struct {
	// The `ReferenceSet` to search.
	ReferenceSetId string `protobuf:"bytes,1,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// If specified, return the references for which the
	// `md5checksum` matches this string (case-sensitive, exact match).
	// See `ReferenceSet::md5checksum` for details.
	Md5Checksum string `protobuf:"bytes,2,opt,name=md5checksum,proto3" json:"md5checksum,omitempty"`
	// If specified, return the references for which the `accession`
	// matches this string (case-sensitive, exact match).
	Accession string `protobuf:"bytes,3,opt,name=accession,proto3" json:"accession,omitempty"`
	// Specifies the maximum number of results to return in a single page.
	// If unspecified, a system default will be used.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `next_page_token` from the previous response.
	PageToken            string   `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReferencesRequest) Reset()         { *m = SearchReferencesRequest{} }
func (m *SearchReferencesRequest) String() string { return proto.CompactTextString(m) }
func (*SearchReferencesRequest) ProtoMessage()    {}
func (*SearchReferencesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{3}
}

func (m *SearchReferencesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReferencesRequest.Unmarshal(m, b)
}
func (m *SearchReferencesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReferencesRequest.Marshal(b, m, deterministic)
}
func (m *SearchReferencesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReferencesRequest.Merge(m, src)
}
func (m *SearchReferencesRequest) XXX_Size() int {
	return xxx_messageInfo_SearchReferencesRequest.Size(m)
}
func (m *SearchReferencesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReferencesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReferencesRequest proto.InternalMessageInfo

func (m *SearchReferencesRequest) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *SearchReferencesRequest) GetMd5Checksum() string {
	if m != nil {
		return m.Md5Checksum
	}
	return ""
}

func (m *SearchReferencesRequest) GetAccession() string {
	if m != nil {
		return m.Accession
	}
	return ""
}

func (m *SearchReferencesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchReferencesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// This is the response from `POST /references/search` expressed as JSON.
type SearchReferencesResponse struct {
	// The list of matching references.
	References []*Reference `protobuf:"bytes,1,rep,name=references,proto3" json:"references,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReferencesResponse) Reset()         { *m = SearchReferencesResponse{} }
func (m *SearchReferencesResponse) String() string { return proto.CompactTextString(m) }
func (*SearchReferencesResponse) ProtoMessage()    {}
func (*SearchReferencesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{4}
}

func (m *SearchReferencesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReferencesResponse.Unmarshal(m, b)
}
func (m *SearchReferencesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReferencesResponse.Marshal(b, m, deterministic)
}
func (m *SearchReferencesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReferencesResponse.Merge(m, src)
}
func (m *SearchReferencesResponse) XXX_Size() int {
	return xxx_messageInfo_SearchReferencesResponse.Size(m)
}
func (m *SearchReferencesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReferencesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReferencesResponse proto.InternalMessageInfo

func (m *SearchReferencesResponse) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *SearchReferencesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// This request maps to the URL `GET /references/{reference_id}`.
type GetReferenceRequest struct {
	// The ID of the `Reference` to be retrieved.
	ReferenceId          string   `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReferenceRequest) Reset()         { *m = GetReferenceRequest{} }
func (m *GetReferenceRequest) String() string { return proto.CompactTextString(m) }
func (*GetReferenceRequest) ProtoMessage()    {}
func (*GetReferenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{5}
}

func (m *GetReferenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReferenceRequest.Unmarshal(m, b)
}
func (m *GetReferenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReferenceRequest.Marshal(b, m, deterministic)
}
func (m *GetReferenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReferenceRequest.Merge(m, src)
}
func (m *GetReferenceRequest) XXX_Size() int {
	return xxx_messageInfo_GetReferenceRequest.Size(m)
}
func (m *GetReferenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReferenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReferenceRequest proto.InternalMessageInfo

func (m *GetReferenceRequest) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

// This request retrieves a region of a reference genome when sent to
// the `/listreferencebases` endpoint.
type ListReferenceBasesRequest struct {
	// The ID of the `Reference` to be retrieved.
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	// The start position (0-based) of this query. Defaults to 0.
	// Genomic positions are non-negative integers less than reference length.
	// Requests spanning the join of circular genomes are represented as
	// two requests one on each side of the join (position 0).
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// The end position (0-based, exclusive) of this query. Defaults
	// to the length of this `Reference`.
	End int64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `next_page_token` from the previous response.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReferenceBasesRequest) Reset()         { *m = ListReferenceBasesRequest{} }
func (m *ListReferenceBasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListReferenceBasesRequest) ProtoMessage()    {}
func (*ListReferenceBasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{6}
}

func (m *ListReferenceBasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReferenceBasesRequest.Unmarshal(m, b)
}
func (m *ListReferenceBasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReferenceBasesRequest.Marshal(b, m, deterministic)
}
func (m *ListReferenceBasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReferenceBasesRequest.Merge(m, src)
}
func (m *ListReferenceBasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListReferenceBasesRequest.Size(m)
}
func (m *ListReferenceBasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReferenceBasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReferenceBasesRequest proto.InternalMessageInfo

func (m *ListReferenceBasesRequest) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

func (m *ListReferenceBasesRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListReferenceBasesRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ListReferenceBasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response from `POST /listreferencebases` expressed as JSON.
type ListReferenceBasesResponse struct {
	// The offset position (0-based) of the given sequence from the start of this
	// `Reference`. This value will differ for each page in a paginated request.
	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// A substring of the bases that make up this reference. Bases are represented
	// as IUPAC-IUB codes; this string matches the regexp `[ACGTMRWSYKVHDBN]*`.
	Sequence string `protobuf:"bytes,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReferenceBasesResponse) Reset()         { *m = ListReferenceBasesResponse{} }
func (m *ListReferenceBasesResponse) String() string { return proto.CompactTextString(m) }
func (*ListReferenceBasesResponse) ProtoMessage()    {}
func (*ListReferenceBasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41dd437aeb1914d, []int{7}
}

func (m *ListReferenceBasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReferenceBasesResponse.Unmarshal(m, b)
}
func (m *ListReferenceBasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReferenceBasesResponse.Marshal(b, m, deterministic)
}
func (m *ListReferenceBasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReferenceBasesResponse.Merge(m, src)
}
func (m *ListReferenceBasesResponse) XXX_Size() int {
	return xxx_messageInfo_ListReferenceBasesResponse.Size(m)
}
func (m *ListReferenceBasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReferenceBasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListReferenceBasesResponse proto.InternalMessageInfo

func (m *ListReferenceBasesResponse) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListReferenceBasesResponse) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

func (m *ListReferenceBasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchReferenceSetsRequest)(nil), "candig.SearchReferenceSetsRequest")
	proto.RegisterType((*SearchReferenceSetsResponse)(nil), "candig.SearchReferenceSetsResponse")
	proto.RegisterType((*GetReferenceSetRequest)(nil), "candig.GetReferenceSetRequest")
	proto.RegisterType((*SearchReferencesRequest)(nil), "candig.SearchReferencesRequest")
	proto.RegisterType((*SearchReferencesResponse)(nil), "candig.SearchReferencesResponse")
	proto.RegisterType((*GetReferenceRequest)(nil), "candig.GetReferenceRequest")
	proto.RegisterType((*ListReferenceBasesRequest)(nil), "candig.ListReferenceBasesRequest")
	proto.RegisterType((*ListReferenceBasesResponse)(nil), "candig.ListReferenceBasesResponse")
}

func init() { proto.RegisterFile("candig/reference_service.proto", fileDescriptor_e41dd437aeb1914d) }

var fileDescriptor_e41dd437aeb1914d = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xd6, 0x62, 0x82, 0x60, 0xc2, 0x4f, 0xba, 0x20, 0x30, 0x86, 0x42, 0xb2, 0x54, 0x25, 0x45,
	0x55, 0x4c, 0x53, 0x55, 0x42, 0xf4, 0xc6, 0xa5, 0x8a, 0xd4, 0x43, 0xe5, 0xf4, 0x1e, 0x39, 0xf6,
	0x24, 0x58, 0x10, 0x3b, 0xf5, 0x6e, 0x10, 0x10, 0x71, 0x41, 0xaa, 0x7a, 0xac, 0xaa, 0x3e, 0x4d,
	0x4f, 0x7d, 0x88, 0xbe, 0x42, 0x1f, 0xa4, 0xf2, 0xda, 0x4e, 0x16, 0xc7, 0x21, 0xad, 0x7a, 0x8b,
	0xe7, 0xf7, 0xfb, 0xbe, 0x99, 0xd9, 0xc0, 0x9e, 0x63, 0xfb, 0xae, 0xd7, 0x35, 0x43, 0xec, 0x60,
	0x88, 0xbe, 0x83, 0x2d, 0x8e, 0xe1, 0x95, 0xe7, 0x60, 0xad, 0x1f, 0x06, 0x22, 0xa0, 0x0b, 0xb1,
	0xdf, 0xd8, 0xca, 0xc6, 0xf1, 0x38, 0xc0, 0xd8, 0xed, 0x06, 0x41, 0xf7, 0x12, 0x4d, 0xbb, 0xef,
	0x99, 0xb6, 0xef, 0x07, 0xc2, 0x16, 0x5e, 0xe0, 0x27, 0x5e, 0xf6, 0x83, 0x80, 0xd1, 0x44, 0x3b,
	0x74, 0xce, 0xad, 0x34, 0xb1, 0x89, 0x82, 0x5b, 0xf8, 0x69, 0x80, 0x5c, 0xd0, 0x32, 0x14, 0x7b,
	0xee, 0x1b, 0xe7, 0x1c, 0x9d, 0x0b, 0x3e, 0xe8, 0xe9, 0xa4, 0x4c, 0xaa, 0x4b, 0x96, 0x6a, 0xa2,
	0xbb, 0xb0, 0x64, 0x3b, 0x0e, 0x72, 0xee, 0x05, 0xbe, 0x3e, 0x27, 0xfd, 0x63, 0x03, 0xdd, 0x87,
	0xa2, 0xcd, 0x39, 0xf6, 0xda, 0x97, 0x37, 0x2d, 0xcf, 0xd5, 0x35, 0xe9, 0x87, 0xd4, 0xd4, 0x70,
	0xe9, 0x0e, 0x2c, 0xf5, 0xed, 0x2e, 0xb6, 0xb8, 0x77, 0x8b, 0xfa, 0x7c, 0x99, 0x54, 0x0b, 0xd6,
	0x62, 0x64, 0x68, 0x7a, 0xb7, 0x48, 0x9f, 0x02, 0x48, 0xa7, 0x08, 0x2e, 0xd0, 0xd7, 0x0b, 0x71,
	0xf1, 0xc8, 0xf2, 0x31, 0x32, 0xb0, 0x7b, 0x02, 0x3b, 0xb9, 0xd8, 0x79, 0x3f, 0xf0, 0x39, 0xd2,
	0xb7, 0xb0, 0xaa, 0xaa, 0x26, 0xb8, 0x4e, 0xca, 0x5a, 0xb5, 0x58, 0xdf, 0xa8, 0xc5, 0x5a, 0xd5,
	0xd4, 0x34, 0x6b, 0x25, 0x54, 0x8b, 0xd0, 0xe7, 0xb0, 0xe6, 0xe3, 0xb5, 0x68, 0x29, 0x00, 0x62,
	0x76, 0x2b, 0x91, 0xf9, 0xc3, 0x08, 0xc4, 0x19, 0x6c, 0xbe, 0x43, 0xf1, 0xa0, 0x52, 0xa2, 0x5d,
	0x15, 0x4a, 0x0f, 0xda, 0x47, 0x02, 0xc4, 0x02, 0xae, 0xaa, 0xad, 0x1a, 0x2e, 0xfb, 0x49, 0x60,
	0x2b, 0x43, 0x84, 0xff, 0x73, 0x95, 0xec, 0xac, 0xe6, 0x66, 0xcc, 0x4a, 0xcb, 0xce, 0xea, 0x7f,
	0x46, 0x31, 0x00, 0x7d, 0x92, 0x40, 0x32, 0x86, 0x57, 0x00, 0xe3, 0xa5, 0x4c, 0x46, 0xf0, 0x64,
	0x62, 0x04, 0x96, 0x12, 0xf4, 0xd7, 0xe2, 0x9f, 0xc0, 0xba, 0x2a, 0x7e, 0xaa, 0x59, 0x05, 0x96,
	0xc7, 0x9a, 0x8d, 0xf4, 0x2a, 0x8e, 0x6c, 0x0d, 0x97, 0x7d, 0x26, 0xb0, 0xfd, 0xde, 0xe3, 0xe3,
	0xdc, 0x33, 0x9b, 0x8f, 0x45, 0x9f, 0x5d, 0x80, 0x6e, 0x40, 0x81, 0x0b, 0x3b, 0x14, 0x12, 0x98,
	0x66, 0xc5, 0x1f, 0xb4, 0x04, 0x1a, 0xfa, 0xf1, 0x9e, 0x6b, 0x56, 0xf4, 0x33, 0x23, 0xdc, 0x7c,
	0x56, 0xb8, 0x6b, 0x30, 0xf2, 0x60, 0x24, 0xd2, 0x6d, 0xc2, 0x42, 0xd0, 0xe9, 0x70, 0x14, 0x12,
	0x81, 0x66, 0x25, 0x5f, 0xd4, 0x80, 0x45, 0x1e, 0x41, 0xf5, 0x1d, 0x4c, 0x84, 0x19, 0x7d, 0xe7,
	0x69, 0xa7, 0xe5, 0x68, 0x57, 0xff, 0x56, 0x80, 0x92, 0xb2, 0xb6, 0xf2, 0x4d, 0xa1, 0x5f, 0x08,
	0xac, 0xe7, 0x9c, 0x14, 0x65, 0xe9, 0xbc, 0xa6, 0xbf, 0x15, 0xc6, 0xc1, 0xa3, 0x31, 0x31, 0x23,
	0x76, 0x78, 0xff, 0xeb, 0xf7, 0xf7, 0xb9, 0x0a, 0xdb, 0x35, 0xaf, 0x8e, 0x6b, 0x27, 0xb5, 0x63,
	0xe5, 0xbd, 0x42, 0xc1, 0x4d, 0x2e, 0x53, 0x4f, 0xc9, 0x11, 0xbd, 0x81, 0xb5, 0xcc, 0x5d, 0xd1,
	0xbd, 0xb4, 0x41, 0xfe, 0xc1, 0x19, 0xb9, 0x77, 0xcd, 0x8e, 0x65, 0xc7, 0x23, 0x5a, 0xcd, 0xef,
	0x38, 0xcc, 0x5e, 0xd7, 0x1d, 0x1d, 0x42, 0x29, 0xbb, 0xcc, 0x74, 0x7f, 0x0a, 0xb9, 0x11, 0xfb,
	0xf2, 0xf4, 0x80, 0x84, 0xfa, 0x33, 0x09, 0x64, 0x8f, 0x6d, 0x4f, 0x02, 0x51, 0x78, 0x7b, 0xb0,
	0xac, 0xd2, 0xa3, 0x3b, 0x79, 0xa4, 0xd3, 0xa6, 0x93, 0x67, 0xc4, 0x5e, 0xc8, 0x2e, 0x07, 0xb4,
	0x92, 0xd3, 0x65, 0xa8, 0x2e, 0xf5, 0x1d, 0xfd, 0x4a, 0x80, 0x4e, 0x2e, 0x1f, 0xad, 0xa4, 0x45,
	0xa7, 0xde, 0x87, 0xc1, 0x1e, 0x0b, 0x49, 0xe8, 0xd6, 0x25, 0x90, 0x97, 0xec, 0x70, 0x26, 0x10,
	0xb3, 0x1d, 0x25, 0x9e, 0x92, 0xa3, 0xf6, 0x82, 0xfc, 0x53, 0x7a, 0xfd, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xc8, 0x59, 0xae, 0xcc, 0xf5, 0x06, 0x00, 0x00,
}