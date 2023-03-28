// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// ConsequenceQueryObject struct for ConsequenceQueryObject
type ConsequenceQueryObject struct {
	// Words to remove.
	Remove []string `json:"remove,omitempty"`
	// Edits to apply.
	Edits []Edit `json:"edits,omitempty"`
}

type ConsequenceQueryObjectOption func(f *ConsequenceQueryObject)

func WithConsequenceQueryObjectRemove(val []string) ConsequenceQueryObjectOption {
	return func(f *ConsequenceQueryObject) {
		f.Remove = val
	}
}

func WithConsequenceQueryObjectEdits(val []Edit) ConsequenceQueryObjectOption {
	return func(f *ConsequenceQueryObject) {
		f.Edits = val
	}
}

// NewConsequenceQueryObject instantiates a new ConsequenceQueryObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsequenceQueryObject(opts ...ConsequenceQueryObjectOption) *ConsequenceQueryObject {
	this := &ConsequenceQueryObject{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewConsequenceQueryObjectWithDefaults instantiates a new ConsequenceQueryObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsequenceQueryObjectWithDefaults() *ConsequenceQueryObject {
	this := &ConsequenceQueryObject{}
	return this
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *ConsequenceQueryObject) GetRemove() []string {
	if o == nil || o.Remove == nil {
		var ret []string
		return ret
	}
	return o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsequenceQueryObject) GetRemoveOk() ([]string, bool) {
	if o == nil || o.Remove == nil {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *ConsequenceQueryObject) HasRemove() bool {
	if o != nil && o.Remove != nil {
		return true
	}

	return false
}

// SetRemove gets a reference to the given []string and assigns it to the Remove field.
func (o *ConsequenceQueryObject) SetRemove(v []string) {
	o.Remove = v
}

// GetEdits returns the Edits field value if set, zero value otherwise.
func (o *ConsequenceQueryObject) GetEdits() []Edit {
	if o == nil || o.Edits == nil {
		var ret []Edit
		return ret
	}
	return o.Edits
}

// GetEditsOk returns a tuple with the Edits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsequenceQueryObject) GetEditsOk() ([]Edit, bool) {
	if o == nil || o.Edits == nil {
		return nil, false
	}
	return o.Edits, true
}

// HasEdits returns a boolean if a field has been set.
func (o *ConsequenceQueryObject) HasEdits() bool {
	if o != nil && o.Edits != nil {
		return true
	}

	return false
}

// SetEdits gets a reference to the given []Edit and assigns it to the Edits field.
func (o *ConsequenceQueryObject) SetEdits(v []Edit) {
	o.Edits = v
}

func (o ConsequenceQueryObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Remove != nil {
		toSerialize["remove"] = o.Remove
	}
	if o.Edits != nil {
		toSerialize["edits"] = o.Edits
	}
	return json.Marshal(toSerialize)
}

func (o ConsequenceQueryObject) String() string {
	out := "ConsequenceQueryObject {\n"
	out += fmt.Sprintf("  remove=%v\n", o.Remove)
	out += fmt.Sprintf("  edits=%v\n", o.Edits)
	out += "}"
	return out
}

type NullableConsequenceQueryObject struct {
	value *ConsequenceQueryObject
	isSet bool
}

func (v NullableConsequenceQueryObject) Get() *ConsequenceQueryObject {
	return v.value
}

func (v *NullableConsequenceQueryObject) Set(val *ConsequenceQueryObject) {
	v.value = val
	v.isSet = true
}

func (v NullableConsequenceQueryObject) IsSet() bool {
	return v.isSet
}

func (v *NullableConsequenceQueryObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsequenceQueryObject(val *ConsequenceQueryObject) *NullableConsequenceQueryObject {
	return &NullableConsequenceQueryObject{value: val, isSet: true}
}

func (v NullableConsequenceQueryObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsequenceQueryObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
