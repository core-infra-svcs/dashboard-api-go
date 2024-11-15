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

// InlineResponse20057ProtectedNetworks Networks included in and excluded from the detection engine
type InlineResponse20057ProtectedNetworks struct {
	// Whether special IPv4 addresses should be used (see: https://tools.ietf.org/html/rfc5735)
	UseDefault *bool `json:"useDefault,omitempty"`
	// List of IP addresses or subnets being protected
	IncludedCidr []string `json:"includedCidr,omitempty"`
	// List of IP addresses or subnets being excluded from protection
	ExcludedCidr []string `json:"excludedCidr,omitempty"`
}

// NewInlineResponse20057ProtectedNetworks instantiates a new InlineResponse20057ProtectedNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20057ProtectedNetworks() *InlineResponse20057ProtectedNetworks {
	this := InlineResponse20057ProtectedNetworks{}
	return &this
}

// NewInlineResponse20057ProtectedNetworksWithDefaults instantiates a new InlineResponse20057ProtectedNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20057ProtectedNetworksWithDefaults() *InlineResponse20057ProtectedNetworks {
	this := InlineResponse20057ProtectedNetworks{}
	return &this
}

// GetUseDefault returns the UseDefault field value if set, zero value otherwise.
func (o *InlineResponse20057ProtectedNetworks) GetUseDefault() bool {
	if o == nil || isNil(o.UseDefault) {
		var ret bool
		return ret
	}
	return *o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057ProtectedNetworks) GetUseDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefault) {
    return nil, false
	}
	return o.UseDefault, true
}

// HasUseDefault returns a boolean if a field has been set.
func (o *InlineResponse20057ProtectedNetworks) HasUseDefault() bool {
	if o != nil && !isNil(o.UseDefault) {
		return true
	}

	return false
}

// SetUseDefault gets a reference to the given bool and assigns it to the UseDefault field.
func (o *InlineResponse20057ProtectedNetworks) SetUseDefault(v bool) {
	o.UseDefault = &v
}

// GetIncludedCidr returns the IncludedCidr field value if set, zero value otherwise.
func (o *InlineResponse20057ProtectedNetworks) GetIncludedCidr() []string {
	if o == nil || isNil(o.IncludedCidr) {
		var ret []string
		return ret
	}
	return o.IncludedCidr
}

// GetIncludedCidrOk returns a tuple with the IncludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057ProtectedNetworks) GetIncludedCidrOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedCidr) {
    return nil, false
	}
	return o.IncludedCidr, true
}

// HasIncludedCidr returns a boolean if a field has been set.
func (o *InlineResponse20057ProtectedNetworks) HasIncludedCidr() bool {
	if o != nil && !isNil(o.IncludedCidr) {
		return true
	}

	return false
}

// SetIncludedCidr gets a reference to the given []string and assigns it to the IncludedCidr field.
func (o *InlineResponse20057ProtectedNetworks) SetIncludedCidr(v []string) {
	o.IncludedCidr = v
}

// GetExcludedCidr returns the ExcludedCidr field value if set, zero value otherwise.
func (o *InlineResponse20057ProtectedNetworks) GetExcludedCidr() []string {
	if o == nil || isNil(o.ExcludedCidr) {
		var ret []string
		return ret
	}
	return o.ExcludedCidr
}

// GetExcludedCidrOk returns a tuple with the ExcludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057ProtectedNetworks) GetExcludedCidrOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedCidr) {
    return nil, false
	}
	return o.ExcludedCidr, true
}

// HasExcludedCidr returns a boolean if a field has been set.
func (o *InlineResponse20057ProtectedNetworks) HasExcludedCidr() bool {
	if o != nil && !isNil(o.ExcludedCidr) {
		return true
	}

	return false
}

// SetExcludedCidr gets a reference to the given []string and assigns it to the ExcludedCidr field.
func (o *InlineResponse20057ProtectedNetworks) SetExcludedCidr(v []string) {
	o.ExcludedCidr = v
}

func (o InlineResponse20057ProtectedNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UseDefault) {
		toSerialize["useDefault"] = o.UseDefault
	}
	if !isNil(o.IncludedCidr) {
		toSerialize["includedCidr"] = o.IncludedCidr
	}
	if !isNil(o.ExcludedCidr) {
		toSerialize["excludedCidr"] = o.ExcludedCidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20057ProtectedNetworks struct {
	value *InlineResponse20057ProtectedNetworks
	isSet bool
}

func (v NullableInlineResponse20057ProtectedNetworks) Get() *InlineResponse20057ProtectedNetworks {
	return v.value
}

func (v *NullableInlineResponse20057ProtectedNetworks) Set(val *InlineResponse20057ProtectedNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20057ProtectedNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20057ProtectedNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20057ProtectedNetworks(val *InlineResponse20057ProtectedNetworks) *NullableInlineResponse20057ProtectedNetworks {
	return &NullableInlineResponse20057ProtectedNetworks{value: val, isSet: true}
}

func (v NullableInlineResponse20057ProtectedNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20057ProtectedNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


