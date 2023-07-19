// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchDictionaryEntriesParams `searchDictionaryEntries` parameters.
type SearchDictionaryEntriesParams struct {
	// Text to search for in an index.
	Query string `json:"query" validate:"required"`
	// Page to retrieve (the first page is `0`, not `1`).
	Page *int32 `json:"page,omitempty"`
	// Number of hits per page.
	HitsPerPage *int32 `json:"hitsPerPage,omitempty"`
	// [Supported language ISO code](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/supported-languages/).
	Language *string `json:"language,omitempty"`
}

type SearchDictionaryEntriesParamsOption func(f *SearchDictionaryEntriesParams)

func WithSearchDictionaryEntriesParamsPage(val int32) SearchDictionaryEntriesParamsOption {
	return func(f *SearchDictionaryEntriesParams) {
		f.Page = &val
	}
}

func WithSearchDictionaryEntriesParamsHitsPerPage(val int32) SearchDictionaryEntriesParamsOption {
	return func(f *SearchDictionaryEntriesParams) {
		f.HitsPerPage = &val
	}
}

func WithSearchDictionaryEntriesParamsLanguage(val string) SearchDictionaryEntriesParamsOption {
	return func(f *SearchDictionaryEntriesParams) {
		f.Language = &val
	}
}

// NewSearchDictionaryEntriesParams instantiates a new SearchDictionaryEntriesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchDictionaryEntriesParams(query string, opts ...SearchDictionaryEntriesParamsOption) *SearchDictionaryEntriesParams {
	this := &SearchDictionaryEntriesParams{}
	this.Query = query
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewSearchDictionaryEntriesParamsWithDefaults instantiates a new SearchDictionaryEntriesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchDictionaryEntriesParamsWithDefaults() *SearchDictionaryEntriesParams {
	this := &SearchDictionaryEntriesParams{}
	var query string = ""
	this.Query = query
	var page int32 = 0
	this.Page = &page
	var hitsPerPage int32 = 20
	this.HitsPerPage = &hitsPerPage
	return this
}

// GetQuery returns the Query field value
func (o *SearchDictionaryEntriesParams) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchDictionaryEntriesParams) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchDictionaryEntriesParams) SetQuery(v string) {
	o.Query = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SearchDictionaryEntriesParams) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDictionaryEntriesParams) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SearchDictionaryEntriesParams) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *SearchDictionaryEntriesParams) SetPage(v int32) {
	o.Page = &v
}

// GetHitsPerPage returns the HitsPerPage field value if set, zero value otherwise.
func (o *SearchDictionaryEntriesParams) GetHitsPerPage() int32 {
	if o == nil || o.HitsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.HitsPerPage
}

// GetHitsPerPageOk returns a tuple with the HitsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDictionaryEntriesParams) GetHitsPerPageOk() (*int32, bool) {
	if o == nil || o.HitsPerPage == nil {
		return nil, false
	}
	return o.HitsPerPage, true
}

// HasHitsPerPage returns a boolean if a field has been set.
func (o *SearchDictionaryEntriesParams) HasHitsPerPage() bool {
	if o != nil && o.HitsPerPage != nil {
		return true
	}

	return false
}

// SetHitsPerPage gets a reference to the given int32 and assigns it to the HitsPerPage field.
func (o *SearchDictionaryEntriesParams) SetHitsPerPage(v int32) {
	o.HitsPerPage = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *SearchDictionaryEntriesParams) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDictionaryEntriesParams) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SearchDictionaryEntriesParams) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *SearchDictionaryEntriesParams) SetLanguage(v string) {
	o.Language = &v
}

func (o SearchDictionaryEntriesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.HitsPerPage != nil {
		toSerialize["hitsPerPage"] = o.HitsPerPage
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

func (o SearchDictionaryEntriesParams) String() string {
	out := ""
	out += fmt.Sprintf("  query=%v\n", o.Query)
	out += fmt.Sprintf("  page=%v\n", o.Page)
	out += fmt.Sprintf("  hitsPerPage=%v\n", o.HitsPerPage)
	out += fmt.Sprintf("  language=%v\n", o.Language)
	return fmt.Sprintf("SearchDictionaryEntriesParams {\n%s}", out)
}

type NullableSearchDictionaryEntriesParams struct {
	value *SearchDictionaryEntriesParams
	isSet bool
}

func (v NullableSearchDictionaryEntriesParams) Get() *SearchDictionaryEntriesParams {
	return v.value
}

func (v *NullableSearchDictionaryEntriesParams) Set(val *SearchDictionaryEntriesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDictionaryEntriesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDictionaryEntriesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDictionaryEntriesParams(val *SearchDictionaryEntriesParams) *NullableSearchDictionaryEntriesParams {
	return &NullableSearchDictionaryEntriesParams{value: val, isSet: true}
}

func (v NullableSearchDictionaryEntriesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDictionaryEntriesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
