// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Cursor struct for Cursor
type Cursor struct {
	// Cursor indicating the location to resume browsing from. Must match the value returned by the previous call.
	Cursor *string `json:"cursor,omitempty"`
}

type CursorOption func(f *Cursor)

func WithCursorCursor(val string) CursorOption {
	return func(f *Cursor) {
		f.Cursor = &val
	}
}

// NewCursor instantiates a new Cursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursor(opts ...CursorOption) *Cursor {
	this := &Cursor{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewCursorWithDefaults instantiates a new Cursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorWithDefaults() *Cursor {
	this := &Cursor{}
	return this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *Cursor) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cursor) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *Cursor) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *Cursor) SetCursor(v string) {
	o.Cursor = &v
}

func (o Cursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

func (o Cursor) String() string {
	out := "Cursor {\n"
	out += fmt.Sprintf("  cursor=%v\n", o.Cursor)
	out += "}"
	return out
}

type NullableCursor struct {
	value *Cursor
	isSet bool
}

func (v NullableCursor) Get() *Cursor {
	return v.value
}

func (v *NullableCursor) Set(val *Cursor) {
	v.value = val
	v.isSet = true
}

func (v NullableCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursor(val *Cursor) *NullableCursor {
	return &NullableCursor{value: val, isSet: true}
}

func (v NullableCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
