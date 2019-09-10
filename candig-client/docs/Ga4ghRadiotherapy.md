# Ga4ghRadiotherapy

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
**TherapeuticModality** | Pointer to **string** |  | [optional] 
**Baseline** | Pointer to **string** |  | [optional] 
**TestResult** | Pointer to **string** |  | [optional] 
**TestResultStd** | Pointer to **string** |  | [optional] 
**TreatingCentreName** | Pointer to **string** |  | [optional] 
**StartIntervalRad** | Pointer to **string** |  | [optional] 
**StartIntervalRadRaw** | Pointer to **string** |  | [optional] 
**RecordingDate** | Pointer to **string** |  | [optional] 
**AdjacentFields** | Pointer to **string** |  | [optional] 
**AdjacentFractions** | Pointer to **string** |  | [optional] 
**Complete** | Pointer to **string** |  | [optional] 
**BrachytherapyDose** | Pointer to **string** |  | [optional] 
**RadiotherapyDose** | Pointer to **string** |  | [optional] 
**SiteNumber** | Pointer to **string** |  | [optional] 
**Technique** | Pointer to **string** |  | [optional] 
**TreatedRegion** | Pointer to **string** |  | [optional] 
**TreatmentPlanId** | Pointer to **string** |  | [optional] 
**RadiationType** | Pointer to **string** |  | [optional] 
**RadiationSite** | Pointer to **string** |  | [optional] 
**TotalDose** | Pointer to **string** |  | [optional] 
**BoostSite** | Pointer to **string** |  | [optional] 
**BoostDose** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghRadiotherapy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghRadiotherapy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghRadiotherapy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghRadiotherapy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghRadiotherapy) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghRadiotherapy) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghRadiotherapy) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghRadiotherapy) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghRadiotherapy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghRadiotherapy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghRadiotherapy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghRadiotherapy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghRadiotherapy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghRadiotherapy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghRadiotherapy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghRadiotherapy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghRadiotherapy) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghRadiotherapy) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghRadiotherapy) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghRadiotherapy) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghRadiotherapy) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghRadiotherapy) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghRadiotherapy) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghRadiotherapy) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghRadiotherapy) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghRadiotherapy) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghRadiotherapy) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghRadiotherapy) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetCourseNumber

`func (o *Ga4ghRadiotherapy) GetCourseNumber() string`

GetCourseNumber returns the CourseNumber field if non-nil, zero value otherwise.

### GetCourseNumberOk

`func (o *Ga4ghRadiotherapy) GetCourseNumberOk() (string, bool)`

GetCourseNumberOk returns a tuple with the CourseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCourseNumber

`func (o *Ga4ghRadiotherapy) HasCourseNumber() bool`

HasCourseNumber returns a boolean if a field has been set.

### SetCourseNumber

`func (o *Ga4ghRadiotherapy) SetCourseNumber(v string)`

SetCourseNumber gets a reference to the given string and assigns it to the CourseNumber field.

### GetStartDate

`func (o *Ga4ghRadiotherapy) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Ga4ghRadiotherapy) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *Ga4ghRadiotherapy) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *Ga4ghRadiotherapy) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetStopDate

`func (o *Ga4ghRadiotherapy) GetStopDate() string`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *Ga4ghRadiotherapy) GetStopDateOk() (string, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStopDate

`func (o *Ga4ghRadiotherapy) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.

### SetStopDate

`func (o *Ga4ghRadiotherapy) SetStopDate(v string)`

SetStopDate gets a reference to the given string and assigns it to the StopDate field.

### GetTherapeuticModality

`func (o *Ga4ghRadiotherapy) GetTherapeuticModality() string`

GetTherapeuticModality returns the TherapeuticModality field if non-nil, zero value otherwise.

### GetTherapeuticModalityOk

`func (o *Ga4ghRadiotherapy) GetTherapeuticModalityOk() (string, bool)`

GetTherapeuticModalityOk returns a tuple with the TherapeuticModality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTherapeuticModality

`func (o *Ga4ghRadiotherapy) HasTherapeuticModality() bool`

HasTherapeuticModality returns a boolean if a field has been set.

### SetTherapeuticModality

`func (o *Ga4ghRadiotherapy) SetTherapeuticModality(v string)`

SetTherapeuticModality gets a reference to the given string and assigns it to the TherapeuticModality field.

### GetBaseline

`func (o *Ga4ghRadiotherapy) GetBaseline() string`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *Ga4ghRadiotherapy) GetBaselineOk() (string, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBaseline

`func (o *Ga4ghRadiotherapy) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### SetBaseline

`func (o *Ga4ghRadiotherapy) SetBaseline(v string)`

SetBaseline gets a reference to the given string and assigns it to the Baseline field.

### GetTestResult

`func (o *Ga4ghRadiotherapy) GetTestResult() string`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *Ga4ghRadiotherapy) GetTestResultOk() (string, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTestResult

`func (o *Ga4ghRadiotherapy) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.

### SetTestResult

`func (o *Ga4ghRadiotherapy) SetTestResult(v string)`

SetTestResult gets a reference to the given string and assigns it to the TestResult field.

### GetTestResultStd

`func (o *Ga4ghRadiotherapy) GetTestResultStd() string`

GetTestResultStd returns the TestResultStd field if non-nil, zero value otherwise.

### GetTestResultStdOk

`func (o *Ga4ghRadiotherapy) GetTestResultStdOk() (string, bool)`

GetTestResultStdOk returns a tuple with the TestResultStd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTestResultStd

`func (o *Ga4ghRadiotherapy) HasTestResultStd() bool`

HasTestResultStd returns a boolean if a field has been set.

### SetTestResultStd

`func (o *Ga4ghRadiotherapy) SetTestResultStd(v string)`

SetTestResultStd gets a reference to the given string and assigns it to the TestResultStd field.

### GetTreatingCentreName

`func (o *Ga4ghRadiotherapy) GetTreatingCentreName() string`

GetTreatingCentreName returns the TreatingCentreName field if non-nil, zero value otherwise.

### GetTreatingCentreNameOk

`func (o *Ga4ghRadiotherapy) GetTreatingCentreNameOk() (string, bool)`

GetTreatingCentreNameOk returns a tuple with the TreatingCentreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatingCentreName

`func (o *Ga4ghRadiotherapy) HasTreatingCentreName() bool`

HasTreatingCentreName returns a boolean if a field has been set.

### SetTreatingCentreName

`func (o *Ga4ghRadiotherapy) SetTreatingCentreName(v string)`

SetTreatingCentreName gets a reference to the given string and assigns it to the TreatingCentreName field.

### GetStartIntervalRad

`func (o *Ga4ghRadiotherapy) GetStartIntervalRad() string`

GetStartIntervalRad returns the StartIntervalRad field if non-nil, zero value otherwise.

### GetStartIntervalRadOk

`func (o *Ga4ghRadiotherapy) GetStartIntervalRadOk() (string, bool)`

GetStartIntervalRadOk returns a tuple with the StartIntervalRad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartIntervalRad

`func (o *Ga4ghRadiotherapy) HasStartIntervalRad() bool`

HasStartIntervalRad returns a boolean if a field has been set.

### SetStartIntervalRad

`func (o *Ga4ghRadiotherapy) SetStartIntervalRad(v string)`

SetStartIntervalRad gets a reference to the given string and assigns it to the StartIntervalRad field.

### GetStartIntervalRadRaw

`func (o *Ga4ghRadiotherapy) GetStartIntervalRadRaw() string`

GetStartIntervalRadRaw returns the StartIntervalRadRaw field if non-nil, zero value otherwise.

### GetStartIntervalRadRawOk

`func (o *Ga4ghRadiotherapy) GetStartIntervalRadRawOk() (string, bool)`

GetStartIntervalRadRawOk returns a tuple with the StartIntervalRadRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartIntervalRadRaw

`func (o *Ga4ghRadiotherapy) HasStartIntervalRadRaw() bool`

HasStartIntervalRadRaw returns a boolean if a field has been set.

### SetStartIntervalRadRaw

`func (o *Ga4ghRadiotherapy) SetStartIntervalRadRaw(v string)`

SetStartIntervalRadRaw gets a reference to the given string and assigns it to the StartIntervalRadRaw field.

### GetRecordingDate

`func (o *Ga4ghRadiotherapy) GetRecordingDate() string`

GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.

### GetRecordingDateOk

`func (o *Ga4ghRadiotherapy) GetRecordingDateOk() (string, bool)`

GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordingDate

`func (o *Ga4ghRadiotherapy) HasRecordingDate() bool`

HasRecordingDate returns a boolean if a field has been set.

### SetRecordingDate

`func (o *Ga4ghRadiotherapy) SetRecordingDate(v string)`

SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.

### GetAdjacentFields

`func (o *Ga4ghRadiotherapy) GetAdjacentFields() string`

GetAdjacentFields returns the AdjacentFields field if non-nil, zero value otherwise.

### GetAdjacentFieldsOk

`func (o *Ga4ghRadiotherapy) GetAdjacentFieldsOk() (string, bool)`

GetAdjacentFieldsOk returns a tuple with the AdjacentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdjacentFields

`func (o *Ga4ghRadiotherapy) HasAdjacentFields() bool`

HasAdjacentFields returns a boolean if a field has been set.

### SetAdjacentFields

`func (o *Ga4ghRadiotherapy) SetAdjacentFields(v string)`

SetAdjacentFields gets a reference to the given string and assigns it to the AdjacentFields field.

### GetAdjacentFractions

`func (o *Ga4ghRadiotherapy) GetAdjacentFractions() string`

GetAdjacentFractions returns the AdjacentFractions field if non-nil, zero value otherwise.

### GetAdjacentFractionsOk

`func (o *Ga4ghRadiotherapy) GetAdjacentFractionsOk() (string, bool)`

GetAdjacentFractionsOk returns a tuple with the AdjacentFractions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdjacentFractions

`func (o *Ga4ghRadiotherapy) HasAdjacentFractions() bool`

HasAdjacentFractions returns a boolean if a field has been set.

### SetAdjacentFractions

`func (o *Ga4ghRadiotherapy) SetAdjacentFractions(v string)`

SetAdjacentFractions gets a reference to the given string and assigns it to the AdjacentFractions field.

### GetComplete

`func (o *Ga4ghRadiotherapy) GetComplete() string`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *Ga4ghRadiotherapy) GetCompleteOk() (string, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplete

`func (o *Ga4ghRadiotherapy) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### SetComplete

`func (o *Ga4ghRadiotherapy) SetComplete(v string)`

SetComplete gets a reference to the given string and assigns it to the Complete field.

### GetBrachytherapyDose

`func (o *Ga4ghRadiotherapy) GetBrachytherapyDose() string`

GetBrachytherapyDose returns the BrachytherapyDose field if non-nil, zero value otherwise.

### GetBrachytherapyDoseOk

`func (o *Ga4ghRadiotherapy) GetBrachytherapyDoseOk() (string, bool)`

GetBrachytherapyDoseOk returns a tuple with the BrachytherapyDose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrachytherapyDose

`func (o *Ga4ghRadiotherapy) HasBrachytherapyDose() bool`

HasBrachytherapyDose returns a boolean if a field has been set.

### SetBrachytherapyDose

`func (o *Ga4ghRadiotherapy) SetBrachytherapyDose(v string)`

SetBrachytherapyDose gets a reference to the given string and assigns it to the BrachytherapyDose field.

### GetRadiotherapyDose

`func (o *Ga4ghRadiotherapy) GetRadiotherapyDose() string`

GetRadiotherapyDose returns the RadiotherapyDose field if non-nil, zero value otherwise.

### GetRadiotherapyDoseOk

`func (o *Ga4ghRadiotherapy) GetRadiotherapyDoseOk() (string, bool)`

GetRadiotherapyDoseOk returns a tuple with the RadiotherapyDose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRadiotherapyDose

`func (o *Ga4ghRadiotherapy) HasRadiotherapyDose() bool`

HasRadiotherapyDose returns a boolean if a field has been set.

### SetRadiotherapyDose

`func (o *Ga4ghRadiotherapy) SetRadiotherapyDose(v string)`

SetRadiotherapyDose gets a reference to the given string and assigns it to the RadiotherapyDose field.

### GetSiteNumber

`func (o *Ga4ghRadiotherapy) GetSiteNumber() string`

GetSiteNumber returns the SiteNumber field if non-nil, zero value otherwise.

### GetSiteNumberOk

`func (o *Ga4ghRadiotherapy) GetSiteNumberOk() (string, bool)`

GetSiteNumberOk returns a tuple with the SiteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteNumber

`func (o *Ga4ghRadiotherapy) HasSiteNumber() bool`

HasSiteNumber returns a boolean if a field has been set.

### SetSiteNumber

`func (o *Ga4ghRadiotherapy) SetSiteNumber(v string)`

SetSiteNumber gets a reference to the given string and assigns it to the SiteNumber field.

### GetTechnique

`func (o *Ga4ghRadiotherapy) GetTechnique() string`

GetTechnique returns the Technique field if non-nil, zero value otherwise.

### GetTechniqueOk

`func (o *Ga4ghRadiotherapy) GetTechniqueOk() (string, bool)`

GetTechniqueOk returns a tuple with the Technique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTechnique

`func (o *Ga4ghRadiotherapy) HasTechnique() bool`

HasTechnique returns a boolean if a field has been set.

### SetTechnique

`func (o *Ga4ghRadiotherapy) SetTechnique(v string)`

SetTechnique gets a reference to the given string and assigns it to the Technique field.

### GetTreatedRegion

`func (o *Ga4ghRadiotherapy) GetTreatedRegion() string`

GetTreatedRegion returns the TreatedRegion field if non-nil, zero value otherwise.

### GetTreatedRegionOk

`func (o *Ga4ghRadiotherapy) GetTreatedRegionOk() (string, bool)`

GetTreatedRegionOk returns a tuple with the TreatedRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatedRegion

`func (o *Ga4ghRadiotherapy) HasTreatedRegion() bool`

HasTreatedRegion returns a boolean if a field has been set.

### SetTreatedRegion

`func (o *Ga4ghRadiotherapy) SetTreatedRegion(v string)`

SetTreatedRegion gets a reference to the given string and assigns it to the TreatedRegion field.

### GetTreatmentPlanId

`func (o *Ga4ghRadiotherapy) GetTreatmentPlanId() string`

GetTreatmentPlanId returns the TreatmentPlanId field if non-nil, zero value otherwise.

### GetTreatmentPlanIdOk

`func (o *Ga4ghRadiotherapy) GetTreatmentPlanIdOk() (string, bool)`

GetTreatmentPlanIdOk returns a tuple with the TreatmentPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatmentPlanId

`func (o *Ga4ghRadiotherapy) HasTreatmentPlanId() bool`

HasTreatmentPlanId returns a boolean if a field has been set.

### SetTreatmentPlanId

`func (o *Ga4ghRadiotherapy) SetTreatmentPlanId(v string)`

SetTreatmentPlanId gets a reference to the given string and assigns it to the TreatmentPlanId field.

### GetRadiationType

`func (o *Ga4ghRadiotherapy) GetRadiationType() string`

GetRadiationType returns the RadiationType field if non-nil, zero value otherwise.

### GetRadiationTypeOk

`func (o *Ga4ghRadiotherapy) GetRadiationTypeOk() (string, bool)`

GetRadiationTypeOk returns a tuple with the RadiationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRadiationType

`func (o *Ga4ghRadiotherapy) HasRadiationType() bool`

HasRadiationType returns a boolean if a field has been set.

### SetRadiationType

`func (o *Ga4ghRadiotherapy) SetRadiationType(v string)`

SetRadiationType gets a reference to the given string and assigns it to the RadiationType field.

### GetRadiationSite

`func (o *Ga4ghRadiotherapy) GetRadiationSite() string`

GetRadiationSite returns the RadiationSite field if non-nil, zero value otherwise.

### GetRadiationSiteOk

`func (o *Ga4ghRadiotherapy) GetRadiationSiteOk() (string, bool)`

GetRadiationSiteOk returns a tuple with the RadiationSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRadiationSite

`func (o *Ga4ghRadiotherapy) HasRadiationSite() bool`

HasRadiationSite returns a boolean if a field has been set.

### SetRadiationSite

`func (o *Ga4ghRadiotherapy) SetRadiationSite(v string)`

SetRadiationSite gets a reference to the given string and assigns it to the RadiationSite field.

### GetTotalDose

`func (o *Ga4ghRadiotherapy) GetTotalDose() string`

GetTotalDose returns the TotalDose field if non-nil, zero value otherwise.

### GetTotalDoseOk

`func (o *Ga4ghRadiotherapy) GetTotalDoseOk() (string, bool)`

GetTotalDoseOk returns a tuple with the TotalDose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDose

`func (o *Ga4ghRadiotherapy) HasTotalDose() bool`

HasTotalDose returns a boolean if a field has been set.

### SetTotalDose

`func (o *Ga4ghRadiotherapy) SetTotalDose(v string)`

SetTotalDose gets a reference to the given string and assigns it to the TotalDose field.

### GetBoostSite

`func (o *Ga4ghRadiotherapy) GetBoostSite() string`

GetBoostSite returns the BoostSite field if non-nil, zero value otherwise.

### GetBoostSiteOk

`func (o *Ga4ghRadiotherapy) GetBoostSiteOk() (string, bool)`

GetBoostSiteOk returns a tuple with the BoostSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBoostSite

`func (o *Ga4ghRadiotherapy) HasBoostSite() bool`

HasBoostSite returns a boolean if a field has been set.

### SetBoostSite

`func (o *Ga4ghRadiotherapy) SetBoostSite(v string)`

SetBoostSite gets a reference to the given string and assigns it to the BoostSite field.

### GetBoostDose

`func (o *Ga4ghRadiotherapy) GetBoostDose() string`

GetBoostDose returns the BoostDose field if non-nil, zero value otherwise.

### GetBoostDoseOk

`func (o *Ga4ghRadiotherapy) GetBoostDoseOk() (string, bool)`

GetBoostDoseOk returns a tuple with the BoostDose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBoostDose

`func (o *Ga4ghRadiotherapy) HasBoostDose() bool`

HasBoostDose returns a boolean if a field has been set.

### SetBoostDose

`func (o *Ga4ghRadiotherapy) SetBoostDose(v string)`

SetBoostDose gets a reference to the given string and assigns it to the BoostDose field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


