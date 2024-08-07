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

// InlineResponse200307BasicServiceSets struct for InlineResponse200307BasicServiceSets
type InlineResponse200307BasicServiceSets struct {
	// Unique identifier for wireless access point.
	Bssid *string `json:"bssid,omitempty"`
	Ssid *InlineResponse200307Ssid `json:"ssid,omitempty"`
	Radio *InlineResponse200307Radio `json:"radio,omitempty"`
}

// NewInlineResponse200307BasicServiceSets instantiates a new InlineResponse200307BasicServiceSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200307BasicServiceSets() *InlineResponse200307BasicServiceSets {
	this := InlineResponse200307BasicServiceSets{}
	return &this
}

// NewInlineResponse200307BasicServiceSetsWithDefaults instantiates a new InlineResponse200307BasicServiceSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200307BasicServiceSetsWithDefaults() *InlineResponse200307BasicServiceSets {
	this := InlineResponse200307BasicServiceSets{}
	return &this
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *InlineResponse200307BasicServiceSets) GetBssid() string {
	if o == nil || isNil(o.Bssid) {
		var ret string
		return ret
	}
	return *o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307BasicServiceSets) GetBssidOk() (*string, bool) {
	if o == nil || isNil(o.Bssid) {
    return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *InlineResponse200307BasicServiceSets) HasBssid() bool {
	if o != nil && !isNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given string and assigns it to the Bssid field.
func (o *InlineResponse200307BasicServiceSets) SetBssid(v string) {
	o.Bssid = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse200307BasicServiceSets) GetSsid() InlineResponse200307Ssid {
	if o == nil || isNil(o.Ssid) {
		var ret InlineResponse200307Ssid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307BasicServiceSets) GetSsidOk() (*InlineResponse200307Ssid, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse200307BasicServiceSets) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given InlineResponse200307Ssid and assigns it to the Ssid field.
func (o *InlineResponse200307BasicServiceSets) SetSsid(v InlineResponse200307Ssid) {
	o.Ssid = &v
}

// GetRadio returns the Radio field value if set, zero value otherwise.
func (o *InlineResponse200307BasicServiceSets) GetRadio() InlineResponse200307Radio {
	if o == nil || isNil(o.Radio) {
		var ret InlineResponse200307Radio
		return ret
	}
	return *o.Radio
}

// GetRadioOk returns a tuple with the Radio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307BasicServiceSets) GetRadioOk() (*InlineResponse200307Radio, bool) {
	if o == nil || isNil(o.Radio) {
    return nil, false
	}
	return o.Radio, true
}

// HasRadio returns a boolean if a field has been set.
func (o *InlineResponse200307BasicServiceSets) HasRadio() bool {
	if o != nil && !isNil(o.Radio) {
		return true
	}

	return false
}

// SetRadio gets a reference to the given InlineResponse200307Radio and assigns it to the Radio field.
func (o *InlineResponse200307BasicServiceSets) SetRadio(v InlineResponse200307Radio) {
	o.Radio = &v
}

func (o InlineResponse200307BasicServiceSets) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200307BasicServiceSets struct {
	value *InlineResponse200307BasicServiceSets
	isSet bool
}

func (v NullableInlineResponse200307BasicServiceSets) Get() *InlineResponse200307BasicServiceSets {
	return v.value
}

func (v *NullableInlineResponse200307BasicServiceSets) Set(val *InlineResponse200307BasicServiceSets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200307BasicServiceSets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200307BasicServiceSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200307BasicServiceSets(val *InlineResponse200307BasicServiceSets) *NullableInlineResponse200307BasicServiceSets {
	return &NullableInlineResponse200307BasicServiceSets{value: val, isSet: true}
}

func (v NullableInlineResponse200307BasicServiceSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200307BasicServiceSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


