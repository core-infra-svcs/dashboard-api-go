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

// InlineResponse200339RedundancyManagementAddresses struct for InlineResponse200339RedundancyManagementAddresses
type InlineResponse200339RedundancyManagementAddresses struct {
	// Wireless LAN controller redundancy management interface ip address
	Address *string `json:"address,omitempty"`
}

// NewInlineResponse200339RedundancyManagementAddresses instantiates a new InlineResponse200339RedundancyManagementAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339RedundancyManagementAddresses() *InlineResponse200339RedundancyManagementAddresses {
	this := InlineResponse200339RedundancyManagementAddresses{}
	return &this
}

// NewInlineResponse200339RedundancyManagementAddressesWithDefaults instantiates a new InlineResponse200339RedundancyManagementAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339RedundancyManagementAddressesWithDefaults() *InlineResponse200339RedundancyManagementAddresses {
	this := InlineResponse200339RedundancyManagementAddresses{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse200339RedundancyManagementAddresses) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339RedundancyManagementAddresses) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse200339RedundancyManagementAddresses) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse200339RedundancyManagementAddresses) SetAddress(v string) {
	o.Address = &v
}

func (o InlineResponse200339RedundancyManagementAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339RedundancyManagementAddresses struct {
	value *InlineResponse200339RedundancyManagementAddresses
	isSet bool
}

func (v NullableInlineResponse200339RedundancyManagementAddresses) Get() *InlineResponse200339RedundancyManagementAddresses {
	return v.value
}

func (v *NullableInlineResponse200339RedundancyManagementAddresses) Set(val *InlineResponse200339RedundancyManagementAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339RedundancyManagementAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339RedundancyManagementAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339RedundancyManagementAddresses(val *InlineResponse200339RedundancyManagementAddresses) *NullableInlineResponse200339RedundancyManagementAddresses {
	return &NullableInlineResponse200339RedundancyManagementAddresses{value: val, isSet: true}
}

func (v NullableInlineResponse200339RedundancyManagementAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339RedundancyManagementAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


