# Ga4ghReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The reference ID. Unique within the repository. | [optional] 
**Length** | Pointer to **string** | The length of this reference&#39;s sequence. | [optional] 
**Md5checksum** | Pointer to **string** | The MD5 checksum uniquely representing this &#x60;Reference&#x60; as a lower-case hexadecimal string, calculated as the MD5 of the upper-case sequence excluding all whitespace characters (this is equivalent to SQ:M5 in SAM). | [optional] 
**Name** | Pointer to **string** | The unique name of this reference within the Reference Set (e.g. &#39;22&#39;). | [optional] 
**SourceUri** | Pointer to **string** | The URI from which the sequence was obtained. Specifies a FASTA format file/string with one name, sequence pair. In most cases, clients should call the &#x60;getReferenceBases()&#x60; method to obtain sequence bases for a &#x60;Reference&#x60; instead of attempting to retrieve this URI. | [optional] 
**SourceAccessions** | Pointer to **[]string** | All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) which must include a version number, e.g. &#x60;GCF_000001405.26&#x60;. | [optional] 
**IsDerived** | Pointer to **bool** | A sequence X is said to be derived from source sequence Y, if X and Y are of the same length and the per-base sequence divergence at A/C/G/T bases is sufficiently small. Two sequences derived from the same official sequence share the same coordinates and annotations, and can be replaced with the official sequence for certain use cases. | [optional] 
**SourceDivergence** | Pointer to **float32** | The &#x60;sourceDivergence&#x60; is the fraction of non-indel bases that do not match the reference this message was derived from. | [optional] 
**Species** | Pointer to [**Ga4ghOntologyTerm**](ga4ghOntologyTerm.md) |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghReference) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghReference) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLength

`func (o *Ga4ghReference) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Ga4ghReference) GetLengthOk() (string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLength

`func (o *Ga4ghReference) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLength

`func (o *Ga4ghReference) SetLength(v string)`

SetLength gets a reference to the given string and assigns it to the Length field.

### GetMd5checksum

`func (o *Ga4ghReference) GetMd5checksum() string`

GetMd5checksum returns the Md5checksum field if non-nil, zero value otherwise.

### GetMd5checksumOk

`func (o *Ga4ghReference) GetMd5checksumOk() (string, bool)`

GetMd5checksumOk returns a tuple with the Md5checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMd5checksum

`func (o *Ga4ghReference) HasMd5checksum() bool`

HasMd5checksum returns a boolean if a field has been set.

### SetMd5checksum

`func (o *Ga4ghReference) SetMd5checksum(v string)`

SetMd5checksum gets a reference to the given string and assigns it to the Md5checksum field.

### GetName

`func (o *Ga4ghReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghReference) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghReference) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghReference) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSourceUri

`func (o *Ga4ghReference) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *Ga4ghReference) GetSourceUriOk() (string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceUri

`func (o *Ga4ghReference) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### SetSourceUri

`func (o *Ga4ghReference) SetSourceUri(v string)`

SetSourceUri gets a reference to the given string and assigns it to the SourceUri field.

### GetSourceAccessions

`func (o *Ga4ghReference) GetSourceAccessions() []string`

GetSourceAccessions returns the SourceAccessions field if non-nil, zero value otherwise.

### GetSourceAccessionsOk

`func (o *Ga4ghReference) GetSourceAccessionsOk() ([]string, bool)`

GetSourceAccessionsOk returns a tuple with the SourceAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceAccessions

`func (o *Ga4ghReference) HasSourceAccessions() bool`

HasSourceAccessions returns a boolean if a field has been set.

### SetSourceAccessions

`func (o *Ga4ghReference) SetSourceAccessions(v []string)`

SetSourceAccessions gets a reference to the given []string and assigns it to the SourceAccessions field.

### GetIsDerived

`func (o *Ga4ghReference) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *Ga4ghReference) GetIsDerivedOk() (bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDerived

`func (o *Ga4ghReference) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerived

`func (o *Ga4ghReference) SetIsDerived(v bool)`

SetIsDerived gets a reference to the given bool and assigns it to the IsDerived field.

### GetSourceDivergence

`func (o *Ga4ghReference) GetSourceDivergence() float32`

GetSourceDivergence returns the SourceDivergence field if non-nil, zero value otherwise.

### GetSourceDivergenceOk

`func (o *Ga4ghReference) GetSourceDivergenceOk() (float32, bool)`

GetSourceDivergenceOk returns a tuple with the SourceDivergence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceDivergence

`func (o *Ga4ghReference) HasSourceDivergence() bool`

HasSourceDivergence returns a boolean if a field has been set.

### SetSourceDivergence

`func (o *Ga4ghReference) SetSourceDivergence(v float32)`

SetSourceDivergence gets a reference to the given float32 and assigns it to the SourceDivergence field.

### GetSpecies

`func (o *Ga4ghReference) GetSpecies() Ga4ghOntologyTerm`

GetSpecies returns the Species field if non-nil, zero value otherwise.

### GetSpeciesOk

`func (o *Ga4ghReference) GetSpeciesOk() (Ga4ghOntologyTerm, bool)`

GetSpeciesOk returns a tuple with the Species field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecies

`func (o *Ga4ghReference) HasSpecies() bool`

HasSpecies returns a boolean if a field has been set.

### SetSpecies

`func (o *Ga4ghReference) SetSpecies(v Ga4ghOntologyTerm)`

SetSpecies gets a reference to the given Ga4ghOntologyTerm and assigns it to the Species field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


