// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetAverageClickPositionResponse struct for GetAverageClickPositionResponse.
type GetAverageClickPositionResponse struct {
	// Average count of all click events.
	Average float64 `json:"average"`
	// Number of click events.
	ClickCount int32 `json:"clickCount"`
	// Average click positions.
	Dates []AverageClickEvent `json:"dates"`
}

// NewGetAverageClickPositionResponse instantiates a new GetAverageClickPositionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAverageClickPositionResponse(average float64, clickCount int32, dates []AverageClickEvent) *GetAverageClickPositionResponse {
	this := &GetAverageClickPositionResponse{}
	this.Average = average
	this.ClickCount = clickCount
	this.Dates = dates
	return this
}

// NewEmptyGetAverageClickPositionResponse return a pointer to an empty GetAverageClickPositionResponse object.
func NewEmptyGetAverageClickPositionResponse() *GetAverageClickPositionResponse {
	return &GetAverageClickPositionResponse{}
}

// GetAverage returns the Average field value.
func (o *GetAverageClickPositionResponse) GetAverage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Average
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
func (o *GetAverageClickPositionResponse) GetAverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Average, true
}

// SetAverage sets field value.
func (o *GetAverageClickPositionResponse) SetAverage(v float64) *GetAverageClickPositionResponse {
	o.Average = v
	return o
}

// GetClickCount returns the ClickCount field value.
func (o *GetAverageClickPositionResponse) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *GetAverageClickPositionResponse) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value.
func (o *GetAverageClickPositionResponse) SetClickCount(v int32) *GetAverageClickPositionResponse {
	o.ClickCount = v
	return o
}

// GetDates returns the Dates field value.
func (o *GetAverageClickPositionResponse) GetDates() []AverageClickEvent {
	if o == nil {
		var ret []AverageClickEvent
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *GetAverageClickPositionResponse) GetDatesOk() ([]AverageClickEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dates, true
}

// SetDates sets field value.
func (o *GetAverageClickPositionResponse) SetDates(v []AverageClickEvent) *GetAverageClickPositionResponse {
	o.Dates = v
	return o
}

func (o GetAverageClickPositionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["average"] = o.Average
	}
	if true {
		toSerialize["clickCount"] = o.ClickCount
	}
	if true {
		toSerialize["dates"] = o.Dates
	}
	return json.Marshal(toSerialize)
}

func (o GetAverageClickPositionResponse) String() string {
	out := ""
	out += fmt.Sprintf("  average=%v\n", o.Average)
	out += fmt.Sprintf("  clickCount=%v\n", o.ClickCount)
	out += fmt.Sprintf("  dates=%v\n", o.Dates)
	return fmt.Sprintf("GetAverageClickPositionResponse {\n%s}", out)
}

type NullableGetAverageClickPositionResponse struct {
	value *GetAverageClickPositionResponse
	isSet bool
}

func (v NullableGetAverageClickPositionResponse) Get() *GetAverageClickPositionResponse {
	return v.value
}

func (v *NullableGetAverageClickPositionResponse) Set(val *GetAverageClickPositionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAverageClickPositionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAverageClickPositionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAverageClickPositionResponse(val *GetAverageClickPositionResponse) *NullableGetAverageClickPositionResponse {
	return &NullableGetAverageClickPositionResponse{value: val, isSet: true}
}

func (v NullableGetAverageClickPositionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAverageClickPositionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
