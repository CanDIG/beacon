# Ga4ghSearchSlidesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slides** | Pointer to [**[]Ga4ghSlide**](ga4ghSlide.md) | The list of slides. | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetSlides

`func (o *Ga4ghSearchSlidesResponse) GetSlides() []Ga4ghSlide`

GetSlides returns the Slides field if non-nil, zero value otherwise.

### GetSlidesOk

`func (o *Ga4ghSearchSlidesResponse) GetSlidesOk() ([]Ga4ghSlide, bool)`

GetSlidesOk returns a tuple with the Slides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSlides

`func (o *Ga4ghSearchSlidesResponse) HasSlides() bool`

HasSlides returns a boolean if a field has been set.

### SetSlides

`func (o *Ga4ghSearchSlidesResponse) SetSlides(v []Ga4ghSlide)`

SetSlides gets a reference to the given []Ga4ghSlide and assigns it to the Slides field.

### GetNextPageToken

`func (o *Ga4ghSearchSlidesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchSlidesResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchSlidesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchSlidesResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


