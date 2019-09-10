# Ga4ghExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the experiment. | [optional] 
**Description** | Pointer to **string** | A description of the experiment. | [optional] 
**MessageCreateTime** | Pointer to **string** |  | [optional] 
**MessageUpdateTime** | Pointer to **string** |  | [optional] 
**RunTime** | Pointer to **string** |  | [optional] 
**Molecule** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Selection** | Pointer to **string** |  | [optional] 
**Library** | Pointer to **string** | The name of the library used as part of this experiment. | [optional] 
**LibraryLayout** | Pointer to **string** | The configuration of sequenced reads. (e.g. Single or Paired). | [optional] 
**InstrumentModel** | Pointer to **string** | The instrument model used as part of this experiment. This maps to sequencing technology in BAM. | [optional] 
**InstrumentDataFile** | Pointer to **string** |  | [optional] 
**SequencingCenter** | Pointer to **string** | The sequencing center used as part of this experiment. | [optional] 
**PlatformUnit** | Pointer to **string** | The platform unit used as part of this experiment. This is a flowcell-barcode or slide unique identifier. | [optional] 
**DatasetId** | Pointer to **string** | ### &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; ### # PROFYLE MODIFICATION BEGIN ### &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; ### The ID of the dataset this Experiment belongs to. | [optional] 
**BiosampleId** | Pointer to **string** |  | [optional] 
**DnaLibraryConstructionMethod** | Pointer to **string** |  | [optional] 
**WgsSequencingCompletionDate** | Pointer to **string** |  | [optional] 
**RnaLibraryConstructionMethod** | Pointer to **string** |  | [optional] 
**RnaSequencingCompletionDate** | Pointer to **string** |  | [optional] 
**PanelCompletionDate** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Ga4ghExperiment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ga4ghExperiment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ga4ghExperiment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ga4ghExperiment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Ga4ghExperiment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ga4ghExperiment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Ga4ghExperiment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Ga4ghExperiment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Ga4ghExperiment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Ga4ghExperiment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Ga4ghExperiment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Ga4ghExperiment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetMessageCreateTime

`func (o *Ga4ghExperiment) GetMessageCreateTime() string`

GetMessageCreateTime returns the MessageCreateTime field if non-nil, zero value otherwise.

### GetMessageCreateTimeOk

`func (o *Ga4ghExperiment) GetMessageCreateTimeOk() (string, bool)`

GetMessageCreateTimeOk returns a tuple with the MessageCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageCreateTime

`func (o *Ga4ghExperiment) HasMessageCreateTime() bool`

HasMessageCreateTime returns a boolean if a field has been set.

### SetMessageCreateTime

`func (o *Ga4ghExperiment) SetMessageCreateTime(v string)`

SetMessageCreateTime gets a reference to the given string and assigns it to the MessageCreateTime field.

### GetMessageUpdateTime

`func (o *Ga4ghExperiment) GetMessageUpdateTime() string`

GetMessageUpdateTime returns the MessageUpdateTime field if non-nil, zero value otherwise.

### GetMessageUpdateTimeOk

`func (o *Ga4ghExperiment) GetMessageUpdateTimeOk() (string, bool)`

GetMessageUpdateTimeOk returns a tuple with the MessageUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageUpdateTime

`func (o *Ga4ghExperiment) HasMessageUpdateTime() bool`

HasMessageUpdateTime returns a boolean if a field has been set.

### SetMessageUpdateTime

`func (o *Ga4ghExperiment) SetMessageUpdateTime(v string)`

SetMessageUpdateTime gets a reference to the given string and assigns it to the MessageUpdateTime field.

### GetRunTime

`func (o *Ga4ghExperiment) GetRunTime() string`

GetRunTime returns the RunTime field if non-nil, zero value otherwise.

### GetRunTimeOk

`func (o *Ga4ghExperiment) GetRunTimeOk() (string, bool)`

GetRunTimeOk returns a tuple with the RunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRunTime

`func (o *Ga4ghExperiment) HasRunTime() bool`

HasRunTime returns a boolean if a field has been set.

### SetRunTime

`func (o *Ga4ghExperiment) SetRunTime(v string)`

SetRunTime gets a reference to the given string and assigns it to the RunTime field.

### GetMolecule

`func (o *Ga4ghExperiment) GetMolecule() string`

GetMolecule returns the Molecule field if non-nil, zero value otherwise.

### GetMoleculeOk

`func (o *Ga4ghExperiment) GetMoleculeOk() (string, bool)`

GetMoleculeOk returns a tuple with the Molecule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMolecule

`func (o *Ga4ghExperiment) HasMolecule() bool`

HasMolecule returns a boolean if a field has been set.

### SetMolecule

`func (o *Ga4ghExperiment) SetMolecule(v string)`

SetMolecule gets a reference to the given string and assigns it to the Molecule field.

### GetStrategy

`func (o *Ga4ghExperiment) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Ga4ghExperiment) GetStrategyOk() (string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStrategy

`func (o *Ga4ghExperiment) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### SetStrategy

`func (o *Ga4ghExperiment) SetStrategy(v string)`

SetStrategy gets a reference to the given string and assigns it to the Strategy field.

### GetSelection

`func (o *Ga4ghExperiment) GetSelection() string`

GetSelection returns the Selection field if non-nil, zero value otherwise.

### GetSelectionOk

`func (o *Ga4ghExperiment) GetSelectionOk() (string, bool)`

GetSelectionOk returns a tuple with the Selection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelection

`func (o *Ga4ghExperiment) HasSelection() bool`

HasSelection returns a boolean if a field has been set.

### SetSelection

`func (o *Ga4ghExperiment) SetSelection(v string)`

SetSelection gets a reference to the given string and assigns it to the Selection field.

### GetLibrary

`func (o *Ga4ghExperiment) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *Ga4ghExperiment) GetLibraryOk() (string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLibrary

`func (o *Ga4ghExperiment) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### SetLibrary

`func (o *Ga4ghExperiment) SetLibrary(v string)`

SetLibrary gets a reference to the given string and assigns it to the Library field.

### GetLibraryLayout

`func (o *Ga4ghExperiment) GetLibraryLayout() string`

GetLibraryLayout returns the LibraryLayout field if non-nil, zero value otherwise.

### GetLibraryLayoutOk

`func (o *Ga4ghExperiment) GetLibraryLayoutOk() (string, bool)`

GetLibraryLayoutOk returns a tuple with the LibraryLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLibraryLayout

`func (o *Ga4ghExperiment) HasLibraryLayout() bool`

HasLibraryLayout returns a boolean if a field has been set.

### SetLibraryLayout

`func (o *Ga4ghExperiment) SetLibraryLayout(v string)`

SetLibraryLayout gets a reference to the given string and assigns it to the LibraryLayout field.

### GetInstrumentModel

`func (o *Ga4ghExperiment) GetInstrumentModel() string`

GetInstrumentModel returns the InstrumentModel field if non-nil, zero value otherwise.

### GetInstrumentModelOk

`func (o *Ga4ghExperiment) GetInstrumentModelOk() (string, bool)`

GetInstrumentModelOk returns a tuple with the InstrumentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstrumentModel

`func (o *Ga4ghExperiment) HasInstrumentModel() bool`

HasInstrumentModel returns a boolean if a field has been set.

### SetInstrumentModel

`func (o *Ga4ghExperiment) SetInstrumentModel(v string)`

SetInstrumentModel gets a reference to the given string and assigns it to the InstrumentModel field.

### GetInstrumentDataFile

`func (o *Ga4ghExperiment) GetInstrumentDataFile() string`

GetInstrumentDataFile returns the InstrumentDataFile field if non-nil, zero value otherwise.

### GetInstrumentDataFileOk

`func (o *Ga4ghExperiment) GetInstrumentDataFileOk() (string, bool)`

GetInstrumentDataFileOk returns a tuple with the InstrumentDataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstrumentDataFile

`func (o *Ga4ghExperiment) HasInstrumentDataFile() bool`

HasInstrumentDataFile returns a boolean if a field has been set.

### SetInstrumentDataFile

`func (o *Ga4ghExperiment) SetInstrumentDataFile(v string)`

SetInstrumentDataFile gets a reference to the given string and assigns it to the InstrumentDataFile field.

### GetSequencingCenter

`func (o *Ga4ghExperiment) GetSequencingCenter() string`

GetSequencingCenter returns the SequencingCenter field if non-nil, zero value otherwise.

### GetSequencingCenterOk

`func (o *Ga4ghExperiment) GetSequencingCenterOk() (string, bool)`

GetSequencingCenterOk returns a tuple with the SequencingCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSequencingCenter

`func (o *Ga4ghExperiment) HasSequencingCenter() bool`

HasSequencingCenter returns a boolean if a field has been set.

### SetSequencingCenter

`func (o *Ga4ghExperiment) SetSequencingCenter(v string)`

SetSequencingCenter gets a reference to the given string and assigns it to the SequencingCenter field.

### GetPlatformUnit

`func (o *Ga4ghExperiment) GetPlatformUnit() string`

GetPlatformUnit returns the PlatformUnit field if non-nil, zero value otherwise.

### GetPlatformUnitOk

`func (o *Ga4ghExperiment) GetPlatformUnitOk() (string, bool)`

GetPlatformUnitOk returns a tuple with the PlatformUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformUnit

`func (o *Ga4ghExperiment) HasPlatformUnit() bool`

HasPlatformUnit returns a boolean if a field has been set.

### SetPlatformUnit

`func (o *Ga4ghExperiment) SetPlatformUnit(v string)`

SetPlatformUnit gets a reference to the given string and assigns it to the PlatformUnit field.

### GetDatasetId

`func (o *Ga4ghExperiment) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Ga4ghExperiment) GetDatasetIdOk() (string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetId

`func (o *Ga4ghExperiment) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetId

`func (o *Ga4ghExperiment) SetDatasetId(v string)`

SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.

### GetBiosampleId

`func (o *Ga4ghExperiment) GetBiosampleId() string`

GetBiosampleId returns the BiosampleId field if non-nil, zero value otherwise.

### GetBiosampleIdOk

`func (o *Ga4ghExperiment) GetBiosampleIdOk() (string, bool)`

GetBiosampleIdOk returns a tuple with the BiosampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBiosampleId

`func (o *Ga4ghExperiment) HasBiosampleId() bool`

HasBiosampleId returns a boolean if a field has been set.

### SetBiosampleId

`func (o *Ga4ghExperiment) SetBiosampleId(v string)`

SetBiosampleId gets a reference to the given string and assigns it to the BiosampleId field.

### GetDnaLibraryConstructionMethod

`func (o *Ga4ghExperiment) GetDnaLibraryConstructionMethod() string`

GetDnaLibraryConstructionMethod returns the DnaLibraryConstructionMethod field if non-nil, zero value otherwise.

### GetDnaLibraryConstructionMethodOk

`func (o *Ga4ghExperiment) GetDnaLibraryConstructionMethodOk() (string, bool)`

GetDnaLibraryConstructionMethodOk returns a tuple with the DnaLibraryConstructionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDnaLibraryConstructionMethod

`func (o *Ga4ghExperiment) HasDnaLibraryConstructionMethod() bool`

HasDnaLibraryConstructionMethod returns a boolean if a field has been set.

### SetDnaLibraryConstructionMethod

`func (o *Ga4ghExperiment) SetDnaLibraryConstructionMethod(v string)`

SetDnaLibraryConstructionMethod gets a reference to the given string and assigns it to the DnaLibraryConstructionMethod field.

### GetWgsSequencingCompletionDate

`func (o *Ga4ghExperiment) GetWgsSequencingCompletionDate() string`

GetWgsSequencingCompletionDate returns the WgsSequencingCompletionDate field if non-nil, zero value otherwise.

### GetWgsSequencingCompletionDateOk

`func (o *Ga4ghExperiment) GetWgsSequencingCompletionDateOk() (string, bool)`

GetWgsSequencingCompletionDateOk returns a tuple with the WgsSequencingCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWgsSequencingCompletionDate

`func (o *Ga4ghExperiment) HasWgsSequencingCompletionDate() bool`

HasWgsSequencingCompletionDate returns a boolean if a field has been set.

### SetWgsSequencingCompletionDate

`func (o *Ga4ghExperiment) SetWgsSequencingCompletionDate(v string)`

SetWgsSequencingCompletionDate gets a reference to the given string and assigns it to the WgsSequencingCompletionDate field.

### GetRnaLibraryConstructionMethod

`func (o *Ga4ghExperiment) GetRnaLibraryConstructionMethod() string`

GetRnaLibraryConstructionMethod returns the RnaLibraryConstructionMethod field if non-nil, zero value otherwise.

### GetRnaLibraryConstructionMethodOk

`func (o *Ga4ghExperiment) GetRnaLibraryConstructionMethodOk() (string, bool)`

GetRnaLibraryConstructionMethodOk returns a tuple with the RnaLibraryConstructionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaLibraryConstructionMethod

`func (o *Ga4ghExperiment) HasRnaLibraryConstructionMethod() bool`

HasRnaLibraryConstructionMethod returns a boolean if a field has been set.

### SetRnaLibraryConstructionMethod

`func (o *Ga4ghExperiment) SetRnaLibraryConstructionMethod(v string)`

SetRnaLibraryConstructionMethod gets a reference to the given string and assigns it to the RnaLibraryConstructionMethod field.

### GetRnaSequencingCompletionDate

`func (o *Ga4ghExperiment) GetRnaSequencingCompletionDate() string`

GetRnaSequencingCompletionDate returns the RnaSequencingCompletionDate field if non-nil, zero value otherwise.

### GetRnaSequencingCompletionDateOk

`func (o *Ga4ghExperiment) GetRnaSequencingCompletionDateOk() (string, bool)`

GetRnaSequencingCompletionDateOk returns a tuple with the RnaSequencingCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRnaSequencingCompletionDate

`func (o *Ga4ghExperiment) HasRnaSequencingCompletionDate() bool`

HasRnaSequencingCompletionDate returns a boolean if a field has been set.

### SetRnaSequencingCompletionDate

`func (o *Ga4ghExperiment) SetRnaSequencingCompletionDate(v string)`

SetRnaSequencingCompletionDate gets a reference to the given string and assigns it to the RnaSequencingCompletionDate field.

### GetPanelCompletionDate

`func (o *Ga4ghExperiment) GetPanelCompletionDate() string`

GetPanelCompletionDate returns the PanelCompletionDate field if non-nil, zero value otherwise.

### GetPanelCompletionDateOk

`func (o *Ga4ghExperiment) GetPanelCompletionDateOk() (string, bool)`

GetPanelCompletionDateOk returns a tuple with the PanelCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPanelCompletionDate

`func (o *Ga4ghExperiment) HasPanelCompletionDate() bool`

HasPanelCompletionDate returns a boolean if a field has been set.

### SetPanelCompletionDate

`func (o *Ga4ghExperiment) SetPanelCompletionDate(v string)`

SetPanelCompletionDate gets a reference to the given string and assigns it to the PanelCompletionDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


