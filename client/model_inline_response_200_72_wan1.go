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

// InlineResponse20072Wan1 WAN 1 IP and subnet
type InlineResponse20072Wan1 struct {
	// IP address used for WAN 1
	Ip *string `json:"ip,omitempty"`
	// Subnet used for WAN 1
	Subnet *string `json:"subnet,omitempty"`
}

// NewInlineResponse20072Wan1 instantiates a new InlineResponse20072Wan1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20072Wan1() *InlineResponse20072Wan1 {
	this := InlineResponse20072Wan1{}
	return &this
}

// NewInlineResponse20072Wan1WithDefaults instantiates a new InlineResponse20072Wan1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20072Wan1WithDefaults() *InlineResponse20072Wan1 {
	this := InlineResponse20072Wan1{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20072Wan1) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Wan1) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20072Wan1) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20072Wan1) SetIp(v string) {
	o.Ip = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20072Wan1) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Wan1) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20072Wan1) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20072Wan1) SetSubnet(v string) {
	o.Subnet = &v
}

func (o InlineResponse20072Wan1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20072Wan1 struct {
	value *InlineResponse20072Wan1
	isSet bool
}

func (v NullableInlineResponse20072Wan1) Get() *InlineResponse20072Wan1 {
	return v.value
}

func (v *NullableInlineResponse20072Wan1) Set(val *InlineResponse20072Wan1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20072Wan1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20072Wan1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20072Wan1(val *InlineResponse20072Wan1) *NullableInlineResponse20072Wan1 {
	return &NullableInlineResponse20072Wan1{value: val, isSet: true}
}

func (v NullableInlineResponse20072Wan1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20072Wan1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


