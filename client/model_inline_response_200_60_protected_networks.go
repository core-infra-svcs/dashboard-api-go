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

// InlineResponse20060ProtectedNetworks Networks included in and excluded from the detection engine
type InlineResponse20060ProtectedNetworks struct {
	// Whether special IPv4 addresses should be used (see: https://tools.ietf.org/html/rfc5735)
	UseDefault *bool `json:"useDefault,omitempty"`
	// List of IP addresses or subnets being protected
	IncludedCidr []string `json:"includedCidr,omitempty"`
	// List of IP addresses or subnets being excluded from protection
	ExcludedCidr []string `json:"excludedCidr,omitempty"`
}

// NewInlineResponse20060ProtectedNetworks instantiates a new InlineResponse20060ProtectedNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060ProtectedNetworks() *InlineResponse20060ProtectedNetworks {
	this := InlineResponse20060ProtectedNetworks{}
	return &this
}

// NewInlineResponse20060ProtectedNetworksWithDefaults instantiates a new InlineResponse20060ProtectedNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060ProtectedNetworksWithDefaults() *InlineResponse20060ProtectedNetworks {
	this := InlineResponse20060ProtectedNetworks{}
	return &this
}

// GetUseDefault returns the UseDefault field value if set, zero value otherwise.
func (o *InlineResponse20060ProtectedNetworks) GetUseDefault() bool {
	if o == nil || isNil(o.UseDefault) {
		var ret bool
		return ret
	}
	return *o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProtectedNetworks) GetUseDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefault) {
    return nil, false
	}
	return o.UseDefault, true
}

// HasUseDefault returns a boolean if a field has been set.
func (o *InlineResponse20060ProtectedNetworks) HasUseDefault() bool {
	if o != nil && !isNil(o.UseDefault) {
		return true
	}

	return false
}

// SetUseDefault gets a reference to the given bool and assigns it to the UseDefault field.
func (o *InlineResponse20060ProtectedNetworks) SetUseDefault(v bool) {
	o.UseDefault = &v
}

// GetIncludedCidr returns the IncludedCidr field value if set, zero value otherwise.
func (o *InlineResponse20060ProtectedNetworks) GetIncludedCidr() []string {
	if o == nil || isNil(o.IncludedCidr) {
		var ret []string
		return ret
	}
	return o.IncludedCidr
}

// GetIncludedCidrOk returns a tuple with the IncludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProtectedNetworks) GetIncludedCidrOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedCidr) {
    return nil, false
	}
	return o.IncludedCidr, true
}

// HasIncludedCidr returns a boolean if a field has been set.
func (o *InlineResponse20060ProtectedNetworks) HasIncludedCidr() bool {
	if o != nil && !isNil(o.IncludedCidr) {
		return true
	}

	return false
}

// SetIncludedCidr gets a reference to the given []string and assigns it to the IncludedCidr field.
func (o *InlineResponse20060ProtectedNetworks) SetIncludedCidr(v []string) {
	o.IncludedCidr = v
}

// GetExcludedCidr returns the ExcludedCidr field value if set, zero value otherwise.
func (o *InlineResponse20060ProtectedNetworks) GetExcludedCidr() []string {
	if o == nil || isNil(o.ExcludedCidr) {
		var ret []string
		return ret
	}
	return o.ExcludedCidr
}

// GetExcludedCidrOk returns a tuple with the ExcludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060ProtectedNetworks) GetExcludedCidrOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedCidr) {
    return nil, false
	}
	return o.ExcludedCidr, true
}

// HasExcludedCidr returns a boolean if a field has been set.
func (o *InlineResponse20060ProtectedNetworks) HasExcludedCidr() bool {
	if o != nil && !isNil(o.ExcludedCidr) {
		return true
	}

	return false
}

// SetExcludedCidr gets a reference to the given []string and assigns it to the ExcludedCidr field.
func (o *InlineResponse20060ProtectedNetworks) SetExcludedCidr(v []string) {
	o.ExcludedCidr = v
}

func (o InlineResponse20060ProtectedNetworks) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20060ProtectedNetworks struct {
	value *InlineResponse20060ProtectedNetworks
	isSet bool
}

func (v NullableInlineResponse20060ProtectedNetworks) Get() *InlineResponse20060ProtectedNetworks {
	return v.value
}

func (v *NullableInlineResponse20060ProtectedNetworks) Set(val *InlineResponse20060ProtectedNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060ProtectedNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060ProtectedNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060ProtectedNetworks(val *InlineResponse20060ProtectedNetworks) *NullableInlineResponse20060ProtectedNetworks {
	return &NullableInlineResponse20060ProtectedNetworks{value: val, isSet: true}
}

func (v NullableInlineResponse20060ProtectedNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060ProtectedNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


