/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20015SimFailover SIM Failover settings.
type InlineResponse20015SimFailover struct {
	// Failover to secondary SIM
	Enabled *bool `json:"enabled,omitempty"`
	// Failover timeout in seconds
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewInlineResponse20015SimFailover instantiates a new InlineResponse20015SimFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015SimFailover() *InlineResponse20015SimFailover {
	this := InlineResponse20015SimFailover{}
	return &this
}

// NewInlineResponse20015SimFailoverWithDefaults instantiates a new InlineResponse20015SimFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015SimFailoverWithDefaults() *InlineResponse20015SimFailover {
	this := InlineResponse20015SimFailover{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20015SimFailover) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015SimFailover) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20015SimFailover) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20015SimFailover) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse20015SimFailover) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015SimFailover) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse20015SimFailover) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse20015SimFailover) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o InlineResponse20015SimFailover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20015SimFailover struct {
	value *InlineResponse20015SimFailover
	isSet bool
}

func (v NullableInlineResponse20015SimFailover) Get() *InlineResponse20015SimFailover {
	return v.value
}

func (v *NullableInlineResponse20015SimFailover) Set(val *InlineResponse20015SimFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20015SimFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20015SimFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20015SimFailover(val *InlineResponse20015SimFailover) *NullableInlineResponse20015SimFailover {
	return &NullableInlineResponse20015SimFailover{value: val, isSet: true}
}

func (v NullableInlineResponse20015SimFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20015SimFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


