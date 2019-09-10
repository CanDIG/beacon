# Ga4ghSearchFeaturesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureSetId** | Pointer to **string** | The annotation set to search within. Either &#x60;feature_set_id&#x60; or &#x60;parent_id&#x60; must be non-empty. | [optional] 
**Name** | Pointer to **string** | Only returns features with this name (case-sensitive, exact match). | [optional] 
**GeneSymbol** | Pointer to **string** | Only return features with matching the provided gene symbol (case-sensitive, exact match). This field may be replaced with a more generic representation in a future version. | [optional] 
**ParentId** | Pointer to **string** | Restricts the search to direct children of the given parent &#x60;feature&#x60; ID. Either &#x60;feature_set_id&#x60; or &#x60;parent_id&#x60; must be non-empty. | [optional] 
**ReferenceName** | Pointer to **string** | Only return features on the reference with this name (matched to literal reference name as imported from the GFF3). | [optional] 
**Start** | Pointer to **string** | Required, if name or symbol not provided. The beginning of the window (0-based, inclusive) for which overlapping features should be returned.  Genomic positions are non-negative integers less than reference length.  Requests spanning the join of circular genomes are represented as two requests one on each side of the join (position 0). | [optional] 
**End** | Pointer to **string** | Required, if name or symbol not provided. The end of the window (0-based, exclusive) for which overlapping features should be returned. | [optional] 
**FeatureTypes** | Pointer to **[]string** | TODO: To be replaced with a fully featured ontology search once the Metadata definitions are rounded out. If specified, this query matches only annotations whose &#x60;feature_type&#x60; matches one of the provided ontology terms. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 

## Methods

### GetFeatureSetId

`func (o *Ga4ghSearchFeaturesRequest) GetFeatureSetId() string`

GetFeatureSetId returns the FeatureSetId field if non-nil, zero value otherwise.

### GetFeatureSetIdOk

`func (o *Ga4ghSearchFeaturesRequest) GetFeatureSetIdOk() (string, bool)`

GetFeatureSetIdOk returns a tuple with the FeatureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureSetId

`func (o *Ga4ghSearchFeaturesRequest) HasFeatureSetId() bool`

HasFeatureSetId returns a boolean if a field has been set.

### SetFeatureSetId

`func (o *Ga4ghSearchFeaturesRequest) SetFeatureSetId(v string)`

SetFeatureSetId gets a reference to the given string and assigns it to the FeatureSetId field.

### GetName

`func (o *Ga4ghSearchFeaturesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghSearchFeaturesRequest) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghSearchFeaturesRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghSearchFeaturesRequest) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetGeneSymbol

`func (o *Ga4ghSearchFeaturesRequest) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *Ga4ghSearchFeaturesRequest) GetGeneSymbolOk() (string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeneSymbol

`func (o *Ga4ghSearchFeaturesRequest) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbol

`func (o *Ga4ghSearchFeaturesRequest) SetGeneSymbol(v string)`

SetGeneSymbol gets a reference to the given string and assigns it to the GeneSymbol field.

### GetParentId

`func (o *Ga4ghSearchFeaturesRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Ga4ghSearchFeaturesRequest) GetParentIdOk() (string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *Ga4ghSearchFeaturesRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *Ga4ghSearchFeaturesRequest) SetParentId(v string)`

SetParentId gets a reference to the given string and assigns it to the ParentId field.

### GetReferenceName

`func (o *Ga4ghSearchFeaturesRequest) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghSearchFeaturesRequest) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghSearchFeaturesRequest) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghSearchFeaturesRequest) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetStart

`func (o *Ga4ghSearchFeaturesRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghSearchFeaturesRequest) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghSearchFeaturesRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghSearchFeaturesRequest) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghSearchFeaturesRequest) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghSearchFeaturesRequest) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghSearchFeaturesRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghSearchFeaturesRequest) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetFeatureTypes

`func (o *Ga4ghSearchFeaturesRequest) GetFeatureTypes() []string`

GetFeatureTypes returns the FeatureTypes field if non-nil, zero value otherwise.

### GetFeatureTypesOk

`func (o *Ga4ghSearchFeaturesRequest) GetFeatureTypesOk() ([]string, bool)`

GetFeatureTypesOk returns a tuple with the FeatureTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureTypes

`func (o *Ga4ghSearchFeaturesRequest) HasFeatureTypes() bool`

HasFeatureTypes returns a boolean if a field has been set.

### SetFeatureTypes

`func (o *Ga4ghSearchFeaturesRequest) SetFeatureTypes(v []string)`

SetFeatureTypes gets a reference to the given []string and assigns it to the FeatureTypes field.

### GetPageSize

`func (o *Ga4ghSearchFeaturesRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchFeaturesRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchFeaturesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchFeaturesRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchFeaturesRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchFeaturesRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchFeaturesRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchFeaturesRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


