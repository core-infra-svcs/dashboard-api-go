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

// InlineResponse20026Ports lldp and/or cdp information, keyed by port
type InlineResponse20026Ports struct {
	Lldp *InlineResponse20026Lldp `json:"lldp,omitempty"`
	Cdp *InlineResponse20026Cdp `json:"cdp,omitempty"`
	// MAC address for the device
	DeviceMac *string `json:"deviceMac,omitempty"`
	Device *InlineResponse20026Device `json:"device,omitempty"`
}

// NewInlineResponse20026Ports instantiates a new InlineResponse20026Ports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026Ports() *InlineResponse20026Ports {
	this := InlineResponse20026Ports{}
	return &this
}

// NewInlineResponse20026PortsWithDefaults instantiates a new InlineResponse20026Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026PortsWithDefaults() *InlineResponse20026Ports {
	this := InlineResponse20026Ports{}
	return &this
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *InlineResponse20026Ports) GetLldp() InlineResponse20026Lldp {
	if o == nil || isNil(o.Lldp) {
		var ret InlineResponse20026Lldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Ports) GetLldpOk() (*InlineResponse20026Lldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *InlineResponse20026Ports) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given InlineResponse20026Lldp and assigns it to the Lldp field.
func (o *InlineResponse20026Ports) SetLldp(v InlineResponse20026Lldp) {
	o.Lldp = &v
}

// GetCdp returns the Cdp field value if set, zero value otherwise.
func (o *InlineResponse20026Ports) GetCdp() InlineResponse20026Cdp {
	if o == nil || isNil(o.Cdp) {
		var ret InlineResponse20026Cdp
		return ret
	}
	return *o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Ports) GetCdpOk() (*InlineResponse20026Cdp, bool) {
	if o == nil || isNil(o.Cdp) {
    return nil, false
	}
	return o.Cdp, true
}

// HasCdp returns a boolean if a field has been set.
func (o *InlineResponse20026Ports) HasCdp() bool {
	if o != nil && !isNil(o.Cdp) {
		return true
	}

	return false
}

// SetCdp gets a reference to the given InlineResponse20026Cdp and assigns it to the Cdp field.
func (o *InlineResponse20026Ports) SetCdp(v InlineResponse20026Cdp) {
	o.Cdp = &v
}

// GetDeviceMac returns the DeviceMac field value if set, zero value otherwise.
func (o *InlineResponse20026Ports) GetDeviceMac() string {
	if o == nil || isNil(o.DeviceMac) {
		var ret string
		return ret
	}
	return *o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Ports) GetDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.DeviceMac) {
    return nil, false
	}
	return o.DeviceMac, true
}

// HasDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse20026Ports) HasDeviceMac() bool {
	if o != nil && !isNil(o.DeviceMac) {
		return true
	}

	return false
}

// SetDeviceMac gets a reference to the given string and assigns it to the DeviceMac field.
func (o *InlineResponse20026Ports) SetDeviceMac(v string) {
	o.DeviceMac = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse20026Ports) GetDevice() InlineResponse20026Device {
	if o == nil || isNil(o.Device) {
		var ret InlineResponse20026Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Ports) GetDeviceOk() (*InlineResponse20026Device, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse20026Ports) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given InlineResponse20026Device and assigns it to the Device field.
func (o *InlineResponse20026Ports) SetDevice(v InlineResponse20026Device) {
	o.Device = &v
}

func (o InlineResponse20026Ports) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20026Ports struct {
	value *InlineResponse20026Ports
	isSet bool
}

func (v NullableInlineResponse20026Ports) Get() *InlineResponse20026Ports {
	return v.value
}

func (v *NullableInlineResponse20026Ports) Set(val *InlineResponse20026Ports) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026Ports) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026Ports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026Ports(val *InlineResponse20026Ports) *NullableInlineResponse20026Ports {
	return &NullableInlineResponse20026Ports{value: val, isSet: true}
}

func (v NullableInlineResponse20026Ports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026Ports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


