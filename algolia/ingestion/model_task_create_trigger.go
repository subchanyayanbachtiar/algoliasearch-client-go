// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// TaskCreateTrigger - struct for TaskCreateTrigger
type TaskCreateTrigger struct {
	OnDemandTriggerInput *OnDemandTriggerInput
	ScheduleTriggerInput *ScheduleTriggerInput
	SubscriptionTrigger  *SubscriptionTrigger
}

// OnDemandTriggerInputAsTaskCreateTrigger is a convenience function that returns OnDemandTriggerInput wrapped in TaskCreateTrigger
func OnDemandTriggerInputAsTaskCreateTrigger(v *OnDemandTriggerInput) TaskCreateTrigger {
	return TaskCreateTrigger{
		OnDemandTriggerInput: v,
	}
}

// ScheduleTriggerInputAsTaskCreateTrigger is a convenience function that returns ScheduleTriggerInput wrapped in TaskCreateTrigger
func ScheduleTriggerInputAsTaskCreateTrigger(v *ScheduleTriggerInput) TaskCreateTrigger {
	return TaskCreateTrigger{
		ScheduleTriggerInput: v,
	}
}

// SubscriptionTriggerAsTaskCreateTrigger is a convenience function that returns SubscriptionTrigger wrapped in TaskCreateTrigger
func SubscriptionTriggerAsTaskCreateTrigger(v *SubscriptionTrigger) TaskCreateTrigger {
	return TaskCreateTrigger{
		SubscriptionTrigger: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TaskCreateTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into OnDemandTriggerInput
	err = newStrictDecoder(data).Decode(&dst.OnDemandTriggerInput)
	if err == nil && validateStruct(dst.OnDemandTriggerInput) == nil {
		jsonOnDemandTriggerInput, _ := json.Marshal(dst.OnDemandTriggerInput)
		if string(jsonOnDemandTriggerInput) == "{}" { // empty struct
			dst.OnDemandTriggerInput = nil
		} else {
			return nil
		}
	} else {
		dst.OnDemandTriggerInput = nil
	}

	// try to unmarshal data into ScheduleTriggerInput
	err = newStrictDecoder(data).Decode(&dst.ScheduleTriggerInput)
	if err == nil && validateStruct(dst.ScheduleTriggerInput) == nil {
		jsonScheduleTriggerInput, _ := json.Marshal(dst.ScheduleTriggerInput)
		if string(jsonScheduleTriggerInput) == "{}" { // empty struct
			dst.ScheduleTriggerInput = nil
		} else {
			return nil
		}
	} else {
		dst.ScheduleTriggerInput = nil
	}

	// try to unmarshal data into SubscriptionTrigger
	err = newStrictDecoder(data).Decode(&dst.SubscriptionTrigger)
	if err == nil && validateStruct(dst.SubscriptionTrigger) == nil {
		jsonSubscriptionTrigger, _ := json.Marshal(dst.SubscriptionTrigger)
		if string(jsonSubscriptionTrigger) == "{}" { // empty struct
			dst.SubscriptionTrigger = nil
		} else {
			return nil
		}
	} else {
		dst.SubscriptionTrigger = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(TaskCreateTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TaskCreateTrigger) MarshalJSON() ([]byte, error) {
	if src.OnDemandTriggerInput != nil {
		return json.Marshal(&src.OnDemandTriggerInput)
	}

	if src.ScheduleTriggerInput != nil {
		return json.Marshal(&src.ScheduleTriggerInput)
	}

	if src.SubscriptionTrigger != nil {
		return json.Marshal(&src.SubscriptionTrigger)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TaskCreateTrigger) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.OnDemandTriggerInput != nil {
		return obj.OnDemandTriggerInput
	}

	if obj.ScheduleTriggerInput != nil {
		return obj.ScheduleTriggerInput
	}

	if obj.SubscriptionTrigger != nil {
		return obj.SubscriptionTrigger
	}

	// all schemas are nil
	return nil
}

type NullableTaskCreateTrigger struct {
	value *TaskCreateTrigger
	isSet bool
}

func (v NullableTaskCreateTrigger) Get() *TaskCreateTrigger {
	return v.value
}

func (v *NullableTaskCreateTrigger) Set(val *TaskCreateTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCreateTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCreateTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCreateTrigger(val *TaskCreateTrigger) *NullableTaskCreateTrigger {
	return &NullableTaskCreateTrigger{value: val, isSet: true}
}

func (v NullableTaskCreateTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCreateTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
