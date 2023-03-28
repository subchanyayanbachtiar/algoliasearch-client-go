// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// SegmentOperandProperty Operand for user profile properties.
type SegmentOperandProperty struct {
	Name    string                  `json:"name"`
	Filters []SegmentPropertyFilter `json:"filters"`
}

// NewSegmentOperandProperty instantiates a new SegmentOperandProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentOperandProperty(name string, filters []SegmentPropertyFilter) *SegmentOperandProperty {
	this := &SegmentOperandProperty{}
	this.Name = name
	this.Filters = filters
	return this
}

// NewSegmentOperandPropertyWithDefaults instantiates a new SegmentOperandProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentOperandPropertyWithDefaults() *SegmentOperandProperty {
	this := &SegmentOperandProperty{}
	return this
}

// GetName returns the Name field value
func (o *SegmentOperandProperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SegmentOperandProperty) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SegmentOperandProperty) SetName(v string) {
	o.Name = v
}

// GetFilters returns the Filters field value
func (o *SegmentOperandProperty) GetFilters() []SegmentPropertyFilter {
	if o == nil {
		var ret []SegmentPropertyFilter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *SegmentOperandProperty) GetFiltersOk() ([]SegmentPropertyFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *SegmentOperandProperty) SetFilters(v []SegmentPropertyFilter) {
	o.Filters = v
}

func (o SegmentOperandProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

func (o SegmentOperandProperty) String() string {
	out := "SegmentOperandProperty {\n"
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  filters=%v\n", o.Filters)
	out += "}"
	return out
}

type NullableSegmentOperandProperty struct {
	value *SegmentOperandProperty
	isSet bool
}

func (v NullableSegmentOperandProperty) Get() *SegmentOperandProperty {
	return v.value
}

func (v *NullableSegmentOperandProperty) Set(val *SegmentOperandProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentOperandProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentOperandProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentOperandProperty(val *SegmentOperandProperty) *NullableSegmentOperandProperty {
	return &NullableSegmentOperandProperty{value: val, isSet: true}
}

func (v NullableSegmentOperandProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentOperandProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
