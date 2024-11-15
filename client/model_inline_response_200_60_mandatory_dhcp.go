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

// InlineResponse20060MandatoryDhcp Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
type InlineResponse20060MandatoryDhcp struct {
	// Enable Mandatory DHCP on single LAN.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20060MandatoryDhcp instantiates a new InlineResponse20060MandatoryDhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060MandatoryDhcp() *InlineResponse20060MandatoryDhcp {
	this := InlineResponse20060MandatoryDhcp{}
	return &this
}

// NewInlineResponse20060MandatoryDhcpWithDefaults instantiates a new InlineResponse20060MandatoryDhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060MandatoryDhcpWithDefaults() *InlineResponse20060MandatoryDhcp {
	this := InlineResponse20060MandatoryDhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20060MandatoryDhcp) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060MandatoryDhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20060MandatoryDhcp) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20060MandatoryDhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20060MandatoryDhcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20060MandatoryDhcp struct {
	value *InlineResponse20060MandatoryDhcp
	isSet bool
}

func (v NullableInlineResponse20060MandatoryDhcp) Get() *InlineResponse20060MandatoryDhcp {
	return v.value
}

func (v *NullableInlineResponse20060MandatoryDhcp) Set(val *InlineResponse20060MandatoryDhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060MandatoryDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060MandatoryDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060MandatoryDhcp(val *InlineResponse20060MandatoryDhcp) *NullableInlineResponse20060MandatoryDhcp {
	return &NullableInlineResponse20060MandatoryDhcp{value: val, isSet: true}
}

func (v NullableInlineResponse20060MandatoryDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060MandatoryDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


