# Ga4ghSample

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
**SampleId** | Pointer to **string** |  | [optional] 
**DiagnosisId** | Pointer to **string** |  | [optional] 
**LocalBiobankId** | Pointer to **string** |  | [optional] 
**CollectionDate** | Pointer to **string** |  | [optional] 
**CollectionHospital** | Pointer to **string** |  | [optional] 
**SampleType** | Pointer to **string** |  | [optional] 
**TissueDiseaseState** | Pointer to **string** |  | [optional] 
**AnatomicSiteTheSampleObtainedFrom** | Pointer to **string** |  | [optional] 
**CancerType** | Pointer to **string** |  | [optional] 
**CancerSubtype** | Pointer to **string** |  | [optional] 
**PathologyReportId** | Pointer to **string** |  | [optional] 
**MorphologicalCode** | Pointer to **string** |  | [optional] 
**TopologicalCode** | Pointer to **string** |  | [optional] 
**ShippingDate** | Pointer to **string** |  | [optional] 
**ReceivedDate** | Pointer to **string** |  | [optional] 
**QualityControlPerformed** | Pointer to **string** |  | [optional] 
**EstimatedTumorContent** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **string** |  | [optional] 
**AssociatedBiobank** | Pointer to **string** |  | [optional] 
**OtherBiobank** | Pointer to **string** |  | [optional] 
**SopFollowed** | Pointer to **string** |  | [optional] 
**IfNotExplainAnyDeviation** | Pointer to **string** |  | [optional] 
**RecordingDate** | Pointer to **string** |  | [optional] 
**StartInterval** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghSample) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghSample) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghSample) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghSample) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDatasetId

`func (o *Ga4ghSample) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghSample) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghSample) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghSample) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetName

`func (o *Ga4ghSample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghSample) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghSample) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghSample) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghSample) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghSample) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghSample) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghSample) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCreated

`func (o *Ga4ghSample) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ga4ghSample) GetCreatedOk() (string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ga4ghSample) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ga4ghSample) SetCreated(v string)`

SetCreated gets a reference to the given string and assigns it to the Created field.

### GetUpdated

`func (o *Ga4ghSample) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Ga4ghSample) GetUpdatedOk() (string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Ga4ghSample) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Ga4ghSample) SetUpdated(v string)`

SetUpdated gets a reference to the given string and assigns it to the Updated field.

### GetPatientId

`func (o *Ga4ghSample) GetPatientId() string`

GetPatientId returns the PatientId field if non-nil, zero value otherwise.

### GetPatientIdOk

`func (o *Ga4ghSample) GetPatientIdOk() (string, bool)`

GetPatientIdOk returns a tuple with the PatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPatientId

`func (o *Ga4ghSample) HasPatientId() bool`

HasPatientId returns a boolean if a field has been set.

### SetPatientId

`func (o *Ga4ghSample) SetPatientId(v string)`

SetPatientId gets a reference to the given string and assigns it to the PatientId field.

### GetSampleId

`func (o *Ga4ghSample) GetSampleId() string`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Ga4ghSample) GetSampleIdOk() (string, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleId

`func (o *Ga4ghSample) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### SetSampleId

`func (o *Ga4ghSample) SetSampleId(v string)`

SetSampleId gets a reference to the given string and assigns it to the SampleId field.

### GetDiagnosisId

`func (o *Ga4ghSample) GetDiagnosisId() string`

GetDiagnosisId returns the DiagnosisId field if non-nil, zero value otherwise.

### GetDiagnosisIdOk

`func (o *Ga4ghSample) GetDiagnosisIdOk() (string, bool)`

GetDiagnosisIdOk returns a tuple with the DiagnosisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosisId

`func (o *Ga4ghSample) HasDiagnosisId() bool`

HasDiagnosisId returns a boolean if a field has been set.

### SetDiagnosisId

`func (o *Ga4ghSample) SetDiagnosisId(v string)`

SetDiagnosisId gets a reference to the given string and assigns it to the DiagnosisId field.

### GetLocalBiobankId

`func (o *Ga4ghSample) GetLocalBiobankId() string`

GetLocalBiobankId returns the LocalBiobankId field if non-nil, zero value otherwise.

### GetLocalBiobankIdOk

`func (o *Ga4ghSample) GetLocalBiobankIdOk() (string, bool)`

GetLocalBiobankIdOk returns a tuple with the LocalBiobankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocalBiobankId

`func (o *Ga4ghSample) HasLocalBiobankId() bool`

HasLocalBiobankId returns a boolean if a field has been set.

### SetLocalBiobankId

`func (o *Ga4ghSample) SetLocalBiobankId(v string)`

SetLocalBiobankId gets a reference to the given string and assigns it to the LocalBiobankId field.

### GetCollectionDate

`func (o *Ga4ghSample) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *Ga4ghSample) GetCollectionDateOk() (string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCollectionDate

`func (o *Ga4ghSample) HasCollectionDate() bool`

HasCollectionDate returns a boolean if a field has been set.

### SetCollectionDate

`func (o *Ga4ghSample) SetCollectionDate(v string)`

SetCollectionDate gets a reference to the given string and assigns it to the CollectionDate field.

### GetCollectionHospital

`func (o *Ga4ghSample) GetCollectionHospital() string`

GetCollectionHospital returns the CollectionHospital field if non-nil, zero value otherwise.

### GetCollectionHospitalOk

`func (o *Ga4ghSample) GetCollectionHospitalOk() (string, bool)`

GetCollectionHospitalOk returns a tuple with the CollectionHospital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCollectionHospital

`func (o *Ga4ghSample) HasCollectionHospital() bool`

HasCollectionHospital returns a boolean if a field has been set.

### SetCollectionHospital

`func (o *Ga4ghSample) SetCollectionHospital(v string)`

SetCollectionHospital gets a reference to the given string and assigns it to the CollectionHospital field.

### GetSampleType

`func (o *Ga4ghSample) GetSampleType() string`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *Ga4ghSample) GetSampleTypeOk() (string, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSampleType

`func (o *Ga4ghSample) HasSampleType() bool`

HasSampleType returns a boolean if a field has been set.

### SetSampleType

`func (o *Ga4ghSample) SetSampleType(v string)`

SetSampleType gets a reference to the given string and assigns it to the SampleType field.

### GetTissueDiseaseState

`func (o *Ga4ghSample) GetTissueDiseaseState() string`

GetTissueDiseaseState returns the TissueDiseaseState field if non-nil, zero value otherwise.

### GetTissueDiseaseStateOk

`func (o *Ga4ghSample) GetTissueDiseaseStateOk() (string, bool)`

GetTissueDiseaseStateOk returns a tuple with the TissueDiseaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTissueDiseaseState

`func (o *Ga4ghSample) HasTissueDiseaseState() bool`

HasTissueDiseaseState returns a boolean if a field has been set.

### SetTissueDiseaseState

`func (o *Ga4ghSample) SetTissueDiseaseState(v string)`

SetTissueDiseaseState gets a reference to the given string and assigns it to the TissueDiseaseState field.

### GetAnatomicSiteTheSampleObtainedFrom

`func (o *Ga4ghSample) GetAnatomicSiteTheSampleObtainedFrom() string`

GetAnatomicSiteTheSampleObtainedFrom returns the AnatomicSiteTheSampleObtainedFrom field if non-nil, zero value otherwise.

### GetAnatomicSiteTheSampleObtainedFromOk

`func (o *Ga4ghSample) GetAnatomicSiteTheSampleObtainedFromOk() (string, bool)`

GetAnatomicSiteTheSampleObtainedFromOk returns a tuple with the AnatomicSiteTheSampleObtainedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnatomicSiteTheSampleObtainedFrom

`func (o *Ga4ghSample) HasAnatomicSiteTheSampleObtainedFrom() bool`

HasAnatomicSiteTheSampleObtainedFrom returns a boolean if a field has been set.

### SetAnatomicSiteTheSampleObtainedFrom

`func (o *Ga4ghSample) SetAnatomicSiteTheSampleObtainedFrom(v string)`

SetAnatomicSiteTheSampleObtainedFrom gets a reference to the given string and assigns it to the AnatomicSiteTheSampleObtainedFrom field.

### GetCancerType

`func (o *Ga4ghSample) GetCancerType() string`

GetCancerType returns the CancerType field if non-nil, zero value otherwise.

### GetCancerTypeOk

`func (o *Ga4ghSample) GetCancerTypeOk() (string, bool)`

GetCancerTypeOk returns a tuple with the CancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancerType

`func (o *Ga4ghSample) HasCancerType() bool`

HasCancerType returns a boolean if a field has been set.

### SetCancerType

`func (o *Ga4ghSample) SetCancerType(v string)`

SetCancerType gets a reference to the given string and assigns it to the CancerType field.

### GetCancerSubtype

`func (o *Ga4ghSample) GetCancerSubtype() string`

GetCancerSubtype returns the CancerSubtype field if non-nil, zero value otherwise.

### GetCancerSubtypeOk

`func (o *Ga4ghSample) GetCancerSubtypeOk() (string, bool)`

GetCancerSubtypeOk returns a tuple with the CancerSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCancerSubtype

`func (o *Ga4ghSample) HasCancerSubtype() bool`

HasCancerSubtype returns a boolean if a field has been set.

### SetCancerSubtype

`func (o *Ga4ghSample) SetCancerSubtype(v string)`

SetCancerSubtype gets a reference to the given string and assigns it to the CancerSubtype field.

### GetPathologyReportId

`func (o *Ga4ghSample) GetPathologyReportId() string`

GetPathologyReportId returns the PathologyReportId field if non-nil, zero value otherwise.

### GetPathologyReportIdOk

`func (o *Ga4ghSample) GetPathologyReportIdOk() (string, bool)`

GetPathologyReportIdOk returns a tuple with the PathologyReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPathologyReportId

`func (o *Ga4ghSample) HasPathologyReportId() bool`

HasPathologyReportId returns a boolean if a field has been set.

### SetPathologyReportId

`func (o *Ga4ghSample) SetPathologyReportId(v string)`

SetPathologyReportId gets a reference to the given string and assigns it to the PathologyReportId field.

### GetMorphologicalCode

`func (o *Ga4ghSample) GetMorphologicalCode() string`

GetMorphologicalCode returns the MorphologicalCode field if non-nil, zero value otherwise.

### GetMorphologicalCodeOk

`func (o *Ga4ghSample) GetMorphologicalCodeOk() (string, bool)`

GetMorphologicalCodeOk returns a tuple with the MorphologicalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMorphologicalCode

`func (o *Ga4ghSample) HasMorphologicalCode() bool`

HasMorphologicalCode returns a boolean if a field has been set.

### SetMorphologicalCode

`func (o *Ga4ghSample) SetMorphologicalCode(v string)`

SetMorphologicalCode gets a reference to the given string and assigns it to the MorphologicalCode field.

### GetTopologicalCode

`func (o *Ga4ghSample) GetTopologicalCode() string`

GetTopologicalCode returns the TopologicalCode field if non-nil, zero value otherwise.

### GetTopologicalCodeOk

`func (o *Ga4ghSample) GetTopologicalCodeOk() (string, bool)`

GetTopologicalCodeOk returns a tuple with the TopologicalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTopologicalCode

`func (o *Ga4ghSample) HasTopologicalCode() bool`

HasTopologicalCode returns a boolean if a field has been set.

### SetTopologicalCode

`func (o *Ga4ghSample) SetTopologicalCode(v string)`

SetTopologicalCode gets a reference to the given string and assigns it to the TopologicalCode field.

### GetShippingDate

`func (o *Ga4ghSample) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *Ga4ghSample) GetShippingDateOk() (string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShippingDate

`func (o *Ga4ghSample) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### SetShippingDate

`func (o *Ga4ghSample) SetShippingDate(v string)`

SetShippingDate gets a reference to the given string and assigns it to the ShippingDate field.

### GetReceivedDate

`func (o *Ga4ghSample) GetReceivedDate() string`

GetReceivedDate returns the ReceivedDate field if non-nil, zero value otherwise.

### GetReceivedDateOk

`func (o *Ga4ghSample) GetReceivedDateOk() (string, bool)`

GetReceivedDateOk returns a tuple with the ReceivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedDate

`func (o *Ga4ghSample) HasReceivedDate() bool`

HasReceivedDate returns a boolean if a field has been set.

### SetReceivedDate

`func (o *Ga4ghSample) SetReceivedDate(v string)`

SetReceivedDate gets a reference to the given string and assigns it to the ReceivedDate field.

### GetQualityControlPerformed

`func (o *Ga4ghSample) GetQualityControlPerformed() string`

GetQualityControlPerformed returns the QualityControlPerformed field if non-nil, zero value otherwise.

### GetQualityControlPerformedOk

`func (o *Ga4ghSample) GetQualityControlPerformedOk() (string, bool)`

GetQualityControlPerformedOk returns a tuple with the QualityControlPerformed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityControlPerformed

`func (o *Ga4ghSample) HasQualityControlPerformed() bool`

HasQualityControlPerformed returns a boolean if a field has been set.

### SetQualityControlPerformed

`func (o *Ga4ghSample) SetQualityControlPerformed(v string)`

SetQualityControlPerformed gets a reference to the given string and assigns it to the QualityControlPerformed field.

### GetEstimatedTumorContent

`func (o *Ga4ghSample) GetEstimatedTumorContent() string`

GetEstimatedTumorContent returns the EstimatedTumorContent field if non-nil, zero value otherwise.

### GetEstimatedTumorContentOk

`func (o *Ga4ghSample) GetEstimatedTumorContentOk() (string, bool)`

GetEstimatedTumorContentOk returns a tuple with the EstimatedTumorContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEstimatedTumorContent

`func (o *Ga4ghSample) HasEstimatedTumorContent() bool`

HasEstimatedTumorContent returns a boolean if a field has been set.

### SetEstimatedTumorContent

`func (o *Ga4ghSample) SetEstimatedTumorContent(v string)`

SetEstimatedTumorContent gets a reference to the given string and assigns it to the EstimatedTumorContent field.

### GetQuantity

`func (o *Ga4ghSample) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Ga4ghSample) GetQuantityOk() (string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuantity

`func (o *Ga4ghSample) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantity

`func (o *Ga4ghSample) SetQuantity(v string)`

SetQuantity gets a reference to the given string and assigns it to the Quantity field.

### GetUnits

`func (o *Ga4ghSample) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Ga4ghSample) GetUnitsOk() (string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnits

`func (o *Ga4ghSample) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnits

`func (o *Ga4ghSample) SetUnits(v string)`

SetUnits gets a reference to the given string and assigns it to the Units field.

### GetAssociatedBiobank

`func (o *Ga4ghSample) GetAssociatedBiobank() string`

GetAssociatedBiobank returns the AssociatedBiobank field if non-nil, zero value otherwise.

### GetAssociatedBiobankOk

`func (o *Ga4ghSample) GetAssociatedBiobankOk() (string, bool)`

GetAssociatedBiobankOk returns a tuple with the AssociatedBiobank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssociatedBiobank

`func (o *Ga4ghSample) HasAssociatedBiobank() bool`

HasAssociatedBiobank returns a boolean if a field has been set.

### SetAssociatedBiobank

`func (o *Ga4ghSample) SetAssociatedBiobank(v string)`

SetAssociatedBiobank gets a reference to the given string and assigns it to the AssociatedBiobank field.

### GetOtherBiobank

`func (o *Ga4ghSample) GetOtherBiobank() string`

GetOtherBiobank returns the OtherBiobank field if non-nil, zero value otherwise.

### GetOtherBiobankOk

`func (o *Ga4ghSample) GetOtherBiobankOk() (string, bool)`

GetOtherBiobankOk returns a tuple with the OtherBiobank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherBiobank

`func (o *Ga4ghSample) HasOtherBiobank() bool`

HasOtherBiobank returns a boolean if a field has been set.

### SetOtherBiobank

`func (o *Ga4ghSample) SetOtherBiobank(v string)`

SetOtherBiobank gets a reference to the given string and assigns it to the OtherBiobank field.

### GetSopFollowed

`func (o *Ga4ghSample) GetSopFollowed() string`

GetSopFollowed returns the SopFollowed field if non-nil, zero value otherwise.

### GetSopFollowedOk

`func (o *Ga4ghSample) GetSopFollowedOk() (string, bool)`

GetSopFollowedOk returns a tuple with the SopFollowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSopFollowed

`func (o *Ga4ghSample) HasSopFollowed() bool`

HasSopFollowed returns a boolean if a field has been set.

### SetSopFollowed

`func (o *Ga4ghSample) SetSopFollowed(v string)`

SetSopFollowed gets a reference to the given string and assigns it to the SopFollowed field.

### GetIfNotExplainAnyDeviation

`func (o *Ga4ghSample) GetIfNotExplainAnyDeviation() string`

GetIfNotExplainAnyDeviation returns the IfNotExplainAnyDeviation field if non-nil, zero value otherwise.

### GetIfNotExplainAnyDeviationOk

`func (o *Ga4ghSample) GetIfNotExplainAnyDeviationOk() (string, bool)`

GetIfNotExplainAnyDeviationOk returns a tuple with the IfNotExplainAnyDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIfNotExplainAnyDeviation

`func (o *Ga4ghSample) HasIfNotExplainAnyDeviation() bool`

HasIfNotExplainAnyDeviation returns a boolean if a field has been set.

### SetIfNotExplainAnyDeviation

`func (o *Ga4ghSample) SetIfNotExplainAnyDeviation(v string)`

SetIfNotExplainAnyDeviation gets a reference to the given string and assigns it to the IfNotExplainAnyDeviation field.

### GetRecordingDate

`func (o *Ga4ghSample) GetRecordingDate() string`

GetRecordingDate returns the RecordingDate field if non-nil, zero value otherwise.

### GetRecordingDateOk

`func (o *Ga4ghSample) GetRecordingDateOk() (string, bool)`

GetRecordingDateOk returns a tuple with the RecordingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordingDate

`func (o *Ga4ghSample) HasRecordingDate() bool`

HasRecordingDate returns a boolean if a field has been set.

### SetRecordingDate

`func (o *Ga4ghSample) SetRecordingDate(v string)`

SetRecordingDate gets a reference to the given string and assigns it to the RecordingDate field.

### GetStartInterval

`func (o *Ga4ghSample) GetStartInterval() string`

GetStartInterval returns the StartInterval field if non-nil, zero value otherwise.

### GetStartIntervalOk

`func (o *Ga4ghSample) GetStartIntervalOk() (string, bool)`

GetStartIntervalOk returns a tuple with the StartInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartInterval

`func (o *Ga4ghSample) HasStartInterval() bool`

HasStartInterval returns a boolean if a field has been set.

### SetStartInterval

`func (o *Ga4ghSample) SetStartInterval(v string)`

SetStartInterval gets a reference to the given string and assigns it to the StartInterval field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


