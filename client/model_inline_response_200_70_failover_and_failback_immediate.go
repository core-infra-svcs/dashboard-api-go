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

// InlineResponse20070FailoverAndFailbackImmediate Immediate WAN failover and failback
type InlineResponse20070FailoverAndFailbackImmediate struct {
	// Whether immediate WAN failover and failback is enabled
	Enabled bool `json:"enabled"`
}

// NewInlineResponse20070FailoverAndFailbackImmediate instantiates a new InlineResponse20070FailoverAndFailbackImmediate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070FailoverAndFailbackImmediate(enabled bool) *InlineResponse20070FailoverAndFailbackImmediate {
	this := InlineResponse20070FailoverAndFailbackImmediate{}
	this.Enabled = enabled
	return &this
}

// NewInlineResponse20070FailoverAndFailbackImmediateWithDefaults instantiates a new InlineResponse20070FailoverAndFailbackImmediate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070FailoverAndFailbackImmediateWithDefaults() *InlineResponse20070FailoverAndFailbackImmediate {
	this := InlineResponse20070FailoverAndFailbackImmediate{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineResponse20070FailoverAndFailbackImmediate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20070FailoverAndFailbackImmediate) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineResponse20070FailoverAndFailbackImmediate) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineResponse20070FailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20070FailoverAndFailbackImmediate struct {
	value *InlineResponse20070FailoverAndFailbackImmediate
	isSet bool
}

func (v NullableInlineResponse20070FailoverAndFailbackImmediate) Get() *InlineResponse20070FailoverAndFailbackImmediate {
	return v.value
}

func (v *NullableInlineResponse20070FailoverAndFailbackImmediate) Set(val *InlineResponse20070FailoverAndFailbackImmediate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070FailoverAndFailbackImmediate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070FailoverAndFailbackImmediate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070FailoverAndFailbackImmediate(val *InlineResponse20070FailoverAndFailbackImmediate) *NullableInlineResponse20070FailoverAndFailbackImmediate {
	return &NullableInlineResponse20070FailoverAndFailbackImmediate{value: val, isSet: true}
}

func (v NullableInlineResponse20070FailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070FailoverAndFailbackImmediate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


