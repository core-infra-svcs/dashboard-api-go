/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	// The target's VLAN (1 to 4094)
	VlanId int32 `json:"vlanId"`
	// The target's MAC address
	Mac string `json:"mac"`
	Callback *DevicesSerialLiveToolsArpTableCallback `json:"callback,omitempty"`
}

// NewInlineObject23 instantiates a new InlineObject23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject23(vlanId int32, mac string) *InlineObject23 {
	this := InlineObject23{}
	this.VlanId = vlanId
	this.Mac = mac
	return &this
}

// NewInlineObject23WithDefaults instantiates a new InlineObject23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject23WithDefaults() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// GetVlanId returns the VlanId field value
func (o *InlineObject23) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetVlanIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *InlineObject23) SetVlanId(v int32) {
	o.VlanId = v
}

// GetMac returns the Mac field value
func (o *InlineObject23) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *InlineObject23) SetMac(v string) {
	o.Mac = v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject23) GetCallback() DevicesSerialLiveToolsArpTableCallback {
	if o == nil || isNil(o.Callback) {
		var ret DevicesSerialLiveToolsArpTableCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject23) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given DevicesSerialLiveToolsArpTableCallback and assigns it to the Callback field.
func (o *InlineObject23) SetCallback(v DevicesSerialLiveToolsArpTableCallback) {
	o.Callback = &v
}

func (o InlineObject23) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vlanId"] = o.VlanId
	}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject23 struct {
	value *InlineObject23
	isSet bool
}

func (v NullableInlineObject23) Get() *InlineObject23 {
	return v.value
}

func (v *NullableInlineObject23) Set(val *InlineObject23) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject23) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject23(val *InlineObject23) *NullableInlineObject23 {
	return &NullableInlineObject23{value: val, isSet: true}
}

func (v NullableInlineObject23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


