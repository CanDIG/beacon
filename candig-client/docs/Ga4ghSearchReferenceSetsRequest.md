# Ga4ghSearchReferenceSetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Md5checksum** | Pointer to **string** | If unset, return the reference sets for which the &#x60;md5checksum&#x60; matches this string (case-sensitive, exact match). See &#x60;ReferenceSet::md5checksum&#x60; for details. | [optional] 
**Accession** | Pointer to **string** | If unset, return the reference sets for which the &#x60;accession&#x60; matches this string (case-sensitive, exact match). | [optional] 
**AssemblyId** | Pointer to **string** | If unset, return the reference sets for which the &#x60;assemblyId&#x60; matches this string (case-sensitive, exact match). | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 

## Methods

### GetMd5checksum

`func (o *Ga4ghSearchReferenceSetsRequest) GetMd5checksum() string`

GetMd5checksum returns the Md5checksum field if non-nil, zero value otherwise.

### GetMd5checksumOk

`func (o *Ga4ghSearchReferenceSetsRequest) GetMd5checksumOk() (string, bool)`

GetMd5checksumOk returns a tuple with the Md5checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMd5checksum

`func (o *Ga4ghSearchReferenceSetsRequest) HasMd5checksum() bool`

HasMd5checksum returns a boolean if a field has been set.

### SetMd5checksum

`func (o *Ga4ghSearchReferenceSetsRequest) SetMd5checksum(v string)`

SetMd5checksum gets a reference to the given string and assigns it to the Md5checksum field.

### GetAccession

`func (o *Ga4ghSearchReferenceSetsRequest) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *Ga4ghSearchReferenceSetsRequest) GetAccessionOk() (string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccession

`func (o *Ga4ghSearchReferenceSetsRequest) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### SetAccession

`func (o *Ga4ghSearchReferenceSetsRequest) SetAccession(v string)`

SetAccession gets a reference to the given string and assigns it to the Accession field.

### GetAssemblyId

`func (o *Ga4ghSearchReferenceSetsRequest) GetAssemblyId() string`

GetAssemblyId returns the AssemblyId field if non-nil, zero value otherwise.

### GetAssemblyIdOk

`func (o *Ga4ghSearchReferenceSetsRequest) GetAssemblyIdOk() (string, bool)`

GetAssemblyIdOk returns a tuple with the AssemblyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssemblyId

`func (o *Ga4ghSearchReferenceSetsRequest) HasAssemblyId() bool`

HasAssemblyId returns a boolean if a field has been set.

### SetAssemblyId

`func (o *Ga4ghSearchReferenceSetsRequest) SetAssemblyId(v string)`

SetAssemblyId gets a reference to the given string and assigns it to the AssemblyId field.

### GetPageSize

`func (o *Ga4ghSearchReferenceSetsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchReferenceSetsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchReferenceSetsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchReferenceSetsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchReferenceSetsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchReferenceSetsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchReferenceSetsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchReferenceSetsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


