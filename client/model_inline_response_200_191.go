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

// InlineResponse200191 struct for InlineResponse200191
type InlineResponse200191 struct {
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
	NamedVlans *InlineResponse200191NamedVlans `json:"namedVlans,omitempty"`
	RegulatoryDomain *InlineResponse200191RegulatoryDomain `json:"regulatoryDomain,omitempty"`
}

// NewInlineResponse200191 instantiates a new InlineResponse200191 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200191() *InlineResponse200191 {
	this := InlineResponse200191{}
	return &this
}

// NewInlineResponse200191WithDefaults instantiates a new InlineResponse200191 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200191WithDefaults() *InlineResponse200191 {
	this := InlineResponse200191{}
	return &this
}

// GetMeshingEnabled returns the MeshingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200191) GetMeshingEnabled() bool {
	if o == nil || isNil(o.MeshingEnabled) {
		var ret bool
		return ret
	}
	return *o.MeshingEnabled
}

// GetMeshingEnabledOk returns a tuple with the MeshingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetMeshingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MeshingEnabled) {
    return nil, false
	}
	return o.MeshingEnabled, true
}

// HasMeshingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200191) HasMeshingEnabled() bool {
	if o != nil && !isNil(o.MeshingEnabled) {
		return true
	}

	return false
}

// SetMeshingEnabled gets a reference to the given bool and assigns it to the MeshingEnabled field.
func (o *InlineResponse200191) SetMeshingEnabled(v bool) {
	o.MeshingEnabled = &v
}

// GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field value if set, zero value otherwise.
func (o *InlineResponse200191) GetIpv6BridgeEnabled() bool {
	if o == nil || isNil(o.Ipv6BridgeEnabled) {
		var ret bool
		return ret
	}
	return *o.Ipv6BridgeEnabled
}

// GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetIpv6BridgeEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Ipv6BridgeEnabled) {
    return nil, false
	}
	return o.Ipv6BridgeEnabled, true
}

// HasIpv6BridgeEnabled returns a boolean if a field has been set.
func (o *InlineResponse200191) HasIpv6BridgeEnabled() bool {
	if o != nil && !isNil(o.Ipv6BridgeEnabled) {
		return true
	}

	return false
}

// SetIpv6BridgeEnabled gets a reference to the given bool and assigns it to the Ipv6BridgeEnabled field.
func (o *InlineResponse200191) SetIpv6BridgeEnabled(v bool) {
	o.Ipv6BridgeEnabled = &v
}

// GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field value if set, zero value otherwise.
func (o *InlineResponse200191) GetLocationAnalyticsEnabled() bool {
	if o == nil || isNil(o.LocationAnalyticsEnabled) {
		var ret bool
		return ret
	}
	return *o.LocationAnalyticsEnabled
}

// GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetLocationAnalyticsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocationAnalyticsEnabled) {
    return nil, false
	}
	return o.LocationAnalyticsEnabled, true
}

// HasLocationAnalyticsEnabled returns a boolean if a field has been set.
func (o *InlineResponse200191) HasLocationAnalyticsEnabled() bool {
	if o != nil && !isNil(o.LocationAnalyticsEnabled) {
		return true
	}

	return false
}

// SetLocationAnalyticsEnabled gets a reference to the given bool and assigns it to the LocationAnalyticsEnabled field.
func (o *InlineResponse200191) SetLocationAnalyticsEnabled(v bool) {
	o.LocationAnalyticsEnabled = &v
}

// GetUpgradeStrategy returns the UpgradeStrategy field value if set, zero value otherwise.
func (o *InlineResponse200191) GetUpgradeStrategy() string {
	if o == nil || isNil(o.UpgradeStrategy) {
		var ret string
		return ret
	}
	return *o.UpgradeStrategy
}

// GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetUpgradeStrategyOk() (*string, bool) {
	if o == nil || isNil(o.UpgradeStrategy) {
    return nil, false
	}
	return o.UpgradeStrategy, true
}

// HasUpgradeStrategy returns a boolean if a field has been set.
func (o *InlineResponse200191) HasUpgradeStrategy() bool {
	if o != nil && !isNil(o.UpgradeStrategy) {
		return true
	}

	return false
}

// SetUpgradeStrategy gets a reference to the given string and assigns it to the UpgradeStrategy field.
func (o *InlineResponse200191) SetUpgradeStrategy(v string) {
	o.UpgradeStrategy = &v
}

// GetLedLightsOn returns the LedLightsOn field value if set, zero value otherwise.
func (o *InlineResponse200191) GetLedLightsOn() bool {
	if o == nil || isNil(o.LedLightsOn) {
		var ret bool
		return ret
	}
	return *o.LedLightsOn
}

// GetLedLightsOnOk returns a tuple with the LedLightsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetLedLightsOnOk() (*bool, bool) {
	if o == nil || isNil(o.LedLightsOn) {
    return nil, false
	}
	return o.LedLightsOn, true
}

// HasLedLightsOn returns a boolean if a field has been set.
func (o *InlineResponse200191) HasLedLightsOn() bool {
	if o != nil && !isNil(o.LedLightsOn) {
		return true
	}

	return false
}

// SetLedLightsOn gets a reference to the given bool and assigns it to the LedLightsOn field.
func (o *InlineResponse200191) SetLedLightsOn(v bool) {
	o.LedLightsOn = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *InlineResponse200191) GetNamedVlans() InlineResponse200191NamedVlans {
	if o == nil || isNil(o.NamedVlans) {
		var ret InlineResponse200191NamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetNamedVlansOk() (*InlineResponse200191NamedVlans, bool) {
	if o == nil || isNil(o.NamedVlans) {
    return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *InlineResponse200191) HasNamedVlans() bool {
	if o != nil && !isNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given InlineResponse200191NamedVlans and assigns it to the NamedVlans field.
func (o *InlineResponse200191) SetNamedVlans(v InlineResponse200191NamedVlans) {
	o.NamedVlans = &v
}

// GetRegulatoryDomain returns the RegulatoryDomain field value if set, zero value otherwise.
func (o *InlineResponse200191) GetRegulatoryDomain() InlineResponse200191RegulatoryDomain {
	if o == nil || isNil(o.RegulatoryDomain) {
		var ret InlineResponse200191RegulatoryDomain
		return ret
	}
	return *o.RegulatoryDomain
}

// GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191) GetRegulatoryDomainOk() (*InlineResponse200191RegulatoryDomain, bool) {
	if o == nil || isNil(o.RegulatoryDomain) {
    return nil, false
	}
	return o.RegulatoryDomain, true
}

// HasRegulatoryDomain returns a boolean if a field has been set.
func (o *InlineResponse200191) HasRegulatoryDomain() bool {
	if o != nil && !isNil(o.RegulatoryDomain) {
		return true
	}

	return false
}

// SetRegulatoryDomain gets a reference to the given InlineResponse200191RegulatoryDomain and assigns it to the RegulatoryDomain field.
func (o *InlineResponse200191) SetRegulatoryDomain(v InlineResponse200191RegulatoryDomain) {
	o.RegulatoryDomain = &v
}

func (o InlineResponse200191) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RegulatoryDomain) {
		toSerialize["regulatoryDomain"] = o.RegulatoryDomain
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200191 struct {
	value *InlineResponse200191
	isSet bool
}

func (v NullableInlineResponse200191) Get() *InlineResponse200191 {
	return v.value
}

func (v *NullableInlineResponse200191) Set(val *InlineResponse200191) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200191) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200191) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200191(val *InlineResponse200191) *NullableInlineResponse200191 {
	return &NullableInlineResponse200191{value: val, isSet: true}
}

func (v NullableInlineResponse200191) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200191) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


