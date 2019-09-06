// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candig/genotype_phenotype_service.proto

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

// This request maps to the body of `POST /phenotypeassociationsets/search` as JSON.
type SearchPhenotypeAssociationSetsRequest struct {
	// The `Dataset` to search. Mandatory
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Specifies the maximum number of results to return in a single page.
	// If unspecified, a system default will be used.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPhenotypeAssociationSetsRequest) Reset()         { *m = SearchPhenotypeAssociationSetsRequest{} }
func (m *SearchPhenotypeAssociationSetsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPhenotypeAssociationSetsRequest) ProtoMessage()    {}
func (*SearchPhenotypeAssociationSetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{0}
}

func (m *SearchPhenotypeAssociationSetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPhenotypeAssociationSetsRequest.Unmarshal(m, b)
}
func (m *SearchPhenotypeAssociationSetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPhenotypeAssociationSetsRequest.Marshal(b, m, deterministic)
}
func (m *SearchPhenotypeAssociationSetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPhenotypeAssociationSetsRequest.Merge(m, src)
}
func (m *SearchPhenotypeAssociationSetsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPhenotypeAssociationSetsRequest.Size(m)
}
func (m *SearchPhenotypeAssociationSetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPhenotypeAssociationSetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPhenotypeAssociationSetsRequest proto.InternalMessageInfo

func (m *SearchPhenotypeAssociationSetsRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *SearchPhenotypeAssociationSetsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchPhenotypeAssociationSetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// This is the response from `POST /phenotypeassociationsets/search` expressed as JSON.
type SearchPhenotypeAssociationSetsResponse struct {
	// The list of matching phenotype association sets.
	PhenotypeAssociationSets []*PhenotypeAssociationSet `protobuf:"bytes,1,rep,name=phenotype_association_sets,json=phenotypeAssociationSets,proto3" json:"phenotype_association_sets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPhenotypeAssociationSetsResponse) Reset() {
	*m = SearchPhenotypeAssociationSetsResponse{}
}
func (m *SearchPhenotypeAssociationSetsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchPhenotypeAssociationSetsResponse) ProtoMessage()    {}
func (*SearchPhenotypeAssociationSetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{1}
}

func (m *SearchPhenotypeAssociationSetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPhenotypeAssociationSetsResponse.Unmarshal(m, b)
}
func (m *SearchPhenotypeAssociationSetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPhenotypeAssociationSetsResponse.Marshal(b, m, deterministic)
}
func (m *SearchPhenotypeAssociationSetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPhenotypeAssociationSetsResponse.Merge(m, src)
}
func (m *SearchPhenotypeAssociationSetsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchPhenotypeAssociationSetsResponse.Size(m)
}
func (m *SearchPhenotypeAssociationSetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPhenotypeAssociationSetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPhenotypeAssociationSetsResponse proto.InternalMessageInfo

func (m *SearchPhenotypeAssociationSetsResponse) GetPhenotypeAssociationSets() []*PhenotypeAssociationSet {
	if m != nil {
		return m.PhenotypeAssociationSets
	}
	return nil
}

func (m *SearchPhenotypeAssociationSetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// One or more ontology terms can be queried together.
type OntologyTermQuery struct {
	Terms                []*OntologyTerm `protobuf:"bytes,1,rep,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OntologyTermQuery) Reset()         { *m = OntologyTermQuery{} }
func (m *OntologyTermQuery) String() string { return proto.CompactTextString(m) }
func (*OntologyTermQuery) ProtoMessage()    {}
func (*OntologyTermQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{2}
}

func (m *OntologyTermQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OntologyTermQuery.Unmarshal(m, b)
}
func (m *OntologyTermQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OntologyTermQuery.Marshal(b, m, deterministic)
}
func (m *OntologyTermQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OntologyTermQuery.Merge(m, src)
}
func (m *OntologyTermQuery) XXX_Size() int {
	return xxx_messageInfo_OntologyTermQuery.Size(m)
}
func (m *OntologyTermQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_OntologyTermQuery.DiscardUnknown(m)
}

var xxx_messageInfo_OntologyTermQuery proto.InternalMessageInfo

func (m *OntologyTermQuery) GetTerms() []*OntologyTerm {
	if m != nil {
		return m.Terms
	}
	return nil
}

// One or more ids can be queried together.  Generally used for instances
// of a particular class of object (e.g. a specific gene or SNP).
type ExternalIdentifierQuery struct {
	Ids                  []*ExternalIdentifier `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExternalIdentifierQuery) Reset()         { *m = ExternalIdentifierQuery{} }
func (m *ExternalIdentifierQuery) String() string { return proto.CompactTextString(m) }
func (*ExternalIdentifierQuery) ProtoMessage()    {}
func (*ExternalIdentifierQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{3}
}

func (m *ExternalIdentifierQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalIdentifierQuery.Unmarshal(m, b)
}
func (m *ExternalIdentifierQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalIdentifierQuery.Marshal(b, m, deterministic)
}
func (m *ExternalIdentifierQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalIdentifierQuery.Merge(m, src)
}
func (m *ExternalIdentifierQuery) XXX_Size() int {
	return xxx_messageInfo_ExternalIdentifierQuery.Size(m)
}
func (m *ExternalIdentifierQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalIdentifierQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalIdentifierQuery proto.InternalMessageInfo

func (m *ExternalIdentifierQuery) GetIds() []*ExternalIdentifier {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Evidence for the phenotype association.
type EvidenceQuery struct {
	// ECO or OBI is recommended
	EvidenceType *OntologyTerm `protobuf:"bytes,1,opt,name=evidenceType,proto3" json:"evidenceType,omitempty"`
	// The system may support regex. https://www.w3.org/TR/xpath-functions/#regex-syntax
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Only match Evidence messages that have any of these external identifiers
	ExternalIdentifiers  []*ExternalIdentifier `protobuf:"bytes,3,rep,name=external_identifiers,json=externalIdentifiers,proto3" json:"external_identifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EvidenceQuery) Reset()         { *m = EvidenceQuery{} }
func (m *EvidenceQuery) String() string { return proto.CompactTextString(m) }
func (*EvidenceQuery) ProtoMessage()    {}
func (*EvidenceQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{4}
}

func (m *EvidenceQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvidenceQuery.Unmarshal(m, b)
}
func (m *EvidenceQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvidenceQuery.Marshal(b, m, deterministic)
}
func (m *EvidenceQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvidenceQuery.Merge(m, src)
}
func (m *EvidenceQuery) XXX_Size() int {
	return xxx_messageInfo_EvidenceQuery.Size(m)
}
func (m *EvidenceQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_EvidenceQuery.DiscardUnknown(m)
}

var xxx_messageInfo_EvidenceQuery proto.InternalMessageInfo

func (m *EvidenceQuery) GetEvidenceType() *OntologyTerm {
	if m != nil {
		return m.EvidenceType
	}
	return nil
}

func (m *EvidenceQuery) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EvidenceQuery) GetExternalIdentifiers() []*ExternalIdentifier {
	if m != nil {
		return m.ExternalIdentifiers
	}
	return nil
}

type SearchPhenotypesRequest struct {
	// The `PhenotypeAssociationSet` to search. Mandatory
	PhenotypeAssociationSetId string `protobuf:"bytes,1,opt,name=phenotype_association_set_id,json=phenotypeAssociationSetId,proto3" json:"phenotype_association_set_id,omitempty"`
	// Phenotype ID
	// TODO remove if a get-by-phenotype ID endpoint is added
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The system may support regex. https://www.w3.org/TR/xpath-functions/#regex-syntax
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Only return results that match this type
	Type *OntologyTerm `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// terms should be OR'd together. e.g. (severe OR abnormal)
	Qualifiers []*OntologyTerm `protobuf:"bytes,5,rep,name=qualifiers,proto3" json:"qualifiers,omitempty"`
	// Only return results that match this age of onset
	AgeOfOnset *OntologyTerm `protobuf:"bytes,6,opt,name=age_of_onset,json=ageOfOnset,proto3" json:"age_of_onset,omitempty"`
	// Specifies the maximum number of results to return in a single page.
	// If unspecified, a system default will be used.
	PageSize int64 `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `next_page_token` from the previous response.
	PageToken            string   `protobuf:"bytes,8,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPhenotypesRequest) Reset()         { *m = SearchPhenotypesRequest{} }
func (m *SearchPhenotypesRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPhenotypesRequest) ProtoMessage()    {}
func (*SearchPhenotypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{5}
}

func (m *SearchPhenotypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPhenotypesRequest.Unmarshal(m, b)
}
func (m *SearchPhenotypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPhenotypesRequest.Marshal(b, m, deterministic)
}
func (m *SearchPhenotypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPhenotypesRequest.Merge(m, src)
}
func (m *SearchPhenotypesRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPhenotypesRequest.Size(m)
}
func (m *SearchPhenotypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPhenotypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPhenotypesRequest proto.InternalMessageInfo

func (m *SearchPhenotypesRequest) GetPhenotypeAssociationSetId() string {
	if m != nil {
		return m.PhenotypeAssociationSetId
	}
	return ""
}

func (m *SearchPhenotypesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SearchPhenotypesRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SearchPhenotypesRequest) GetType() *OntologyTerm {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SearchPhenotypesRequest) GetQualifiers() []*OntologyTerm {
	if m != nil {
		return m.Qualifiers
	}
	return nil
}

func (m *SearchPhenotypesRequest) GetAgeOfOnset() *OntologyTerm {
	if m != nil {
		return m.AgeOfOnset
	}
	return nil
}

func (m *SearchPhenotypesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchPhenotypesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// This is the response from `POST /phenotypes/search` expressed as JSON.
type SearchPhenotypesResponse struct {
	// The list of matching PhenotypeInstances.
	Phenotypes []*PhenotypeInstance `protobuf:"bytes,1,rep,name=phenotypes,proto3" json:"phenotypes,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPhenotypesResponse) Reset()         { *m = SearchPhenotypesResponse{} }
func (m *SearchPhenotypesResponse) String() string { return proto.CompactTextString(m) }
func (*SearchPhenotypesResponse) ProtoMessage()    {}
func (*SearchPhenotypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{6}
}

func (m *SearchPhenotypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPhenotypesResponse.Unmarshal(m, b)
}
func (m *SearchPhenotypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPhenotypesResponse.Marshal(b, m, deterministic)
}
func (m *SearchPhenotypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPhenotypesResponse.Merge(m, src)
}
func (m *SearchPhenotypesResponse) XXX_Size() int {
	return xxx_messageInfo_SearchPhenotypesResponse.Size(m)
}
func (m *SearchPhenotypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPhenotypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPhenotypesResponse proto.InternalMessageInfo

func (m *SearchPhenotypesResponse) GetPhenotypes() []*PhenotypeInstance {
	if m != nil {
		return m.Phenotypes
	}
	return nil
}

func (m *SearchPhenotypesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

//
//This request maps to the body of `POST /featurephenotypeassociations/search` as JSON.
//
//The goal here is to allow users to query using one or more of
//Genotype, Phenotype, Environment, and Evidence.
//
//A query using one of the above items is to mean, by default,
//that the remainder of the query is as a "wildcard", such
//that all matches to just that query term would come back.
//Combinations of the above are to act like AND rather than OR.
//
//The "genotype" part of the query methods can be one or more
//genomic features.  Associations can be made at many
//levels of granularity (from whole genotypes down to individual
//SNVs), but users may use these methods with partial or
//inexact information.  Therefore, the query methods must be
//able to support query of some or all of the associated features.
//Furthermore, use of the relationships between genomic features
//means that when querying for a gene, any variants to that
//gene are also returned.  For example, a query with
//BRCA2 would mean that in addition to any direct associations
//to the BRCA2, all associations to sequence variants of BRCA2 would also
//be returned.  Similarly, queries with OntologyTerms should perform
//the subclass closure.
//
//Each query can be made against a string, an array of external
//identifers (such as for gene or SNP ids), ontology term ids, or
//full feature/phenotype/evidence objects.
type SearchGenotypePhenotypeRequest struct {
	// The `PhenotypeAssociationSet` to search. Mandatory
	PhenotypeAssociationSetId string `protobuf:"bytes,1,opt,name=phenotype_association_set_id,json=phenotypeAssociationSetId,proto3" json:"phenotype_association_set_id,omitempty"`
	// At least one feature_id or phenotype_id must be provided.
	FeatureIds []string `protobuf:"bytes,2,rep,name=feature_ids,json=featureIds,proto3" json:"feature_ids,omitempty"`
	// Phenotype IDs
	PhenotypeIds []string `protobuf:"bytes,3,rep,name=phenotype_ids,json=phenotypeIds,proto3" json:"phenotype_ids,omitempty"`
	// evidence
	Evidence []*EvidenceQuery `protobuf:"bytes,4,rep,name=evidence,proto3" json:"evidence,omitempty"`
	// Specifies the maximum number of results to return in a single page.
	// If unspecified, a system default will be used.
	PageSize int64 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken            string   `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchGenotypePhenotypeRequest) Reset()         { *m = SearchGenotypePhenotypeRequest{} }
func (m *SearchGenotypePhenotypeRequest) String() string { return proto.CompactTextString(m) }
func (*SearchGenotypePhenotypeRequest) ProtoMessage()    {}
func (*SearchGenotypePhenotypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{7}
}

func (m *SearchGenotypePhenotypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGenotypePhenotypeRequest.Unmarshal(m, b)
}
func (m *SearchGenotypePhenotypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGenotypePhenotypeRequest.Marshal(b, m, deterministic)
}
func (m *SearchGenotypePhenotypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGenotypePhenotypeRequest.Merge(m, src)
}
func (m *SearchGenotypePhenotypeRequest) XXX_Size() int {
	return xxx_messageInfo_SearchGenotypePhenotypeRequest.Size(m)
}
func (m *SearchGenotypePhenotypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGenotypePhenotypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGenotypePhenotypeRequest proto.InternalMessageInfo

func (m *SearchGenotypePhenotypeRequest) GetPhenotypeAssociationSetId() string {
	if m != nil {
		return m.PhenotypeAssociationSetId
	}
	return ""
}

func (m *SearchGenotypePhenotypeRequest) GetFeatureIds() []string {
	if m != nil {
		return m.FeatureIds
	}
	return nil
}

func (m *SearchGenotypePhenotypeRequest) GetPhenotypeIds() []string {
	if m != nil {
		return m.PhenotypeIds
	}
	return nil
}

func (m *SearchGenotypePhenotypeRequest) GetEvidence() []*EvidenceQuery {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *SearchGenotypePhenotypeRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchGenotypePhenotypeRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// This is the response from `POST /genotypephenotype/search` expressed as JSON.
type SearchGenotypePhenotypeResponse struct {
	// The list of matching FeaturePhenotypeAssociation.
	Associations []*FeaturePhenotypeAssociation `protobuf:"bytes,1,rep,name=associations,proto3" json:"associations,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchGenotypePhenotypeResponse) Reset()         { *m = SearchGenotypePhenotypeResponse{} }
func (m *SearchGenotypePhenotypeResponse) String() string { return proto.CompactTextString(m) }
func (*SearchGenotypePhenotypeResponse) ProtoMessage()    {}
func (*SearchGenotypePhenotypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a94b51187eeb9944, []int{8}
}

func (m *SearchGenotypePhenotypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGenotypePhenotypeResponse.Unmarshal(m, b)
}
func (m *SearchGenotypePhenotypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGenotypePhenotypeResponse.Marshal(b, m, deterministic)
}
func (m *SearchGenotypePhenotypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGenotypePhenotypeResponse.Merge(m, src)
}
func (m *SearchGenotypePhenotypeResponse) XXX_Size() int {
	return xxx_messageInfo_SearchGenotypePhenotypeResponse.Size(m)
}
func (m *SearchGenotypePhenotypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGenotypePhenotypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGenotypePhenotypeResponse proto.InternalMessageInfo

func (m *SearchGenotypePhenotypeResponse) GetAssociations() []*FeaturePhenotypeAssociation {
	if m != nil {
		return m.Associations
	}
	return nil
}

func (m *SearchGenotypePhenotypeResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchPhenotypeAssociationSetsRequest)(nil), "candig.SearchPhenotypeAssociationSetsRequest")
	proto.RegisterType((*SearchPhenotypeAssociationSetsResponse)(nil), "candig.SearchPhenotypeAssociationSetsResponse")
	proto.RegisterType((*OntologyTermQuery)(nil), "candig.OntologyTermQuery")
	proto.RegisterType((*ExternalIdentifierQuery)(nil), "candig.ExternalIdentifierQuery")
	proto.RegisterType((*EvidenceQuery)(nil), "candig.EvidenceQuery")
	proto.RegisterType((*SearchPhenotypesRequest)(nil), "candig.SearchPhenotypesRequest")
	proto.RegisterType((*SearchPhenotypesResponse)(nil), "candig.SearchPhenotypesResponse")
	proto.RegisterType((*SearchGenotypePhenotypeRequest)(nil), "candig.SearchGenotypePhenotypeRequest")
	proto.RegisterType((*SearchGenotypePhenotypeResponse)(nil), "candig.SearchGenotypePhenotypeResponse")
}

func init() {
	proto.RegisterFile("candig/genotype_phenotype_service.proto", fileDescriptor_a94b51187eeb9944)
}

var fileDescriptor_a94b51187eeb9944 = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x4f, 0xdb, 0x58,
	0x10, 0x96, 0x6d, 0xc8, 0x92, 0x21, 0x2c, 0xda, 0x07, 0x2b, 0x4c, 0x60, 0x49, 0x64, 0x76, 0x21,
	0x62, 0xd9, 0x84, 0x65, 0x57, 0x88, 0xdd, 0x0b, 0xea, 0x81, 0xa2, 0x1c, 0x2a, 0xa8, 0xc3, 0xb5,
	0xb2, 0x5c, 0x7b, 0x12, 0x9e, 0x9a, 0xf8, 0x19, 0xbf, 0x17, 0x44, 0x90, 0x7a, 0xe1, 0x27, 0xb4,
	0x52, 0x7f, 0x42, 0xaf, 0x3d, 0xf4, 0xd6, 0xbf, 0xd1, 0x3f, 0xd0, 0x43, 0xff, 0x42, 0xef, 0x95,
	0xed, 0x67, 0xc7, 0x49, 0x70, 0xc8, 0xa1, 0xb7, 0x68, 0xfc, 0xcd, 0xcc, 0x37, 0xf3, 0xcd, 0x67,
	0x07, 0x76, 0x1d, 0xdb, 0x73, 0x69, 0xa7, 0xd1, 0x41, 0x8f, 0x89, 0x81, 0x8f, 0x96, 0x7f, 0x95,
	0xfc, 0xe2, 0x18, 0xdc, 0x50, 0x07, 0xeb, 0x7e, 0xc0, 0x04, 0x23, 0x85, 0x18, 0x58, 0x5e, 0x91,
	0x09, 0x0e, 0xeb, 0xf5, 0x98, 0x17, 0x3f, 0x2c, 0x57, 0x72, 0xab, 0x48, 0xc0, 0x66, 0x87, 0xb1,
	0x4e, 0x17, 0x1b, 0xb6, 0x4f, 0x1b, 0xb6, 0xe7, 0x31, 0x61, 0x0b, 0xca, 0x3c, 0x1e, 0x3f, 0x35,
	0xee, 0x15, 0xf8, 0xa3, 0x85, 0x76, 0xe0, 0x5c, 0x5d, 0x24, 0x79, 0x4f, 0x38, 0x67, 0x0e, 0x8d,
	0x50, 0x2d, 0x14, 0xdc, 0xc4, 0xeb, 0x3e, 0x72, 0x41, 0x7e, 0x03, 0x70, 0x6d, 0x61, 0x73, 0x14,
	0x16, 0x75, 0x75, 0xa5, 0xaa, 0xd4, 0x8a, 0x66, 0x51, 0x46, 0x9a, 0x2e, 0xd9, 0x80, 0xa2, 0x6f,
	0x77, 0xd0, 0xe2, 0xf4, 0x0e, 0x75, 0xb5, 0xaa, 0xd4, 0x34, 0x73, 0x21, 0x0c, 0xb4, 0xe8, 0x1d,
	0x86, 0xb9, 0xd1, 0x43, 0xc1, 0x5e, 0xa1, 0xa7, 0x6b, 0x71, 0x6e, 0x18, 0xb9, 0x0c, 0x03, 0xc6,
	0x07, 0x05, 0x76, 0x1e, 0x23, 0xc1, 0x7d, 0xe6, 0x71, 0x24, 0x2f, 0xa0, 0x3c, 0x5c, 0x93, 0x3d,
	0x04, 0x59, 0x1c, 0x05, 0xd7, 0x95, 0xaa, 0x56, 0x5b, 0x3c, 0xac, 0xd4, 0xe3, 0x9d, 0xd4, 0x73,
	0xaa, 0x99, 0xba, 0x9f, 0xd3, 0x86, 0xec, 0xc0, 0xb2, 0x87, 0xb7, 0xc2, 0xca, 0xb0, 0x55, 0x23,
	0xb6, 0x4b, 0x61, 0xf8, 0x22, 0x65, 0x7c, 0x02, 0xbf, 0x9c, 0x7b, 0x82, 0x75, 0x59, 0x67, 0x70,
	0x89, 0x41, 0xef, 0x79, 0x1f, 0x83, 0x01, 0xd9, 0x83, 0x79, 0x81, 0x41, 0x2f, 0xa1, 0xb1, 0x9a,
	0xd0, 0xc8, 0x22, 0xcd, 0x18, 0x62, 0x9c, 0xc1, 0xda, 0xe9, 0xad, 0xc0, 0xc0, 0xb3, 0xbb, 0x4d,
	0x17, 0x3d, 0x41, 0xdb, 0x14, 0x83, 0xb8, 0xcc, 0x3e, 0x68, 0xd4, 0x4d, 0x8a, 0x94, 0x93, 0x22,
	0x93, 0x68, 0x33, 0x84, 0x19, 0x9f, 0x14, 0x58, 0x3a, 0xbd, 0xa1, 0x2e, 0x7a, 0x0e, 0xc6, 0xf9,
	0xc7, 0x50, 0x42, 0x19, 0xb8, 0x1c, 0xf8, 0x18, 0x49, 0x95, 0xc7, 0x66, 0x04, 0x49, 0xaa, 0xb0,
	0xe8, 0x22, 0x77, 0x02, 0xea, 0x87, 0x0b, 0x91, 0x93, 0x67, 0x43, 0xe4, 0x19, 0xac, 0xa2, 0x24,
	0x62, 0xd1, 0x94, 0x09, 0xd7, 0xb5, 0x47, 0xc9, 0xae, 0xe0, 0x44, 0x8c, 0x1b, 0x5f, 0x54, 0x58,
	0x1b, 0x13, 0x3e, 0xbd, 0xb7, 0x13, 0xd8, 0xcc, 0x55, 0x7a, 0x78, 0x81, 0xeb, 0x39, 0x52, 0x36,
	0x5d, 0xf2, 0x33, 0xa8, 0xd4, 0x95, 0x43, 0xa8, 0xd4, 0x1d, 0x9f, 0x4e, 0x9b, 0x9c, 0xae, 0x06,
	0x73, 0x61, 0x25, 0x7d, 0x6e, 0xca, 0xc6, 0x22, 0x04, 0xf9, 0x17, 0xe0, 0xba, 0x6f, 0x77, 0xe5,
	0xf4, 0xf3, 0x53, 0xf4, 0xce, 0xe0, 0xc8, 0x11, 0x94, 0xc2, 0xbb, 0x62, 0x6d, 0x2b, 0xbc, 0x65,
	0xa1, 0x17, 0xa6, 0xf4, 0x01, 0xbb, 0x83, 0xe7, 0xed, 0xf3, 0x10, 0x37, 0xea, 0xad, 0x9f, 0xa6,
	0x7a, 0x6b, 0x61, 0xdc, 0x5b, 0xaf, 0x41, 0x9f, 0xdc, 0xb0, 0x34, 0xd3, 0x7f, 0x00, 0xe9, 0xfa,
	0x92, 0x83, 0x5b, 0x9f, 0x30, 0x4f, 0xd3, 0xe3, 0xc2, 0xf6, 0x1c, 0x34, 0x33, 0xe0, 0x99, 0x8d,
	0xf2, 0x4e, 0x85, 0xad, 0xb8, 0xff, 0x99, 0xcc, 0x4d, 0xeb, 0xfe, 0x30, 0xa1, 0x2b, 0xb0, 0xd8,
	0x46, 0x5b, 0xf4, 0x03, 0xb4, 0x42, 0xe3, 0xa8, 0x55, 0xad, 0x56, 0x34, 0x41, 0x86, 0x9a, 0x2e,
	0x27, 0xdb, 0xb0, 0x34, 0xec, 0x10, 0x42, 0xb4, 0x08, 0x52, 0x4a, 0x83, 0x21, 0xe8, 0x6f, 0x58,
	0x48, 0xcc, 0xa0, 0xcf, 0x45, 0xab, 0xf8, 0x35, 0x3d, 0xe7, 0xac, 0xbf, 0xcc, 0x14, 0x36, 0xaa,
	0xcb, 0xfc, 0x54, 0x5d, 0x0a, 0xe3, 0xba, 0xbc, 0x51, 0xa0, 0x92, 0xbb, 0x18, 0xa9, 0xcf, 0x19,
	0x94, 0x32, 0xfb, 0x48, 0x14, 0xda, 0x4e, 0x68, 0x3d, 0x8d, 0x27, 0x7c, 0xe8, 0x2d, 0x67, 0x8e,
	0x24, 0xce, 0xaa, 0xd6, 0xe1, 0x37, 0x0d, 0xf4, 0x09, 0x3a, 0xad, 0xf8, 0x63, 0x44, 0x3e, 0x2a,
	0x89, 0x94, 0x79, 0x6f, 0x69, 0xf2, 0x57, 0x42, 0x6d, 0xa6, 0x4f, 0x4a, 0xb9, 0x3e, 0x2b, 0x3c,
	0xde, 0x87, 0x71, 0x78, 0xff, 0xf9, 0xeb, 0x5b, 0x75, 0xdf, 0xd8, 0x6d, 0xdc, 0x1c, 0xd4, 0x8f,
	0xeb, 0x07, 0x8d, 0x54, 0xc0, 0xec, 0xb4, 0x28, 0x78, 0x83, 0x47, 0x05, 0xff, 0x57, 0xf6, 0xc8,
	0x1d, 0x2c, 0x8f, 0x55, 0x27, 0x95, 0x9c, 0xb6, 0x29, 0xaf, 0x6a, 0x3e, 0x40, 0x32, 0xf9, 0x3d,
	0x62, 0xb2, 0x65, 0xac, 0x4f, 0x30, 0xc9, 0xf6, 0x7e, 0xaf, 0xc0, 0x46, 0xfe, 0x68, 0x9c, 0xec,
	0x8c, 0xf6, 0xc9, 0x73, 0x48, 0x79, 0xf7, 0x51, 0x9c, 0xa4, 0x75, 0x14, 0xd1, 0x3a, 0x30, 0xfe,
	0x4c, 0x68, 0x49, 0x13, 0x3c, 0xb8, 0xa7, 0x21, 0xd1, 0x97, 0x85, 0xe8, 0xcf, 0xc0, 0x3f, 0xdf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x67, 0x65, 0x6d, 0x93, 0x08, 0x00, 0x00,
}
