// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchForFacetValuesResponse struct for SearchForFacetValuesResponse
type SearchForFacetValuesResponse struct {
	FacetHits []FacetHits `json:"facetHits"`
}

// NewSearchForFacetValuesResponse instantiates a new SearchForFacetValuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchForFacetValuesResponse(facetHits []FacetHits) *SearchForFacetValuesResponse {
	this := &SearchForFacetValuesResponse{}
	this.FacetHits = facetHits
	return this
}

// NewSearchForFacetValuesResponseWithDefaults instantiates a new SearchForFacetValuesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchForFacetValuesResponseWithDefaults() *SearchForFacetValuesResponse {
	this := &SearchForFacetValuesResponse{}
	return this
}

// GetFacetHits returns the FacetHits field value
func (o *SearchForFacetValuesResponse) GetFacetHits() []FacetHits {
	if o == nil {
		var ret []FacetHits
		return ret
	}

	return o.FacetHits
}

// GetFacetHitsOk returns a tuple with the FacetHits field value
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesResponse) GetFacetHitsOk() ([]FacetHits, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacetHits, true
}

// SetFacetHits sets field value
func (o *SearchForFacetValuesResponse) SetFacetHits(v []FacetHits) {
	o.FacetHits = v
}

func (o SearchForFacetValuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["facetHits"] = o.FacetHits
	}
	return json.Marshal(toSerialize)
}

func (o SearchForFacetValuesResponse) String() string {
	out := "SearchForFacetValuesResponse {\n"
	out += fmt.Sprintf("  facetHits=%v\n", o.FacetHits)
	out += "}"
	return out
}

type NullableSearchForFacetValuesResponse struct {
	value *SearchForFacetValuesResponse
	isSet bool
}

func (v NullableSearchForFacetValuesResponse) Get() *SearchForFacetValuesResponse {
	return v.value
}

func (v *NullableSearchForFacetValuesResponse) Set(val *SearchForFacetValuesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchForFacetValuesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchForFacetValuesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchForFacetValuesResponse(val *SearchForFacetValuesResponse) *NullableSearchForFacetValuesResponse {
	return &NullableSearchForFacetValuesResponse{value: val, isSet: true}
}

func (v NullableSearchForFacetValuesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchForFacetValuesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
