/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2003ApiAuthenticationIpRestrictionsForKeys Details for API-only IP restrictions.
type InlineResponse2003ApiAuthenticationIpRestrictionsForKeys struct {
	// Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Enabled *bool `json:"enabled,omitempty"`
	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	Ranges *[]string `json:"ranges,omitempty"`
}

// NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeys instantiates a new InlineResponse2003ApiAuthenticationIpRestrictionsForKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeys() *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys {
	this := InlineResponse2003ApiAuthenticationIpRestrictionsForKeys{}
	return &this
}

// NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeysWithDefaults instantiates a new InlineResponse2003ApiAuthenticationIpRestrictionsForKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeysWithDefaults() *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys {
	this := InlineResponse2003ApiAuthenticationIpRestrictionsForKeys{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetRanges() []string {
	if o == nil || o.Ranges == nil {
		var ret []string
		return ret
	}
	return *o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetRangesOk() (*[]string, bool) {
	if o == nil || o.Ranges == nil {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []string and assigns it to the Ranges field.
func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) SetRanges(v []string) {
	o.Ranges = &v
}

func (o InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys struct {
	value *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys
	isSet bool
}

func (v NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys) Get() *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys {
	return v.value
}

func (v *NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys) Set(val *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys(val *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) *NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys {
	return &NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys{value: val, isSet: true}
}

func (v NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003ApiAuthenticationIpRestrictionsForKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


