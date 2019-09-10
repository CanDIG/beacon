# Ga4ghEnrollment

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
**EnrollmentInstitution** | Pointer to **string** |  | [optional] 
**EnrollmentApprovalDate** | Pointer to **string** |  | [optional] 
**CrossEnrollment** | Pointer to **string** |  | [optional] 
**OtherPersonalizedMedicineStudyName** | Pointer to **string** |  | [optional] 
**OtherPersonalizedMedicineStudyId** | Pointer to **string** |  | [optional] 
**AgeAtEnrollment** | Pointer to **string** |  | [optional] 
**EligibilityCategory** | Pointer to **string** |  | [optional] 
**StatusAtEnrollment** | Pointer to **string** |  | [optional] 
**PrimaryOncologistName** | Pointer to **string** |  | [optional] 
**PrimaryOncologistContact** | Pointer to **string** |  | [optional] 
**ReferringPhysicianName** | Pointer to **string** |  | [optional] 
**ReferringPhysicianContact** | Pointer to **string** |  | [optional] 
**SummaryOfIdRequest** | Pointer to **string** |  | [optional] 
**TreatingCentreName** | Pointer to **string** |  | [optional] 
**TreatingCentreProvince** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghEnrollment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghEnrollment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghEnrollment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghEnrollment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghEnrollment) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghEnrollment) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghEnrollment) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghEnrollment) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghEnrollment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghEnrollment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghEnrollment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghEnrollment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghEnrollment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghEnrollment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghEnrollment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghEnrollment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghEnrollment) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghEnrollment) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghEnrollment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghEnrollment) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghEnrollment) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghEnrollment) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghEnrollment) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghEnrollment) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghEnrollment) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghEnrollment) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghEnrollment) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghEnrollment) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetEnrollmentInstitution

`func (o *Ga4ghEnrollment) GetEnrollmentInstitution() string`

GetEnrollmentInstitution returns the EnrollmentInstitution field if non-nil, zero value otherwise.

### GetEnrollmentInstitutionOk

`func (o *Ga4ghEnrollment) GetEnrollmentInstitutionOk() (string, bool)`

GetEnrollmentInstitutionOk returns a tuple with the EnrollmentInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollmentInstitution

`func (o *Ga4ghEnrollment) HasEnrollmentInstitution() bool`

HasEnrollmentInstitution returns a boolean if a field has been set.

### SetEnrollmentInstitution

`func (o *Ga4ghEnrollment) SetEnrollmentInstitution(v string)`

SetEnrollmentInstitution gets a reference to the given string and assigns it to the EnrollmentInstitution field.

### GetEnrollmentApprovalDate

`func (o *Ga4ghEnrollment) GetEnrollmentApprovalDate() string`

GetEnrollmentApprovalDate returns the EnrollmentApprovalDate field if non-nil, zero value otherwise.

### GetEnrollmentApprovalDateOk

`func (o *Ga4ghEnrollment) GetEnrollmentApprovalDateOk() (string, bool)`

GetEnrollmentApprovalDateOk returns a tuple with the EnrollmentApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollmentApprovalDate

`func (o *Ga4ghEnrollment) HasEnrollmentApprovalDate() bool`

HasEnrollmentApprovalDate returns a boolean if a field has been set.

### SetEnrollmentApprovalDate

`func (o *Ga4ghEnrollment) SetEnrollmentApprovalDate(v string)`

SetEnrollmentApprovalDate gets a reference to the given string and assigns it to the EnrollmentApprovalDate field.

### GetCrossEnrollment

`func (o *Ga4ghEnrollment) GetCrossEnrollment() string`

GetCrossEnrollment returns the CrossEnrollment field if non-nil, zero value otherwise.

### GetCrossEnrollmentOk

`func (o *Ga4ghEnrollment) GetCrossEnrollmentOk() (string, bool)`

GetCrossEnrollmentOk returns a tuple with the CrossEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCrossEnrollment

`func (o *Ga4ghEnrollment) HasCrossEnrollment() bool`

HasCrossEnrollment returns a boolean if a field has been set.

### SetCrossEnrollment

`func (o *Ga4ghEnrollment) SetCrossEnrollment(v string)`

SetCrossEnrollment gets a reference to the given string and assigns it to the CrossEnrollment field.

### GetOtherPersonalizedMedicineStudyName

`func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyName() string`

GetOtherPersonalizedMedicineStudyName returns the OtherPersonalizedMedicineStudyName field if non-nil, zero value otherwise.

### GetOtherPersonalizedMedicineStudyNameOk

`func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyNameOk() (string, bool)`

GetOtherPersonalizedMedicineStudyNameOk returns a tuple with the OtherPersonalizedMedicineStudyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherPersonalizedMedicineStudyName

`func (o *Ga4ghEnrollment) HasOtherPersonalizedMedicineStudyName() bool`

HasOtherPersonalizedMedicineStudyName returns a boolean if a field has been set.

### SetOtherPersonalizedMedicineStudyName

`func (o *Ga4ghEnrollment) SetOtherPersonalizedMedicineStudyName(v string)`

SetOtherPersonalizedMedicineStudyName gets a reference to the given string and assigns it to the OtherPersonalizedMedicineStudyName field.

### GetOtherPersonalizedMedicineStudyId

`func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyId() string`

GetOtherPersonalizedMedicineStudyId returns the OtherPersonalizedMedicineStudyId field if non-nil, zero value otherwise.

### GetOtherPersonalizedMedicineStudyIdOk

`func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyIdOk() (string, bool)`

GetOtherPersonalizedMedicineStudyIdOk returns a tuple with the OtherPersonalizedMedicineStudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherPersonalizedMedicineStudyId

`func (o *Ga4ghEnrollment) HasOtherPersonalizedMedicineStudyId() bool`

HasOtherPersonalizedMedicineStudyId returns a boolean if a field has been set.

### SetOtherPersonalizedMedicineStudyId

`func (o *Ga4ghEnrollment) SetOtherPersonalizedMedicineStudyId(v string)`

SetOtherPersonalizedMedicineStudyId gets a reference to the given string and assigns it to the OtherPersonalizedMedicineStudyId field.

### GetAgeAtEnrollment

`func (o *Ga4ghEnrollment) GetAgeAtEnrollment() string`

GetAgeAtEnrollment returns the AgeAtEnrollment field if non-nil, zero value otherwise.

### GetAgeAtEnrollmentOk

`func (o *Ga4ghEnrollment) GetAgeAtEnrollmentOk() (string, bool)`

GetAgeAtEnrollmentOk returns a tuple with the AgeAtEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgeAtEnrollment

`func (o *Ga4ghEnrollment) HasAgeAtEnrollment() bool`

HasAgeAtEnrollment returns a boolean if a field has been set.

### SetAgeAtEnrollment

`func (o *Ga4ghEnrollment) SetAgeAtEnrollment(v string)`

SetAgeAtEnrollment gets a reference to the given string and assigns it to the AgeAtEnrollment field.

### GetEligibilityCategory

`func (o *Ga4ghEnrollment) GetEligibilityCategory() string`

GetEligibilityCategory returns the EligibilityCategory field if non-nil, zero value otherwise.

### GetEligibilityCategoryOk

`func (o *Ga4ghEnrollment) GetEligibilityCategoryOk() (string, bool)`

GetEligibilityCategoryOk returns a tuple with the EligibilityCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEligibilityCategory

`func (o *Ga4ghEnrollment) HasEligibilityCategory() bool`

HasEligibilityCategory returns a boolean if a field has been set.

### SetEligibilityCategory

`func (o *Ga4ghEnrollment) SetEligibilityCategory(v string)`

SetEligibilityCategory gets a reference to the given string and assigns it to the EligibilityCategory field.

### GetStatusAtEnrollment

`func (o *Ga4ghEnrollment) GetStatusAtEnrollment() string`

GetStatusAtEnrollment returns the StatusAtEnrollment field if non-nil, zero value otherwise.

### GetStatusAtEnrollmentOk

`func (o *Ga4ghEnrollment) GetStatusAtEnrollmentOk() (string, bool)`

GetStatusAtEnrollmentOk returns a tuple with the StatusAtEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatusAtEnrollment

`func (o *Ga4ghEnrollment) HasStatusAtEnrollment() bool`

HasStatusAtEnrollment returns a boolean if a field has been set.

### SetStatusAtEnrollment

`func (o *Ga4ghEnrollment) SetStatusAtEnrollment(v string)`

SetStatusAtEnrollment gets a reference to the given string and assigns it to the StatusAtEnrollment field.

### GetPrimaryOncologistName

`func (o *Ga4ghEnrollment) GetPrimaryOncologistName() string`

GetPrimaryOncologistName returns the PrimaryOncologistName field if non-nil, zero value otherwise.

### GetPrimaryOncologistNameOk

`func (o *Ga4ghEnrollment) GetPrimaryOncologistNameOk() (string, bool)`

GetPrimaryOncologistNameOk returns a tuple with the PrimaryOncologistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimaryOncologistName

`func (o *Ga4ghEnrollment) HasPrimaryOncologistName() bool`

HasPrimaryOncologistName returns a boolean if a field has been set.

### SetPrimaryOncologistName

`func (o *Ga4ghEnrollment) SetPrimaryOncologistName(v string)`

SetPrimaryOncologistName gets a reference to the given string and assigns it to the PrimaryOncologistName field.

### GetPrimaryOncologistContact

`func (o *Ga4ghEnrollment) GetPrimaryOncologistContact() string`

GetPrimaryOncologistContact returns the PrimaryOncologistContact field if non-nil, zero value otherwise.

### GetPrimaryOncologistContactOk

`func (o *Ga4ghEnrollment) GetPrimaryOncologistContactOk() (string, bool)`

GetPrimaryOncologistContactOk returns a tuple with the PrimaryOncologistContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimaryOncologistContact

`func (o *Ga4ghEnrollment) HasPrimaryOncologistContact() bool`

HasPrimaryOncologistContact returns a boolean if a field has been set.

### SetPrimaryOncologistContact

`func (o *Ga4ghEnrollment) SetPrimaryOncologistContact(v string)`

SetPrimaryOncologistContact gets a reference to the given string and assigns it to the PrimaryOncologistContact field.

### GetReferringPhysicianName

`func (o *Ga4ghEnrollment) GetReferringPhysicianName() string`

GetReferringPhysicianName returns the ReferringPhysicianName field if non-nil, zero value otherwise.

### GetReferringPhysicianNameOk

`func (o *Ga4ghEnrollment) GetReferringPhysicianNameOk() (string, bool)`

GetReferringPhysicianNameOk returns a tuple with the ReferringPhysicianName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferringPhysicianName

`func (o *Ga4ghEnrollment) HasReferringPhysicianName() bool`

HasReferringPhysicianName returns a boolean if a field has been set.

### SetReferringPhysicianName

`func (o *Ga4ghEnrollment) SetReferringPhysicianName(v string)`

SetReferringPhysicianName gets a reference to the given string and assigns it to the ReferringPhysicianName field.

### GetReferringPhysicianContact

`func (o *Ga4ghEnrollment) GetReferringPhysicianContact() string`

GetReferringPhysicianContact returns the ReferringPhysicianContact field if non-nil, zero value otherwise.

### GetReferringPhysicianContactOk

`func (o *Ga4ghEnrollment) GetReferringPhysicianContactOk() (string, bool)`

GetReferringPhysicianContactOk returns a tuple with the ReferringPhysicianContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferringPhysicianContact

`func (o *Ga4ghEnrollment) HasReferringPhysicianContact() bool`

HasReferringPhysicianContact returns a boolean if a field has been set.

### SetReferringPhysicianContact

`func (o *Ga4ghEnrollment) SetReferringPhysicianContact(v string)`

SetReferringPhysicianContact gets a reference to the given string and assigns it to the ReferringPhysicianContact field.

### GetSummaryOfIdRequest

`func (o *Ga4ghEnrollment) GetSummaryOfIdRequest() string`

GetSummaryOfIdRequest returns the SummaryOfIdRequest field if non-nil, zero value otherwise.

### GetSummaryOfIdRequestOk

`func (o *Ga4ghEnrollment) GetSummaryOfIdRequestOk() (string, bool)`

GetSummaryOfIdRequestOk returns a tuple with the SummaryOfIdRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSummaryOfIdRequest

`func (o *Ga4ghEnrollment) HasSummaryOfIdRequest() bool`

HasSummaryOfIdRequest returns a boolean if a field has been set.

### SetSummaryOfIdRequest

`func (o *Ga4ghEnrollment) SetSummaryOfIdRequest(v string)`

SetSummaryOfIdRequest gets a reference to the given string and assigns it to the SummaryOfIdRequest field.

### GetTreatingCentreName

`func (o *Ga4ghEnrollment) GetTreatingCentreName() string`

GetTreatingCentreName returns the TreatingCentreName field if non-nil, zero value otherwise.

### GetTreatingCentreNameOk

`func (o *Ga4ghEnrollment) GetTreatingCentreNameOk() (string, bool)`

GetTreatingCentreNameOk returns a tuple with the TreatingCentreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatingCentreName

`func (o *Ga4ghEnrollment) HasTreatingCentreName() bool`

HasTreatingCentreName returns a boolean if a field has been set.

### SetTreatingCentreName

`func (o *Ga4ghEnrollment) SetTreatingCentreName(v string)`

SetTreatingCentreName gets a reference to the given string and assigns it to the TreatingCentreName field.

### GetTreatingCentreProvince

`func (o *Ga4ghEnrollment) GetTreatingCentreProvince() string`

GetTreatingCentreProvince returns the TreatingCentreProvince field if non-nil, zero value otherwise.

### GetTreatingCentreProvinceOk

`func (o *Ga4ghEnrollment) GetTreatingCentreProvinceOk() (string, bool)`

GetTreatingCentreProvinceOk returns a tuple with the TreatingCentreProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTreatingCentreProvince

`func (o *Ga4ghEnrollment) HasTreatingCentreProvince() bool`

HasTreatingCentreProvince returns a boolean if a field has been set.

### SetTreatingCentreProvince

`func (o *Ga4ghEnrollment) SetTreatingCentreProvince(v string)`

SetTreatingCentreProvince gets a reference to the given string and assigns it to the TreatingCentreProvince field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


