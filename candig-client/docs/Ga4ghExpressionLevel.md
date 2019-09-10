# Ga4ghExpressionLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RnaQuantificationId** | Pointer to **string** |  | [optional] 
**RawReadCount** | Pointer to **float32** | The number of reads mapped to this feature. | [optional] 
**Expression** | Pointer to **float32** | Numerical expression value. | [optional] 
**IsNormalized** | Pointer to **bool** | True if the expression value is a normalized value. | [optional] 
**Units** | Pointer to [**Ga4ghExpressionUnit**](ga4ghExpressionUnit.md) |  | [optional] 
**Score** | Pointer to **float32** | Weighted score for the expression value. | [optional] 
**ConfIntervalLow** | Pointer to **float32** | Lower bound of the confidence interval on the expression value. | [optional] 
**ConfIntervalHigh** | Pointer to **float32** | Upper bound of the confidence interval on the expression value. | [optional] 

## Methods

### GetId

`func (o *Ga4ghExpressionLevel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghExpressionLevel) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghExpressionLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghExpressionLevel) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghExpressionLevel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghExpressionLevel) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghExpressionLevel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghExpressionLevel) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetRnaQuantificationId

`func (o *Ga4ghExpressionLevel) GetRnaQuantificationId() string`

GetRnaQuantificationId returns the RnaQuantificationId field if non-nil, zero value otherwise.

### GetRnaQuantificationIdOk

`func (o *Ga4ghExpressionLevel) GetRnaQuantificationIdOk() (string, bool)`

GetRnaQuantificationIdOk returns a tuple with the RnaQuantificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaQuantificationId

`func (o *Ga4ghExpressionLevel) HasRnaQuantificationId() bool`

HasRnaQuantificationId returns a boolean if a field has been set.

### SetRnaQuantificationId

`func (o *Ga4ghExpressionLevel) SetRnaQuantificationId(v string)`

SetRnaQuantificationId gets a reference to the given string and assigns it to the RnaQuantificationId field.

### GetRawReadCount

`func (o *Ga4ghExpressionLevel) GetRawReadCount() float32`

GetRawReadCount returns the RawReadCount field if non-nil, zero value otherwise.

### GetRawReadCountOk

`func (o *Ga4ghExpressionLevel) GetRawReadCountOk() (float32, bool)`

GetRawReadCountOk returns a tuple with the RawReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRawReadCount

`func (o *Ga4ghExpressionLevel) HasRawReadCount() bool`

HasRawReadCount returns a boolean if a field has been set.

### SetRawReadCount

`func (o *Ga4ghExpressionLevel) SetRawReadCount(v float32)`

SetRawReadCount gets a reference to the given float32 and assigns it to the RawReadCount field.

### GetExpression

`func (o *Ga4ghExpressionLevel) GetExpression() float32`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Ga4ghExpressionLevel) GetExpressionOk() (float32, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpression

`func (o *Ga4ghExpressionLevel) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpression

`func (o *Ga4ghExpressionLevel) SetExpression(v float32)`

SetExpression gets a reference to the given float32 and assigns it to the Expression field.

### GetIsNormalized

`func (o *Ga4ghExpressionLevel) GetIsNormalized() bool`

GetIsNormalized returns the IsNormalized field if non-nil, zero value otherwise.

### GetIsNormalizedOk

`func (o *Ga4ghExpressionLevel) GetIsNormalizedOk() (bool, bool)`

GetIsNormalizedOk returns a tuple with the IsNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsNormalized

`func (o *Ga4ghExpressionLevel) HasIsNormalized() bool`

HasIsNormalized returns a boolean if a field has been set.

### SetIsNormalized

`func (o *Ga4ghExpressionLevel) SetIsNormalized(v bool)`

SetIsNormalized gets a reference to the given bool and assigns it to the IsNormalized field.

### GetUnits

`func (o *Ga4ghExpressionLevel) GetUnits() Ga4ghExpressionUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Ga4ghExpressionLevel) GetUnitsOk() (Ga4ghExpressionUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnits

`func (o *Ga4ghExpressionLevel) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnits

`func (o *Ga4ghExpressionLevel) SetUnits(v Ga4ghExpressionUnit)`

SetUnits gets a reference to the given Ga4ghExpressionUnit and assigns it to the Units field.

### GetScore

`func (o *Ga4ghExpressionLevel) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Ga4ghExpressionLevel) GetScoreOk() (float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScore

`func (o *Ga4ghExpressionLevel) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScore

`func (o *Ga4ghExpressionLevel) SetScore(v float32)`

SetScore gets a reference to the given float32 and assigns it to the Score field.

### GetConfIntervalLow

`func (o *Ga4ghExpressionLevel) GetConfIntervalLow() float32`

GetConfIntervalLow returns the ConfIntervalLow field if non-nil, zero value otherwise.

### GetConfIntervalLowOk

`func (o *Ga4ghExpressionLevel) GetConfIntervalLowOk() (float32, bool)`

GetConfIntervalLowOk returns a tuple with the ConfIntervalLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfIntervalLow

`func (o *Ga4ghExpressionLevel) HasConfIntervalLow() bool`

HasConfIntervalLow returns a boolean if a field has been set.

### SetConfIntervalLow

`func (o *Ga4ghExpressionLevel) SetConfIntervalLow(v float32)`

SetConfIntervalLow gets a reference to the given float32 and assigns it to the ConfIntervalLow field.

### GetConfIntervalHigh

`func (o *Ga4ghExpressionLevel) GetConfIntervalHigh() float32`

GetConfIntervalHigh returns the ConfIntervalHigh field if non-nil, zero value otherwise.

### GetConfIntervalHighOk

`func (o *Ga4ghExpressionLevel) GetConfIntervalHighOk() (float32, bool)`

GetConfIntervalHighOk returns a tuple with the ConfIntervalHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfIntervalHigh

`func (o *Ga4ghExpressionLevel) HasConfIntervalHigh() bool`

HasConfIntervalHigh returns a boolean if a field has been set.

### SetConfIntervalHigh

`func (o *Ga4ghExpressionLevel) SetConfIntervalHigh(v float32)`

SetConfIntervalHigh gets a reference to the given float32 and assigns it to the ConfIntervalHigh field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


