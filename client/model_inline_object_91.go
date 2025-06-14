/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject91 struct for InlineObject91
type InlineObject91 struct {
	// The array of clients to provision
	Clients []NetworksNetworkIdClientsProvisionClients `json:"clients"`
	// The policy to apply to the specified client. Can be 'Group policy', 'Allowed', 'Blocked', 'Per connection' or 'Normal'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to \"Group policy\". Otherwise this is ignored.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	PoliciesBySecurityAppliance *NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance `json:"policiesBySecurityAppliance,omitempty"`
	PoliciesBySsid *NetworksNetworkIdClientsProvisionPoliciesBySsid `json:"policiesBySsid,omitempty"`
}

// NewInlineObject91 instantiates a new InlineObject91 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject91(clients []NetworksNetworkIdClientsProvisionClients, devicePolicy string) *InlineObject91 {
	this := InlineObject91{}
	this.Clients = clients
	this.DevicePolicy = devicePolicy
	return &this
}

// NewInlineObject91WithDefaults instantiates a new InlineObject91 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject91WithDefaults() *InlineObject91 {
	this := InlineObject91{}
	return &this
}

// GetClients returns the Clients field value
func (o *InlineObject91) GetClients() []NetworksNetworkIdClientsProvisionClients {
	if o == nil {
		var ret []NetworksNetworkIdClientsProvisionClients
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetClientsOk() ([]NetworksNetworkIdClientsProvisionClients, bool) {
	if o == nil {
    return nil, false
	}
	return o.Clients, true
}

// SetClients sets field value
func (o *InlineObject91) SetClients(v []NetworksNetworkIdClientsProvisionClients) {
	o.Clients = v
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *InlineObject91) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *InlineObject91) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineObject91) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineObject91) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineObject91) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPoliciesBySecurityAppliance returns the PoliciesBySecurityAppliance field value if set, zero value otherwise.
func (o *InlineObject91) GetPoliciesBySecurityAppliance() NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance {
	if o == nil || isNil(o.PoliciesBySecurityAppliance) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance
		return ret
	}
	return *o.PoliciesBySecurityAppliance
}

// GetPoliciesBySecurityApplianceOk returns a tuple with the PoliciesBySecurityAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetPoliciesBySecurityApplianceOk() (*NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance, bool) {
	if o == nil || isNil(o.PoliciesBySecurityAppliance) {
    return nil, false
	}
	return o.PoliciesBySecurityAppliance, true
}

// HasPoliciesBySecurityAppliance returns a boolean if a field has been set.
func (o *InlineObject91) HasPoliciesBySecurityAppliance() bool {
	if o != nil && !isNil(o.PoliciesBySecurityAppliance) {
		return true
	}

	return false
}

// SetPoliciesBySecurityAppliance gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance and assigns it to the PoliciesBySecurityAppliance field.
func (o *InlineObject91) SetPoliciesBySecurityAppliance(v NetworksNetworkIdClientsProvisionPoliciesBySecurityAppliance) {
	o.PoliciesBySecurityAppliance = &v
}

// GetPoliciesBySsid returns the PoliciesBySsid field value if set, zero value otherwise.
func (o *InlineObject91) GetPoliciesBySsid() NetworksNetworkIdClientsProvisionPoliciesBySsid {
	if o == nil || isNil(o.PoliciesBySsid) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid
		return ret
	}
	return *o.PoliciesBySsid
}

// GetPoliciesBySsidOk returns a tuple with the PoliciesBySsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetPoliciesBySsidOk() (*NetworksNetworkIdClientsProvisionPoliciesBySsid, bool) {
	if o == nil || isNil(o.PoliciesBySsid) {
    return nil, false
	}
	return o.PoliciesBySsid, true
}

// HasPoliciesBySsid returns a boolean if a field has been set.
func (o *InlineObject91) HasPoliciesBySsid() bool {
	if o != nil && !isNil(o.PoliciesBySsid) {
		return true
	}

	return false
}

// SetPoliciesBySsid gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid and assigns it to the PoliciesBySsid field.
func (o *InlineObject91) SetPoliciesBySsid(v NetworksNetworkIdClientsProvisionPoliciesBySsid) {
	o.PoliciesBySsid = &v
}

func (o InlineObject91) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject91 struct {
	value *InlineObject91
	isSet bool
}

func (v NullableInlineObject91) Get() *InlineObject91 {
	return v.value
}

func (v *NullableInlineObject91) Set(val *InlineObject91) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject91) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject91) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject91(val *InlineObject91) *NullableInlineObject91 {
	return &NullableInlineObject91{value: val, isSet: true}
}

func (v NullableInlineObject91) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject91) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


