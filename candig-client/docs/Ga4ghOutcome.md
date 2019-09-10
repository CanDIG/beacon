# Ga4ghOutcome

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
**PhysicalExamId** | Pointer to **string** |  | [optional] 
**DateOfAssessment** | Pointer to **string** |  | [optional] 
**DiseaseResponseOrStatus** | Pointer to **string** |  | [optional] 
**OtherResponseClassification** | Pointer to **string** |  | [optional] 
**MinimalResidualDiseaseAssessment** | Pointer to **string** |  | [optional] 
**MethodOfResponseEvaluation** | Pointer to **string** |  | [optional] 
**ResponseCriteriaUsed** | Pointer to **string** |  | [optional] 
**SummaryStage** | Pointer to **string** |  | [optional] 
**SitesOfAnyProgressionOrRecurrence** | Pointer to **string** |  | [optional] 
**VitalStatus** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **string** |  | [optional] 
**HeightUnits** | Pointer to **string** |  | [optional] 
**WeightUnits** | Pointer to **string** |  | [optional] 
**PerformanceStatus** | Pointer to **string** |  | [optional] 
**OverallSurvivalInMonths** | Pointer to **string** |  | [optional] 
**DiseaseFreeSurvivalInMonths** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghOutcome) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghOutcome) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghOutcome) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghOutcome) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghOutcome) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghOutcome) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghOutcome) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghOutcome) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghOutcome) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghOutcome) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghOutcome) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghOutcome) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghOutcome) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghOutcome) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghOutcome) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghOutcome) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghOutcome) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghOutcome) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghOutcome) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghOutcome) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghOutcome) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghOutcome) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghOutcome) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghOutcome) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghOutcome) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghOutcome) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghOutcome) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghOutcome) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetPhysicalExamId

`func (o *Ga4ghOutcome) GetPhysicalExamId() string`

GetPhysicalExamId returns the PhysicalExamId field if non-nil, zero value otherwise.

### GetPhysicalExamIdOk

`func (o *Ga4ghOutcome) GetPhysicalExamIdOk() (string, bool)`

GetPhysicalExamIdOk returns a tuple with the PhysicalExamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhysicalExamId

`func (o *Ga4ghOutcome) HasPhysicalExamId() bool`

HasPhysicalExamId returns a boolean if a field has been set.

### SetPhysicalExamId

`func (o *Ga4ghOutcome) SetPhysicalExamId(v string)`

SetPhysicalExamId gets a reference to the given string and assigns it to the PhysicalExamId field.

### GetDateOfAssessment

`func (o *Ga4ghOutcome) GetDateOfAssessment() string`

GetDateOfAssessment returns the DateOfAssessment field if non-nil, zero value otherwise.

### GetDateOfAssessmentOk

`func (o *Ga4ghOutcome) GetDateOfAssessmentOk() (string, bool)`

GetDateOfAssessmentOk returns a tuple with the DateOfAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfAssessment

`func (o *Ga4ghOutcome) HasDateOfAssessment() bool`

HasDateOfAssessment returns a boolean if a field has been set.

### SetDateOfAssessment

`func (o *Ga4ghOutcome) SetDateOfAssessment(v string)`

SetDateOfAssessment gets a reference to the given string and assigns it to the DateOfAssessment field.

### GetDiseaseResponseOrStatus

`func (o *Ga4ghOutcome) GetDiseaseResponseOrStatus() string`

GetDiseaseResponseOrStatus returns the DiseaseResponseOrStatus field if non-nil, zero value otherwise.

### GetDiseaseResponseOrStatusOk

`func (o *Ga4ghOutcome) GetDiseaseResponseOrStatusOk() (string, bool)`

GetDiseaseResponseOrStatusOk returns a tuple with the DiseaseResponseOrStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiseaseResponseOrStatus

`func (o *Ga4ghOutcome) HasDiseaseResponseOrStatus() bool`

HasDiseaseResponseOrStatus returns a boolean if a field has been set.

### SetDiseaseResponseOrStatus

`func (o *Ga4ghOutcome) SetDiseaseResponseOrStatus(v string)`

SetDiseaseResponseOrStatus gets a reference to the given string and assigns it to the DiseaseResponseOrStatus field.

### GetOtherResponseClassification

`func (o *Ga4ghOutcome) GetOtherResponseClassification() string`

GetOtherResponseClassification returns the OtherResponseClassification field if non-nil, zero value otherwise.

### GetOtherResponseClassificationOk

`func (o *Ga4ghOutcome) GetOtherResponseClassificationOk() (string, bool)`

GetOtherResponseClassificationOk returns a tuple with the OtherResponseClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherResponseClassification

`func (o *Ga4ghOutcome) HasOtherResponseClassification() bool`

HasOtherResponseClassification returns a boolean if a field has been set.

### SetOtherResponseClassification

`func (o *Ga4ghOutcome) SetOtherResponseClassification(v string)`

SetOtherResponseClassification gets a reference to the given string and assigns it to the OtherResponseClassification field.

### GetMinimalResidualDiseaseAssessment

`func (o *Ga4ghOutcome) GetMinimalResidualDiseaseAssessment() string`

GetMinimalResidualDiseaseAssessment returns the MinimalResidualDiseaseAssessment field if non-nil, zero value otherwise.

### GetMinimalResidualDiseaseAssessmentOk

`func (o *Ga4ghOutcome) GetMinimalResidualDiseaseAssessmentOk() (string, bool)`

GetMinimalResidualDiseaseAssessmentOk returns a tuple with the MinimalResidualDiseaseAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimalResidualDiseaseAssessment

`func (o *Ga4ghOutcome) HasMinimalResidualDiseaseAssessment() bool`

HasMinimalResidualDiseaseAssessment returns a boolean if a field has been set.

### SetMinimalResidualDiseaseAssessment

`func (o *Ga4ghOutcome) SetMinimalResidualDiseaseAssessment(v string)`

SetMinimalResidualDiseaseAssessment gets a reference to the given string and assigns it to the MinimalResidualDiseaseAssessment field.

### GetMethodOfResponseEvaluation

`func (o *Ga4ghOutcome) GetMethodOfResponseEvaluation() string`

GetMethodOfResponseEvaluation returns the MethodOfResponseEvaluation field if non-nil, zero value otherwise.

### GetMethodOfResponseEvaluationOk

`func (o *Ga4ghOutcome) GetMethodOfResponseEvaluationOk() (string, bool)`

GetMethodOfResponseEvaluationOk returns a tuple with the MethodOfResponseEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMethodOfResponseEvaluation

`func (o *Ga4ghOutcome) HasMethodOfResponseEvaluation() bool`

HasMethodOfResponseEvaluation returns a boolean if a field has been set.

### SetMethodOfResponseEvaluation

`func (o *Ga4ghOutcome) SetMethodOfResponseEvaluation(v string)`

SetMethodOfResponseEvaluation gets a reference to the given string and assigns it to the MethodOfResponseEvaluation field.

### GetResponseCriteriaUsed

`func (o *Ga4ghOutcome) GetResponseCriteriaUsed() string`

GetResponseCriteriaUsed returns the ResponseCriteriaUsed field if non-nil, zero value otherwise.

### GetResponseCriteriaUsedOk

`func (o *Ga4ghOutcome) GetResponseCriteriaUsedOk() (string, bool)`

GetResponseCriteriaUsedOk returns a tuple with the ResponseCriteriaUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseCriteriaUsed

`func (o *Ga4ghOutcome) HasResponseCriteriaUsed() bool`

HasResponseCriteriaUsed returns a boolean if a field has been set.

### SetResponseCriteriaUsed

`func (o *Ga4ghOutcome) SetResponseCriteriaUsed(v string)`

SetResponseCriteriaUsed gets a reference to the given string and assigns it to the ResponseCriteriaUsed field.

### GetSummaryStage

`func (o *Ga4ghOutcome) GetSummaryStage() string`

GetSummaryStage returns the SummaryStage field if non-nil, zero value otherwise.

### GetSummaryStageOk

`func (o *Ga4ghOutcome) GetSummaryStageOk() (string, bool)`

GetSummaryStageOk returns a tuple with the SummaryStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSummaryStage

`func (o *Ga4ghOutcome) HasSummaryStage() bool`

HasSummaryStage returns a boolean if a field has been set.

### SetSummaryStage

`func (o *Ga4ghOutcome) SetSummaryStage(v string)`

SetSummaryStage gets a reference to the given string and assigns it to the SummaryStage field.

### GetSitesOfAnyProgressionOrRecurrence

`func (o *Ga4ghOutcome) GetSitesOfAnyProgressionOrRecurrence() string`

GetSitesOfAnyProgressionOrRecurrence returns the SitesOfAnyProgressionOrRecurrence field if non-nil, zero value otherwise.

### GetSitesOfAnyProgressionOrRecurrenceOk

`func (o *Ga4ghOutcome) GetSitesOfAnyProgressionOrRecurrenceOk() (string, bool)`

GetSitesOfAnyProgressionOrRecurrenceOk returns a tuple with the SitesOfAnyProgressionOrRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSitesOfAnyProgressionOrRecurrence

`func (o *Ga4ghOutcome) HasSitesOfAnyProgressionOrRecurrence() bool`

HasSitesOfAnyProgressionOrRecurrence returns a boolean if a field has been set.

### SetSitesOfAnyProgressionOrRecurrence

`func (o *Ga4ghOutcome) SetSitesOfAnyProgressionOrRecurrence(v string)`

SetSitesOfAnyProgressionOrRecurrence gets a reference to the given string and assigns it to the SitesOfAnyProgressionOrRecurrence field.

### GetVitalStatus

`func (o *Ga4ghOutcome) GetVitalStatus() string`

GetVitalStatus returns the VitalStatus field if non-nil, zero value otherwise.

### GetVitalStatusOk

`func (o *Ga4ghOutcome) GetVitalStatusOk() (string, bool)`

GetVitalStatusOk returns a tuple with the VitalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVitalStatus

`func (o *Ga4ghOutcome) HasVitalStatus() bool`

HasVitalStatus returns a boolean if a field has been set.

### SetVitalStatus

`func (o *Ga4ghOutcome) SetVitalStatus(v string)`

SetVitalStatus gets a reference to the given string and assigns it to the VitalStatus field.

### GetHeight

`func (o *Ga4ghOutcome) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Ga4ghOutcome) GetHeightOk() (string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeight

`func (o *Ga4ghOutcome) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeight

`func (o *Ga4ghOutcome) SetHeight(v string)`

SetHeight gets a reference to the given string and assigns it to the Height field.

### GetWeight

`func (o *Ga4ghOutcome) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Ga4ghOutcome) GetWeightOk() (string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeight

`func (o *Ga4ghOutcome) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeight

`func (o *Ga4ghOutcome) SetWeight(v string)`

SetWeight gets a reference to the given string and assigns it to the Weight field.

### GetHeightUnits

`func (o *Ga4ghOutcome) GetHeightUnits() string`

GetHeightUnits returns the HeightUnits field if non-nil, zero value otherwise.

### GetHeightUnitsOk

`func (o *Ga4ghOutcome) GetHeightUnitsOk() (string, bool)`

GetHeightUnitsOk returns a tuple with the HeightUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeightUnits

`func (o *Ga4ghOutcome) HasHeightUnits() bool`

HasHeightUnits returns a boolean if a field has been set.

### SetHeightUnits

`func (o *Ga4ghOutcome) SetHeightUnits(v string)`

SetHeightUnits gets a reference to the given string and assigns it to the HeightUnits field.

### GetWeightUnits

`func (o *Ga4ghOutcome) GetWeightUnits() string`

GetWeightUnits returns the WeightUnits field if non-nil, zero value otherwise.

### GetWeightUnitsOk

`func (o *Ga4ghOutcome) GetWeightUnitsOk() (string, bool)`

GetWeightUnitsOk returns a tuple with the WeightUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeightUnits

`func (o *Ga4ghOutcome) HasWeightUnits() bool`

HasWeightUnits returns a boolean if a field has been set.

### SetWeightUnits

`func (o *Ga4ghOutcome) SetWeightUnits(v string)`

SetWeightUnits gets a reference to the given string and assigns it to the WeightUnits field.

### GetPerformanceStatus

`func (o *Ga4ghOutcome) GetPerformanceStatus() string`

GetPerformanceStatus returns the PerformanceStatus field if non-nil, zero value otherwise.

### GetPerformanceStatusOk

`func (o *Ga4ghOutcome) GetPerformanceStatusOk() (string, bool)`

GetPerformanceStatusOk returns a tuple with the PerformanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPerformanceStatus

`func (o *Ga4ghOutcome) HasPerformanceStatus() bool`

HasPerformanceStatus returns a boolean if a field has been set.

### SetPerformanceStatus

`func (o *Ga4ghOutcome) SetPerformanceStatus(v string)`

SetPerformanceStatus gets a reference to the given string and assigns it to the PerformanceStatus field.

### GetOverallSurvivalInMonths

`func (o *Ga4ghOutcome) GetOverallSurvivalInMonths() string`

GetOverallSurvivalInMonths returns the OverallSurvivalInMonths field if non-nil, zero value otherwise.

### GetOverallSurvivalInMonthsOk

`func (o *Ga4ghOutcome) GetOverallSurvivalInMonthsOk() (string, bool)`

GetOverallSurvivalInMonthsOk returns a tuple with the OverallSurvivalInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverallSurvivalInMonths

`func (o *Ga4ghOutcome) HasOverallSurvivalInMonths() bool`

HasOverallSurvivalInMonths returns a boolean if a field has been set.

### SetOverallSurvivalInMonths

`func (o *Ga4ghOutcome) SetOverallSurvivalInMonths(v string)`

SetOverallSurvivalInMonths gets a reference to the given string and assigns it to the OverallSurvivalInMonths field.

### GetDiseaseFreeSurvivalInMonths

`func (o *Ga4ghOutcome) GetDiseaseFreeSurvivalInMonths() string`

GetDiseaseFreeSurvivalInMonths returns the DiseaseFreeSurvivalInMonths field if non-nil, zero value otherwise.

### GetDiseaseFreeSurvivalInMonthsOk

`func (o *Ga4ghOutcome) GetDiseaseFreeSurvivalInMonthsOk() (string, bool)`

GetDiseaseFreeSurvivalInMonthsOk returns a tuple with the DiseaseFreeSurvivalInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiseaseFreeSurvivalInMonths

`func (o *Ga4ghOutcome) HasDiseaseFreeSurvivalInMonths() bool`

HasDiseaseFreeSurvivalInMonths returns a boolean if a field has been set.

### SetDiseaseFreeSurvivalInMonths

`func (o *Ga4ghOutcome) SetDiseaseFreeSurvivalInMonths(v string)`

SetDiseaseFreeSurvivalInMonths gets a reference to the given string and assigns it to the DiseaseFreeSurvivalInMonths field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


