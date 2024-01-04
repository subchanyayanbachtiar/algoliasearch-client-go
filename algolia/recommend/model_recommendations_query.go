// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RecommendationsQuery struct for RecommendationsQuery.
type RecommendationsQuery struct {
	// Algolia index name.
	IndexName string `json:"indexName"`
	// Recommendations with a confidence score lower than `threshold` won't appear in results. > **Note**: Each recommendation has a confidence score of 0 to 100. The closer the score is to 100, the more relevant the recommendations are.
	Threshold *int32 `json:"threshold,omitempty"`
	// Maximum number of recommendations to retrieve. If 0, all recommendations will be returned.
	MaxRecommendations *int32               `json:"maxRecommendations,omitempty"`
	Model              RecommendationModels `json:"model"`
	// Unique object identifier.
	ObjectID           string              `json:"objectID"`
	QueryParameters    *SearchParamsObject `json:"queryParameters,omitempty"`
	FallbackParameters *SearchParamsObject `json:"fallbackParameters,omitempty"`
}

type RecommendationsQueryOption func(f *RecommendationsQuery)

func WithRecommendationsQueryThreshold(val int32) RecommendationsQueryOption {
	return func(f *RecommendationsQuery) {
		f.Threshold = &val
	}
}

func WithRecommendationsQueryMaxRecommendations(val int32) RecommendationsQueryOption {
	return func(f *RecommendationsQuery) {
		f.MaxRecommendations = &val
	}
}

func WithRecommendationsQueryQueryParameters(val SearchParamsObject) RecommendationsQueryOption {
	return func(f *RecommendationsQuery) {
		f.QueryParameters = &val
	}
}

func WithRecommendationsQueryFallbackParameters(val SearchParamsObject) RecommendationsQueryOption {
	return func(f *RecommendationsQuery) {
		f.FallbackParameters = &val
	}
}

// NewRecommendationsQuery instantiates a new RecommendationsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationsQuery(indexName string, model RecommendationModels, objectID string, opts ...RecommendationsQueryOption) *RecommendationsQuery {
	this := &RecommendationsQuery{}
	this.IndexName = indexName
	this.Model = model
	this.ObjectID = objectID
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewRecommendationsQueryWithDefaults instantiates a new RecommendationsQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationsQueryWithDefaults() *RecommendationsQuery {
	this := &RecommendationsQuery{}
	var maxRecommendations int32 = 0
	this.MaxRecommendations = &maxRecommendations
	return this
}

// GetIndexName returns the IndexName field value.
func (o *RecommendationsQuery) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *RecommendationsQuery) SetIndexName(v string) {
	o.IndexName = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RecommendationsQuery) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RecommendationsQuery) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *RecommendationsQuery) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *RecommendationsQuery) GetMaxRecommendations() int32 {
	if o == nil || o.MaxRecommendations == nil {
		var ret int32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetMaxRecommendationsOk() (*int32, bool) {
	if o == nil || o.MaxRecommendations == nil {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *RecommendationsQuery) HasMaxRecommendations() bool {
	if o != nil && o.MaxRecommendations != nil {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given int32 and assigns it to the MaxRecommendations field.
func (o *RecommendationsQuery) SetMaxRecommendations(v int32) {
	o.MaxRecommendations = &v
}

// GetModel returns the Model field value.
func (o *RecommendationsQuery) GetModel() RecommendationModels {
	if o == nil {
		var ret RecommendationModels
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetModelOk() (*RecommendationModels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *RecommendationsQuery) SetModel(v RecommendationModels) {
	o.Model = v
}

// GetObjectID returns the ObjectID field value.
func (o *RecommendationsQuery) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value.
func (o *RecommendationsQuery) SetObjectID(v string) {
	o.ObjectID = v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *RecommendationsQuery) GetQueryParameters() SearchParamsObject {
	if o == nil || o.QueryParameters == nil {
		var ret SearchParamsObject
		return ret
	}
	return *o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetQueryParametersOk() (*SearchParamsObject, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *RecommendationsQuery) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given SearchParamsObject and assigns it to the QueryParameters field.
func (o *RecommendationsQuery) SetQueryParameters(v SearchParamsObject) {
	o.QueryParameters = &v
}

// GetFallbackParameters returns the FallbackParameters field value if set, zero value otherwise.
func (o *RecommendationsQuery) GetFallbackParameters() SearchParamsObject {
	if o == nil || o.FallbackParameters == nil {
		var ret SearchParamsObject
		return ret
	}
	return *o.FallbackParameters
}

// GetFallbackParametersOk returns a tuple with the FallbackParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsQuery) GetFallbackParametersOk() (*SearchParamsObject, bool) {
	if o == nil || o.FallbackParameters == nil {
		return nil, false
	}
	return o.FallbackParameters, true
}

// HasFallbackParameters returns a boolean if a field has been set.
func (o *RecommendationsQuery) HasFallbackParameters() bool {
	if o != nil && o.FallbackParameters != nil {
		return true
	}

	return false
}

// SetFallbackParameters gets a reference to the given SearchParamsObject and assigns it to the FallbackParameters field.
func (o *RecommendationsQuery) SetFallbackParameters(v SearchParamsObject) {
	o.FallbackParameters = &v
}

func (o RecommendationsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.MaxRecommendations != nil {
		toSerialize["maxRecommendations"] = o.MaxRecommendations
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["objectID"] = o.ObjectID
	}
	if o.QueryParameters != nil {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if o.FallbackParameters != nil {
		toSerialize["fallbackParameters"] = o.FallbackParameters
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationsQuery) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  threshold=%v\n", o.Threshold)
	out += fmt.Sprintf("  maxRecommendations=%v\n", o.MaxRecommendations)
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	out += fmt.Sprintf("  queryParameters=%v\n", o.QueryParameters)
	out += fmt.Sprintf("  fallbackParameters=%v\n", o.FallbackParameters)
	return fmt.Sprintf("RecommendationsQuery {\n%s}", out)
}

type NullableRecommendationsQuery struct {
	value *RecommendationsQuery
	isSet bool
}

func (v NullableRecommendationsQuery) Get() *RecommendationsQuery {
	return v.value
}

func (v *NullableRecommendationsQuery) Set(val *RecommendationsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationsQuery(val *RecommendationsQuery) *NullableRecommendationsQuery {
	return &NullableRecommendationsQuery{value: val, isSet: true}
}

func (v NullableRecommendationsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
