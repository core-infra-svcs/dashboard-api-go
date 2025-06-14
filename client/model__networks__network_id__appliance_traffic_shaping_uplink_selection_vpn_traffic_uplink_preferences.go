/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct for NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	// Array of traffic filters for this uplink preference rule
	TrafficFilters []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 `json:"trafficFilters"`
	// Preferred uplink for this uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	PreferredUplink string `json:"preferredUplink"`
	// Fail over criterion for this uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'
	FailOverCriterion *string `json:"failOverCriterion,omitempty"`
	PerformanceClass *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass `json:"performanceClass,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences(trafficFilters []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1, preferredUplink string) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetTrafficFilters() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 {
	if o == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetTrafficFiltersOk() ([]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetTrafficFilters(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

// GetFailOverCriterion returns the FailOverCriterion field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetFailOverCriterion() string {
	if o == nil || isNil(o.FailOverCriterion) {
		var ret string
		return ret
	}
	return *o.FailOverCriterion
}

// GetFailOverCriterionOk returns a tuple with the FailOverCriterion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool) {
	if o == nil || isNil(o.FailOverCriterion) {
    return nil, false
	}
	return o.FailOverCriterion, true
}

// HasFailOverCriterion returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) HasFailOverCriterion() bool {
	if o != nil && !isNil(o.FailOverCriterion) {
		return true
	}

	return false
}

// SetFailOverCriterion gets a reference to the given string and assigns it to the FailOverCriterion field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetFailOverCriterion(v string) {
	o.FailOverCriterion = &v
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPerformanceClass() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass {
	if o == nil || isNil(o.PerformanceClass) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass
		return ret
	}
	return *o.PerformanceClass
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) GetPerformanceClassOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass, bool) {
	if o == nil || isNil(o.PerformanceClass) {
    return nil, false
	}
	return o.PerformanceClass, true
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) HasPerformanceClass() bool {
	if o != nil && !isNil(o.PerformanceClass) {
		return true
	}

	return false
}

// SetPerformanceClass gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass and assigns it to the PerformanceClass field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) SetPerformanceClass(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionPerformanceClass) {
	o.PerformanceClass = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	if !isNil(o.FailOverCriterion) {
		toSerialize["failOverCriterion"] = o.FailOverCriterion
	}
	if !isNil(o.PerformanceClass) {
		toSerialize["performanceClass"] = o.PerformanceClass
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


