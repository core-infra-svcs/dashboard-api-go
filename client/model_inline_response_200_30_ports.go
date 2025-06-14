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

// InlineResponse20030Ports lldp and/or cdp information, keyed by port
type InlineResponse20030Ports struct {
	Lldp *InlineResponse20030Lldp `json:"lldp,omitempty"`
	Cdp *InlineResponse20030Cdp `json:"cdp,omitempty"`
	// MAC address for the device
	DeviceMac *string `json:"deviceMac,omitempty"`
	Device *InlineResponse20030Device `json:"device,omitempty"`
}

// NewInlineResponse20030Ports instantiates a new InlineResponse20030Ports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030Ports() *InlineResponse20030Ports {
	this := InlineResponse20030Ports{}
	return &this
}

// NewInlineResponse20030PortsWithDefaults instantiates a new InlineResponse20030Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030PortsWithDefaults() *InlineResponse20030Ports {
	this := InlineResponse20030Ports{}
	return &this
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *InlineResponse20030Ports) GetLldp() InlineResponse20030Lldp {
	if o == nil || isNil(o.Lldp) {
		var ret InlineResponse20030Lldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Ports) GetLldpOk() (*InlineResponse20030Lldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *InlineResponse20030Ports) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given InlineResponse20030Lldp and assigns it to the Lldp field.
func (o *InlineResponse20030Ports) SetLldp(v InlineResponse20030Lldp) {
	o.Lldp = &v
}

// GetCdp returns the Cdp field value if set, zero value otherwise.
func (o *InlineResponse20030Ports) GetCdp() InlineResponse20030Cdp {
	if o == nil || isNil(o.Cdp) {
		var ret InlineResponse20030Cdp
		return ret
	}
	return *o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Ports) GetCdpOk() (*InlineResponse20030Cdp, bool) {
	if o == nil || isNil(o.Cdp) {
    return nil, false
	}
	return o.Cdp, true
}

// HasCdp returns a boolean if a field has been set.
func (o *InlineResponse20030Ports) HasCdp() bool {
	if o != nil && !isNil(o.Cdp) {
		return true
	}

	return false
}

// SetCdp gets a reference to the given InlineResponse20030Cdp and assigns it to the Cdp field.
func (o *InlineResponse20030Ports) SetCdp(v InlineResponse20030Cdp) {
	o.Cdp = &v
}

// GetDeviceMac returns the DeviceMac field value if set, zero value otherwise.
func (o *InlineResponse20030Ports) GetDeviceMac() string {
	if o == nil || isNil(o.DeviceMac) {
		var ret string
		return ret
	}
	return *o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Ports) GetDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.DeviceMac) {
    return nil, false
	}
	return o.DeviceMac, true
}

// HasDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse20030Ports) HasDeviceMac() bool {
	if o != nil && !isNil(o.DeviceMac) {
		return true
	}

	return false
}

// SetDeviceMac gets a reference to the given string and assigns it to the DeviceMac field.
func (o *InlineResponse20030Ports) SetDeviceMac(v string) {
	o.DeviceMac = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse20030Ports) GetDevice() InlineResponse20030Device {
	if o == nil || isNil(o.Device) {
		var ret InlineResponse20030Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030Ports) GetDeviceOk() (*InlineResponse20030Device, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse20030Ports) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given InlineResponse20030Device and assigns it to the Device field.
func (o *InlineResponse20030Ports) SetDevice(v InlineResponse20030Device) {
	o.Device = &v
}

func (o InlineResponse20030Ports) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	if !isNil(o.Cdp) {
		toSerialize["cdp"] = o.Cdp
	}
	if !isNil(o.DeviceMac) {
		toSerialize["deviceMac"] = o.DeviceMac
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030Ports struct {
	value *InlineResponse20030Ports
	isSet bool
}

func (v NullableInlineResponse20030Ports) Get() *InlineResponse20030Ports {
	return v.value
}

func (v *NullableInlineResponse20030Ports) Set(val *InlineResponse20030Ports) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030Ports) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030Ports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030Ports(val *InlineResponse20030Ports) *NullableInlineResponse20030Ports {
	return &NullableInlineResponse20030Ports{value: val, isSet: true}
}

func (v NullableInlineResponse20030Ports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030Ports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


