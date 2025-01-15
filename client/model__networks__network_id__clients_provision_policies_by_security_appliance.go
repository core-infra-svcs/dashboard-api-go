/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance An object, describing what the policy-connection association is for the security appliance. (Only relevant if the security appliance is actually within the network)
type NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance struct {
	// The policy to apply to the specified client. Can be 'Allowed', 'Blocked' or 'Normal'. Required.
	DevicePolicy *string `json:"devicePolicy,omitempty"`
}

// NewNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance() *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance {
	this := NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance{}
	return &this
}

// NewNetworksNetworkIdClientsProvisionPoliciesBySecurityApplianceWithDefaults instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsProvisionPoliciesBySecurityApplianceWithDefaults() *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance {
	this := NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance{}
	return &this
}

// GetDevicePolicy returns the DevicePolicy field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) GetDevicePolicy() string {
	if o == nil || isNil(o.DevicePolicy) {
		var ret string
		return ret
	}
	return *o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) GetDevicePolicyOk() (*string, bool) {
	if o == nil || isNil(o.DevicePolicy) {
    return nil, false
	}
	return o.DevicePolicy, true
}

// HasDevicePolicy returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) HasDevicePolicy() bool {
	if o != nil && !isNil(o.DevicePolicy) {
		return true
	}

	return false
}

// SetDevicePolicy gets a reference to the given string and assigns it to the DevicePolicy field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) SetDevicePolicy(v string) {
	o.DevicePolicy = &v
}

func (o NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DevicePolicy) {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance struct {
	value *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance
	isSet bool
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) Get() *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) Set(val *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance(val *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) *NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance {
	return &NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


