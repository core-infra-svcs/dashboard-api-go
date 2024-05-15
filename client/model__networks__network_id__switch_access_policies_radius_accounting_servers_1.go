/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 struct for NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1
type NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 struct {
	// unique ID of the RADIUS accounting server
	ServerId *string `json:"serverId,omitempty"`
	// Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored
	OrganizationRadiusServerId *string `json:"organizationRadiusServerId,omitempty"`
	// Public IP address of the RADIUS accounting server
	Host *string `json:"host,omitempty"`
	// UDP port that the RADIUS Accounting server listens on for access requests
	Port *int32 `json:"port,omitempty"`
	// RADIUS client shared secret
	Secret *string `json:"secret,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetServerId() string {
	if o == nil || isNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetServerIdOk() (*string, bool) {
	if o == nil || isNil(o.ServerId) {
    return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasServerId() bool {
	if o != nil && !isNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetServerId(v string) {
	o.ServerId = &v
}

// GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetOrganizationRadiusServerId() string {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
		var ret string
		return ret
	}
	return *o.OrganizationRadiusServerId
}

// GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetOrganizationRadiusServerIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationRadiusServerId) {
    return nil, false
	}
	return o.OrganizationRadiusServerId, true
}

// HasOrganizationRadiusServerId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasOrganizationRadiusServerId() bool {
	if o != nil && !isNil(o.OrganizationRadiusServerId) {
		return true
	}

	return false
}

// SetOrganizationRadiusServerId gets a reference to the given string and assigns it to the OrganizationRadiusServerId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetOrganizationRadiusServerId(v string) {
	o.OrganizationRadiusServerId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetSecret(v string) {
	o.Secret = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) Get() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1(val *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


