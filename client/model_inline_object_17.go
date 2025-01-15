/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject17 struct for InlineObject17
type InlineObject17 struct {
	// A list of ports for which to perform the cable test.  For Catalyst switches, IOS interface names are also supported, such as \"GigabitEthernet1/0/8\", \"Gi1/0/8\", or even \"1/0/8\".
	Ports []string `json:"ports"`
	Callback *DevicesSerialLiveToolsArpTableCallback `json:"callback,omitempty"`
}

// NewInlineObject17 instantiates a new InlineObject17 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject17(ports []string) *InlineObject17 {
	this := InlineObject17{}
	this.Ports = ports
	return &this
}

// NewInlineObject17WithDefaults instantiates a new InlineObject17 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject17WithDefaults() *InlineObject17 {
	this := InlineObject17{}
	return &this
}

// GetPorts returns the Ports field value
func (o *InlineObject17) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *InlineObject17) GetPortsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *InlineObject17) SetPorts(v []string) {
	o.Ports = v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject17) GetCallback() DevicesSerialLiveToolsArpTableCallback {
	if o == nil || isNil(o.Callback) {
		var ret DevicesSerialLiveToolsArpTableCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject17) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject17) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given DevicesSerialLiveToolsArpTableCallback and assigns it to the Callback field.
func (o *InlineObject17) SetCallback(v DevicesSerialLiveToolsArpTableCallback) {
	o.Callback = &v
}

func (o InlineObject17) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject17 struct {
	value *InlineObject17
	isSet bool
}

func (v NullableInlineObject17) Get() *InlineObject17 {
	return v.value
}

func (v *NullableInlineObject17) Set(val *InlineObject17) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject17) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject17) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject17(val *InlineObject17) *NullableInlineObject17 {
	return &NullableInlineObject17{value: val, isSet: true}
}

func (v NullableInlineObject17) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject17) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


