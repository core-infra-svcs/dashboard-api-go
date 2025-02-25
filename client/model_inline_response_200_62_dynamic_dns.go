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

// InlineResponse20062DynamicDns Dynamic DNS settings for a network
type InlineResponse20062DynamicDns struct {
	// Dynamic DNS enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Dynamic DNS url prefix. DDNS must be enabled to update
	Prefix *string `json:"prefix,omitempty"`
	// Dynamic DNS url. DDNS must be enabled to update
	Url *string `json:"url,omitempty"`
}

// NewInlineResponse20062DynamicDns instantiates a new InlineResponse20062DynamicDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20062DynamicDns() *InlineResponse20062DynamicDns {
	this := InlineResponse20062DynamicDns{}
	return &this
}

// NewInlineResponse20062DynamicDnsWithDefaults instantiates a new InlineResponse20062DynamicDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20062DynamicDnsWithDefaults() *InlineResponse20062DynamicDns {
	this := InlineResponse20062DynamicDns{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20062DynamicDns) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062DynamicDns) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20062DynamicDns) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20062DynamicDns) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InlineResponse20062DynamicDns) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062DynamicDns) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InlineResponse20062DynamicDns) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InlineResponse20062DynamicDns) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20062DynamicDns) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062DynamicDns) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20062DynamicDns) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20062DynamicDns) SetUrl(v string) {
	o.Url = &v
}

func (o InlineResponse20062DynamicDns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20062DynamicDns struct {
	value *InlineResponse20062DynamicDns
	isSet bool
}

func (v NullableInlineResponse20062DynamicDns) Get() *InlineResponse20062DynamicDns {
	return v.value
}

func (v *NullableInlineResponse20062DynamicDns) Set(val *InlineResponse20062DynamicDns) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20062DynamicDns) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20062DynamicDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20062DynamicDns(val *InlineResponse20062DynamicDns) *NullableInlineResponse20062DynamicDns {
	return &NullableInlineResponse20062DynamicDns{value: val, isSet: true}
}

func (v NullableInlineResponse20062DynamicDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20062DynamicDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


