// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// GetApiKeyResponse struct for GetApiKeyResponse
type GetApiKeyResponse struct {
	// The API key.
	Value *string `json:"value,omitempty"`
	// Time of the event expressed in milliseconds since the Unix epoch.
	CreatedAt int64 `json:"createdAt"`
	// Set of permissions associated with the key.
	Acl []Acl `json:"acl"`
	// A comment used to identify a key more easily in the dashboard. It is not interpreted by the API.
	Description *string `json:"description,omitempty"`
	// Restrict this new API key to a list of indices or index patterns. If the list is empty, all indices are allowed.
	Indexes []string `json:"indexes,omitempty"`
	// Maximum number of hits this API key can retrieve in one query. If zero, no limit is enforced.
	MaxHitsPerQuery *int32 `json:"maxHitsPerQuery,omitempty"`
	// Maximum number of API calls per hour allowed from a given IP address or a user token.
	MaxQueriesPerIPPerHour *int32 `json:"maxQueriesPerIPPerHour,omitempty"`
	// URL-encoded query string. Force some query parameters to be applied for each query made with this API key.
	QueryParameters *string `json:"queryParameters,omitempty"`
	// Restrict this new API key to specific referers. If empty or blank, defaults to all referers.
	Referers []string `json:"referers,omitempty"`
	// Validity limit for this key in seconds. The key will automatically be removed after this period of time.
	Validity *int32 `json:"validity,omitempty"`
}

type GetApiKeyResponseOption func(f *GetApiKeyResponse)

func WithGetApiKeyResponseValue(val string) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.Value = &val
	}
}

func WithGetApiKeyResponseDescription(val string) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.Description = &val
	}
}

func WithGetApiKeyResponseIndexes(val []string) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.Indexes = val
	}
}

func WithGetApiKeyResponseMaxHitsPerQuery(val int32) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.MaxHitsPerQuery = &val
	}
}

func WithGetApiKeyResponseMaxQueriesPerIPPerHour(val int32) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.MaxQueriesPerIPPerHour = &val
	}
}

func WithGetApiKeyResponseQueryParameters(val string) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.QueryParameters = &val
	}
}

func WithGetApiKeyResponseReferers(val []string) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.Referers = val
	}
}

func WithGetApiKeyResponseValidity(val int32) GetApiKeyResponseOption {
	return func(f *GetApiKeyResponse) {
		f.Validity = &val
	}
}

// NewGetApiKeyResponse instantiates a new GetApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiKeyResponse(createdAt int64, acl []Acl, opts ...GetApiKeyResponseOption) *GetApiKeyResponse {
	this := &GetApiKeyResponse{}
	this.CreatedAt = createdAt
	this.Acl = acl
	var description string = ""
	this.Description = &description
	var maxHitsPerQuery int32 = 0
	this.MaxHitsPerQuery = &maxHitsPerQuery
	var maxQueriesPerIPPerHour int32 = 0
	this.MaxQueriesPerIPPerHour = &maxQueriesPerIPPerHour
	var queryParameters string = ""
	this.QueryParameters = &queryParameters
	var validity int32 = 0
	this.Validity = &validity
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewGetApiKeyResponseWithDefaults instantiates a new GetApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiKeyResponseWithDefaults() *GetApiKeyResponse {
	this := &GetApiKeyResponse{}
	var description string = ""
	this.Description = &description
	var maxHitsPerQuery int32 = 0
	this.MaxHitsPerQuery = &maxHitsPerQuery
	var maxQueriesPerIPPerHour int32 = 0
	this.MaxQueriesPerIPPerHour = &maxQueriesPerIPPerHour
	var queryParameters string = ""
	this.QueryParameters = &queryParameters
	var validity int32 = 0
	this.Validity = &validity
	return this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetApiKeyResponse) SetValue(v string) {
	o.Value = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetApiKeyResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetApiKeyResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetAcl returns the Acl field value
func (o *GetApiKeyResponse) GetAcl() []Acl {
	if o == nil {
		var ret []Acl
		return ret
	}

	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetAclOk() ([]Acl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Acl, true
}

// SetAcl sets field value
func (o *GetApiKeyResponse) SetAcl(v []Acl) {
	o.Acl = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetApiKeyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetIndexesOk() ([]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasIndexes() bool {
	if o != nil && o.Indexes != nil {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *GetApiKeyResponse) SetIndexes(v []string) {
	o.Indexes = v
}

// GetMaxHitsPerQuery returns the MaxHitsPerQuery field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetMaxHitsPerQuery() int32 {
	if o == nil || o.MaxHitsPerQuery == nil {
		var ret int32
		return ret
	}
	return *o.MaxHitsPerQuery
}

// GetMaxHitsPerQueryOk returns a tuple with the MaxHitsPerQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetMaxHitsPerQueryOk() (*int32, bool) {
	if o == nil || o.MaxHitsPerQuery == nil {
		return nil, false
	}
	return o.MaxHitsPerQuery, true
}

// HasMaxHitsPerQuery returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasMaxHitsPerQuery() bool {
	if o != nil && o.MaxHitsPerQuery != nil {
		return true
	}

	return false
}

// SetMaxHitsPerQuery gets a reference to the given int32 and assigns it to the MaxHitsPerQuery field.
func (o *GetApiKeyResponse) SetMaxHitsPerQuery(v int32) {
	o.MaxHitsPerQuery = &v
}

// GetMaxQueriesPerIPPerHour returns the MaxQueriesPerIPPerHour field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetMaxQueriesPerIPPerHour() int32 {
	if o == nil || o.MaxQueriesPerIPPerHour == nil {
		var ret int32
		return ret
	}
	return *o.MaxQueriesPerIPPerHour
}

// GetMaxQueriesPerIPPerHourOk returns a tuple with the MaxQueriesPerIPPerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetMaxQueriesPerIPPerHourOk() (*int32, bool) {
	if o == nil || o.MaxQueriesPerIPPerHour == nil {
		return nil, false
	}
	return o.MaxQueriesPerIPPerHour, true
}

// HasMaxQueriesPerIPPerHour returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasMaxQueriesPerIPPerHour() bool {
	if o != nil && o.MaxQueriesPerIPPerHour != nil {
		return true
	}

	return false
}

// SetMaxQueriesPerIPPerHour gets a reference to the given int32 and assigns it to the MaxQueriesPerIPPerHour field.
func (o *GetApiKeyResponse) SetMaxQueriesPerIPPerHour(v int32) {
	o.MaxQueriesPerIPPerHour = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetQueryParameters() string {
	if o == nil || o.QueryParameters == nil {
		var ret string
		return ret
	}
	return *o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetQueryParametersOk() (*string, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given string and assigns it to the QueryParameters field.
func (o *GetApiKeyResponse) SetQueryParameters(v string) {
	o.QueryParameters = &v
}

// GetReferers returns the Referers field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetReferers() []string {
	if o == nil || o.Referers == nil {
		var ret []string
		return ret
	}
	return o.Referers
}

// GetReferersOk returns a tuple with the Referers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetReferersOk() ([]string, bool) {
	if o == nil || o.Referers == nil {
		return nil, false
	}
	return o.Referers, true
}

// HasReferers returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasReferers() bool {
	if o != nil && o.Referers != nil {
		return true
	}

	return false
}

// SetReferers gets a reference to the given []string and assigns it to the Referers field.
func (o *GetApiKeyResponse) SetReferers(v []string) {
	o.Referers = v
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *GetApiKeyResponse) GetValidity() int32 {
	if o == nil || o.Validity == nil {
		var ret int32
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeyResponse) GetValidityOk() (*int32, bool) {
	if o == nil || o.Validity == nil {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *GetApiKeyResponse) HasValidity() bool {
	if o != nil && o.Validity != nil {
		return true
	}

	return false
}

// SetValidity gets a reference to the given int32 and assigns it to the Validity field.
func (o *GetApiKeyResponse) SetValidity(v int32) {
	o.Validity = &v
}

func (o GetApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["acl"] = o.Acl
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.MaxHitsPerQuery != nil {
		toSerialize["maxHitsPerQuery"] = o.MaxHitsPerQuery
	}
	if o.MaxQueriesPerIPPerHour != nil {
		toSerialize["maxQueriesPerIPPerHour"] = o.MaxQueriesPerIPPerHour
	}
	if o.QueryParameters != nil {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if o.Referers != nil {
		toSerialize["referers"] = o.Referers
	}
	if o.Validity != nil {
		toSerialize["validity"] = o.Validity
	}
	return json.Marshal(toSerialize)
}

func (o GetApiKeyResponse) String() string {
	out := "GetApiKeyResponse {\n"
	out += fmt.Sprintf("  value=%v\n", o.Value)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	out += fmt.Sprintf("  acl=%v\n", o.Acl)
	out += fmt.Sprintf("  description=%v\n", o.Description)
	out += fmt.Sprintf("  indexes=%v\n", o.Indexes)
	out += fmt.Sprintf("  maxHitsPerQuery=%v\n", o.MaxHitsPerQuery)
	out += fmt.Sprintf("  maxQueriesPerIPPerHour=%v\n", o.MaxQueriesPerIPPerHour)
	out += fmt.Sprintf("  queryParameters=%v\n", o.QueryParameters)
	out += fmt.Sprintf("  referers=%v\n", o.Referers)
	out += fmt.Sprintf("  validity=%v\n", o.Validity)
	out += "}"
	return out
}

type NullableGetApiKeyResponse struct {
	value *GetApiKeyResponse
	isSet bool
}

func (v NullableGetApiKeyResponse) Get() *GetApiKeyResponse {
	return v.value
}

func (v *NullableGetApiKeyResponse) Set(val *GetApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiKeyResponse(val *GetApiKeyResponse) *NullableGetApiKeyResponse {
	return &NullableGetApiKeyResponse{value: val, isSet: true}
}

func (v NullableGetApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
