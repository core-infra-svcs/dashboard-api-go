/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject183 struct for InlineObject183
type InlineObject183 struct {
	// Toggle for enabling or disabling meshing in a network
	MeshingEnabled *bool `json:"meshingEnabled,omitempty"`
	// Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	Ipv6BridgeEnabled *bool `json:"ipv6BridgeEnabled,omitempty"`
	// Toggle for enabling or disabling location analytics for your network
	LocationAnalyticsEnabled *bool `json:"locationAnalyticsEnabled,omitempty"`
	// The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher.
	UpgradeStrategy *string `json:"upgradeStrategy,omitempty"`
	// Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LedLightsOn *bool `json:"ledLightsOn,omitempty"`
	NamedVlans *NetworksNetworkIdWirelessSettingsNamedVlans `json:"namedVlans,omitempty"`
}

// NewInlineObject183 instantiates a new InlineObject183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject183() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// NewInlineObject183WithDefaults instantiates a new InlineObject183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject183WithDefaults() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// GetMeshingEnabled returns the MeshingEnabled field value if set, zero value otherwise.
func (o *InlineObject183) GetMeshingEnabled() bool {
	if o == nil || isNil(o.MeshingEnabled) {
		var ret bool
		return ret
	}
	return *o.MeshingEnabled
}

// GetMeshingEnabledOk returns a tuple with the MeshingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetMeshingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MeshingEnabled) {
    return nil, false
	}
	return o.MeshingEnabled, true
}

// HasMeshingEnabled returns a boolean if a field has been set.
func (o *InlineObject183) HasMeshingEnabled() bool {
	if o != nil && !isNil(o.MeshingEnabled) {
		return true
	}

	return false
}

// SetMeshingEnabled gets a reference to the given bool and assigns it to the MeshingEnabled field.
func (o *InlineObject183) SetMeshingEnabled(v bool) {
	o.MeshingEnabled = &v
}

// GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field value if set, zero value otherwise.
func (o *InlineObject183) GetIpv6BridgeEnabled() bool {
	if o == nil || isNil(o.Ipv6BridgeEnabled) {
		var ret bool
		return ret
	}
	return *o.Ipv6BridgeEnabled
}

// GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetIpv6BridgeEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Ipv6BridgeEnabled) {
    return nil, false
	}
	return o.Ipv6BridgeEnabled, true
}

// HasIpv6BridgeEnabled returns a boolean if a field has been set.
func (o *InlineObject183) HasIpv6BridgeEnabled() bool {
	if o != nil && !isNil(o.Ipv6BridgeEnabled) {
		return true
	}

	return false
}

// SetIpv6BridgeEnabled gets a reference to the given bool and assigns it to the Ipv6BridgeEnabled field.
func (o *InlineObject183) SetIpv6BridgeEnabled(v bool) {
	o.Ipv6BridgeEnabled = &v
}

// GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field value if set, zero value otherwise.
func (o *InlineObject183) GetLocationAnalyticsEnabled() bool {
	if o == nil || isNil(o.LocationAnalyticsEnabled) {
		var ret bool
		return ret
	}
	return *o.LocationAnalyticsEnabled
}

// GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetLocationAnalyticsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocationAnalyticsEnabled) {
    return nil, false
	}
	return o.LocationAnalyticsEnabled, true
}

// HasLocationAnalyticsEnabled returns a boolean if a field has been set.
func (o *InlineObject183) HasLocationAnalyticsEnabled() bool {
	if o != nil && !isNil(o.LocationAnalyticsEnabled) {
		return true
	}

	return false
}

// SetLocationAnalyticsEnabled gets a reference to the given bool and assigns it to the LocationAnalyticsEnabled field.
func (o *InlineObject183) SetLocationAnalyticsEnabled(v bool) {
	o.LocationAnalyticsEnabled = &v
}

// GetUpgradeStrategy returns the UpgradeStrategy field value if set, zero value otherwise.
func (o *InlineObject183) GetUpgradeStrategy() string {
	if o == nil || isNil(o.UpgradeStrategy) {
		var ret string
		return ret
	}
	return *o.UpgradeStrategy
}

// GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetUpgradeStrategyOk() (*string, bool) {
	if o == nil || isNil(o.UpgradeStrategy) {
    return nil, false
	}
	return o.UpgradeStrategy, true
}

// HasUpgradeStrategy returns a boolean if a field has been set.
func (o *InlineObject183) HasUpgradeStrategy() bool {
	if o != nil && !isNil(o.UpgradeStrategy) {
		return true
	}

	return false
}

// SetUpgradeStrategy gets a reference to the given string and assigns it to the UpgradeStrategy field.
func (o *InlineObject183) SetUpgradeStrategy(v string) {
	o.UpgradeStrategy = &v
}

// GetLedLightsOn returns the LedLightsOn field value if set, zero value otherwise.
func (o *InlineObject183) GetLedLightsOn() bool {
	if o == nil || isNil(o.LedLightsOn) {
		var ret bool
		return ret
	}
	return *o.LedLightsOn
}

// GetLedLightsOnOk returns a tuple with the LedLightsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetLedLightsOnOk() (*bool, bool) {
	if o == nil || isNil(o.LedLightsOn) {
    return nil, false
	}
	return o.LedLightsOn, true
}

// HasLedLightsOn returns a boolean if a field has been set.
func (o *InlineObject183) HasLedLightsOn() bool {
	if o != nil && !isNil(o.LedLightsOn) {
		return true
	}

	return false
}

// SetLedLightsOn gets a reference to the given bool and assigns it to the LedLightsOn field.
func (o *InlineObject183) SetLedLightsOn(v bool) {
	o.LedLightsOn = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *InlineObject183) GetNamedVlans() NetworksNetworkIdWirelessSettingsNamedVlans {
	if o == nil || isNil(o.NamedVlans) {
		var ret NetworksNetworkIdWirelessSettingsNamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetNamedVlansOk() (*NetworksNetworkIdWirelessSettingsNamedVlans, bool) {
	if o == nil || isNil(o.NamedVlans) {
    return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *InlineObject183) HasNamedVlans() bool {
	if o != nil && !isNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given NetworksNetworkIdWirelessSettingsNamedVlans and assigns it to the NamedVlans field.
func (o *InlineObject183) SetNamedVlans(v NetworksNetworkIdWirelessSettingsNamedVlans) {
	o.NamedVlans = &v
}

func (o InlineObject183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MeshingEnabled) {
		toSerialize["meshingEnabled"] = o.MeshingEnabled
	}
	if !isNil(o.Ipv6BridgeEnabled) {
		toSerialize["ipv6BridgeEnabled"] = o.Ipv6BridgeEnabled
	}
	if !isNil(o.LocationAnalyticsEnabled) {
		toSerialize["locationAnalyticsEnabled"] = o.LocationAnalyticsEnabled
	}
	if !isNil(o.UpgradeStrategy) {
		toSerialize["upgradeStrategy"] = o.UpgradeStrategy
	}
	if !isNil(o.LedLightsOn) {
		toSerialize["ledLightsOn"] = o.LedLightsOn
	}
	if !isNil(o.NamedVlans) {
		toSerialize["namedVlans"] = o.NamedVlans
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject183 struct {
	value *InlineObject183
	isSet bool
}

func (v NullableInlineObject183) Get() *InlineObject183 {
	return v.value
}

func (v *NullableInlineObject183) Set(val *InlineObject183) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject183) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject183(val *InlineObject183) *NullableInlineObject183 {
	return &NullableInlineObject183{value: val, isSet: true}
}

func (v NullableInlineObject183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


