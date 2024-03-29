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

// This is the response from `POST /datasets/search` expressed as JSON.
type Ga4ghSearchDatasetsResponse struct {
	// The list of datasets.
	Datasets *[]Ga4ghDataset `json:"datasets,omitempty"`

	// The continuation token, which is used to page through large result sets. Provide this value in a subsequent request to return the next page of results. This field will be empty if there aren't any additional results.
	NextPageToken *string `json:"next_page_token,omitempty"`

}

// GetDatasets returns the Datasets field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDatasetsResponse) GetDatasets() []Ga4ghDataset {
	if o == nil || o.Datasets == nil {
		var ret []Ga4ghDataset
		return ret
	}
	return *o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDatasetsResponse) GetDatasetsOk() ([]Ga4ghDataset, bool) {
	if o == nil || o.Datasets == nil {
		var ret []Ga4ghDataset
		return ret, false
	}
	return *o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *Ga4ghSearchDatasetsResponse) HasDatasets() bool {
	if o != nil && o.Datasets != nil {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []Ga4ghDataset and assigns it to the Datasets field.
func (o *Ga4ghSearchDatasetsResponse) SetDatasets(v []Ga4ghDataset) {
	o.Datasets = &v
}

// GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.
func (o *Ga4ghSearchDatasetsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ga4ghSearchDatasetsResponse) GetNextPageTokenOk() (string, bool) {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret, false
	}
	return *o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Ga4ghSearchDatasetsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Ga4ghSearchDatasetsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}


func (o Ga4ghSearchDatasetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datasets != nil {
		toSerialize["datasets"] = o.Datasets
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}


