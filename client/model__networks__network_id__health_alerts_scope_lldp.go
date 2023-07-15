/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdHealthAlertsScopeLldp Lldp information
type NetworksNetworkIdHealthAlertsScopeLldp struct {
	// Port Id
	PortId *string `json:"portId,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScopeLldp instantiates a new NetworksNetworkIdHealthAlertsScopeLldp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScopeLldp() *NetworksNetworkIdHealthAlertsScopeLldp {
	this := NetworksNetworkIdHealthAlertsScopeLldp{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopeLldpWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopeLldp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopeLldpWithDefaults() *NetworksNetworkIdHealthAlertsScopeLldp {
	this := NetworksNetworkIdHealthAlertsScopeLldp{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeLldp) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeLldp) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeLldp) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *NetworksNetworkIdHealthAlertsScopeLldp) SetPortId(v string) {
	o.PortId = &v
}

func (o NetworksNetworkIdHealthAlertsScopeLldp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScopeLldp struct {
	value *NetworksNetworkIdHealthAlertsScopeLldp
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScopeLldp) Get() *NetworksNetworkIdHealthAlertsScopeLldp {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeLldp) Set(val *NetworksNetworkIdHealthAlertsScopeLldp) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScopeLldp) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeLldp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScopeLldp(val *NetworksNetworkIdHealthAlertsScopeLldp) *NullableNetworksNetworkIdHealthAlertsScopeLldp {
	return &NullableNetworksNetworkIdHealthAlertsScopeLldp{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScopeLldp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeLldp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


