# Ga4ghSearchFusionDetectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fusiondetection** | Pointer to [**[]Ga4ghFusionDetection**](ga4ghFusionDetection.md) | The list of fusion/sv detections. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetFusiondetection

`func (o *Ga4ghSearchFusionDetectionResponse) GetFusiondetection() []Ga4ghFusionDetection`

GetFusiondetection returns the Fusiondetection field if non-nil, zero value otherwise.

### GetFusiondetectionOk

`func (o *Ga4ghSearchFusionDetectionResponse) GetFusiondetectionOk() ([]Ga4ghFusionDetection, bool)`

GetFusiondetectionOk returns a tuple with the Fusiondetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFusiondetection

`func (o *Ga4ghSearchFusionDetectionResponse) HasFusiondetection() bool`

HasFusiondetection returns a boolean if a field has been set.

### SetFusiondetection

`func (o *Ga4ghSearchFusionDetectionResponse) SetFusiondetection(v []Ga4ghFusionDetection)`

SetFusiondetection gets a reference to the given []Ga4ghFusionDetection and assigns it to the Fusiondetection field.

### GetNextPageToken

`func (o *Ga4ghSearchFusionDetectionResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchFusionDetectionResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchFusionDetectionResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchFusionDetectionResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


