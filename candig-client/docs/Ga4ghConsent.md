# Ga4ghConsent

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
**ConsentId** | Pointer to **string** |  | [optional] 
**ConsentDate** | Pointer to **string** |  | [optional] 
**ConsentVersion** | Pointer to **string** |  | [optional] 
**PatientConsentedTo** | Pointer to **string** |  | [optional] 
**ReasonForRejection** | Pointer to **string** |  | [optional] 
**WasAssentObtained** | Pointer to **string** |  | [optional] 
**DateOfAssent** | Pointer to **string** |  | [optional] 
**AssentFormVersion** | Pointer to **string** |  | [optional] 
**IfAssentNotObtainedWhyNot** | Pointer to **string** |  | [optional] 
**ReconsentDate** | Pointer to **string** |  | [optional] 
**ReconsentVersion** | Pointer to **string** |  | [optional] 
**ConsentingCoordinatorName** | Pointer to **string** |  | [optional] 
**PreviouslyConsented** | Pointer to **string** |  | [optional] 
**NameOfOtherBiobank** | Pointer to **string** |  | [optional] 
**HasConsentBeenWithdrawn** | Pointer to **string** |  | [optional] 
**DateOfConsentWithdrawal** | Pointer to **string** |  | [optional] 
**TypeOfConsentWithdrawal** | Pointer to **string** |  | [optional] 
**ReasonForConsentWithdrawal** | Pointer to **string** |  | [optional] 
**ConsentFormComplete** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghConsent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghConsent) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghConsent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghConsent) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghConsent) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghConsent) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghConsent) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghConsent) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghConsent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghConsent) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghConsent) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghConsent) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghConsent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghConsent) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghConsent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghConsent) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghConsent) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghConsent) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghConsent) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghConsent) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghConsent) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghConsent) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghConsent) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghConsent) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghConsent) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghConsent) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghConsent) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghConsent) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetConsentId

`func (o *Ga4ghConsent) GetConsentId() string`

GetConsentId returns the ConsentId field if non-nil, zero value otherwise.

### GetConsentIdOk

`func (o *Ga4ghConsent) GetConsentIdOk() (string, bool)`

GetConsentIdOk returns a tuple with the ConsentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentId

`func (o *Ga4ghConsent) HasConsentId() bool`

HasConsentId returns a boolean if a field has been set.

### SetConsentId

`func (o *Ga4ghConsent) SetConsentId(v string)`

SetConsentId gets a reference to the given string and assigns it to the ConsentId field.

### GetConsentDate

`func (o *Ga4ghConsent) GetConsentDate() string`

GetConsentDate returns the ConsentDate field if non-nil, zero value otherwise.

### GetConsentDateOk

`func (o *Ga4ghConsent) GetConsentDateOk() (string, bool)`

GetConsentDateOk returns a tuple with the ConsentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentDate

`func (o *Ga4ghConsent) HasConsentDate() bool`

HasConsentDate returns a boolean if a field has been set.

### SetConsentDate

`func (o *Ga4ghConsent) SetConsentDate(v string)`

SetConsentDate gets a reference to the given string and assigns it to the ConsentDate field.

### GetConsentVersion

`func (o *Ga4ghConsent) GetConsentVersion() string`

GetConsentVersion returns the ConsentVersion field if non-nil, zero value otherwise.

### GetConsentVersionOk

`func (o *Ga4ghConsent) GetConsentVersionOk() (string, bool)`

GetConsentVersionOk returns a tuple with the ConsentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentVersion

`func (o *Ga4ghConsent) HasConsentVersion() bool`

HasConsentVersion returns a boolean if a field has been set.

### SetConsentVersion

`func (o *Ga4ghConsent) SetConsentVersion(v string)`

SetConsentVersion gets a reference to the given string and assigns it to the ConsentVersion field.

### GetPatientConsentedTo

`func (o *Ga4ghConsent) GetPatientConsentedTo() string`

GetPatientConsentedTo returns the PatientConsentedTo field if non-nil, zero value otherwise.

### GetPatientConsentedToOk

`func (o *Ga4ghConsent) GetPatientConsentedToOk() (string, bool)`

GetPatientConsentedToOk returns a tuple with the PatientConsentedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientConsentedTo

`func (o *Ga4ghConsent) HasPatientConsentedTo() bool`

HasPatientConsentedTo returns a boolean if a field has been set.

### SetPatientConsentedTo

`func (o *Ga4ghConsent) SetPatientConsentedTo(v string)`

SetPatientConsentedTo gets a reference to the given string and assigns it to the PatientConsentedTo field.

### GetReasonForRejection

`func (o *Ga4ghConsent) GetReasonForRejection() string`

GetReasonForRejection returns the ReasonForRejection field if non-nil, zero value otherwise.

### GetReasonForRejectionOk

`func (o *Ga4ghConsent) GetReasonForRejectionOk() (string, bool)`

GetReasonForRejectionOk returns a tuple with the ReasonForRejection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReasonForRejection

`func (o *Ga4ghConsent) HasReasonForRejection() bool`

HasReasonForRejection returns a boolean if a field has been set.

### SetReasonForRejection

`func (o *Ga4ghConsent) SetReasonForRejection(v string)`

SetReasonForRejection gets a reference to the given string and assigns it to the ReasonForRejection field.

### GetWasAssentObtained

`func (o *Ga4ghConsent) GetWasAssentObtained() string`

GetWasAssentObtained returns the WasAssentObtained field if non-nil, zero value otherwise.

### GetWasAssentObtainedOk

`func (o *Ga4ghConsent) GetWasAssentObtainedOk() (string, bool)`

GetWasAssentObtainedOk returns a tuple with the WasAssentObtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWasAssentObtained

`func (o *Ga4ghConsent) HasWasAssentObtained() bool`

HasWasAssentObtained returns a boolean if a field has been set.

### SetWasAssentObtained

`func (o *Ga4ghConsent) SetWasAssentObtained(v string)`

SetWasAssentObtained gets a reference to the given string and assigns it to the WasAssentObtained field.

### GetDateOfAssent

`func (o *Ga4ghConsent) GetDateOfAssent() string`

GetDateOfAssent returns the DateOfAssent field if non-nil, zero value otherwise.

### GetDateOfAssentOk

`func (o *Ga4ghConsent) GetDateOfAssentOk() (string, bool)`

GetDateOfAssentOk returns a tuple with the DateOfAssent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfAssent

`func (o *Ga4ghConsent) HasDateOfAssent() bool`

HasDateOfAssent returns a boolean if a field has been set.

### SetDateOfAssent

`func (o *Ga4ghConsent) SetDateOfAssent(v string)`

SetDateOfAssent gets a reference to the given string and assigns it to the DateOfAssent field.

### GetAssentFormVersion

`func (o *Ga4ghConsent) GetAssentFormVersion() string`

GetAssentFormVersion returns the AssentFormVersion field if non-nil, zero value otherwise.

### GetAssentFormVersionOk

`func (o *Ga4ghConsent) GetAssentFormVersionOk() (string, bool)`

GetAssentFormVersionOk returns a tuple with the AssentFormVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssentFormVersion

`func (o *Ga4ghConsent) HasAssentFormVersion() bool`

HasAssentFormVersion returns a boolean if a field has been set.

### SetAssentFormVersion

`func (o *Ga4ghConsent) SetAssentFormVersion(v string)`

SetAssentFormVersion gets a reference to the given string and assigns it to the AssentFormVersion field.

### GetIfAssentNotObtainedWhyNot

`func (o *Ga4ghConsent) GetIfAssentNotObtainedWhyNot() string`

GetIfAssentNotObtainedWhyNot returns the IfAssentNotObtainedWhyNot field if non-nil, zero value otherwise.

### GetIfAssentNotObtainedWhyNotOk

`func (o *Ga4ghConsent) GetIfAssentNotObtainedWhyNotOk() (string, bool)`

GetIfAssentNotObtainedWhyNotOk returns a tuple with the IfAssentNotObtainedWhyNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIfAssentNotObtainedWhyNot

`func (o *Ga4ghConsent) HasIfAssentNotObtainedWhyNot() bool`

HasIfAssentNotObtainedWhyNot returns a boolean if a field has been set.

### SetIfAssentNotObtainedWhyNot

`func (o *Ga4ghConsent) SetIfAssentNotObtainedWhyNot(v string)`

SetIfAssentNotObtainedWhyNot gets a reference to the given string and assigns it to the IfAssentNotObtainedWhyNot field.

### GetReconsentDate

`func (o *Ga4ghConsent) GetReconsentDate() string`

GetReconsentDate returns the ReconsentDate field if non-nil, zero value otherwise.

### GetReconsentDateOk

`func (o *Ga4ghConsent) GetReconsentDateOk() (string, bool)`

GetReconsentDateOk returns a tuple with the ReconsentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReconsentDate

`func (o *Ga4ghConsent) HasReconsentDate() bool`

HasReconsentDate returns a boolean if a field has been set.

### SetReconsentDate

`func (o *Ga4ghConsent) SetReconsentDate(v string)`

SetReconsentDate gets a reference to the given string and assigns it to the ReconsentDate field.

### GetReconsentVersion

`func (o *Ga4ghConsent) GetReconsentVersion() string`

GetReconsentVersion returns the ReconsentVersion field if non-nil, zero value otherwise.

### GetReconsentVersionOk

`func (o *Ga4ghConsent) GetReconsentVersionOk() (string, bool)`

GetReconsentVersionOk returns a tuple with the ReconsentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReconsentVersion

`func (o *Ga4ghConsent) HasReconsentVersion() bool`

HasReconsentVersion returns a boolean if a field has been set.

### SetReconsentVersion

`func (o *Ga4ghConsent) SetReconsentVersion(v string)`

SetReconsentVersion gets a reference to the given string and assigns it to the ReconsentVersion field.

### GetConsentingCoordinatorName

`func (o *Ga4ghConsent) GetConsentingCoordinatorName() string`

GetConsentingCoordinatorName returns the ConsentingCoordinatorName field if non-nil, zero value otherwise.

### GetConsentingCoordinatorNameOk

`func (o *Ga4ghConsent) GetConsentingCoordinatorNameOk() (string, bool)`

GetConsentingCoordinatorNameOk returns a tuple with the ConsentingCoordinatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentingCoordinatorName

`func (o *Ga4ghConsent) HasConsentingCoordinatorName() bool`

HasConsentingCoordinatorName returns a boolean if a field has been set.

### SetConsentingCoordinatorName

`func (o *Ga4ghConsent) SetConsentingCoordinatorName(v string)`

SetConsentingCoordinatorName gets a reference to the given string and assigns it to the ConsentingCoordinatorName field.

### GetPreviouslyConsented

`func (o *Ga4ghConsent) GetPreviouslyConsented() string`

GetPreviouslyConsented returns the PreviouslyConsented field if non-nil, zero value otherwise.

### GetPreviouslyConsentedOk

`func (o *Ga4ghConsent) GetPreviouslyConsentedOk() (string, bool)`

GetPreviouslyConsentedOk returns a tuple with the PreviouslyConsented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviouslyConsented

`func (o *Ga4ghConsent) HasPreviouslyConsented() bool`

HasPreviouslyConsented returns a boolean if a field has been set.

### SetPreviouslyConsented

`func (o *Ga4ghConsent) SetPreviouslyConsented(v string)`

SetPreviouslyConsented gets a reference to the given string and assigns it to the PreviouslyConsented field.

### GetNameOfOtherBiobank

`func (o *Ga4ghConsent) GetNameOfOtherBiobank() string`

GetNameOfOtherBiobank returns the NameOfOtherBiobank field if non-nil, zero value otherwise.

### GetNameOfOtherBiobankOk

`func (o *Ga4ghConsent) GetNameOfOtherBiobankOk() (string, bool)`

GetNameOfOtherBiobankOk returns a tuple with the NameOfOtherBiobank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNameOfOtherBiobank

`func (o *Ga4ghConsent) HasNameOfOtherBiobank() bool`

HasNameOfOtherBiobank returns a boolean if a field has been set.

### SetNameOfOtherBiobank

`func (o *Ga4ghConsent) SetNameOfOtherBiobank(v string)`

SetNameOfOtherBiobank gets a reference to the given string and assigns it to the NameOfOtherBiobank field.

### GetHasConsentBeenWithdrawn

`func (o *Ga4ghConsent) GetHasConsentBeenWithdrawn() string`

GetHasConsentBeenWithdrawn returns the HasConsentBeenWithdrawn field if non-nil, zero value otherwise.

### GetHasConsentBeenWithdrawnOk

`func (o *Ga4ghConsent) GetHasConsentBeenWithdrawnOk() (string, bool)`

GetHasConsentBeenWithdrawnOk returns a tuple with the HasConsentBeenWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasConsentBeenWithdrawn

`func (o *Ga4ghConsent) HasHasConsentBeenWithdrawn() bool`

HasHasConsentBeenWithdrawn returns a boolean if a field has been set.

### SetHasConsentBeenWithdrawn

`func (o *Ga4ghConsent) SetHasConsentBeenWithdrawn(v string)`

SetHasConsentBeenWithdrawn gets a reference to the given string and assigns it to the HasConsentBeenWithdrawn field.

### GetDateOfConsentWithdrawal

`func (o *Ga4ghConsent) GetDateOfConsentWithdrawal() string`

GetDateOfConsentWithdrawal returns the DateOfConsentWithdrawal field if non-nil, zero value otherwise.

### GetDateOfConsentWithdrawalOk

`func (o *Ga4ghConsent) GetDateOfConsentWithdrawalOk() (string, bool)`

GetDateOfConsentWithdrawalOk returns a tuple with the DateOfConsentWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateOfConsentWithdrawal

`func (o *Ga4ghConsent) HasDateOfConsentWithdrawal() bool`

HasDateOfConsentWithdrawal returns a boolean if a field has been set.

### SetDateOfConsentWithdrawal

`func (o *Ga4ghConsent) SetDateOfConsentWithdrawal(v string)`

SetDateOfConsentWithdrawal gets a reference to the given string and assigns it to the DateOfConsentWithdrawal field.

### GetTypeOfConsentWithdrawal

`func (o *Ga4ghConsent) GetTypeOfConsentWithdrawal() string`

GetTypeOfConsentWithdrawal returns the TypeOfConsentWithdrawal field if non-nil, zero value otherwise.

### GetTypeOfConsentWithdrawalOk

`func (o *Ga4ghConsent) GetTypeOfConsentWithdrawalOk() (string, bool)`

GetTypeOfConsentWithdrawalOk returns a tuple with the TypeOfConsentWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeOfConsentWithdrawal

`func (o *Ga4ghConsent) HasTypeOfConsentWithdrawal() bool`

HasTypeOfConsentWithdrawal returns a boolean if a field has been set.

### SetTypeOfConsentWithdrawal

`func (o *Ga4ghConsent) SetTypeOfConsentWithdrawal(v string)`

SetTypeOfConsentWithdrawal gets a reference to the given string and assigns it to the TypeOfConsentWithdrawal field.

### GetReasonForConsentWithdrawal

`func (o *Ga4ghConsent) GetReasonForConsentWithdrawal() string`

GetReasonForConsentWithdrawal returns the ReasonForConsentWithdrawal field if non-nil, zero value otherwise.

### GetReasonForConsentWithdrawalOk

`func (o *Ga4ghConsent) GetReasonForConsentWithdrawalOk() (string, bool)`

GetReasonForConsentWithdrawalOk returns a tuple with the ReasonForConsentWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReasonForConsentWithdrawal

`func (o *Ga4ghConsent) HasReasonForConsentWithdrawal() bool`

HasReasonForConsentWithdrawal returns a boolean if a field has been set.

### SetReasonForConsentWithdrawal

`func (o *Ga4ghConsent) SetReasonForConsentWithdrawal(v string)`

SetReasonForConsentWithdrawal gets a reference to the given string and assigns it to the ReasonForConsentWithdrawal field.

### GetConsentFormComplete

`func (o *Ga4ghConsent) GetConsentFormComplete() string`

GetConsentFormComplete returns the ConsentFormComplete field if non-nil, zero value otherwise.

### GetConsentFormCompleteOk

`func (o *Ga4ghConsent) GetConsentFormCompleteOk() (string, bool)`

GetConsentFormCompleteOk returns a tuple with the ConsentFormComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentFormComplete

`func (o *Ga4ghConsent) HasConsentFormComplete() bool`

HasConsentFormComplete returns a boolean if a field has been set.

### SetConsentFormComplete

`func (o *Ga4ghConsent) SetConsentFormComplete(v string)`

SetConsentFormComplete gets a reference to the given string and assigns it to the ConsentFormComplete field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


