# Ga4ghSearchReadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignments** | Pointer to [**[]Ga4ghReadAlignment**](ga4ghReadAlignment.md) | The list of matching alignment messages, sorted by position. Unmapped reads, which have no position, are returned last. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetAlignments

`func (o *Ga4ghSearchReadsResponse) GetAlignments() []Ga4ghReadAlignment`

GetAlignments returns the Alignments field if non-nil, zero value otherwise.

### GetAlignmentsOk

`func (o *Ga4ghSearchReadsResponse) GetAlignmentsOk() ([]Ga4ghReadAlignment, bool)`

GetAlignmentsOk returns a tuple with the Alignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlignments

`func (o *Ga4ghSearchReadsResponse) HasAlignments() bool`

HasAlignments returns a boolean if a field has been set.

### SetAlignments

`func (o *Ga4ghSearchReadsResponse) SetAlignments(v []Ga4ghReadAlignment)`

SetAlignments gets a reference to the given []Ga4ghReadAlignment and assigns it to the Alignments field.

### GetNextPageToken

`func (o *Ga4ghSearchReadsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchReadsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchReadsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchReadsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


