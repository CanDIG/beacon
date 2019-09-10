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

// A `Reference` is a canonical assembled contig, intended to act as a reference coordinate space for other genomic annotations. A single `Reference` might represent the human chromosome 1, for instance.  `Reference` s are designed to be immutable.
type Ga4ghReference struct {
	// The reference ID. Unique within the repository.
	Id *string `json:"id,omitempty"`

	// The length of this reference's sequence.
	Length *string `json:"length,omitempty"`

	// The MD5 checksum uniquely representing this `Reference` as a lower-case hexadecimal string, calculated as the MD5 of the upper-case sequence excluding all whitespace characters (this is equivalent to SQ:M5 in SAM).
	Md5checksum *string `json:"md5checksum,omitempty"`

	// The unique name of this reference within the Reference Set (e.g. '22').
	Name *string `json:"name,omitempty"`

	// The URI from which the sequence was obtained. Specifies a FASTA format file/string with one name, sequence pair. In most cases, clients should call the `getReferenceBases()` method to obtain sequence bases for a `Reference` instead of attempting to retrieve this URI.
	SourceUri *string `json:"source_uri,omitempty"`

	// All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) which must include a version number, e.g. `GCF_000001405.26`.
	SourceAccessions *[]string `json:"source_accessions,omitempty"`

	// A sequence X is said to be derived from source sequence Y, if X and Y are of the same length and the per-base sequence divergence at A/C/G/T bases is sufficiently small. Two sequences derived from the same official sequence share the same coordinates and annotations, and can be replaced with the official sequence for certain use cases.
	IsDerived *bool `json:"is_derived,omitempty"`

	// The `sourceDivergence` is the fraction of non-indel bases that do not match the reference this message was derived from.
	SourceDivergence *float32 `json:"source_divergence,omitempty"`

	Species *Ga4ghOntologyTerm `json:"species,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ga4ghReference) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ga4ghReference) SetId(v string) {
	o.Id = &v
}

// GetLength returns the Length field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetLength() string {
	if o == nil || o.Length == nil {
		var ret string
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetLengthOk() (string, bool) {
	if o == nil || o.Length == nil {
		var ret string
		return ret, false
	}
	return *o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Ga4ghReference) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given string and assigns it to the Length field.
func (o *Ga4ghReference) SetLength(v string) {
	o.Length = &v
}

// GetMd5checksum returns the Md5checksum field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetMd5checksum() string {
	if o == nil || o.Md5checksum == nil {
		var ret string
		return ret
	}
	return *o.Md5checksum
}

// GetMd5checksumOk returns a tuple with the Md5checksum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetMd5checksumOk() (string, bool) {
	if o == nil || o.Md5checksum == nil {
		var ret string
		return ret, false
	}
	return *o.Md5checksum, true
}

// HasMd5checksum returns a boolean if a field has been set.
func (o *Ga4ghReference) HasMd5checksum() bool {
	if o != nil && o.Md5checksum != nil {
		return true
	}

	return false
}

// SetMd5checksum gets a reference to the given string and assigns it to the Md5checksum field.
func (o *Ga4ghReference) SetMd5checksum(v string) {
	o.Md5checksum = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ga4ghReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ga4ghReference) SetName(v string) {
	o.Name = &v
}

// GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetSourceUri() string {
	if o == nil || o.SourceUri == nil {
		var ret string
		return ret
	}
	return *o.SourceUri
}

// GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetSourceUriOk() (string, bool) {
	if o == nil || o.SourceUri == nil {
		var ret string
		return ret, false
	}
	return *o.SourceUri, true
}

// HasSourceUri returns a boolean if a field has been set.
func (o *Ga4ghReference) HasSourceUri() bool {
	if o != nil && o.SourceUri != nil {
		return true
	}

	return false
}

// SetSourceUri gets a reference to the given string and assigns it to the SourceUri field.
func (o *Ga4ghReference) SetSourceUri(v string) {
	o.SourceUri = &v
}

// GetSourceAccessions returns the SourceAccessions field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetSourceAccessions() []string {
	if o == nil || o.SourceAccessions == nil {
		var ret []string
		return ret
	}
	return *o.SourceAccessions
}

// GetSourceAccessionsOk returns a tuple with the SourceAccessions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetSourceAccessionsOk() ([]string, bool) {
	if o == nil || o.SourceAccessions == nil {
		var ret []string
		return ret, false
	}
	return *o.SourceAccessions, true
}

// HasSourceAccessions returns a boolean if a field has been set.
func (o *Ga4ghReference) HasSourceAccessions() bool {
	if o != nil && o.SourceAccessions != nil {
		return true
	}

	return false
}

// SetSourceAccessions gets a reference to the given []string and assigns it to the SourceAccessions field.
func (o *Ga4ghReference) SetSourceAccessions(v []string) {
	o.SourceAccessions = &v
}

// GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetIsDerived() bool {
	if o == nil || o.IsDerived == nil {
		var ret bool
		return ret
	}
	return *o.IsDerived
}

// GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetIsDerivedOk() (bool, bool) {
	if o == nil || o.IsDerived == nil {
		var ret bool
		return ret, false
	}
	return *o.IsDerived, true
}

// HasIsDerived returns a boolean if a field has been set.
func (o *Ga4ghReference) HasIsDerived() bool {
	if o != nil && o.IsDerived != nil {
		return true
	}

	return false
}

// SetIsDerived gets a reference to the given bool and assigns it to the IsDerived field.
func (o *Ga4ghReference) SetIsDerived(v bool) {
	o.IsDerived = &v
}

// GetSourceDivergence returns the SourceDivergence field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetSourceDivergence() float32 {
	if o == nil || o.SourceDivergence == nil {
		var ret float32
		return ret
	}
	return *o.SourceDivergence
}

// GetSourceDivergenceOk returns a tuple with the SourceDivergence field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetSourceDivergenceOk() (float32, bool) {
	if o == nil || o.SourceDivergence == nil {
		var ret float32
		return ret, false
	}
	return *o.SourceDivergence, true
}

// HasSourceDivergence returns a boolean if a field has been set.
func (o *Ga4ghReference) HasSourceDivergence() bool {
	if o != nil && o.SourceDivergence != nil {
		return true
	}

	return false
}

// SetSourceDivergence gets a reference to the given float32 and assigns it to the SourceDivergence field.
func (o *Ga4ghReference) SetSourceDivergence(v float32) {
	o.SourceDivergence = &v
}

// GetSpecies returns the Species field if non-nil, zero value otherwise.
func (o *Ga4ghReference) GetSpecies() Ga4ghOntologyTerm {
	if o == nil || o.Species == nil {
		var ret Ga4ghOntologyTerm
		return ret
	}
	return *o.Species
}

// GetSpeciesOk returns a tuple with the Species field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghReference) GetSpeciesOk() (Ga4ghOntologyTerm, bool) {
	if o == nil || o.Species == nil {
		var ret Ga4ghOntologyTerm
		return ret, false
	}
	return *o.Species, true
}

// HasSpecies returns a boolean if a field has been set.
func (o *Ga4ghReference) HasSpecies() bool {
	if o != nil && o.Species != nil {
		return true
	}

	return false
}

// SetSpecies gets a reference to the given Ga4ghOntologyTerm and assigns it to the Species field.
func (o *Ga4ghReference) SetSpecies(v Ga4ghOntologyTerm) {
	o.Species = &v
}


func (o Ga4ghReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Md5checksum != nil {
		toSerialize["md5checksum"] = o.Md5checksum
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SourceUri != nil {
		toSerialize["source_uri"] = o.SourceUri
	}
	if o.SourceAccessions != nil {
		toSerialize["source_accessions"] = o.SourceAccessions
	}
	if o.IsDerived != nil {
		toSerialize["is_derived"] = o.IsDerived
	}
	if o.SourceDivergence != nil {
		toSerialize["source_divergence"] = o.SourceDivergence
	}
	if o.Species != nil {
		toSerialize["species"] = o.Species
	}
	return json.Marshal(toSerialize)
}


