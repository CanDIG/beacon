# Ga4ghSearchContinuousRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousSetId** | Pointer to **string** | The annotation set to search within. Required value. | [optional] 
**ReferenceName** | Pointer to **string** | Get continuous values on this reference. Required value. | [optional] 
**Start** | Pointer to **string** | The beginning of the window (0-based, inclusive) for which continuous values should be returned. Required value. | [optional] 
**End** | Pointer to **string** | The end of the window (0-based, exclusive) for which continuous values should be returned.  Required value. | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#x60;next_page_token&#x60; from the previous response. | [optional] 

## Methods

### GetContinuousSetId

`func (o *Ga4ghSearchContinuousRequest) GetContinuousSetId() string`

GetContinuousSetId returns the ContinuousSetId field if non-nil, zero value otherwise.

### GetContinuousSetIdOk

`func (o *Ga4ghSearchContinuousRequest) GetContinuousSetIdOk() (string, bool)`

GetContinuousSetIdOk returns a tuple with the ContinuousSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContinuousSetId

`func (o *Ga4ghSearchContinuousRequest) HasContinuousSetId() bool`

HasContinuousSetId returns a boolean if a field has been set.

### SetContinuousSetId

`func (o *Ga4ghSearchContinuousRequest) SetContinuousSetId(v string)`

SetContinuousSetId gets a reference to the given string and assigns it to the ContinuousSetId field.

### GetReferenceName

`func (o *Ga4ghSearchContinuousRequest) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Ga4ghSearchContinuousRequest) GetReferenceNameOk() (string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceName

`func (o *Ga4ghSearchContinuousRequest) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### SetReferenceName

`func (o *Ga4ghSearchContinuousRequest) SetReferenceName(v string)`

SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.

### GetStart

`func (o *Ga4ghSearchContinuousRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Ga4ghSearchContinuousRequest) GetStartOk() (string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Ga4ghSearchContinuousRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Ga4ghSearchContinuousRequest) SetStart(v string)`

SetStart gets a reference to the given string and assigns it to the Start field.

### GetEnd

`func (o *Ga4ghSearchContinuousRequest) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Ga4ghSearchContinuousRequest) GetEndOk() (string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Ga4ghSearchContinuousRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Ga4ghSearchContinuousRequest) SetEnd(v string)`

SetEnd gets a reference to the given string and assigns it to the End field.

### GetPageSize

`func (o *Ga4ghSearchContinuousRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchContinuousRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchContinuousRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchContinuousRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchContinuousRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchContinuousRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchContinuousRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchContinuousRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


