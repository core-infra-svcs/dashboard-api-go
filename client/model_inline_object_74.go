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

// InlineObject74 struct for InlineObject74
type InlineObject74 struct {
	// The array of clients to provision
	Clients []NetworksNetworkIdClientsProvisionClients `json:"clients"`
	// The policy to apply to the specified client. Can be 'Group policy', 'Allowed', 'Blocked', 'Per connection' or 'Normal'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to \"Group policy\". Otherwise this is ignored.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	PoliciesBySecurityAppliance *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance `json:"policiesBySecurityAppliance,omitempty"`
	PoliciesBySsid *NetworksNetworkIdClientsProvisionPoliciesBySsid `json:"policiesBySsid,omitempty"`
}

// NewInlineObject74 instantiates a new InlineObject74 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject74(clients []NetworksNetworkIdClientsProvisionClients, devicePolicy string) *InlineObject74 {
	this := InlineObject74{}
	this.Clients = clients
	this.DevicePolicy = devicePolicy
	return &this
}

// NewInlineObject74WithDefaults instantiates a new InlineObject74 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject74WithDefaults() *InlineObject74 {
	this := InlineObject74{}
	return &this
}

// GetClients returns the Clients field value
func (o *InlineObject74) GetClients() []NetworksNetworkIdClientsProvisionClients {
	if o == nil {
		var ret []NetworksNetworkIdClientsProvisionClients
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetClientsOk() ([]NetworksNetworkIdClientsProvisionClients, bool) {
	if o == nil {
    return nil, false
	}
	return o.Clients, true
}

// SetClients sets field value
func (o *InlineObject74) SetClients(v []NetworksNetworkIdClientsProvisionClients) {
	o.Clients = v
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *InlineObject74) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *InlineObject74) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineObject74) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineObject74) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineObject74) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPoliciesBySecurityAppliance returns the PoliciesBySecurityAppliance field value if set, zero value otherwise.
func (o *InlineObject74) GetPoliciesBySecurityAppliance() NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance {
	if o == nil || isNil(o.PoliciesBySecurityAppliance) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance
		return ret
	}
	return *o.PoliciesBySecurityAppliance
}

// GetPoliciesBySecurityApplianceOk returns a tuple with the PoliciesBySecurityAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetPoliciesBySecurityApplianceOk() (*NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance, bool) {
	if o == nil || isNil(o.PoliciesBySecurityAppliance) {
    return nil, false
	}
	return o.PoliciesBySecurityAppliance, true
}

// HasPoliciesBySecurityAppliance returns a boolean if a field has been set.
func (o *InlineObject74) HasPoliciesBySecurityAppliance() bool {
	if o != nil && !isNil(o.PoliciesBySecurityAppliance) {
		return true
	}

	return false
}

// SetPoliciesBySecurityAppliance gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance and assigns it to the PoliciesBySecurityAppliance field.
func (o *InlineObject74) SetPoliciesBySecurityAppliance(v NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) {
	o.PoliciesBySecurityAppliance = &v
}

// GetPoliciesBySsid returns the PoliciesBySsid field value if set, zero value otherwise.
func (o *InlineObject74) GetPoliciesBySsid() NetworksNetworkIdClientsProvisionPoliciesBySsid {
	if o == nil || isNil(o.PoliciesBySsid) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid
		return ret
	}
	return *o.PoliciesBySsid
}

// GetPoliciesBySsidOk returns a tuple with the PoliciesBySsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetPoliciesBySsidOk() (*NetworksNetworkIdClientsProvisionPoliciesBySsid, bool) {
	if o == nil || isNil(o.PoliciesBySsid) {
    return nil, false
	}
	return o.PoliciesBySsid, true
}

// HasPoliciesBySsid returns a boolean if a field has been set.
func (o *InlineObject74) HasPoliciesBySsid() bool {
	if o != nil && !isNil(o.PoliciesBySsid) {
		return true
	}

	return false
}

// SetPoliciesBySsid gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid and assigns it to the PoliciesBySsid field.
func (o *InlineObject74) SetPoliciesBySsid(v NetworksNetworkIdClientsProvisionPoliciesBySsid) {
	o.PoliciesBySsid = &v
}

func (o InlineObject74) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clients"] = o.Clients
	}
	if true {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.PoliciesBySecurityAppliance) {
		toSerialize["policiesBySecurityAppliance"] = o.PoliciesBySecurityAppliance
	}
	if !isNil(o.PoliciesBySsid) {
		toSerialize["policiesBySsid"] = o.PoliciesBySsid
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject74 struct {
	value *InlineObject74
	isSet bool
}

func (v NullableInlineObject74) Get() *InlineObject74 {
	return v.value
}

func (v *NullableInlineObject74) Set(val *InlineObject74) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject74) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject74) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject74(val *InlineObject74) *NullableInlineObject74 {
	return &NullableInlineObject74{value: val, isSet: true}
}

func (v NullableInlineObject74) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject74) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


