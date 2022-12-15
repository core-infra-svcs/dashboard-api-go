/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20012MandatoryDhcp Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
type InlineResponse20012MandatoryDhcp struct {
	// Enable Mandatory DHCP on single LAN.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20012MandatoryDhcp instantiates a new InlineResponse20012MandatoryDhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012MandatoryDhcp() *InlineResponse20012MandatoryDhcp {
	this := InlineResponse20012MandatoryDhcp{}
	return &this
}

// NewInlineResponse20012MandatoryDhcpWithDefaults instantiates a new InlineResponse20012MandatoryDhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012MandatoryDhcpWithDefaults() *InlineResponse20012MandatoryDhcp {
	this := InlineResponse20012MandatoryDhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20012MandatoryDhcp) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012MandatoryDhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20012MandatoryDhcp) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20012MandatoryDhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20012MandatoryDhcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012MandatoryDhcp struct {
	value *InlineResponse20012MandatoryDhcp
	isSet bool
}

func (v NullableInlineResponse20012MandatoryDhcp) Get() *InlineResponse20012MandatoryDhcp {
	return v.value
}

func (v *NullableInlineResponse20012MandatoryDhcp) Set(val *InlineResponse20012MandatoryDhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012MandatoryDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012MandatoryDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012MandatoryDhcp(val *InlineResponse20012MandatoryDhcp) *NullableInlineResponse20012MandatoryDhcp {
	return &NullableInlineResponse20012MandatoryDhcp{value: val, isSet: true}
}

func (v NullableInlineResponse20012MandatoryDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012MandatoryDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


