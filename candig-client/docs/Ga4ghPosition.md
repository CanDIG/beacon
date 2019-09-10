# Ga4ghPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceName** | Pointer to **string** | The name of the &#x60;Reference&#x60; on which the &#x60;Position&#x60; is located. | [optional] 
**Position** | Pointer to **string** | The 0-based offset from the start of the forward strand for that &#x60;Reference&#x60;. Genomic positions are non-negative integers less than &#x60;Reference&#x60; length. | [optional] 
**Strand** | Pointer to [**Ga4ghStrand**](ga4ghStrand.md) |  | [optional] 

## Methods

### GetReferenceName

`func (o *Ga4ghPosition) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghPosition) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghPosition) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghPosition) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetPosition

`func (o *Ga4ghPosition) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Ga4ghPosition) GetPositionOk() (string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *Ga4ghPosition) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *Ga4ghPosition) SetPosition(v string)`

SetPosition gets a reference to the given string and assigns it to the Position field.

### GetStrand

`func (o *Ga4ghPosition) GetStrand() Ga4ghStrand`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *Ga4ghPosition) GetStrandOk() (Ga4ghStrand, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStrand

`func (o *Ga4ghPosition) HasStrand() bool`

HasStrand returns a boolean if a field has been set.

### SetStrand

`func (o *Ga4ghPosition) SetStrand(v Ga4ghStrand)`

SetStrand gets a reference to the given Ga4ghStrand and assigns it to the Strand field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


