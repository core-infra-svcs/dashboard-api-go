/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20068FailoverAndFailbackImmediate Immediate WAN failover and failback
type InlineResponse20068FailoverAndFailbackImmediate struct {
	// Whether immediate WAN failover and failback is enabled
	Enabled bool `json:"enabled"`
}

// NewInlineResponse20068FailoverAndFailbackImmediate instantiates a new InlineResponse20068FailoverAndFailbackImmediate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20068FailoverAndFailbackImmediate(enabled bool) *InlineResponse20068FailoverAndFailbackImmediate {
	this := InlineResponse20068FailoverAndFailbackImmediate{}
	this.Enabled = enabled
	return &this
}

// NewInlineResponse20068FailoverAndFailbackImmediateWithDefaults instantiates a new InlineResponse20068FailoverAndFailbackImmediate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20068FailoverAndFailbackImmediateWithDefaults() *InlineResponse20068FailoverAndFailbackImmediate {
	this := InlineResponse20068FailoverAndFailbackImmediate{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineResponse20068FailoverAndFailbackImmediate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20068FailoverAndFailbackImmediate) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineResponse20068FailoverAndFailbackImmediate) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineResponse20068FailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20068FailoverAndFailbackImmediate struct {
	value *InlineResponse20068FailoverAndFailbackImmediate
	isSet bool
}

func (v NullableInlineResponse20068FailoverAndFailbackImmediate) Get() *InlineResponse20068FailoverAndFailbackImmediate {
	return v.value
}

func (v *NullableInlineResponse20068FailoverAndFailbackImmediate) Set(val *InlineResponse20068FailoverAndFailbackImmediate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20068FailoverAndFailbackImmediate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20068FailoverAndFailbackImmediate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20068FailoverAndFailbackImmediate(val *InlineResponse20068FailoverAndFailbackImmediate) *NullableInlineResponse20068FailoverAndFailbackImmediate {
	return &NullableInlineResponse20068FailoverAndFailbackImmediate{value: val, isSet: true}
}

func (v NullableInlineResponse20068FailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20068FailoverAndFailbackImmediate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


