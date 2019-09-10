# Ga4ghFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of this annotation node. | [optional] 
**Name** | Pointer to **string** | An optional name to provide for the feature. | [optional] 
**GeneSymbol** | Pointer to **string** | The gene symbol the feature occurs on. This field may be replaced with a more generic representation in a future version. | [optional] 
**ParentId** | Pointer to **string** | Parent Id of this node. Set to empty string if node has no parent. | [optional] 
**ChildIds** | Pointer to **[]string** | Ordered array of Child Ids of this node. Since not all child nodes are ordered by genomic coordinates, this can&#39;t always be reconstructed from parent_id&#39;s of the children alone. | [optional] 
**FeatureSetId** | Pointer to **string** | Identifier for the containing feature set. | [optional] 
**ReferenceName** | Pointer to **string** | The reference on which this feature occurs (e.g. &#x60;chr20&#x60; or &#x60;X&#x60;). | [optional] 
**Start** | Pointer to **string** | The start position at which this feature occurs (0-based). This corresponds to the first base of the string of reference bases. Genomic positions are non-negative integers less than reference length. Features spanning the join of circular genomes are represented as two features one on each side of the join (position 0). | [optional] 
**End** | Pointer to **string** | The end position (exclusive), resulting in [start, end) closed-open interval. This is typically calculated by &#x60;start + reference_bases.length&#x60;. | [optional] 
**Strand** | Pointer to [**Ga4ghStrand**](ga4ghStrand.md) |  | [optional] 
**FeatureType** | Pointer to [**Ga4ghOntologyTerm**](ga4ghOntologyTerm.md) |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghFeature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghFeature) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghFeature) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghFeature) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghFeature) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghFeature) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetGeneSymbol

`func (o *Ga4ghFeature) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *Ga4ghFeature) GetGeneSymbolOk() (string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeneSymbol

`func (o *Ga4ghFeature) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbol

`func (o *Ga4ghFeature) SetGeneSymbol(v string)`

SetGeneSymbol gets a reference to the given string and assigns it to the GeneSymbol field.

### GetParentId

`func (o *Ga4ghFeature) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Ga4ghFeature) GetParentIdOk() (string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *Ga4ghFeature) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *Ga4ghFeature) SetParentId(v string)`

SetParentId gets a reference to the given string and assigns it to the ParentId field.

### GetChildIds

`func (o *Ga4ghFeature) GetChildIds() []string`

GetChildIds returns the ChildIds field if non-nil, zero value otherwise.

### GetChildIdsOk

`func (o *Ga4ghFeature) GetChildIdsOk() ([]string, bool)`

GetChildIdsOk returns a tuple with the ChildIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildIds

`func (o *Ga4ghFeature) HasChildIds() bool`

HasChildIds returns a boolean if a field has been set.

### SetChildIds

`func (o *Ga4ghFeature) SetChildIds(v []string)`

SetChildIds gets a reference to the given []string and assigns it to the ChildIds field.

### GetFeatureSetId

`func (o *Ga4ghFeature) GetFeatureSetId() string`

GetFeatureSetId returns the FeatureSetId field if non-nil, zero value otherwise.

### GetFeatureSetIdOk

`func (o *Ga4ghFeature) GetFeatureSetIdOk() (string, bool)`

GetFeatureSetIdOk returns a tuple with the FeatureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureSetId

`func (o *Ga4ghFeature) HasFeatureSetId() bool`

HasFeatureSetId returns a boolean if a field has been set.

### SetFeatureSetId

`func (o *Ga4ghFeature) SetFeatureSetId(v string)`

SetFeatureSetId gets a reference to the given string and assigns it to the FeatureSetId field.

### GetReferenceName

`func (o *Ga4ghFeature) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghFeature) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghFeature) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghFeature) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetStart

`func (o *Ga4ghFeature) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghFeature) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghFeature) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghFeature) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghFeature) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghFeature) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghFeature) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghFeature) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetStrand

`func (o *Ga4ghFeature) GetStrand() Ga4ghStrand`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *Ga4ghFeature) GetStrandOk() (Ga4ghStrand, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStrand

`func (o *Ga4ghFeature) HasStrand() bool`

HasStrand returns a boolean if a field has been set.

### SetStrand

`func (o *Ga4ghFeature) SetStrand(v Ga4ghStrand)`

SetStrand gets a reference to the given Ga4ghStrand and assigns it to the Strand field.

### GetFeatureType

`func (o *Ga4ghFeature) GetFeatureType() Ga4ghOntologyTerm`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *Ga4ghFeature) GetFeatureTypeOk() (Ga4ghOntologyTerm, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureType

`func (o *Ga4ghFeature) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### SetFeatureType

`func (o *Ga4ghFeature) SetFeatureType(v Ga4ghOntologyTerm)`

SetFeatureType gets a reference to the given Ga4ghOntologyTerm and assigns it to the FeatureType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


