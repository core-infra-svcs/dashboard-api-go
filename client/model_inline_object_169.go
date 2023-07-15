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

// InlineObject169 struct for InlineObject169
type InlineObject169 struct {
	Concentrator *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator `json:"concentrator,omitempty"`
	SplitTunnel *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel `json:"splitTunnel,omitempty"`
	Failover *NetworksNetworkIdWirelessSsidsNumberVpnFailover `json:"failover,omitempty"`
}

// NewInlineObject169 instantiates a new InlineObject169 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject169() *InlineObject169 {
	this := InlineObject169{}
	return &this
}

// NewInlineObject169WithDefaults instantiates a new InlineObject169 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject169WithDefaults() *InlineObject169 {
	this := InlineObject169{}
	return &this
}

// GetConcentrator returns the Concentrator field value if set, zero value otherwise.
func (o *InlineObject169) GetConcentrator() NetworksNetworkIdWirelessSsidsNumberVpnConcentrator {
	if o == nil || isNil(o.Concentrator) {
		var ret NetworksNetworkIdWirelessSsidsNumberVpnConcentrator
		return ret
	}
	return *o.Concentrator
}

// GetConcentratorOk returns a tuple with the Concentrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject169) GetConcentratorOk() (*NetworksNetworkIdWirelessSsidsNumberVpnConcentrator, bool) {
	if o == nil || isNil(o.Concentrator) {
    return nil, false
	}
	return o.Concentrator, true
}

// HasConcentrator returns a boolean if a field has been set.
func (o *InlineObject169) HasConcentrator() bool {
	if o != nil && !isNil(o.Concentrator) {
		return true
	}

	return false
}

// SetConcentrator gets a reference to the given NetworksNetworkIdWirelessSsidsNumberVpnConcentrator and assigns it to the Concentrator field.
func (o *InlineObject169) SetConcentrator(v NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) {
	o.Concentrator = &v
}

// GetSplitTunnel returns the SplitTunnel field value if set, zero value otherwise.
func (o *InlineObject169) GetSplitTunnel() NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel {
	if o == nil || isNil(o.SplitTunnel) {
		var ret NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel
		return ret
	}
	return *o.SplitTunnel
}

// GetSplitTunnelOk returns a tuple with the SplitTunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject169) GetSplitTunnelOk() (*NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel, bool) {
	if o == nil || isNil(o.SplitTunnel) {
    return nil, false
	}
	return o.SplitTunnel, true
}

// HasSplitTunnel returns a boolean if a field has been set.
func (o *InlineObject169) HasSplitTunnel() bool {
	if o != nil && !isNil(o.SplitTunnel) {
		return true
	}

	return false
}

// SetSplitTunnel gets a reference to the given NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel and assigns it to the SplitTunnel field.
func (o *InlineObject169) SetSplitTunnel(v NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) {
	o.SplitTunnel = &v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *InlineObject169) GetFailover() NetworksNetworkIdWirelessSsidsNumberVpnFailover {
	if o == nil || isNil(o.Failover) {
		var ret NetworksNetworkIdWirelessSsidsNumberVpnFailover
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject169) GetFailoverOk() (*NetworksNetworkIdWirelessSsidsNumberVpnFailover, bool) {
	if o == nil || isNil(o.Failover) {
    return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *InlineObject169) HasFailover() bool {
	if o != nil && !isNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given NetworksNetworkIdWirelessSsidsNumberVpnFailover and assigns it to the Failover field.
func (o *InlineObject169) SetFailover(v NetworksNetworkIdWirelessSsidsNumberVpnFailover) {
	o.Failover = &v
}

func (o InlineObject169) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentrator) {
		toSerialize["concentrator"] = o.Concentrator
	}
	if !isNil(o.SplitTunnel) {
		toSerialize["splitTunnel"] = o.SplitTunnel
	}
	if !isNil(o.Failover) {
		toSerialize["failover"] = o.Failover
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject169 struct {
	value *InlineObject169
	isSet bool
}

func (v NullableInlineObject169) Get() *InlineObject169 {
	return v.value
}

func (v *NullableInlineObject169) Set(val *InlineObject169) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject169) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject169) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject169(val *InlineObject169) *NullableInlineObject169 {
	return &NullableInlineObject169{value: val, isSet: true}
}

func (v NullableInlineObject169) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject169) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


