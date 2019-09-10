# Ga4ghLinearAlignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to [**Ga4ghPosition**](ga4ghPosition.md) |  | [optional] 
**MappingQuality** | Pointer to **int32** | The mapping quality of this alignment, meaning the likelihood that the read maps to this position.  Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to the nearest integer. | [optional] 
**Cigar** | Pointer to [**[]Ga4ghCigarUnit**](ga4ghCigarUnit.md) | Represents the local alignment of this sequence (alignment matches, indels, etc) versus the reference. | [optional] 

## Methods

### GetPosition

`func (o *Ga4ghLinearAlignment) GetPosition() Ga4ghPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Ga4ghLinearAlignment) GetPositionOk() (Ga4ghPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *Ga4ghLinearAlignment) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *Ga4ghLinearAlignment) SetPosition(v Ga4ghPosition)`

SetPosition gets a reference to the given Ga4ghPosition and assigns it to the Position field.

### GetMappingQuality

`func (o *Ga4ghLinearAlignment) GetMappingQuality() int32`

GetMappingQuality returns the MappingQuality field if non-nil, zero value otherwise.

### GetMappingQualityOk

`func (o *Ga4ghLinearAlignment) GetMappingQualityOk() (int32, bool)`

GetMappingQualityOk returns a tuple with the MappingQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMappingQuality

`func (o *Ga4ghLinearAlignment) HasMappingQuality() bool`

HasMappingQuality returns a boolean if a field has been set.

### SetMappingQuality

`func (o *Ga4ghLinearAlignment) SetMappingQuality(v int32)`

SetMappingQuality gets a reference to the given int32 and assigns it to the MappingQuality field.

### GetCigar

`func (o *Ga4ghLinearAlignment) GetCigar() []Ga4ghCigarUnit`

GetCigar returns the Cigar field if non-nil, zero value otherwise.

### GetCigarOk

`func (o *Ga4ghLinearAlignment) GetCigarOk() ([]Ga4ghCigarUnit, bool)`

GetCigarOk returns a tuple with the Cigar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCigar

`func (o *Ga4ghLinearAlignment) HasCigar() bool`

HasCigar returns a boolean if a field has been set.

### SetCigar

`func (o *Ga4ghLinearAlignment) SetCigar(v []Ga4ghCigarUnit)`

SetCigar gets a reference to the given []Ga4ghCigarUnit and assigns it to the Cigar field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


