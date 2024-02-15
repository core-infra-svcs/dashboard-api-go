/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20035 struct for InlineResponse20035
type InlineResponse20035 struct {
	// The site-to-site VPN mode.
	Mode *string `json:"mode,omitempty"`
	// The list of VPN hubs, in order of preference.
	Hubs []InlineResponse20035Hubs `json:"hubs,omitempty"`
	// The list of subnets and their VPN presence.
	Subnets []InlineResponse20035Subnets `json:"subnets,omitempty"`
}

// NewInlineResponse20035 instantiates a new InlineResponse20035 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035WithDefaults() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse20035) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse20035) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse20035) SetMode(v string) {
	o.Mode = &v
}

// GetHubs returns the Hubs field value if set, zero value otherwise.
func (o *InlineResponse20035) GetHubs() []InlineResponse20035Hubs {
	if o == nil || isNil(o.Hubs) {
		var ret []InlineResponse20035Hubs
		return ret
	}
	return o.Hubs
}

// GetHubsOk returns a tuple with the Hubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetHubsOk() ([]InlineResponse20035Hubs, bool) {
	if o == nil || isNil(o.Hubs) {
    return nil, false
	}
	return o.Hubs, true
}

// HasHubs returns a boolean if a field has been set.
func (o *InlineResponse20035) HasHubs() bool {
	if o != nil && !isNil(o.Hubs) {
		return true
	}

	return false
}

// SetHubs gets a reference to the given []InlineResponse20035Hubs and assigns it to the Hubs field.
func (o *InlineResponse20035) SetHubs(v []InlineResponse20035Hubs) {
	o.Hubs = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *InlineResponse20035) GetSubnets() []InlineResponse20035Subnets {
	if o == nil || isNil(o.Subnets) {
		var ret []InlineResponse20035Subnets
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetSubnetsOk() ([]InlineResponse20035Subnets, bool) {
	if o == nil || isNil(o.Subnets) {
    return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *InlineResponse20035) HasSubnets() bool {
	if o != nil && !isNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []InlineResponse20035Subnets and assigns it to the Subnets field.
func (o *InlineResponse20035) SetSubnets(v []InlineResponse20035Subnets) {
	o.Subnets = v
}

func (o InlineResponse20035) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Hubs) {
		toSerialize["hubs"] = o.Hubs
	}
	if !isNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035 struct {
	value *InlineResponse20035
	isSet bool
}

func (v NullableInlineResponse20035) Get() *InlineResponse20035 {
	return v.value
}

func (v *NullableInlineResponse20035) Set(val *InlineResponse20035) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035(val *InlineResponse20035) *NullableInlineResponse20035 {
	return &NullableInlineResponse20035{value: val, isSet: true}
}

func (v NullableInlineResponse20035) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


