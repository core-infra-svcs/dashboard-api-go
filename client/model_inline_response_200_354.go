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

// InlineResponse200354 struct for InlineResponse200354
type InlineResponse200354 struct {
	Downstream *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream `json:"downstream,omitempty"`
	Upstream *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream `json:"upstream,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork `json:"network,omitempty"`
	Device *OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice `json:"device,omitempty"`
}

// NewInlineResponse200354 instantiates a new InlineResponse200354 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200354() *InlineResponse200354 {
	this := InlineResponse200354{}
	return &this
}

// NewInlineResponse200354WithDefaults instantiates a new InlineResponse200354 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200354WithDefaults() *InlineResponse200354 {
	this := InlineResponse200354{}
	return &this
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse200354) GetDownstream() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream {
	if o == nil || isNil(o.Downstream) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354) GetDownstreamOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse200354) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream and assigns it to the Downstream field.
func (o *InlineResponse200354) SetDownstream(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientDownstream) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse200354) GetUpstream() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	if o == nil || isNil(o.Upstream) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354) GetUpstreamOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse200354) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream and assigns it to the Upstream field.
func (o *InlineResponse200354) SetUpstream(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) {
	o.Upstream = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200354) GetNetwork() OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200354) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork and assigns it to the Network field.
func (o *InlineResponse200354) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesPacketLossByClientNetwork) {
	o.Network = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200354) GetDevice() OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354) GetDeviceOk() (*OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200354) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice and assigns it to the Device field.
func (o *InlineResponse200354) SetDevice(v OrganizationsOrganizationIdWirelessDevicesPacketLossByDeviceDevice) {
	o.Device = &v
}

func (o InlineResponse200354) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200354 struct {
	value *InlineResponse200354
	isSet bool
}

func (v NullableInlineResponse200354) Get() *InlineResponse200354 {
	return v.value
}

func (v *NullableInlineResponse200354) Set(val *InlineResponse200354) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200354) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200354) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200354(val *InlineResponse200354) *NullableInlineResponse200354 {
	return &NullableInlineResponse200354{value: val, isSet: true}
}

func (v NullableInlineResponse200354) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200354) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


