/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject127 struct for InlineObject127
type InlineObject127 struct {
	// The IP address of the interface to use
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject127 instantiates a new InlineObject127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject127(interfaceIp string, multicastGroup string) *InlineObject127 {
	this := InlineObject127{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject127WithDefaults instantiates a new InlineObject127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject127WithDefaults() *InlineObject127 {
	this := InlineObject127{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject127) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetInterfaceIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject127) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject127) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetMulticastGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject127) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject127 struct {
	value *InlineObject127
	isSet bool
}

func (v NullableInlineObject127) Get() *InlineObject127 {
	return v.value
}

func (v *NullableInlineObject127) Set(val *InlineObject127) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject127) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject127(val *InlineObject127) *NullableInlineObject127 {
	return &NullableInlineObject127{value: val, isSet: true}
}

func (v NullableInlineObject127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


