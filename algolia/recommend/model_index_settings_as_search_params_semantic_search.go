// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// IndexSettingsAsSearchParamsSemanticSearch The settings relevant for configuration of the semantic search engine. These settings are only used when the mode is set to 'neuralSearch'.
type IndexSettingsAsSearchParamsSemanticSearch struct {
	// When null, the current index / replica group will be used as the event source.
	EventSources []string `json:"eventSources,omitempty"`
}

type IndexSettingsAsSearchParamsSemanticSearchOption func(f *IndexSettingsAsSearchParamsSemanticSearch)

func WithIndexSettingsAsSearchParamsSemanticSearchEventSources(val []string) IndexSettingsAsSearchParamsSemanticSearchOption {
	return func(f *IndexSettingsAsSearchParamsSemanticSearch) {
		f.EventSources = val
	}
}

// NewIndexSettingsAsSearchParamsSemanticSearch instantiates a new IndexSettingsAsSearchParamsSemanticSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexSettingsAsSearchParamsSemanticSearch(opts ...IndexSettingsAsSearchParamsSemanticSearchOption) *IndexSettingsAsSearchParamsSemanticSearch {
	this := &IndexSettingsAsSearchParamsSemanticSearch{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewIndexSettingsAsSearchParamsSemanticSearchWithDefaults instantiates a new IndexSettingsAsSearchParamsSemanticSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexSettingsAsSearchParamsSemanticSearchWithDefaults() *IndexSettingsAsSearchParamsSemanticSearch {
	this := &IndexSettingsAsSearchParamsSemanticSearch{}
	return this
}

// GetEventSources returns the EventSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexSettingsAsSearchParamsSemanticSearch) GetEventSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EventSources
}

// GetEventSourcesOk returns a tuple with the EventSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexSettingsAsSearchParamsSemanticSearch) GetEventSourcesOk() ([]string, bool) {
	if o == nil || o.EventSources == nil {
		return nil, false
	}
	return o.EventSources, true
}

// HasEventSources returns a boolean if a field has been set.
func (o *IndexSettingsAsSearchParamsSemanticSearch) HasEventSources() bool {
	if o != nil && o.EventSources != nil {
		return true
	}

	return false
}

// SetEventSources gets a reference to the given []string and assigns it to the EventSources field.
func (o *IndexSettingsAsSearchParamsSemanticSearch) SetEventSources(v []string) {
	o.EventSources = v
}

func (o IndexSettingsAsSearchParamsSemanticSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.EventSources != nil {
		toSerialize["eventSources"] = o.EventSources
	}
	return json.Marshal(toSerialize)
}

func (o IndexSettingsAsSearchParamsSemanticSearch) String() string {
	out := ""
	out += fmt.Sprintf("  eventSources=%v\n", o.EventSources)
	return fmt.Sprintf("IndexSettingsAsSearchParamsSemanticSearch {\n%s}", out)
}

type NullableIndexSettingsAsSearchParamsSemanticSearch struct {
	value *IndexSettingsAsSearchParamsSemanticSearch
	isSet bool
}

func (v NullableIndexSettingsAsSearchParamsSemanticSearch) Get() *IndexSettingsAsSearchParamsSemanticSearch {
	return v.value
}

func (v *NullableIndexSettingsAsSearchParamsSemanticSearch) Set(val *IndexSettingsAsSearchParamsSemanticSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexSettingsAsSearchParamsSemanticSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexSettingsAsSearchParamsSemanticSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexSettingsAsSearchParamsSemanticSearch(val *IndexSettingsAsSearchParamsSemanticSearch) *NullableIndexSettingsAsSearchParamsSemanticSearch {
	return &NullableIndexSettingsAsSearchParamsSemanticSearch{value: val, isSet: true}
}

func (v NullableIndexSettingsAsSearchParamsSemanticSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexSettingsAsSearchParamsSemanticSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
