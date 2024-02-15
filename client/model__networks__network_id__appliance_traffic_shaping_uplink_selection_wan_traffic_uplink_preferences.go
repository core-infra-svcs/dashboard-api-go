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

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct for NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	// Array of traffic filters for this uplink preference rule
	TrafficFilters []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters `json:"trafficFilters"`
	// Preferred uplink for this uplink preference rule. Must be one of: 'wan1' or 'wan2'
	PreferredUplink string `json:"preferredUplink"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences(trafficFilters []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters, preferredUplink string) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetTrafficFilters() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters {
	if o == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetTrafficFiltersOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) SetTrafficFilters(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


