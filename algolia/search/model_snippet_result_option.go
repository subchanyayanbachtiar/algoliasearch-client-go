// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SnippetResultOption Snippeted attributes show parts of the matched attributes. Only returned when attributesToSnippet is non-empty.
type SnippetResultOption struct {
	// Markup text with `facetQuery` matches highlighted.
	Value      string     `json:"value" validate:"required"`
	MatchLevel MatchLevel `json:"matchLevel" validate:"required"`
}

// NewSnippetResultOption instantiates a new SnippetResultOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetResultOption(value string, matchLevel MatchLevel) *SnippetResultOption {
	this := &SnippetResultOption{}
	this.Value = value
	this.MatchLevel = matchLevel
	return this
}

// NewSnippetResultOptionWithDefaults instantiates a new SnippetResultOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetResultOptionWithDefaults() *SnippetResultOption {
	this := &SnippetResultOption{}
	return this
}

// GetValue returns the Value field value
func (o *SnippetResultOption) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SnippetResultOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SnippetResultOption) SetValue(v string) {
	o.Value = v
}

// GetMatchLevel returns the MatchLevel field value
func (o *SnippetResultOption) GetMatchLevel() MatchLevel {
	if o == nil {
		var ret MatchLevel
		return ret
	}

	return o.MatchLevel
}

// GetMatchLevelOk returns a tuple with the MatchLevel field value
// and a boolean to check if the value has been set.
func (o *SnippetResultOption) GetMatchLevelOk() (*MatchLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchLevel, true
}

// SetMatchLevel sets field value
func (o *SnippetResultOption) SetMatchLevel(v MatchLevel) {
	o.MatchLevel = v
}

func (o SnippetResultOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["matchLevel"] = o.MatchLevel
	}
	return json.Marshal(toSerialize)
}

func (o SnippetResultOption) String() string {
	out := ""
	out += fmt.Sprintf("  value=%v\n", o.Value)
	out += fmt.Sprintf("  matchLevel=%v\n", o.MatchLevel)
	return fmt.Sprintf("SnippetResultOption {\n%s}", out)
}

type NullableSnippetResultOption struct {
	value *SnippetResultOption
	isSet bool
}

func (v NullableSnippetResultOption) Get() *SnippetResultOption {
	return v.value
}

func (v *NullableSnippetResultOption) Set(val *SnippetResultOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippetResultOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippetResultOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippetResultOption(val *SnippetResultOption) *NullableSnippetResultOption {
	return &NullableSnippetResultOption{value: val, isSet: true}
}

func (v NullableSnippetResultOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippetResultOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
