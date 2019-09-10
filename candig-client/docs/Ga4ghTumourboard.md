# Ga4ghTumourboard

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
**DateOfMolecularTumorBoard** | Pointer to **string** |  | [optional] 
**TypeOfSampleAnalyzed** | Pointer to **string** |  | [optional] 
**TypeOfTumourSampleAnalyzed** | Pointer to **string** |  | [optional] 
**AnalysesDiscussed** | Pointer to **string** |  | [optional] 
**SomaticSampleType** | Pointer to **string** |  | [optional] 
**NormalExpressionComparator** | Pointer to **string** |  | [optional] 
**DiseaseExpressionComparator** | Pointer to **string** |  | [optional] 
**HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer** | Pointer to **string** |  | [optional] 
**ActionableTargetFound** | Pointer to **string** |  | [optional] 
**MolecularTumorBoardRecommendation** | Pointer to **string** |  | [optional] 
**GermlineDnaSampleId** | Pointer to **string** |  | [optional] 
**TumorDnaSampleId** | Pointer to **string** |  | [optional] 
**TumorRnaSampleId** | Pointer to **string** |  | [optional] 
**GermlineSnvDiscussed** | Pointer to **string** |  | [optional] 
**SomaticSnvDiscussed** | Pointer to **string** |  | [optional] 
**CnvsDiscussed** | Pointer to **string** |  | [optional] 
**StructuralVariantDiscussed** | Pointer to **string** |  | [optional] 
**ClassificationOfVariants** | Pointer to **string** |  | [optional] 
**ClinicalValidationProgress** | Pointer to **string** |  | [optional] 
**TypeOfValidation** | Pointer to **string** |  | [optional] 
**AgentOrDrugClass** | Pointer to **string** |  | [optional] 
**LevelOfEvidenceForExpressionTargetAgentMatch** | Pointer to **string** |  | [optional] 
**DidTreatmentPlanChangeBasedOnProfilingResult** | Pointer to **string** |  | [optional] 
**HowTreatmentHasAlteredBasedOnProfiling** | Pointer to **string** |  | [optional] 
**ReasonTreatmentPlanDidNotChangeBasedOnProfiling** | Pointer to **string** |  | [optional] 
**DetailsOfTreatmentPlanImpact** | Pointer to **string** |  | [optional] 
**PatientOrFamilyInformedOfGermlineVariant** | Pointer to **string** |  | [optional] 
**PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling** | Pointer to **string** |  | [optional] 
**SummaryReport** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghTumourboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghTumourboard) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghTumourboard) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghTumourboard) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghTumourboard) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghTumourboard) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghTumourboard) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghTumourboard) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghTumourboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghTumourboard) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghTumourboard) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghTumourboard) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghTumourboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghTumourboard) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghTumourboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghTumourboard) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghTumourboard) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghTumourboard) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghTumourboard) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghTumourboard) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghTumourboard) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghTumourboard) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghTumourboard) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghTumourboard) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghTumourboard) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghTumourboard) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghTumourboard) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghTumourboard) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetDateOfMolecularTumorBoard

`func (o *Ga4ghTumourboard) GetDateOfMolecularTumorBoard() string`

GetDateOfMolecularTumorBoard returns the DateOfMolecularTumorBoard field if non-nil, zero value otherwise.

### GetDateOfMolecularTumorBoardOk

`func (o *Ga4ghTumourboard) GetDateOfMolecularTumorBoardOk() (string, bool)`

GetDateOfMolecularTumorBoardOk returns a tuple with the DateOfMolecularTumorBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfMolecularTumorBoard

`func (o *Ga4ghTumourboard) HasDateOfMolecularTumorBoard() bool`

HasDateOfMolecularTumorBoard returns a boolean if a field has been set.

### SetDateOfMolecularTumorBoard

`func (o *Ga4ghTumourboard) SetDateOfMolecularTumorBoard(v string)`

SetDateOfMolecularTumorBoard gets a reference to the given string and assigns it to the DateOfMolecularTumorBoard field.

### GetTypeOfSampleAnalyzed

`func (o *Ga4ghTumourboard) GetTypeOfSampleAnalyzed() string`

GetTypeOfSampleAnalyzed returns the TypeOfSampleAnalyzed field if non-nil, zero value otherwise.

### GetTypeOfSampleAnalyzedOk

`func (o *Ga4ghTumourboard) GetTypeOfSampleAnalyzedOk() (string, bool)`

GetTypeOfSampleAnalyzedOk returns a tuple with the TypeOfSampleAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeOfSampleAnalyzed

`func (o *Ga4ghTumourboard) HasTypeOfSampleAnalyzed() bool`

HasTypeOfSampleAnalyzed returns a boolean if a field has been set.

### SetTypeOfSampleAnalyzed

`func (o *Ga4ghTumourboard) SetTypeOfSampleAnalyzed(v string)`

SetTypeOfSampleAnalyzed gets a reference to the given string and assigns it to the TypeOfSampleAnalyzed field.

### GetTypeOfTumourSampleAnalyzed

`func (o *Ga4ghTumourboard) GetTypeOfTumourSampleAnalyzed() string`

GetTypeOfTumourSampleAnalyzed returns the TypeOfTumourSampleAnalyzed field if non-nil, zero value otherwise.

### GetTypeOfTumourSampleAnalyzedOk

`func (o *Ga4ghTumourboard) GetTypeOfTumourSampleAnalyzedOk() (string, bool)`

GetTypeOfTumourSampleAnalyzedOk returns a tuple with the TypeOfTumourSampleAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeOfTumourSampleAnalyzed

`func (o *Ga4ghTumourboard) HasTypeOfTumourSampleAnalyzed() bool`

HasTypeOfTumourSampleAnalyzed returns a boolean if a field has been set.

### SetTypeOfTumourSampleAnalyzed

`func (o *Ga4ghTumourboard) SetTypeOfTumourSampleAnalyzed(v string)`

SetTypeOfTumourSampleAnalyzed gets a reference to the given string and assigns it to the TypeOfTumourSampleAnalyzed field.

### GetAnalysesDiscussed

`func (o *Ga4ghTumourboard) GetAnalysesDiscussed() string`

GetAnalysesDiscussed returns the AnalysesDiscussed field if non-nil, zero value otherwise.

### GetAnalysesDiscussedOk

`func (o *Ga4ghTumourboard) GetAnalysesDiscussedOk() (string, bool)`

GetAnalysesDiscussedOk returns a tuple with the AnalysesDiscussed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnalysesDiscussed

`func (o *Ga4ghTumourboard) HasAnalysesDiscussed() bool`

HasAnalysesDiscussed returns a boolean if a field has been set.

### SetAnalysesDiscussed

`func (o *Ga4ghTumourboard) SetAnalysesDiscussed(v string)`

SetAnalysesDiscussed gets a reference to the given string and assigns it to the AnalysesDiscussed field.

### GetSomaticSampleType

`func (o *Ga4ghTumourboard) GetSomaticSampleType() string`

GetSomaticSampleType returns the SomaticSampleType field if non-nil, zero value otherwise.

### GetSomaticSampleTypeOk

`func (o *Ga4ghTumourboard) GetSomaticSampleTypeOk() (string, bool)`

GetSomaticSampleTypeOk returns a tuple with the SomaticSampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSomaticSampleType

`func (o *Ga4ghTumourboard) HasSomaticSampleType() bool`

HasSomaticSampleType returns a boolean if a field has been set.

### SetSomaticSampleType

`func (o *Ga4ghTumourboard) SetSomaticSampleType(v string)`

SetSomaticSampleType gets a reference to the given string and assigns it to the SomaticSampleType field.

### GetNormalExpressionComparator

`func (o *Ga4ghTumourboard) GetNormalExpressionComparator() string`

GetNormalExpressionComparator returns the NormalExpressionComparator field if non-nil, zero value otherwise.

### GetNormalExpressionComparatorOk

`func (o *Ga4ghTumourboard) GetNormalExpressionComparatorOk() (string, bool)`

GetNormalExpressionComparatorOk returns a tuple with the NormalExpressionComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNormalExpressionComparator

`func (o *Ga4ghTumourboard) HasNormalExpressionComparator() bool`

HasNormalExpressionComparator returns a boolean if a field has been set.

### SetNormalExpressionComparator

`func (o *Ga4ghTumourboard) SetNormalExpressionComparator(v string)`

SetNormalExpressionComparator gets a reference to the given string and assigns it to the NormalExpressionComparator field.

### GetDiseaseExpressionComparator

`func (o *Ga4ghTumourboard) GetDiseaseExpressionComparator() string`

GetDiseaseExpressionComparator returns the DiseaseExpressionComparator field if non-nil, zero value otherwise.

### GetDiseaseExpressionComparatorOk

`func (o *Ga4ghTumourboard) GetDiseaseExpressionComparatorOk() (string, bool)`

GetDiseaseExpressionComparatorOk returns a tuple with the DiseaseExpressionComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiseaseExpressionComparator

`func (o *Ga4ghTumourboard) HasDiseaseExpressionComparator() bool`

HasDiseaseExpressionComparator returns a boolean if a field has been set.

### SetDiseaseExpressionComparator

`func (o *Ga4ghTumourboard) SetDiseaseExpressionComparator(v string)`

SetDiseaseExpressionComparator gets a reference to the given string and assigns it to the DiseaseExpressionComparator field.

### GetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer

`func (o *Ga4ghTumourboard) GetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer() string`

GetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer returns the HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer field if non-nil, zero value otherwise.

### GetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancerOk

`func (o *Ga4ghTumourboard) GetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancerOk() (string, bool)`

GetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancerOk returns a tuple with the HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer

`func (o *Ga4ghTumourboard) HasHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer() bool`

HasHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer returns a boolean if a field has been set.

### SetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer

`func (o *Ga4ghTumourboard) SetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer(v string)`

SetHasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer gets a reference to the given string and assigns it to the HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer field.

### GetActionableTargetFound

`func (o *Ga4ghTumourboard) GetActionableTargetFound() string`

GetActionableTargetFound returns the ActionableTargetFound field if non-nil, zero value otherwise.

### GetActionableTargetFoundOk

`func (o *Ga4ghTumourboard) GetActionableTargetFoundOk() (string, bool)`

GetActionableTargetFoundOk returns a tuple with the ActionableTargetFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionableTargetFound

`func (o *Ga4ghTumourboard) HasActionableTargetFound() bool`

HasActionableTargetFound returns a boolean if a field has been set.

### SetActionableTargetFound

`func (o *Ga4ghTumourboard) SetActionableTargetFound(v string)`

SetActionableTargetFound gets a reference to the given string and assigns it to the ActionableTargetFound field.

### GetMolecularTumorBoardRecommendation

`func (o *Ga4ghTumourboard) GetMolecularTumorBoardRecommendation() string`

GetMolecularTumorBoardRecommendation returns the MolecularTumorBoardRecommendation field if non-nil, zero value otherwise.

### GetMolecularTumorBoardRecommendationOk

`func (o *Ga4ghTumourboard) GetMolecularTumorBoardRecommendationOk() (string, bool)`

GetMolecularTumorBoardRecommendationOk returns a tuple with the MolecularTumorBoardRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMolecularTumorBoardRecommendation

`func (o *Ga4ghTumourboard) HasMolecularTumorBoardRecommendation() bool`

HasMolecularTumorBoardRecommendation returns a boolean if a field has been set.

### SetMolecularTumorBoardRecommendation

`func (o *Ga4ghTumourboard) SetMolecularTumorBoardRecommendation(v string)`

SetMolecularTumorBoardRecommendation gets a reference to the given string and assigns it to the MolecularTumorBoardRecommendation field.

### GetGermlineDnaSampleId

`func (o *Ga4ghTumourboard) GetGermlineDnaSampleId() string`

GetGermlineDnaSampleId returns the GermlineDnaSampleId field if non-nil, zero value otherwise.

### GetGermlineDnaSampleIdOk

`func (o *Ga4ghTumourboard) GetGermlineDnaSampleIdOk() (string, bool)`

GetGermlineDnaSampleIdOk returns a tuple with the GermlineDnaSampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGermlineDnaSampleId

`func (o *Ga4ghTumourboard) HasGermlineDnaSampleId() bool`

HasGermlineDnaSampleId returns a boolean if a field has been set.

### SetGermlineDnaSampleId

`func (o *Ga4ghTumourboard) SetGermlineDnaSampleId(v string)`

SetGermlineDnaSampleId gets a reference to the given string and assigns it to the GermlineDnaSampleId field.

### GetTumorDnaSampleId

`func (o *Ga4ghTumourboard) GetTumorDnaSampleId() string`

GetTumorDnaSampleId returns the TumorDnaSampleId field if non-nil, zero value otherwise.

### GetTumorDnaSampleIdOk

`func (o *Ga4ghTumourboard) GetTumorDnaSampleIdOk() (string, bool)`

GetTumorDnaSampleIdOk returns a tuple with the TumorDnaSampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTumorDnaSampleId

`func (o *Ga4ghTumourboard) HasTumorDnaSampleId() bool`

HasTumorDnaSampleId returns a boolean if a field has been set.

### SetTumorDnaSampleId

`func (o *Ga4ghTumourboard) SetTumorDnaSampleId(v string)`

SetTumorDnaSampleId gets a reference to the given string and assigns it to the TumorDnaSampleId field.

### GetTumorRnaSampleId

`func (o *Ga4ghTumourboard) GetTumorRnaSampleId() string`

GetTumorRnaSampleId returns the TumorRnaSampleId field if non-nil, zero value otherwise.

### GetTumorRnaSampleIdOk

`func (o *Ga4ghTumourboard) GetTumorRnaSampleIdOk() (string, bool)`

GetTumorRnaSampleIdOk returns a tuple with the TumorRnaSampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTumorRnaSampleId

`func (o *Ga4ghTumourboard) HasTumorRnaSampleId() bool`

HasTumorRnaSampleId returns a boolean if a field has been set.

### SetTumorRnaSampleId

`func (o *Ga4ghTumourboard) SetTumorRnaSampleId(v string)`

SetTumorRnaSampleId gets a reference to the given string and assigns it to the TumorRnaSampleId field.

### GetGermlineSnvDiscussed

`func (o *Ga4ghTumourboard) GetGermlineSnvDiscussed() string`

GetGermlineSnvDiscussed returns the GermlineSnvDiscussed field if non-nil, zero value otherwise.

### GetGermlineSnvDiscussedOk

`func (o *Ga4ghTumourboard) GetGermlineSnvDiscussedOk() (string, bool)`

GetGermlineSnvDiscussedOk returns a tuple with the GermlineSnvDiscussed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGermlineSnvDiscussed

`func (o *Ga4ghTumourboard) HasGermlineSnvDiscussed() bool`

HasGermlineSnvDiscussed returns a boolean if a field has been set.

### SetGermlineSnvDiscussed

`func (o *Ga4ghTumourboard) SetGermlineSnvDiscussed(v string)`

SetGermlineSnvDiscussed gets a reference to the given string and assigns it to the GermlineSnvDiscussed field.

### GetSomaticSnvDiscussed

`func (o *Ga4ghTumourboard) GetSomaticSnvDiscussed() string`

GetSomaticSnvDiscussed returns the SomaticSnvDiscussed field if non-nil, zero value otherwise.

### GetSomaticSnvDiscussedOk

`func (o *Ga4ghTumourboard) GetSomaticSnvDiscussedOk() (string, bool)`

GetSomaticSnvDiscussedOk returns a tuple with the SomaticSnvDiscussed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSomaticSnvDiscussed

`func (o *Ga4ghTumourboard) HasSomaticSnvDiscussed() bool`

HasSomaticSnvDiscussed returns a boolean if a field has been set.

### SetSomaticSnvDiscussed

`func (o *Ga4ghTumourboard) SetSomaticSnvDiscussed(v string)`

SetSomaticSnvDiscussed gets a reference to the given string and assigns it to the SomaticSnvDiscussed field.

### GetCnvsDiscussed

`func (o *Ga4ghTumourboard) GetCnvsDiscussed() string`

GetCnvsDiscussed returns the CnvsDiscussed field if non-nil, zero value otherwise.

### GetCnvsDiscussedOk

`func (o *Ga4ghTumourboard) GetCnvsDiscussedOk() (string, bool)`

GetCnvsDiscussedOk returns a tuple with the CnvsDiscussed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCnvsDiscussed

`func (o *Ga4ghTumourboard) HasCnvsDiscussed() bool`

HasCnvsDiscussed returns a boolean if a field has been set.

### SetCnvsDiscussed

`func (o *Ga4ghTumourboard) SetCnvsDiscussed(v string)`

SetCnvsDiscussed gets a reference to the given string and assigns it to the CnvsDiscussed field.

### GetStructuralVariantDiscussed

`func (o *Ga4ghTumourboard) GetStructuralVariantDiscussed() string`

GetStructuralVariantDiscussed returns the StructuralVariantDiscussed field if non-nil, zero value otherwise.

### GetStructuralVariantDiscussedOk

`func (o *Ga4ghTumourboard) GetStructuralVariantDiscussedOk() (string, bool)`

GetStructuralVariantDiscussedOk returns a tuple with the StructuralVariantDiscussed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStructuralVariantDiscussed

`func (o *Ga4ghTumourboard) HasStructuralVariantDiscussed() bool`

HasStructuralVariantDiscussed returns a boolean if a field has been set.

### SetStructuralVariantDiscussed

`func (o *Ga4ghTumourboard) SetStructuralVariantDiscussed(v string)`

SetStructuralVariantDiscussed gets a reference to the given string and assigns it to the StructuralVariantDiscussed field.

### GetClassificationOfVariants

`func (o *Ga4ghTumourboard) GetClassificationOfVariants() string`

GetClassificationOfVariants returns the ClassificationOfVariants field if non-nil, zero value otherwise.

### GetClassificationOfVariantsOk

`func (o *Ga4ghTumourboard) GetClassificationOfVariantsOk() (string, bool)`

GetClassificationOfVariantsOk returns a tuple with the ClassificationOfVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassificationOfVariants

`func (o *Ga4ghTumourboard) HasClassificationOfVariants() bool`

HasClassificationOfVariants returns a boolean if a field has been set.

### SetClassificationOfVariants

`func (o *Ga4ghTumourboard) SetClassificationOfVariants(v string)`

SetClassificationOfVariants gets a reference to the given string and assigns it to the ClassificationOfVariants field.

### GetClinicalValidationProgress

`func (o *Ga4ghTumourboard) GetClinicalValidationProgress() string`

GetClinicalValidationProgress returns the ClinicalValidationProgress field if non-nil, zero value otherwise.

### GetClinicalValidationProgressOk

`func (o *Ga4ghTumourboard) GetClinicalValidationProgressOk() (string, bool)`

GetClinicalValidationProgressOk returns a tuple with the ClinicalValidationProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClinicalValidationProgress

`func (o *Ga4ghTumourboard) HasClinicalValidationProgress() bool`

HasClinicalValidationProgress returns a boolean if a field has been set.

### SetClinicalValidationProgress

`func (o *Ga4ghTumourboard) SetClinicalValidationProgress(v string)`

SetClinicalValidationProgress gets a reference to the given string and assigns it to the ClinicalValidationProgress field.

### GetTypeOfValidation

`func (o *Ga4ghTumourboard) GetTypeOfValidation() string`

GetTypeOfValidation returns the TypeOfValidation field if non-nil, zero value otherwise.

### GetTypeOfValidationOk

`func (o *Ga4ghTumourboard) GetTypeOfValidationOk() (string, bool)`

GetTypeOfValidationOk returns a tuple with the TypeOfValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeOfValidation

`func (o *Ga4ghTumourboard) HasTypeOfValidation() bool`

HasTypeOfValidation returns a boolean if a field has been set.

### SetTypeOfValidation

`func (o *Ga4ghTumourboard) SetTypeOfValidation(v string)`

SetTypeOfValidation gets a reference to the given string and assigns it to the TypeOfValidation field.

### GetAgentOrDrugClass

`func (o *Ga4ghTumourboard) GetAgentOrDrugClass() string`

GetAgentOrDrugClass returns the AgentOrDrugClass field if non-nil, zero value otherwise.

### GetAgentOrDrugClassOk

`func (o *Ga4ghTumourboard) GetAgentOrDrugClassOk() (string, bool)`

GetAgentOrDrugClassOk returns a tuple with the AgentOrDrugClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgentOrDrugClass

`func (o *Ga4ghTumourboard) HasAgentOrDrugClass() bool`

HasAgentOrDrugClass returns a boolean if a field has been set.

### SetAgentOrDrugClass

`func (o *Ga4ghTumourboard) SetAgentOrDrugClass(v string)`

SetAgentOrDrugClass gets a reference to the given string and assigns it to the AgentOrDrugClass field.

### GetLevelOfEvidenceForExpressionTargetAgentMatch

`func (o *Ga4ghTumourboard) GetLevelOfEvidenceForExpressionTargetAgentMatch() string`

GetLevelOfEvidenceForExpressionTargetAgentMatch returns the LevelOfEvidenceForExpressionTargetAgentMatch field if non-nil, zero value otherwise.

### GetLevelOfEvidenceForExpressionTargetAgentMatchOk

`func (o *Ga4ghTumourboard) GetLevelOfEvidenceForExpressionTargetAgentMatchOk() (string, bool)`

GetLevelOfEvidenceForExpressionTargetAgentMatchOk returns a tuple with the LevelOfEvidenceForExpressionTargetAgentMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLevelOfEvidenceForExpressionTargetAgentMatch

`func (o *Ga4ghTumourboard) HasLevelOfEvidenceForExpressionTargetAgentMatch() bool`

HasLevelOfEvidenceForExpressionTargetAgentMatch returns a boolean if a field has been set.

### SetLevelOfEvidenceForExpressionTargetAgentMatch

`func (o *Ga4ghTumourboard) SetLevelOfEvidenceForExpressionTargetAgentMatch(v string)`

SetLevelOfEvidenceForExpressionTargetAgentMatch gets a reference to the given string and assigns it to the LevelOfEvidenceForExpressionTargetAgentMatch field.

### GetDidTreatmentPlanChangeBasedOnProfilingResult

`func (o *Ga4ghTumourboard) GetDidTreatmentPlanChangeBasedOnProfilingResult() string`

GetDidTreatmentPlanChangeBasedOnProfilingResult returns the DidTreatmentPlanChangeBasedOnProfilingResult field if non-nil, zero value otherwise.

### GetDidTreatmentPlanChangeBasedOnProfilingResultOk

`func (o *Ga4ghTumourboard) GetDidTreatmentPlanChangeBasedOnProfilingResultOk() (string, bool)`

GetDidTreatmentPlanChangeBasedOnProfilingResultOk returns a tuple with the DidTreatmentPlanChangeBasedOnProfilingResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDidTreatmentPlanChangeBasedOnProfilingResult

`func (o *Ga4ghTumourboard) HasDidTreatmentPlanChangeBasedOnProfilingResult() bool`

HasDidTreatmentPlanChangeBasedOnProfilingResult returns a boolean if a field has been set.

### SetDidTreatmentPlanChangeBasedOnProfilingResult

`func (o *Ga4ghTumourboard) SetDidTreatmentPlanChangeBasedOnProfilingResult(v string)`

SetDidTreatmentPlanChangeBasedOnProfilingResult gets a reference to the given string and assigns it to the DidTreatmentPlanChangeBasedOnProfilingResult field.

### GetHowTreatmentHasAlteredBasedOnProfiling

`func (o *Ga4ghTumourboard) GetHowTreatmentHasAlteredBasedOnProfiling() string`

GetHowTreatmentHasAlteredBasedOnProfiling returns the HowTreatmentHasAlteredBasedOnProfiling field if non-nil, zero value otherwise.

### GetHowTreatmentHasAlteredBasedOnProfilingOk

`func (o *Ga4ghTumourboard) GetHowTreatmentHasAlteredBasedOnProfilingOk() (string, bool)`

GetHowTreatmentHasAlteredBasedOnProfilingOk returns a tuple with the HowTreatmentHasAlteredBasedOnProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHowTreatmentHasAlteredBasedOnProfiling

`func (o *Ga4ghTumourboard) HasHowTreatmentHasAlteredBasedOnProfiling() bool`

HasHowTreatmentHasAlteredBasedOnProfiling returns a boolean if a field has been set.

### SetHowTreatmentHasAlteredBasedOnProfiling

`func (o *Ga4ghTumourboard) SetHowTreatmentHasAlteredBasedOnProfiling(v string)`

SetHowTreatmentHasAlteredBasedOnProfiling gets a reference to the given string and assigns it to the HowTreatmentHasAlteredBasedOnProfiling field.

### GetReasonTreatmentPlanDidNotChangeBasedOnProfiling

`func (o *Ga4ghTumourboard) GetReasonTreatmentPlanDidNotChangeBasedOnProfiling() string`

GetReasonTreatmentPlanDidNotChangeBasedOnProfiling returns the ReasonTreatmentPlanDidNotChangeBasedOnProfiling field if non-nil, zero value otherwise.

### GetReasonTreatmentPlanDidNotChangeBasedOnProfilingOk

`func (o *Ga4ghTumourboard) GetReasonTreatmentPlanDidNotChangeBasedOnProfilingOk() (string, bool)`

GetReasonTreatmentPlanDidNotChangeBasedOnProfilingOk returns a tuple with the ReasonTreatmentPlanDidNotChangeBasedOnProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReasonTreatmentPlanDidNotChangeBasedOnProfiling

`func (o *Ga4ghTumourboard) HasReasonTreatmentPlanDidNotChangeBasedOnProfiling() bool`

HasReasonTreatmentPlanDidNotChangeBasedOnProfiling returns a boolean if a field has been set.

### SetReasonTreatmentPlanDidNotChangeBasedOnProfiling

`func (o *Ga4ghTumourboard) SetReasonTreatmentPlanDidNotChangeBasedOnProfiling(v string)`

SetReasonTreatmentPlanDidNotChangeBasedOnProfiling gets a reference to the given string and assigns it to the ReasonTreatmentPlanDidNotChangeBasedOnProfiling field.

### GetDetailsOfTreatmentPlanImpact

`func (o *Ga4ghTumourboard) GetDetailsOfTreatmentPlanImpact() string`

GetDetailsOfTreatmentPlanImpact returns the DetailsOfTreatmentPlanImpact field if non-nil, zero value otherwise.

### GetDetailsOfTreatmentPlanImpactOk

`func (o *Ga4ghTumourboard) GetDetailsOfTreatmentPlanImpactOk() (string, bool)`

GetDetailsOfTreatmentPlanImpactOk returns a tuple with the DetailsOfTreatmentPlanImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetailsOfTreatmentPlanImpact

`func (o *Ga4ghTumourboard) HasDetailsOfTreatmentPlanImpact() bool`

HasDetailsOfTreatmentPlanImpact returns a boolean if a field has been set.

### SetDetailsOfTreatmentPlanImpact

`func (o *Ga4ghTumourboard) SetDetailsOfTreatmentPlanImpact(v string)`

SetDetailsOfTreatmentPlanImpact gets a reference to the given string and assigns it to the DetailsOfTreatmentPlanImpact field.

### GetPatientOrFamilyInformedOfGermlineVariant

`func (o *Ga4ghTumourboard) GetPatientOrFamilyInformedOfGermlineVariant() string`

GetPatientOrFamilyInformedOfGermlineVariant returns the PatientOrFamilyInformedOfGermlineVariant field if non-nil, zero value otherwise.

### GetPatientOrFamilyInformedOfGermlineVariantOk

`func (o *Ga4ghTumourboard) GetPatientOrFamilyInformedOfGermlineVariantOk() (string, bool)`

GetPatientOrFamilyInformedOfGermlineVariantOk returns a tuple with the PatientOrFamilyInformedOfGermlineVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientOrFamilyInformedOfGermlineVariant

`func (o *Ga4ghTumourboard) HasPatientOrFamilyInformedOfGermlineVariant() bool`

HasPatientOrFamilyInformedOfGermlineVariant returns a boolean if a field has been set.

### SetPatientOrFamilyInformedOfGermlineVariant

`func (o *Ga4ghTumourboard) SetPatientOrFamilyInformedOfGermlineVariant(v string)`

SetPatientOrFamilyInformedOfGermlineVariant gets a reference to the given string and assigns it to the PatientOrFamilyInformedOfGermlineVariant field.

### GetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling

`func (o *Ga4ghTumourboard) GetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling() string`

GetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling returns the PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling field if non-nil, zero value otherwise.

### GetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfilingOk

`func (o *Ga4ghTumourboard) GetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfilingOk() (string, bool)`

GetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfilingOk returns a tuple with the PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling

`func (o *Ga4ghTumourboard) HasPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling() bool`

HasPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling returns a boolean if a field has been set.

### SetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling

`func (o *Ga4ghTumourboard) SetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling(v string)`

SetPatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling gets a reference to the given string and assigns it to the PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling field.

### GetSummaryReport

`func (o *Ga4ghTumourboard) GetSummaryReport() string`

GetSummaryReport returns the SummaryReport field if non-nil, zero value otherwise.

### GetSummaryReportOk

`func (o *Ga4ghTumourboard) GetSummaryReportOk() (string, bool)`

GetSummaryReportOk returns a tuple with the SummaryReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSummaryReport

`func (o *Ga4ghTumourboard) HasSummaryReport() bool`

HasSummaryReport returns a boolean if a field has been set.

### SetSummaryReport

`func (o *Ga4ghTumourboard) SetSummaryReport(v string)`

SetSummaryReport gets a reference to the given string and assigns it to the SummaryReport field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


