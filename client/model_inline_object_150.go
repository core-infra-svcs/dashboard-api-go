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

// InlineObject150 struct for InlineObject150
type InlineObject150 struct {
	// The IP address of the interface where the RP needs to be created.
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject150 instantiates a new InlineObject150 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject150(interfaceIp string, multicastGroup string) *InlineObject150 {
	this := InlineObject150{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject150WithDefaults instantiates a new InlineObject150 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject150WithDefaults() *InlineObject150 {
	this := InlineObject150{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject150) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject150) GetInterfaceIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject150) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject150) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject150) GetMulticastGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject150) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject150) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject150 struct {
	value *InlineObject150
	isSet bool
}

func (v NullableInlineObject150) Get() *InlineObject150 {
	return v.value
}

func (v *NullableInlineObject150) Set(val *InlineObject150) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject150) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject150) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject150(val *InlineObject150) *NullableInlineObject150 {
	return &NullableInlineObject150{value: val, isSet: true}
}

func (v NullableInlineObject150) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject150) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


