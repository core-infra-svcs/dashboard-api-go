/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	// The serial number for the device
	Serial *string `json:"serial,omitempty"`
	ConnectionStats *InlineResponse2009ConnectionStats `json:"connectionStats,omitempty"`
}

// NewInlineResponse2009 instantiates a new InlineResponse2009 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009WithDefaults() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2009) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2009) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2009) SetSerial(v string) {
	o.Serial = &v
}

// GetConnectionStats returns the ConnectionStats field value if set, zero value otherwise.
func (o *InlineResponse2009) GetConnectionStats() InlineResponse2009ConnectionStats {
	if o == nil || isNil(o.ConnectionStats) {
		var ret InlineResponse2009ConnectionStats
		return ret
	}
	return *o.ConnectionStats
}

// GetConnectionStatsOk returns a tuple with the ConnectionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetConnectionStatsOk() (*InlineResponse2009ConnectionStats, bool) {
	if o == nil || isNil(o.ConnectionStats) {
    return nil, false
	}
	return o.ConnectionStats, true
}

// HasConnectionStats returns a boolean if a field has been set.
func (o *InlineResponse2009) HasConnectionStats() bool {
	if o != nil && !isNil(o.ConnectionStats) {
		return true
	}

	return false
}

// SetConnectionStats gets a reference to the given InlineResponse2009ConnectionStats and assigns it to the ConnectionStats field.
func (o *InlineResponse2009) SetConnectionStats(v InlineResponse2009ConnectionStats) {
	o.ConnectionStats = &v
}

func (o InlineResponse2009) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ConnectionStats) {
		toSerialize["connectionStats"] = o.ConnectionStats
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009 struct {
	value *InlineResponse2009
	isSet bool
}

func (v NullableInlineResponse2009) Get() *InlineResponse2009 {
	return v.value
}

func (v *NullableInlineResponse2009) Set(val *InlineResponse2009) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009(val *InlineResponse2009) *NullableInlineResponse2009 {
	return &NullableInlineResponse2009{value: val, isSet: true}
}

func (v NullableInlineResponse2009) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


