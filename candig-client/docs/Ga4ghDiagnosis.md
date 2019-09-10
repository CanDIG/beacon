# Ga4ghDiagnosis

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
**DiagnosisId** | Pointer to **string** |  | [optional] 
**DiagnosisDate** | Pointer to **string** |  | [optional] 
**AgeAtDiagnosis** | Pointer to **string** |  | [optional] 
**CancerType** | Pointer to **string** |  | [optional] 
**Classification** | Pointer to **string** |  | [optional] 
**CancerSite** | Pointer to **string** |  | [optional] 
**Histology** | Pointer to **string** |  | [optional] 
**MethodOfDefinitiveDiagnosis** | Pointer to **string** |  | [optional] 
**SampleType** | Pointer to **string** |  | [optional] 
**SampleSite** | Pointer to **string** |  | [optional] 
**TumorGrade** | Pointer to **string** |  | [optional] 
**GradingSystemUsed** | Pointer to **string** |  | [optional] 
**SitesOfMetastases** | Pointer to **string** |  | [optional] 
**StagingSystem** | Pointer to **string** |  | [optional] 
**VersionOrEditionOfTheStagingSystem** | Pointer to **string** |  | [optional] 
**SpecificTumorStageAtDiagnosis** | Pointer to **string** |  | [optional] 
**PrognosticBiomarkers** | Pointer to **string** |  | [optional] 
**BiomarkerQuantification** | Pointer to **string** |  | [optional] 
**AdditionalMolecularTesting** | Pointer to **string** |  | [optional] 
**AdditionalTestType** | Pointer to **string** |  | [optional] 
**LaboratoryName** | Pointer to **string** |  | [optional] 
**LaboratoryAddress** | Pointer to **string** |  | [optional] 
**SiteOfMetastases** | Pointer to **string** |  | [optional] 
**StagingSystemVersion** | Pointer to **string** |  | [optional] 
**SpecificStage** | Pointer to **string** |  | [optional] 
**CancerSpecificBiomarkers** | Pointer to **string** |  | [optional] 
**AdditionalMolecularDiagnosticTestingPerformed** | Pointer to **string** |  | [optional] 
**AdditionalTest** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghDiagnosis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghDiagnosis) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghDiagnosis) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghDiagnosis) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghDiagnosis) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghDiagnosis) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghDiagnosis) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghDiagnosis) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghDiagnosis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghDiagnosis) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghDiagnosis) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghDiagnosis) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghDiagnosis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghDiagnosis) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghDiagnosis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghDiagnosis) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghDiagnosis) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghDiagnosis) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghDiagnosis) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghDiagnosis) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghDiagnosis) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghDiagnosis) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghDiagnosis) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghDiagnosis) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghDiagnosis) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghDiagnosis) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghDiagnosis) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghDiagnosis) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetDiagnosisId

`func (o *Ga4ghDiagnosis) GetDiagnosisId() string`

GetDiagnosisId returns the DiagnosisId field if non-nil, zero value otherwise.

### GetDiagnosisIdOk

`func (o *Ga4ghDiagnosis) GetDiagnosisIdOk() (string, bool)`

GetDiagnosisIdOk returns a tuple with the DiagnosisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosisId

`func (o *Ga4ghDiagnosis) HasDiagnosisId() bool`

HasDiagnosisId returns a boolean if a field has been set.

### SetDiagnosisId

`func (o *Ga4ghDiagnosis) SetDiagnosisId(v string)`

SetDiagnosisId gets a reference to the given string and assigns it to the DiagnosisId field.

### GetDiagnosisDate

`func (o *Ga4ghDiagnosis) GetDiagnosisDate() string`

GetDiagnosisDate returns the DiagnosisDate field if non-nil, zero value otherwise.

### GetDiagnosisDateOk

`func (o *Ga4ghDiagnosis) GetDiagnosisDateOk() (string, bool)`

GetDiagnosisDateOk returns a tuple with the DiagnosisDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosisDate

`func (o *Ga4ghDiagnosis) HasDiagnosisDate() bool`

HasDiagnosisDate returns a boolean if a field has been set.

### SetDiagnosisDate

`func (o *Ga4ghDiagnosis) SetDiagnosisDate(v string)`

SetDiagnosisDate gets a reference to the given string and assigns it to the DiagnosisDate field.

### GetAgeAtDiagnosis

`func (o *Ga4ghDiagnosis) GetAgeAtDiagnosis() string`

GetAgeAtDiagnosis returns the AgeAtDiagnosis field if non-nil, zero value otherwise.

### GetAgeAtDiagnosisOk

`func (o *Ga4ghDiagnosis) GetAgeAtDiagnosisOk() (string, bool)`

GetAgeAtDiagnosisOk returns a tuple with the AgeAtDiagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgeAtDiagnosis

`func (o *Ga4ghDiagnosis) HasAgeAtDiagnosis() bool`

HasAgeAtDiagnosis returns a boolean if a field has been set.

### SetAgeAtDiagnosis

`func (o *Ga4ghDiagnosis) SetAgeAtDiagnosis(v string)`

SetAgeAtDiagnosis gets a reference to the given string and assigns it to the AgeAtDiagnosis field.

### GetCancerType

`func (o *Ga4ghDiagnosis) GetCancerType() string`

GetCancerType returns the CancerType field if non-nil, zero value otherwise.

### GetCancerTypeOk

`func (o *Ga4ghDiagnosis) GetCancerTypeOk() (string, bool)`

GetCancerTypeOk returns a tuple with the CancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancerType

`func (o *Ga4ghDiagnosis) HasCancerType() bool`

HasCancerType returns a boolean if a field has been set.

### SetCancerType

`func (o *Ga4ghDiagnosis) SetCancerType(v string)`

SetCancerType gets a reference to the given string and assigns it to the CancerType field.

### GetClassification

`func (o *Ga4ghDiagnosis) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Ga4ghDiagnosis) GetClassificationOk() (string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassification

`func (o *Ga4ghDiagnosis) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassification

`func (o *Ga4ghDiagnosis) SetClassification(v string)`

SetClassification gets a reference to the given string and assigns it to the Classification field.

### GetCancerSite

`func (o *Ga4ghDiagnosis) GetCancerSite() string`

GetCancerSite returns the CancerSite field if non-nil, zero value otherwise.

### GetCancerSiteOk

`func (o *Ga4ghDiagnosis) GetCancerSiteOk() (string, bool)`

GetCancerSiteOk returns a tuple with the CancerSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancerSite

`func (o *Ga4ghDiagnosis) HasCancerSite() bool`

HasCancerSite returns a boolean if a field has been set.

### SetCancerSite

`func (o *Ga4ghDiagnosis) SetCancerSite(v string)`

SetCancerSite gets a reference to the given string and assigns it to the CancerSite field.

### GetHistology

`func (o *Ga4ghDiagnosis) GetHistology() string`

GetHistology returns the Histology field if non-nil, zero value otherwise.

### GetHistologyOk

`func (o *Ga4ghDiagnosis) GetHistologyOk() (string, bool)`

GetHistologyOk returns a tuple with the Histology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHistology

`func (o *Ga4ghDiagnosis) HasHistology() bool`

HasHistology returns a boolean if a field has been set.

### SetHistology

`func (o *Ga4ghDiagnosis) SetHistology(v string)`

SetHistology gets a reference to the given string and assigns it to the Histology field.

### GetMethodOfDefinitiveDiagnosis

`func (o *Ga4ghDiagnosis) GetMethodOfDefinitiveDiagnosis() string`

GetMethodOfDefinitiveDiagnosis returns the MethodOfDefinitiveDiagnosis field if non-nil, zero value otherwise.

### GetMethodOfDefinitiveDiagnosisOk

`func (o *Ga4ghDiagnosis) GetMethodOfDefinitiveDiagnosisOk() (string, bool)`

GetMethodOfDefinitiveDiagnosisOk returns a tuple with the MethodOfDefinitiveDiagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMethodOfDefinitiveDiagnosis

`func (o *Ga4ghDiagnosis) HasMethodOfDefinitiveDiagnosis() bool`

HasMethodOfDefinitiveDiagnosis returns a boolean if a field has been set.

### SetMethodOfDefinitiveDiagnosis

`func (o *Ga4ghDiagnosis) SetMethodOfDefinitiveDiagnosis(v string)`

SetMethodOfDefinitiveDiagnosis gets a reference to the given string and assigns it to the MethodOfDefinitiveDiagnosis field.

### GetSampleType

`func (o *Ga4ghDiagnosis) GetSampleType() string`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *Ga4ghDiagnosis) GetSampleTypeOk() (string, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleType

`func (o *Ga4ghDiagnosis) HasSampleType() bool`

HasSampleType returns a boolean if a field has been set.

### SetSampleType

`func (o *Ga4ghDiagnosis) SetSampleType(v string)`

SetSampleType gets a reference to the given string and assigns it to the SampleType field.

### GetSampleSite

`func (o *Ga4ghDiagnosis) GetSampleSite() string`

GetSampleSite returns the SampleSite field if non-nil, zero value otherwise.

### GetSampleSiteOk

`func (o *Ga4ghDiagnosis) GetSampleSiteOk() (string, bool)`

GetSampleSiteOk returns a tuple with the SampleSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleSite

`func (o *Ga4ghDiagnosis) HasSampleSite() bool`

HasSampleSite returns a boolean if a field has been set.

### SetSampleSite

`func (o *Ga4ghDiagnosis) SetSampleSite(v string)`

SetSampleSite gets a reference to the given string and assigns it to the SampleSite field.

### GetTumorGrade

`func (o *Ga4ghDiagnosis) GetTumorGrade() string`

GetTumorGrade returns the TumorGrade field if non-nil, zero value otherwise.

### GetTumorGradeOk

`func (o *Ga4ghDiagnosis) GetTumorGradeOk() (string, bool)`

GetTumorGradeOk returns a tuple with the TumorGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTumorGrade

`func (o *Ga4ghDiagnosis) HasTumorGrade() bool`

HasTumorGrade returns a boolean if a field has been set.

### SetTumorGrade

`func (o *Ga4ghDiagnosis) SetTumorGrade(v string)`

SetTumorGrade gets a reference to the given string and assigns it to the TumorGrade field.

### GetGradingSystemUsed

`func (o *Ga4ghDiagnosis) GetGradingSystemUsed() string`

GetGradingSystemUsed returns the GradingSystemUsed field if non-nil, zero value otherwise.

### GetGradingSystemUsedOk

`func (o *Ga4ghDiagnosis) GetGradingSystemUsedOk() (string, bool)`

GetGradingSystemUsedOk returns a tuple with the GradingSystemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGradingSystemUsed

`func (o *Ga4ghDiagnosis) HasGradingSystemUsed() bool`

HasGradingSystemUsed returns a boolean if a field has been set.

### SetGradingSystemUsed

`func (o *Ga4ghDiagnosis) SetGradingSystemUsed(v string)`

SetGradingSystemUsed gets a reference to the given string and assigns it to the GradingSystemUsed field.

### GetSitesOfMetastases

`func (o *Ga4ghDiagnosis) GetSitesOfMetastases() string`

GetSitesOfMetastases returns the SitesOfMetastases field if non-nil, zero value otherwise.

### GetSitesOfMetastasesOk

`func (o *Ga4ghDiagnosis) GetSitesOfMetastasesOk() (string, bool)`

GetSitesOfMetastasesOk returns a tuple with the SitesOfMetastases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSitesOfMetastases

`func (o *Ga4ghDiagnosis) HasSitesOfMetastases() bool`

HasSitesOfMetastases returns a boolean if a field has been set.

### SetSitesOfMetastases

`func (o *Ga4ghDiagnosis) SetSitesOfMetastases(v string)`

SetSitesOfMetastases gets a reference to the given string and assigns it to the SitesOfMetastases field.

### GetStagingSystem

`func (o *Ga4ghDiagnosis) GetStagingSystem() string`

GetStagingSystem returns the StagingSystem field if non-nil, zero value otherwise.

### GetStagingSystemOk

`func (o *Ga4ghDiagnosis) GetStagingSystemOk() (string, bool)`

GetStagingSystemOk returns a tuple with the StagingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStagingSystem

`func (o *Ga4ghDiagnosis) HasStagingSystem() bool`

HasStagingSystem returns a boolean if a field has been set.

### SetStagingSystem

`func (o *Ga4ghDiagnosis) SetStagingSystem(v string)`

SetStagingSystem gets a reference to the given string and assigns it to the StagingSystem field.

### GetVersionOrEditionOfTheStagingSystem

`func (o *Ga4ghDiagnosis) GetVersionOrEditionOfTheStagingSystem() string`

GetVersionOrEditionOfTheStagingSystem returns the VersionOrEditionOfTheStagingSystem field if non-nil, zero value otherwise.

### GetVersionOrEditionOfTheStagingSystemOk

`func (o *Ga4ghDiagnosis) GetVersionOrEditionOfTheStagingSystemOk() (string, bool)`

GetVersionOrEditionOfTheStagingSystemOk returns a tuple with the VersionOrEditionOfTheStagingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersionOrEditionOfTheStagingSystem

`func (o *Ga4ghDiagnosis) HasVersionOrEditionOfTheStagingSystem() bool`

HasVersionOrEditionOfTheStagingSystem returns a boolean if a field has been set.

### SetVersionOrEditionOfTheStagingSystem

`func (o *Ga4ghDiagnosis) SetVersionOrEditionOfTheStagingSystem(v string)`

SetVersionOrEditionOfTheStagingSystem gets a reference to the given string and assigns it to the VersionOrEditionOfTheStagingSystem field.

### GetSpecificTumorStageAtDiagnosis

`func (o *Ga4ghDiagnosis) GetSpecificTumorStageAtDiagnosis() string`

GetSpecificTumorStageAtDiagnosis returns the SpecificTumorStageAtDiagnosis field if non-nil, zero value otherwise.

### GetSpecificTumorStageAtDiagnosisOk

`func (o *Ga4ghDiagnosis) GetSpecificTumorStageAtDiagnosisOk() (string, bool)`

GetSpecificTumorStageAtDiagnosisOk returns a tuple with the SpecificTumorStageAtDiagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecificTumorStageAtDiagnosis

`func (o *Ga4ghDiagnosis) HasSpecificTumorStageAtDiagnosis() bool`

HasSpecificTumorStageAtDiagnosis returns a boolean if a field has been set.

### SetSpecificTumorStageAtDiagnosis

`func (o *Ga4ghDiagnosis) SetSpecificTumorStageAtDiagnosis(v string)`

SetSpecificTumorStageAtDiagnosis gets a reference to the given string and assigns it to the SpecificTumorStageAtDiagnosis field.

### GetPrognosticBiomarkers

`func (o *Ga4ghDiagnosis) GetPrognosticBiomarkers() string`

GetPrognosticBiomarkers returns the PrognosticBiomarkers field if non-nil, zero value otherwise.

### GetPrognosticBiomarkersOk

`func (o *Ga4ghDiagnosis) GetPrognosticBiomarkersOk() (string, bool)`

GetPrognosticBiomarkersOk returns a tuple with the PrognosticBiomarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrognosticBiomarkers

`func (o *Ga4ghDiagnosis) HasPrognosticBiomarkers() bool`

HasPrognosticBiomarkers returns a boolean if a field has been set.

### SetPrognosticBiomarkers

`func (o *Ga4ghDiagnosis) SetPrognosticBiomarkers(v string)`

SetPrognosticBiomarkers gets a reference to the given string and assigns it to the PrognosticBiomarkers field.

### GetBiomarkerQuantification

`func (o *Ga4ghDiagnosis) GetBiomarkerQuantification() string`

GetBiomarkerQuantification returns the BiomarkerQuantification field if non-nil, zero value otherwise.

### GetBiomarkerQuantificationOk

`func (o *Ga4ghDiagnosis) GetBiomarkerQuantificationOk() (string, bool)`

GetBiomarkerQuantificationOk returns a tuple with the BiomarkerQuantification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiomarkerQuantification

`func (o *Ga4ghDiagnosis) HasBiomarkerQuantification() bool`

HasBiomarkerQuantification returns a boolean if a field has been set.

### SetBiomarkerQuantification

`func (o *Ga4ghDiagnosis) SetBiomarkerQuantification(v string)`

SetBiomarkerQuantification gets a reference to the given string and assigns it to the BiomarkerQuantification field.

### GetAdditionalMolecularTesting

`func (o *Ga4ghDiagnosis) GetAdditionalMolecularTesting() string`

GetAdditionalMolecularTesting returns the AdditionalMolecularTesting field if non-nil, zero value otherwise.

### GetAdditionalMolecularTestingOk

`func (o *Ga4ghDiagnosis) GetAdditionalMolecularTestingOk() (string, bool)`

GetAdditionalMolecularTestingOk returns a tuple with the AdditionalMolecularTesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalMolecularTesting

`func (o *Ga4ghDiagnosis) HasAdditionalMolecularTesting() bool`

HasAdditionalMolecularTesting returns a boolean if a field has been set.

### SetAdditionalMolecularTesting

`func (o *Ga4ghDiagnosis) SetAdditionalMolecularTesting(v string)`

SetAdditionalMolecularTesting gets a reference to the given string and assigns it to the AdditionalMolecularTesting field.

### GetAdditionalTestType

`func (o *Ga4ghDiagnosis) GetAdditionalTestType() string`

GetAdditionalTestType returns the AdditionalTestType field if non-nil, zero value otherwise.

### GetAdditionalTestTypeOk

`func (o *Ga4ghDiagnosis) GetAdditionalTestTypeOk() (string, bool)`

GetAdditionalTestTypeOk returns a tuple with the AdditionalTestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalTestType

`func (o *Ga4ghDiagnosis) HasAdditionalTestType() bool`

HasAdditionalTestType returns a boolean if a field has been set.

### SetAdditionalTestType

`func (o *Ga4ghDiagnosis) SetAdditionalTestType(v string)`

SetAdditionalTestType gets a reference to the given string and assigns it to the AdditionalTestType field.

### GetLaboratoryName

`func (o *Ga4ghDiagnosis) GetLaboratoryName() string`

GetLaboratoryName returns the LaboratoryName field if non-nil, zero value otherwise.

### GetLaboratoryNameOk

`func (o *Ga4ghDiagnosis) GetLaboratoryNameOk() (string, bool)`

GetLaboratoryNameOk returns a tuple with the LaboratoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLaboratoryName

`func (o *Ga4ghDiagnosis) HasLaboratoryName() bool`

HasLaboratoryName returns a boolean if a field has been set.

### SetLaboratoryName

`func (o *Ga4ghDiagnosis) SetLaboratoryName(v string)`

SetLaboratoryName gets a reference to the given string and assigns it to the LaboratoryName field.

### GetLaboratoryAddress

`func (o *Ga4ghDiagnosis) GetLaboratoryAddress() string`

GetLaboratoryAddress returns the LaboratoryAddress field if non-nil, zero value otherwise.

### GetLaboratoryAddressOk

`func (o *Ga4ghDiagnosis) GetLaboratoryAddressOk() (string, bool)`

GetLaboratoryAddressOk returns a tuple with the LaboratoryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLaboratoryAddress

`func (o *Ga4ghDiagnosis) HasLaboratoryAddress() bool`

HasLaboratoryAddress returns a boolean if a field has been set.

### SetLaboratoryAddress

`func (o *Ga4ghDiagnosis) SetLaboratoryAddress(v string)`

SetLaboratoryAddress gets a reference to the given string and assigns it to the LaboratoryAddress field.

### GetSiteOfMetastases

`func (o *Ga4ghDiagnosis) GetSiteOfMetastases() string`

GetSiteOfMetastases returns the SiteOfMetastases field if non-nil, zero value otherwise.

### GetSiteOfMetastasesOk

`func (o *Ga4ghDiagnosis) GetSiteOfMetastasesOk() (string, bool)`

GetSiteOfMetastasesOk returns a tuple with the SiteOfMetastases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteOfMetastases

`func (o *Ga4ghDiagnosis) HasSiteOfMetastases() bool`

HasSiteOfMetastases returns a boolean if a field has been set.

### SetSiteOfMetastases

`func (o *Ga4ghDiagnosis) SetSiteOfMetastases(v string)`

SetSiteOfMetastases gets a reference to the given string and assigns it to the SiteOfMetastases field.

### GetStagingSystemVersion

`func (o *Ga4ghDiagnosis) GetStagingSystemVersion() string`

GetStagingSystemVersion returns the StagingSystemVersion field if non-nil, zero value otherwise.

### GetStagingSystemVersionOk

`func (o *Ga4ghDiagnosis) GetStagingSystemVersionOk() (string, bool)`

GetStagingSystemVersionOk returns a tuple with the StagingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStagingSystemVersion

`func (o *Ga4ghDiagnosis) HasStagingSystemVersion() bool`

HasStagingSystemVersion returns a boolean if a field has been set.

### SetStagingSystemVersion

`func (o *Ga4ghDiagnosis) SetStagingSystemVersion(v string)`

SetStagingSystemVersion gets a reference to the given string and assigns it to the StagingSystemVersion field.

### GetSpecificStage

`func (o *Ga4ghDiagnosis) GetSpecificStage() string`

GetSpecificStage returns the SpecificStage field if non-nil, zero value otherwise.

### GetSpecificStageOk

`func (o *Ga4ghDiagnosis) GetSpecificStageOk() (string, bool)`

GetSpecificStageOk returns a tuple with the SpecificStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecificStage

`func (o *Ga4ghDiagnosis) HasSpecificStage() bool`

HasSpecificStage returns a boolean if a field has been set.

### SetSpecificStage

`func (o *Ga4ghDiagnosis) SetSpecificStage(v string)`

SetSpecificStage gets a reference to the given string and assigns it to the SpecificStage field.

### GetCancerSpecificBiomarkers

`func (o *Ga4ghDiagnosis) GetCancerSpecificBiomarkers() string`

GetCancerSpecificBiomarkers returns the CancerSpecificBiomarkers field if non-nil, zero value otherwise.

### GetCancerSpecificBiomarkersOk

`func (o *Ga4ghDiagnosis) GetCancerSpecificBiomarkersOk() (string, bool)`

GetCancerSpecificBiomarkersOk returns a tuple with the CancerSpecificBiomarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancerSpecificBiomarkers

`func (o *Ga4ghDiagnosis) HasCancerSpecificBiomarkers() bool`

HasCancerSpecificBiomarkers returns a boolean if a field has been set.

### SetCancerSpecificBiomarkers

`func (o *Ga4ghDiagnosis) SetCancerSpecificBiomarkers(v string)`

SetCancerSpecificBiomarkers gets a reference to the given string and assigns it to the CancerSpecificBiomarkers field.

### GetAdditionalMolecularDiagnosticTestingPerformed

`func (o *Ga4ghDiagnosis) GetAdditionalMolecularDiagnosticTestingPerformed() string`

GetAdditionalMolecularDiagnosticTestingPerformed returns the AdditionalMolecularDiagnosticTestingPerformed field if non-nil, zero value otherwise.

### GetAdditionalMolecularDiagnosticTestingPerformedOk

`func (o *Ga4ghDiagnosis) GetAdditionalMolecularDiagnosticTestingPerformedOk() (string, bool)`

GetAdditionalMolecularDiagnosticTestingPerformedOk returns a tuple with the AdditionalMolecularDiagnosticTestingPerformed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalMolecularDiagnosticTestingPerformed

`func (o *Ga4ghDiagnosis) HasAdditionalMolecularDiagnosticTestingPerformed() bool`

HasAdditionalMolecularDiagnosticTestingPerformed returns a boolean if a field has been set.

### SetAdditionalMolecularDiagnosticTestingPerformed

`func (o *Ga4ghDiagnosis) SetAdditionalMolecularDiagnosticTestingPerformed(v string)`

SetAdditionalMolecularDiagnosticTestingPerformed gets a reference to the given string and assigns it to the AdditionalMolecularDiagnosticTestingPerformed field.

### GetAdditionalTest

`func (o *Ga4ghDiagnosis) GetAdditionalTest() string`

GetAdditionalTest returns the AdditionalTest field if non-nil, zero value otherwise.

### GetAdditionalTestOk

`func (o *Ga4ghDiagnosis) GetAdditionalTestOk() (string, bool)`

GetAdditionalTestOk returns a tuple with the AdditionalTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalTest

`func (o *Ga4ghDiagnosis) HasAdditionalTest() bool`

HasAdditionalTest returns a boolean if a field has been set.

### SetAdditionalTest

`func (o *Ga4ghDiagnosis) SetAdditionalTest(v string)`

SetAdditionalTest gets a reference to the given string and assigns it to the AdditionalTest field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


