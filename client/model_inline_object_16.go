/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	// FQDN, IPv4 or IPv6 address
	Target string `json:"target"`
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
	Callback *DevicesSerialLiveToolsPingCallback `json:"callback,omitempty"`
}

// NewInlineObject16 instantiates a new InlineObject16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16(target string) *InlineObject16 {
	this := InlineObject16{}
	this.Target = target
	return &this
}

// NewInlineObject16WithDefaults instantiates a new InlineObject16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16WithDefaults() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// GetTarget returns the Target field value
func (o *InlineObject16) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetTargetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *InlineObject16) SetTarget(v string) {
	o.Target = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject16) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject16) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject16) SetCount(v int32) {
	o.Count = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject16) GetCallback() DevicesSerialLiveToolsPingCallback {
	if o == nil || isNil(o.Callback) {
		var ret DevicesSerialLiveToolsPingCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetCallbackOk() (*DevicesSerialLiveToolsPingCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject16) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given DevicesSerialLiveToolsPingCallback and assigns it to the Callback field.
func (o *InlineObject16) SetCallback(v DevicesSerialLiveToolsPingCallback) {
	o.Callback = &v
}

func (o InlineObject16) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject16 struct {
	value *InlineObject16
	isSet bool
}

func (v NullableInlineObject16) Get() *InlineObject16 {
	return v.value
}

func (v *NullableInlineObject16) Set(val *InlineObject16) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16(val *InlineObject16) *NullableInlineObject16 {
	return &NullableInlineObject16{value: val, isSet: true}
}

func (v NullableInlineObject16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


