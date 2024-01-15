// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// AssignUserIdParams Assign userID parameters.
type AssignUserIdParams struct {
	// Cluster name.
	Cluster string `json:"cluster"`
}

// NewAssignUserIdParams instantiates a new AssignUserIdParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssignUserIdParams(cluster string) *AssignUserIdParams {
	this := &AssignUserIdParams{}
	this.Cluster = cluster
	return this
}

// NewEmptyAssignUserIdParams return a pointer to an empty AssignUserIdParams object.
func NewEmptyAssignUserIdParams() *AssignUserIdParams {
	return &AssignUserIdParams{}
}

// GetCluster returns the Cluster field value.
func (o *AssignUserIdParams) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *AssignUserIdParams) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *AssignUserIdParams) SetCluster(v string) *AssignUserIdParams {
	o.Cluster = v
	return o
}

func (o AssignUserIdParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

func (o AssignUserIdParams) String() string {
	out := ""
	out += fmt.Sprintf("  cluster=%v\n", o.Cluster)
	return fmt.Sprintf("AssignUserIdParams {\n%s}", out)
}

type NullableAssignUserIdParams struct {
	value *AssignUserIdParams
	isSet bool
}

func (v NullableAssignUserIdParams) Get() *AssignUserIdParams {
	return v.value
}

func (v *NullableAssignUserIdParams) Set(val *AssignUserIdParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignUserIdParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignUserIdParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignUserIdParams(val *AssignUserIdParams) *NullableAssignUserIdParams {
	return &NullableAssignUserIdParams{value: val, isSet: true}
}

func (v NullableAssignUserIdParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignUserIdParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
