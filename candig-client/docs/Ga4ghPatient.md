# Ga4ghPatient

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
**OtherIds** | Pointer to **string** |  | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Ethnicity** | Pointer to **string** |  | [optional] 
**Race** | Pointer to **string** |  | [optional] 
**ProvinceOfResidence** | Pointer to **string** |  | [optional] 
**DateOfDeath** | Pointer to **string** |  | [optional] 
**CauseOfDeath** | Pointer to **string** |  | [optional] 
**AutopsyTissueForResearch** | Pointer to **string** |  | [optional] 
**PriorMalignancy** | Pointer to **string** |  | [optional] 
**DateOfPriorMalignancy** | Pointer to **string** |  | [optional] 
**FamilyHistoryAndRiskFactors** | Pointer to **string** |  | [optional] 
**FamilyHistoryOfPredispositionSyndrome** | Pointer to **string** |  | [optional] 
**DetailsOfPredispositionSyndrome** | Pointer to **string** |  | [optional] 
**GeneticCancerSyndrome** | Pointer to **string** |  | [optional] 
**OtherGeneticConditionOrSignificantComorbidity** | Pointer to **string** |  | [optional] 
**OccupationalOrEnvironmentalExposure** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghPatient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghPatient) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghPatient) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghPatient) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghPatient) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghPatient) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghPatient) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghPatient) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghPatient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghPatient) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghPatient) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghPatient) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghPatient) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghPatient) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghPatient) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghPatient) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghPatient) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghPatient) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghPatient) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghPatient) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghPatient) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghPatient) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghPatient) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghPatient) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghPatient) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghPatient) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghPatient) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghPatient) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetOtherIds

`func (o *Ga4ghPatient) GetOtherIds() string`

GetOtherIds returns the OtherIds field if non-nil, zero value otherwise.

### GetOtherIdsOk

`func (o *Ga4ghPatient) GetOtherIdsOk() (string, bool)`

GetOtherIdsOk returns a tuple with the OtherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherIds

`func (o *Ga4ghPatient) HasOtherIds() bool`

HasOtherIds returns a boolean if a field has been set.

### SetOtherIds

`func (o *Ga4ghPatient) SetOtherIds(v string)`

SetOtherIds gets a reference to the given string and assigns it to the OtherIds field.

### GetDateOfBirth

`func (o *Ga4ghPatient) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Ga4ghPatient) GetDateOfBirthOk() (string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfBirth

`func (o *Ga4ghPatient) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirth

`func (o *Ga4ghPatient) SetDateOfBirth(v string)`

SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.

### GetGender

`func (o *Ga4ghPatient) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Ga4ghPatient) GetGenderOk() (string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGender

`func (o *Ga4ghPatient) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGender

`func (o *Ga4ghPatient) SetGender(v string)`

SetGender gets a reference to the given string and assigns it to the Gender field.

### GetEthnicity

`func (o *Ga4ghPatient) GetEthnicity() string`

GetEthnicity returns the Ethnicity field if non-nil, zero value otherwise.

### GetEthnicityOk

`func (o *Ga4ghPatient) GetEthnicityOk() (string, bool)`

GetEthnicityOk returns a tuple with the Ethnicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEthnicity

`func (o *Ga4ghPatient) HasEthnicity() bool`

HasEthnicity returns a boolean if a field has been set.

### SetEthnicity

`func (o *Ga4ghPatient) SetEthnicity(v string)`

SetEthnicity gets a reference to the given string and assigns it to the Ethnicity field.

### GetRace

`func (o *Ga4ghPatient) GetRace() string`

GetRace returns the Race field if non-nil, zero value otherwise.

### GetRaceOk

`func (o *Ga4ghPatient) GetRaceOk() (string, bool)`

GetRaceOk returns a tuple with the Race field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRace

`func (o *Ga4ghPatient) HasRace() bool`

HasRace returns a boolean if a field has been set.

### SetRace

`func (o *Ga4ghPatient) SetRace(v string)`

SetRace gets a reference to the given string and assigns it to the Race field.

### GetProvinceOfResidence

`func (o *Ga4ghPatient) GetProvinceOfResidence() string`

GetProvinceOfResidence returns the ProvinceOfResidence field if non-nil, zero value otherwise.

### GetProvinceOfResidenceOk

`func (o *Ga4ghPatient) GetProvinceOfResidenceOk() (string, bool)`

GetProvinceOfResidenceOk returns a tuple with the ProvinceOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvinceOfResidence

`func (o *Ga4ghPatient) HasProvinceOfResidence() bool`

HasProvinceOfResidence returns a boolean if a field has been set.

### SetProvinceOfResidence

`func (o *Ga4ghPatient) SetProvinceOfResidence(v string)`

SetProvinceOfResidence gets a reference to the given string and assigns it to the ProvinceOfResidence field.

### GetDateOfDeath

`func (o *Ga4ghPatient) GetDateOfDeath() string`

GetDateOfDeath returns the DateOfDeath field if non-nil, zero value otherwise.

### GetDateOfDeathOk

`func (o *Ga4ghPatient) GetDateOfDeathOk() (string, bool)`

GetDateOfDeathOk returns a tuple with the DateOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfDeath

`func (o *Ga4ghPatient) HasDateOfDeath() bool`

HasDateOfDeath returns a boolean if a field has been set.

### SetDateOfDeath

`func (o *Ga4ghPatient) SetDateOfDeath(v string)`

SetDateOfDeath gets a reference to the given string and assigns it to the DateOfDeath field.

### GetCauseOfDeath

`func (o *Ga4ghPatient) GetCauseOfDeath() string`

GetCauseOfDeath returns the CauseOfDeath field if non-nil, zero value otherwise.

### GetCauseOfDeathOk

`func (o *Ga4ghPatient) GetCauseOfDeathOk() (string, bool)`

GetCauseOfDeathOk returns a tuple with the CauseOfDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCauseOfDeath

`func (o *Ga4ghPatient) HasCauseOfDeath() bool`

HasCauseOfDeath returns a boolean if a field has been set.

### SetCauseOfDeath

`func (o *Ga4ghPatient) SetCauseOfDeath(v string)`

SetCauseOfDeath gets a reference to the given string and assigns it to the CauseOfDeath field.

### GetAutopsyTissueForResearch

`func (o *Ga4ghPatient) GetAutopsyTissueForResearch() string`

GetAutopsyTissueForResearch returns the AutopsyTissueForResearch field if non-nil, zero value otherwise.

### GetAutopsyTissueForResearchOk

`func (o *Ga4ghPatient) GetAutopsyTissueForResearchOk() (string, bool)`

GetAutopsyTissueForResearchOk returns a tuple with the AutopsyTissueForResearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutopsyTissueForResearch

`func (o *Ga4ghPatient) HasAutopsyTissueForResearch() bool`

HasAutopsyTissueForResearch returns a boolean if a field has been set.

### SetAutopsyTissueForResearch

`func (o *Ga4ghPatient) SetAutopsyTissueForResearch(v string)`

SetAutopsyTissueForResearch gets a reference to the given string and assigns it to the AutopsyTissueForResearch field.

### GetPriorMalignancy

`func (o *Ga4ghPatient) GetPriorMalignancy() string`

GetPriorMalignancy returns the PriorMalignancy field if non-nil, zero value otherwise.

### GetPriorMalignancyOk

`func (o *Ga4ghPatient) GetPriorMalignancyOk() (string, bool)`

GetPriorMalignancyOk returns a tuple with the PriorMalignancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriorMalignancy

`func (o *Ga4ghPatient) HasPriorMalignancy() bool`

HasPriorMalignancy returns a boolean if a field has been set.

### SetPriorMalignancy

`func (o *Ga4ghPatient) SetPriorMalignancy(v string)`

SetPriorMalignancy gets a reference to the given string and assigns it to the PriorMalignancy field.

### GetDateOfPriorMalignancy

`func (o *Ga4ghPatient) GetDateOfPriorMalignancy() string`

GetDateOfPriorMalignancy returns the DateOfPriorMalignancy field if non-nil, zero value otherwise.

### GetDateOfPriorMalignancyOk

`func (o *Ga4ghPatient) GetDateOfPriorMalignancyOk() (string, bool)`

GetDateOfPriorMalignancyOk returns a tuple with the DateOfPriorMalignancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfPriorMalignancy

`func (o *Ga4ghPatient) HasDateOfPriorMalignancy() bool`

HasDateOfPriorMalignancy returns a boolean if a field has been set.

### SetDateOfPriorMalignancy

`func (o *Ga4ghPatient) SetDateOfPriorMalignancy(v string)`

SetDateOfPriorMalignancy gets a reference to the given string and assigns it to the DateOfPriorMalignancy field.

### GetFamilyHistoryAndRiskFactors

`func (o *Ga4ghPatient) GetFamilyHistoryAndRiskFactors() string`

GetFamilyHistoryAndRiskFactors returns the FamilyHistoryAndRiskFactors field if non-nil, zero value otherwise.

### GetFamilyHistoryAndRiskFactorsOk

`func (o *Ga4ghPatient) GetFamilyHistoryAndRiskFactorsOk() (string, bool)`

GetFamilyHistoryAndRiskFactorsOk returns a tuple with the FamilyHistoryAndRiskFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFamilyHistoryAndRiskFactors

`func (o *Ga4ghPatient) HasFamilyHistoryAndRiskFactors() bool`

HasFamilyHistoryAndRiskFactors returns a boolean if a field has been set.

### SetFamilyHistoryAndRiskFactors

`func (o *Ga4ghPatient) SetFamilyHistoryAndRiskFactors(v string)`

SetFamilyHistoryAndRiskFactors gets a reference to the given string and assigns it to the FamilyHistoryAndRiskFactors field.

### GetFamilyHistoryOfPredispositionSyndrome

`func (o *Ga4ghPatient) GetFamilyHistoryOfPredispositionSyndrome() string`

GetFamilyHistoryOfPredispositionSyndrome returns the FamilyHistoryOfPredispositionSyndrome field if non-nil, zero value otherwise.

### GetFamilyHistoryOfPredispositionSyndromeOk

`func (o *Ga4ghPatient) GetFamilyHistoryOfPredispositionSyndromeOk() (string, bool)`

GetFamilyHistoryOfPredispositionSyndromeOk returns a tuple with the FamilyHistoryOfPredispositionSyndrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFamilyHistoryOfPredispositionSyndrome

`func (o *Ga4ghPatient) HasFamilyHistoryOfPredispositionSyndrome() bool`

HasFamilyHistoryOfPredispositionSyndrome returns a boolean if a field has been set.

### SetFamilyHistoryOfPredispositionSyndrome

`func (o *Ga4ghPatient) SetFamilyHistoryOfPredispositionSyndrome(v string)`

SetFamilyHistoryOfPredispositionSyndrome gets a reference to the given string and assigns it to the FamilyHistoryOfPredispositionSyndrome field.

### GetDetailsOfPredispositionSyndrome

`func (o *Ga4ghPatient) GetDetailsOfPredispositionSyndrome() string`

GetDetailsOfPredispositionSyndrome returns the DetailsOfPredispositionSyndrome field if non-nil, zero value otherwise.

### GetDetailsOfPredispositionSyndromeOk

`func (o *Ga4ghPatient) GetDetailsOfPredispositionSyndromeOk() (string, bool)`

GetDetailsOfPredispositionSyndromeOk returns a tuple with the DetailsOfPredispositionSyndrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetailsOfPredispositionSyndrome

`func (o *Ga4ghPatient) HasDetailsOfPredispositionSyndrome() bool`

HasDetailsOfPredispositionSyndrome returns a boolean if a field has been set.

### SetDetailsOfPredispositionSyndrome

`func (o *Ga4ghPatient) SetDetailsOfPredispositionSyndrome(v string)`

SetDetailsOfPredispositionSyndrome gets a reference to the given string and assigns it to the DetailsOfPredispositionSyndrome field.

### GetGeneticCancerSyndrome

`func (o *Ga4ghPatient) GetGeneticCancerSyndrome() string`

GetGeneticCancerSyndrome returns the GeneticCancerSyndrome field if non-nil, zero value otherwise.

### GetGeneticCancerSyndromeOk

`func (o *Ga4ghPatient) GetGeneticCancerSyndromeOk() (string, bool)`

GetGeneticCancerSyndromeOk returns a tuple with the GeneticCancerSyndrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeneticCancerSyndrome

`func (o *Ga4ghPatient) HasGeneticCancerSyndrome() bool`

HasGeneticCancerSyndrome returns a boolean if a field has been set.

### SetGeneticCancerSyndrome

`func (o *Ga4ghPatient) SetGeneticCancerSyndrome(v string)`

SetGeneticCancerSyndrome gets a reference to the given string and assigns it to the GeneticCancerSyndrome field.

### GetOtherGeneticConditionOrSignificantComorbidity

`func (o *Ga4ghPatient) GetOtherGeneticConditionOrSignificantComorbidity() string`

GetOtherGeneticConditionOrSignificantComorbidity returns the OtherGeneticConditionOrSignificantComorbidity field if non-nil, zero value otherwise.

### GetOtherGeneticConditionOrSignificantComorbidityOk

`func (o *Ga4ghPatient) GetOtherGeneticConditionOrSignificantComorbidityOk() (string, bool)`

GetOtherGeneticConditionOrSignificantComorbidityOk returns a tuple with the OtherGeneticConditionOrSignificantComorbidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherGeneticConditionOrSignificantComorbidity

`func (o *Ga4ghPatient) HasOtherGeneticConditionOrSignificantComorbidity() bool`

HasOtherGeneticConditionOrSignificantComorbidity returns a boolean if a field has been set.

### SetOtherGeneticConditionOrSignificantComorbidity

`func (o *Ga4ghPatient) SetOtherGeneticConditionOrSignificantComorbidity(v string)`

SetOtherGeneticConditionOrSignificantComorbidity gets a reference to the given string and assigns it to the OtherGeneticConditionOrSignificantComorbidity field.

### GetOccupationalOrEnvironmentalExposure

`func (o *Ga4ghPatient) GetOccupationalOrEnvironmentalExposure() string`

GetOccupationalOrEnvironmentalExposure returns the OccupationalOrEnvironmentalExposure field if non-nil, zero value otherwise.

### GetOccupationalOrEnvironmentalExposureOk

`func (o *Ga4ghPatient) GetOccupationalOrEnvironmentalExposureOk() (string, bool)`

GetOccupationalOrEnvironmentalExposureOk returns a tuple with the OccupationalOrEnvironmentalExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOccupationalOrEnvironmentalExposure

`func (o *Ga4ghPatient) HasOccupationalOrEnvironmentalExposure() bool`

HasOccupationalOrEnvironmentalExposure returns a boolean if a field has been set.

### SetOccupationalOrEnvironmentalExposure

`func (o *Ga4ghPatient) SetOccupationalOrEnvironmentalExposure(v string)`

SetOccupationalOrEnvironmentalExposure gets a reference to the given string and assigns it to the OccupationalOrEnvironmentalExposure field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


