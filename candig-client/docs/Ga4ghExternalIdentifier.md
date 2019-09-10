# Ga4ghExternalIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **string** | The source of the identifier, e.g. &#x60;Ensembl&#x60;. | [optional] 
**Identifier** | Pointer to **string** | The ID defined by the external database, e.g. &#x60;ENST00000000000&#x60;. | [optional] 
**Version** | Pointer to **string** | The version of the object or the database, e.g. &#x60;78&#x60;. | [optional] 

## Methods

### GetDatabase

`func (o *Ga4ghExternalIdentifier) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Ga4ghExternalIdentifier) GetDatabaseOk() (string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatabase

`func (o *Ga4ghExternalIdentifier) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabase

`func (o *Ga4ghExternalIdentifier) SetDatabase(v string)`

SetDatabase gets a reference to the given string and assigns it to the Database field.

### GetIdentifier

`func (o *Ga4ghExternalIdentifier) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Ga4ghExternalIdentifier) GetIdentifierOk() (string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifier

`func (o *Ga4ghExternalIdentifier) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifier

`func (o *Ga4ghExternalIdentifier) SetIdentifier(v string)`

SetIdentifier gets a reference to the given string and assigns it to the Identifier field.

### GetVersion

`func (o *Ga4ghExternalIdentifier) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Ga4ghExternalIdentifier) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *Ga4ghExternalIdentifier) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *Ga4ghExternalIdentifier) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


