// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// Discount - The absolute value of the discount for this product, in units of `currency`.
type Discount struct {
	Float64 *float64
	String  *string
}

// float64AsDiscount is a convenience function that returns float64 wrapped in Discount.
func Float64AsDiscount(v float64) *Discount {
	return &Discount{
		Float64: &v,
	}
}

// stringAsDiscount is a convenience function that returns string wrapped in Discount.
func StringAsDiscount(v string) *Discount {
	return &Discount{
		String: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *Discount) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into Float64
	err = newStrictDecoder(data).Decode(&dst.Float64)
	if err == nil && validateStruct(dst.Float64) == nil {
		jsonFloat64, _ := json.Marshal(dst.Float64)
		if string(jsonFloat64) == "{}" { // empty struct
			dst.Float64 = nil
		} else {
			return nil
		}
	} else {
		dst.Float64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil && validateStruct(dst.String) == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(Discount)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src Discount) MarshalJSON() ([]byte, error) {
	if src.Float64 != nil {
		return json.Marshal(&src.Float64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj *Discount) GetActualInstance() any {
	if obj == nil {
		return nil
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

type NullableDiscount struct {
	value *Discount
	isSet bool
}

func (v NullableDiscount) Get() *Discount {
	return v.value
}

func (v *NullableDiscount) Set(val *Discount) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscount(val *Discount) *NullableDiscount {
	return &NullableDiscount{value: val, isSet: true}
}

func (v NullableDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
