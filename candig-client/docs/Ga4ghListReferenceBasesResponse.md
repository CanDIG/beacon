# Ga4ghListReferenceBasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **string** | The offset position (0-based) of the given sequence from the start of this &#x60;Reference&#x60;. This value will differ for each page in a paginated request. | [optional] 
**Sequence** | Pointer to **string** | A substring of the bases that make up this reference. Bases are represented as IUPAC-IUB codes; this string matches the regexp &#x60;[ACGTMRWSYKVHDBN]*&#x60;. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetOffset

`func (o *Ga4ghListReferenceBasesResponse) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Ga4ghListReferenceBasesResponse) GetOffsetOk() (string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOffset

`func (o *Ga4ghListReferenceBasesResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffset

`func (o *Ga4ghListReferenceBasesResponse) SetOffset(v string)`

SetOffset gets a reference to the given string and assigns it to the Offset field.

### GetSequence

`func (o *Ga4ghListReferenceBasesResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *Ga4ghListReferenceBasesResponse) GetSequenceOk() (string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSequence

`func (o *Ga4ghListReferenceBasesResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequence

`func (o *Ga4ghListReferenceBasesResponse) SetSequence(v string)`

SetSequence gets a reference to the given string and assigns it to the Sequence field.

### GetNextPageToken

`func (o *Ga4ghListReferenceBasesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghListReferenceBasesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghListReferenceBasesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghListReferenceBasesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


