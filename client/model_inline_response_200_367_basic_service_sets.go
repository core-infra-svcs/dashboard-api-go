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

// InlineResponse200367BasicServiceSets struct for InlineResponse200367BasicServiceSets
type InlineResponse200367BasicServiceSets struct {
	// Unique identifier for wireless access point.
	Bssid *string `json:"bssid,omitempty"`
	Ssid *InlineResponse200367Ssid `json:"ssid,omitempty"`
	Radio *InlineResponse200367Radio `json:"radio,omitempty"`
}

// NewInlineResponse200367BasicServiceSets instantiates a new InlineResponse200367BasicServiceSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200367BasicServiceSets() *InlineResponse200367BasicServiceSets {
	this := InlineResponse200367BasicServiceSets{}
	return &this
}

// NewInlineResponse200367BasicServiceSetsWithDefaults instantiates a new InlineResponse200367BasicServiceSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200367BasicServiceSetsWithDefaults() *InlineResponse200367BasicServiceSets {
	this := InlineResponse200367BasicServiceSets{}
	return &this
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *InlineResponse200367BasicServiceSets) GetBssid() string {
	if o == nil || isNil(o.Bssid) {
		var ret string
		return ret
	}
	return *o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367BasicServiceSets) GetBssidOk() (*string, bool) {
	if o == nil || isNil(o.Bssid) {
    return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *InlineResponse200367BasicServiceSets) HasBssid() bool {
	if o != nil && !isNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given string and assigns it to the Bssid field.
func (o *InlineResponse200367BasicServiceSets) SetBssid(v string) {
	o.Bssid = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse200367BasicServiceSets) GetSsid() InlineResponse200367Ssid {
	if o == nil || isNil(o.Ssid) {
		var ret InlineResponse200367Ssid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367BasicServiceSets) GetSsidOk() (*InlineResponse200367Ssid, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse200367BasicServiceSets) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given InlineResponse200367Ssid and assigns it to the Ssid field.
func (o *InlineResponse200367BasicServiceSets) SetSsid(v InlineResponse200367Ssid) {
	o.Ssid = &v
}

// GetRadio returns the Radio field value if set, zero value otherwise.
func (o *InlineResponse200367BasicServiceSets) GetRadio() InlineResponse200367Radio {
	if o == nil || isNil(o.Radio) {
		var ret InlineResponse200367Radio
		return ret
	}
	return *o.Radio
}

// GetRadioOk returns a tuple with the Radio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200367BasicServiceSets) GetRadioOk() (*InlineResponse200367Radio, bool) {
	if o == nil || isNil(o.Radio) {
    return nil, false
	}
	return o.Radio, true
}

// HasRadio returns a boolean if a field has been set.
func (o *InlineResponse200367BasicServiceSets) HasRadio() bool {
	if o != nil && !isNil(o.Radio) {
		return true
	}

	return false
}

// SetRadio gets a reference to the given InlineResponse200367Radio and assigns it to the Radio field.
func (o *InlineResponse200367BasicServiceSets) SetRadio(v InlineResponse200367Radio) {
	o.Radio = &v
}

func (o InlineResponse200367BasicServiceSets) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200367BasicServiceSets struct {
	value *InlineResponse200367BasicServiceSets
	isSet bool
}

func (v NullableInlineResponse200367BasicServiceSets) Get() *InlineResponse200367BasicServiceSets {
	return v.value
}

func (v *NullableInlineResponse200367BasicServiceSets) Set(val *InlineResponse200367BasicServiceSets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200367BasicServiceSets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200367BasicServiceSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200367BasicServiceSets(val *InlineResponse200367BasicServiceSets) *NullableInlineResponse200367BasicServiceSets {
	return &NullableInlineResponse200367BasicServiceSets{value: val, isSet: true}
}

func (v NullableInlineResponse200367BasicServiceSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200367BasicServiceSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


