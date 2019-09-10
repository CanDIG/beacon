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

type Ga4ghExternalIdentifier struct {
	// The source of the identifier, e.g. `Ensembl`.
	Database *string `json:"database,omitempty"`

	// The ID defined by the external database, e.g. `ENST00000000000`.
	Identifier *string `json:"identifier,omitempty"`

	// The version of the object or the database, e.g. `78`.
	Version *string `json:"version,omitempty"`

}

// GetDatabase returns the Database field if non-nil, zero value otherwise.
func (o *Ga4ghExternalIdentifier) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghExternalIdentifier) GetDatabaseOk() (string, bool) {
	if o == nil || o.Database == nil {
		var ret string
		return ret, false
	}
	return *o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Ga4ghExternalIdentifier) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *Ga4ghExternalIdentifier) SetDatabase(v string) {
	o.Database = &v
}

// GetIdentifier returns the Identifier field if non-nil, zero value otherwise.
func (o *Ga4ghExternalIdentifier) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghExternalIdentifier) GetIdentifierOk() (string, bool) {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret, false
	}
	return *o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Ga4ghExternalIdentifier) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Ga4ghExternalIdentifier) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *Ga4ghExternalIdentifier) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghExternalIdentifier) GetVersionOk() (string, bool) {
	if o == nil || o.Version == nil {
		var ret string
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Ga4ghExternalIdentifier) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Ga4ghExternalIdentifier) SetVersion(v string) {
	o.Version = &v
}


func (o Ga4ghExternalIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}


