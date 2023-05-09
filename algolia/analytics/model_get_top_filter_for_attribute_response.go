// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetTopFilterForAttributeResponse struct for GetTopFilterForAttributeResponse
type GetTopFilterForAttributeResponse struct {
	// A list of filters for the given attributes.
	Values []GetTopFilterForAttribute `json:"values" validate:"required"`
}

// NewGetTopFilterForAttributeResponse instantiates a new GetTopFilterForAttributeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTopFilterForAttributeResponse(values []GetTopFilterForAttribute) *GetTopFilterForAttributeResponse {
	this := &GetTopFilterForAttributeResponse{}
	this.Values = values
	return this
}

// NewGetTopFilterForAttributeResponseWithDefaults instantiates a new GetTopFilterForAttributeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTopFilterForAttributeResponseWithDefaults() *GetTopFilterForAttributeResponse {
	this := &GetTopFilterForAttributeResponse{}
	return this
}

// GetValues returns the Values field value
func (o *GetTopFilterForAttributeResponse) GetValues() []GetTopFilterForAttribute {
	if o == nil {
		var ret []GetTopFilterForAttribute
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTopFilterForAttributeResponse) GetValuesOk() ([]GetTopFilterForAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTopFilterForAttributeResponse) SetValues(v []GetTopFilterForAttribute) {
	o.Values = v
}

func (o GetTopFilterForAttributeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

func (o GetTopFilterForAttributeResponse) String() string {
	out := ""
	out += fmt.Sprintf("  values=%v\n", o.Values)
	return fmt.Sprintf("GetTopFilterForAttributeResponse {\n%s}", out)
}

type NullableGetTopFilterForAttributeResponse struct {
	value *GetTopFilterForAttributeResponse
	isSet bool
}

func (v NullableGetTopFilterForAttributeResponse) Get() *GetTopFilterForAttributeResponse {
	return v.value
}

func (v *NullableGetTopFilterForAttributeResponse) Set(val *GetTopFilterForAttributeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTopFilterForAttributeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTopFilterForAttributeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTopFilterForAttributeResponse(val *GetTopFilterForAttributeResponse) *NullableGetTopFilterForAttributeResponse {
	return &NullableGetTopFilterForAttributeResponse{value: val, isSet: true}
}

func (v NullableGetTopFilterForAttributeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTopFilterForAttributeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
