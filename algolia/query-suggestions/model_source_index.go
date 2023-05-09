// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// SourceIndex struct for SourceIndex
type SourceIndex struct {
	// Source index name.
	IndexName string `json:"indexName" validate:"required"`
	// List of analytics tags to filter the popular searches per tag.
	AnalyticsTags []string `json:"analyticsTags,omitempty"`
	// List of facets to define as categories for the query suggestions.
	Facets []map[string]interface{} `json:"facets,omitempty"`
	// Minimum number of hits (e.g., matching records in the source index) to generate a suggestions.
	MinHits *int32 `json:"minHits,omitempty"`
	// Minimum number of required letters for a suggestion to remain.
	MinLetters *int32 `json:"minLetters,omitempty"`
	// List of facet attributes used to generate Query Suggestions. The resulting suggestions are every combination of the facets in the nested list (e.g., (facetA and facetB) and facetC).
	Generate [][]string `json:"generate,omitempty"`
	// List of external indices to use to generate custom Query Suggestions.
	External []SourceIndexExternal `json:"external,omitempty"`
}

type SourceIndexOption func(f *SourceIndex)

func WithSourceIndexAnalyticsTags(val []string) SourceIndexOption {
	return func(f *SourceIndex) {
		f.AnalyticsTags = val
	}
}

func WithSourceIndexFacets(val []map[string]interface{}) SourceIndexOption {
	return func(f *SourceIndex) {
		f.Facets = val
	}
}

func WithSourceIndexMinHits(val int32) SourceIndexOption {
	return func(f *SourceIndex) {
		f.MinHits = &val
	}
}

func WithSourceIndexMinLetters(val int32) SourceIndexOption {
	return func(f *SourceIndex) {
		f.MinLetters = &val
	}
}

func WithSourceIndexGenerate(val [][]string) SourceIndexOption {
	return func(f *SourceIndex) {
		f.Generate = val
	}
}

func WithSourceIndexExternal(val []SourceIndexExternal) SourceIndexOption {
	return func(f *SourceIndex) {
		f.External = val
	}
}

// NewSourceIndex instantiates a new SourceIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceIndex(indexName string, opts ...SourceIndexOption) *SourceIndex {
	this := &SourceIndex{}
	this.IndexName = indexName
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewSourceIndexWithDefaults instantiates a new SourceIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceIndexWithDefaults() *SourceIndex {
	this := &SourceIndex{}
	return this
}

// GetIndexName returns the IndexName field value
func (o *SourceIndex) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value
func (o *SourceIndex) SetIndexName(v string) {
	o.IndexName = v
}

// GetAnalyticsTags returns the AnalyticsTags field value if set, zero value otherwise.
func (o *SourceIndex) GetAnalyticsTags() []string {
	if o == nil || o.AnalyticsTags == nil {
		var ret []string
		return ret
	}
	return o.AnalyticsTags
}

// GetAnalyticsTagsOk returns a tuple with the AnalyticsTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetAnalyticsTagsOk() ([]string, bool) {
	if o == nil || o.AnalyticsTags == nil {
		return nil, false
	}
	return o.AnalyticsTags, true
}

// HasAnalyticsTags returns a boolean if a field has been set.
func (o *SourceIndex) HasAnalyticsTags() bool {
	if o != nil && o.AnalyticsTags != nil {
		return true
	}

	return false
}

// SetAnalyticsTags gets a reference to the given []string and assigns it to the AnalyticsTags field.
func (o *SourceIndex) SetAnalyticsTags(v []string) {
	o.AnalyticsTags = v
}

// GetFacets returns the Facets field value if set, zero value otherwise.
func (o *SourceIndex) GetFacets() []map[string]interface{} {
	if o == nil || o.Facets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetFacetsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Facets == nil {
		return nil, false
	}
	return o.Facets, true
}

// HasFacets returns a boolean if a field has been set.
func (o *SourceIndex) HasFacets() bool {
	if o != nil && o.Facets != nil {
		return true
	}

	return false
}

// SetFacets gets a reference to the given []map[string]interface{} and assigns it to the Facets field.
func (o *SourceIndex) SetFacets(v []map[string]interface{}) {
	o.Facets = v
}

// GetMinHits returns the MinHits field value if set, zero value otherwise.
func (o *SourceIndex) GetMinHits() int32 {
	if o == nil || o.MinHits == nil {
		var ret int32
		return ret
	}
	return *o.MinHits
}

// GetMinHitsOk returns a tuple with the MinHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetMinHitsOk() (*int32, bool) {
	if o == nil || o.MinHits == nil {
		return nil, false
	}
	return o.MinHits, true
}

// HasMinHits returns a boolean if a field has been set.
func (o *SourceIndex) HasMinHits() bool {
	if o != nil && o.MinHits != nil {
		return true
	}

	return false
}

// SetMinHits gets a reference to the given int32 and assigns it to the MinHits field.
func (o *SourceIndex) SetMinHits(v int32) {
	o.MinHits = &v
}

// GetMinLetters returns the MinLetters field value if set, zero value otherwise.
func (o *SourceIndex) GetMinLetters() int32 {
	if o == nil || o.MinLetters == nil {
		var ret int32
		return ret
	}
	return *o.MinLetters
}

// GetMinLettersOk returns a tuple with the MinLetters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetMinLettersOk() (*int32, bool) {
	if o == nil || o.MinLetters == nil {
		return nil, false
	}
	return o.MinLetters, true
}

// HasMinLetters returns a boolean if a field has been set.
func (o *SourceIndex) HasMinLetters() bool {
	if o != nil && o.MinLetters != nil {
		return true
	}

	return false
}

// SetMinLetters gets a reference to the given int32 and assigns it to the MinLetters field.
func (o *SourceIndex) SetMinLetters(v int32) {
	o.MinLetters = &v
}

// GetGenerate returns the Generate field value if set, zero value otherwise.
func (o *SourceIndex) GetGenerate() [][]string {
	if o == nil || o.Generate == nil {
		var ret [][]string
		return ret
	}
	return o.Generate
}

// GetGenerateOk returns a tuple with the Generate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetGenerateOk() ([][]string, bool) {
	if o == nil || o.Generate == nil {
		return nil, false
	}
	return o.Generate, true
}

// HasGenerate returns a boolean if a field has been set.
func (o *SourceIndex) HasGenerate() bool {
	if o != nil && o.Generate != nil {
		return true
	}

	return false
}

// SetGenerate gets a reference to the given [][]string and assigns it to the Generate field.
func (o *SourceIndex) SetGenerate(v [][]string) {
	o.Generate = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *SourceIndex) GetExternal() []SourceIndexExternal {
	if o == nil || o.External == nil {
		var ret []SourceIndexExternal
		return ret
	}
	return o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetExternalOk() ([]SourceIndexExternal, bool) {
	if o == nil || o.External == nil {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *SourceIndex) HasExternal() bool {
	if o != nil && o.External != nil {
		return true
	}

	return false
}

// SetExternal gets a reference to the given []SourceIndexExternal and assigns it to the External field.
func (o *SourceIndex) SetExternal(v []SourceIndexExternal) {
	o.External = v
}

func (o SourceIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	if o.AnalyticsTags != nil {
		toSerialize["analyticsTags"] = o.AnalyticsTags
	}
	if o.Facets != nil {
		toSerialize["facets"] = o.Facets
	}
	if o.MinHits != nil {
		toSerialize["minHits"] = o.MinHits
	}
	if o.MinLetters != nil {
		toSerialize["minLetters"] = o.MinLetters
	}
	if o.Generate != nil {
		toSerialize["generate"] = o.Generate
	}
	if o.External != nil {
		toSerialize["external"] = o.External
	}
	return json.Marshal(toSerialize)
}

func (o SourceIndex) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  analyticsTags=%v\n", o.AnalyticsTags)
	out += fmt.Sprintf("  facets=%v\n", o.Facets)
	out += fmt.Sprintf("  minHits=%v\n", o.MinHits)
	out += fmt.Sprintf("  minLetters=%v\n", o.MinLetters)
	out += fmt.Sprintf("  generate=%v\n", o.Generate)
	out += fmt.Sprintf("  external=%v\n", o.External)
	return fmt.Sprintf("SourceIndex {\n%s}", out)
}

type NullableSourceIndex struct {
	value *SourceIndex
	isSet bool
}

func (v NullableSourceIndex) Get() *SourceIndex {
	return v.value
}

func (v *NullableSourceIndex) Set(val *SourceIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceIndex(val *SourceIndex) *NullableSourceIndex {
	return &NullableSourceIndex{value: val, isSet: true}
}

func (v NullableSourceIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
