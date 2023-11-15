/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023MandatoryDhcp Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
type InlineResponse20023MandatoryDhcp struct {
	// Enable Mandatory DHCP on single LAN.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20023MandatoryDhcp instantiates a new InlineResponse20023MandatoryDhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023MandatoryDhcp() *InlineResponse20023MandatoryDhcp {
	this := InlineResponse20023MandatoryDhcp{}
	return &this
}

// NewInlineResponse20023MandatoryDhcpWithDefaults instantiates a new InlineResponse20023MandatoryDhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023MandatoryDhcpWithDefaults() *InlineResponse20023MandatoryDhcp {
	this := InlineResponse20023MandatoryDhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20023MandatoryDhcp) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023MandatoryDhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20023MandatoryDhcp) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20023MandatoryDhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20023MandatoryDhcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023MandatoryDhcp struct {
	value *InlineResponse20023MandatoryDhcp
	isSet bool
}

func (v NullableInlineResponse20023MandatoryDhcp) Get() *InlineResponse20023MandatoryDhcp {
	return v.value
}

func (v *NullableInlineResponse20023MandatoryDhcp) Set(val *InlineResponse20023MandatoryDhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023MandatoryDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023MandatoryDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023MandatoryDhcp(val *InlineResponse20023MandatoryDhcp) *NullableInlineResponse20023MandatoryDhcp {
	return &NullableInlineResponse20023MandatoryDhcp{value: val, isSet: true}
}

func (v NullableInlineResponse20023MandatoryDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023MandatoryDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

