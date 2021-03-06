/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject50 struct for InlineObject50
type InlineObject50 struct {
	// Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured.
	Enabled bool `json:"enabled"`
	// An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512.
	AsNumber *int32 `json:"asNumber,omitempty"`
	// The IBGP holdtimer in seconds. The IBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.
	IbgpHoldTimer *int32 `json:"ibgpHoldTimer,omitempty"`
	// List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.
	Neighbors *[]NetworksNetworkIdApplianceVpnBgpNeighbors `json:"neighbors,omitempty"`
}

// NewInlineObject50 instantiates a new InlineObject50 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject50(enabled bool) *InlineObject50 {
	this := InlineObject50{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject50WithDefaults instantiates a new InlineObject50 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject50WithDefaults() *InlineObject50 {
	this := InlineObject50{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject50) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject50) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject50) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise.
func (o *InlineObject50) GetAsNumber() int32 {
	if o == nil || o.AsNumber == nil {
		var ret int32
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject50) GetAsNumberOk() (*int32, bool) {
	if o == nil || o.AsNumber == nil {
		return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *InlineObject50) HasAsNumber() bool {
	if o != nil && o.AsNumber != nil {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given int32 and assigns it to the AsNumber field.
func (o *InlineObject50) SetAsNumber(v int32) {
	o.AsNumber = &v
}

// GetIbgpHoldTimer returns the IbgpHoldTimer field value if set, zero value otherwise.
func (o *InlineObject50) GetIbgpHoldTimer() int32 {
	if o == nil || o.IbgpHoldTimer == nil {
		var ret int32
		return ret
	}
	return *o.IbgpHoldTimer
}

// GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject50) GetIbgpHoldTimerOk() (*int32, bool) {
	if o == nil || o.IbgpHoldTimer == nil {
		return nil, false
	}
	return o.IbgpHoldTimer, true
}

// HasIbgpHoldTimer returns a boolean if a field has been set.
func (o *InlineObject50) HasIbgpHoldTimer() bool {
	if o != nil && o.IbgpHoldTimer != nil {
		return true
	}

	return false
}

// SetIbgpHoldTimer gets a reference to the given int32 and assigns it to the IbgpHoldTimer field.
func (o *InlineObject50) SetIbgpHoldTimer(v int32) {
	o.IbgpHoldTimer = &v
}

// GetNeighbors returns the Neighbors field value if set, zero value otherwise.
func (o *InlineObject50) GetNeighbors() []NetworksNetworkIdApplianceVpnBgpNeighbors {
	if o == nil || o.Neighbors == nil {
		var ret []NetworksNetworkIdApplianceVpnBgpNeighbors
		return ret
	}
	return *o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject50) GetNeighborsOk() (*[]NetworksNetworkIdApplianceVpnBgpNeighbors, bool) {
	if o == nil || o.Neighbors == nil {
		return nil, false
	}
	return o.Neighbors, true
}

// HasNeighbors returns a boolean if a field has been set.
func (o *InlineObject50) HasNeighbors() bool {
	if o != nil && o.Neighbors != nil {
		return true
	}

	return false
}

// SetNeighbors gets a reference to the given []NetworksNetworkIdApplianceVpnBgpNeighbors and assigns it to the Neighbors field.
func (o *InlineObject50) SetNeighbors(v []NetworksNetworkIdApplianceVpnBgpNeighbors) {
	o.Neighbors = &v
}

func (o InlineObject50) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AsNumber != nil {
		toSerialize["asNumber"] = o.AsNumber
	}
	if o.IbgpHoldTimer != nil {
		toSerialize["ibgpHoldTimer"] = o.IbgpHoldTimer
	}
	if o.Neighbors != nil {
		toSerialize["neighbors"] = o.Neighbors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject50 struct {
	value *InlineObject50
	isSet bool
}

func (v NullableInlineObject50) Get() *InlineObject50 {
	return v.value
}

func (v *NullableInlineObject50) Set(val *InlineObject50) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject50) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject50) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject50(val *InlineObject50) *NullableInlineObject50 {
	return &NullableInlineObject50{value: val, isSet: true}
}

func (v NullableInlineObject50) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject50) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


