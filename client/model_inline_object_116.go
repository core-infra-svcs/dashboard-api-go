/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject116 struct for InlineObject116
type InlineObject116 struct {
	// The wifiMacs of the devices to be locked.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the devices to be locked.
	Ids []string `json:"ids,omitempty"`
	// The serials of the devices to be locked.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be locked.
	Scope []string `json:"scope,omitempty"`
	// The pin number for locking macOS devices (a six digit number). Required only for macOS devices.
	Pin *int32 `json:"pin,omitempty"`
}

// NewInlineObject116 instantiates a new InlineObject116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject116() *InlineObject116 {
	this := InlineObject116{}
	return &this
}

// NewInlineObject116WithDefaults instantiates a new InlineObject116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject116WithDefaults() *InlineObject116 {
	this := InlineObject116{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *InlineObject116) GetWifiMacs() []string {
	if o == nil || isNil(o.WifiMacs) {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetWifiMacsOk() ([]string, bool) {
	if o == nil || isNil(o.WifiMacs) {
    return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *InlineObject116) HasWifiMacs() bool {
	if o != nil && !isNil(o.WifiMacs) {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *InlineObject116) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject116) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject116) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject116) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject116) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject116) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject116) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject116) GetScope() []string {
	if o == nil || isNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject116) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *InlineObject116) SetScope(v []string) {
	o.Scope = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *InlineObject116) GetPin() int32 {
	if o == nil || isNil(o.Pin) {
		var ret int32
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetPinOk() (*int32, bool) {
	if o == nil || isNil(o.Pin) {
    return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *InlineObject116) HasPin() bool {
	if o != nil && !isNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given int32 and assigns it to the Pin field.
func (o *InlineObject116) SetPin(v int32) {
	o.Pin = &v
}

func (o InlineObject116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMacs) {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject116 struct {
	value *InlineObject116
	isSet bool
}

func (v NullableInlineObject116) Get() *InlineObject116 {
	return v.value
}

func (v *NullableInlineObject116) Set(val *InlineObject116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject116(val *InlineObject116) *NullableInlineObject116 {
	return &NullableInlineObject116{value: val, isSet: true}
}

func (v NullableInlineObject116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


