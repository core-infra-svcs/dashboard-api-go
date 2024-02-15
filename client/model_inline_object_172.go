/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject172 struct for InlineObject172
type InlineObject172 struct {
	// AP port profile name
	Name *string `json:"name,omitempty"`
	// AP ports configuration
	Ports []NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts `json:"ports,omitempty"`
	// AP usb ports configuration
	UsbPorts []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1 `json:"usbPorts,omitempty"`
}

// NewInlineObject172 instantiates a new InlineObject172 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject172() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// NewInlineObject172WithDefaults instantiates a new InlineObject172 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject172WithDefaults() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject172) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject172) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject172) SetName(v string) {
	o.Name = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineObject172) GetPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts {
	if o == nil || isNil(o.Ports) {
		var ret []NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetPortsOk() ([]NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineObject172) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts and assigns it to the Ports field.
func (o *InlineObject172) SetPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesProfileIdPorts) {
	o.Ports = v
}

// GetUsbPorts returns the UsbPorts field value if set, zero value otherwise.
func (o *InlineObject172) GetUsbPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1 {
	if o == nil || isNil(o.UsbPorts) {
		var ret []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1
		return ret
	}
	return o.UsbPorts
}

// GetUsbPortsOk returns a tuple with the UsbPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetUsbPortsOk() ([]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1, bool) {
	if o == nil || isNil(o.UsbPorts) {
    return nil, false
	}
	return o.UsbPorts, true
}

// HasUsbPorts returns a boolean if a field has been set.
func (o *InlineObject172) HasUsbPorts() bool {
	if o != nil && !isNil(o.UsbPorts) {
		return true
	}

	return false
}

// SetUsbPorts gets a reference to the given []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1 and assigns it to the UsbPorts field.
func (o *InlineObject172) SetUsbPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1) {
	o.UsbPorts = v
}

func (o InlineObject172) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.UsbPorts) {
		toSerialize["usbPorts"] = o.UsbPorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject172 struct {
	value *InlineObject172
	isSet bool
}

func (v NullableInlineObject172) Get() *InlineObject172 {
	return v.value
}

func (v *NullableInlineObject172) Set(val *InlineObject172) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject172) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject172) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject172(val *InlineObject172) *NullableInlineObject172 {
	return &NullableInlineObject172{value: val, isSet: true}
}

func (v NullableInlineObject172) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject172) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


