/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200185 struct for InlineResponse200185
type InlineResponse200185 struct {
	Network *OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand `json:"byBand,omitempty"`
}

// NewInlineResponse200185 instantiates a new InlineResponse200185 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200185() *InlineResponse200185 {
	this := InlineResponse200185{}
	return &this
}

// NewInlineResponse200185WithDefaults instantiates a new InlineResponse200185 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200185WithDefaults() *InlineResponse200185 {
	this := InlineResponse200185{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200185) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200185) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork and assigns it to the Network field.
func (o *InlineResponse200185) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *InlineResponse200185) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand {
	if o == nil || isNil(o.ByBand) {
		var ret []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetByBandOk() ([]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool) {
	if o == nil || isNil(o.ByBand) {
    return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *InlineResponse200185) HasByBand() bool {
	if o != nil && !isNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand and assigns it to the ByBand field.
func (o *InlineResponse200185) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand) {
	o.ByBand = v
}

func (o InlineResponse200185) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200185 struct {
	value *InlineResponse200185
	isSet bool
}

func (v NullableInlineResponse200185) Get() *InlineResponse200185 {
	return v.value
}

func (v *NullableInlineResponse200185) Set(val *InlineResponse200185) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200185) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200185) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200185(val *InlineResponse200185) *NullableInlineResponse200185 {
	return &NullableInlineResponse200185{value: val, isSet: true}
}

func (v NullableInlineResponse200185) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200185) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

