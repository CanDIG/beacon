# Ga4ghContinuous

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | The start position at which this signal occurs (0-based). This corresponds to the first base of the string of reference bases. Genomic positions are non-negative integers less than the reference length. | [optional] 
**Values** | Pointer to **[]float64** | The contiguous data values. Unsampled bases are given as NaN. | [optional] 
**ContinuousSetId** | Pointer to **string** | Identifier for the containing continous set. | [optional] 
**ReferenceName** | Pointer to **string** | The reference on which this signal is defined (e.g. &#x60;chr20&#x60; or &#x60;X&#x60;). | [optional] 

## Methods

### GetStart

`func (o *Ga4ghContinuous) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghContinuous) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghContinuous) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghContinuous) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetValues

`func (o *Ga4ghContinuous) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Ga4ghContinuous) GetValuesOk() ([]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *Ga4ghContinuous) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *Ga4ghContinuous) SetValues(v []float64)`

SetValues gets a reference to the given []float64 and assigns it to the Values field.

### GetContinuousSetId

`func (o *Ga4ghContinuous) GetContinuousSetId() string`

GetContinuousSetId returns the ContinuousSetId field if non-nil, zero value otherwise.

### GetContinuousSetIdOk

`func (o *Ga4ghContinuous) GetContinuousSetIdOk() (string, bool)`

GetContinuousSetIdOk returns a tuple with the ContinuousSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContinuousSetId

`func (o *Ga4ghContinuous) HasContinuousSetId() bool`

HasContinuousSetId returns a boolean if a field has been set.

### SetContinuousSetId

`func (o *Ga4ghContinuous) SetContinuousSetId(v string)`

SetContinuousSetId gets a reference to the given string and assigns it to the ContinuousSetId field.

### GetReferenceName

`func (o *Ga4ghContinuous) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghContinuous) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghContinuous) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghContinuous) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


