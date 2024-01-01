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

// InlineObject136 struct for InlineObject136
type InlineObject136 struct {
	// The IP address of the interface where the RP needs to be created.
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject136 instantiates a new InlineObject136 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject136(interfaceIp string, multicastGroup string) *InlineObject136 {
	this := InlineObject136{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject136WithDefaults instantiates a new InlineObject136 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject136WithDefaults() *InlineObject136 {
	this := InlineObject136{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject136) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject136) GetInterfaceIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject136) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject136) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject136) GetMulticastGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject136) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject136 struct {
	value *InlineObject136
	isSet bool
}

func (v NullableInlineObject136) Get() *InlineObject136 {
	return v.value
}

func (v *NullableInlineObject136) Set(val *InlineObject136) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject136) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject136) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject136(val *InlineObject136) *NullableInlineObject136 {
	return &NullableInlineObject136{value: val, isSet: true}
}

func (v NullableInlineObject136) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject136) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


