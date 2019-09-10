# Ga4ghResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**ReferenceName** | Pointer to **string** | Reference genome. Used if the user is requesting variants endpoint. | [optional] 
**Gene** | Pointer to **string** | Gene name. Used if the user is requesting variantsByGene endpoint. | [optional] 
**Count** | Pointer to **[]string** | List of fields to aggregate by and return counts. Overrides response format if included. | [optional] 

## Methods

### GetTable

`func (o *Ga4ghResult) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *Ga4ghResult) GetTableOk() (string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTable

`func (o *Ga4ghResult) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTable

`func (o *Ga4ghResult) SetTable(v string)`

SetTable gets a reference to the given string and assigns it to the Table field.

### GetFields

`func (o *Ga4ghResult) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Ga4ghResult) GetFieldsOk() ([]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFields

`func (o *Ga4ghResult) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFields

`func (o *Ga4ghResult) SetFields(v []string)`

SetFields gets a reference to the given []string and assigns it to the Fields field.

### GetStart

`func (o *Ga4ghResult) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghResult) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghResult) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghResult) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghResult) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghResult) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghResult) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghResult) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetReferenceName

`func (o *Ga4ghResult) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghResult) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghResult) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghResult) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetGene

`func (o *Ga4ghResult) GetGene() string`

GetGene returns the Gene field if non-nil, zero value otherwise.

### GetGeneOk

`func (o *Ga4ghResult) GetGeneOk() (string, bool)`

GetGeneOk returns a tuple with the Gene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGene

`func (o *Ga4ghResult) HasGene() bool`

HasGene returns a boolean if a field has been set.

### SetGene

`func (o *Ga4ghResult) SetGene(v string)`

SetGene gets a reference to the given string and assigns it to the Gene field.

### GetCount

`func (o *Ga4ghResult) GetCount() []string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Ga4ghResult) GetCountOk() ([]string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *Ga4ghResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *Ga4ghResult) SetCount(v []string)`

SetCount gets a reference to the given []string and assigns it to the Count field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


