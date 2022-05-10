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

// InlineObject107 struct for InlineObject107
type InlineObject107 struct {
	// The IP address of the interface to use
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject107 instantiates a new InlineObject107 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject107(interfaceIp string, multicastGroup string) *InlineObject107 {
	this := InlineObject107{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject107WithDefaults instantiates a new InlineObject107 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject107WithDefaults() *InlineObject107 {
	this := InlineObject107{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject107) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject107) GetInterfaceIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject107) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject107) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject107) GetMulticastGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject107) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject107) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject107 struct {
	value *InlineObject107
	isSet bool
}

func (v NullableInlineObject107) Get() *InlineObject107 {
	return v.value
}

func (v *NullableInlineObject107) Set(val *InlineObject107) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject107) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject107) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject107(val *InlineObject107) *NullableInlineObject107 {
	return &NullableInlineObject107{value: val, isSet: true}
}

func (v NullableInlineObject107) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject107) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


