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

// InlineResponse200337Addresses struct for InlineResponse200337Addresses
type InlineResponse200337Addresses struct {
	// Type of address for the device uplink. Available options are: ipv4, ipv6. enum = [ipv4, ipv6]
	Protocol *string `json:"protocol,omitempty"`
	// The address of the wireless LAN controller interface
	Address *string `json:"address,omitempty"`
	// The address of the wireless LAN controller interface
	Subnet *string `json:"subnet,omitempty"`
}

// NewInlineResponse200337Addresses instantiates a new InlineResponse200337Addresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200337Addresses() *InlineResponse200337Addresses {
	this := InlineResponse200337Addresses{}
	return &this
}

// NewInlineResponse200337AddressesWithDefaults instantiates a new InlineResponse200337Addresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200337AddressesWithDefaults() *InlineResponse200337Addresses {
	this := InlineResponse200337Addresses{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse200337Addresses) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Addresses) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse200337Addresses) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse200337Addresses) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse200337Addresses) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Addresses) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse200337Addresses) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse200337Addresses) SetAddress(v string) {
	o.Address = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse200337Addresses) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Addresses) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse200337Addresses) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse200337Addresses) SetSubnet(v string) {
	o.Subnet = &v
}

func (o InlineResponse200337Addresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200337Addresses struct {
	value *InlineResponse200337Addresses
	isSet bool
}

func (v NullableInlineResponse200337Addresses) Get() *InlineResponse200337Addresses {
	return v.value
}

func (v *NullableInlineResponse200337Addresses) Set(val *InlineResponse200337Addresses) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200337Addresses) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200337Addresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200337Addresses(val *InlineResponse200337Addresses) *NullableInlineResponse200337Addresses {
	return &NullableInlineResponse200337Addresses{value: val, isSet: true}
}

func (v NullableInlineResponse200337Addresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200337Addresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

