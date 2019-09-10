# Ga4ghCigarUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**CigarUnitOperation**](CigarUnitOperation.md) |  | [optional] 
**OperationLength** | Pointer to **string** | The number of genomic bases that the operation runs for. Required. | [optional] 
**ReferenceSequence** | Pointer to **string** | &#x60;referenceSequence&#x60; is only used at mismatches (&#x60;SEQUENCE_MISMATCH&#x60;) and deletions (&#x60;DELETE&#x60;). Filling this field replaces SAM&#39;s MD tag. If the relevant information is not available, this field is unset. | [optional] 

## Methods

### GetOperation

`func (o *Ga4ghCigarUnit) GetOperation() CigarUnitOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Ga4ghCigarUnit) GetOperationOk() (CigarUnitOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperation

`func (o *Ga4ghCigarUnit) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperation

`func (o *Ga4ghCigarUnit) SetOperation(v CigarUnitOperation)`

SetOperation gets a reference to the given CigarUnitOperation and assigns it to the Operation field.

### GetOperationLength

`func (o *Ga4ghCigarUnit) GetOperationLength() string`

GetOperationLength returns the OperationLength field if non-nil, zero value otherwise.

### GetOperationLengthOk

`func (o *Ga4ghCigarUnit) GetOperationLengthOk() (string, bool)`

GetOperationLengthOk returns a tuple with the OperationLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperationLength

`func (o *Ga4ghCigarUnit) HasOperationLength() bool`

HasOperationLength returns a boolean if a field has been set.

### SetOperationLength

`func (o *Ga4ghCigarUnit) SetOperationLength(v string)`

SetOperationLength gets a reference to the given string and assigns it to the OperationLength field.

### GetReferenceSequence

`func (o *Ga4ghCigarUnit) GetReferenceSequence() string`

GetReferenceSequence returns the ReferenceSequence field if non-nil, zero value otherwise.

### GetReferenceSequenceOk

`func (o *Ga4ghCigarUnit) GetReferenceSequenceOk() (string, bool)`

GetReferenceSequenceOk returns a tuple with the ReferenceSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceSequence

`func (o *Ga4ghCigarUnit) HasReferenceSequence() bool`

HasReferenceSequence returns a boolean if a field has been set.

### SetReferenceSequence

`func (o *Ga4ghCigarUnit) SetReferenceSequence(v string)`

SetReferenceSequence gets a reference to the given string and assigns it to the ReferenceSequence field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


