/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200170 struct for InlineResponse200170
type InlineResponse200170 struct {
	// The serial number for the device.
	Serial *string `json:"serial,omitempty"`
	// The MAC address of the device.
	Mac *string `json:"mac,omitempty"`
	Network *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`
}

// NewInlineResponse200170 instantiates a new InlineResponse200170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200170() *InlineResponse200170 {
	this := InlineResponse200170{}
	return &this
}

// NewInlineResponse200170WithDefaults instantiates a new InlineResponse200170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200170WithDefaults() *InlineResponse200170 {
	this := InlineResponse200170{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200170) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200170) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200170) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200170) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200170) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200170) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200170) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200170) SetMac(v string) {
	o.Mac = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200170) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200170) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200170) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork and assigns it to the Network field.
func (o *InlineResponse200170) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *InlineResponse200170) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	if o == nil || isNil(o.ByBand) {
		var ret []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200170) GetByBandOk() ([]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool) {
	if o == nil || isNil(o.ByBand) {
    return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *InlineResponse200170) HasByBand() bool {
	if o != nil && !isNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand and assigns it to the ByBand field.
func (o *InlineResponse200170) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) {
	o.ByBand = v
}

func (o InlineResponse200170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200170 struct {
	value *InlineResponse200170
	isSet bool
}

func (v NullableInlineResponse200170) Get() *InlineResponse200170 {
	return v.value
}

func (v *NullableInlineResponse200170) Set(val *InlineResponse200170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200170(val *InlineResponse200170) *NullableInlineResponse200170 {
	return &NullableInlineResponse200170{value: val, isSet: true}
}

func (v NullableInlineResponse200170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


