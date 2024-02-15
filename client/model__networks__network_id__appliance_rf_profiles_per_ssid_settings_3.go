/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 Settings for SSID 3
type NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 struct {
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewNetworksNetworkIdApplianceRfProfilesPerSsidSettings3 instantiates a new NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceRfProfilesPerSsidSettings3() *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 {
	this := NetworksNetworkIdApplianceRfProfilesPerSsidSettings3{}
	return &this
}

// NewNetworksNetworkIdApplianceRfProfilesPerSsidSettings3WithDefaults instantiates a new NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceRfProfilesPerSsidSettings3WithDefaults() *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 {
	this := NetworksNetworkIdApplianceRfProfilesPerSsidSettings3{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3 struct {
	value *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3) Get() *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3) Set(val *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3(val *NetworksNetworkIdApplianceRfProfilesPerSsidSettings3) *NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3 {
	return &NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesPerSsidSettings3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


