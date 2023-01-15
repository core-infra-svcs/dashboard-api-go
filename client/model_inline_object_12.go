/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	// FQDN, IPv4 or IPv6 address
	Target string `json:"target"`
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
}

// NewInlineObject12 instantiates a new InlineObject12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject12(target string) *InlineObject12 {
	this := InlineObject12{}
	this.Target = target
	return &this
}

// NewInlineObject12WithDefaults instantiates a new InlineObject12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject12WithDefaults() *InlineObject12 {
	this := InlineObject12{}
	return &this
}

// GetTarget returns the Target field value
func (o *InlineObject12) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetTargetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *InlineObject12) SetTarget(v string) {
	o.Target = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject12) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject12) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject12) SetCount(v int32) {
	o.Count = &v
}

func (o InlineObject12) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject12 struct {
	value *InlineObject12
	isSet bool
}

func (v NullableInlineObject12) Get() *InlineObject12 {
	return v.value
}

func (v *NullableInlineObject12) Set(val *InlineObject12) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject12) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject12(val *InlineObject12) *NullableInlineObject12 {
	return &NullableInlineObject12{value: val, isSet: true}
}

func (v NullableInlineObject12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


