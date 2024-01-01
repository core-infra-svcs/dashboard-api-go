/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions struct for NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions struct {
	// Option name.
	Name *string `json:"name,omitempty"`
	// Option value.
	Value *string `json:"value,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptionsWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptionsWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) SetValue(v string) {
	o.Value = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFieldsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


