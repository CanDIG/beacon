# Ga4ghOntologyTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermId** | Pointer to **string** | Ontology term identifier - the CURIE for an ontology term. It differs from the standard GA4GH schema&#39;s :ref:&#x60;id &lt;apidesign_object_ids&gt;&#x60; in that it is a CURIE pointing to an information resource outside of the scope of the schema or its resource implementation. | [optional] 
**Term** | Pointer to **string** | Ontology term - the label of the ontology term the termId is pointing to. | [optional] 

## Methods

### GetTermId

`func (o *Ga4ghOntologyTerm) GetTermId() string`

GetTermId returns the TermId field if non-nil, zero value otherwise.

### GetTermIdOk

`func (o *Ga4ghOntologyTerm) GetTermIdOk() (string, bool)`

GetTermIdOk returns a tuple with the TermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTermId

`func (o *Ga4ghOntologyTerm) HasTermId() bool`

HasTermId returns a boolean if a field has been set.

### SetTermId

`func (o *Ga4ghOntologyTerm) SetTermId(v string)`

SetTermId gets a reference to the given string and assigns it to the TermId field.

### GetTerm

`func (o *Ga4ghOntologyTerm) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *Ga4ghOntologyTerm) GetTermOk() (string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTerm

`func (o *Ga4ghOntologyTerm) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTerm

`func (o *Ga4ghOntologyTerm) SetTerm(v string)`

SetTerm gets a reference to the given string and assigns it to the Term field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


