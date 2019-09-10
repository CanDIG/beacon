# \ClinicalMetadataServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCelltransplant**](ClinicalMetadataServiceApi.md#GetCelltransplant) | **Get** /celltransplants/{celltransplant_id} |  Gets &#x60;Celltransplant&#x60; by ID. 
[**GetChemotherapy**](ClinicalMetadataServiceApi.md#GetChemotherapy) | **Get** /chemotherapies/{chemotherapy_id} |  Gets &#x60;Chemotherapy&#x60; by ID. 
[**GetComplication**](ClinicalMetadataServiceApi.md#GetComplication) | **Get** /complications/{complication_id} |  Gets &#x60;Complication&#x60; by ID. 
[**GetConsent**](ClinicalMetadataServiceApi.md#GetConsent) | **Get** /consents/{consent_id} |  Gets &#x60;Consent&#x60; by ID. 
[**GetDiagnosis**](ClinicalMetadataServiceApi.md#GetDiagnosis) | **Get** /diagnoses/{diagnosis_id} |  Gets &#x60;Diagnosis&#x60; by ID. 
[**GetEnrollment**](ClinicalMetadataServiceApi.md#GetEnrollment) | **Get** /enrollments/{enrollment_id} |  Gets &#x60;Enrollment&#x60; by ID. 
[**GetImmunotherapy**](ClinicalMetadataServiceApi.md#GetImmunotherapy) | **Get** /immunotherapies/{immunotherapy_id} |  Gets &#x60;Immunotherapy&#x60; by ID. 
[**GetLabtest**](ClinicalMetadataServiceApi.md#GetLabtest) | **Get** /labtests/{labtest_id} |  Gets &#x60;Labtest&#x60; by ID. 
[**GetOutcome**](ClinicalMetadataServiceApi.md#GetOutcome) | **Get** /outcomes/{outcome_id} |  Gets &#x60;Outcome&#x60; by ID. 
[**GetPatient**](ClinicalMetadataServiceApi.md#GetPatient) | **Get** /patients/{patient_id} |  Gets &#x60;Patient&#x60; by ID. 
[**GetRadiotherapy**](ClinicalMetadataServiceApi.md#GetRadiotherapy) | **Get** /radiotherapies/{radiotherapy_id} |  Gets &#x60;Radiotherapy&#x60; by ID. 
[**GetSample**](ClinicalMetadataServiceApi.md#GetSample) | **Get** /samples/{sample_id} |  Gets &#x60;Sample&#x60; by ID. 
[**GetSlide**](ClinicalMetadataServiceApi.md#GetSlide) | **Get** /slides/{slide_id} |  Gets &#x60;Slide&#x60; by ID. 
[**GetStudy**](ClinicalMetadataServiceApi.md#GetStudy) | **Get** /studies/{study_id} |  Gets &#x60;Study&#x60; by ID. 
[**GetSurgery**](ClinicalMetadataServiceApi.md#GetSurgery) | **Get** /surgeries/{surgery_id} |  Gets &#x60;Surgery&#x60; by ID. 
[**GetTreatment**](ClinicalMetadataServiceApi.md#GetTreatment) | **Get** /treatments/{treatment_id} |  Gets &#x60;Treatment&#x60; by ID. 
[**GetTumourboard**](ClinicalMetadataServiceApi.md#GetTumourboard) | **Get** /tumourboards/{tumourboard_id} |  Gets &#x60;Tumourboard&#x60; by ID. 
[**SearchCelltransplants**](ClinicalMetadataServiceApi.md#SearchCelltransplants) | **Post** /celltransplants/search |  Gets a list of Celltransplants accessible through the API. 
[**SearchChemotherapies**](ClinicalMetadataServiceApi.md#SearchChemotherapies) | **Post** /chemotherapies/search |  Gets a list of Chemotherapies accessible through the API. 
[**SearchComplications**](ClinicalMetadataServiceApi.md#SearchComplications) | **Post** /complications/search |  Gets a list of Complications accessible through the API. 
[**SearchConsents**](ClinicalMetadataServiceApi.md#SearchConsents) | **Post** /consents/search |  Gets a list of Consents accessible through the API. 
[**SearchDiagnoses**](ClinicalMetadataServiceApi.md#SearchDiagnoses) | **Post** /diagnoses/search |  Gets a list of Diagnoses accessible through the API. 
[**SearchEnrollments**](ClinicalMetadataServiceApi.md#SearchEnrollments) | **Post** /enrollments/search |  Gets a list of Enrollments accessible through the API. 
[**SearchImmunotherapies**](ClinicalMetadataServiceApi.md#SearchImmunotherapies) | **Post** /immunotherapies/search |  Gets a list of Immunotherapies accessible through the API. 
[**SearchLabtests**](ClinicalMetadataServiceApi.md#SearchLabtests) | **Post** /labtests/search |  Gets a list of Labtests accessible through the API. 
[**SearchOutcomes**](ClinicalMetadataServiceApi.md#SearchOutcomes) | **Post** /outcomes/search |  Gets a list of Outcomes accessible through the API. 
[**SearchPatients**](ClinicalMetadataServiceApi.md#SearchPatients) | **Post** /patients/search |  Gets a list of Patients accessible through the API. 
[**SearchRadiotherapies**](ClinicalMetadataServiceApi.md#SearchRadiotherapies) | **Post** /radiotherapies/search |  Gets a list of Radiotherapies accessible through the API. 
[**SearchSamples**](ClinicalMetadataServiceApi.md#SearchSamples) | **Post** /samples/search |  Gets a list of Samples accessible through the API. 
[**SearchSlides**](ClinicalMetadataServiceApi.md#SearchSlides) | **Post** /slides/search |  Gets a list of Slides accessible through the API. 
[**SearchStudies**](ClinicalMetadataServiceApi.md#SearchStudies) | **Post** /studies/search |  Gets a list of Studies accessible through the API. 
[**SearchSurgeries**](ClinicalMetadataServiceApi.md#SearchSurgeries) | **Post** /surgeries/search |  Gets a list of Surgeries accessible through the API. 
[**SearchTreatments**](ClinicalMetadataServiceApi.md#SearchTreatments) | **Post** /treatments/search |  Gets a list of Treatments accessible through the API. 
[**SearchTumourboards**](ClinicalMetadataServiceApi.md#SearchTumourboards) | **Post** /tumourboards/search |  Gets a list of Tumourboards accessible through the API. 



## GetCelltransplant

> Ga4ghCelltransplant GetCelltransplant(ctx, celltransplantId)
 Gets `Celltransplant` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**celltransplantId** | **string**| The ID of the celltransplant requested | 

### Return type

[**Ga4ghCelltransplant**](ga4ghCelltransplant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChemotherapy

> Ga4ghChemotherapy GetChemotherapy(ctx, chemotherapyId)
 Gets `Chemotherapy` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chemotherapyId** | **string**| The ID of the chemotherapy requested | 

### Return type

[**Ga4ghChemotherapy**](ga4ghChemotherapy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplication

> Ga4ghComplication GetComplication(ctx, complicationId)
 Gets `Complication` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complicationId** | **string**| The ID of the complication requested | 

### Return type

[**Ga4ghComplication**](ga4ghComplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsent

> Ga4ghConsent GetConsent(ctx, consentId)
 Gets `Consent` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentId** | **string**| The ID of the consent requested | 

### Return type

[**Ga4ghConsent**](ga4ghConsent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiagnosis

> Ga4ghDiagnosis GetDiagnosis(ctx, diagnosisId)
 Gets `Diagnosis` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diagnosisId** | **string**| The ID of the diagnosis requested | 

### Return type

[**Ga4ghDiagnosis**](ga4ghDiagnosis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollment

> Ga4ghEnrollment GetEnrollment(ctx, enrollmentId)
 Gets `Enrollment` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enrollmentId** | **string**| The ID of the enrollment requested | 

### Return type

[**Ga4ghEnrollment**](ga4ghEnrollment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImmunotherapy

> Ga4ghImmunotherapy GetImmunotherapy(ctx, immunotherapyId)
 Gets `Immunotherapy` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**immunotherapyId** | **string**| The ID of the immunotherapy requested | 

### Return type

[**Ga4ghImmunotherapy**](ga4ghImmunotherapy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabtest

> Ga4ghLabtest GetLabtest(ctx, labtestId)
 Gets `Labtest` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labtestId** | **string**| The ID of the labtest requested | 

### Return type

[**Ga4ghLabtest**](ga4ghLabtest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutcome

> Ga4ghOutcome GetOutcome(ctx, outcomeId)
 Gets `Outcome` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcomeId** | **string**| The ID of the outcome requested | 

### Return type

[**Ga4ghOutcome**](ga4ghOutcome.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatient

> Ga4ghPatient GetPatient(ctx, patientId)
 Gets `Patient` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patientId** | **string**| The ID of the patient requested | 

### Return type

[**Ga4ghPatient**](ga4ghPatient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadiotherapy

> Ga4ghRadiotherapy GetRadiotherapy(ctx, radiotherapyId)
 Gets `Radiotherapy` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radiotherapyId** | **string**| The ID of the radiotherapy requested | 

### Return type

[**Ga4ghRadiotherapy**](ga4ghRadiotherapy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSample

> Ga4ghSample GetSample(ctx, sampleId)
 Gets `Sample` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sampleId** | **string**| The ID of the sample requested | 

### Return type

[**Ga4ghSample**](ga4ghSample.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlide

> Ga4ghSlide GetSlide(ctx, slideId)
 Gets `Slide` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slideId** | **string**| The ID of the slide requested | 

### Return type

[**Ga4ghSlide**](ga4ghSlide.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudy

> Ga4ghStudy GetStudy(ctx, studyId)
 Gets `Study` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string**| The ID of the study requested | 

### Return type

[**Ga4ghStudy**](ga4ghStudy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSurgery

> Ga4ghSurgery GetSurgery(ctx, surgeryId)
 Gets `Surgery` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surgeryId** | **string**| The ID of the surgery requested | 

### Return type

[**Ga4ghSurgery**](ga4ghSurgery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTreatment

> Ga4ghTreatment GetTreatment(ctx, treatmentId)
 Gets `Treatment` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**treatmentId** | **string**| The ID of the treatment requested | 

### Return type

[**Ga4ghTreatment**](ga4ghTreatment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTumourboard

> Ga4ghTumourboard GetTumourboard(ctx, tumourboardId)
 Gets `Tumourboard` by ID. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tumourboardId** | **string**| The ID of the tumourboard requested | 

### Return type

[**Ga4ghTumourboard**](ga4ghTumourboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCelltransplants

> Ga4ghSearchCelltransplantsResponse SearchCelltransplants(ctx, body)
 Gets a list of Celltransplants accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchCelltransplantsRequest**](Ga4ghSearchCelltransplantsRequest.md)|  | 

### Return type

[**Ga4ghSearchCelltransplantsResponse**](ga4ghSearchCelltransplantsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchChemotherapies

> Ga4ghSearchChemotherapiesResponse SearchChemotherapies(ctx, body)
 Gets a list of Chemotherapies accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchChemotherapiesRequest**](Ga4ghSearchChemotherapiesRequest.md)|  | 

### Return type

[**Ga4ghSearchChemotherapiesResponse**](ga4ghSearchChemotherapiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchComplications

> Ga4ghSearchComplicationsResponse SearchComplications(ctx, body)
 Gets a list of Complications accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchComplicationsRequest**](Ga4ghSearchComplicationsRequest.md)|  | 

### Return type

[**Ga4ghSearchComplicationsResponse**](ga4ghSearchComplicationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchConsents

> Ga4ghSearchConsentsResponse SearchConsents(ctx, body)
 Gets a list of Consents accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchConsentsRequest**](Ga4ghSearchConsentsRequest.md)|  | 

### Return type

[**Ga4ghSearchConsentsResponse**](ga4ghSearchConsentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDiagnoses

> Ga4ghSearchDiagnosesResponse SearchDiagnoses(ctx, body)
 Gets a list of Diagnoses accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchDiagnosesRequest**](Ga4ghSearchDiagnosesRequest.md)|  | 

### Return type

[**Ga4ghSearchDiagnosesResponse**](ga4ghSearchDiagnosesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEnrollments

> Ga4ghSearchEnrollmentsResponse SearchEnrollments(ctx, body)
 Gets a list of Enrollments accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchEnrollmentsRequest**](Ga4ghSearchEnrollmentsRequest.md)|  | 

### Return type

[**Ga4ghSearchEnrollmentsResponse**](ga4ghSearchEnrollmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchImmunotherapies

> Ga4ghSearchImmunotherapiesResponse SearchImmunotherapies(ctx, body)
 Gets a list of Immunotherapies accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchImmunotherapiesRequest**](Ga4ghSearchImmunotherapiesRequest.md)|  | 

### Return type

[**Ga4ghSearchImmunotherapiesResponse**](ga4ghSearchImmunotherapiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLabtests

> Ga4ghSearchLabtestsResponse SearchLabtests(ctx, body)
 Gets a list of Labtests accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchLabtestsRequest**](Ga4ghSearchLabtestsRequest.md)|  | 

### Return type

[**Ga4ghSearchLabtestsResponse**](ga4ghSearchLabtestsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOutcomes

> Ga4ghSearchOutcomesResponse SearchOutcomes(ctx, body)
 Gets a list of Outcomes accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchOutcomesRequest**](Ga4ghSearchOutcomesRequest.md)|  | 

### Return type

[**Ga4ghSearchOutcomesResponse**](ga4ghSearchOutcomesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPatients

> Ga4ghSearchPatientsResponse SearchPatients(ctx, body)
 Gets a list of Patients accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchPatientsRequest**](Ga4ghSearchPatientsRequest.md)|  | 

### Return type

[**Ga4ghSearchPatientsResponse**](ga4ghSearchPatientsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRadiotherapies

> Ga4ghSearchRadiotherapiesResponse SearchRadiotherapies(ctx, body)
 Gets a list of Radiotherapies accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchRadiotherapiesRequest**](Ga4ghSearchRadiotherapiesRequest.md)|  | 

### Return type

[**Ga4ghSearchRadiotherapiesResponse**](ga4ghSearchRadiotherapiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSamples

> Ga4ghSearchSamplesResponse SearchSamples(ctx, body)
 Gets a list of Samples accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchSamplesRequest**](Ga4ghSearchSamplesRequest.md)|  | 

### Return type

[**Ga4ghSearchSamplesResponse**](ga4ghSearchSamplesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSlides

> Ga4ghSearchSlidesResponse SearchSlides(ctx, body)
 Gets a list of Slides accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchSlidesRequest**](Ga4ghSearchSlidesRequest.md)|  | 

### Return type

[**Ga4ghSearchSlidesResponse**](ga4ghSearchSlidesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStudies

> Ga4ghSearchStudiesResponse SearchStudies(ctx, body)
 Gets a list of Studies accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchStudiesRequest**](Ga4ghSearchStudiesRequest.md)|  | 

### Return type

[**Ga4ghSearchStudiesResponse**](ga4ghSearchStudiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSurgeries

> Ga4ghSearchSurgeriesResponse SearchSurgeries(ctx, body)
 Gets a list of Surgeries accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchSurgeriesRequest**](Ga4ghSearchSurgeriesRequest.md)|  | 

### Return type

[**Ga4ghSearchSurgeriesResponse**](ga4ghSearchSurgeriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTreatments

> Ga4ghSearchTreatmentsResponse SearchTreatments(ctx, body)
 Gets a list of Treatments accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchTreatmentsRequest**](Ga4ghSearchTreatmentsRequest.md)|  | 

### Return type

[**Ga4ghSearchTreatmentsResponse**](ga4ghSearchTreatmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTumourboards

> Ga4ghSearchTumourboardsResponse SearchTumourboards(ctx, body)
 Gets a list of Tumourboards accessible through the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Ga4ghSearchTumourboardsRequest**](Ga4ghSearchTumourboardsRequest.md)|  | 

### Return type

[**Ga4ghSearchTumourboardsResponse**](ga4ghSearchTumourboardsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

