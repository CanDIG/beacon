# Ga4ghSearchVariantCallingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variantcalling** | Pointer to [**[]Ga4ghVariantCalling**](ga4ghVariantCalling.md) | The list of variant calling metadata. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets.Provide this value in a subsequent request to return the next page ofresults. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetVariantcalling

`func (o *Ga4ghSearchVariantCallingResponse) GetVariantcalling() []Ga4ghVariantCalling`

GetVariantcalling returns the Variantcalling field if non-nil, zero value otherwise.

### GetVariantcallingOk

`func (o *Ga4ghSearchVariantCallingResponse) GetVariantcallingOk() ([]Ga4ghVariantCalling, bool)`

GetVariantcallingOk returns a tuple with the Variantcalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantcalling

`func (o *Ga4ghSearchVariantCallingResponse) HasVariantcalling() bool`

HasVariantcalling returns a boolean if a field has been set.

### SetVariantcalling

`func (o *Ga4ghSearchVariantCallingResponse) SetVariantcalling(v []Ga4ghVariantCalling)`

SetVariantcalling gets a reference to the given []Ga4ghVariantCalling and assigns it to the Variantcalling field.

### GetNextPageToken

`func (o *Ga4ghSearchVariantCallingResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchVariantCallingResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchVariantCallingResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchVariantCallingResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


