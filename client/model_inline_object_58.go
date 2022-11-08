/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject58 struct for InlineObject58
type InlineObject58 struct {
	// Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured.
	Enabled bool `json:"enabled"`
	// An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512.
	AsNumber *int32 `json:"asNumber,omitempty"`
	// The IBGP holdtimer in seconds. The IBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.
	IbgpHoldTimer *int32 `json:"ibgpHoldTimer,omitempty"`
	// List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.
	Neighbors []NetworksNetworkIdApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`
}

// NewInlineObject58 instantiates a new InlineObject58 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject58(enabled bool) *InlineObject58 {
	this := InlineObject58{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject58WithDefaults instantiates a new InlineObject58 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject58WithDefaults() *InlineObject58 {
	this := InlineObject58{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject58) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject58) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise.
func (o *InlineObject58) GetAsNumber() int32 {
	if o == nil || isNil(o.AsNumber) {
		var ret int32
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetAsNumberOk() (*int32, bool) {
	if o == nil || isNil(o.AsNumber) {
    return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *InlineObject58) HasAsNumber() bool {
	if o != nil && !isNil(o.AsNumber) {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given int32 and assigns it to the AsNumber field.
func (o *InlineObject58) SetAsNumber(v int32) {
	o.AsNumber = &v
}

// GetIbgpHoldTimer returns the IbgpHoldTimer field value if set, zero value otherwise.
func (o *InlineObject58) GetIbgpHoldTimer() int32 {
	if o == nil || isNil(o.IbgpHoldTimer) {
		var ret int32
		return ret
	}
	return *o.IbgpHoldTimer
}

// GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetIbgpHoldTimerOk() (*int32, bool) {
	if o == nil || isNil(o.IbgpHoldTimer) {
    return nil, false
	}
	return o.IbgpHoldTimer, true
}

// HasIbgpHoldTimer returns a boolean if a field has been set.
func (o *InlineObject58) HasIbgpHoldTimer() bool {
	if o != nil && !isNil(o.IbgpHoldTimer) {
		return true
	}

	return false
}

// SetIbgpHoldTimer gets a reference to the given int32 and assigns it to the IbgpHoldTimer field.
func (o *InlineObject58) SetIbgpHoldTimer(v int32) {
	o.IbgpHoldTimer = &v
}

// GetNeighbors returns the Neighbors field value if set, zero value otherwise.
func (o *InlineObject58) GetNeighbors() []NetworksNetworkIdApplianceVpnBgpNeighbors {
	if o == nil || isNil(o.Neighbors) {
		var ret []NetworksNetworkIdApplianceVpnBgpNeighbors
		return ret
	}
	return o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject58) GetNeighborsOk() ([]NetworksNetworkIdApplianceVpnBgpNeighbors, bool) {
	if o == nil || isNil(o.Neighbors) {
    return nil, false
	}
	return o.Neighbors, true
}

// HasNeighbors returns a boolean if a field has been set.
func (o *InlineObject58) HasNeighbors() bool {
	if o != nil && !isNil(o.Neighbors) {
		return true
	}

	return false
}

// SetNeighbors gets a reference to the given []NetworksNetworkIdApplianceVpnBgpNeighbors and assigns it to the Neighbors field.
func (o *InlineObject58) SetNeighbors(v []NetworksNetworkIdApplianceVpnBgpNeighbors) {
	o.Neighbors = v
}

func (o InlineObject58) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AsNumber) {
		toSerialize["asNumber"] = o.AsNumber
	}
	if !isNil(o.IbgpHoldTimer) {
		toSerialize["ibgpHoldTimer"] = o.IbgpHoldTimer
	}
	if !isNil(o.Neighbors) {
		toSerialize["neighbors"] = o.Neighbors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject58 struct {
	value *InlineObject58
	isSet bool
}

func (v NullableInlineObject58) Get() *InlineObject58 {
	return v.value
}

func (v *NullableInlineObject58) Set(val *InlineObject58) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject58) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject58) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject58(val *InlineObject58) *NullableInlineObject58 {
	return &NullableInlineObject58{value: val, isSet: true}
}

func (v NullableInlineObject58) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject58) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


