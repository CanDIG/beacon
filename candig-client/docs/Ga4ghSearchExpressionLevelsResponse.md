# Ga4ghSearchExpressionLevelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpressionLevels** | Pointer to [**[]Ga4ghExpressionLevel**](ga4ghExpressionLevel.md) | The list of matching quantifications. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of &#39;nextPageToken&#39; from the previous response. | [optional] 

## Methods

### GetExpressionLevels

`func (o *Ga4ghSearchExpressionLevelsResponse) GetExpressionLevels() []Ga4ghExpressionLevel`

GetExpressionLevels returns the ExpressionLevels field if non-nil, zero value otherwise.

### GetExpressionLevelsOk

`func (o *Ga4ghSearchExpressionLevelsResponse) GetExpressionLevelsOk() ([]Ga4ghExpressionLevel, bool)`

GetExpressionLevelsOk returns a tuple with the ExpressionLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpressionLevels

`func (o *Ga4ghSearchExpressionLevelsResponse) HasExpressionLevels() bool`

HasExpressionLevels returns a boolean if a field has been set.

### SetExpressionLevels

`func (o *Ga4ghSearchExpressionLevelsResponse) SetExpressionLevels(v []Ga4ghExpressionLevel)`

SetExpressionLevels gets a reference to the given []Ga4ghExpressionLevel and assigns it to the ExpressionLevels field.

### GetNextPageToken

`func (o *Ga4ghSearchExpressionLevelsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchExpressionLevelsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchExpressionLevelsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchExpressionLevelsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


