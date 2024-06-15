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

// NetworksNetworkIdSwitchAccessPoliciesRadiusServers struct for NetworksNetworkIdSwitchAccessPoliciesRadiusServers
type NetworksNetworkIdSwitchAccessPoliciesRadiusServers struct {
	// Unique ID of the RADIUS server
	ServerId *string `json:"serverId,omitempty"`
	// Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
	OrganizationRadiusServerId *string `json:"organizationRadiusServerId,omitempty"`
	// Public IP address of the RADIUS server
	Host *string `json:"host,omitempty"`
	// UDP port that the RADIUS server listens on for access requests
	Port *int32 `json:"port,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers() *NetworksNetworkIdSwitchAccessPoliciesRadiusServers {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusServers{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusServersWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusServersWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusServers {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusServers{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetServerId() string {
	if o == nil || isNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetServerIdOk() (*string, bool) {
	if o == nil || isNil(o.ServerId) {
    return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) HasServerId() bool {
	if o != nil && !isNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) SetServerId(v string) {
	o.ServerId = &v
}

// GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetOrganizationRadiusServerId() string {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
		var ret string
		return ret
	}
	return *o.OrganizationRadiusServerId
}

// GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetOrganizationRadiusServerIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
    return nil, false
	}
	return o.OrganizationRadiusServerId, true
}

// HasOrganizationRadiusServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) HasOrganizationRadiusServerId() bool {
	if o != nil && !isNil(o.OrganizationRadiusServerId) {
		return true
	}

	return false
}

// SetOrganizationRadiusServerId gets a reference to the given string and assigns it to the OrganizationRadiusServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) SetOrganizationRadiusServerId(v string) {
	o.OrganizationRadiusServerId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) SetPort(v int32) {
	o.Port = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadiusServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !isNil(o.OrganizationRadiusServerId) {
		toSerialize["organizationRadiusServerId"] = o.OrganizationRadiusServerId
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadiusServers
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers) Get() *NetworksNetworkIdSwitchAccessPoliciesRadiusServers {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers(val *NetworksNetworkIdSwitchAccessPoliciesRadiusServers) *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


