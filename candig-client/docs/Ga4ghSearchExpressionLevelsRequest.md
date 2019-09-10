# Ga4ghSearchExpressionLevelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RnaQuantificationId** | Pointer to **string** | The rnaQuantification to restrict search to. | [optional] 
**Names** | Pointer to **[]string** | Only return expressions with any of the names (strict string matching). | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **int32** | Specifies the maximum number of results to return in a single page. If unspecified, a system default will be used. | [optional] 
**PageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#39;nextPageToken&#39; from the previous response. | [optional] 

## Methods

### GetRnaQuantificationId

`func (o *Ga4ghSearchExpressionLevelsRequest) GetRnaQuantificationId() string`

GetRnaQuantificationId returns the RnaQuantificationId field if non-nil, zero value otherwise.

### GetRnaQuantificationIdOk

`func (o *Ga4ghSearchExpressionLevelsRequest) GetRnaQuantificationIdOk() (string, bool)`

GetRnaQuantificationIdOk returns a tuple with the RnaQuantificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaQuantificationId

`func (o *Ga4ghSearchExpressionLevelsRequest) HasRnaQuantificationId() bool`

HasRnaQuantificationId returns a boolean if a field has been set.

### SetRnaQuantificationId

`func (o *Ga4ghSearchExpressionLevelsRequest) SetRnaQuantificationId(v string)`

SetRnaQuantificationId gets a reference to the given string and assigns it to the RnaQuantificationId field.

### GetNames

`func (o *Ga4ghSearchExpressionLevelsRequest) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Ga4ghSearchExpressionLevelsRequest) GetNamesOk() ([]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *Ga4ghSearchExpressionLevelsRequest) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *Ga4ghSearchExpressionLevelsRequest) SetNames(v []string)`

SetNames gets a reference to the given []string and assigns it to the Names field.

### GetThreshold

`func (o *Ga4ghSearchExpressionLevelsRequest) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Ga4ghSearchExpressionLevelsRequest) GetThresholdOk() (float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThreshold

`func (o *Ga4ghSearchExpressionLevelsRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### SetThreshold

`func (o *Ga4ghSearchExpressionLevelsRequest) SetThreshold(v float32)`

SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.

### GetPageSize

`func (o *Ga4ghSearchExpressionLevelsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Ga4ghSearchExpressionLevelsRequest) GetPageSizeOk() (int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageSize

`func (o *Ga4ghSearchExpressionLevelsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSize

`func (o *Ga4ghSearchExpressionLevelsRequest) SetPageSize(v int32)`

SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.

### GetPageToken

`func (o *Ga4ghSearchExpressionLevelsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *Ga4ghSearchExpressionLevelsRequest) GetPageTokenOk() (string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPageToken

`func (o *Ga4ghSearchExpressionLevelsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### SetPageToken

`func (o *Ga4ghSearchExpressionLevelsRequest) SetPageToken(v string)`

SetPageToken gets a reference to the given string and assigns it to the PageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


