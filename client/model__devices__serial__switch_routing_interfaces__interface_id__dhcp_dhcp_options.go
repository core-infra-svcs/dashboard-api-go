/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions struct for DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions
type DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions struct {
	// The code for DHCP option which should be from 2 to 254
	Code string `json:"code"`
	// The type of the DHCP option which should be one of           ('text', 'ip', 'integer' or 'hex')
	Type string `json:"type"`
	// The value of the DHCP option
	Value string `json:"value"`
}

// NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions instantiates a new DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions(code string, type_ string, value string) *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions {
	this := DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions{}
	this.Code = code
	this.Type = type_
	this.Value = value
	return &this
}

// NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptionsWithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptionsWithDefaults() *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions {
	this := DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions{}
	return &this
}

// GetCode returns the Code field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) GetCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) SetValue(v string) {
	o.Value = v
}

func (o DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions struct {
	value *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions
	isSet bool
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) Get() *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions {
	return v.value
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) Set(val *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions(val *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions {
	return &NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpDhcpOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


