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

// InlineResponse200287 struct for InlineResponse200287
type InlineResponse200287 struct {
	Downstream *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream `json:"downstream,omitempty"`
	Upstream *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream `json:"upstream,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork `json:"network,omitempty"`
	Device *OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice `json:"device,omitempty"`
}

// NewInlineResponse200287 instantiates a new InlineResponse200287 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200287() *InlineResponse200287 {
	this := InlineResponse200287{}
	return &this
}

// NewInlineResponse200287WithDefaults instantiates a new InlineResponse200287 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200287WithDefaults() *InlineResponse200287 {
	this := InlineResponse200287{}
	return &this
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse200287) GetDownstream() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	if o == nil || isNil(o.Downstream) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetDownstreamOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse200287) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream and assigns it to the Downstream field.
func (o *InlineResponse200287) SetDownstream(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse200287) GetUpstream() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	if o == nil || isNil(o.Upstream) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetUpstreamOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse200287) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream and assigns it to the Upstream field.
func (o *InlineResponse200287) SetUpstream(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) {
	o.Upstream = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200287) GetNetwork() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200287) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork and assigns it to the Network field.
func (o *InlineResponse200287) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) {
	o.Network = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200287) GetDevice() OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetDeviceOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200287) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice and assigns it to the Device field.
func (o *InlineResponse200287) SetDevice(v OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice) {
	o.Device = &v
}

func (o InlineResponse200287) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200287 struct {
	value *InlineResponse200287
	isSet bool
}

func (v NullableInlineResponse200287) Get() *InlineResponse200287 {
	return v.value
}

func (v *NullableInlineResponse200287) Set(val *InlineResponse200287) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200287) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200287) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200287(val *InlineResponse200287) *NullableInlineResponse200287 {
	return &NullableInlineResponse200287{value: val, isSet: true}
}

func (v NullableInlineResponse200287) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200287) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


