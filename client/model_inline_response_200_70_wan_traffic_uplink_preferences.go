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

// InlineResponse20070WanTrafficUplinkPreferences struct for InlineResponse20070WanTrafficUplinkPreferences
type InlineResponse20070WanTrafficUplinkPreferences struct {
	// Traffic filters
	TrafficFilters []InlineResponse20070TrafficFilters `json:"trafficFilters"`
	// Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
	PreferredUplink string `json:"preferredUplink"`
}

// NewInlineResponse20070WanTrafficUplinkPreferences instantiates a new InlineResponse20070WanTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070WanTrafficUplinkPreferences(trafficFilters []InlineResponse20070TrafficFilters, preferredUplink string) *InlineResponse20070WanTrafficUplinkPreferences {
	this := InlineResponse20070WanTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewInlineResponse20070WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20070WanTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20070WanTrafficUplinkPreferences {
	this := InlineResponse20070WanTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *InlineResponse20070WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20070TrafficFilters {
	if o == nil {
		var ret []InlineResponse20070TrafficFilters
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20070WanTrafficUplinkPreferences) GetTrafficFiltersOk() ([]InlineResponse20070TrafficFilters, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *InlineResponse20070WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20070TrafficFilters) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *InlineResponse20070WanTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20070WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *InlineResponse20070WanTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

func (o InlineResponse20070WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20070WanTrafficUplinkPreferences struct {
	value *InlineResponse20070WanTrafficUplinkPreferences
	isSet bool
}

func (v NullableInlineResponse20070WanTrafficUplinkPreferences) Get() *InlineResponse20070WanTrafficUplinkPreferences {
	return v.value
}

func (v *NullableInlineResponse20070WanTrafficUplinkPreferences) Set(val *InlineResponse20070WanTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070WanTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070WanTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070WanTrafficUplinkPreferences(val *InlineResponse20070WanTrafficUplinkPreferences) *NullableInlineResponse20070WanTrafficUplinkPreferences {
	return &NullableInlineResponse20070WanTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableInlineResponse20070WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070WanTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


