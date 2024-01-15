// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// ConvertedFilters struct for ConvertedFilters.
type ConvertedFilters struct {
	// The name of the event, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName string          `json:"eventName"`
	EventType ConversionEvent `json:"eventType"`
	// The name of an Algolia index.
	Index string `json:"index"`
	// Facet filters.  Each facet filter string must be URL-encoded, such as, `discount:10%25`.
	Filters []string `json:"filters"`
	// An anonymous or pseudonymous user identifier.  > **Note**: Never include personally identifiable information in user tokens.
	UserToken string `json:"userToken"`
	// An identifier for authenticated users.  > **Note**: Never include personally identifiable information in user tokens.
	AuthenticatedUserToken *string `json:"authenticatedUserToken,omitempty"`
	// The timestamp of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type ConvertedFiltersOption func(f *ConvertedFilters)

func WithConvertedFiltersAuthenticatedUserToken(val string) ConvertedFiltersOption {
	return func(f *ConvertedFilters) {
		f.AuthenticatedUserToken = &val
	}
}

func WithConvertedFiltersTimestamp(val int64) ConvertedFiltersOption {
	return func(f *ConvertedFilters) {
		f.Timestamp = &val
	}
}

// NewConvertedFilters instantiates a new ConvertedFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConvertedFilters(eventName string, eventType ConversionEvent, index string, filters []string, userToken string, opts ...ConvertedFiltersOption) *ConvertedFilters {
	this := &ConvertedFilters{}
	this.EventName = eventName
	this.EventType = eventType
	this.Index = index
	this.Filters = filters
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyConvertedFilters return a pointer to an empty ConvertedFilters object.
func NewEmptyConvertedFilters() *ConvertedFilters {
	return &ConvertedFilters{}
}

// GetEventName returns the EventName field value.
func (o *ConvertedFilters) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value.
func (o *ConvertedFilters) SetEventName(v string) *ConvertedFilters {
	o.EventName = v
	return o
}

// GetEventType returns the EventType field value.
func (o *ConvertedFilters) GetEventType() ConversionEvent {
	if o == nil {
		var ret ConversionEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetEventTypeOk() (*ConversionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *ConvertedFilters) SetEventType(v ConversionEvent) *ConvertedFilters {
	o.EventType = v
	return o
}

// GetIndex returns the Index field value.
func (o *ConvertedFilters) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *ConvertedFilters) SetIndex(v string) *ConvertedFilters {
	o.Index = v
	return o
}

// GetFilters returns the Filters field value.
func (o *ConvertedFilters) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetFiltersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value.
func (o *ConvertedFilters) SetFilters(v []string) *ConvertedFilters {
	o.Filters = v
	return o
}

// GetUserToken returns the UserToken field value.
func (o *ConvertedFilters) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *ConvertedFilters) SetUserToken(v string) *ConvertedFilters {
	o.UserToken = v
	return o
}

// GetAuthenticatedUserToken returns the AuthenticatedUserToken field value if set, zero value otherwise.
func (o *ConvertedFilters) GetAuthenticatedUserToken() string {
	if o == nil || o.AuthenticatedUserToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedUserToken
}

// GetAuthenticatedUserTokenOk returns a tuple with the AuthenticatedUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetAuthenticatedUserTokenOk() (*string, bool) {
	if o == nil || o.AuthenticatedUserToken == nil {
		return nil, false
	}
	return o.AuthenticatedUserToken, true
}

// HasAuthenticatedUserToken returns a boolean if a field has been set.
func (o *ConvertedFilters) HasAuthenticatedUserToken() bool {
	if o != nil && o.AuthenticatedUserToken != nil {
		return true
	}

	return false
}

// SetAuthenticatedUserToken gets a reference to the given string and assigns it to the AuthenticatedUserToken field.
func (o *ConvertedFilters) SetAuthenticatedUserToken(v string) *ConvertedFilters {
	o.AuthenticatedUserToken = &v
	return o
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ConvertedFilters) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedFilters) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ConvertedFilters) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ConvertedFilters) SetTimestamp(v int64) *ConvertedFilters {
	o.Timestamp = &v
	return o
}

func (o ConvertedFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["filters"] = o.Filters
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.AuthenticatedUserToken != nil {
		toSerialize["authenticatedUserToken"] = o.AuthenticatedUserToken
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

func (o ConvertedFilters) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  filters=%v\n", o.Filters)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  authenticatedUserToken=%v\n", o.AuthenticatedUserToken)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	return fmt.Sprintf("ConvertedFilters {\n%s}", out)
}

type NullableConvertedFilters struct {
	value *ConvertedFilters
	isSet bool
}

func (v NullableConvertedFilters) Get() *ConvertedFilters {
	return v.value
}

func (v *NullableConvertedFilters) Set(val *ConvertedFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertedFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertedFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertedFilters(val *ConvertedFilters) *NullableConvertedFilters {
	return &NullableConvertedFilters{value: val, isSet: true}
}

func (v NullableConvertedFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertedFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
