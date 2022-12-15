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

// NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies struct for NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies
type NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	// The device type. Can be one of 'Android', 'BlackBerry', 'Chrome OS', 'iPad', 'iPhone', 'iPod', 'Mac OS X', 'Windows', 'Windows Phone', 'B&N Nook' or 'Other OS'
	DeviceType string `json:"deviceType"`
	// The device policy. Can be one of 'Allowed', 'Blocked' or 'Group policy'
	DevicePolicy string `json:"devicePolicy"`
	// ID of the group policy object.
	GroupPolicyId *int32 `json:"groupPolicyId,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies instantiates a new NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies(deviceType string, devicePolicy string) *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	this := NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies{}
	this.DeviceType = deviceType
	this.DevicePolicy = devicePolicy
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePoliciesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePoliciesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	this := NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies{}
	return &this
}

// GetDeviceType returns the DeviceType field value
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) GetGroupPolicyId() int32 {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret int32
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) GetGroupPolicyIdOk() (*int32, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given int32 and assigns it to the GroupPolicyId field.
func (o *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) SetGroupPolicyId(v int32) {
	o.GroupPolicyId = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deviceType"] = o.DeviceType
	}
	if true {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	value *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) Get() *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) Set(val *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies(val *NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) *NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	return &NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


