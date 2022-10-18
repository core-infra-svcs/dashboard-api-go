/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject113 struct for InlineObject113
type InlineObject113 struct {
	Alerts *NetworksNetworkIdSwitchDhcpServerPolicyAlerts `json:"alerts,omitempty"`
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	AllowedServers *[]string `json:"allowedServers,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	BlockedServers *[]string `json:"blockedServers,omitempty"`
	ArpInspection *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection `json:"arpInspection,omitempty"`
}

// NewInlineObject113 instantiates a new InlineObject113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject113() *InlineObject113 {
	this := InlineObject113{}
	return &this
}

// NewInlineObject113WithDefaults instantiates a new InlineObject113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject113WithDefaults() *InlineObject113 {
	this := InlineObject113{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineObject113) GetAlerts() NetworksNetworkIdSwitchDhcpServerPolicyAlerts {
	if o == nil || o.Alerts == nil {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetAlertsOk() (*NetworksNetworkIdSwitchDhcpServerPolicyAlerts, bool) {
	if o == nil || o.Alerts == nil {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineObject113) HasAlerts() bool {
	if o != nil && o.Alerts != nil {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyAlerts and assigns it to the Alerts field.
func (o *InlineObject113) SetAlerts(v NetworksNetworkIdSwitchDhcpServerPolicyAlerts) {
	o.Alerts = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *InlineObject113) GetDefaultPolicy() string {
	if o == nil || o.DefaultPolicy == nil {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || o.DefaultPolicy == nil {
		return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *InlineObject113) HasDefaultPolicy() bool {
	if o != nil && o.DefaultPolicy != nil {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *InlineObject113) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *InlineObject113) GetAllowedServers() []string {
	if o == nil || o.AllowedServers == nil {
		var ret []string
		return ret
	}
	return *o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetAllowedServersOk() (*[]string, bool) {
	if o == nil || o.AllowedServers == nil {
		return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *InlineObject113) HasAllowedServers() bool {
	if o != nil && o.AllowedServers != nil {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *InlineObject113) SetAllowedServers(v []string) {
	o.AllowedServers = &v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *InlineObject113) GetBlockedServers() []string {
	if o == nil || o.BlockedServers == nil {
		var ret []string
		return ret
	}
	return *o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetBlockedServersOk() (*[]string, bool) {
	if o == nil || o.BlockedServers == nil {
		return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *InlineObject113) HasBlockedServers() bool {
	if o != nil && o.BlockedServers != nil {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *InlineObject113) SetBlockedServers(v []string) {
	o.BlockedServers = &v
}

// GetArpInspection returns the ArpInspection field value if set, zero value otherwise.
func (o *InlineObject113) GetArpInspection() NetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	if o == nil || o.ArpInspection == nil {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspection
		return ret
	}
	return *o.ArpInspection
}

// GetArpInspectionOk returns a tuple with the ArpInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetArpInspectionOk() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspection, bool) {
	if o == nil || o.ArpInspection == nil {
		return nil, false
	}
	return o.ArpInspection, true
}

// HasArpInspection returns a boolean if a field has been set.
func (o *InlineObject113) HasArpInspection() bool {
	if o != nil && o.ArpInspection != nil {
		return true
	}

	return false
}

// SetArpInspection gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyArpInspection and assigns it to the ArpInspection field.
func (o *InlineObject113) SetArpInspection(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) {
	o.ArpInspection = &v
}

func (o InlineObject113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alerts != nil {
		toSerialize["alerts"] = o.Alerts
	}
	if o.DefaultPolicy != nil {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	if o.AllowedServers != nil {
		toSerialize["allowedServers"] = o.AllowedServers
	}
	if o.BlockedServers != nil {
		toSerialize["blockedServers"] = o.BlockedServers
	}
	if o.ArpInspection != nil {
		toSerialize["arpInspection"] = o.ArpInspection
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject113 struct {
	value *InlineObject113
	isSet bool
}

func (v NullableInlineObject113) Get() *InlineObject113 {
	return v.value
}

func (v *NullableInlineObject113) Set(val *InlineObject113) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject113) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject113(val *InlineObject113) *NullableInlineObject113 {
	return &NullableInlineObject113{value: val, isSet: true}
}

func (v NullableInlineObject113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


