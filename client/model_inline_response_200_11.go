/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	// The current time
	Ts *string `json:"ts,omitempty"`
	Zones *InlineResponse20011Zones `json:"zones,omitempty"`
}

// NewInlineResponse20011 instantiates a new InlineResponse20011 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011WithDefaults() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20011) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20011) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse20011) SetTs(v string) {
	o.Ts = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *InlineResponse20011) GetZones() InlineResponse20011Zones {
	if o == nil || isNil(o.Zones) {
		var ret InlineResponse20011Zones
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetZonesOk() (*InlineResponse20011Zones, bool) {
	if o == nil || isNil(o.Zones) {
    return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *InlineResponse20011) HasZones() bool {
	if o != nil && !isNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given InlineResponse20011Zones and assigns it to the Zones field.
func (o *InlineResponse20011) SetZones(v InlineResponse20011Zones) {
	o.Zones = &v
}

func (o InlineResponse20011) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011 struct {
	value *InlineResponse20011
	isSet bool
}

func (v NullableInlineResponse20011) Get() *InlineResponse20011 {
	return v.value
}

func (v *NullableInlineResponse20011) Set(val *InlineResponse20011) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011(val *InlineResponse20011) *NullableInlineResponse20011 {
	return &NullableInlineResponse20011{value: val, isSet: true}
}

func (v NullableInlineResponse20011) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


