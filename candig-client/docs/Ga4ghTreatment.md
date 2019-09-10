# Ga4ghTreatment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is unique in the context of the server instance. | [optional] 
**DatasetId** | Pointer to **string** | The ID of the dataset this object belongs to. | [optional] 
**Name** | Pointer to **string** | This is a label or symbolic identifier for this object. | [optional] 
**Description** | Pointer to **string** | This attribute contains human readable text. | [optional] 
**Created** | Pointer to **string** | The :ref:&#x60;ISO 8601&lt;metadata_date_time&gt;&#x60; time at which this record was created. | [optional] 
**Updated** | Pointer to **string** | The :ref:&#x60;ISO 8601&lt;metadata_date_time&gt;&#x60; time at which this record was updated. | [optional] 
**PatientId** | Pointer to **string** |  | [optional] 
**CourseNumber** | Pointer to **string** |  | [optional] 
**TherapeuticModality** | Pointer to **string** |  | [optional] 
**TreatmentPlanType** | Pointer to **string** |  | [optional] 
**TreatmentIntent** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**StopDate** | Pointer to **string** |  | [optional] 
**ReasonForEndingTheTreatment** | Pointer to **string** |  | [optional] 
**ResponseToTreatment** | Pointer to **string** |  | [optional] 
**ResponseCriteriaUsed** | Pointer to **string** |  | [optional] 
**DateOfRecurrenceOrProgressionAfterThisTreatment** | Pointer to **string** |  | [optional] 
**UnexpectedOrUnusualToxicityDuringTreatment** | Pointer to **string** |  | [optional] 
**DiagnosisId** | Pointer to **string** |  | [optional] 
**TreatmentPlanId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghTreatment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghTreatment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghTreatment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghTreatment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghTreatment) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghTreatment) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghTreatment) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghTreatment) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghTreatment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghTreatment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghTreatment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghTreatment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghTreatment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghTreatment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghTreatment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghTreatment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghTreatment) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghTreatment) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghTreatment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghTreatment) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghTreatment) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghTreatment) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghTreatment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghTreatment) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghTreatment) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghTreatment) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghTreatment) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghTreatment) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetCourseNumber

`func (o *Ga4ghTreatment) GetCourseNumber() string`

GetCourseNumber returns the CourseNumber field if non-nil, zero value otherwise.

### GetCourseNumberOk

`func (o *Ga4ghTreatment) GetCourseNumberOk() (string, bool)`

GetCourseNumberOk returns a tuple with the CourseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCourseNumber

`func (o *Ga4ghTreatment) HasCourseNumber() bool`

HasCourseNumber returns a boolean if a field has been set.

### SetCourseNumber

`func (o *Ga4ghTreatment) SetCourseNumber(v string)`

SetCourseNumber gets a reference to the given string and assigns it to the CourseNumber field.

### GetTherapeuticModality

`func (o *Ga4ghTreatment) GetTherapeuticModality() string`

GetTherapeuticModality returns the TherapeuticModality field if non-nil, zero value otherwise.

### GetTherapeuticModalityOk

`func (o *Ga4ghTreatment) GetTherapeuticModalityOk() (string, bool)`

GetTherapeuticModalityOk returns a tuple with the TherapeuticModality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTherapeuticModality

`func (o *Ga4ghTreatment) HasTherapeuticModality() bool`

HasTherapeuticModality returns a boolean if a field has been set.

### SetTherapeuticModality

`func (o *Ga4ghTreatment) SetTherapeuticModality(v string)`

SetTherapeuticModality gets a reference to the given string and assigns it to the TherapeuticModality field.

### GetTreatmentPlanType

`func (o *Ga4ghTreatment) GetTreatmentPlanType() string`

GetTreatmentPlanType returns the TreatmentPlanType field if non-nil, zero value otherwise.

### GetTreatmentPlanTypeOk

`func (o *Ga4ghTreatment) GetTreatmentPlanTypeOk() (string, bool)`

GetTreatmentPlanTypeOk returns a tuple with the TreatmentPlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatmentPlanType

`func (o *Ga4ghTreatment) HasTreatmentPlanType() bool`

HasTreatmentPlanType returns a boolean if a field has been set.

### SetTreatmentPlanType

`func (o *Ga4ghTreatment) SetTreatmentPlanType(v string)`

SetTreatmentPlanType gets a reference to the given string and assigns it to the TreatmentPlanType field.

### GetTreatmentIntent

`func (o *Ga4ghTreatment) GetTreatmentIntent() string`

GetTreatmentIntent returns the TreatmentIntent field if non-nil, zero value otherwise.

### GetTreatmentIntentOk

`func (o *Ga4ghTreatment) GetTreatmentIntentOk() (string, bool)`

GetTreatmentIntentOk returns a tuple with the TreatmentIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatmentIntent

`func (o *Ga4ghTreatment) HasTreatmentIntent() bool`

HasTreatmentIntent returns a boolean if a field has been set.

### SetTreatmentIntent

`func (o *Ga4ghTreatment) SetTreatmentIntent(v string)`

SetTreatmentIntent gets a reference to the given string and assigns it to the TreatmentIntent field.

### GetStartDate

`func (o *Ga4ghTreatment) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Ga4ghTreatment) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *Ga4ghTreatment) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *Ga4ghTreatment) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetStopDate

`func (o *Ga4ghTreatment) GetStopDate() string`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *Ga4ghTreatment) GetStopDateOk() (string, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStopDate

`func (o *Ga4ghTreatment) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.

### SetStopDate

`func (o *Ga4ghTreatment) SetStopDate(v string)`

SetStopDate gets a reference to the given string and assigns it to the StopDate field.

### GetReasonForEndingTheTreatment

`func (o *Ga4ghTreatment) GetReasonForEndingTheTreatment() string`

GetReasonForEndingTheTreatment returns the ReasonForEndingTheTreatment field if non-nil, zero value otherwise.

### GetReasonForEndingTheTreatmentOk

`func (o *Ga4ghTreatment) GetReasonForEndingTheTreatmentOk() (string, bool)`

GetReasonForEndingTheTreatmentOk returns a tuple with the ReasonForEndingTheTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReasonForEndingTheTreatment

`func (o *Ga4ghTreatment) HasReasonForEndingTheTreatment() bool`

HasReasonForEndingTheTreatment returns a boolean if a field has been set.

### SetReasonForEndingTheTreatment

`func (o *Ga4ghTreatment) SetReasonForEndingTheTreatment(v string)`

SetReasonForEndingTheTreatment gets a reference to the given string and assigns it to the ReasonForEndingTheTreatment field.

### GetResponseToTreatment

`func (o *Ga4ghTreatment) GetResponseToTreatment() string`

GetResponseToTreatment returns the ResponseToTreatment field if non-nil, zero value otherwise.

### GetResponseToTreatmentOk

`func (o *Ga4ghTreatment) GetResponseToTreatmentOk() (string, bool)`

GetResponseToTreatmentOk returns a tuple with the ResponseToTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseToTreatment

`func (o *Ga4ghTreatment) HasResponseToTreatment() bool`

HasResponseToTreatment returns a boolean if a field has been set.

### SetResponseToTreatment

`func (o *Ga4ghTreatment) SetResponseToTreatment(v string)`

SetResponseToTreatment gets a reference to the given string and assigns it to the ResponseToTreatment field.

### GetResponseCriteriaUsed

`func (o *Ga4ghTreatment) GetResponseCriteriaUsed() string`

GetResponseCriteriaUsed returns the ResponseCriteriaUsed field if non-nil, zero value otherwise.

### GetResponseCriteriaUsedOk

`func (o *Ga4ghTreatment) GetResponseCriteriaUsedOk() (string, bool)`

GetResponseCriteriaUsedOk returns a tuple with the ResponseCriteriaUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseCriteriaUsed

`func (o *Ga4ghTreatment) HasResponseCriteriaUsed() bool`

HasResponseCriteriaUsed returns a boolean if a field has been set.

### SetResponseCriteriaUsed

`func (o *Ga4ghTreatment) SetResponseCriteriaUsed(v string)`

SetResponseCriteriaUsed gets a reference to the given string and assigns it to the ResponseCriteriaUsed field.

### GetDateOfRecurrenceOrProgressionAfterThisTreatment

`func (o *Ga4ghTreatment) GetDateOfRecurrenceOrProgressionAfterThisTreatment() string`

GetDateOfRecurrenceOrProgressionAfterThisTreatment returns the DateOfRecurrenceOrProgressionAfterThisTreatment field if non-nil, zero value otherwise.

### GetDateOfRecurrenceOrProgressionAfterThisTreatmentOk

`func (o *Ga4ghTreatment) GetDateOfRecurrenceOrProgressionAfterThisTreatmentOk() (string, bool)`

GetDateOfRecurrenceOrProgressionAfterThisTreatmentOk returns a tuple with the DateOfRecurrenceOrProgressionAfterThisTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfRecurrenceOrProgressionAfterThisTreatment

`func (o *Ga4ghTreatment) HasDateOfRecurrenceOrProgressionAfterThisTreatment() bool`

HasDateOfRecurrenceOrProgressionAfterThisTreatment returns a boolean if a field has been set.

### SetDateOfRecurrenceOrProgressionAfterThisTreatment

`func (o *Ga4ghTreatment) SetDateOfRecurrenceOrProgressionAfterThisTreatment(v string)`

SetDateOfRecurrenceOrProgressionAfterThisTreatment gets a reference to the given string and assigns it to the DateOfRecurrenceOrProgressionAfterThisTreatment field.

### GetUnexpectedOrUnusualToxicityDuringTreatment

`func (o *Ga4ghTreatment) GetUnexpectedOrUnusualToxicityDuringTreatment() string`

GetUnexpectedOrUnusualToxicityDuringTreatment returns the UnexpectedOrUnusualToxicityDuringTreatment field if non-nil, zero value otherwise.

### GetUnexpectedOrUnusualToxicityDuringTreatmentOk

`func (o *Ga4ghTreatment) GetUnexpectedOrUnusualToxicityDuringTreatmentOk() (string, bool)`

GetUnexpectedOrUnusualToxicityDuringTreatmentOk returns a tuple with the UnexpectedOrUnusualToxicityDuringTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnexpectedOrUnusualToxicityDuringTreatment

`func (o *Ga4ghTreatment) HasUnexpectedOrUnusualToxicityDuringTreatment() bool`

HasUnexpectedOrUnusualToxicityDuringTreatment returns a boolean if a field has been set.

### SetUnexpectedOrUnusualToxicityDuringTreatment

`func (o *Ga4ghTreatment) SetUnexpectedOrUnusualToxicityDuringTreatment(v string)`

SetUnexpectedOrUnusualToxicityDuringTreatment gets a reference to the given string and assigns it to the UnexpectedOrUnusualToxicityDuringTreatment field.

### GetDiagnosisId

`func (o *Ga4ghTreatment) GetDiagnosisId() string`

GetDiagnosisId returns the DiagnosisId field if non-nil, zero value otherwise.

### GetDiagnosisIdOk

`func (o *Ga4ghTreatment) GetDiagnosisIdOk() (string, bool)`

GetDiagnosisIdOk returns a tuple with the DiagnosisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosisId

`func (o *Ga4ghTreatment) HasDiagnosisId() bool`

HasDiagnosisId returns a boolean if a field has been set.

### SetDiagnosisId

`func (o *Ga4ghTreatment) SetDiagnosisId(v string)`

SetDiagnosisId gets a reference to the given string and assigns it to the DiagnosisId field.

### GetTreatmentPlanId

`func (o *Ga4ghTreatment) GetTreatmentPlanId() string`

GetTreatmentPlanId returns the TreatmentPlanId field if non-nil, zero value otherwise.

### GetTreatmentPlanIdOk

`func (o *Ga4ghTreatment) GetTreatmentPlanIdOk() (string, bool)`

GetTreatmentPlanIdOk returns a tuple with the TreatmentPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatmentPlanId

`func (o *Ga4ghTreatment) HasTreatmentPlanId() bool`

HasTreatmentPlanId returns a boolean if a field has been set.

### SetTreatmentPlanId

`func (o *Ga4ghTreatment) SetTreatmentPlanId(v string)`

SetTreatmentPlanId gets a reference to the given string and assigns it to the TreatmentPlanId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


