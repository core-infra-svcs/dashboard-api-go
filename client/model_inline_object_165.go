/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject165 struct for InlineObject165
type InlineObject165 struct {
	// A list of branding policy IDs arranged in ascending priority order (IDs later in the array have higher priority).
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

// NewInlineObject165 instantiates a new InlineObject165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject165(brandingPolicyIds []string) *InlineObject165 {
	this := InlineObject165{}
	this.BrandingPolicyIds = brandingPolicyIds
	return &this
}

// NewInlineObject165WithDefaults instantiates a new InlineObject165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject165WithDefaults() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// GetBrandingPolicyIds returns the BrandingPolicyIds field value
func (o *InlineObject165) GetBrandingPolicyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BrandingPolicyIds
}

// GetBrandingPolicyIdsOk returns a tuple with the BrandingPolicyIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetBrandingPolicyIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BrandingPolicyIds, true
}

// SetBrandingPolicyIds sets field value
func (o *InlineObject165) SetBrandingPolicyIds(v []string) {
	o.BrandingPolicyIds = v
}

func (o InlineObject165) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brandingPolicyIds"] = o.BrandingPolicyIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject165 struct {
	value *InlineObject165
	isSet bool
}

func (v NullableInlineObject165) Get() *InlineObject165 {
	return v.value
}

func (v *NullableInlineObject165) Set(val *InlineObject165) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject165) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject165(val *InlineObject165) *NullableInlineObject165 {
	return &NullableInlineObject165{value: val, isSet: true}
}

func (v NullableInlineObject165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

