/*
 * CanDIG Services
 *
 * Below is a list of APIs that CanDIG currently supports.<br/><br/>For /search and /count endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/search_count_services_usage.pdf) for instructions and sample queries.<br/>For all metadata and variant services endpoints, refer to [this documentation](https://www.distributedgenomics.ca/static/metadata_variants_services_sample_queries.pdf) for sample queries.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"encoding/json"
)

// Ga4ghConsent struct for Ga4ghConsent
type Ga4ghConsent struct {
	// This is unique in the context of the server instance.
	Id *string `json:"id,omitempty"`

	// The ID of the dataset this object belongs to.
	DatasetId *string `json:"datasetId,omitempty"`

	// This is a label or symbolic identifier for this object.
	Name *string `json:"name,omitempty"`

	// This attribute contains human readable text.
	Description *string `json:"description,omitempty"`

	// The :ref:`ISO 8601<metadata_date_time>` time at which this record was created.
	Created *string `json:"created,omitempty"`

	// The :ref:`ISO 8601<metadata_date_time>` time at which this record was updated.
	Updated *string `json:"updated,omitempty"`

	PatientId *string `json:"patientId,omitempty"`

	ConsentId *string `json:"consentId,omitempty"`

	ConsentDate *string `json:"consentDate,omitempty"`

	ConsentVersion *string `json:"consentVersion,omitempty"`

	PatientConsentedTo *string `json:"patientConsentedTo,omitempty"`

	ReasonForRejection *string `json:"reasonForRejection,omitempty"`

	WasAssentObtained *string `json:"wasAssentObtained,omitempty"`

	DateOfAssent *string `json:"dateOfAssent,omitempty"`

	AssentFormVersion *string `json:"assentFormVersion,omitempty"`

	IfAssentNotObtainedWhyNot *string `json:"ifAssentNotObtainedWhyNot,omitempty"`

	ReconsentDate *string `json:"reconsentDate,omitempty"`

	ReconsentVersion *string `json:"reconsentVersion,omitempty"`

	ConsentingCoordinatorName *string `json:"consentingCoordinatorName,omitempty"`

	PreviouslyConsented *string `json:"previouslyConsented,omitempty"`

	NameOfOtherBiobank *string `json:"nameOfOtherBiobank,omitempty"`

	HasConsentBeenWithdrawn *string `json:"hasConsentBeenWithdrawn,omitempty"`

	DateOfConsentWithdrawal *string `json:"dateOfConsentWithdrawal,omitempty"`

	TypeOfConsentWithdrawal *string `json:"typeOfConsentWithdrawal,omitempty"`

	ReasonForConsentWithdrawal *string `json:"reasonForConsentWithdrawal,omitempty"`

	ConsentFormComplete *string `json:"consentFormComplete,omitempty"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghConsent) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghConsent) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghConsent) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghConsent) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghConsent) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghConsent) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghConsent) SetPatientId(v string) {
	o.PatientId = &v
}

// GetConsentId returns the ConsentId field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetConsentId() string {
	if o == nil || o.ConsentId == nil {
		var ret string
		return ret
	}
	return *o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetConsentIdOk() (string, bool) {
	if o == nil || o.ConsentId == nil {
		var ret string
		return ret, false
	}
	return *o.ConsentId, true
}

// HasConsentId returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasConsentId() bool {
	if o != nil && o.ConsentId != nil {
		return true
	}

	return false
}

// SetConsentId gets a reference to the given string and assigns it to the ConsentId field.
func (o *Ga4ghConsent) SetConsentId(v string) {
	o.ConsentId = &v
}

// GetConsentDate returns the ConsentDate field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetConsentDate() string {
	if o == nil || o.ConsentDate == nil {
		var ret string
		return ret
	}
	return *o.ConsentDate
}

// GetConsentDateOk returns a tuple with the ConsentDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetConsentDateOk() (string, bool) {
	if o == nil || o.ConsentDate == nil {
		var ret string
		return ret, false
	}
	return *o.ConsentDate, true
}

// HasConsentDate returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasConsentDate() bool {
	if o != nil && o.ConsentDate != nil {
		return true
	}

	return false
}

// SetConsentDate gets a reference to the given string and assigns it to the ConsentDate field.
func (o *Ga4ghConsent) SetConsentDate(v string) {
	o.ConsentDate = &v
}

// GetConsentVersion returns the ConsentVersion field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetConsentVersion() string {
	if o == nil || o.ConsentVersion == nil {
		var ret string
		return ret
	}
	return *o.ConsentVersion
}

// GetConsentVersionOk returns a tuple with the ConsentVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetConsentVersionOk() (string, bool) {
	if o == nil || o.ConsentVersion == nil {
		var ret string
		return ret, false
	}
	return *o.ConsentVersion, true
}

// HasConsentVersion returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasConsentVersion() bool {
	if o != nil && o.ConsentVersion != nil {
		return true
	}

	return false
}

// SetConsentVersion gets a reference to the given string and assigns it to the ConsentVersion field.
func (o *Ga4ghConsent) SetConsentVersion(v string) {
	o.ConsentVersion = &v
}

// GetPatientConsentedTo returns the PatientConsentedTo field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetPatientConsentedTo() string {
	if o == nil || o.PatientConsentedTo == nil {
		var ret string
		return ret
	}
	return *o.PatientConsentedTo
}

// GetPatientConsentedToOk returns a tuple with the PatientConsentedTo field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetPatientConsentedToOk() (string, bool) {
	if o == nil || o.PatientConsentedTo == nil {
		var ret string
		return ret, false
	}
	return *o.PatientConsentedTo, true
}

// HasPatientConsentedTo returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasPatientConsentedTo() bool {
	if o != nil && o.PatientConsentedTo != nil {
		return true
	}

	return false
}

// SetPatientConsentedTo gets a reference to the given string and assigns it to the PatientConsentedTo field.
func (o *Ga4ghConsent) SetPatientConsentedTo(v string) {
	o.PatientConsentedTo = &v
}

// GetReasonForRejection returns the ReasonForRejection field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetReasonForRejection() string {
	if o == nil || o.ReasonForRejection == nil {
		var ret string
		return ret
	}
	return *o.ReasonForRejection
}

// GetReasonForRejectionOk returns a tuple with the ReasonForRejection field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetReasonForRejectionOk() (string, bool) {
	if o == nil || o.ReasonForRejection == nil {
		var ret string
		return ret, false
	}
	return *o.ReasonForRejection, true
}

// HasReasonForRejection returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasReasonForRejection() bool {
	if o != nil && o.ReasonForRejection != nil {
		return true
	}

	return false
}

// SetReasonForRejection gets a reference to the given string and assigns it to the ReasonForRejection field.
func (o *Ga4ghConsent) SetReasonForRejection(v string) {
	o.ReasonForRejection = &v
}

// GetWasAssentObtained returns the WasAssentObtained field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetWasAssentObtained() string {
	if o == nil || o.WasAssentObtained == nil {
		var ret string
		return ret
	}
	return *o.WasAssentObtained
}

// GetWasAssentObtainedOk returns a tuple with the WasAssentObtained field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetWasAssentObtainedOk() (string, bool) {
	if o == nil || o.WasAssentObtained == nil {
		var ret string
		return ret, false
	}
	return *o.WasAssentObtained, true
}

// HasWasAssentObtained returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasWasAssentObtained() bool {
	if o != nil && o.WasAssentObtained != nil {
		return true
	}

	return false
}

// SetWasAssentObtained gets a reference to the given string and assigns it to the WasAssentObtained field.
func (o *Ga4ghConsent) SetWasAssentObtained(v string) {
	o.WasAssentObtained = &v
}

// GetDateOfAssent returns the DateOfAssent field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetDateOfAssent() string {
	if o == nil || o.DateOfAssent == nil {
		var ret string
		return ret
	}
	return *o.DateOfAssent
}

// GetDateOfAssentOk returns a tuple with the DateOfAssent field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetDateOfAssentOk() (string, bool) {
	if o == nil || o.DateOfAssent == nil {
		var ret string
		return ret, false
	}
	return *o.DateOfAssent, true
}

// HasDateOfAssent returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasDateOfAssent() bool {
	if o != nil && o.DateOfAssent != nil {
		return true
	}

	return false
}

// SetDateOfAssent gets a reference to the given string and assigns it to the DateOfAssent field.
func (o *Ga4ghConsent) SetDateOfAssent(v string) {
	o.DateOfAssent = &v
}

// GetAssentFormVersion returns the AssentFormVersion field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetAssentFormVersion() string {
	if o == nil || o.AssentFormVersion == nil {
		var ret string
		return ret
	}
	return *o.AssentFormVersion
}

// GetAssentFormVersionOk returns a tuple with the AssentFormVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetAssentFormVersionOk() (string, bool) {
	if o == nil || o.AssentFormVersion == nil {
		var ret string
		return ret, false
	}
	return *o.AssentFormVersion, true
}

// HasAssentFormVersion returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasAssentFormVersion() bool {
	if o != nil && o.AssentFormVersion != nil {
		return true
	}

	return false
}

// SetAssentFormVersion gets a reference to the given string and assigns it to the AssentFormVersion field.
func (o *Ga4ghConsent) SetAssentFormVersion(v string) {
	o.AssentFormVersion = &v
}

// GetIfAssentNotObtainedWhyNot returns the IfAssentNotObtainedWhyNot field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetIfAssentNotObtainedWhyNot() string {
	if o == nil || o.IfAssentNotObtainedWhyNot == nil {
		var ret string
		return ret
	}
	return *o.IfAssentNotObtainedWhyNot
}

// GetIfAssentNotObtainedWhyNotOk returns a tuple with the IfAssentNotObtainedWhyNot field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetIfAssentNotObtainedWhyNotOk() (string, bool) {
	if o == nil || o.IfAssentNotObtainedWhyNot == nil {
		var ret string
		return ret, false
	}
	return *o.IfAssentNotObtainedWhyNot, true
}

// HasIfAssentNotObtainedWhyNot returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasIfAssentNotObtainedWhyNot() bool {
	if o != nil && o.IfAssentNotObtainedWhyNot != nil {
		return true
	}

	return false
}

// SetIfAssentNotObtainedWhyNot gets a reference to the given string and assigns it to the IfAssentNotObtainedWhyNot field.
func (o *Ga4ghConsent) SetIfAssentNotObtainedWhyNot(v string) {
	o.IfAssentNotObtainedWhyNot = &v
}

// GetReconsentDate returns the ReconsentDate field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetReconsentDate() string {
	if o == nil || o.ReconsentDate == nil {
		var ret string
		return ret
	}
	return *o.ReconsentDate
}

// GetReconsentDateOk returns a tuple with the ReconsentDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetReconsentDateOk() (string, bool) {
	if o == nil || o.ReconsentDate == nil {
		var ret string
		return ret, false
	}
	return *o.ReconsentDate, true
}

// HasReconsentDate returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasReconsentDate() bool {
	if o != nil && o.ReconsentDate != nil {
		return true
	}

	return false
}

// SetReconsentDate gets a reference to the given string and assigns it to the ReconsentDate field.
func (o *Ga4ghConsent) SetReconsentDate(v string) {
	o.ReconsentDate = &v
}

// GetReconsentVersion returns the ReconsentVersion field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetReconsentVersion() string {
	if o == nil || o.ReconsentVersion == nil {
		var ret string
		return ret
	}
	return *o.ReconsentVersion
}

// GetReconsentVersionOk returns a tuple with the ReconsentVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetReconsentVersionOk() (string, bool) {
	if o == nil || o.ReconsentVersion == nil {
		var ret string
		return ret, false
	}
	return *o.ReconsentVersion, true
}

// HasReconsentVersion returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasReconsentVersion() bool {
	if o != nil && o.ReconsentVersion != nil {
		return true
	}

	return false
}

// SetReconsentVersion gets a reference to the given string and assigns it to the ReconsentVersion field.
func (o *Ga4ghConsent) SetReconsentVersion(v string) {
	o.ReconsentVersion = &v
}

// GetConsentingCoordinatorName returns the ConsentingCoordinatorName field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetConsentingCoordinatorName() string {
	if o == nil || o.ConsentingCoordinatorName == nil {
		var ret string
		return ret
	}
	return *o.ConsentingCoordinatorName
}

// GetConsentingCoordinatorNameOk returns a tuple with the ConsentingCoordinatorName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetConsentingCoordinatorNameOk() (string, bool) {
	if o == nil || o.ConsentingCoordinatorName == nil {
		var ret string
		return ret, false
	}
	return *o.ConsentingCoordinatorName, true
}

// HasConsentingCoordinatorName returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasConsentingCoordinatorName() bool {
	if o != nil && o.ConsentingCoordinatorName != nil {
		return true
	}

	return false
}

// SetConsentingCoordinatorName gets a reference to the given string and assigns it to the ConsentingCoordinatorName field.
func (o *Ga4ghConsent) SetConsentingCoordinatorName(v string) {
	o.ConsentingCoordinatorName = &v
}

// GetPreviouslyConsented returns the PreviouslyConsented field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetPreviouslyConsented() string {
	if o == nil || o.PreviouslyConsented == nil {
		var ret string
		return ret
	}
	return *o.PreviouslyConsented
}

// GetPreviouslyConsentedOk returns a tuple with the PreviouslyConsented field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetPreviouslyConsentedOk() (string, bool) {
	if o == nil || o.PreviouslyConsented == nil {
		var ret string
		return ret, false
	}
	return *o.PreviouslyConsented, true
}

// HasPreviouslyConsented returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasPreviouslyConsented() bool {
	if o != nil && o.PreviouslyConsented != nil {
		return true
	}

	return false
}

// SetPreviouslyConsented gets a reference to the given string and assigns it to the PreviouslyConsented field.
func (o *Ga4ghConsent) SetPreviouslyConsented(v string) {
	o.PreviouslyConsented = &v
}

// GetNameOfOtherBiobank returns the NameOfOtherBiobank field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetNameOfOtherBiobank() string {
	if o == nil || o.NameOfOtherBiobank == nil {
		var ret string
		return ret
	}
	return *o.NameOfOtherBiobank
}

// GetNameOfOtherBiobankOk returns a tuple with the NameOfOtherBiobank field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetNameOfOtherBiobankOk() (string, bool) {
	if o == nil || o.NameOfOtherBiobank == nil {
		var ret string
		return ret, false
	}
	return *o.NameOfOtherBiobank, true
}

// HasNameOfOtherBiobank returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasNameOfOtherBiobank() bool {
	if o != nil && o.NameOfOtherBiobank != nil {
		return true
	}

	return false
}

// SetNameOfOtherBiobank gets a reference to the given string and assigns it to the NameOfOtherBiobank field.
func (o *Ga4ghConsent) SetNameOfOtherBiobank(v string) {
	o.NameOfOtherBiobank = &v
}

// GetHasConsentBeenWithdrawn returns the HasConsentBeenWithdrawn field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetHasConsentBeenWithdrawn() string {
	if o == nil || o.HasConsentBeenWithdrawn == nil {
		var ret string
		return ret
	}
	return *o.HasConsentBeenWithdrawn
}

// GetHasConsentBeenWithdrawnOk returns a tuple with the HasConsentBeenWithdrawn field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetHasConsentBeenWithdrawnOk() (string, bool) {
	if o == nil || o.HasConsentBeenWithdrawn == nil {
		var ret string
		return ret, false
	}
	return *o.HasConsentBeenWithdrawn, true
}

// HasHasConsentBeenWithdrawn returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasHasConsentBeenWithdrawn() bool {
	if o != nil && o.HasConsentBeenWithdrawn != nil {
		return true
	}

	return false
}

// SetHasConsentBeenWithdrawn gets a reference to the given string and assigns it to the HasConsentBeenWithdrawn field.
func (o *Ga4ghConsent) SetHasConsentBeenWithdrawn(v string) {
	o.HasConsentBeenWithdrawn = &v
}

// GetDateOfConsentWithdrawal returns the DateOfConsentWithdrawal field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetDateOfConsentWithdrawal() string {
	if o == nil || o.DateOfConsentWithdrawal == nil {
		var ret string
		return ret
	}
	return *o.DateOfConsentWithdrawal
}

// GetDateOfConsentWithdrawalOk returns a tuple with the DateOfConsentWithdrawal field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetDateOfConsentWithdrawalOk() (string, bool) {
	if o == nil || o.DateOfConsentWithdrawal == nil {
		var ret string
		return ret, false
	}
	return *o.DateOfConsentWithdrawal, true
}

// HasDateOfConsentWithdrawal returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasDateOfConsentWithdrawal() bool {
	if o != nil && o.DateOfConsentWithdrawal != nil {
		return true
	}

	return false
}

// SetDateOfConsentWithdrawal gets a reference to the given string and assigns it to the DateOfConsentWithdrawal field.
func (o *Ga4ghConsent) SetDateOfConsentWithdrawal(v string) {
	o.DateOfConsentWithdrawal = &v
}

// GetTypeOfConsentWithdrawal returns the TypeOfConsentWithdrawal field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetTypeOfConsentWithdrawal() string {
	if o == nil || o.TypeOfConsentWithdrawal == nil {
		var ret string
		return ret
	}
	return *o.TypeOfConsentWithdrawal
}

// GetTypeOfConsentWithdrawalOk returns a tuple with the TypeOfConsentWithdrawal field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetTypeOfConsentWithdrawalOk() (string, bool) {
	if o == nil || o.TypeOfConsentWithdrawal == nil {
		var ret string
		return ret, false
	}
	return *o.TypeOfConsentWithdrawal, true
}

// HasTypeOfConsentWithdrawal returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasTypeOfConsentWithdrawal() bool {
	if o != nil && o.TypeOfConsentWithdrawal != nil {
		return true
	}

	return false
}

// SetTypeOfConsentWithdrawal gets a reference to the given string and assigns it to the TypeOfConsentWithdrawal field.
func (o *Ga4ghConsent) SetTypeOfConsentWithdrawal(v string) {
	o.TypeOfConsentWithdrawal = &v
}

// GetReasonForConsentWithdrawal returns the ReasonForConsentWithdrawal field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetReasonForConsentWithdrawal() string {
	if o == nil || o.ReasonForConsentWithdrawal == nil {
		var ret string
		return ret
	}
	return *o.ReasonForConsentWithdrawal
}

// GetReasonForConsentWithdrawalOk returns a tuple with the ReasonForConsentWithdrawal field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetReasonForConsentWithdrawalOk() (string, bool) {
	if o == nil || o.ReasonForConsentWithdrawal == nil {
		var ret string
		return ret, false
	}
	return *o.ReasonForConsentWithdrawal, true
}

// HasReasonForConsentWithdrawal returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasReasonForConsentWithdrawal() bool {
	if o != nil && o.ReasonForConsentWithdrawal != nil {
		return true
	}

	return false
}

// SetReasonForConsentWithdrawal gets a reference to the given string and assigns it to the ReasonForConsentWithdrawal field.
func (o *Ga4ghConsent) SetReasonForConsentWithdrawal(v string) {
	o.ReasonForConsentWithdrawal = &v
}

// GetConsentFormComplete returns the ConsentFormComplete field if non-nil, zero value otherwise.
func (o *Ga4ghConsent) GetConsentFormComplete() string {
	if o == nil || o.ConsentFormComplete == nil {
		var ret string
		return ret
	}
	return *o.ConsentFormComplete
}

// GetConsentFormCompleteOk returns a tuple with the ConsentFormComplete field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghConsent) GetConsentFormCompleteOk() (string, bool) {
	if o == nil || o.ConsentFormComplete == nil {
		var ret string
		return ret, false
	}
	return *o.ConsentFormComplete, true
}

// HasConsentFormComplete returns a boolean if a field has been set.
func (o *Ga4ghConsent) HasConsentFormComplete() bool {
	if o != nil && o.ConsentFormComplete != nil {
		return true
	}

	return false
}

// SetConsentFormComplete gets a reference to the given string and assigns it to the ConsentFormComplete field.
func (o *Ga4ghConsent) SetConsentFormComplete(v string) {
	o.ConsentFormComplete = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o Ga4ghConsent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DatasetId != nil {
		toSerialize["datasetId"] = o.DatasetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.PatientId != nil {
		toSerialize["patientId"] = o.PatientId
	}
	if o.ConsentId != nil {
		toSerialize["consentId"] = o.ConsentId
	}
	if o.ConsentDate != nil {
		toSerialize["consentDate"] = o.ConsentDate
	}
	if o.ConsentVersion != nil {
		toSerialize["consentVersion"] = o.ConsentVersion
	}
	if o.PatientConsentedTo != nil {
		toSerialize["patientConsentedTo"] = o.PatientConsentedTo
	}
	if o.ReasonForRejection != nil {
		toSerialize["reasonForRejection"] = o.ReasonForRejection
	}
	if o.WasAssentObtained != nil {
		toSerialize["wasAssentObtained"] = o.WasAssentObtained
	}
	if o.DateOfAssent != nil {
		toSerialize["dateOfAssent"] = o.DateOfAssent
	}
	if o.AssentFormVersion != nil {
		toSerialize["assentFormVersion"] = o.AssentFormVersion
	}
	if o.IfAssentNotObtainedWhyNot != nil {
		toSerialize["ifAssentNotObtainedWhyNot"] = o.IfAssentNotObtainedWhyNot
	}
	if o.ReconsentDate != nil {
		toSerialize["reconsentDate"] = o.ReconsentDate
	}
	if o.ReconsentVersion != nil {
		toSerialize["reconsentVersion"] = o.ReconsentVersion
	}
	if o.ConsentingCoordinatorName != nil {
		toSerialize["consentingCoordinatorName"] = o.ConsentingCoordinatorName
	}
	if o.PreviouslyConsented != nil {
		toSerialize["previouslyConsented"] = o.PreviouslyConsented
	}
	if o.NameOfOtherBiobank != nil {
		toSerialize["nameOfOtherBiobank"] = o.NameOfOtherBiobank
	}
	if o.HasConsentBeenWithdrawn != nil {
		toSerialize["hasConsentBeenWithdrawn"] = o.HasConsentBeenWithdrawn
	}
	if o.DateOfConsentWithdrawal != nil {
		toSerialize["dateOfConsentWithdrawal"] = o.DateOfConsentWithdrawal
	}
	if o.TypeOfConsentWithdrawal != nil {
		toSerialize["typeOfConsentWithdrawal"] = o.TypeOfConsentWithdrawal
	}
	if o.ReasonForConsentWithdrawal != nil {
		toSerialize["reasonForConsentWithdrawal"] = o.ReasonForConsentWithdrawal
	}
	if o.ConsentFormComplete != nil {
		toSerialize["consentFormComplete"] = o.ConsentFormComplete
	}
	return json.Marshal(toSerialize)
}
