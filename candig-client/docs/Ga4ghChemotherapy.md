# Ga4ghChemotherapy

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
**StartDate** | Pointer to **string** |  | [optional] 
**StopDate** | Pointer to **string** |  | [optional] 
**SystematicTherapyAgentName** | Pointer to **string** |  | [optional] 
**Route** | Pointer to **string** |  | [optional] 
**Dose** | Pointer to **string** |  | [optional] 
**DoseFrequency** | Pointer to **string** |  | [optional] 
**DoseUnit** | Pointer to **string** |  | [optional] 
**DaysPerCycle** | Pointer to **string** |  | [optional] 
**NumberOfCycle** | Pointer to **string** |  | [optional] 
**TreatmentIntent** | Pointer to **string** |  | [optional] 
**TreatingCentreName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ProtocolCode** | Pointer to **string** |  | [optional] 
**RecordingDate** | Pointer to **string** |  | [optional] 
**TreatmentPlanId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghChemotherapy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghChemotherapy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghChemotherapy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghChemotherapy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghChemotherapy) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghChemotherapy) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghChemotherapy) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghChemotherapy) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghChemotherapy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghChemotherapy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghChemotherapy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghChemotherapy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghChemotherapy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghChemotherapy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghChemotherapy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghChemotherapy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghChemotherapy) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghChemotherapy) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghChemotherapy) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghChemotherapy) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghChemotherapy) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghChemotherapy) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghChemotherapy) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghChemotherapy) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghChemotherapy) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghChemotherapy) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghChemotherapy) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghChemotherapy) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetCourseNumber

`func (o *Ga4ghChemotherapy) GetCourseNumber() string`

GetCourseNumber returns the CourseNumber field if non-nil, zero value otherwise.

### GetCourseNumberOk

`func (o *Ga4ghChemotherapy) GetCourseNumberOk() (string, bool)`

GetCourseNumberOk returns a tuple with the CourseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCourseNumber

`func (o *Ga4ghChemotherapy) HasCourseNumber() bool`

HasCourseNumber returns a boolean if a field has been set.

### SetCourseNumber

`func (o *Ga4ghChemotherapy) SetCourseNumber(v string)`

SetCourseNumber gets a reference to the given string and assigns it to the CourseNumber field.

### GetStartDate

`func (o *Ga4ghChemotherapy) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Ga4ghChemotherapy) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *Ga4ghChemotherapy) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *Ga4ghChemotherapy) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetStopDate

`func (o *Ga4ghChemotherapy) GetStopDate() string`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *Ga4ghChemotherapy) GetStopDateOk() (string, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStopDate

`func (o *Ga4ghChemotherapy) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.

### SetStopDate

`func (o *Ga4ghChemotherapy) SetStopDate(v string)`

SetStopDate gets a reference to the given string and assigns it to the StopDate field.

### GetSystematicTherapyAgentName

`func (o *Ga4ghChemotherapy) GetSystematicTherapyAgentName() string`

GetSystematicTherapyAgentName returns the SystematicTherapyAgentName field if non-nil, zero value otherwise.

### GetSystematicTherapyAgentNameOk

`func (o *Ga4ghChemotherapy) GetSystematicTherapyAgentNameOk() (string, bool)`

GetSystematicTherapyAgentNameOk returns a tuple with the SystematicTherapyAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystematicTherapyAgentName

`func (o *Ga4ghChemotherapy) HasSystematicTherapyAgentName() bool`

HasSystematicTherapyAgentName returns a boolean if a field has been set.

### SetSystematicTherapyAgentName

`func (o *Ga4ghChemotherapy) SetSystematicTherapyAgentName(v string)`

SetSystematicTherapyAgentName gets a reference to the given string and assigns it to the SystematicTherapyAgentName field.

### GetRoute

`func (o *Ga4ghChemotherapy) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *Ga4ghChemotherapy) GetRouteOk() (string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoute

`func (o *Ga4ghChemotherapy) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### SetRoute

`func (o *Ga4ghChemotherapy) SetRoute(v string)`

SetRoute gets a reference to the given string and assigns it to the Route field.

### GetDose

`func (o *Ga4ghChemotherapy) GetDose() string`

GetDose returns the Dose field if non-nil, zero value otherwise.

### GetDoseOk

`func (o *Ga4ghChemotherapy) GetDoseOk() (string, bool)`

GetDoseOk returns a tuple with the Dose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDose

`func (o *Ga4ghChemotherapy) HasDose() bool`

HasDose returns a boolean if a field has been set.

### SetDose

`func (o *Ga4ghChemotherapy) SetDose(v string)`

SetDose gets a reference to the given string and assigns it to the Dose field.

### GetDoseFrequency

`func (o *Ga4ghChemotherapy) GetDoseFrequency() string`

GetDoseFrequency returns the DoseFrequency field if non-nil, zero value otherwise.

### GetDoseFrequencyOk

`func (o *Ga4ghChemotherapy) GetDoseFrequencyOk() (string, bool)`

GetDoseFrequencyOk returns a tuple with the DoseFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDoseFrequency

`func (o *Ga4ghChemotherapy) HasDoseFrequency() bool`

HasDoseFrequency returns a boolean if a field has been set.

### SetDoseFrequency

`func (o *Ga4ghChemotherapy) SetDoseFrequency(v string)`

SetDoseFrequency gets a reference to the given string and assigns it to the DoseFrequency field.

### GetDoseUnit

`func (o *Ga4ghChemotherapy) GetDoseUnit() string`

GetDoseUnit returns the DoseUnit field if non-nil, zero value otherwise.

### GetDoseUnitOk

`func (o *Ga4ghChemotherapy) GetDoseUnitOk() (string, bool)`

GetDoseUnitOk returns a tuple with the DoseUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDoseUnit

`func (o *Ga4ghChemotherapy) HasDoseUnit() bool`

HasDoseUnit returns a boolean if a field has been set.

### SetDoseUnit

`func (o *Ga4ghChemotherapy) SetDoseUnit(v string)`

SetDoseUnit gets a reference to the given string and assigns it to the DoseUnit field.

### GetDaysPerCycle

`func (o *Ga4ghChemotherapy) GetDaysPerCycle() string`

GetDaysPerCycle returns the DaysPerCycle field if non-nil, zero value otherwise.

### GetDaysPerCycleOk

`func (o *Ga4ghChemotherapy) GetDaysPerCycleOk() (string, bool)`

GetDaysPerCycleOk returns a tuple with the DaysPerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaysPerCycle

`func (o *Ga4ghChemotherapy) HasDaysPerCycle() bool`

HasDaysPerCycle returns a boolean if a field has been set.

### SetDaysPerCycle

`func (o *Ga4ghChemotherapy) SetDaysPerCycle(v string)`

SetDaysPerCycle gets a reference to the given string and assigns it to the DaysPerCycle field.

### GetNumberOfCycle

`func (o *Ga4ghChemotherapy) GetNumberOfCycle() string`

GetNumberOfCycle returns the NumberOfCycle field if non-nil, zero value otherwise.

### GetNumberOfCycleOk

`func (o *Ga4ghChemotherapy) GetNumberOfCycleOk() (string, bool)`

GetNumberOfCycleOk returns a tuple with the NumberOfCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfCycle

`func (o *Ga4ghChemotherapy) HasNumberOfCycle() bool`

HasNumberOfCycle returns a boolean if a field has been set.

### SetNumberOfCycle

`func (o *Ga4ghChemotherapy) SetNumberOfCycle(v string)`

SetNumberOfCycle gets a reference to the given string and assigns it to the NumberOfCycle field.

### GetTreatmentIntent

`func (o *Ga4ghChemotherapy) GetTreatmentIntent() string`

GetTreatmentIntent returns the TreatmentIntent field if non-nil, zero value otherwise.

### GetTreatmentIntentOk

`func (o *Ga4ghChemotherapy) GetTreatmentIntentOk() (string, bool)`

GetTreatmentIntentOk returns a tuple with the TreatmentIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatmentIntent

`func (o *Ga4ghChemotherapy) HasTreatmentIntent() bool`

HasTreatmentIntent returns a boolean if a field has been set.

### SetTreatmentIntent

`func (o *Ga4ghChemotherapy) SetTreatmentIntent(v string)`

SetTreatmentIntent gets a reference to the given string and assigns it to the TreatmentIntent field.

### GetTreatingCentreName

`func (o *Ga4ghChemotherapy) GetTreatingCentreName() string`

GetTreatingCentreName returns the TreatingCentreName field if non-nil, zero value otherwise.

### GetTreatingCentreNameOk

`func (o *Ga4ghChemotherapy) GetTreatingCentreNameOk() (string, bool)`

GetTreatingCentreNameOk returns a tuple with the TreatingCentreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatingCentreName

`func (o *Ga4ghChemotherapy) HasTreatingCentreName() bool`

HasTreatingCentreName returns a boolean if a field has been set.

### SetTreatingCentreName

`func (o *Ga4ghChemotherapy) SetTreatingCentreName(v string)`

SetTreatingCentreName gets a reference to the given string and assigns it to the TreatingCentreName field.

### GetType

`func (o *Ga4ghChemotherapy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ga4ghChemotherapy) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Ga4ghChemotherapy) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Ga4ghChemotherapy) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetProtocolCode

`func (o *Ga4ghChemotherapy) GetProtocolCode() string`

GetProtocolCode returns the ProtocolCode field if non-nil, zero value otherwise.

### GetProtocolCodeOk

`func (o *Ga4ghChemotherapy) GetProtocolCodeOk() (string, bool)`

GetProtocolCodeOk returns a tuple with the ProtocolCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtocolCode

`func (o *Ga4ghChemotherapy) HasProtocolCode() bool`

HasProtocolCode returns a boolean if a field has been set.

### SetProtocolCode

`func (o *Ga4ghChemotherapy) SetProtocolCode(v string)`

SetProtocolCode gets a reference to the given string and assigns it to the ProtocolCode field.

### GetRecordingDate

`func (o *Ga4ghChemotherapy) GetRecordingDate() string`

GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.

### GetRecordingDateOk

`func (o *Ga4ghChemotherapy) GetRecordingDateOk() (string, bool)`

GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordingDate

`func (o *Ga4ghChemotherapy) HasRecordingDate() bool`

HasRecordingDate returns a boolean if a field has been set.

### SetRecordingDate

`func (o *Ga4ghChemotherapy) SetRecordingDate(v string)`

SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.

### GetTreatmentPlanId

`func (o *Ga4ghChemotherapy) GetTreatmentPlanId() string`

GetTreatmentPlanId returns the TreatmentPlanId field if non-nil, zero value otherwise.

### GetTreatmentPlanIdOk

`func (o *Ga4ghChemotherapy) GetTreatmentPlanIdOk() (string, bool)`

GetTreatmentPlanIdOk returns a tuple with the TreatmentPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatmentPlanId

`func (o *Ga4ghChemotherapy) HasTreatmentPlanId() bool`

HasTreatmentPlanId returns a boolean if a field has been set.

### SetTreatmentPlanId

`func (o *Ga4ghChemotherapy) SetTreatmentPlanId(v string)`

SetTreatmentPlanId gets a reference to the given string and assigns it to the TreatmentPlanId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


