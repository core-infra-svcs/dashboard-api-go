/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchRoutingInterfacesIpv6 IPv6 addressing
type DevicesSerialSwitchRoutingInterfacesIpv6 struct {
	// Assignment mode
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// IPv6 address
	Address *string `json:"address,omitempty"`
	// IPv6 subnet
	Prefix *string `json:"prefix,omitempty"`
	// IPv6 gateway
	Gateway *string `json:"gateway,omitempty"`
}

// NewDevicesSerialSwitchRoutingInterfacesIpv6 instantiates a new DevicesSerialSwitchRoutingInterfacesIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchRoutingInterfacesIpv6() *DevicesSerialSwitchRoutingInterfacesIpv6 {
	this := DevicesSerialSwitchRoutingInterfacesIpv6{}
	return &this
}

// NewDevicesSerialSwitchRoutingInterfacesIpv6WithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchRoutingInterfacesIpv6WithDefaults() *DevicesSerialSwitchRoutingInterfacesIpv6 {
	this := DevicesSerialSwitchRoutingInterfacesIpv6{}
	return &this
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetAssignmentMode() string {
	if o == nil || isNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.AssignmentMode) {
    return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) HasAssignmentMode() bool {
	if o != nil && !isNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) SetAddress(v string) {
	o.Address = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) SetPrefix(v string) {
	o.Prefix = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *DevicesSerialSwitchRoutingInterfacesIpv6) SetGateway(v string) {
	o.Gateway = &v
}

func (o DevicesSerialSwitchRoutingInterfacesIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchRoutingInterfacesIpv6 struct {
	value *DevicesSerialSwitchRoutingInterfacesIpv6
	isSet bool
}

func (v NullableDevicesSerialSwitchRoutingInterfacesIpv6) Get() *DevicesSerialSwitchRoutingInterfacesIpv6 {
	return v.value
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesIpv6) Set(val *DevicesSerialSwitchRoutingInterfacesIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchRoutingInterfacesIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchRoutingInterfacesIpv6(val *DevicesSerialSwitchRoutingInterfacesIpv6) *NullableDevicesSerialSwitchRoutingInterfacesIpv6 {
	return &NullableDevicesSerialSwitchRoutingInterfacesIpv6{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchRoutingInterfacesIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


