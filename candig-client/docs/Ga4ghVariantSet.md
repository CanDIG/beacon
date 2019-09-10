# Ga4ghVariantSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The variant set ID. | [optional] 
**Name** | Pointer to **string** | The variant set name. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the dataset this variant set belongs to. | [optional] 
**ReferenceSetId** | Pointer to **string** | The ID of the reference set that describes the sequences used by the variants in this set. | [optional] 
**PatientId** | Pointer to **string** |  | [optional] 
**SampleId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**[]Ga4ghVariantSetMetadata**](ga4ghVariantSetMetadata.md) | Optional metadata associated with this variant set. This array can be used to store information about the variant set, such as information found in VCF header fields, that isn&#39;t already available in first class fields such as \&quot;name\&quot;. | [optional] 

## Methods

### GetId

`func (o *Ga4ghVariantSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghVariantSet) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghVariantSet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghVariantSet) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghVariantSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghVariantSet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghVariantSet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghVariantSet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDatasetId

`func (o *Ga4ghVariantSet) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghVariantSet) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghVariantSet) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghVariantSet) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetReferenceSetId

`func (o *Ga4ghVariantSet) GetReferenceSetId() string`

GetReferenceSetId returns the ReferenceSetId field if non-nil, zero value otherwise.

### GetReferenceSetIdOk

`func (o *Ga4ghVariantSet) GetReferenceSetIdOk() (string, bool)`

GetReferenceSetIdOk returns a tuple with the ReferenceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceSetId

`func (o *Ga4ghVariantSet) HasReferenceSetId() bool`

HasReferenceSetId returns a boolean if a field has been set.

### SetReferenceSetId

`func (o *Ga4ghVariantSet) SetReferenceSetId(v string)`

SetReferenceSetId gets a reference to the given string and assigns it to the ReferenceSetId field.

### GetPatientId

`func (o *Ga4ghVariantSet) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghVariantSet) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghVariantSet) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghVariantSet) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetSampleId

`func (o *Ga4ghVariantSet) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghVariantSet) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghVariantSet) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghVariantSet) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetMetadata

`func (o *Ga4ghVariantSet) GetMetadata() []Ga4ghVariantSetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Ga4ghVariantSet) GetMetadataOk() ([]Ga4ghVariantSetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetadata

`func (o *Ga4ghVariantSet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadata

`func (o *Ga4ghVariantSet) SetMetadata(v []Ga4ghVariantSetMetadata)`

SetMetadata gets a reference to the given []Ga4ghVariantSetMetadata and assigns it to the Metadata field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


