# Ga4ghCallSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The call set ID. | [optional] 
**Name** | Pointer to **string** | The call set name. | [optional] 
**BiosampleId** | Pointer to **string** | The Biosample the call set data was generated from. | [optional] 
**VariantSetIds** | Pointer to **[]string** | The IDs of the variant sets this call set has calls in. | [optional] 
**Created** | Pointer to **string** | The date this call set was created in milliseconds from the epoch. | [optional] 
**Updated** | Pointer to **string** | The time at which this call set was last updated in milliseconds from the epoch. | [optional] 
**Attributes** | Pointer to [**Ga4ghAttributes**](ga4ghAttributes.md) |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghCallSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghCallSet) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghCallSet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghCallSet) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghCallSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghCallSet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghCallSet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghCallSet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetBiosampleId

`func (o *Ga4ghCallSet) GetBiosampleId() string`

GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.

### GetBiosampleIdOk

`func (o *Ga4ghCallSet) GetBiosampleIdOk() (string, bool)`

GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiosampleId

`func (o *Ga4ghCallSet) HasBiosampleId() bool`

HasBiosampleId returns a boolean if a field has been set.

### SetBiosampleId

`func (o *Ga4ghCallSet) SetBiosampleId(v string)`

SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.

### GetVariantSetIds

`func (o *Ga4ghCallSet) GetVariantSetIds() []string`

GetVariantSetIds returns the VariantSetIds field if non-nil, zero value otherwise.

### GetVariantSetIdsOk

`func (o *Ga4ghCallSet) GetVariantSetIdsOk() ([]string, bool)`

GetVariantSetIdsOk returns a tuple with the VariantSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantSetIds

`func (o *Ga4ghCallSet) HasVariantSetIds() bool`

HasVariantSetIds returns a boolean if a field has been set.

### SetVariantSetIds

`func (o *Ga4ghCallSet) SetVariantSetIds(v []string)`

SetVariantSetIds gets a reference to the given []string and assigns it to the VariantSetIds field.

### GetCreated

`func (o *Ga4ghCallSet) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghCallSet) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghCallSet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghCallSet) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghCallSet) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghCallSet) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghCallSet) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghCallSet) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetAttributes

`func (o *Ga4ghCallSet) GetAttributes() Ga4ghAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Ga4ghCallSet) GetAttributesOk() (Ga4ghAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Ga4ghCallSet) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Ga4ghCallSet) SetAttributes(v Ga4ghAttributes)`

SetAttributes gets a reference to the given Ga4ghAttributes and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


