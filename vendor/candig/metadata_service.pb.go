// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candig/metadata_service.proto

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

// This request maps to the body of `POST /datasets/search` as JSON.
type SearchDatasetsRequest struct {
	// Specifies the maximum number of results to return in a single page.
	// If unspecified, a system default will be used.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `next_page_token` from the previous response.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchDatasetsRequest) Reset()         { *m = SearchDatasetsRequest{} }
func (m *SearchDatasetsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchDatasetsRequest) ProtoMessage()    {}
func (*SearchDatasetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7be1a57dcff81f, []int{0}
}

func (m *SearchDatasetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchDatasetsRequest.Unmarshal(m, b)
}
func (m *SearchDatasetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchDatasetsRequest.Marshal(b, m, deterministic)
}
func (m *SearchDatasetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchDatasetsRequest.Merge(m, src)
}
func (m *SearchDatasetsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchDatasetsRequest.Size(m)
}
func (m *SearchDatasetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchDatasetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchDatasetsRequest proto.InternalMessageInfo

func (m *SearchDatasetsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchDatasetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// This is the response from `POST /datasets/search` expressed as JSON.
type SearchDatasetsResponse struct {
	// The list of datasets.
	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets,proto3" json:"datasets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchDatasetsResponse) Reset()         { *m = SearchDatasetsResponse{} }
func (m *SearchDatasetsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchDatasetsResponse) ProtoMessage()    {}
func (*SearchDatasetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7be1a57dcff81f, []int{1}
}

func (m *SearchDatasetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchDatasetsResponse.Unmarshal(m, b)
}
func (m *SearchDatasetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchDatasetsResponse.Marshal(b, m, deterministic)
}
func (m *SearchDatasetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchDatasetsResponse.Merge(m, src)
}
func (m *SearchDatasetsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchDatasetsResponse.Size(m)
}
func (m *SearchDatasetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchDatasetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchDatasetsResponse proto.InternalMessageInfo

func (m *SearchDatasetsResponse) GetDatasets() []*Dataset {
	if m != nil {
		return m.Datasets
	}
	return nil
}

func (m *SearchDatasetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// This request maps to the URL `GET /datasets/{dataset_id}`.
type GetDatasetRequest struct {
	// The ID of the `Dataset` to be retrieved.
	DatasetId            string   `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatasetRequest) Reset()         { *m = GetDatasetRequest{} }
func (m *GetDatasetRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatasetRequest) ProtoMessage()    {}
func (*GetDatasetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7be1a57dcff81f, []int{2}
}

func (m *GetDatasetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatasetRequest.Unmarshal(m, b)
}
func (m *GetDatasetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatasetRequest.Marshal(b, m, deterministic)
}
func (m *GetDatasetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatasetRequest.Merge(m, src)
}
func (m *GetDatasetRequest) XXX_Size() int {
	return xxx_messageInfo_GetDatasetRequest.Size(m)
}
func (m *GetDatasetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatasetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatasetRequest proto.InternalMessageInfo

func (m *GetDatasetRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchDatasetsRequest)(nil), "candig.SearchDatasetsRequest")
	proto.RegisterType((*SearchDatasetsResponse)(nil), "candig.SearchDatasetsResponse")
	proto.RegisterType((*GetDatasetRequest)(nil), "candig.GetDatasetRequest")
}

func init() { proto.RegisterFile("candig/metadata_service.proto", fileDescriptor_4b7be1a57dcff81f) }

var fileDescriptor_4b7be1a57dcff81f = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0xa5, 0xfb, 0xf9, 0xc7, 0xfa, 0x89, 0x0e, 0x03, 0xd3, 0x59, 0x57, 0x1d, 0x05, 0x65, 0x28,
	0x34, 0x63, 0xde, 0x88, 0xd7, 0x82, 0x78, 0x21, 0x48, 0xeb, 0x7d, 0x89, 0xeb, 0x47, 0x0d, 0xba,
	0xa4, 0x2e, 0x71, 0xc8, 0xc4, 0x1b, 0x5f, 0xc1, 0x47, 0xf3, 0x15, 0x04, 0x5f, 0x43, 0x96, 0xa4,
	0x1b, 0x6e, 0x7a, 0x7b, 0xce, 0xc9, 0xf9, 0xce, 0x39, 0x81, 0x70, 0xc8, 0x44, 0xce, 0x0b, 0x3a,
	0x42, 0xcd, 0x72, 0xa6, 0x59, 0xa6, 0x70, 0x3c, 0xe1, 0x43, 0x8c, 0xcb, 0xb1, 0xd4, 0x92, 0xd4,
	0x2d, 0x1d, 0xb4, 0x96, 0x64, 0x96, 0x0e, 0x3a, 0x85, 0x94, 0xc5, 0x03, 0x52, 0x56, 0x72, 0xca,
	0x84, 0x90, 0x9a, 0x69, 0x2e, 0x85, 0xb2, 0x6c, 0x94, 0x42, 0x2b, 0x45, 0x36, 0x1e, 0xde, 0x9d,
	0x33, 0xcd, 0x14, 0x6a, 0x95, 0xe0, 0xe3, 0x13, 0x2a, 0x4d, 0x76, 0xc1, 0x2f, 0x59, 0x81, 0x99,
	0xe2, 0x53, 0x6c, 0x7b, 0x5d, 0xaf, 0xf7, 0x3f, 0x69, 0xcc, 0x80, 0x94, 0x4f, 0x91, 0x84, 0x00,
	0x86, 0xd4, 0xf2, 0x1e, 0x45, 0xbb, 0xd6, 0xf5, 0x7a, 0x7e, 0x62, 0xe4, 0x37, 0x33, 0x20, 0x1a,
	0xc1, 0xd6, 0xb2, 0xa9, 0x2a, 0xa5, 0x50, 0x48, 0x8e, 0xa1, 0x91, 0x3b, 0xac, 0xed, 0x75, 0xff,
	0xf5, 0xd6, 0x06, 0xcd, 0xd8, 0xc6, 0x8e, 0x9d, 0x36, 0x99, 0x0b, 0xc8, 0x21, 0x34, 0x05, 0x3e,
	0xeb, 0x6c, 0xe5, 0xd4, 0xfa, 0x0c, 0xbe, 0x9e, 0x9f, 0x1b, 0xc0, 0xe6, 0x05, 0xea, 0xea, 0xbd,
	0xcb, 0x1f, 0x02, 0x38, 0xa3, 0x8c, 0xe7, 0xa6, 0x80, 0x9f, 0xf8, 0x0e, 0xb9, 0xcc, 0x07, 0x5f,
	0x1e, 0x34, 0xaf, 0xdc, 0x50, 0xa9, 0x9d, 0x93, 0x28, 0xd8, 0xf8, 0x19, 0x9b, 0x84, 0x55, 0xb8,
	0x5f, 0x37, 0x0a, 0xf6, 0xfe, 0xa2, 0x6d, 0xdb, 0x28, 0x7a, 0xfb, 0xf8, 0x7c, 0xaf, 0x75, 0xa2,
	0x6d, 0x3a, 0xe9, 0xc7, 0xa7, 0x71, 0x9f, 0x56, 0xd5, 0xa8, 0x32, 0x0f, 0xce, 0xbc, 0x23, 0x92,
	0x01, 0x2c, 0xc2, 0x93, 0x9d, 0xca, 0x71, 0xa5, 0x50, 0xb0, 0x3c, 0x54, 0x74, 0x60, 0xdc, 0xf7,
	0x49, 0xb8, 0xe2, 0xfe, 0xb2, 0x68, 0xfe, 0x7a, 0x5b, 0x37, 0x1f, 0x7d, 0xf2, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x83, 0x3c, 0xd5, 0xfa, 0x46, 0x02, 0x00, 0x00,
}