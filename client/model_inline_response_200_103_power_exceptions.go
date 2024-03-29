/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200103PowerExceptions struct for InlineResponse200103PowerExceptions
type InlineResponse200103PowerExceptions struct {
	// Serial number of the switch
	Serial *string `json:"serial,omitempty"`
	// Per switch exception (combined, redundant, useNetworkSetting)
	PowerType *string `json:"powerType,omitempty"`
}

// NewInlineResponse200103PowerExceptions instantiates a new InlineResponse200103PowerExceptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200103PowerExceptions() *InlineResponse200103PowerExceptions {
	this := InlineResponse200103PowerExceptions{}
	return &this
}

// NewInlineResponse200103PowerExceptionsWithDefaults instantiates a new InlineResponse200103PowerExceptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200103PowerExceptionsWithDefaults() *InlineResponse200103PowerExceptions {
	this := InlineResponse200103PowerExceptions{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200103PowerExceptions) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103PowerExceptions) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200103PowerExceptions) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200103PowerExceptions) SetSerial(v string) {
	o.Serial = &v
}

// GetPowerType returns the PowerType field value if set, zero value otherwise.
func (o *InlineResponse200103PowerExceptions) GetPowerType() string {
	if o == nil || isNil(o.PowerType) {
		var ret string
		return ret
	}
	return *o.PowerType
}

// GetPowerTypeOk returns a tuple with the PowerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103PowerExceptions) GetPowerTypeOk() (*string, bool) {
	if o == nil || isNil(o.PowerType) {
    return nil, false
	}
	return o.PowerType, true
}

// HasPowerType returns a boolean if a field has been set.
func (o *InlineResponse200103PowerExceptions) HasPowerType() bool {
	if o != nil && !isNil(o.PowerType) {
		return true
	}

	return false
}

// SetPowerType gets a reference to the given string and assigns it to the PowerType field.
func (o *InlineResponse200103PowerExceptions) SetPowerType(v string) {
	o.PowerType = &v
}

func (o InlineResponse200103PowerExceptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.PowerType) {
		toSerialize["powerType"] = o.PowerType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200103PowerExceptions struct {
	value *InlineResponse200103PowerExceptions
	isSet bool
}

func (v NullableInlineResponse200103PowerExceptions) Get() *InlineResponse200103PowerExceptions {
	return v.value
}

func (v *NullableInlineResponse200103PowerExceptions) Set(val *InlineResponse200103PowerExceptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200103PowerExceptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200103PowerExceptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200103PowerExceptions(val *InlineResponse200103PowerExceptions) *NullableInlineResponse200103PowerExceptions {
	return &NullableInlineResponse200103PowerExceptions{value: val, isSet: true}
}

func (v NullableInlineResponse200103PowerExceptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200103PowerExceptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


