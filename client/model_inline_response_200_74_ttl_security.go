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

// InlineResponse20074TtlSecurity Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks
type InlineResponse20074TtlSecurity struct {
	// Whether BGP TTL security is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20074TtlSecurity instantiates a new InlineResponse20074TtlSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074TtlSecurity() *InlineResponse20074TtlSecurity {
	this := InlineResponse20074TtlSecurity{}
	return &this
}

// NewInlineResponse20074TtlSecurityWithDefaults instantiates a new InlineResponse20074TtlSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074TtlSecurityWithDefaults() *InlineResponse20074TtlSecurity {
	this := InlineResponse20074TtlSecurity{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20074TtlSecurity) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074TtlSecurity) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20074TtlSecurity) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20074TtlSecurity) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20074TtlSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074TtlSecurity struct {
	value *InlineResponse20074TtlSecurity
	isSet bool
}

func (v NullableInlineResponse20074TtlSecurity) Get() *InlineResponse20074TtlSecurity {
	return v.value
}

func (v *NullableInlineResponse20074TtlSecurity) Set(val *InlineResponse20074TtlSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074TtlSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074TtlSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074TtlSecurity(val *InlineResponse20074TtlSecurity) *NullableInlineResponse20074TtlSecurity {
	return &NullableInlineResponse20074TtlSecurity{value: val, isSet: true}
}

func (v NullableInlineResponse20074TtlSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074TtlSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


