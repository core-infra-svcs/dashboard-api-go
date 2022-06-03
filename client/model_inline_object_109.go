/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject109 struct for InlineObject109
type InlineObject109 struct {
	// The IP address of the interface to use
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject109 instantiates a new InlineObject109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject109(interfaceIp string, multicastGroup string) *InlineObject109 {
	this := InlineObject109{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject109WithDefaults instantiates a new InlineObject109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject109WithDefaults() *InlineObject109 {
	this := InlineObject109{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject109) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject109) GetInterfaceIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject109) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject109) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject109) GetMulticastGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject109) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject109 struct {
	value *InlineObject109
	isSet bool
}

func (v NullableInlineObject109) Get() *InlineObject109 {
	return v.value
}

func (v *NullableInlineObject109) Set(val *InlineObject109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject109(val *InlineObject109) *NullableInlineObject109 {
	return &NullableInlineObject109{value: val, isSet: true}
}

func (v NullableInlineObject109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


