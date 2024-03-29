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

type Ga4ghSearchQueryResponse struct {
	Patients *[]Ga4ghPatient `json:"patients,omitempty"`

	Enrollments *[]Ga4ghEnrollment `json:"enrollments,omitempty"`

	Consents *[]Ga4ghConsent `json:"consents,omitempty"`

	Diagnoses *[]Ga4ghDiagnosis `json:"diagnoses,omitempty"`

	Samples *[]Ga4ghSample `json:"samples,omitempty"`

	Treatments *[]Ga4ghTreatment `json:"treatments,omitempty"`

	Outcomes *[]Ga4ghOutcome `json:"outcomes,omitempty"`

	Complications *[]Ga4ghComplication `json:"complications,omitempty"`

	Tumourboards *[]Ga4ghTumourboard `json:"tumourboards,omitempty"`

	Variants *[]Ga4ghVariant `json:"variants,omitempty"`

	Slides *[]Ga4ghSlide `json:"slides,omitempty"`

	Studies *[]Ga4ghStudy `json:"studies,omitempty"`

	Labtests *[]Ga4ghLabtest `json:"labtests,omitempty"`

	Surgeries *[]Ga4ghSurgery `json:"surgeries,omitempty"`

	Chemotherapies *[]Ga4ghChemotherapy `json:"chemotherapies,omitempty"`

	Immunotherapies *[]Ga4ghImmunotherapy `json:"immunotherapies,omitempty"`

	Radiotherapies *[]Ga4ghRadiotherapy `json:"radiotherapies,omitempty"`

	Celltransplants *[]Ga4ghCelltransplant `json:"celltransplants,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetPatients returns the Patients field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetPatients() []Ga4ghPatient {
	if o == nil || o.Patients == nil {
		var ret []Ga4ghPatient
		return ret
	}
	return *o.Patients
}

// GetPatientsOk returns a tuple with the Patients field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetPatientsOk() ([]Ga4ghPatient, bool) {
	if o == nil || o.Patients == nil {
		var ret []Ga4ghPatient
		return ret, false
	}
	return *o.Patients, true
}

// HasPatients returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasPatients() bool {
	if o != nil && o.Patients != nil {
		return true
	}

	return false
}

// SetPatients gets a reference to the given []Ga4ghPatient and assigns it to the Patients field.
func (o *Ga4ghSearchQueryResponse) SetPatients(v []Ga4ghPatient) {
	o.Patients = &v
}

// GetEnrollments returns the Enrollments field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetEnrollments() []Ga4ghEnrollment {
	if o == nil || o.Enrollments == nil {
		var ret []Ga4ghEnrollment
		return ret
	}
	return *o.Enrollments
}

// GetEnrollmentsOk returns a tuple with the Enrollments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetEnrollmentsOk() ([]Ga4ghEnrollment, bool) {
	if o == nil || o.Enrollments == nil {
		var ret []Ga4ghEnrollment
		return ret, false
	}
	return *o.Enrollments, true
}

// HasEnrollments returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasEnrollments() bool {
	if o != nil && o.Enrollments != nil {
		return true
	}

	return false
}

// SetEnrollments gets a reference to the given []Ga4ghEnrollment and assigns it to the Enrollments field.
func (o *Ga4ghSearchQueryResponse) SetEnrollments(v []Ga4ghEnrollment) {
	o.Enrollments = &v
}

// GetConsents returns the Consents field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetConsents() []Ga4ghConsent {
	if o == nil || o.Consents == nil {
		var ret []Ga4ghConsent
		return ret
	}
	return *o.Consents
}

// GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetConsentsOk() ([]Ga4ghConsent, bool) {
	if o == nil || o.Consents == nil {
		var ret []Ga4ghConsent
		return ret, false
	}
	return *o.Consents, true
}

// HasConsents returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasConsents() bool {
	if o != nil && o.Consents != nil {
		return true
	}

	return false
}

// SetConsents gets a reference to the given []Ga4ghConsent and assigns it to the Consents field.
func (o *Ga4ghSearchQueryResponse) SetConsents(v []Ga4ghConsent) {
	o.Consents = &v
}

// GetDiagnoses returns the Diagnoses field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetDiagnoses() []Ga4ghDiagnosis {
	if o == nil || o.Diagnoses == nil {
		var ret []Ga4ghDiagnosis
		return ret
	}
	return *o.Diagnoses
}

// GetDiagnosesOk returns a tuple with the Diagnoses field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetDiagnosesOk() ([]Ga4ghDiagnosis, bool) {
	if o == nil || o.Diagnoses == nil {
		var ret []Ga4ghDiagnosis
		return ret, false
	}
	return *o.Diagnoses, true
}

// HasDiagnoses returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasDiagnoses() bool {
	if o != nil && o.Diagnoses != nil {
		return true
	}

	return false
}

// SetDiagnoses gets a reference to the given []Ga4ghDiagnosis and assigns it to the Diagnoses field.
func (o *Ga4ghSearchQueryResponse) SetDiagnoses(v []Ga4ghDiagnosis) {
	o.Diagnoses = &v
}

// GetSamples returns the Samples field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetSamples() []Ga4ghSample {
	if o == nil || o.Samples == nil {
		var ret []Ga4ghSample
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetSamplesOk() ([]Ga4ghSample, bool) {
	if o == nil || o.Samples == nil {
		var ret []Ga4ghSample
		return ret, false
	}
	return *o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasSamples() bool {
	if o != nil && o.Samples != nil {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []Ga4ghSample and assigns it to the Samples field.
func (o *Ga4ghSearchQueryResponse) SetSamples(v []Ga4ghSample) {
	o.Samples = &v
}

// GetTreatments returns the Treatments field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetTreatments() []Ga4ghTreatment {
	if o == nil || o.Treatments == nil {
		var ret []Ga4ghTreatment
		return ret
	}
	return *o.Treatments
}

// GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetTreatmentsOk() ([]Ga4ghTreatment, bool) {
	if o == nil || o.Treatments == nil {
		var ret []Ga4ghTreatment
		return ret, false
	}
	return *o.Treatments, true
}

// HasTreatments returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasTreatments() bool {
	if o != nil && o.Treatments != nil {
		return true
	}

	return false
}

// SetTreatments gets a reference to the given []Ga4ghTreatment and assigns it to the Treatments field.
func (o *Ga4ghSearchQueryResponse) SetTreatments(v []Ga4ghTreatment) {
	o.Treatments = &v
}

// GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetOutcomes() []Ga4ghOutcome {
	if o == nil || o.Outcomes == nil {
		var ret []Ga4ghOutcome
		return ret
	}
	return *o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetOutcomesOk() ([]Ga4ghOutcome, bool) {
	if o == nil || o.Outcomes == nil {
		var ret []Ga4ghOutcome
		return ret, false
	}
	return *o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given []Ga4ghOutcome and assigns it to the Outcomes field.
func (o *Ga4ghSearchQueryResponse) SetOutcomes(v []Ga4ghOutcome) {
	o.Outcomes = &v
}

// GetComplications returns the Complications field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetComplications() []Ga4ghComplication {
	if o == nil || o.Complications == nil {
		var ret []Ga4ghComplication
		return ret
	}
	return *o.Complications
}

// GetComplicationsOk returns a tuple with the Complications field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetComplicationsOk() ([]Ga4ghComplication, bool) {
	if o == nil || o.Complications == nil {
		var ret []Ga4ghComplication
		return ret, false
	}
	return *o.Complications, true
}

// HasComplications returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasComplications() bool {
	if o != nil && o.Complications != nil {
		return true
	}

	return false
}

// SetComplications gets a reference to the given []Ga4ghComplication and assigns it to the Complications field.
func (o *Ga4ghSearchQueryResponse) SetComplications(v []Ga4ghComplication) {
	o.Complications = &v
}

// GetTumourboards returns the Tumourboards field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetTumourboards() []Ga4ghTumourboard {
	if o == nil || o.Tumourboards == nil {
		var ret []Ga4ghTumourboard
		return ret
	}
	return *o.Tumourboards
}

// GetTumourboardsOk returns a tuple with the Tumourboards field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetTumourboardsOk() ([]Ga4ghTumourboard, bool) {
	if o == nil || o.Tumourboards == nil {
		var ret []Ga4ghTumourboard
		return ret, false
	}
	return *o.Tumourboards, true
}

// HasTumourboards returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasTumourboards() bool {
	if o != nil && o.Tumourboards != nil {
		return true
	}

	return false
}

// SetTumourboards gets a reference to the given []Ga4ghTumourboard and assigns it to the Tumourboards field.
func (o *Ga4ghSearchQueryResponse) SetTumourboards(v []Ga4ghTumourboard) {
	o.Tumourboards = &v
}

// GetVariants returns the Variants field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetVariants() []Ga4ghVariant {
	if o == nil || o.Variants == nil {
		var ret []Ga4ghVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetVariantsOk() ([]Ga4ghVariant, bool) {
	if o == nil || o.Variants == nil {
		var ret []Ga4ghVariant
		return ret, false
	}
	return *o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasVariants() bool {
	if o != nil && o.Variants != nil {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []Ga4ghVariant and assigns it to the Variants field.
func (o *Ga4ghSearchQueryResponse) SetVariants(v []Ga4ghVariant) {
	o.Variants = &v
}

// GetSlides returns the Slides field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetSlides() []Ga4ghSlide {
	if o == nil || o.Slides == nil {
		var ret []Ga4ghSlide
		return ret
	}
	return *o.Slides
}

// GetSlidesOk returns a tuple with the Slides field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetSlidesOk() ([]Ga4ghSlide, bool) {
	if o == nil || o.Slides == nil {
		var ret []Ga4ghSlide
		return ret, false
	}
	return *o.Slides, true
}

// HasSlides returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasSlides() bool {
	if o != nil && o.Slides != nil {
		return true
	}

	return false
}

// SetSlides gets a reference to the given []Ga4ghSlide and assigns it to the Slides field.
func (o *Ga4ghSearchQueryResponse) SetSlides(v []Ga4ghSlide) {
	o.Slides = &v
}

// GetStudies returns the Studies field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetStudies() []Ga4ghStudy {
	if o == nil || o.Studies == nil {
		var ret []Ga4ghStudy
		return ret
	}
	return *o.Studies
}

// GetStudiesOk returns a tuple with the Studies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetStudiesOk() ([]Ga4ghStudy, bool) {
	if o == nil || o.Studies == nil {
		var ret []Ga4ghStudy
		return ret, false
	}
	return *o.Studies, true
}

// HasStudies returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasStudies() bool {
	if o != nil && o.Studies != nil {
		return true
	}

	return false
}

// SetStudies gets a reference to the given []Ga4ghStudy and assigns it to the Studies field.
func (o *Ga4ghSearchQueryResponse) SetStudies(v []Ga4ghStudy) {
	o.Studies = &v
}

// GetLabtests returns the Labtests field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetLabtests() []Ga4ghLabtest {
	if o == nil || o.Labtests == nil {
		var ret []Ga4ghLabtest
		return ret
	}
	return *o.Labtests
}

// GetLabtestsOk returns a tuple with the Labtests field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetLabtestsOk() ([]Ga4ghLabtest, bool) {
	if o == nil || o.Labtests == nil {
		var ret []Ga4ghLabtest
		return ret, false
	}
	return *o.Labtests, true
}

// HasLabtests returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasLabtests() bool {
	if o != nil && o.Labtests != nil {
		return true
	}

	return false
}

// SetLabtests gets a reference to the given []Ga4ghLabtest and assigns it to the Labtests field.
func (o *Ga4ghSearchQueryResponse) SetLabtests(v []Ga4ghLabtest) {
	o.Labtests = &v
}

// GetSurgeries returns the Surgeries field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetSurgeries() []Ga4ghSurgery {
	if o == nil || o.Surgeries == nil {
		var ret []Ga4ghSurgery
		return ret
	}
	return *o.Surgeries
}

// GetSurgeriesOk returns a tuple with the Surgeries field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetSurgeriesOk() ([]Ga4ghSurgery, bool) {
	if o == nil || o.Surgeries == nil {
		var ret []Ga4ghSurgery
		return ret, false
	}
	return *o.Surgeries, true
}

// HasSurgeries returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasSurgeries() bool {
	if o != nil && o.Surgeries != nil {
		return true
	}

	return false
}

// SetSurgeries gets a reference to the given []Ga4ghSurgery and assigns it to the Surgeries field.
func (o *Ga4ghSearchQueryResponse) SetSurgeries(v []Ga4ghSurgery) {
	o.Surgeries = &v
}

// GetChemotherapies returns the Chemotherapies field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetChemotherapies() []Ga4ghChemotherapy {
	if o == nil || o.Chemotherapies == nil {
		var ret []Ga4ghChemotherapy
		return ret
	}
	return *o.Chemotherapies
}

// GetChemotherapiesOk returns a tuple with the Chemotherapies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetChemotherapiesOk() ([]Ga4ghChemotherapy, bool) {
	if o == nil || o.Chemotherapies == nil {
		var ret []Ga4ghChemotherapy
		return ret, false
	}
	return *o.Chemotherapies, true
}

// HasChemotherapies returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasChemotherapies() bool {
	if o != nil && o.Chemotherapies != nil {
		return true
	}

	return false
}

// SetChemotherapies gets a reference to the given []Ga4ghChemotherapy and assigns it to the Chemotherapies field.
func (o *Ga4ghSearchQueryResponse) SetChemotherapies(v []Ga4ghChemotherapy) {
	o.Chemotherapies = &v
}

// GetImmunotherapies returns the Immunotherapies field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetImmunotherapies() []Ga4ghImmunotherapy {
	if o == nil || o.Immunotherapies == nil {
		var ret []Ga4ghImmunotherapy
		return ret
	}
	return *o.Immunotherapies
}

// GetImmunotherapiesOk returns a tuple with the Immunotherapies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetImmunotherapiesOk() ([]Ga4ghImmunotherapy, bool) {
	if o == nil || o.Immunotherapies == nil {
		var ret []Ga4ghImmunotherapy
		return ret, false
	}
	return *o.Immunotherapies, true
}

// HasImmunotherapies returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasImmunotherapies() bool {
	if o != nil && o.Immunotherapies != nil {
		return true
	}

	return false
}

// SetImmunotherapies gets a reference to the given []Ga4ghImmunotherapy and assigns it to the Immunotherapies field.
func (o *Ga4ghSearchQueryResponse) SetImmunotherapies(v []Ga4ghImmunotherapy) {
	o.Immunotherapies = &v
}

// GetRadiotherapies returns the Radiotherapies field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetRadiotherapies() []Ga4ghRadiotherapy {
	if o == nil || o.Radiotherapies == nil {
		var ret []Ga4ghRadiotherapy
		return ret
	}
	return *o.Radiotherapies
}

// GetRadiotherapiesOk returns a tuple with the Radiotherapies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetRadiotherapiesOk() ([]Ga4ghRadiotherapy, bool) {
	if o == nil || o.Radiotherapies == nil {
		var ret []Ga4ghRadiotherapy
		return ret, false
	}
	return *o.Radiotherapies, true
}

// HasRadiotherapies returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasRadiotherapies() bool {
	if o != nil && o.Radiotherapies != nil {
		return true
	}

	return false
}

// SetRadiotherapies gets a reference to the given []Ga4ghRadiotherapy and assigns it to the Radiotherapies field.
func (o *Ga4ghSearchQueryResponse) SetRadiotherapies(v []Ga4ghRadiotherapy) {
	o.Radiotherapies = &v
}

// GetCelltransplants returns the Celltransplants field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetCelltransplants() []Ga4ghCelltransplant {
	if o == nil || o.Celltransplants == nil {
		var ret []Ga4ghCelltransplant
		return ret
	}
	return *o.Celltransplants
}

// GetCelltransplantsOk returns a tuple with the Celltransplants field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetCelltransplantsOk() ([]Ga4ghCelltransplant, bool) {
	if o == nil || o.Celltransplants == nil {
		var ret []Ga4ghCelltransplant
		return ret, false
	}
	return *o.Celltransplants, true
}

// HasCelltransplants returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasCelltransplants() bool {
	if o != nil && o.Celltransplants != nil {
		return true
	}

	return false
}

// SetCelltransplants gets a reference to the given []Ga4ghCelltransplant and assigns it to the Celltransplants field.
func (o *Ga4ghSearchQueryResponse) SetCelltransplants(v []Ga4ghCelltransplant) {
	o.Celltransplants = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchQueryResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchQueryResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchQueryResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchQueryResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Patients != nil {
		toSerialize["patients"] = o.Patients
	}
	if o.Enrollments != nil {
		toSerialize["enrollments"] = o.Enrollments
	}
	if o.Consents != nil {
		toSerialize["consents"] = o.Consents
	}
	if o.Diagnoses != nil {
		toSerialize["diagnoses"] = o.Diagnoses
	}
	if o.Samples != nil {
		toSerialize["samples"] = o.Samples
	}
	if o.Treatments != nil {
		toSerialize["treatments"] = o.Treatments
	}
	if o.Outcomes != nil {
		toSerialize["outcomes"] = o.Outcomes
	}
	if o.Complications != nil {
		toSerialize["complications"] = o.Complications
	}
	if o.Tumourboards != nil {
		toSerialize["tumourboards"] = o.Tumourboards
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.Slides != nil {
		toSerialize["slides"] = o.Slides
	}
	if o.Studies != nil {
		toSerialize["studies"] = o.Studies
	}
	if o.Labtests != nil {
		toSerialize["labtests"] = o.Labtests
	}
	if o.Surgeries != nil {
		toSerialize["surgeries"] = o.Surgeries
	}
	if o.Chemotherapies != nil {
		toSerialize["chemotherapies"] = o.Chemotherapies
	}
	if o.Immunotherapies != nil {
		toSerialize["immunotherapies"] = o.Immunotherapies
	}
	if o.Radiotherapies != nil {
		toSerialize["radiotherapies"] = o.Radiotherapies
	}
	if o.Celltransplants != nil {
		toSerialize["celltransplants"] = o.Celltransplants
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


