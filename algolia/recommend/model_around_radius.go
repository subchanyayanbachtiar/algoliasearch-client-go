// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// AroundRadius - [Maximum radius](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#increase-the-search-radius) for a geographical search (in meters).
type AroundRadius struct {
	AroundRadiusAll *AroundRadiusAll
	Int32           *int32
}

// int32AsAroundRadius is a convenience function that returns int32 wrapped in AroundRadius.
func Int32AsAroundRadius(v int32) *AroundRadius {
	return &AroundRadius{
		Int32: &v,
	}
}

// AroundRadiusAllAsAroundRadius is a convenience function that returns AroundRadiusAll wrapped in AroundRadius.
func AroundRadiusAllAsAroundRadius(v AroundRadiusAll) *AroundRadius {
	return &AroundRadius{
		AroundRadiusAll: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *AroundRadius) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into AroundRadiusAll
	err = newStrictDecoder(data).Decode(&dst.AroundRadiusAll)
	if err == nil && validateStruct(dst.AroundRadiusAll) == nil {
		jsonAroundRadiusAll, _ := json.Marshal(dst.AroundRadiusAll)
		if string(jsonAroundRadiusAll) == "{}" { // empty struct
			dst.AroundRadiusAll = nil
		} else {
			return nil
		}
	} else {
		dst.AroundRadiusAll = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil && validateStruct(dst.Int32) == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil
		}
	} else {
		dst.Int32 = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(AroundRadius)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src AroundRadius) MarshalJSON() ([]byte, error) {
	if src.AroundRadiusAll != nil {
		return json.Marshal(&src.AroundRadiusAll)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj *AroundRadius) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.AroundRadiusAll != nil {
		return obj.AroundRadiusAll
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableAroundRadius struct {
	value *AroundRadius
	isSet bool
}

func (v NullableAroundRadius) Get() *AroundRadius {
	return v.value
}

func (v *NullableAroundRadius) Set(val *AroundRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableAroundRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableAroundRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAroundRadius(val *AroundRadius) *NullableAroundRadius {
	return &NullableAroundRadius{value: val, isSet: true}
}

func (v NullableAroundRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAroundRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
