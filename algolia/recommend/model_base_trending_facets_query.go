// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// BaseTrendingFacetsQuery struct for BaseTrendingFacetsQuery
type BaseTrendingFacetsQuery struct {
	// Facet name for trending models.
	FacetName string               `json:"facetName" validate:"required"`
	Model     *TrendingFacetsModel `json:"model,omitempty"`
}

type BaseTrendingFacetsQueryOption func(f *BaseTrendingFacetsQuery)

func WithBaseTrendingFacetsQueryModel(val TrendingFacetsModel) BaseTrendingFacetsQueryOption {
	return func(f *BaseTrendingFacetsQuery) {
		f.Model = &val
	}
}

// NewBaseTrendingFacetsQuery instantiates a new BaseTrendingFacetsQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTrendingFacetsQuery(facetName string, opts ...BaseTrendingFacetsQueryOption) *BaseTrendingFacetsQuery {
	this := &BaseTrendingFacetsQuery{}
	this.FacetName = facetName
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewBaseTrendingFacetsQueryWithDefaults instantiates a new BaseTrendingFacetsQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTrendingFacetsQueryWithDefaults() *BaseTrendingFacetsQuery {
	this := &BaseTrendingFacetsQuery{}
	return this
}

// GetFacetName returns the FacetName field value
func (o *BaseTrendingFacetsQuery) GetFacetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
func (o *BaseTrendingFacetsQuery) GetFacetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetName, true
}

// SetFacetName sets field value
func (o *BaseTrendingFacetsQuery) SetFacetName(v string) {
	o.FacetName = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *BaseTrendingFacetsQuery) GetModel() TrendingFacetsModel {
	if o == nil || o.Model == nil {
		var ret TrendingFacetsModel
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTrendingFacetsQuery) GetModelOk() (*TrendingFacetsModel, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *BaseTrendingFacetsQuery) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given TrendingFacetsModel and assigns it to the Model field.
func (o *BaseTrendingFacetsQuery) SetModel(v TrendingFacetsModel) {
	o.Model = &v
}

func (o BaseTrendingFacetsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["facetName"] = o.FacetName
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (o BaseTrendingFacetsQuery) String() string {
	out := ""
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	out += fmt.Sprintf("  model=%v\n", o.Model)
	return fmt.Sprintf("BaseTrendingFacetsQuery {\n%s}", out)
}

type NullableBaseTrendingFacetsQuery struct {
	value *BaseTrendingFacetsQuery
	isSet bool
}

func (v NullableBaseTrendingFacetsQuery) Get() *BaseTrendingFacetsQuery {
	return v.value
}

func (v *NullableBaseTrendingFacetsQuery) Set(val *BaseTrendingFacetsQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTrendingFacetsQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTrendingFacetsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTrendingFacetsQuery(val *BaseTrendingFacetsQuery) *NullableBaseTrendingFacetsQuery {
	return &NullableBaseTrendingFacetsQuery{value: val, isSet: true}
}

func (v NullableBaseTrendingFacetsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTrendingFacetsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
