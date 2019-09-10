# Ga4ghReferenceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The reference set ID. Unique in the repository. | [optional] 
**Name** | Pointer to **string** | The reference set name. | [optional] 
**Md5checksum** | Pointer to **string** | Order-independent MD5 checksum which identifies this &#x60;ReferenceSet&#x60;.  To compute this checksum, make a list of &#x60;Reference.md5checksum&#x60; for all &#x60;Reference&#x60; s in this set. Then sort that list, and take the MD5 hash of all the strings concatenated together. Express the hash as a lower-case hexadecimal string. | [optional] 
**Species** | Pointer to [**Ga4ghOntologyTerm**](ga4ghOntologyTerm.md) |  | [optional] 
**Description** | Pointer to **string** | Optional free text description of this reference set. | [optional] 
**AssemblyId** | Pointer to **string** | The remaining information is about the source of the sequences Public id of this reference set, such as &#x60;GRCh37&#x60;. | [optional] 
**SourceUri** | Pointer to **string** | Specifies a FASTA format file/string. | [optional] 
**SourceAccessions** | Pointer to **[]string** | All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) ideally with a version number, e.g. &#x60;NC_000001.11&#x60;. | [optional] 
**IsDerived** | Pointer to **bool** | A reference set may be derived from a source if it contains additional sequences, or some of the sequences within it are derived (see the definition of &#x60;isDerived&#x60; in &#x60;Reference&#x60;). | [optional] 

## Methods

### GetId

`func (o *Ga4ghReferenceSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghReferenceSet) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghReferenceSet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghReferenceSet) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghReferenceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghReferenceSet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghReferenceSet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghReferenceSet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetMd5checksum

`func (o *Ga4ghReferenceSet) GetMd5checksum() string`

GetMd5checksum returns the Md5checksum field if non-nil, zero value otherwise.

### GetMd5checksumOk

`func (o *Ga4ghReferenceSet) GetMd5checksumOk() (string, bool)`

GetMd5checksumOk returns a tuple with the Md5checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMd5checksum

`func (o *Ga4ghReferenceSet) HasMd5checksum() bool`

HasMd5checksum returns a boolean if a field has been set.

### SetMd5checksum

`func (o *Ga4ghReferenceSet) SetMd5checksum(v string)`

SetMd5checksum gets a reference to the given string and assigns it to the Md5checksum field.

### GetSpecies

`func (o *Ga4ghReferenceSet) GetSpecies() Ga4ghOntologyTerm`

GetSpecies returns the Species field if non-nil, zero value otherwise.

### GetSpeciesOk

`func (o *Ga4ghReferenceSet) GetSpeciesOk() (Ga4ghOntologyTerm, bool)`

GetSpeciesOk returns a tuple with the Species field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecies

`func (o *Ga4ghReferenceSet) HasSpecies() bool`

HasSpecies returns a boolean if a field has been set.

### SetSpecies

`func (o *Ga4ghReferenceSet) SetSpecies(v Ga4ghOntologyTerm)`

SetSpecies gets a reference to the given Ga4ghOntologyTerm and assigns it to the Species field.

### GetDescription

`func (o *Ga4ghReferenceSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghReferenceSet) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghReferenceSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghReferenceSet) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetAssemblyId

`func (o *Ga4ghReferenceSet) GetAssemblyId() string`

GetAssemblyId returns the AssemblyId field if non-nil, zero value otherwise.

### GetAssemblyIdOk

`func (o *Ga4ghReferenceSet) GetAssemblyIdOk() (string, bool)`

GetAssemblyIdOk returns a tuple with the AssemblyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssemblyId

`func (o *Ga4ghReferenceSet) HasAssemblyId() bool`

HasAssemblyId returns a boolean if a field has been set.

### SetAssemblyId

`func (o *Ga4ghReferenceSet) SetAssemblyId(v string)`

SetAssemblyId gets a reference to the given string and assigns it to the AssemblyId field.

### GetSourceUri

`func (o *Ga4ghReferenceSet) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *Ga4ghReferenceSet) GetSourceUriOk() (string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceUri

`func (o *Ga4ghReferenceSet) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### SetSourceUri

`func (o *Ga4ghReferenceSet) SetSourceUri(v string)`

SetSourceUri gets a reference to the given string and assigns it to the SourceUri field.

### GetSourceAccessions

`func (o *Ga4ghReferenceSet) GetSourceAccessions() []string`

GetSourceAccessions returns the SourceAccessions field if non-nil, zero value otherwise.

### GetSourceAccessionsOk

`func (o *Ga4ghReferenceSet) GetSourceAccessionsOk() ([]string, bool)`

GetSourceAccessionsOk returns a tuple with the SourceAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceAccessions

`func (o *Ga4ghReferenceSet) HasSourceAccessions() bool`

HasSourceAccessions returns a boolean if a field has been set.

### SetSourceAccessions

`func (o *Ga4ghReferenceSet) SetSourceAccessions(v []string)`

SetSourceAccessions gets a reference to the given []string and assigns it to the SourceAccessions field.

### GetIsDerived

`func (o *Ga4ghReferenceSet) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *Ga4ghReferenceSet) GetIsDerivedOk() (bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDerived

`func (o *Ga4ghReferenceSet) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerived

`func (o *Ga4ghReferenceSet) SetIsDerived(v bool)`

SetIsDerived gets a reference to the given bool and assigns it to the IsDerived field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


