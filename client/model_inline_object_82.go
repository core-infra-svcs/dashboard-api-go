/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject82 struct for InlineObject82
type InlineObject82 struct {
	// The wifiMacs of the devices to be checked-in.
	WifiMacs *[]string `json:"wifiMacs,omitempty"`
	// The ids of the devices to be checked-in.
	Ids *[]string `json:"ids,omitempty"`
	// The serials of the devices to be checked-in.
	Serials *[]string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be checked-in.
	Scope *[]string `json:"scope,omitempty"`
}

// NewInlineObject82 instantiates a new InlineObject82 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject82() *InlineObject82 {
	this := InlineObject82{}
	return &this
}

// NewInlineObject82WithDefaults instantiates a new InlineObject82 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject82WithDefaults() *InlineObject82 {
	this := InlineObject82{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *InlineObject82) GetWifiMacs() []string {
	if o == nil || o.WifiMacs == nil {
		var ret []string
		return ret
	}
	return *o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetWifiMacsOk() (*[]string, bool) {
	if o == nil || o.WifiMacs == nil {
		return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *InlineObject82) HasWifiMacs() bool {
	if o != nil && o.WifiMacs != nil {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *InlineObject82) SetWifiMacs(v []string) {
	o.WifiMacs = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject82) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject82) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject82) SetIds(v []string) {
	o.Ids = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject82) GetSerials() []string {
	if o == nil || o.Serials == nil {
		var ret []string
		return ret
	}
	return *o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetSerialsOk() (*[]string, bool) {
	if o == nil || o.Serials == nil {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject82) HasSerials() bool {
	if o != nil && o.Serials != nil {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject82) SetSerials(v []string) {
	o.Serials = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject82) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject82) GetScopeOk() (*[]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject82) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *InlineObject82) SetScope(v []string) {
	o.Scope = &v
}

func (o InlineObject82) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WifiMacs != nil {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Serials != nil {
		toSerialize["serials"] = o.Serials
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject82 struct {
	value *InlineObject82
	isSet bool
}

func (v NullableInlineObject82) Get() *InlineObject82 {
	return v.value
}

func (v *NullableInlineObject82) Set(val *InlineObject82) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject82) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject82) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject82(val *InlineObject82) *NullableInlineObject82 {
	return &NullableInlineObject82{value: val, isSet: true}
}

func (v NullableInlineObject82) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject82) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


