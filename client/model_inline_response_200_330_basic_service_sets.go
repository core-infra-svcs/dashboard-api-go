/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200330BasicServiceSets struct for InlineResponse200330BasicServiceSets
type InlineResponse200330BasicServiceSets struct {
	// Unique identifier for wireless access point.
	Bssid *string `json:"bssid,omitempty"`
	Ssid *InlineResponse200330Ssid `json:"ssid,omitempty"`
	Radio *InlineResponse200330Radio `json:"radio,omitempty"`
}

// NewInlineResponse200330BasicServiceSets instantiates a new InlineResponse200330BasicServiceSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200330BasicServiceSets() *InlineResponse200330BasicServiceSets {
	this := InlineResponse200330BasicServiceSets{}
	return &this
}

// NewInlineResponse200330BasicServiceSetsWithDefaults instantiates a new InlineResponse200330BasicServiceSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200330BasicServiceSetsWithDefaults() *InlineResponse200330BasicServiceSets {
	this := InlineResponse200330BasicServiceSets{}
	return &this
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *InlineResponse200330BasicServiceSets) GetBssid() string {
	if o == nil || isNil(o.Bssid) {
		var ret string
		return ret
	}
	return *o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200330BasicServiceSets) GetBssidOk() (*string, bool) {
	if o == nil || isNil(o.Bssid) {
    return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *InlineResponse200330BasicServiceSets) HasBssid() bool {
	if o != nil && !isNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given string and assigns it to the Bssid field.
func (o *InlineResponse200330BasicServiceSets) SetBssid(v string) {
	o.Bssid = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse200330BasicServiceSets) GetSsid() InlineResponse200330Ssid {
	if o == nil || isNil(o.Ssid) {
		var ret InlineResponse200330Ssid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200330BasicServiceSets) GetSsidOk() (*InlineResponse200330Ssid, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse200330BasicServiceSets) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given InlineResponse200330Ssid and assigns it to the Ssid field.
func (o *InlineResponse200330BasicServiceSets) SetSsid(v InlineResponse200330Ssid) {
	o.Ssid = &v
}

// GetRadio returns the Radio field value if set, zero value otherwise.
func (o *InlineResponse200330BasicServiceSets) GetRadio() InlineResponse200330Radio {
	if o == nil || isNil(o.Radio) {
		var ret InlineResponse200330Radio
		return ret
	}
	return *o.Radio
}

// GetRadioOk returns a tuple with the Radio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200330BasicServiceSets) GetRadioOk() (*InlineResponse200330Radio, bool) {
	if o == nil || isNil(o.Radio) {
    return nil, false
	}
	return o.Radio, true
}

// HasRadio returns a boolean if a field has been set.
func (o *InlineResponse200330BasicServiceSets) HasRadio() bool {
	if o != nil && !isNil(o.Radio) {
		return true
	}

	return false
}

// SetRadio gets a reference to the given InlineResponse200330Radio and assigns it to the Radio field.
func (o *InlineResponse200330BasicServiceSets) SetRadio(v InlineResponse200330Radio) {
	o.Radio = &v
}

func (o InlineResponse200330BasicServiceSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Bssid) {
		toSerialize["bssid"] = o.Bssid
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.Radio) {
		toSerialize["radio"] = o.Radio
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200330BasicServiceSets struct {
	value *InlineResponse200330BasicServiceSets
	isSet bool
}

func (v NullableInlineResponse200330BasicServiceSets) Get() *InlineResponse200330BasicServiceSets {
	return v.value
}

func (v *NullableInlineResponse200330BasicServiceSets) Set(val *InlineResponse200330BasicServiceSets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200330BasicServiceSets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200330BasicServiceSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200330BasicServiceSets(val *InlineResponse200330BasicServiceSets) *NullableInlineResponse200330BasicServiceSets {
	return &NullableInlineResponse200330BasicServiceSets{value: val, isSet: true}
}

func (v NullableInlineResponse200330BasicServiceSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200330BasicServiceSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

