/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject138 struct for InlineObject138
type InlineObject138 struct {
	Alerts *NetworksNetworkIdSwitchDhcpServerPolicyAlerts `json:"alerts,omitempty"`
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	AllowedServers []string `json:"allowedServers,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	BlockedServers []string `json:"blockedServers,omitempty"`
	ArpInspection *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`
}

// NewInlineObject138 instantiates a new InlineObject138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject138() *InlineObject138 {
	this := InlineObject138{}
	return &this
}

// NewInlineObject138WithDefaults instantiates a new InlineObject138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject138WithDefaults() *InlineObject138 {
	this := InlineObject138{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineObject138) GetAlerts() NetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	if o == nil || isNil(o.Alerts) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetAlertsOk() (*NetworksNetworkIdSwitchDhcpServerPolicyAlerts, bool) {
	if o == nil || isNil(o.Alerts) {
    return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineObject138) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyAlerts and assigns it to the Alerts field.
func (o *InlineObject138) SetAlerts(v NetworksNetworkIdSwitchDhcpServerPolicyAlerts) {
	o.Alerts = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *InlineObject138) GetDefaultPolicy() string {
	if o == nil || isNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || isNil(o.DefaultPolicy) {
    return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *InlineObject138) HasDefaultPolicy() bool {
	if o != nil && !isNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *InlineObject138) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *InlineObject138) GetAllowedServers() []string {
	if o == nil || isNil(o.AllowedServers) {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetAllowedServersOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedServers) {
    return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *InlineObject138) HasAllowedServers() bool {
	if o != nil && !isNil(o.AllowedServers) {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *InlineObject138) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *InlineObject138) GetBlockedServers() []string {
	if o == nil || isNil(o.BlockedServers) {
		var ret []string
		return ret
	}
	return o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetBlockedServersOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedServers) {
    return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *InlineObject138) HasBlockedServers() bool {
	if o != nil && !isNil(o.BlockedServers) {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *InlineObject138) SetBlockedServers(v []string) {
	o.BlockedServers = v
}

// GetArpInspection returns the ArpInspection field value if set, zero value otherwise.
func (o *InlineObject138) GetArpInspection() NetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	if o == nil || isNil(o.ArpInspection) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspection
		return ret
	}
	return *o.ArpInspection
}

// GetArpInspectionOk returns a tuple with the ArpInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetArpInspectionOk() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspection, bool) {
	if o == nil || isNil(o.ArpInspection) {
    return nil, false
	}
	return o.ArpInspection, true
}

// HasArpInspection returns a boolean if a field has been set.
func (o *InlineObject138) HasArpInspection() bool {
	if o != nil && !isNil(o.ArpInspection) {
		return true
	}

	return false
}

// SetArpInspection gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyArpInspection and assigns it to the ArpInspection field.
func (o *InlineObject138) SetArpInspection(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) {
	o.ArpInspection = &v
}

func (o InlineObject138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !isNil(o.DefaultPolicy) {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	if !isNil(o.AllowedServers) {
		toSerialize["allowedServers"] = o.AllowedServers
	}
	if !isNil(o.BlockedServers) {
		toSerialize["blockedServers"] = o.BlockedServers
	}
	if !isNil(o.ArpInspection) {
		toSerialize["arpInspection"] = o.ArpInspection
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject138 struct {
	value *InlineObject138
	isSet bool
}

func (v NullableInlineObject138) Get() *InlineObject138 {
	return v.value
}

func (v *NullableInlineObject138) Set(val *InlineObject138) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject138) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject138(val *InlineObject138) *NullableInlineObject138 {
	return &NullableInlineObject138{value: val, isSet: true}
}

func (v NullableInlineObject138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


