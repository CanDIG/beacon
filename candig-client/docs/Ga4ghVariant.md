# Ga4ghVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The variant ID. | [optional] 
**VariantSetId** | Pointer to **string** | The ID of the &#x60;VariantSet&#x60; this variant belongs to. This transitively defines the &#x60;ReferenceSet&#x60; against which the &#x60;Variant&#x60; is to be interpreted. | [optional] 
**Names** | Pointer to **[]string** | Names for the variant, for example a RefSNP ID. | [optional] 
**Created** | Pointer to **string** | The date this variant was created in milliseconds from the epoch. | [optional] 
**Updated** | Pointer to **string** | The time at which this variant was last updated in milliseconds from the epoch. | [optional] 
**ReferenceName** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** | The start position at which this variant occurs (0-based). This corresponds to the first base of the string of reference bases. Genomic positions are non-negative integers less than reference length. Variants spanning the join of circular genomes are represented as two variants one on each side of the join (position 0). | [optional] 
**End** | Pointer to **string** | The end position (exclusive), resulting in [start, end) closed-open interval. This is typically calculated by &#x60;start + referenceBases.length&#x60;. | [optional] 
**ReferenceBases** | Pointer to **string** | The reference bases for this variant. They start at the given start position. | [optional] 
**AlternateBases** | Pointer to **[]string** | The bases that appear instead of the reference bases. Multiple alternate alleles are possible. | [optional] 
**Attributes** | Pointer to [**Ga4ghAttributes**](ga4ghAttributes.md) |  | [optional] 
**Calls** | Pointer to [**[]Ga4ghCall**](ga4ghCall.md) | The variant calls for this particular variant. Each one represents the determination of genotype with respect to this variant. &#x60;Call&#x60;s in this array are implicitly associated with this &#x60;Variant&#x60;. | [optional] 
**VariantType** | Pointer to **string** |  | [optional] 
**Svlen** | Pointer to **string** |  | [optional] 
**Cipos** | Pointer to **[]int32** |  | [optional] 
**Ciend** | Pointer to **[]int32** |  | [optional] 
**FiltersApplied** | Pointer to **bool** | True if filters were applied for this variant. VCF column 7 \&quot;FILTER\&quot; any value other than the missing value. | [optional] 
**FiltersPassed** | Pointer to **bool** | True if all filters for this variant passed. VCF column 7 \&quot;FILTER\&quot; value PASS. | [optional] 
**FiltersFailed** | Pointer to **[]string** | Zero or more filters that failed for this variant. VCF column 7 \&quot;FILTER\&quot; shared across all alleles in the same VCF record. | [optional] 
**PatientId** | Pointer to **string** | Patient_Id, this field is not populated in the database, but generated on-the-fly when variants are being returned. | [optional] 

## Methods

### GetId

`func (o *Ga4ghVariant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghVariant) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghVariant) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghVariant) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetVariantSetId

`func (o *Ga4ghVariant) GetVariantSetId() string`

GetVariantSetId returns the VariantSetId field if non-nil, zero value otherwise.

### GetVariantSetIdOk

`func (o *Ga4ghVariant) GetVariantSetIdOk() (string, bool)`

GetVariantSetIdOk returns a tuple with the VariantSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantSetId

`func (o *Ga4ghVariant) HasVariantSetId() bool`

HasVariantSetId returns a boolean if a field has been set.

### SetVariantSetId

`func (o *Ga4ghVariant) SetVariantSetId(v string)`

SetVariantSetId gets a reference to the given string and assigns it to the VariantSetId field.

### GetNames

`func (o *Ga4ghVariant) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Ga4ghVariant) GetNamesOk() ([]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *Ga4ghVariant) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *Ga4ghVariant) SetNames(v []string)`

SetNames gets a reference to the given []string and assigns it to the Names field.

### GetCreated

`func (o *Ga4ghVariant) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghVariant) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghVariant) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghVariant) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghVariant) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghVariant) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghVariant) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghVariant) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetReferenceName

`func (o *Ga4ghVariant) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghVariant) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghVariant) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghVariant) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetStart

`func (o *Ga4ghVariant) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghVariant) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghVariant) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghVariant) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghVariant) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghVariant) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghVariant) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghVariant) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetReferenceBases

`func (o *Ga4ghVariant) GetReferenceBases() string`

GetReferenceBases returns the ReferenceBases field if non-nil, zero value otherwise.

### GetReferenceBasesOk

`func (o *Ga4ghVariant) GetReferenceBasesOk() (string, bool)`

GetReferenceBasesOk returns a tuple with the ReferenceBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceBases

`func (o *Ga4ghVariant) HasReferenceBases() bool`

HasReferenceBases returns a boolean if a field has been set.

### SetReferenceBases

`func (o *Ga4ghVariant) SetReferenceBases(v string)`

SetReferenceBases gets a reference to the given string and assigns it to the ReferenceBases field.

### GetAlternateBases

`func (o *Ga4ghVariant) GetAlternateBases() []string`

GetAlternateBases returns the AlternateBases field if non-nil, zero value otherwise.

### GetAlternateBasesOk

`func (o *Ga4ghVariant) GetAlternateBasesOk() ([]string, bool)`

GetAlternateBasesOk returns a tuple with the AlternateBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlternateBases

`func (o *Ga4ghVariant) HasAlternateBases() bool`

HasAlternateBases returns a boolean if a field has been set.

### SetAlternateBases

`func (o *Ga4ghVariant) SetAlternateBases(v []string)`

SetAlternateBases gets a reference to the given []string and assigns it to the AlternateBases field.

### GetAttributes

`func (o *Ga4ghVariant) GetAttributes() Ga4ghAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Ga4ghVariant) GetAttributesOk() (Ga4ghAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Ga4ghVariant) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Ga4ghVariant) SetAttributes(v Ga4ghAttributes)`

SetAttributes gets a reference to the given Ga4ghAttributes and assigns it to the Attributes field.

### GetCalls

`func (o *Ga4ghVariant) GetCalls() []Ga4ghCall`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *Ga4ghVariant) GetCallsOk() ([]Ga4ghCall, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalls

`func (o *Ga4ghVariant) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### SetCalls

`func (o *Ga4ghVariant) SetCalls(v []Ga4ghCall)`

SetCalls gets a reference to the given []Ga4ghCall and assigns it to the Calls field.

### GetVariantType

`func (o *Ga4ghVariant) GetVariantType() string`

GetVariantType returns the VariantType field if non-nil, zero value otherwise.

### GetVariantTypeOk

`func (o *Ga4ghVariant) GetVariantTypeOk() (string, bool)`

GetVariantTypeOk returns a tuple with the VariantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantType

`func (o *Ga4ghVariant) HasVariantType() bool`

HasVariantType returns a boolean if a field has been set.

### SetVariantType

`func (o *Ga4ghVariant) SetVariantType(v string)`

SetVariantType gets a reference to the given string and assigns it to the VariantType field.

### GetSvlen

`func (o *Ga4ghVariant) GetSvlen() string`

GetSvlen returns the Svlen field if non-nil, zero value otherwise.

### GetSvlenOk

`func (o *Ga4ghVariant) GetSvlenOk() (string, bool)`

GetSvlenOk returns a tuple with the Svlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSvlen

`func (o *Ga4ghVariant) HasSvlen() bool`

HasSvlen returns a boolean if a field has been set.

### SetSvlen

`func (o *Ga4ghVariant) SetSvlen(v string)`

SetSvlen gets a reference to the given string and assigns it to the Svlen field.

### GetCipos

`func (o *Ga4ghVariant) GetCipos() []int32`

GetCipos returns the Cipos field if non-nil, zero value otherwise.

### GetCiposOk

`func (o *Ga4ghVariant) GetCiposOk() ([]int32, bool)`

GetCiposOk returns a tuple with the Cipos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCipos

`func (o *Ga4ghVariant) HasCipos() bool`

HasCipos returns a boolean if a field has been set.

### SetCipos

`func (o *Ga4ghVariant) SetCipos(v []int32)`

SetCipos gets a reference to the given []int32 and assigns it to the Cipos field.

### GetCiend

`func (o *Ga4ghVariant) GetCiend() []int32`

GetCiend returns the Ciend field if non-nil, zero value otherwise.

### GetCiendOk

`func (o *Ga4ghVariant) GetCiendOk() ([]int32, bool)`

GetCiendOk returns a tuple with the Ciend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCiend

`func (o *Ga4ghVariant) HasCiend() bool`

HasCiend returns a boolean if a field has been set.

### SetCiend

`func (o *Ga4ghVariant) SetCiend(v []int32)`

SetCiend gets a reference to the given []int32 and assigns it to the Ciend field.

### GetFiltersApplied

`func (o *Ga4ghVariant) GetFiltersApplied() bool`

GetFiltersApplied returns the FiltersApplied field if non-nil, zero value otherwise.

### GetFiltersAppliedOk

`func (o *Ga4ghVariant) GetFiltersAppliedOk() (bool, bool)`

GetFiltersAppliedOk returns a tuple with the FiltersApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFiltersApplied

`func (o *Ga4ghVariant) HasFiltersApplied() bool`

HasFiltersApplied returns a boolean if a field has been set.

### SetFiltersApplied

`func (o *Ga4ghVariant) SetFiltersApplied(v bool)`

SetFiltersApplied gets a reference to the given bool and assigns it to the FiltersApplied field.

### GetFiltersPassed

`func (o *Ga4ghVariant) GetFiltersPassed() bool`

GetFiltersPassed returns the FiltersPassed field if non-nil, zero value otherwise.

### GetFiltersPassedOk

`func (o *Ga4ghVariant) GetFiltersPassedOk() (bool, bool)`

GetFiltersPassedOk returns a tuple with the FiltersPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFiltersPassed

`func (o *Ga4ghVariant) HasFiltersPassed() bool`

HasFiltersPassed returns a boolean if a field has been set.

### SetFiltersPassed

`func (o *Ga4ghVariant) SetFiltersPassed(v bool)`

SetFiltersPassed gets a reference to the given bool and assigns it to the FiltersPassed field.

### GetFiltersFailed

`func (o *Ga4ghVariant) GetFiltersFailed() []string`

GetFiltersFailed returns the FiltersFailed field if non-nil, zero value otherwise.

### GetFiltersFailedOk

`func (o *Ga4ghVariant) GetFiltersFailedOk() ([]string, bool)`

GetFiltersFailedOk returns a tuple with the FiltersFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFiltersFailed

`func (o *Ga4ghVariant) HasFiltersFailed() bool`

HasFiltersFailed returns a boolean if a field has been set.

### SetFiltersFailed

`func (o *Ga4ghVariant) SetFiltersFailed(v []string)`

SetFiltersFailed gets a reference to the given []string and assigns it to the FiltersFailed field.

### GetPatientId

`func (o *Ga4ghVariant) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghVariant) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghVariant) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghVariant) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


