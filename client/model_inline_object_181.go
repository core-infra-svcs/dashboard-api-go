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

// InlineObject181 struct for InlineObject181
type InlineObject181 struct {
	// AP port profile name
	Name string `json:"name"`
	// AP ports configuration
	Ports []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 `json:"ports"`
	// AP usb ports configuration
	UsbPorts []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1 `json:"usbPorts,omitempty"`
}

// NewInlineObject181 instantiates a new InlineObject181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject181(name string, ports []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) *InlineObject181 {
	this := InlineObject181{}
	this.Name = name
	this.Ports = ports
	return &this
}

// NewInlineObject181WithDefaults instantiates a new InlineObject181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject181WithDefaults() *InlineObject181 {
	this := InlineObject181{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject181) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject181) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject181) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value
func (o *InlineObject181) GetPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1 {
	if o == nil {
		var ret []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *InlineObject181) GetPortsOk() ([]NetworksNetworkIdWirelessEthernetPortsProfilesPorts1, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *InlineObject181) SetPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1) {
	o.Ports = v
}

// GetUsbPorts returns the UsbPorts field value if set, zero value otherwise.
func (o *InlineObject181) GetUsbPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1 {
	if o == nil || isNil(o.UsbPorts) {
		var ret []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1
		return ret
	}
	return o.UsbPorts
}

// GetUsbPortsOk returns a tuple with the UsbPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject181) GetUsbPortsOk() ([]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1, bool) {
	if o == nil || isNil(o.UsbPorts) {
    return nil, false
	}
	return o.UsbPorts, true
}

// HasUsbPorts returns a boolean if a field has been set.
func (o *InlineObject181) HasUsbPorts() bool {
	if o != nil && !isNil(o.UsbPorts) {
		return true
	}

	return false
}

// SetUsbPorts gets a reference to the given []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1 and assigns it to the UsbPorts field.
func (o *InlineObject181) SetUsbPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1) {
	o.UsbPorts = v
}

func (o InlineObject181) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.UsbPorts) {
		toSerialize["usbPorts"] = o.UsbPorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject181 struct {
	value *InlineObject181
	isSet bool
}

func (v NullableInlineObject181) Get() *InlineObject181 {
	return v.value
}

func (v *NullableInlineObject181) Set(val *InlineObject181) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject181) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject181) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject181(val *InlineObject181) *NullableInlineObject181 {
	return &NullableInlineObject181{value: val, isSet: true}
}

func (v NullableInlineObject181) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject181) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


