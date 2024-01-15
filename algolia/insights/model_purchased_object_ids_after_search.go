// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// PurchasedObjectIDsAfterSearch Use this event to track when users make a purchase after a previous Algolia request. If you're building your category pages with Algolia, you'll also use this event.
type PurchasedObjectIDsAfterSearch struct {
	// The name of the event, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName    string          `json:"eventName"`
	EventType    ConversionEvent `json:"eventType"`
	EventSubtype PurchaseEvent   `json:"eventSubtype"`
	// The name of an Algolia index.
	Index string `json:"index"`
	// The object IDs of the records that are part of the event.
	ObjectIDs []string `json:"objectIDs"`
	// An anonymous or pseudonymous user identifier.  > **Note**: Never include personally identifiable information in user tokens.
	UserToken string `json:"userToken"`
	// An identifier for authenticated users.  > **Note**: Never include personally identifiable information in user tokens.
	AuthenticatedUserToken *string `json:"authenticatedUserToken,omitempty"`
	// Three-letter [currency code](https://www.iso.org/iso-4217-currency-codes.html).
	Currency *string `json:"currency,omitempty"`
	// Extra information about the records involved in a purchase or add-to-cart events.  If provided, it must be the same length as `objectIDs`.
	ObjectData []ObjectDataAfterSearch `json:"objectData,omitempty"`
	// The timestamp of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
	Value     *Value `json:"value,omitempty"`
}

type PurchasedObjectIDsAfterSearchOption func(f *PurchasedObjectIDsAfterSearch)

func WithPurchasedObjectIDsAfterSearchAuthenticatedUserToken(val string) PurchasedObjectIDsAfterSearchOption {
	return func(f *PurchasedObjectIDsAfterSearch) {
		f.AuthenticatedUserToken = &val
	}
}

func WithPurchasedObjectIDsAfterSearchCurrency(val string) PurchasedObjectIDsAfterSearchOption {
	return func(f *PurchasedObjectIDsAfterSearch) {
		f.Currency = &val
	}
}

func WithPurchasedObjectIDsAfterSearchObjectData(val []ObjectDataAfterSearch) PurchasedObjectIDsAfterSearchOption {
	return func(f *PurchasedObjectIDsAfterSearch) {
		f.ObjectData = val
	}
}

func WithPurchasedObjectIDsAfterSearchTimestamp(val int64) PurchasedObjectIDsAfterSearchOption {
	return func(f *PurchasedObjectIDsAfterSearch) {
		f.Timestamp = &val
	}
}

func WithPurchasedObjectIDsAfterSearchValue(val Value) PurchasedObjectIDsAfterSearchOption {
	return func(f *PurchasedObjectIDsAfterSearch) {
		f.Value = &val
	}
}

// NewPurchasedObjectIDsAfterSearch instantiates a new PurchasedObjectIDsAfterSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPurchasedObjectIDsAfterSearch(eventName string, eventType ConversionEvent, eventSubtype PurchaseEvent, index string, objectIDs []string, userToken string, opts ...PurchasedObjectIDsAfterSearchOption) *PurchasedObjectIDsAfterSearch {
	this := &PurchasedObjectIDsAfterSearch{}
	this.EventName = eventName
	this.EventType = eventType
	this.EventSubtype = eventSubtype
	this.Index = index
	this.ObjectIDs = objectIDs
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewPurchasedObjectIDsAfterSearchWithDefaults instantiates a new PurchasedObjectIDsAfterSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPurchasedObjectIDsAfterSearchWithDefaults() *PurchasedObjectIDsAfterSearch {
	this := &PurchasedObjectIDsAfterSearch{}
	return this
}

// GetEventName returns the EventName field value.
func (o *PurchasedObjectIDsAfterSearch) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value.
func (o *PurchasedObjectIDsAfterSearch) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value.
func (o *PurchasedObjectIDsAfterSearch) GetEventType() ConversionEvent {
	if o == nil {
		var ret ConversionEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetEventTypeOk() (*ConversionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *PurchasedObjectIDsAfterSearch) SetEventType(v ConversionEvent) {
	o.EventType = v
}

// GetEventSubtype returns the EventSubtype field value.
func (o *PurchasedObjectIDsAfterSearch) GetEventSubtype() PurchaseEvent {
	if o == nil {
		var ret PurchaseEvent
		return ret
	}

	return o.EventSubtype
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetEventSubtypeOk() (*PurchaseEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSubtype, true
}

// SetEventSubtype sets field value.
func (o *PurchasedObjectIDsAfterSearch) SetEventSubtype(v PurchaseEvent) {
	o.EventSubtype = v
}

// GetIndex returns the Index field value.
func (o *PurchasedObjectIDsAfterSearch) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *PurchasedObjectIDsAfterSearch) SetIndex(v string) {
	o.Index = v
}

// GetObjectIDs returns the ObjectIDs field value.
func (o *PurchasedObjectIDsAfterSearch) GetObjectIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectIDs
}

// GetObjectIDsOk returns a tuple with the ObjectIDs field value
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetObjectIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectIDs, true
}

// SetObjectIDs sets field value.
func (o *PurchasedObjectIDsAfterSearch) SetObjectIDs(v []string) {
	o.ObjectIDs = v
}

// GetUserToken returns the UserToken field value.
func (o *PurchasedObjectIDsAfterSearch) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *PurchasedObjectIDsAfterSearch) SetUserToken(v string) {
	o.UserToken = v
}

// GetAuthenticatedUserToken returns the AuthenticatedUserToken field value if set, zero value otherwise.
func (o *PurchasedObjectIDsAfterSearch) GetAuthenticatedUserToken() string {
	if o == nil || o.AuthenticatedUserToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedUserToken
}

// GetAuthenticatedUserTokenOk returns a tuple with the AuthenticatedUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetAuthenticatedUserTokenOk() (*string, bool) {
	if o == nil || o.AuthenticatedUserToken == nil {
		return nil, false
	}
	return o.AuthenticatedUserToken, true
}

// HasAuthenticatedUserToken returns a boolean if a field has been set.
func (o *PurchasedObjectIDsAfterSearch) HasAuthenticatedUserToken() bool {
	if o != nil && o.AuthenticatedUserToken != nil {
		return true
	}

	return false
}

// SetAuthenticatedUserToken gets a reference to the given string and assigns it to the AuthenticatedUserToken field.
func (o *PurchasedObjectIDsAfterSearch) SetAuthenticatedUserToken(v string) {
	o.AuthenticatedUserToken = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PurchasedObjectIDsAfterSearch) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PurchasedObjectIDsAfterSearch) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PurchasedObjectIDsAfterSearch) SetCurrency(v string) {
	o.Currency = &v
}

// GetObjectData returns the ObjectData field value if set, zero value otherwise.
func (o *PurchasedObjectIDsAfterSearch) GetObjectData() []ObjectDataAfterSearch {
	if o == nil || o.ObjectData == nil {
		var ret []ObjectDataAfterSearch
		return ret
	}
	return o.ObjectData
}

// GetObjectDataOk returns a tuple with the ObjectData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetObjectDataOk() ([]ObjectDataAfterSearch, bool) {
	if o == nil || o.ObjectData == nil {
		return nil, false
	}
	return o.ObjectData, true
}

// HasObjectData returns a boolean if a field has been set.
func (o *PurchasedObjectIDsAfterSearch) HasObjectData() bool {
	if o != nil && o.ObjectData != nil {
		return true
	}

	return false
}

// SetObjectData gets a reference to the given []ObjectDataAfterSearch and assigns it to the ObjectData field.
func (o *PurchasedObjectIDsAfterSearch) SetObjectData(v []ObjectDataAfterSearch) {
	o.ObjectData = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PurchasedObjectIDsAfterSearch) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PurchasedObjectIDsAfterSearch) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *PurchasedObjectIDsAfterSearch) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PurchasedObjectIDsAfterSearch) GetValue() Value {
	if o == nil || o.Value == nil {
		var ret Value
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchasedObjectIDsAfterSearch) GetValueOk() (*Value, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PurchasedObjectIDsAfterSearch) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given Value and assigns it to the Value field.
func (o *PurchasedObjectIDsAfterSearch) SetValue(v Value) {
	o.Value = &v
}

func (o PurchasedObjectIDsAfterSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["eventSubtype"] = o.EventSubtype
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["objectIDs"] = o.ObjectIDs
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.AuthenticatedUserToken != nil {
		toSerialize["authenticatedUserToken"] = o.AuthenticatedUserToken
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.ObjectData != nil {
		toSerialize["objectData"] = o.ObjectData
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

func (o PurchasedObjectIDsAfterSearch) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  eventSubtype=%v\n", o.EventSubtype)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  objectIDs=%v\n", o.ObjectIDs)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  authenticatedUserToken=%v\n", o.AuthenticatedUserToken)
	out += fmt.Sprintf("  currency=%v\n", o.Currency)
	out += fmt.Sprintf("  objectData=%v\n", o.ObjectData)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	out += fmt.Sprintf("  value=%v\n", o.Value)
	return fmt.Sprintf("PurchasedObjectIDsAfterSearch {\n%s}", out)
}

type NullablePurchasedObjectIDsAfterSearch struct {
	value *PurchasedObjectIDsAfterSearch
	isSet bool
}

func (v NullablePurchasedObjectIDsAfterSearch) Get() *PurchasedObjectIDsAfterSearch {
	return v.value
}

func (v *NullablePurchasedObjectIDsAfterSearch) Set(val *PurchasedObjectIDsAfterSearch) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchasedObjectIDsAfterSearch) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchasedObjectIDsAfterSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchasedObjectIDsAfterSearch(val *PurchasedObjectIDsAfterSearch) *NullablePurchasedObjectIDsAfterSearch {
	return &NullablePurchasedObjectIDsAfterSearch{value: val, isSet: true}
}

func (v NullablePurchasedObjectIDsAfterSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchasedObjectIDsAfterSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
