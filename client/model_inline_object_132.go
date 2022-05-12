/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject132 struct for InlineObject132
type InlineObject132 struct {
	// Toggle for enabling or disabling meshing in a network
	MeshingEnabled *bool `json:"meshingEnabled,omitempty"`
	// Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	Ipv6BridgeEnabled *bool `json:"ipv6BridgeEnabled,omitempty"`
	// Toggle for enabling or disabling location analytics for your network
	LocationAnalyticsEnabled *bool `json:"locationAnalyticsEnabled,omitempty"`
	// The upgrade strategy to apply to the network. Must be one of 'minimizeUpgradeTime' or 'minimizeClientDowntime'. Requires firmware version MR 26.8 or higher'
	UpgradeStrategy *string `json:"upgradeStrategy,omitempty"`
	// Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LedLightsOn *bool `json:"ledLightsOn,omitempty"`
}

// NewInlineObject132 instantiates a new InlineObject132 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject132() *InlineObject132 {
	this := InlineObject132{}
	return &this
}

// NewInlineObject132WithDefaults instantiates a new InlineObject132 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject132WithDefaults() *InlineObject132 {
	this := InlineObject132{}
	return &this
}

// GetMeshingEnabled returns the MeshingEnabled field value if set, zero value otherwise.
func (o *InlineObject132) GetMeshingEnabled() bool {
	if o == nil || o.MeshingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MeshingEnabled
}

// GetMeshingEnabledOk returns a tuple with the MeshingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetMeshingEnabledOk() (*bool, bool) {
	if o == nil || o.MeshingEnabled == nil {
		return nil, false
	}
	return o.MeshingEnabled, true
}

// HasMeshingEnabled returns a boolean if a field has been set.
func (o *InlineObject132) HasMeshingEnabled() bool {
	if o != nil && o.MeshingEnabled != nil {
		return true
	}

	return false
}

// SetMeshingEnabled gets a reference to the given bool and assigns it to the MeshingEnabled field.
func (o *InlineObject132) SetMeshingEnabled(v bool) {
	o.MeshingEnabled = &v
}

// GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field value if set, zero value otherwise.
func (o *InlineObject132) GetIpv6BridgeEnabled() bool {
	if o == nil || o.Ipv6BridgeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6BridgeEnabled
}

// GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetIpv6BridgeEnabledOk() (*bool, bool) {
	if o == nil || o.Ipv6BridgeEnabled == nil {
		return nil, false
	}
	return o.Ipv6BridgeEnabled, true
}

// HasIpv6BridgeEnabled returns a boolean if a field has been set.
func (o *InlineObject132) HasIpv6BridgeEnabled() bool {
	if o != nil && o.Ipv6BridgeEnabled != nil {
		return true
	}

	return false
}

// SetIpv6BridgeEnabled gets a reference to the given bool and assigns it to the Ipv6BridgeEnabled field.
func (o *InlineObject132) SetIpv6BridgeEnabled(v bool) {
	o.Ipv6BridgeEnabled = &v
}

// GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field value if set, zero value otherwise.
func (o *InlineObject132) GetLocationAnalyticsEnabled() bool {
	if o == nil || o.LocationAnalyticsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LocationAnalyticsEnabled
}

// GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetLocationAnalyticsEnabledOk() (*bool, bool) {
	if o == nil || o.LocationAnalyticsEnabled == nil {
		return nil, false
	}
	return o.LocationAnalyticsEnabled, true
}

// HasLocationAnalyticsEnabled returns a boolean if a field has been set.
func (o *InlineObject132) HasLocationAnalyticsEnabled() bool {
	if o != nil && o.LocationAnalyticsEnabled != nil {
		return true
	}

	return false
}

// SetLocationAnalyticsEnabled gets a reference to the given bool and assigns it to the LocationAnalyticsEnabled field.
func (o *InlineObject132) SetLocationAnalyticsEnabled(v bool) {
	o.LocationAnalyticsEnabled = &v
}

// GetUpgradeStrategy returns the UpgradeStrategy field value if set, zero value otherwise.
func (o *InlineObject132) GetUpgradeStrategy() string {
	if o == nil || o.UpgradeStrategy == nil {
		var ret string
		return ret
	}
	return *o.UpgradeStrategy
}

// GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetUpgradeStrategyOk() (*string, bool) {
	if o == nil || o.UpgradeStrategy == nil {
		return nil, false
	}
	return o.UpgradeStrategy, true
}

// HasUpgradeStrategy returns a boolean if a field has been set.
func (o *InlineObject132) HasUpgradeStrategy() bool {
	if o != nil && o.UpgradeStrategy != nil {
		return true
	}

	return false
}

// SetUpgradeStrategy gets a reference to the given string and assigns it to the UpgradeStrategy field.
func (o *InlineObject132) SetUpgradeStrategy(v string) {
	o.UpgradeStrategy = &v
}

// GetLedLightsOn returns the LedLightsOn field value if set, zero value otherwise.
func (o *InlineObject132) GetLedLightsOn() bool {
	if o == nil || o.LedLightsOn == nil {
		var ret bool
		return ret
	}
	return *o.LedLightsOn
}

// GetLedLightsOnOk returns a tuple with the LedLightsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject132) GetLedLightsOnOk() (*bool, bool) {
	if o == nil || o.LedLightsOn == nil {
		return nil, false
	}
	return o.LedLightsOn, true
}

// HasLedLightsOn returns a boolean if a field has been set.
func (o *InlineObject132) HasLedLightsOn() bool {
	if o != nil && o.LedLightsOn != nil {
		return true
	}

	return false
}

// SetLedLightsOn gets a reference to the given bool and assigns it to the LedLightsOn field.
func (o *InlineObject132) SetLedLightsOn(v bool) {
	o.LedLightsOn = &v
}

func (o InlineObject132) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MeshingEnabled != nil {
		toSerialize["meshingEnabled"] = o.MeshingEnabled
	}
	if o.Ipv6BridgeEnabled != nil {
		toSerialize["ipv6BridgeEnabled"] = o.Ipv6BridgeEnabled
	}
	if o.LocationAnalyticsEnabled != nil {
		toSerialize["locationAnalyticsEnabled"] = o.LocationAnalyticsEnabled
	}
	if o.UpgradeStrategy != nil {
		toSerialize["upgradeStrategy"] = o.UpgradeStrategy
	}
	if o.LedLightsOn != nil {
		toSerialize["ledLightsOn"] = o.LedLightsOn
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject132 struct {
	value *InlineObject132
	isSet bool
}

func (v NullableInlineObject132) Get() *InlineObject132 {
	return v.value
}

func (v *NullableInlineObject132) Set(val *InlineObject132) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject132) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject132) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject132(val *InlineObject132) *NullableInlineObject132 {
	return &NullableInlineObject132{value: val, isSet: true}
}

func (v NullableInlineObject132) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject132) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

