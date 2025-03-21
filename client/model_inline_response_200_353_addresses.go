/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200353Addresses struct for InlineResponse200353Addresses
type InlineResponse200353Addresses struct {
	// Type of address for the device uplink. Available options are: ipv4, ipv6. enum = [ipv4, ipv6]
	Protocol *string `json:"protocol,omitempty"`
	// The address of the wireless LAN controller interface
	Address *string `json:"address,omitempty"`
	// The address of the wireless LAN controller interface
	Subnet *string `json:"subnet,omitempty"`
}

// NewInlineResponse200353Addresses instantiates a new InlineResponse200353Addresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200353Addresses() *InlineResponse200353Addresses {
	this := InlineResponse200353Addresses{}
	return &this
}

// NewInlineResponse200353AddressesWithDefaults instantiates a new InlineResponse200353Addresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200353AddressesWithDefaults() *InlineResponse200353Addresses {
	this := InlineResponse200353Addresses{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse200353Addresses) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Addresses) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse200353Addresses) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse200353Addresses) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse200353Addresses) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Addresses) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse200353Addresses) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse200353Addresses) SetAddress(v string) {
	o.Address = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse200353Addresses) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Addresses) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse200353Addresses) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse200353Addresses) SetSubnet(v string) {
	o.Subnet = &v
}

func (o InlineResponse200353Addresses) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200353Addresses struct {
	value *InlineResponse200353Addresses
	isSet bool
}

func (v NullableInlineResponse200353Addresses) Get() *InlineResponse200353Addresses {
	return v.value
}

func (v *NullableInlineResponse200353Addresses) Set(val *InlineResponse200353Addresses) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200353Addresses) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200353Addresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200353Addresses(val *InlineResponse200353Addresses) *NullableInlineResponse200353Addresses {
	return &NullableInlineResponse200353Addresses{value: val, isSet: true}
}

func (v NullableInlineResponse200353Addresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200353Addresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


