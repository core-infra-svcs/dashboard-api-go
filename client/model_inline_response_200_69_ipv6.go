/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20069Ipv6 Information regarding IPv6 address of the neighbor
type InlineResponse20069Ipv6 struct {
	// The IPv6 address of the neighbor
	Address *string `json:"address,omitempty"`
}

// NewInlineResponse20069Ipv6 instantiates a new InlineResponse20069Ipv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069Ipv6() *InlineResponse20069Ipv6 {
	this := InlineResponse20069Ipv6{}
	return &this
}

// NewInlineResponse20069Ipv6WithDefaults instantiates a new InlineResponse20069Ipv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069Ipv6WithDefaults() *InlineResponse20069Ipv6 {
	this := InlineResponse20069Ipv6{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse20069Ipv6) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Ipv6) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse20069Ipv6) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse20069Ipv6) SetAddress(v string) {
	o.Address = &v
}

func (o InlineResponse20069Ipv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069Ipv6 struct {
	value *InlineResponse20069Ipv6
	isSet bool
}

func (v NullableInlineResponse20069Ipv6) Get() *InlineResponse20069Ipv6 {
	return v.value
}

func (v *NullableInlineResponse20069Ipv6) Set(val *InlineResponse20069Ipv6) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069Ipv6) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069Ipv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069Ipv6(val *InlineResponse20069Ipv6) *NullableInlineResponse20069Ipv6 {
	return &NullableInlineResponse20069Ipv6{value: val, isSet: true}
}

func (v NullableInlineResponse20069Ipv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069Ipv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


