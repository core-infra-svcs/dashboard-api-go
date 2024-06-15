/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpServerPolicyAlerts Alert settings for DHCP servers
type NetworksNetworkIdSwitchDhcpServerPolicyAlerts struct {
	Email *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail `json:"email,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyAlerts instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpServerPolicyAlerts() *NetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	this := NetworksNetworkIdSwitchDhcpServerPolicyAlerts{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyAlertsWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpServerPolicyAlertsWithDefaults() *NetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	this := NetworksNetworkIdSwitchDhcpServerPolicyAlerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlerts) GetEmail() NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlerts) GetEmailOk() (*NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail and assigns it to the Email field.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlerts) SetEmail(v NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) {
	o.Email = &v
}

func (o NetworksNetworkIdSwitchDhcpServerPolicyAlerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts struct {
	value *NetworksNetworkIdSwitchDhcpServerPolicyAlerts
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts) Get() *NetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts) Set(val *NetworksNetworkIdSwitchDhcpServerPolicyAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts(val *NetworksNetworkIdSwitchDhcpServerPolicyAlerts) *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	return &NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


