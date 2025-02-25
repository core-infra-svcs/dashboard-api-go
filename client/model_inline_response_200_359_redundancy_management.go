/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200359RedundancyManagement Wireless LAN controller redundancy management interface information
type InlineResponse200359RedundancyManagement struct {
	// Wireless LAN controller redundancy management interface addresses
	Addresses []InlineResponse200359RedundancyManagementAddresses `json:"addresses,omitempty"`
}

// NewInlineResponse200359RedundancyManagement instantiates a new InlineResponse200359RedundancyManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200359RedundancyManagement() *InlineResponse200359RedundancyManagement {
	this := InlineResponse200359RedundancyManagement{}
	return &this
}

// NewInlineResponse200359RedundancyManagementWithDefaults instantiates a new InlineResponse200359RedundancyManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200359RedundancyManagementWithDefaults() *InlineResponse200359RedundancyManagement {
	this := InlineResponse200359RedundancyManagement{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineResponse200359RedundancyManagement) GetAddresses() []InlineResponse200359RedundancyManagementAddresses {
	if o == nil || isNil(o.Addresses) {
		var ret []InlineResponse200359RedundancyManagementAddresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359RedundancyManagement) GetAddressesOk() ([]InlineResponse200359RedundancyManagementAddresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineResponse200359RedundancyManagement) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []InlineResponse200359RedundancyManagementAddresses and assigns it to the Addresses field.
func (o *InlineResponse200359RedundancyManagement) SetAddresses(v []InlineResponse200359RedundancyManagementAddresses) {
	o.Addresses = v
}

func (o InlineResponse200359RedundancyManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200359RedundancyManagement struct {
	value *InlineResponse200359RedundancyManagement
	isSet bool
}

func (v NullableInlineResponse200359RedundancyManagement) Get() *InlineResponse200359RedundancyManagement {
	return v.value
}

func (v *NullableInlineResponse200359RedundancyManagement) Set(val *InlineResponse200359RedundancyManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200359RedundancyManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200359RedundancyManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200359RedundancyManagement(val *InlineResponse200359RedundancyManagement) *NullableInlineResponse200359RedundancyManagement {
	return &NullableInlineResponse200359RedundancyManagement{value: val, isSet: true}
}

func (v NullableInlineResponse200359RedundancyManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200359RedundancyManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


