/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20069Wan2 WAN 2 IP and subnet
type InlineResponse20069Wan2 struct {
	// IP address used for WAN 2
	Ip *string `json:"ip,omitempty"`
	// Subnet used for WAN 2
	Subnet *string `json:"subnet,omitempty"`
}

// NewInlineResponse20069Wan2 instantiates a new InlineResponse20069Wan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069Wan2() *InlineResponse20069Wan2 {
	this := InlineResponse20069Wan2{}
	return &this
}

// NewInlineResponse20069Wan2WithDefaults instantiates a new InlineResponse20069Wan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069Wan2WithDefaults() *InlineResponse20069Wan2 {
	this := InlineResponse20069Wan2{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20069Wan2) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Wan2) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20069Wan2) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20069Wan2) SetIp(v string) {
	o.Ip = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20069Wan2) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Wan2) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20069Wan2) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20069Wan2) SetSubnet(v string) {
	o.Subnet = &v
}

func (o InlineResponse20069Wan2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069Wan2 struct {
	value *InlineResponse20069Wan2
	isSet bool
}

func (v NullableInlineResponse20069Wan2) Get() *InlineResponse20069Wan2 {
	return v.value
}

func (v *NullableInlineResponse20069Wan2) Set(val *InlineResponse20069Wan2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069Wan2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069Wan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069Wan2(val *InlineResponse20069Wan2) *NullableInlineResponse20069Wan2 {
	return &NullableInlineResponse20069Wan2{value: val, isSet: true}
}

func (v NullableInlineResponse20069Wan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069Wan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


