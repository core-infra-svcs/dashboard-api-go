/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 struct for NetworksNetworkIdSwitchAccessPoliciesRadiusServers1
type NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 struct {
	// Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	OrganizationRadiusServerId *string `json:"organizationRadiusServerId,omitempty"`
	// Public IP address of the RADIUS server
	Host *string `json:"host,omitempty"`
	// UDP port that the RADIUS server listens on for access requests
	Port *int32 `json:"port,omitempty"`
	// RADIUS client shared secret
	Secret *string `json:"secret,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers1 instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers1() *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusServers1{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers1WithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers1WithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusServers1{}
	return &this
}

// GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetOrganizationRadiusServerId() string {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
		var ret string
		return ret
	}
	return *o.OrganizationRadiusServerId
}

// GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetOrganizationRadiusServerIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
    return nil, false
	}
	return o.OrganizationRadiusServerId, true
}

// HasOrganizationRadiusServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) HasOrganizationRadiusServerId() bool {
	if o != nil && !isNil(o.OrganizationRadiusServerId) {
		return true
	}

	return false
}

// SetOrganizationRadiusServerId gets a reference to the given string and assigns it to the OrganizationRadiusServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) SetOrganizationRadiusServerId(v string) {
	o.OrganizationRadiusServerId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) SetSecret(v string) {
	o.Secret = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1 struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1) Get() *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1(val *NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1 {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusServers1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


