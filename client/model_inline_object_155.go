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

// InlineObject155 struct for InlineObject155
type InlineObject155 struct {
	// The IP address of the interface where the RP needs to be created.
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject155 instantiates a new InlineObject155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject155(interfaceIp string, multicastGroup string) *InlineObject155 {
	this := InlineObject155{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject155WithDefaults instantiates a new InlineObject155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject155WithDefaults() *InlineObject155 {
	this := InlineObject155{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject155) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetInterfaceIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject155) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject155) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetMulticastGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject155) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject155 struct {
	value *InlineObject155
	isSet bool
}

func (v NullableInlineObject155) Get() *InlineObject155 {
	return v.value
}

func (v *NullableInlineObject155) Set(val *InlineObject155) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject155) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject155(val *InlineObject155) *NullableInlineObject155 {
	return &NullableInlineObject155{value: val, isSet: true}
}

func (v NullableInlineObject155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


