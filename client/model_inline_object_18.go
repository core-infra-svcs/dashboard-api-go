/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject18 struct for InlineObject18
type InlineObject18 struct {
	// FQDN, IPv4 or IPv6 address
	Target string `json:"target"`
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
	Callback *DevicesSerialLiveToolsArpTableCallback `json:"callback,omitempty"`
}

// NewInlineObject18 instantiates a new InlineObject18 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject18(target string) *InlineObject18 {
	this := InlineObject18{}
	this.Target = target
	return &this
}

// NewInlineObject18WithDefaults instantiates a new InlineObject18 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject18WithDefaults() *InlineObject18 {
	this := InlineObject18{}
	return &this
}

// GetTarget returns the Target field value
func (o *InlineObject18) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *InlineObject18) GetTargetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *InlineObject18) SetTarget(v string) {
	o.Target = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject18) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject18) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject18) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject18) SetCount(v int32) {
	o.Count = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject18) GetCallback() DevicesSerialLiveToolsArpTableCallback {
	if o == nil || isNil(o.Callback) {
		var ret DevicesSerialLiveToolsArpTableCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject18) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject18) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given DevicesSerialLiveToolsArpTableCallback and assigns it to the Callback field.
func (o *InlineObject18) SetCallback(v DevicesSerialLiveToolsArpTableCallback) {
	o.Callback = &v
}

func (o InlineObject18) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject18 struct {
	value *InlineObject18
	isSet bool
}

func (v NullableInlineObject18) Get() *InlineObject18 {
	return v.value
}

func (v *NullableInlineObject18) Set(val *InlineObject18) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject18) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject18) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject18(val *InlineObject18) *NullableInlineObject18 {
	return &NullableInlineObject18{value: val, isSet: true}
}

func (v NullableInlineObject18) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject18) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


