/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20025 struct for InlineResponse20025
type InlineResponse20025 struct {
	// Source MAC address
	SourceMac *string `json:"sourceMac,omitempty"`
	// Mapping of ports to lldp and/or cdp information
	Ports *map[string]InlineResponse20025Ports `json:"ports,omitempty"`
}

// NewInlineResponse20025 instantiates a new InlineResponse20025 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025() *InlineResponse20025 {
	this := InlineResponse20025{}
	return &this
}

// NewInlineResponse20025WithDefaults instantiates a new InlineResponse20025 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025WithDefaults() *InlineResponse20025 {
	this := InlineResponse20025{}
	return &this
}

// GetSourceMac returns the SourceMac field value if set, zero value otherwise.
func (o *InlineResponse20025) GetSourceMac() string {
	if o == nil || isNil(o.SourceMac) {
		var ret string
		return ret
	}
	return *o.SourceMac
}

// GetSourceMacOk returns a tuple with the SourceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025) GetSourceMacOk() (*string, bool) {
	if o == nil || isNil(o.SourceMac) {
    return nil, false
	}
	return o.SourceMac, true
}

// HasSourceMac returns a boolean if a field has been set.
func (o *InlineResponse20025) HasSourceMac() bool {
	if o != nil && !isNil(o.SourceMac) {
		return true
	}

	return false
}

// SetSourceMac gets a reference to the given string and assigns it to the SourceMac field.
func (o *InlineResponse20025) SetSourceMac(v string) {
	o.SourceMac = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse20025) GetPorts() map[string]InlineResponse20025Ports {
	if o == nil || isNil(o.Ports) {
		var ret map[string]InlineResponse20025Ports
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025) GetPortsOk() (*map[string]InlineResponse20025Ports, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse20025) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given map[string]InlineResponse20025Ports and assigns it to the Ports field.
func (o *InlineResponse20025) SetPorts(v map[string]InlineResponse20025Ports) {
	o.Ports = &v
}

func (o InlineResponse20025) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceMac) {
		toSerialize["sourceMac"] = o.SourceMac
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025 struct {
	value *InlineResponse20025
	isSet bool
}

func (v NullableInlineResponse20025) Get() *InlineResponse20025 {
	return v.value
}

func (v *NullableInlineResponse20025) Set(val *InlineResponse20025) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025(val *InlineResponse20025) *NullableInlineResponse20025 {
	return &NullableInlineResponse20025{value: val, isSet: true}
}

func (v NullableInlineResponse20025) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


