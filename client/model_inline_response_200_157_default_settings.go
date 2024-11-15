/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200157DefaultSettings Default multicast setting for entire network. IGMP snooping and Flood unknown       multicast traffic settings are enabled by default.
type InlineResponse200157DefaultSettings struct {
	// IGMP snooping enabled for the entire network
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled,omitempty"`
	// Flood unknown multicast traffic enabled for the entire network
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`
}

// NewInlineResponse200157DefaultSettings instantiates a new InlineResponse200157DefaultSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200157DefaultSettings() *InlineResponse200157DefaultSettings {
	this := InlineResponse200157DefaultSettings{}
	return &this
}

// NewInlineResponse200157DefaultSettingsWithDefaults instantiates a new InlineResponse200157DefaultSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200157DefaultSettingsWithDefaults() *InlineResponse200157DefaultSettings {
	this := InlineResponse200157DefaultSettings{}
	return &this
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200157DefaultSettings) GetIgmpSnoopingEnabled() bool {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200157DefaultSettings) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
    return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200157DefaultSettings) HasIgmpSnoopingEnabled() bool {
	if o != nil && !isNil(o.IgmpSnoopingEnabled) {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *InlineResponse200157DefaultSettings) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *InlineResponse200157DefaultSettings) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200157DefaultSettings) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
    return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *InlineResponse200157DefaultSettings) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *InlineResponse200157DefaultSettings) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o InlineResponse200157DefaultSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IgmpSnoopingEnabled) {
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200157DefaultSettings struct {
	value *InlineResponse200157DefaultSettings
	isSet bool
}

func (v NullableInlineResponse200157DefaultSettings) Get() *InlineResponse200157DefaultSettings {
	return v.value
}

func (v *NullableInlineResponse200157DefaultSettings) Set(val *InlineResponse200157DefaultSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200157DefaultSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200157DefaultSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200157DefaultSettings(val *InlineResponse200157DefaultSettings) *NullableInlineResponse200157DefaultSettings {
	return &NullableInlineResponse200157DefaultSettings{value: val, isSet: true}
}

func (v NullableInlineResponse200157DefaultSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200157DefaultSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


