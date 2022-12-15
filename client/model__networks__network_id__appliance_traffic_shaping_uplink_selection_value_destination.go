/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination Destination of this custom type traffic filter
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address, or \"any\". E.g.: \"192.168.10.0/24\",  \"192.168.10.1\" (same as \"192.168.10.1/32\"), \"0.0.0.0/0\" (same as \"any\")
	Cidr *string `json:"cidr,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestinationWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestinationWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) SetCidr(v string) {
	o.Cidr = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


