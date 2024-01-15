// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package analytics

import (
	"encoding/json"
	"fmt"
)

// ConversionRateEvent struct for ConversionRateEvent.
type ConversionRateEvent struct {
	// [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).
	Rate float64 `json:"rate"`
	// Number of tracked searches. This is the number of search requests where the `clickAnalytics` parameter is `true`.
	TrackedSearchCount int32 `json:"trackedSearchCount"`
	// Number of converted clicks.
	ConversionCount int32 `json:"conversionCount"`
	// Date of the event in the format YYYY-MM-DD.
	Date string `json:"date"`
}

// NewConversionRateEvent instantiates a new ConversionRateEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConversionRateEvent(rate float64, trackedSearchCount int32, conversionCount int32, date string) *ConversionRateEvent {
	this := &ConversionRateEvent{}
	this.Rate = rate
	this.TrackedSearchCount = trackedSearchCount
	this.ConversionCount = conversionCount
	this.Date = date
	return this
}

// NewEmptyConversionRateEvent return a pointer to an empty ConversionRateEvent object.
func NewEmptyConversionRateEvent() *ConversionRateEvent {
	return &ConversionRateEvent{}
}

// GetRate returns the Rate field value.
func (o *ConversionRateEvent) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *ConversionRateEvent) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value.
func (o *ConversionRateEvent) SetRate(v float64) *ConversionRateEvent {
	o.Rate = v
	return o
}

// GetTrackedSearchCount returns the TrackedSearchCount field value.
func (o *ConversionRateEvent) GetTrackedSearchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrackedSearchCount
}

// GetTrackedSearchCountOk returns a tuple with the TrackedSearchCount field value
// and a boolean to check if the value has been set.
func (o *ConversionRateEvent) GetTrackedSearchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackedSearchCount, true
}

// SetTrackedSearchCount sets field value.
func (o *ConversionRateEvent) SetTrackedSearchCount(v int32) *ConversionRateEvent {
	o.TrackedSearchCount = v
	return o
}

// GetConversionCount returns the ConversionCount field value.
func (o *ConversionRateEvent) GetConversionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConversionCount
}

// GetConversionCountOk returns a tuple with the ConversionCount field value
// and a boolean to check if the value has been set.
func (o *ConversionRateEvent) GetConversionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionCount, true
}

// SetConversionCount sets field value.
func (o *ConversionRateEvent) SetConversionCount(v int32) *ConversionRateEvent {
	o.ConversionCount = v
	return o
}

// GetDate returns the Date field value.
func (o *ConversionRateEvent) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ConversionRateEvent) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *ConversionRateEvent) SetDate(v string) *ConversionRateEvent {
	o.Date = v
	return o
}

func (o ConversionRateEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if true {
		toSerialize["trackedSearchCount"] = o.TrackedSearchCount
	}
	if true {
		toSerialize["conversionCount"] = o.ConversionCount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

func (o ConversionRateEvent) String() string {
	out := ""
	out += fmt.Sprintf("  rate=%v\n", o.Rate)
	out += fmt.Sprintf("  trackedSearchCount=%v\n", o.TrackedSearchCount)
	out += fmt.Sprintf("  conversionCount=%v\n", o.ConversionCount)
	out += fmt.Sprintf("  date=%v\n", o.Date)
	return fmt.Sprintf("ConversionRateEvent {\n%s}", out)
}

type NullableConversionRateEvent struct {
	value *ConversionRateEvent
	isSet bool
}

func (v NullableConversionRateEvent) Get() *ConversionRateEvent {
	return v.value
}

func (v *NullableConversionRateEvent) Set(val *ConversionRateEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableConversionRateEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableConversionRateEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversionRateEvent(val *ConversionRateEvent) *NullableConversionRateEvent {
	return &NullableConversionRateEvent{value: val, isSet: true}
}

func (v NullableConversionRateEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversionRateEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
