/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject189 struct for InlineObject189
type InlineObject189 struct {
	Concentrator *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator `json:"concentrator,omitempty"`
	SplitTunnel *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel `json:"splitTunnel,omitempty"`
	Failover *NetworksNetworkIdWirelessSsidsNumberVpnFailover `json:"failover,omitempty"`
}

// NewInlineObject189 instantiates a new InlineObject189 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject189() *InlineObject189 {
	this := InlineObject189{}
	return &this
}

// NewInlineObject189WithDefaults instantiates a new InlineObject189 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject189WithDefaults() *InlineObject189 {
	this := InlineObject189{}
	return &this
}

// GetConcentrator returns the Concentrator field value if set, zero value otherwise.
func (o *InlineObject189) GetConcentrator() NetworksNetworkIdWirelessSsidsNumberVpnConcentrator {
	if o == nil || isNil(o.Concentrator) {
		var ret NetworksNetworkIdWirelessSsidsNumberVpnConcentrator
		return ret
	}
	return *o.Concentrator
}

// GetConcentratorOk returns a tuple with the Concentrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetConcentratorOk() (*NetworksNetworkIdWirelessSsidsNumberVpnConcentrator, bool) {
	if o == nil || isNil(o.Concentrator) {
    return nil, false
	}
	return o.Concentrator, true
}

// HasConcentrator returns a boolean if a field has been set.
func (o *InlineObject189) HasConcentrator() bool {
	if o != nil && !isNil(o.Concentrator) {
		return true
	}

	return false
}

// SetConcentrator gets a reference to the given NetworksNetworkIdWirelessSsidsNumberVpnConcentrator and assigns it to the Concentrator field.
func (o *InlineObject189) SetConcentrator(v NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) {
	o.Concentrator = &v
}

// GetSplitTunnel returns the SplitTunnel field value if set, zero value otherwise.
func (o *InlineObject189) GetSplitTunnel() NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel {
	if o == nil || isNil(o.SplitTunnel) {
		var ret NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel
		return ret
	}
	return *o.SplitTunnel
}

// GetSplitTunnelOk returns a tuple with the SplitTunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetSplitTunnelOk() (*NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel, bool) {
	if o == nil || isNil(o.SplitTunnel) {
    return nil, false
	}
	return o.SplitTunnel, true
}

// HasSplitTunnel returns a boolean if a field has been set.
func (o *InlineObject189) HasSplitTunnel() bool {
	if o != nil && !isNil(o.SplitTunnel) {
		return true
	}

	return false
}

// SetSplitTunnel gets a reference to the given NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel and assigns it to the SplitTunnel field.
func (o *InlineObject189) SetSplitTunnel(v NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) {
	o.SplitTunnel = &v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *InlineObject189) GetFailover() NetworksNetworkIdWirelessSsidsNumberVpnFailover {
	if o == nil || isNil(o.Failover) {
		var ret NetworksNetworkIdWirelessSsidsNumberVpnFailover
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetFailoverOk() (*NetworksNetworkIdWirelessSsidsNumberVpnFailover, bool) {
	if o == nil || isNil(o.Failover) {
    return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *InlineObject189) HasFailover() bool {
	if o != nil && !isNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given NetworksNetworkIdWirelessSsidsNumberVpnFailover and assigns it to the Failover field.
func (o *InlineObject189) SetFailover(v NetworksNetworkIdWirelessSsidsNumberVpnFailover) {
	o.Failover = &v
}

func (o InlineObject189) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject189 struct {
	value *InlineObject189
	isSet bool
}

func (v NullableInlineObject189) Get() *InlineObject189 {
	return v.value
}

func (v *NullableInlineObject189) Set(val *InlineObject189) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject189) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject189) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject189(val *InlineObject189) *NullableInlineObject189 {
	return &NullableInlineObject189{value: val, isSet: true}
}

func (v NullableInlineObject189) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject189) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


