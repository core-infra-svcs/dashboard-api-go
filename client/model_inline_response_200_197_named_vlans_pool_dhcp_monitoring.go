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

// InlineResponse200197NamedVlansPoolDhcpMonitoring Named VLAN Pool DHCP Monitoring settings.
type InlineResponse200197NamedVlansPoolDhcpMonitoring struct {
	// Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP
	Enabled *bool `json:"enabled,omitempty"`
	// The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Duration *int32 `json:"duration,omitempty"`
}

// NewInlineResponse200197NamedVlansPoolDhcpMonitoring instantiates a new InlineResponse200197NamedVlansPoolDhcpMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200197NamedVlansPoolDhcpMonitoring() *InlineResponse200197NamedVlansPoolDhcpMonitoring {
	this := InlineResponse200197NamedVlansPoolDhcpMonitoring{}
	return &this
}

// NewInlineResponse200197NamedVlansPoolDhcpMonitoringWithDefaults instantiates a new InlineResponse200197NamedVlansPoolDhcpMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200197NamedVlansPoolDhcpMonitoringWithDefaults() *InlineResponse200197NamedVlansPoolDhcpMonitoring {
	this := InlineResponse200197NamedVlansPoolDhcpMonitoring{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse200197NamedVlansPoolDhcpMonitoring) SetDuration(v int32) {
	o.Duration = &v
}

func (o InlineResponse200197NamedVlansPoolDhcpMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200197NamedVlansPoolDhcpMonitoring struct {
	value *InlineResponse200197NamedVlansPoolDhcpMonitoring
	isSet bool
}

func (v NullableInlineResponse200197NamedVlansPoolDhcpMonitoring) Get() *InlineResponse200197NamedVlansPoolDhcpMonitoring {
	return v.value
}

func (v *NullableInlineResponse200197NamedVlansPoolDhcpMonitoring) Set(val *InlineResponse200197NamedVlansPoolDhcpMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200197NamedVlansPoolDhcpMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200197NamedVlansPoolDhcpMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200197NamedVlansPoolDhcpMonitoring(val *InlineResponse200197NamedVlansPoolDhcpMonitoring) *NullableInlineResponse200197NamedVlansPoolDhcpMonitoring {
	return &NullableInlineResponse200197NamedVlansPoolDhcpMonitoring{value: val, isSet: true}
}

func (v NullableInlineResponse200197NamedVlansPoolDhcpMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200197NamedVlansPoolDhcpMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


