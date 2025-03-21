/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20069VpnTrafficUplinkPreferences struct for InlineResponse20069VpnTrafficUplinkPreferences
type InlineResponse20069VpnTrafficUplinkPreferences struct {
	// Traffic filters
	TrafficFilters []InlineResponse20069TrafficFilters1 `json:"trafficFilters"`
	// Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
	PreferredUplink string `json:"preferredUplink"`
	// Fail over criterion for uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'
	FailOverCriterion *string `json:"failOverCriterion,omitempty"`
	PerformanceClass *InlineResponse20069PerformanceClass `json:"performanceClass,omitempty"`
}

// NewInlineResponse20069VpnTrafficUplinkPreferences instantiates a new InlineResponse20069VpnTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069VpnTrafficUplinkPreferences(trafficFilters []InlineResponse20069TrafficFilters1, preferredUplink string) *InlineResponse20069VpnTrafficUplinkPreferences {
	this := InlineResponse20069VpnTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewInlineResponse20069VpnTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20069VpnTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069VpnTrafficUplinkPreferencesWithDefaults() *InlineResponse20069VpnTrafficUplinkPreferences {
	this := InlineResponse20069VpnTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20069TrafficFilters1 {
	if o == nil {
		var ret []InlineResponse20069TrafficFilters1
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetTrafficFiltersOk() ([]InlineResponse20069TrafficFilters1, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *InlineResponse20069VpnTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20069TrafficFilters1) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *InlineResponse20069VpnTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

// GetFailOverCriterion returns the FailOverCriterion field value if set, zero value otherwise.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetFailOverCriterion() string {
	if o == nil || isNil(o.FailOverCriterion) {
		var ret string
		return ret
	}
	return *o.FailOverCriterion
}

// GetFailOverCriterionOk returns a tuple with the FailOverCriterion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetFailOverCriterionOk() (*string, bool) {
	if o == nil || isNil(o.FailOverCriterion) {
    return nil, false
	}
	return o.FailOverCriterion, true
}

// HasFailOverCriterion returns a boolean if a field has been set.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) HasFailOverCriterion() bool {
	if o != nil && !isNil(o.FailOverCriterion) {
		return true
	}

	return false
}

// SetFailOverCriterion gets a reference to the given string and assigns it to the FailOverCriterion field.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) SetFailOverCriterion(v string) {
	o.FailOverCriterion = &v
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetPerformanceClass() InlineResponse20069PerformanceClass {
	if o == nil || isNil(o.PerformanceClass) {
		var ret InlineResponse20069PerformanceClass
		return ret
	}
	return *o.PerformanceClass
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) GetPerformanceClassOk() (*InlineResponse20069PerformanceClass, bool) {
	if o == nil || isNil(o.PerformanceClass) {
    return nil, false
	}
	return o.PerformanceClass, true
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) HasPerformanceClass() bool {
	if o != nil && !isNil(o.PerformanceClass) {
		return true
	}

	return false
}

// SetPerformanceClass gets a reference to the given InlineResponse20069PerformanceClass and assigns it to the PerformanceClass field.
func (o *InlineResponse20069VpnTrafficUplinkPreferences) SetPerformanceClass(v InlineResponse20069PerformanceClass) {
	o.PerformanceClass = &v
}

func (o InlineResponse20069VpnTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20069VpnTrafficUplinkPreferences struct {
	value *InlineResponse20069VpnTrafficUplinkPreferences
	isSet bool
}

func (v NullableInlineResponse20069VpnTrafficUplinkPreferences) Get() *InlineResponse20069VpnTrafficUplinkPreferences {
	return v.value
}

func (v *NullableInlineResponse20069VpnTrafficUplinkPreferences) Set(val *InlineResponse20069VpnTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069VpnTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069VpnTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069VpnTrafficUplinkPreferences(val *InlineResponse20069VpnTrafficUplinkPreferences) *NullableInlineResponse20069VpnTrafficUplinkPreferences {
	return &NullableInlineResponse20069VpnTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableInlineResponse20069VpnTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069VpnTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


