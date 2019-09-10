# Ga4ghSearchQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patients** | Pointer to [**[]Ga4ghPatient**](ga4ghPatient.md) |  | [optional] 
**Enrollments** | Pointer to [**[]Ga4ghEnrollment**](ga4ghEnrollment.md) |  | [optional] 
**Consents** | Pointer to [**[]Ga4ghConsent**](ga4ghConsent.md) |  | [optional] 
**Diagnoses** | Pointer to [**[]Ga4ghDiagnosis**](ga4ghDiagnosis.md) |  | [optional] 
**Samples** | Pointer to [**[]Ga4ghSample**](ga4ghSample.md) |  | [optional] 
**Treatments** | Pointer to [**[]Ga4ghTreatment**](ga4ghTreatment.md) |  | [optional] 
**Outcomes** | Pointer to [**[]Ga4ghOutcome**](ga4ghOutcome.md) |  | [optional] 
**Complications** | Pointer to [**[]Ga4ghComplication**](ga4ghComplication.md) |  | [optional] 
**Tumourboards** | Pointer to [**[]Ga4ghTumourboard**](ga4ghTumourboard.md) |  | [optional] 
**Variants** | Pointer to [**[]Ga4ghVariant**](ga4ghVariant.md) |  | [optional] 
**Slides** | Pointer to [**[]Ga4ghSlide**](ga4ghSlide.md) |  | [optional] 
**Studies** | Pointer to [**[]Ga4ghStudy**](ga4ghStudy.md) |  | [optional] 
**Labtests** | Pointer to [**[]Ga4ghLabtest**](ga4ghLabtest.md) |  | [optional] 
**Surgeries** | Pointer to [**[]Ga4ghSurgery**](ga4ghSurgery.md) |  | [optional] 
**Chemotherapies** | Pointer to [**[]Ga4ghChemotherapy**](ga4ghChemotherapy.md) |  | [optional] 
**Immunotherapies** | Pointer to [**[]Ga4ghImmunotherapy**](ga4ghImmunotherapy.md) |  | [optional] 
**Radiotherapies** | Pointer to [**[]Ga4ghRadiotherapy**](ga4ghRadiotherapy.md) |  | [optional] 
**Celltransplants** | Pointer to [**[]Ga4ghCelltransplant**](ga4ghCelltransplant.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren&#39;t any additional results. | [optional] 

## Methods

### GetPatients

`func (o *Ga4ghSearchQueryResponse) GetPatients() []Ga4ghPatient`

GetPatients returns the Patients field if non-nil, zero value otherwise.

### GetPatientsOk

`func (o *Ga4ghSearchQueryResponse) GetPatientsOk() ([]Ga4ghPatient, bool)`

GetPatientsOk returns a tuple with the Patients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatients

`func (o *Ga4ghSearchQueryResponse) HasPatients() bool`

HasPatients returns a boolean if a field has been set.

### SetPatients

`func (o *Ga4ghSearchQueryResponse) SetPatients(v []Ga4ghPatient)`

SetPatients gets a reference to the given []Ga4ghPatient and assigns it to the Patients field.

### GetEnrollments

`func (o *Ga4ghSearchQueryResponse) GetEnrollments() []Ga4ghEnrollment`

GetEnrollments returns the Enrollments field if non-nil, zero value otherwise.

### GetEnrollmentsOk

`func (o *Ga4ghSearchQueryResponse) GetEnrollmentsOk() ([]Ga4ghEnrollment, bool)`

GetEnrollmentsOk returns a tuple with the Enrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollments

`func (o *Ga4ghSearchQueryResponse) HasEnrollments() bool`

HasEnrollments returns a boolean if a field has been set.

### SetEnrollments

`func (o *Ga4ghSearchQueryResponse) SetEnrollments(v []Ga4ghEnrollment)`

SetEnrollments gets a reference to the given []Ga4ghEnrollment and assigns it to the Enrollments field.

### GetConsents

`func (o *Ga4ghSearchQueryResponse) GetConsents() []Ga4ghConsent`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *Ga4ghSearchQueryResponse) GetConsentsOk() ([]Ga4ghConsent, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsents

`func (o *Ga4ghSearchQueryResponse) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### SetConsents

`func (o *Ga4ghSearchQueryResponse) SetConsents(v []Ga4ghConsent)`

SetConsents gets a reference to the given []Ga4ghConsent and assigns it to the Consents field.

### GetDiagnoses

`func (o *Ga4ghSearchQueryResponse) GetDiagnoses() []Ga4ghDiagnosis`

GetDiagnoses returns the Diagnoses field if non-nil, zero value otherwise.

### GetDiagnosesOk

`func (o *Ga4ghSearchQueryResponse) GetDiagnosesOk() ([]Ga4ghDiagnosis, bool)`

GetDiagnosesOk returns a tuple with the Diagnoses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnoses

`func (o *Ga4ghSearchQueryResponse) HasDiagnoses() bool`

HasDiagnoses returns a boolean if a field has been set.

### SetDiagnoses

`func (o *Ga4ghSearchQueryResponse) SetDiagnoses(v []Ga4ghDiagnosis)`

SetDiagnoses gets a reference to the given []Ga4ghDiagnosis and assigns it to the Diagnoses field.

### GetSamples

`func (o *Ga4ghSearchQueryResponse) GetSamples() []Ga4ghSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *Ga4ghSearchQueryResponse) GetSamplesOk() ([]Ga4ghSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamples

`func (o *Ga4ghSearchQueryResponse) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### SetSamples

`func (o *Ga4ghSearchQueryResponse) SetSamples(v []Ga4ghSample)`

SetSamples gets a reference to the given []Ga4ghSample and assigns it to the Samples field.

### GetTreatments

`func (o *Ga4ghSearchQueryResponse) GetTreatments() []Ga4ghTreatment`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *Ga4ghSearchQueryResponse) GetTreatmentsOk() ([]Ga4ghTreatment, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatments

`func (o *Ga4ghSearchQueryResponse) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### SetTreatments

`func (o *Ga4ghSearchQueryResponse) SetTreatments(v []Ga4ghTreatment)`

SetTreatments gets a reference to the given []Ga4ghTreatment and assigns it to the Treatments field.

### GetOutcomes

`func (o *Ga4ghSearchQueryResponse) GetOutcomes() []Ga4ghOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *Ga4ghSearchQueryResponse) GetOutcomesOk() ([]Ga4ghOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutcomes

`func (o *Ga4ghSearchQueryResponse) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomes

`func (o *Ga4ghSearchQueryResponse) SetOutcomes(v []Ga4ghOutcome)`

SetOutcomes gets a reference to the given []Ga4ghOutcome and assigns it to the Outcomes field.

### GetComplications

`func (o *Ga4ghSearchQueryResponse) GetComplications() []Ga4ghComplication`

GetComplications returns the Complications field if non-nil, zero value otherwise.

### GetComplicationsOk

`func (o *Ga4ghSearchQueryResponse) GetComplicationsOk() ([]Ga4ghComplication, bool)`

GetComplicationsOk returns a tuple with the Complications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplications

`func (o *Ga4ghSearchQueryResponse) HasComplications() bool`

HasComplications returns a boolean if a field has been set.

### SetComplications

`func (o *Ga4ghSearchQueryResponse) SetComplications(v []Ga4ghComplication)`

SetComplications gets a reference to the given []Ga4ghComplication and assigns it to the Complications field.

### GetTumourboards

`func (o *Ga4ghSearchQueryResponse) GetTumourboards() []Ga4ghTumourboard`

GetTumourboards returns the Tumourboards field if non-nil, zero value otherwise.

### GetTumourboardsOk

`func (o *Ga4ghSearchQueryResponse) GetTumourboardsOk() ([]Ga4ghTumourboard, bool)`

GetTumourboardsOk returns a tuple with the Tumourboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTumourboards

`func (o *Ga4ghSearchQueryResponse) HasTumourboards() bool`

HasTumourboards returns a boolean if a field has been set.

### SetTumourboards

`func (o *Ga4ghSearchQueryResponse) SetTumourboards(v []Ga4ghTumourboard)`

SetTumourboards gets a reference to the given []Ga4ghTumourboard and assigns it to the Tumourboards field.

### GetVariants

`func (o *Ga4ghSearchQueryResponse) GetVariants() []Ga4ghVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Ga4ghSearchQueryResponse) GetVariantsOk() ([]Ga4ghVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariants

`func (o *Ga4ghSearchQueryResponse) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariants

`func (o *Ga4ghSearchQueryResponse) SetVariants(v []Ga4ghVariant)`

SetVariants gets a reference to the given []Ga4ghVariant and assigns it to the Variants field.

### GetSlides

`func (o *Ga4ghSearchQueryResponse) GetSlides() []Ga4ghSlide`

GetSlides returns the Slides field if non-nil, zero value otherwise.

### GetSlidesOk

`func (o *Ga4ghSearchQueryResponse) GetSlidesOk() ([]Ga4ghSlide, bool)`

GetSlidesOk returns a tuple with the Slides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSlides

`func (o *Ga4ghSearchQueryResponse) HasSlides() bool`

HasSlides returns a boolean if a field has been set.

### SetSlides

`func (o *Ga4ghSearchQueryResponse) SetSlides(v []Ga4ghSlide)`

SetSlides gets a reference to the given []Ga4ghSlide and assigns it to the Slides field.

### GetStudies

`func (o *Ga4ghSearchQueryResponse) GetStudies() []Ga4ghStudy`

GetStudies returns the Studies field if non-nil, zero value otherwise.

### GetStudiesOk

`func (o *Ga4ghSearchQueryResponse) GetStudiesOk() ([]Ga4ghStudy, bool)`

GetStudiesOk returns a tuple with the Studies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStudies

`func (o *Ga4ghSearchQueryResponse) HasStudies() bool`

HasStudies returns a boolean if a field has been set.

### SetStudies

`func (o *Ga4ghSearchQueryResponse) SetStudies(v []Ga4ghStudy)`

SetStudies gets a reference to the given []Ga4ghStudy and assigns it to the Studies field.

### GetLabtests

`func (o *Ga4ghSearchQueryResponse) GetLabtests() []Ga4ghLabtest`

GetLabtests returns the Labtests field if non-nil, zero value otherwise.

### GetLabtestsOk

`func (o *Ga4ghSearchQueryResponse) GetLabtestsOk() ([]Ga4ghLabtest, bool)`

GetLabtestsOk returns a tuple with the Labtests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabtests

`func (o *Ga4ghSearchQueryResponse) HasLabtests() bool`

HasLabtests returns a boolean if a field has been set.

### SetLabtests

`func (o *Ga4ghSearchQueryResponse) SetLabtests(v []Ga4ghLabtest)`

SetLabtests gets a reference to the given []Ga4ghLabtest and assigns it to the Labtests field.

### GetSurgeries

`func (o *Ga4ghSearchQueryResponse) GetSurgeries() []Ga4ghSurgery`

GetSurgeries returns the Surgeries field if non-nil, zero value otherwise.

### GetSurgeriesOk

`func (o *Ga4ghSearchQueryResponse) GetSurgeriesOk() ([]Ga4ghSurgery, bool)`

GetSurgeriesOk returns a tuple with the Surgeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurgeries

`func (o *Ga4ghSearchQueryResponse) HasSurgeries() bool`

HasSurgeries returns a boolean if a field has been set.

### SetSurgeries

`func (o *Ga4ghSearchQueryResponse) SetSurgeries(v []Ga4ghSurgery)`

SetSurgeries gets a reference to the given []Ga4ghSurgery and assigns it to the Surgeries field.

### GetChemotherapies

`func (o *Ga4ghSearchQueryResponse) GetChemotherapies() []Ga4ghChemotherapy`

GetChemotherapies returns the Chemotherapies field if non-nil, zero value otherwise.

### GetChemotherapiesOk

`func (o *Ga4ghSearchQueryResponse) GetChemotherapiesOk() ([]Ga4ghChemotherapy, bool)`

GetChemotherapiesOk returns a tuple with the Chemotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChemotherapies

`func (o *Ga4ghSearchQueryResponse) HasChemotherapies() bool`

HasChemotherapies returns a boolean if a field has been set.

### SetChemotherapies

`func (o *Ga4ghSearchQueryResponse) SetChemotherapies(v []Ga4ghChemotherapy)`

SetChemotherapies gets a reference to the given []Ga4ghChemotherapy and assigns it to the Chemotherapies field.

### GetImmunotherapies

`func (o *Ga4ghSearchQueryResponse) GetImmunotherapies() []Ga4ghImmunotherapy`

GetImmunotherapies returns the Immunotherapies field if non-nil, zero value otherwise.

### GetImmunotherapiesOk

`func (o *Ga4ghSearchQueryResponse) GetImmunotherapiesOk() ([]Ga4ghImmunotherapy, bool)`

GetImmunotherapiesOk returns a tuple with the Immunotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImmunotherapies

`func (o *Ga4ghSearchQueryResponse) HasImmunotherapies() bool`

HasImmunotherapies returns a boolean if a field has been set.

### SetImmunotherapies

`func (o *Ga4ghSearchQueryResponse) SetImmunotherapies(v []Ga4ghImmunotherapy)`

SetImmunotherapies gets a reference to the given []Ga4ghImmunotherapy and assigns it to the Immunotherapies field.

### GetRadiotherapies

`func (o *Ga4ghSearchQueryResponse) GetRadiotherapies() []Ga4ghRadiotherapy`

GetRadiotherapies returns the Radiotherapies field if non-nil, zero value otherwise.

### GetRadiotherapiesOk

`func (o *Ga4ghSearchQueryResponse) GetRadiotherapiesOk() ([]Ga4ghRadiotherapy, bool)`

GetRadiotherapiesOk returns a tuple with the Radiotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRadiotherapies

`func (o *Ga4ghSearchQueryResponse) HasRadiotherapies() bool`

HasRadiotherapies returns a boolean if a field has been set.

### SetRadiotherapies

`func (o *Ga4ghSearchQueryResponse) SetRadiotherapies(v []Ga4ghRadiotherapy)`

SetRadiotherapies gets a reference to the given []Ga4ghRadiotherapy and assigns it to the Radiotherapies field.

### GetCelltransplants

`func (o *Ga4ghSearchQueryResponse) GetCelltransplants() []Ga4ghCelltransplant`

GetCelltransplants returns the Celltransplants field if non-nil, zero value otherwise.

### GetCelltransplantsOk

`func (o *Ga4ghSearchQueryResponse) GetCelltransplantsOk() ([]Ga4ghCelltransplant, bool)`

GetCelltransplantsOk returns a tuple with the Celltransplants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCelltransplants

`func (o *Ga4ghSearchQueryResponse) HasCelltransplants() bool`

HasCelltransplants returns a boolean if a field has been set.

### SetCelltransplants

`func (o *Ga4ghSearchQueryResponse) SetCelltransplants(v []Ga4ghCelltransplant)`

SetCelltransplants gets a reference to the given []Ga4ghCelltransplant and assigns it to the Celltransplants field.

### GetNextPageToken

`func (o *Ga4ghSearchQueryResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Ga4ghSearchQueryResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *Ga4ghSearchQueryResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *Ga4ghSearchQueryResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


