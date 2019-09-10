# Ga4ghSearchExpressionAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressionanalysis** | Pointer to [**[]Ga4ghExpressionAnalysis**](ga4ghExpressionAnalysis.md) | The list of expression analysis. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetExpressionanalysis

`func (o *Ga4ghSearchExpressionAnalysisResponse) GetExpressionanalysis() []Ga4ghExpressionAnalysis`

GetExpressionanalysis returns the Expressionanalysis field if non-nil, zero value otherwise.

### GetExpressionanalysisOk

`func (o *Ga4ghSearchExpressionAnalysisResponse) GetExpressionanalysisOk() ([]Ga4ghExpressionAnalysis, bool)`

GetExpressionanalysisOk returns a tuple with the Expressionanalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpressionanalysis

`func (o *Ga4ghSearchExpressionAnalysisResponse) HasExpressionanalysis() bool`

HasExpressionanalysis returns a boolean if a field has been set.

### SetExpressionanalysis

`func (o *Ga4ghSearchExpressionAnalysisResponse) SetExpressionanalysis(v []Ga4ghExpressionAnalysis)`

SetExpressionanalysis gets a reference to the given []Ga4ghExpressionAnalysis and assigns it to the Expressionanalysis field.

### GetNextPageToken

`func (o *Ga4ghSearchExpressionAnalysisResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchExpressionAnalysisResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchExpressionAnalysisResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchExpressionAnalysisResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


