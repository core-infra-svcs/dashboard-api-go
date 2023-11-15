/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026WanTrafficUplinkPreferences struct for InlineResponse20026WanTrafficUplinkPreferences
type InlineResponse20026WanTrafficUplinkPreferences struct {
	// Traffic filters
	TrafficFilters []InlineResponse20026TrafficFilters `json:"trafficFilters"`
	// Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
	PreferredUplink string `json:"preferredUplink"`
}

// NewInlineResponse20026WanTrafficUplinkPreferences instantiates a new InlineResponse20026WanTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026WanTrafficUplinkPreferences(trafficFilters []InlineResponse20026TrafficFilters, preferredUplink string) *InlineResponse20026WanTrafficUplinkPreferences {
	this := InlineResponse20026WanTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewInlineResponse20026WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20026WanTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20026WanTrafficUplinkPreferences {
	this := InlineResponse20026WanTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *InlineResponse20026WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20026TrafficFilters {
	if o == nil {
		var ret []InlineResponse20026TrafficFilters
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026WanTrafficUplinkPreferences) GetTrafficFiltersOk() ([]InlineResponse20026TrafficFilters, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *InlineResponse20026WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20026TrafficFilters) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *InlineResponse20026WanTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *InlineResponse20026WanTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

func (o InlineResponse20026WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026WanTrafficUplinkPreferences struct {
	value *InlineResponse20026WanTrafficUplinkPreferences
	isSet bool
}

func (v NullableInlineResponse20026WanTrafficUplinkPreferences) Get() *InlineResponse20026WanTrafficUplinkPreferences {
	return v.value
}

func (v *NullableInlineResponse20026WanTrafficUplinkPreferences) Set(val *InlineResponse20026WanTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026WanTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026WanTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026WanTrafficUplinkPreferences(val *InlineResponse20026WanTrafficUplinkPreferences) *NullableInlineResponse20026WanTrafficUplinkPreferences {
	return &NullableInlineResponse20026WanTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableInlineResponse20026WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026WanTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

