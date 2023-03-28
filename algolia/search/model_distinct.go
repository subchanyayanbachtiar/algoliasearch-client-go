// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Distinct - Enables de-duplication or grouping of results.
type Distinct struct {
	Bool  *bool
	Int32 *int32
}

// boolAsDistinct is a convenience function that returns bool wrapped in Distinct
func BoolAsDistinct(v *bool) Distinct {
	return Distinct{
		Bool: v,
	}
}

// int32AsDistinct is a convenience function that returns int32 wrapped in Distinct
func Int32AsDistinct(v *int32) Distinct {
	return Distinct{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Distinct) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Int32 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Distinct)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Distinct)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Distinct) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Distinct) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableDistinct struct {
	value *Distinct
	isSet bool
}

func (v NullableDistinct) Get() *Distinct {
	return v.value
}

func (v *NullableDistinct) Set(val *Distinct) {
	v.value = val
	v.isSet = true
}

func (v NullableDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistinct(val *Distinct) *NullableDistinct {
	return &NullableDistinct{value: val, isSet: true}
}

func (v NullableDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
