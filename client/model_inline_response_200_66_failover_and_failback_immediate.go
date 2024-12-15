/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20066FailoverAndFailbackImmediate Immediate WAN failover and failback
type InlineResponse20066FailoverAndFailbackImmediate struct {
	// Whether immediate WAN failover and failback is enabled
	Enabled bool `json:"enabled"`
}

// NewInlineResponse20066FailoverAndFailbackImmediate instantiates a new InlineResponse20066FailoverAndFailbackImmediate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066FailoverAndFailbackImmediate(enabled bool) *InlineResponse20066FailoverAndFailbackImmediate {
	this := InlineResponse20066FailoverAndFailbackImmediate{}
	this.Enabled = enabled
	return &this
}

// NewInlineResponse20066FailoverAndFailbackImmediateWithDefaults instantiates a new InlineResponse20066FailoverAndFailbackImmediate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066FailoverAndFailbackImmediateWithDefaults() *InlineResponse20066FailoverAndFailbackImmediate {
	this := InlineResponse20066FailoverAndFailbackImmediate{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineResponse20066FailoverAndFailbackImmediate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066FailoverAndFailbackImmediate) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineResponse20066FailoverAndFailbackImmediate) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineResponse20066FailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20066FailoverAndFailbackImmediate struct {
	value *InlineResponse20066FailoverAndFailbackImmediate
	isSet bool
}

func (v NullableInlineResponse20066FailoverAndFailbackImmediate) Get() *InlineResponse20066FailoverAndFailbackImmediate {
	return v.value
}

func (v *NullableInlineResponse20066FailoverAndFailbackImmediate) Set(val *InlineResponse20066FailoverAndFailbackImmediate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066FailoverAndFailbackImmediate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066FailoverAndFailbackImmediate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066FailoverAndFailbackImmediate(val *InlineResponse20066FailoverAndFailbackImmediate) *NullableInlineResponse20066FailoverAndFailbackImmediate {
	return &NullableInlineResponse20066FailoverAndFailbackImmediate{value: val, isSet: true}
}

func (v NullableInlineResponse20066FailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066FailoverAndFailbackImmediate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


