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

// InlineResponse20110Uplinks struct for InlineResponse20110Uplinks
type InlineResponse20110Uplinks struct {
	// Name of the interface
	Interface *string `json:"interface,omitempty"`
	// VLAN of the interface
	Vlan *int32 `json:"vlan,omitempty"`
	// Addresses of the interface
	Addresses []InlineResponse20110Addresses `json:"addresses,omitempty"`
}

// NewInlineResponse20110Uplinks instantiates a new InlineResponse20110Uplinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20110Uplinks() *InlineResponse20110Uplinks {
	this := InlineResponse20110Uplinks{}
	return &this
}

// NewInlineResponse20110UplinksWithDefaults instantiates a new InlineResponse20110Uplinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20110UplinksWithDefaults() *InlineResponse20110Uplinks {
	this := InlineResponse20110Uplinks{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *InlineResponse20110Uplinks) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Uplinks) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *InlineResponse20110Uplinks) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *InlineResponse20110Uplinks) SetInterface(v string) {
	o.Interface = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20110Uplinks) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Uplinks) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20110Uplinks) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse20110Uplinks) SetVlan(v int32) {
	o.Vlan = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineResponse20110Uplinks) GetAddresses() []InlineResponse20110Addresses {
	if o == nil || isNil(o.Addresses) {
		var ret []InlineResponse20110Addresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Uplinks) GetAddressesOk() ([]InlineResponse20110Addresses, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineResponse20110Uplinks) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []InlineResponse20110Addresses and assigns it to the Addresses field.
func (o *InlineResponse20110Uplinks) SetAddresses(v []InlineResponse20110Addresses) {
	o.Addresses = v
}

func (o InlineResponse20110Uplinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20110Uplinks struct {
	value *InlineResponse20110Uplinks
	isSet bool
}

func (v NullableInlineResponse20110Uplinks) Get() *InlineResponse20110Uplinks {
	return v.value
}

func (v *NullableInlineResponse20110Uplinks) Set(val *InlineResponse20110Uplinks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20110Uplinks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20110Uplinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20110Uplinks(val *InlineResponse20110Uplinks) *NullableInlineResponse20110Uplinks {
	return &NullableInlineResponse20110Uplinks{value: val, isSet: true}
}

func (v NullableInlineResponse20110Uplinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20110Uplinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


