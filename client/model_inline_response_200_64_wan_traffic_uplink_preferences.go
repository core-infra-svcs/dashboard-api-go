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

// InlineResponse20064WanTrafficUplinkPreferences struct for InlineResponse20064WanTrafficUplinkPreferences
type InlineResponse20064WanTrafficUplinkPreferences struct {
	// Traffic filters
	TrafficFilters []InlineResponse20064TrafficFilters `json:"trafficFilters"`
	// Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
	PreferredUplink string `json:"preferredUplink"`
}

// NewInlineResponse20064WanTrafficUplinkPreferences instantiates a new InlineResponse20064WanTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20064WanTrafficUplinkPreferences(trafficFilters []InlineResponse20064TrafficFilters, preferredUplink string) *InlineResponse20064WanTrafficUplinkPreferences {
	this := InlineResponse20064WanTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewInlineResponse20064WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20064WanTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20064WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20064WanTrafficUplinkPreferences {
	this := InlineResponse20064WanTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *InlineResponse20064WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20064TrafficFilters {
	if o == nil {
		var ret []InlineResponse20064TrafficFilters
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064WanTrafficUplinkPreferences) GetTrafficFiltersOk() ([]InlineResponse20064TrafficFilters, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *InlineResponse20064WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20064TrafficFilters) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *InlineResponse20064WanTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *InlineResponse20064WanTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

func (o InlineResponse20064WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20064WanTrafficUplinkPreferences struct {
	value *InlineResponse20064WanTrafficUplinkPreferences
	isSet bool
}

func (v NullableInlineResponse20064WanTrafficUplinkPreferences) Get() *InlineResponse20064WanTrafficUplinkPreferences {
	return v.value
}

func (v *NullableInlineResponse20064WanTrafficUplinkPreferences) Set(val *InlineResponse20064WanTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20064WanTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20064WanTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20064WanTrafficUplinkPreferences(val *InlineResponse20064WanTrafficUplinkPreferences) *NullableInlineResponse20064WanTrafficUplinkPreferences {
	return &NullableInlineResponse20064WanTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableInlineResponse20064WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20064WanTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


