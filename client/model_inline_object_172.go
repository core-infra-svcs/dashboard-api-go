/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject172 struct for InlineObject172
type InlineObject172 struct {
	// Allows clients to access rogue networks. Blocked by default.
	DefaultPolicy string `json:"defaultPolicy"`
}

// NewInlineObject172 instantiates a new InlineObject172 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject172(defaultPolicy string) *InlineObject172 {
	this := InlineObject172{}
	this.DefaultPolicy = defaultPolicy
	return &this
}

// NewInlineObject172WithDefaults instantiates a new InlineObject172 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject172WithDefaults() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// GetDefaultPolicy returns the DefaultPolicy field value
func (o *InlineObject172) GetDefaultPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetDefaultPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultPolicy, true
}

// SetDefaultPolicy sets field value
func (o *InlineObject172) SetDefaultPolicy(v string) {
	o.DefaultPolicy = v
}

func (o InlineObject172) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject172 struct {
	value *InlineObject172
	isSet bool
}

func (v NullableInlineObject172) Get() *InlineObject172 {
	return v.value
}

func (v *NullableInlineObject172) Set(val *InlineObject172) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject172) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject172) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject172(val *InlineObject172) *NullableInlineObject172 {
	return &NullableInlineObject172{value: val, isSet: true}
}

func (v NullableInlineObject172) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject172) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


