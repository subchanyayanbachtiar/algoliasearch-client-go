// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package monitoring

import (
	"encoding/json"
	"fmt"
)

// IndexingTimeResponse struct for IndexingTimeResponse.
type IndexingTimeResponse struct {
	Metrics *IndexingTimeResponseMetrics `json:"metrics,omitempty"`
}

type IndexingTimeResponseOption func(f *IndexingTimeResponse)

func WithIndexingTimeResponseMetrics(val IndexingTimeResponseMetrics) IndexingTimeResponseOption {
	return func(f *IndexingTimeResponse) {
		f.Metrics = &val
	}
}

// NewIndexingTimeResponse instantiates a new IndexingTimeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIndexingTimeResponse(opts ...IndexingTimeResponseOption) *IndexingTimeResponse {
	this := &IndexingTimeResponse{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyIndexingTimeResponse return a pointer to an empty IndexingTimeResponse object.
func NewEmptyIndexingTimeResponse() *IndexingTimeResponse {
	return &IndexingTimeResponse{}
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *IndexingTimeResponse) GetMetrics() IndexingTimeResponseMetrics {
	if o == nil || o.Metrics == nil {
		var ret IndexingTimeResponseMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexingTimeResponse) GetMetricsOk() (*IndexingTimeResponseMetrics, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *IndexingTimeResponse) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given IndexingTimeResponseMetrics and assigns it to the Metrics field.
func (o *IndexingTimeResponse) SetMetrics(v *IndexingTimeResponseMetrics) *IndexingTimeResponse {
	o.Metrics = v
	return o
}

func (o IndexingTimeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

func (o IndexingTimeResponse) String() string {
	out := ""
	out += fmt.Sprintf("  metrics=%v\n", o.Metrics)
	return fmt.Sprintf("IndexingTimeResponse {\n%s}", out)
}

type NullableIndexingTimeResponse struct {
	value *IndexingTimeResponse
	isSet bool
}

func (v NullableIndexingTimeResponse) Get() *IndexingTimeResponse {
	return v.value
}

func (v *NullableIndexingTimeResponse) Set(val *IndexingTimeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexingTimeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexingTimeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexingTimeResponse(val *IndexingTimeResponse) *NullableIndexingTimeResponse {
	return &NullableIndexingTimeResponse{value: val, isSet: true}
}

func (v NullableIndexingTimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexingTimeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
