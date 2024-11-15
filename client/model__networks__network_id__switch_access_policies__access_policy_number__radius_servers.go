/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers struct for NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers
type NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers struct {
	// Unique ID of the RADIUS server. When provided, the existing RADIUS server will be updated instead of creating a new one
	ServerId *string `json:"serverId,omitempty"`
	// Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	OrganizationRadiusServerId *string `json:"organizationRadiusServerId,omitempty"`
	// Public IP address of the RADIUS server
	Host *string `json:"host,omitempty"`
	// UDP port that the RADIUS server listens on for access requests
	Port *int32 `json:"port,omitempty"`
	// RADIUS client shared secret
	Secret *string `json:"secret,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers instantiates a new NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers() *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers {
	this := NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServersWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServersWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers {
	this := NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetServerId() string {
	if o == nil || isNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetServerIdOk() (*string, bool) {
	if o == nil || isNil(o.ServerId) {
    return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) HasServerId() bool {
	if o != nil && !isNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) SetServerId(v string) {
	o.ServerId = &v
}

// GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetOrganizationRadiusServerId() string {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
		var ret string
		return ret
	}
	return *o.OrganizationRadiusServerId
}

// GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetOrganizationRadiusServerIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
    return nil, false
	}
	return o.OrganizationRadiusServerId, true
}

// HasOrganizationRadiusServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) HasOrganizationRadiusServerId() bool {
	if o != nil && !isNil(o.OrganizationRadiusServerId) {
		return true
	}

	return false
}

// SetOrganizationRadiusServerId gets a reference to the given string and assigns it to the OrganizationRadiusServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) SetOrganizationRadiusServerId(v string) {
	o.OrganizationRadiusServerId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) SetSecret(v string) {
	o.Secret = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers struct {
	value *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) Get() *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) Set(val *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers(val *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) *NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


