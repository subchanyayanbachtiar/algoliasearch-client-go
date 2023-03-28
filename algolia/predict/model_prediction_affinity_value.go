// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// PredictionAffinityValue - struct for PredictionAffinityValue
type PredictionAffinityValue struct {
	Bool    *bool
	Float64 *float64
	String  *string
}

// boolAsPredictionAffinityValue is a convenience function that returns bool wrapped in PredictionAffinityValue
func BoolAsPredictionAffinityValue(v *bool) PredictionAffinityValue {
	return PredictionAffinityValue{
		Bool: v,
	}
}

// float64AsPredictionAffinityValue is a convenience function that returns float64 wrapped in PredictionAffinityValue
func Float64AsPredictionAffinityValue(v *float64) PredictionAffinityValue {
	return PredictionAffinityValue{
		Float64: v,
	}
}

// stringAsPredictionAffinityValue is a convenience function that returns string wrapped in PredictionAffinityValue
func StringAsPredictionAffinityValue(v *string) PredictionAffinityValue {
	return PredictionAffinityValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PredictionAffinityValue) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into Float64
	err = newStrictDecoder(data).Decode(&dst.Float64)
	if err == nil {
		jsonFloat64, _ := json.Marshal(dst.Float64)
		if string(jsonFloat64) == "{}" { // empty struct
			dst.Float64 = nil
		} else {
			match++
		}
	} else {
		dst.Float64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Float64 = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PredictionAffinityValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PredictionAffinityValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PredictionAffinityValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float64 != nil {
		return json.Marshal(&src.Float64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PredictionAffinityValue) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePredictionAffinityValue struct {
	value *PredictionAffinityValue
	isSet bool
}

func (v NullablePredictionAffinityValue) Get() *PredictionAffinityValue {
	return v.value
}

func (v *NullablePredictionAffinityValue) Set(val *PredictionAffinityValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionAffinityValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionAffinityValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionAffinityValue(val *PredictionAffinityValue) *NullablePredictionAffinityValue {
	return &NullablePredictionAffinityValue{value: val, isSet: true}
}

func (v NullablePredictionAffinityValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionAffinityValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
