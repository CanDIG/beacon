# Ga4ghCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallSetName** | Pointer to **string** | The name of the call set this variant call belongs to. If this field is not present, the ordering of the call sets from a &#x60;SearchCallSetsRequest&#x60; over this &#x60;VariantSet&#x60; is guaranteed to match the ordering of the calls on this &#x60;Variant&#x60;. The number of results will also be the same. | [optional] 
**CallSetId** | Pointer to **string** | The ID of the call set this variant call belongs to.  If this field is not present, the ordering of the call sets from a &#x60;SearchCallSetsRequest&#x60; over this &#x60;VariantSet&#x60; is guaranteed to match the ordering of the calls on this &#x60;Variant&#x60;. The number of results will also be the same. | [optional] 
**Genotype** | Pointer to [**ProtobufListValue**](protobufListValue.md) |  | [optional] 
**Phaseset** | Pointer to **string** | If this field is populated, this variant call&#39;s genotype ordering implies the phase of the bases and is consistent with any other variant calls on the same contig which have the same phaseset string. | [optional] 
**GenotypeLikelihood** | Pointer to **[]float64** | The genotype likelihoods for this variant call. Each array entry represents how likely a specific genotype is for this call as log10(P(data | genotype)), analogous to the GL tag in the VCF spec. The value ordering is defined by the GL tag in the VCF spec. | [optional] 
**Attributes** | Pointer to [**Ga4ghAttributes**](ga4ghAttributes.md) |  | [optional] 

## Methods

### GetCallSetName

`func (o *Ga4ghCall) GetCallSetName() string`

GetCallSetName returns the CallSetName field if non-nil, zero value otherwise.

### GetCallSetNameOk

`func (o *Ga4ghCall) GetCallSetNameOk() (string, bool)`

GetCallSetNameOk returns a tuple with the CallSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallSetName

`func (o *Ga4ghCall) HasCallSetName() bool`

HasCallSetName returns a boolean if a field has been set.

### SetCallSetName

`func (o *Ga4ghCall) SetCallSetName(v string)`

SetCallSetName gets a reference to the given string and assigns it to the CallSetName field.

### GetCallSetId

`func (o *Ga4ghCall) GetCallSetId() string`

GetCallSetId returns the CallSetId field if non-nil, zero value otherwise.

### GetCallSetIdOk

`func (o *Ga4ghCall) GetCallSetIdOk() (string, bool)`

GetCallSetIdOk returns a tuple with the CallSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallSetId

`func (o *Ga4ghCall) HasCallSetId() bool`

HasCallSetId returns a boolean if a field has been set.

### SetCallSetId

`func (o *Ga4ghCall) SetCallSetId(v string)`

SetCallSetId gets a reference to the given string and assigns it to the CallSetId field.

### GetGenotype

`func (o *Ga4ghCall) GetGenotype() ProtobufListValue`

GetGenotype returns the Genotype field if non-nil, zero value otherwise.

### GetGenotypeOk

`func (o *Ga4ghCall) GetGenotypeOk() (ProtobufListValue, bool)`

GetGenotypeOk returns a tuple with the Genotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGenotype

`func (o *Ga4ghCall) HasGenotype() bool`

HasGenotype returns a boolean if a field has been set.

### SetGenotype

`func (o *Ga4ghCall) SetGenotype(v ProtobufListValue)`

SetGenotype gets a reference to the given ProtobufListValue and assigns it to the Genotype field.

### GetPhaseset

`func (o *Ga4ghCall) GetPhaseset() string`

GetPhaseset returns the Phaseset field if non-nil, zero value otherwise.

### GetPhasesetOk

`func (o *Ga4ghCall) GetPhasesetOk() (string, bool)`

GetPhasesetOk returns a tuple with the Phaseset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhaseset

`func (o *Ga4ghCall) HasPhaseset() bool`

HasPhaseset returns a boolean if a field has been set.

### SetPhaseset

`func (o *Ga4ghCall) SetPhaseset(v string)`

SetPhaseset gets a reference to the given string and assigns it to the Phaseset field.

### GetGenotypeLikelihood

`func (o *Ga4ghCall) GetGenotypeLikelihood() []float64`

GetGenotypeLikelihood returns the GenotypeLikelihood field if non-nil, zero value otherwise.

### GetGenotypeLikelihoodOk

`func (o *Ga4ghCall) GetGenotypeLikelihoodOk() ([]float64, bool)`

GetGenotypeLikelihoodOk returns a tuple with the GenotypeLikelihood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGenotypeLikelihood

`func (o *Ga4ghCall) HasGenotypeLikelihood() bool`

HasGenotypeLikelihood returns a boolean if a field has been set.

### SetGenotypeLikelihood

`func (o *Ga4ghCall) SetGenotypeLikelihood(v []float64)`

SetGenotypeLikelihood gets a reference to the given []float64 and assigns it to the GenotypeLikelihood field.

### GetAttributes

`func (o *Ga4ghCall) GetAttributes() Ga4ghAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Ga4ghCall) GetAttributesOk() (Ga4ghAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Ga4ghCall) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Ga4ghCall) SetAttributes(v Ga4ghAttributes)`

SetAttributes gets a reference to the given Ga4ghAttributes and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


