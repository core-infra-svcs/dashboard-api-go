/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20013 struct for InlineResponse20013
type InlineResponse20013 struct {
	// The serial number for the device
	Serial *string `json:"serial,omitempty"`
	ConnectionStats *InlineResponse20013ConnectionStats `json:"connectionStats,omitempty"`
}

// NewInlineResponse20013 instantiates a new InlineResponse20013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20013() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20013WithDefaults() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20013) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20013) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20013) SetSerial(v string) {
	o.Serial = &v
}

// GetConnectionStats returns the ConnectionStats field value if set, zero value otherwise.
func (o *InlineResponse20013) GetConnectionStats() InlineResponse20013ConnectionStats {
	if o == nil || isNil(o.ConnectionStats) {
		var ret InlineResponse20013ConnectionStats
		return ret
	}
	return *o.ConnectionStats
}

// GetConnectionStatsOk returns a tuple with the ConnectionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetConnectionStatsOk() (*InlineResponse20013ConnectionStats, bool) {
	if o == nil || isNil(o.ConnectionStats) {
    return nil, false
	}
	return o.ConnectionStats, true
}

// HasConnectionStats returns a boolean if a field has been set.
func (o *InlineResponse20013) HasConnectionStats() bool {
	if o != nil && !isNil(o.ConnectionStats) {
		return true
	}

	return false
}

// SetConnectionStats gets a reference to the given InlineResponse20013ConnectionStats and assigns it to the ConnectionStats field.
func (o *InlineResponse20013) SetConnectionStats(v InlineResponse20013ConnectionStats) {
	o.ConnectionStats = &v
}

func (o InlineResponse20013) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ConnectionStats) {
		toSerialize["connectionStats"] = o.ConnectionStats
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20013 struct {
	value *InlineResponse20013
	isSet bool
}

func (v NullableInlineResponse20013) Get() *InlineResponse20013 {
	return v.value
}

func (v *NullableInlineResponse20013) Set(val *InlineResponse20013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013(val *InlineResponse20013) *NullableInlineResponse20013 {
	return &NullableInlineResponse20013{value: val, isSet: true}
}

func (v NullableInlineResponse20013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


