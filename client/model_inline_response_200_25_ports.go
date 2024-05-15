/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20025Ports lldp and/or cdp information, keyed by port
type InlineResponse20025Ports struct {
	Lldp *InlineResponse20025Lldp `json:"lldp,omitempty"`
	Cdp *InlineResponse20025Cdp `json:"cdp,omitempty"`
}

// NewInlineResponse20025Ports instantiates a new InlineResponse20025Ports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025Ports() *InlineResponse20025Ports {
	this := InlineResponse20025Ports{}
	return &this
}

// NewInlineResponse20025PortsWithDefaults instantiates a new InlineResponse20025Ports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025PortsWithDefaults() *InlineResponse20025Ports {
	this := InlineResponse20025Ports{}
	return &this
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *InlineResponse20025Ports) GetLldp() InlineResponse20025Lldp {
	if o == nil || isNil(o.Lldp) {
		var ret InlineResponse20025Lldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Ports) GetLldpOk() (*InlineResponse20025Lldp, bool) {
	if o == nil || isNil(o.Lldp) {
    return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *InlineResponse20025Ports) HasLldp() bool {
	if o != nil && !isNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given InlineResponse20025Lldp and assigns it to the Lldp field.
func (o *InlineResponse20025Ports) SetLldp(v InlineResponse20025Lldp) {
	o.Lldp = &v
}

// GetCdp returns the Cdp field value if set, zero value otherwise.
func (o *InlineResponse20025Ports) GetCdp() InlineResponse20025Cdp {
	if o == nil || isNil(o.Cdp) {
		var ret InlineResponse20025Cdp
		return ret
	}
	return *o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Ports) GetCdpOk() (*InlineResponse20025Cdp, bool) {
	if o == nil || isNil(o.Cdp) {
    return nil, false
	}
	return o.Cdp, true
}

// HasCdp returns a boolean if a field has been set.
func (o *InlineResponse20025Ports) HasCdp() bool {
	if o != nil && !isNil(o.Cdp) {
		return true
	}

	return false
}

// SetCdp gets a reference to the given InlineResponse20025Cdp and assigns it to the Cdp field.
func (o *InlineResponse20025Ports) SetCdp(v InlineResponse20025Cdp) {
	o.Cdp = &v
}

func (o InlineResponse20025Ports) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	if !isNil(o.Cdp) {
		toSerialize["cdp"] = o.Cdp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025Ports struct {
	value *InlineResponse20025Ports
	isSet bool
}

func (v NullableInlineResponse20025Ports) Get() *InlineResponse20025Ports {
	return v.value
}

func (v *NullableInlineResponse20025Ports) Set(val *InlineResponse20025Ports) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025Ports) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025Ports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025Ports(val *InlineResponse20025Ports) *NullableInlineResponse20025Ports {
	return &NullableInlineResponse20025Ports{value: val, isSet: true}
}

func (v NullableInlineResponse20025Ports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025Ports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


