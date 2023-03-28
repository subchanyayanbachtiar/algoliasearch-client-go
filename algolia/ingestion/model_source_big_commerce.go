// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SourceBigCommerce struct for SourceBigCommerce
type SourceBigCommerce struct {
	// The store hash identifying the store the shopper is signing in to.
	StoreHash *string             `json:"storeHash,omitempty"`
	Channel   *BigCommerceChannel `json:"channel,omitempty"`
}

type SourceBigCommerceOption func(f *SourceBigCommerce)

func WithSourceBigCommerceStoreHash(val string) SourceBigCommerceOption {
	return func(f *SourceBigCommerce) {
		f.StoreHash = &val
	}
}

func WithSourceBigCommerceChannel(val BigCommerceChannel) SourceBigCommerceOption {
	return func(f *SourceBigCommerce) {
		f.Channel = &val
	}
}

// NewSourceBigCommerce instantiates a new SourceBigCommerce object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceBigCommerce(opts ...SourceBigCommerceOption) *SourceBigCommerce {
	this := &SourceBigCommerce{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewSourceBigCommerceWithDefaults instantiates a new SourceBigCommerce object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceBigCommerceWithDefaults() *SourceBigCommerce {
	this := &SourceBigCommerce{}
	return this
}

// GetStoreHash returns the StoreHash field value if set, zero value otherwise.
func (o *SourceBigCommerce) GetStoreHash() string {
	if o == nil || o.StoreHash == nil {
		var ret string
		return ret
	}
	return *o.StoreHash
}

// GetStoreHashOk returns a tuple with the StoreHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBigCommerce) GetStoreHashOk() (*string, bool) {
	if o == nil || o.StoreHash == nil {
		return nil, false
	}
	return o.StoreHash, true
}

// HasStoreHash returns a boolean if a field has been set.
func (o *SourceBigCommerce) HasStoreHash() bool {
	if o != nil && o.StoreHash != nil {
		return true
	}

	return false
}

// SetStoreHash gets a reference to the given string and assigns it to the StoreHash field.
func (o *SourceBigCommerce) SetStoreHash(v string) {
	o.StoreHash = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *SourceBigCommerce) GetChannel() BigCommerceChannel {
	if o == nil || o.Channel == nil {
		var ret BigCommerceChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBigCommerce) GetChannelOk() (*BigCommerceChannel, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *SourceBigCommerce) HasChannel() bool {
	if o != nil && o.Channel != nil {
		return true
	}

	return false
}

// SetChannel gets a reference to the given BigCommerceChannel and assigns it to the Channel field.
func (o *SourceBigCommerce) SetChannel(v BigCommerceChannel) {
	o.Channel = &v
}

func (o SourceBigCommerce) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.StoreHash != nil {
		toSerialize["storeHash"] = o.StoreHash
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	return json.Marshal(toSerialize)
}

func (o SourceBigCommerce) String() string {
	out := "SourceBigCommerce {\n"
	out += fmt.Sprintf("  storeHash=%v\n", o.StoreHash)
	out += fmt.Sprintf("  channel=%v\n", o.Channel)
	out += "}"
	return out
}

type NullableSourceBigCommerce struct {
	value *SourceBigCommerce
	isSet bool
}

func (v NullableSourceBigCommerce) Get() *SourceBigCommerce {
	return v.value
}

func (v *NullableSourceBigCommerce) Set(val *SourceBigCommerce) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceBigCommerce) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceBigCommerce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceBigCommerce(val *SourceBigCommerce) *NullableSourceBigCommerce {
	return &NullableSourceBigCommerce{value: val, isSet: true}
}

func (v NullableSourceBigCommerce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceBigCommerce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
