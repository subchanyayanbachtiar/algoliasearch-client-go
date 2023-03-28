// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// DictionaryEntryState The state of the dictionary entry.
type DictionaryEntryState string

// List of dictionaryEntryState
const (
	DICTIONARYENTRYSTATE_ENABLED  DictionaryEntryState = "enabled"
	DICTIONARYENTRYSTATE_DISABLED DictionaryEntryState = "disabled"
)

// All allowed values of DictionaryEntryState enum
var AllowedDictionaryEntryStateEnumValues = []DictionaryEntryState{
	"enabled",
	"disabled",
}

func (v *DictionaryEntryState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DictionaryEntryState(value)
	for _, existing := range AllowedDictionaryEntryStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DictionaryEntryState", value)
}

// NewDictionaryEntryStateFromValue returns a pointer to a valid DictionaryEntryState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDictionaryEntryStateFromValue(v string) (*DictionaryEntryState, error) {
	ev := DictionaryEntryState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DictionaryEntryState: valid values are %v", v, AllowedDictionaryEntryStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DictionaryEntryState) IsValid() bool {
	for _, existing := range AllowedDictionaryEntryStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dictionaryEntryState value
func (v DictionaryEntryState) Ptr() *DictionaryEntryState {
	return &v
}

type NullableDictionaryEntryState struct {
	value *DictionaryEntryState
	isSet bool
}

func (v NullableDictionaryEntryState) Get() *DictionaryEntryState {
	return v.value
}

func (v *NullableDictionaryEntryState) Set(val *DictionaryEntryState) {
	v.value = val
	v.isSet = true
}

func (v NullableDictionaryEntryState) IsSet() bool {
	return v.isSet
}

func (v *NullableDictionaryEntryState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDictionaryEntryState(val *DictionaryEntryState) *NullableDictionaryEntryState {
	return &NullableDictionaryEntryState{value: val, isSet: true}
}

func (v NullableDictionaryEntryState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDictionaryEntryState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
