# Ga4ghSearchVariantsByGeneNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** | The &#x60;VariantSet&#x60; to search. | [optional] 
**Gene** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 
**PatientList** | Pointer to **[]string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**ReferenceName** | Pointer to **string** | Reference genome. Used in advanced searchesif the user is requesting variants endpoint. | [optional] 
**CallSetIds** | Pointer to **[]string** | Only return variant calls which belong to call sets with these IDs. If unspecified, return all variants and no variant call objects. | [optional] 

## Methods

### GetDatasetId

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetGene

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetGene() string`

GetGene returns the Gene field if non-nil, zero value otherwise.

### GetGeneOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetGeneOk() (string, bool)`

GetGeneOk returns a tuple with the Gene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGene

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasGene() bool`

HasGene returns a boolean if a field has been set.

### SetGene

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetGene(v string)`

SetGene gets a reference to the given string and assigns it to the Gene field.

### GetPageSize

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.

### GetPatientList

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetPatientList() []string`

GetPatientList returns the PatientList field if non-nil, zero value otherwise.

### GetPatientListOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetPatientListOk() ([]string, bool)`

GetPatientListOk returns a tuple with the PatientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientList

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasPatientList() bool`

HasPatientList returns a boolean if a field has been set.

### SetPatientList

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetPatientList(v []string)`

SetPatientList gets a reference to the given []string and assigns it to the PatientList field.

### GetStart

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetReferenceName

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetCallSetIds

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetCallSetIds() []string`

GetCallSetIds returns the CallSetIds field if non-nil, zero value otherwise.

### GetCallSetIdsOk

`func (o *Ga4ghSearchVariantsByGeneNameRequest) GetCallSetIdsOk() ([]string, bool)`

GetCallSetIdsOk returns a tuple with the CallSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallSetIds

`func (o *Ga4ghSearchVariantsByGeneNameRequest) HasCallSetIds() bool`

HasCallSetIds returns a boolean if a field has been set.

### SetCallSetIds

`func (o *Ga4ghSearchVariantsByGeneNameRequest) SetCallSetIds(v []string)`

SetCallSetIds gets a reference to the given []string and assigns it to the CallSetIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


