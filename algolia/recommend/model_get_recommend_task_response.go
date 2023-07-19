// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// GetRecommendTaskResponse struct for GetRecommendTaskResponse
type GetRecommendTaskResponse struct {
	Status TaskStatus `json:"status" validate:"required"`
}

// NewGetRecommendTaskResponse instantiates a new GetRecommendTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecommendTaskResponse(status TaskStatus) *GetRecommendTaskResponse {
	this := &GetRecommendTaskResponse{}
	this.Status = status
	return this
}

// NewGetRecommendTaskResponseWithDefaults instantiates a new GetRecommendTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecommendTaskResponseWithDefaults() *GetRecommendTaskResponse {
	this := &GetRecommendTaskResponse{}
	return this
}

// GetStatus returns the Status field value
func (o *GetRecommendTaskResponse) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetRecommendTaskResponse) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetRecommendTaskResponse) SetStatus(v TaskStatus) {
	o.Status = v
}

func (o GetRecommendTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

func (o GetRecommendTaskResponse) String() string {
	out := ""
	out += fmt.Sprintf("  status=%v\n", o.Status)
	return fmt.Sprintf("GetRecommendTaskResponse {\n%s}", out)
}

type NullableGetRecommendTaskResponse struct {
	value *GetRecommendTaskResponse
	isSet bool
}

func (v NullableGetRecommendTaskResponse) Get() *GetRecommendTaskResponse {
	return v.value
}

func (v *NullableGetRecommendTaskResponse) Set(val *GetRecommendTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecommendTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecommendTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecommendTaskResponse(val *GetRecommendTaskResponse) *NullableGetRecommendTaskResponse {
	return &NullableGetRecommendTaskResponse{value: val, isSet: true}
}

func (v NullableGetRecommendTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecommendTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
