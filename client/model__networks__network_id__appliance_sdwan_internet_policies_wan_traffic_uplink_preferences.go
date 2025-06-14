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

// NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences struct for NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences
type NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences struct {
	// Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	PreferredUplink string `json:"preferredUplink"`
	// WAN failover and failback behavior
	FailOverCriterion *string `json:"failOverCriterion,omitempty"`
	PerformanceClass *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass `json:"performanceClass,omitempty"`
	// Traffic filters
	TrafficFilters []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters `json:"trafficFilters"`
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences(preferredUplink string, trafficFilters []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences{}
	this.PreferredUplink = preferredUplink
	this.TrafficFilters = trafficFilters
	return &this
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences{}
	return &this
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

// GetFailOverCriterion returns the FailOverCriterion field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetFailOverCriterion() string {
	if o == nil || isNil(o.FailOverCriterion) {
		var ret string
		return ret
	}
	return *o.FailOverCriterion
}

// GetFailOverCriterionOk returns a tuple with the FailOverCriterion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool) {
	if o == nil || isNil(o.FailOverCriterion) {
    return nil, false
	}
	return o.FailOverCriterion, true
}

// HasFailOverCriterion returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) HasFailOverCriterion() bool {
	if o != nil && !isNil(o.FailOverCriterion) {
		return true
	}

	return false
}

// SetFailOverCriterion gets a reference to the given string and assigns it to the FailOverCriterion field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) SetFailOverCriterion(v string) {
	o.FailOverCriterion = &v
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetPerformanceClass() NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass {
	if o == nil || isNil(o.PerformanceClass) {
		var ret NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass
		return ret
	}
	return *o.PerformanceClass
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetPerformanceClassOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass, bool) {
	if o == nil || isNil(o.PerformanceClass) {
    return nil, false
	}
	return o.PerformanceClass, true
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) HasPerformanceClass() bool {
	if o != nil && !isNil(o.PerformanceClass) {
		return true
	}

	return false
}

// SetPerformanceClass gets a reference to the given NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass and assigns it to the PerformanceClass field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) SetPerformanceClass(v NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) {
	o.PerformanceClass = &v
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetTrafficFilters() []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters {
	if o == nil {
		var ret []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) GetTrafficFiltersOk() ([]NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) SetTrafficFilters(v []NetworksNetworkIdApplianceSdwanInternetPoliciesTrafficFilters) {
	o.TrafficFilters = v
}

func (o NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	if !isNil(o.FailOverCriterion) {
		toSerialize["failOverCriterion"] = o.FailOverCriterion
	}
	if !isNil(o.PerformanceClass) {
		toSerialize["performanceClass"] = o.PerformanceClass
	}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences struct {
	value *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) Get() *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) Set(val *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences(val *NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	return &NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


