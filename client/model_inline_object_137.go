/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject137 struct for InlineObject137
type InlineObject137 struct {
	Alerts *NetworksNetworkIdSwitchDhcpServerPolicyAlerts `json:"alerts,omitempty"`
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	AllowedServers []string `json:"allowedServers,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	BlockedServers []string `json:"blockedServers,omitempty"`
	ArpInspection *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`
}

// NewInlineObject137 instantiates a new InlineObject137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject137() *InlineObject137 {
	this := InlineObject137{}
	return &this
}

// NewInlineObject137WithDefaults instantiates a new InlineObject137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject137WithDefaults() *InlineObject137 {
	this := InlineObject137{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineObject137) GetAlerts() NetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	if o == nil || isNil(o.Alerts) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject137) GetAlertsOk() (*NetworksNetworkIdSwitchDhcpServerPolicyAlerts, bool) {
	if o == nil || isNil(o.Alerts) {
    return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineObject137) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyAlerts and assigns it to the Alerts field.
func (o *InlineObject137) SetAlerts(v NetworksNetworkIdSwitchDhcpServerPolicyAlerts) {
	o.Alerts = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *InlineObject137) GetDefaultPolicy() string {
	if o == nil || isNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject137) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || isNil(o.DefaultPolicy) {
    return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *InlineObject137) HasDefaultPolicy() bool {
	if o != nil && !isNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *InlineObject137) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *InlineObject137) GetAllowedServers() []string {
	if o == nil || isNil(o.AllowedServers) {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject137) GetAllowedServersOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedServers) {
    return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *InlineObject137) HasAllowedServers() bool {
	if o != nil && !isNil(o.AllowedServers) {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *InlineObject137) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *InlineObject137) GetBlockedServers() []string {
	if o == nil || isNil(o.BlockedServers) {
		var ret []string
		return ret
	}
	return o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject137) GetBlockedServersOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedServers) {
    return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *InlineObject137) HasBlockedServers() bool {
	if o != nil && !isNil(o.BlockedServers) {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *InlineObject137) SetBlockedServers(v []string) {
	o.BlockedServers = v
}

// GetArpInspection returns the ArpInspection field value if set, zero value otherwise.
func (o *InlineObject137) GetArpInspection() NetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	if o == nil || isNil(o.ArpInspection) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspection
		return ret
	}
	return *o.ArpInspection
}

// GetArpInspectionOk returns a tuple with the ArpInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject137) GetArpInspectionOk() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspection, bool) {
	if o == nil || isNil(o.ArpInspection) {
    return nil, false
	}
	return o.ArpInspection, true
}

// HasArpInspection returns a boolean if a field has been set.
func (o *InlineObject137) HasArpInspection() bool {
	if o != nil && !isNil(o.ArpInspection) {
		return true
	}

	return false
}

// SetArpInspection gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyArpInspection and assigns it to the ArpInspection field.
func (o *InlineObject137) SetArpInspection(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) {
	o.ArpInspection = &v
}

func (o InlineObject137) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject137 struct {
	value *InlineObject137
	isSet bool
}

func (v NullableInlineObject137) Get() *InlineObject137 {
	return v.value
}

func (v *NullableInlineObject137) Set(val *InlineObject137) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject137) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject137(val *InlineObject137) *NullableInlineObject137 {
	return &NullableInlineObject137{value: val, isSet: true}
}

func (v NullableInlineObject137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


