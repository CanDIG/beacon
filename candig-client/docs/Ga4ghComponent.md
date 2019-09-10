# Ga4ghComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Patients** | Pointer to [**Ga4ghSearchPatientsRequest**](ga4ghSearchPatientsRequest.md) |  | [optional] 
**Enrollments** | Pointer to [**Ga4ghSearchEnrollmentsRequest**](ga4ghSearchEnrollmentsRequest.md) |  | [optional] 
**Consents** | Pointer to [**Ga4ghSearchConsentsRequest**](ga4ghSearchConsentsRequest.md) |  | [optional] 
**Diagnoses** | Pointer to [**Ga4ghSearchDiagnosesRequest**](ga4ghSearchDiagnosesRequest.md) |  | [optional] 
**Samples** | Pointer to [**Ga4ghSearchSamplesRequest**](ga4ghSearchSamplesRequest.md) |  | [optional] 
**Treatments** | Pointer to [**Ga4ghSearchTreatmentsRequest**](ga4ghSearchTreatmentsRequest.md) |  | [optional] 
**Outcomes** | Pointer to [**Ga4ghSearchOutcomesRequest**](ga4ghSearchOutcomesRequest.md) |  | [optional] 
**Complications** | Pointer to [**Ga4ghSearchComplicationsRequest**](ga4ghSearchComplicationsRequest.md) |  | [optional] 
**Tumourboards** | Pointer to [**Ga4ghSearchTumourboardsRequest**](ga4ghSearchTumourboardsRequest.md) |  | [optional] 
**Variants** | Pointer to [**Ga4ghSearchVariantsRequest**](ga4ghSearchVariantsRequest.md) |  | [optional] 
**VariantsByGene** | Pointer to [**Ga4ghSearchVariantsByGeneNameRequest**](ga4ghSearchVariantsByGeneNameRequest.md) |  | [optional] 
**Slides** | Pointer to [**Ga4ghSearchSlidesRequest**](ga4ghSearchSlidesRequest.md) |  | [optional] 
**Studies** | Pointer to [**Ga4ghSearchStudiesRequest**](ga4ghSearchStudiesRequest.md) |  | [optional] 
**Labtests** | Pointer to [**Ga4ghSearchLabtestsRequest**](ga4ghSearchLabtestsRequest.md) |  | [optional] 
**Chemotherapies** | Pointer to [**Ga4ghSearchChemotherapiesRequest**](ga4ghSearchChemotherapiesRequest.md) |  | [optional] 
**Radiotherapies** | Pointer to [**Ga4ghSearchRadiotherapiesRequest**](ga4ghSearchRadiotherapiesRequest.md) |  | [optional] 
**Immunotherapies** | Pointer to [**Ga4ghSearchImmunotherapiesRequest**](ga4ghSearchImmunotherapiesRequest.md) |  | [optional] 
**Surgeries** | Pointer to [**Ga4ghSearchSurgeriesRequest**](ga4ghSearchSurgeriesRequest.md) |  | [optional] 
**Celltransplants** | Pointer to [**Ga4ghSearchCelltransplantsRequest**](ga4ghSearchCelltransplantsRequest.md) |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghComponent) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghComponent) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetPatients

`func (o *Ga4ghComponent) GetPatients() Ga4ghSearchPatientsRequest`

GetPatients returns the Patients field if non-nil, zero value otherwise.

### GetPatientsOk

`func (o *Ga4ghComponent) GetPatientsOk() (Ga4ghSearchPatientsRequest, bool)`

GetPatientsOk returns a tuple with the Patients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatients

`func (o *Ga4ghComponent) HasPatients() bool`

HasPatients returns a boolean if a field has been set.

### SetPatients

`func (o *Ga4ghComponent) SetPatients(v Ga4ghSearchPatientsRequest)`

SetPatients gets a reference to the given Ga4ghSearchPatientsRequest and assigns it to the Patients field.

### GetEnrollments

`func (o *Ga4ghComponent) GetEnrollments() Ga4ghSearchEnrollmentsRequest`

GetEnrollments returns the Enrollments field if non-nil, zero value otherwise.

### GetEnrollmentsOk

`func (o *Ga4ghComponent) GetEnrollmentsOk() (Ga4ghSearchEnrollmentsRequest, bool)`

GetEnrollmentsOk returns a tuple with the Enrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollments

`func (o *Ga4ghComponent) HasEnrollments() bool`

HasEnrollments returns a boolean if a field has been set.

### SetEnrollments

`func (o *Ga4ghComponent) SetEnrollments(v Ga4ghSearchEnrollmentsRequest)`

SetEnrollments gets a reference to the given Ga4ghSearchEnrollmentsRequest and assigns it to the Enrollments field.

### GetConsents

`func (o *Ga4ghComponent) GetConsents() Ga4ghSearchConsentsRequest`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *Ga4ghComponent) GetConsentsOk() (Ga4ghSearchConsentsRequest, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsents

`func (o *Ga4ghComponent) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### SetConsents

`func (o *Ga4ghComponent) SetConsents(v Ga4ghSearchConsentsRequest)`

SetConsents gets a reference to the given Ga4ghSearchConsentsRequest and assigns it to the Consents field.

### GetDiagnoses

`func (o *Ga4ghComponent) GetDiagnoses() Ga4ghSearchDiagnosesRequest`

GetDiagnoses returns the Diagnoses field if non-nil, zero value otherwise.

### GetDiagnosesOk

`func (o *Ga4ghComponent) GetDiagnosesOk() (Ga4ghSearchDiagnosesRequest, bool)`

GetDiagnosesOk returns a tuple with the Diagnoses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnoses

`func (o *Ga4ghComponent) HasDiagnoses() bool`

HasDiagnoses returns a boolean if a field has been set.

### SetDiagnoses

`func (o *Ga4ghComponent) SetDiagnoses(v Ga4ghSearchDiagnosesRequest)`

SetDiagnoses gets a reference to the given Ga4ghSearchDiagnosesRequest and assigns it to the Diagnoses field.

### GetSamples

`func (o *Ga4ghComponent) GetSamples() Ga4ghSearchSamplesRequest`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *Ga4ghComponent) GetSamplesOk() (Ga4ghSearchSamplesRequest, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamples

`func (o *Ga4ghComponent) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### SetSamples

`func (o *Ga4ghComponent) SetSamples(v Ga4ghSearchSamplesRequest)`

SetSamples gets a reference to the given Ga4ghSearchSamplesRequest and assigns it to the Samples field.

### GetTreatments

`func (o *Ga4ghComponent) GetTreatments() Ga4ghSearchTreatmentsRequest`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *Ga4ghComponent) GetTreatmentsOk() (Ga4ghSearchTreatmentsRequest, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatments

`func (o *Ga4ghComponent) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### SetTreatments

`func (o *Ga4ghComponent) SetTreatments(v Ga4ghSearchTreatmentsRequest)`

SetTreatments gets a reference to the given Ga4ghSearchTreatmentsRequest and assigns it to the Treatments field.

### GetOutcomes

`func (o *Ga4ghComponent) GetOutcomes() Ga4ghSearchOutcomesRequest`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *Ga4ghComponent) GetOutcomesOk() (Ga4ghSearchOutcomesRequest, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutcomes

`func (o *Ga4ghComponent) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomes

`func (o *Ga4ghComponent) SetOutcomes(v Ga4ghSearchOutcomesRequest)`

SetOutcomes gets a reference to the given Ga4ghSearchOutcomesRequest and assigns it to the Outcomes field.

### GetComplications

`func (o *Ga4ghComponent) GetComplications() Ga4ghSearchComplicationsRequest`

GetComplications returns the Complications field if non-nil, zero value otherwise.

### GetComplicationsOk

`func (o *Ga4ghComponent) GetComplicationsOk() (Ga4ghSearchComplicationsRequest, bool)`

GetComplicationsOk returns a tuple with the Complications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplications

`func (o *Ga4ghComponent) HasComplications() bool`

HasComplications returns a boolean if a field has been set.

### SetComplications

`func (o *Ga4ghComponent) SetComplications(v Ga4ghSearchComplicationsRequest)`

SetComplications gets a reference to the given Ga4ghSearchComplicationsRequest and assigns it to the Complications field.

### GetTumourboards

`func (o *Ga4ghComponent) GetTumourboards() Ga4ghSearchTumourboardsRequest`

GetTumourboards returns the Tumourboards field if non-nil, zero value otherwise.

### GetTumourboardsOk

`func (o *Ga4ghComponent) GetTumourboardsOk() (Ga4ghSearchTumourboardsRequest, bool)`

GetTumourboardsOk returns a tuple with the Tumourboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTumourboards

`func (o *Ga4ghComponent) HasTumourboards() bool`

HasTumourboards returns a boolean if a field has been set.

### SetTumourboards

`func (o *Ga4ghComponent) SetTumourboards(v Ga4ghSearchTumourboardsRequest)`

SetTumourboards gets a reference to the given Ga4ghSearchTumourboardsRequest and assigns it to the Tumourboards field.

### GetVariants

`func (o *Ga4ghComponent) GetVariants() Ga4ghSearchVariantsRequest`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Ga4ghComponent) GetVariantsOk() (Ga4ghSearchVariantsRequest, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariants

`func (o *Ga4ghComponent) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariants

`func (o *Ga4ghComponent) SetVariants(v Ga4ghSearchVariantsRequest)`

SetVariants gets a reference to the given Ga4ghSearchVariantsRequest and assigns it to the Variants field.

### GetVariantsByGene

`func (o *Ga4ghComponent) GetVariantsByGene() Ga4ghSearchVariantsByGeneNameRequest`

GetVariantsByGene returns the VariantsByGene field if non-nil, zero value otherwise.

### GetVariantsByGeneOk

`func (o *Ga4ghComponent) GetVariantsByGeneOk() (Ga4ghSearchVariantsByGeneNameRequest, bool)`

GetVariantsByGeneOk returns a tuple with the VariantsByGene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariantsByGene

`func (o *Ga4ghComponent) HasVariantsByGene() bool`

HasVariantsByGene returns a boolean if a field has been set.

### SetVariantsByGene

`func (o *Ga4ghComponent) SetVariantsByGene(v Ga4ghSearchVariantsByGeneNameRequest)`

SetVariantsByGene gets a reference to the given Ga4ghSearchVariantsByGeneNameRequest and assigns it to the VariantsByGene field.

### GetSlides

`func (o *Ga4ghComponent) GetSlides() Ga4ghSearchSlidesRequest`

GetSlides returns the Slides field if non-nil, zero value otherwise.

### GetSlidesOk

`func (o *Ga4ghComponent) GetSlidesOk() (Ga4ghSearchSlidesRequest, bool)`

GetSlidesOk returns a tuple with the Slides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSlides

`func (o *Ga4ghComponent) HasSlides() bool`

HasSlides returns a boolean if a field has been set.

### SetSlides

`func (o *Ga4ghComponent) SetSlides(v Ga4ghSearchSlidesRequest)`

SetSlides gets a reference to the given Ga4ghSearchSlidesRequest and assigns it to the Slides field.

### GetStudies

`func (o *Ga4ghComponent) GetStudies() Ga4ghSearchStudiesRequest`

GetStudies returns the Studies field if non-nil, zero value otherwise.

### GetStudiesOk

`func (o *Ga4ghComponent) GetStudiesOk() (Ga4ghSearchStudiesRequest, bool)`

GetStudiesOk returns a tuple with the Studies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStudies

`func (o *Ga4ghComponent) HasStudies() bool`

HasStudies returns a boolean if a field has been set.

### SetStudies

`func (o *Ga4ghComponent) SetStudies(v Ga4ghSearchStudiesRequest)`

SetStudies gets a reference to the given Ga4ghSearchStudiesRequest and assigns it to the Studies field.

### GetLabtests

`func (o *Ga4ghComponent) GetLabtests() Ga4ghSearchLabtestsRequest`

GetLabtests returns the Labtests field if non-nil, zero value otherwise.

### GetLabtestsOk

`func (o *Ga4ghComponent) GetLabtestsOk() (Ga4ghSearchLabtestsRequest, bool)`

GetLabtestsOk returns a tuple with the Labtests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabtests

`func (o *Ga4ghComponent) HasLabtests() bool`

HasLabtests returns a boolean if a field has been set.

### SetLabtests

`func (o *Ga4ghComponent) SetLabtests(v Ga4ghSearchLabtestsRequest)`

SetLabtests gets a reference to the given Ga4ghSearchLabtestsRequest and assigns it to the Labtests field.

### GetChemotherapies

`func (o *Ga4ghComponent) GetChemotherapies() Ga4ghSearchChemotherapiesRequest`

GetChemotherapies returns the Chemotherapies field if non-nil, zero value otherwise.

### GetChemotherapiesOk

`func (o *Ga4ghComponent) GetChemotherapiesOk() (Ga4ghSearchChemotherapiesRequest, bool)`

GetChemotherapiesOk returns a tuple with the Chemotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChemotherapies

`func (o *Ga4ghComponent) HasChemotherapies() bool`

HasChemotherapies returns a boolean if a field has been set.

### SetChemotherapies

`func (o *Ga4ghComponent) SetChemotherapies(v Ga4ghSearchChemotherapiesRequest)`

SetChemotherapies gets a reference to the given Ga4ghSearchChemotherapiesRequest and assigns it to the Chemotherapies field.

### GetRadiotherapies

`func (o *Ga4ghComponent) GetRadiotherapies() Ga4ghSearchRadiotherapiesRequest`

GetRadiotherapies returns the Radiotherapies field if non-nil, zero value otherwise.

### GetRadiotherapiesOk

`func (o *Ga4ghComponent) GetRadiotherapiesOk() (Ga4ghSearchRadiotherapiesRequest, bool)`

GetRadiotherapiesOk returns a tuple with the Radiotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRadiotherapies

`func (o *Ga4ghComponent) HasRadiotherapies() bool`

HasRadiotherapies returns a boolean if a field has been set.

### SetRadiotherapies

`func (o *Ga4ghComponent) SetRadiotherapies(v Ga4ghSearchRadiotherapiesRequest)`

SetRadiotherapies gets a reference to the given Ga4ghSearchRadiotherapiesRequest and assigns it to the Radiotherapies field.

### GetImmunotherapies

`func (o *Ga4ghComponent) GetImmunotherapies() Ga4ghSearchImmunotherapiesRequest`

GetImmunotherapies returns the Immunotherapies field if non-nil, zero value otherwise.

### GetImmunotherapiesOk

`func (o *Ga4ghComponent) GetImmunotherapiesOk() (Ga4ghSearchImmunotherapiesRequest, bool)`

GetImmunotherapiesOk returns a tuple with the Immunotherapies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImmunotherapies

`func (o *Ga4ghComponent) HasImmunotherapies() bool`

HasImmunotherapies returns a boolean if a field has been set.

### SetImmunotherapies

`func (o *Ga4ghComponent) SetImmunotherapies(v Ga4ghSearchImmunotherapiesRequest)`

SetImmunotherapies gets a reference to the given Ga4ghSearchImmunotherapiesRequest and assigns it to the Immunotherapies field.

### GetSurgeries

`func (o *Ga4ghComponent) GetSurgeries() Ga4ghSearchSurgeriesRequest`

GetSurgeries returns the Surgeries field if non-nil, zero value otherwise.

### GetSurgeriesOk

`func (o *Ga4ghComponent) GetSurgeriesOk() (Ga4ghSearchSurgeriesRequest, bool)`

GetSurgeriesOk returns a tuple with the Surgeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurgeries

`func (o *Ga4ghComponent) HasSurgeries() bool`

HasSurgeries returns a boolean if a field has been set.

### SetSurgeries

`func (o *Ga4ghComponent) SetSurgeries(v Ga4ghSearchSurgeriesRequest)`

SetSurgeries gets a reference to the given Ga4ghSearchSurgeriesRequest and assigns it to the Surgeries field.

### GetCelltransplants

`func (o *Ga4ghComponent) GetCelltransplants() Ga4ghSearchCelltransplantsRequest`

GetCelltransplants returns the Celltransplants field if non-nil, zero value otherwise.

### GetCelltransplantsOk

`func (o *Ga4ghComponent) GetCelltransplantsOk() (Ga4ghSearchCelltransplantsRequest, bool)`

GetCelltransplantsOk returns a tuple with the Celltransplants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCelltransplants

`func (o *Ga4ghComponent) HasCelltransplants() bool`

HasCelltransplants returns a boolean if a field has been set.

### SetCelltransplants

`func (o *Ga4ghComponent) SetCelltransplants(v Ga4ghSearchCelltransplantsRequest)`

SetCelltransplants gets a reference to the given Ga4ghSearchCelltransplantsRequest and assigns it to the Celltransplants field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


