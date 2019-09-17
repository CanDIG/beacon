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

type Ga4ghEnrollment struct {
	// This is unique in the context of the server instance.
	Id *string `json:"id,omitempty"`

	// The ID of the dataset this object belongs to.
	DatasetId *string `json:"dataset_id,omitempty"`

	// This is a label or symbolic identifier for this object.
	Name *string `json:"name,omitempty"`

	// This attribute contains human readable text.
	Description *string `json:"description,omitempty"`

	// The :ref:`ISO 8601<metadata_date_time>` time at which this record was created.
	Created *string `json:"created,omitempty"`

	// The :ref:`ISO 8601<metadata_date_time>` time at which this record was updated.
	Updated *string `json:"updated,omitempty"`

	PatientId *string `json:"patientId,omitempty"`

	EnrollmentInstitution *string `json:"enrollmentInstitution,omitempty"`

	EnrollmentApprovalDate *string `json:"enrollmentApprovalDate,omitempty"`

	CrossEnrollment *string `json:"crossEnrollment,omitempty"`

	OtherPersonalizedMedicineStudyName *string `json:"otherPersonalizedMedicineStudyName,omitempty"`

	OtherPersonalizedMedicineStudyId *string `json:"otherPersonalizedMedicineStudyId,omitempty"`

	AgeAtEnrollment *string `json:"ageAtEnrollment,omitempty"`

	EligibilityCategory *string `json:"eligibilityCategory,omitempty"`

	StatusAtEnrollment *string `json:"statusAtEnrollment,omitempty"`

	PrimaryOncologistName *string `json:"primaryOncologistName,omitempty"`

	PrimaryOncologistContact *string `json:"primaryOncologistContact,omitempty"`

	ReferringPhysicianName *string `json:"referringPhysicianName,omitempty"`

	ReferringPhysicianContact *string `json:"referringPhysicianContact,omitempty"`

	SummaryOfIdRequest *string `json:"summaryOfIdRequest,omitempty"`

	TreatingCentreName *string `json:"treatingCentreName,omitempty"`

	TreatingCentreProvince *string `json:"treatingCentreProvince,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghEnrollment) SetId(v string) {
	o.Id = &v
}

// GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetDatasetIdOk() (string, bool) {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret, false
	}
	return *o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Ga4ghEnrollment) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghEnrollment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Ga4ghEnrollment) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetCreatedOk() (string, bool) {
	if o == nil || o.Created == nil {
		var ret string
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Ga4ghEnrollment) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetUpdatedOk() (string, bool) {
	if o == nil || o.Updated == nil {
		var ret string
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Ga4ghEnrollment) SetUpdated(v string) {
	o.Updated = &v
}

// GetPatientId returns the PatientId field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetPatientId() string {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret
	}
	return *o.PatientId
}

// GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetPatientIdOk() (string, bool) {
	if o == nil || o.PatientId == nil {
		var ret string
		return ret, false
	}
	return *o.PatientId, true
}

// HasPatientId returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasPatientId() bool {
	if o != nil && o.PatientId != nil {
		return true
	}

	return false
}

// SetPatientId gets a reference to the given string and assigns it to the PatientId field.
func (o *Ga4ghEnrollment) SetPatientId(v string) {
	o.PatientId = &v
}

// GetEnrollmentInstitution returns the EnrollmentInstitution field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetEnrollmentInstitution() string {
	if o == nil || o.EnrollmentInstitution == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentInstitution
}

// GetEnrollmentInstitutionOk returns a tuple with the EnrollmentInstitution field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetEnrollmentInstitutionOk() (string, bool) {
	if o == nil || o.EnrollmentInstitution == nil {
		var ret string
		return ret, false
	}
	return *o.EnrollmentInstitution, true
}

// HasEnrollmentInstitution returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasEnrollmentInstitution() bool {
	if o != nil && o.EnrollmentInstitution != nil {
		return true
	}

	return false
}

// SetEnrollmentInstitution gets a reference to the given string and assigns it to the EnrollmentInstitution field.
func (o *Ga4ghEnrollment) SetEnrollmentInstitution(v string) {
	o.EnrollmentInstitution = &v
}

// GetEnrollmentApprovalDate returns the EnrollmentApprovalDate field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetEnrollmentApprovalDate() string {
	if o == nil || o.EnrollmentApprovalDate == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentApprovalDate
}

// GetEnrollmentApprovalDateOk returns a tuple with the EnrollmentApprovalDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetEnrollmentApprovalDateOk() (string, bool) {
	if o == nil || o.EnrollmentApprovalDate == nil {
		var ret string
		return ret, false
	}
	return *o.EnrollmentApprovalDate, true
}

// HasEnrollmentApprovalDate returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasEnrollmentApprovalDate() bool {
	if o != nil && o.EnrollmentApprovalDate != nil {
		return true
	}

	return false
}

// SetEnrollmentApprovalDate gets a reference to the given string and assigns it to the EnrollmentApprovalDate field.
func (o *Ga4ghEnrollment) SetEnrollmentApprovalDate(v string) {
	o.EnrollmentApprovalDate = &v
}

// GetCrossEnrollment returns the CrossEnrollment field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetCrossEnrollment() string {
	if o == nil || o.CrossEnrollment == nil {
		var ret string
		return ret
	}
	return *o.CrossEnrollment
}

// GetCrossEnrollmentOk returns a tuple with the CrossEnrollment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetCrossEnrollmentOk() (string, bool) {
	if o == nil || o.CrossEnrollment == nil {
		var ret string
		return ret, false
	}
	return *o.CrossEnrollment, true
}

// HasCrossEnrollment returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasCrossEnrollment() bool {
	if o != nil && o.CrossEnrollment != nil {
		return true
	}

	return false
}

// SetCrossEnrollment gets a reference to the given string and assigns it to the CrossEnrollment field.
func (o *Ga4ghEnrollment) SetCrossEnrollment(v string) {
	o.CrossEnrollment = &v
}

// GetOtherPersonalizedMedicineStudyName returns the OtherPersonalizedMedicineStudyName field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyName() string {
	if o == nil || o.OtherPersonalizedMedicineStudyName == nil {
		var ret string
		return ret
	}
	return *o.OtherPersonalizedMedicineStudyName
}

// GetOtherPersonalizedMedicineStudyNameOk returns a tuple with the OtherPersonalizedMedicineStudyName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyNameOk() (string, bool) {
	if o == nil || o.OtherPersonalizedMedicineStudyName == nil {
		var ret string
		return ret, false
	}
	return *o.OtherPersonalizedMedicineStudyName, true
}

// HasOtherPersonalizedMedicineStudyName returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasOtherPersonalizedMedicineStudyName() bool {
	if o != nil && o.OtherPersonalizedMedicineStudyName != nil {
		return true
	}

	return false
}

// SetOtherPersonalizedMedicineStudyName gets a reference to the given string and assigns it to the OtherPersonalizedMedicineStudyName field.
func (o *Ga4ghEnrollment) SetOtherPersonalizedMedicineStudyName(v string) {
	o.OtherPersonalizedMedicineStudyName = &v
}

// GetOtherPersonalizedMedicineStudyId returns the OtherPersonalizedMedicineStudyId field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyId() string {
	if o == nil || o.OtherPersonalizedMedicineStudyId == nil {
		var ret string
		return ret
	}
	return *o.OtherPersonalizedMedicineStudyId
}

// GetOtherPersonalizedMedicineStudyIdOk returns a tuple with the OtherPersonalizedMedicineStudyId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetOtherPersonalizedMedicineStudyIdOk() (string, bool) {
	if o == nil || o.OtherPersonalizedMedicineStudyId == nil {
		var ret string
		return ret, false
	}
	return *o.OtherPersonalizedMedicineStudyId, true
}

// HasOtherPersonalizedMedicineStudyId returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasOtherPersonalizedMedicineStudyId() bool {
	if o != nil && o.OtherPersonalizedMedicineStudyId != nil {
		return true
	}

	return false
}

// SetOtherPersonalizedMedicineStudyId gets a reference to the given string and assigns it to the OtherPersonalizedMedicineStudyId field.
func (o *Ga4ghEnrollment) SetOtherPersonalizedMedicineStudyId(v string) {
	o.OtherPersonalizedMedicineStudyId = &v
}

// GetAgeAtEnrollment returns the AgeAtEnrollment field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetAgeAtEnrollment() string {
	if o == nil || o.AgeAtEnrollment == nil {
		var ret string
		return ret
	}
	return *o.AgeAtEnrollment
}

// GetAgeAtEnrollmentOk returns a tuple with the AgeAtEnrollment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetAgeAtEnrollmentOk() (string, bool) {
	if o == nil || o.AgeAtEnrollment == nil {
		var ret string
		return ret, false
	}
	return *o.AgeAtEnrollment, true
}

// HasAgeAtEnrollment returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasAgeAtEnrollment() bool {
	if o != nil && o.AgeAtEnrollment != nil {
		return true
	}

	return false
}

// SetAgeAtEnrollment gets a reference to the given string and assigns it to the AgeAtEnrollment field.
func (o *Ga4ghEnrollment) SetAgeAtEnrollment(v string) {
	o.AgeAtEnrollment = &v
}

// GetEligibilityCategory returns the EligibilityCategory field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetEligibilityCategory() string {
	if o == nil || o.EligibilityCategory == nil {
		var ret string
		return ret
	}
	return *o.EligibilityCategory
}

// GetEligibilityCategoryOk returns a tuple with the EligibilityCategory field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetEligibilityCategoryOk() (string, bool) {
	if o == nil || o.EligibilityCategory == nil {
		var ret string
		return ret, false
	}
	return *o.EligibilityCategory, true
}

// HasEligibilityCategory returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasEligibilityCategory() bool {
	if o != nil && o.EligibilityCategory != nil {
		return true
	}

	return false
}

// SetEligibilityCategory gets a reference to the given string and assigns it to the EligibilityCategory field.
func (o *Ga4ghEnrollment) SetEligibilityCategory(v string) {
	o.EligibilityCategory = &v
}

// GetStatusAtEnrollment returns the StatusAtEnrollment field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetStatusAtEnrollment() string {
	if o == nil || o.StatusAtEnrollment == nil {
		var ret string
		return ret
	}
	return *o.StatusAtEnrollment
}

// GetStatusAtEnrollmentOk returns a tuple with the StatusAtEnrollment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetStatusAtEnrollmentOk() (string, bool) {
	if o == nil || o.StatusAtEnrollment == nil {
		var ret string
		return ret, false
	}
	return *o.StatusAtEnrollment, true
}

// HasStatusAtEnrollment returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasStatusAtEnrollment() bool {
	if o != nil && o.StatusAtEnrollment != nil {
		return true
	}

	return false
}

// SetStatusAtEnrollment gets a reference to the given string and assigns it to the StatusAtEnrollment field.
func (o *Ga4ghEnrollment) SetStatusAtEnrollment(v string) {
	o.StatusAtEnrollment = &v
}

// GetPrimaryOncologistName returns the PrimaryOncologistName field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetPrimaryOncologistName() string {
	if o == nil || o.PrimaryOncologistName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryOncologistName
}

// GetPrimaryOncologistNameOk returns a tuple with the PrimaryOncologistName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetPrimaryOncologistNameOk() (string, bool) {
	if o == nil || o.PrimaryOncologistName == nil {
		var ret string
		return ret, false
	}
	return *o.PrimaryOncologistName, true
}

// HasPrimaryOncologistName returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasPrimaryOncologistName() bool {
	if o != nil && o.PrimaryOncologistName != nil {
		return true
	}

	return false
}

// SetPrimaryOncologistName gets a reference to the given string and assigns it to the PrimaryOncologistName field.
func (o *Ga4ghEnrollment) SetPrimaryOncologistName(v string) {
	o.PrimaryOncologistName = &v
}

// GetPrimaryOncologistContact returns the PrimaryOncologistContact field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetPrimaryOncologistContact() string {
	if o == nil || o.PrimaryOncologistContact == nil {
		var ret string
		return ret
	}
	return *o.PrimaryOncologistContact
}

// GetPrimaryOncologistContactOk returns a tuple with the PrimaryOncologistContact field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetPrimaryOncologistContactOk() (string, bool) {
	if o == nil || o.PrimaryOncologistContact == nil {
		var ret string
		return ret, false
	}
	return *o.PrimaryOncologistContact, true
}

// HasPrimaryOncologistContact returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasPrimaryOncologistContact() bool {
	if o != nil && o.PrimaryOncologistContact != nil {
		return true
	}

	return false
}

// SetPrimaryOncologistContact gets a reference to the given string and assigns it to the PrimaryOncologistContact field.
func (o *Ga4ghEnrollment) SetPrimaryOncologistContact(v string) {
	o.PrimaryOncologistContact = &v
}

// GetReferringPhysicianName returns the ReferringPhysicianName field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetReferringPhysicianName() string {
	if o == nil || o.ReferringPhysicianName == nil {
		var ret string
		return ret
	}
	return *o.ReferringPhysicianName
}

// GetReferringPhysicianNameOk returns a tuple with the ReferringPhysicianName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetReferringPhysicianNameOk() (string, bool) {
	if o == nil || o.ReferringPhysicianName == nil {
		var ret string
		return ret, false
	}
	return *o.ReferringPhysicianName, true
}

// HasReferringPhysicianName returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasReferringPhysicianName() bool {
	if o != nil && o.ReferringPhysicianName != nil {
		return true
	}

	return false
}

// SetReferringPhysicianName gets a reference to the given string and assigns it to the ReferringPhysicianName field.
func (o *Ga4ghEnrollment) SetReferringPhysicianName(v string) {
	o.ReferringPhysicianName = &v
}

// GetReferringPhysicianContact returns the ReferringPhysicianContact field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetReferringPhysicianContact() string {
	if o == nil || o.ReferringPhysicianContact == nil {
		var ret string
		return ret
	}
	return *o.ReferringPhysicianContact
}

// GetReferringPhysicianContactOk returns a tuple with the ReferringPhysicianContact field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetReferringPhysicianContactOk() (string, bool) {
	if o == nil || o.ReferringPhysicianContact == nil {
		var ret string
		return ret, false
	}
	return *o.ReferringPhysicianContact, true
}

// HasReferringPhysicianContact returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasReferringPhysicianContact() bool {
	if o != nil && o.ReferringPhysicianContact != nil {
		return true
	}

	return false
}

// SetReferringPhysicianContact gets a reference to the given string and assigns it to the ReferringPhysicianContact field.
func (o *Ga4ghEnrollment) SetReferringPhysicianContact(v string) {
	o.ReferringPhysicianContact = &v
}

// GetSummaryOfIdRequest returns the SummaryOfIdRequest field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetSummaryOfIdRequest() string {
	if o == nil || o.SummaryOfIdRequest == nil {
		var ret string
		return ret
	}
	return *o.SummaryOfIdRequest
}

// GetSummaryOfIdRequestOk returns a tuple with the SummaryOfIdRequest field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetSummaryOfIdRequestOk() (string, bool) {
	if o == nil || o.SummaryOfIdRequest == nil {
		var ret string
		return ret, false
	}
	return *o.SummaryOfIdRequest, true
}

// HasSummaryOfIdRequest returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasSummaryOfIdRequest() bool {
	if o != nil && o.SummaryOfIdRequest != nil {
		return true
	}

	return false
}

// SetSummaryOfIdRequest gets a reference to the given string and assigns it to the SummaryOfIdRequest field.
func (o *Ga4ghEnrollment) SetSummaryOfIdRequest(v string) {
	o.SummaryOfIdRequest = &v
}

// GetTreatingCentreName returns the TreatingCentreName field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetTreatingCentreName() string {
	if o == nil || o.TreatingCentreName == nil {
		var ret string
		return ret
	}
	return *o.TreatingCentreName
}

// GetTreatingCentreNameOk returns a tuple with the TreatingCentreName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetTreatingCentreNameOk() (string, bool) {
	if o == nil || o.TreatingCentreName == nil {
		var ret string
		return ret, false
	}
	return *o.TreatingCentreName, true
}

// HasTreatingCentreName returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasTreatingCentreName() bool {
	if o != nil && o.TreatingCentreName != nil {
		return true
	}

	return false
}

// SetTreatingCentreName gets a reference to the given string and assigns it to the TreatingCentreName field.
func (o *Ga4ghEnrollment) SetTreatingCentreName(v string) {
	o.TreatingCentreName = &v
}

// GetTreatingCentreProvince returns the TreatingCentreProvince field if non-nil, zero value otherwise.
func (o *Ga4ghEnrollment) GetTreatingCentreProvince() string {
	if o == nil || o.TreatingCentreProvince == nil {
		var ret string
		return ret
	}
	return *o.TreatingCentreProvince
}

// GetTreatingCentreProvinceOk returns a tuple with the TreatingCentreProvince field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghEnrollment) GetTreatingCentreProvinceOk() (string, bool) {
	if o == nil || o.TreatingCentreProvince == nil {
		var ret string
		return ret, false
	}
	return *o.TreatingCentreProvince, true
}

// HasTreatingCentreProvince returns a boolean if a field has been set.
func (o *Ga4ghEnrollment) HasTreatingCentreProvince() bool {
	if o != nil && o.TreatingCentreProvince != nil {
		return true
	}

	return false
}

// SetTreatingCentreProvince gets a reference to the given string and assigns it to the TreatingCentreProvince field.
func (o *Ga4ghEnrollment) SetTreatingCentreProvince(v string) {
	o.TreatingCentreProvince = &v
}


func (o Ga4ghEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
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
	if o.EnrollmentInstitution != nil {
		toSerialize["enrollmentInstitution"] = o.EnrollmentInstitution
	}
	if o.EnrollmentApprovalDate != nil {
		toSerialize["enrollmentApprovalDate"] = o.EnrollmentApprovalDate
	}
	if o.CrossEnrollment != nil {
		toSerialize["crossEnrollment"] = o.CrossEnrollment
	}
	if o.OtherPersonalizedMedicineStudyName != nil {
		toSerialize["otherPersonalizedMedicineStudyName"] = o.OtherPersonalizedMedicineStudyName
	}
	if o.OtherPersonalizedMedicineStudyId != nil {
		toSerialize["otherPersonalizedMedicineStudyId"] = o.OtherPersonalizedMedicineStudyId
	}
	if o.AgeAtEnrollment != nil {
		toSerialize["ageAtEnrollment"] = o.AgeAtEnrollment
	}
	if o.EligibilityCategory != nil {
		toSerialize["eligibilityCategory"] = o.EligibilityCategory
	}
	if o.StatusAtEnrollment != nil {
		toSerialize["statusAtEnrollment"] = o.StatusAtEnrollment
	}
	if o.PrimaryOncologistName != nil {
		toSerialize["primaryOncologistName"] = o.PrimaryOncologistName
	}
	if o.PrimaryOncologistContact != nil {
		toSerialize["primaryOncologistContact"] = o.PrimaryOncologistContact
	}
	if o.ReferringPhysicianName != nil {
		toSerialize["referringPhysicianName"] = o.ReferringPhysicianName
	}
	if o.ReferringPhysicianContact != nil {
		toSerialize["referringPhysicianContact"] = o.ReferringPhysicianContact
	}
	if o.SummaryOfIdRequest != nil {
		toSerialize["summaryOfIdRequest"] = o.SummaryOfIdRequest
	}
	if o.TreatingCentreName != nil {
		toSerialize["treatingCentreName"] = o.TreatingCentreName
	}
	if o.TreatingCentreProvince != nil {
		toSerialize["treatingCentreProvince"] = o.TreatingCentreProvince
	}
	return json.Marshal(toSerialize)
}

