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

// InlineObject238 struct for InlineObject238
type InlineObject238 struct {
	// A list of Meraki Serials to migrate
	Serials []string `json:"serials"`
	// The controller or management mode to which the devices will be migrated
	Target string `json:"target"`
}

// NewInlineObject238 instantiates a new InlineObject238 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject238(serials []string, target string) *InlineObject238 {
	this := InlineObject238{}
	this.Serials = serials
	this.Target = target
	return &this
}

// NewInlineObject238WithDefaults instantiates a new InlineObject238 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject238WithDefaults() *InlineObject238 {
	this := InlineObject238{}
	return &this
}

// GetSerials returns the Serials field value
func (o *InlineObject238) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineObject238) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineObject238) SetSerials(v []string) {
	o.Serials = v
}

// GetTarget returns the Target field value
func (o *InlineObject238) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *InlineObject238) GetTargetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *InlineObject238) SetTarget(v string) {
	o.Target = v
}

func (o InlineObject238) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serials"] = o.Serials
	}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject238 struct {
	value *InlineObject238
	isSet bool
}

func (v NullableInlineObject238) Get() *InlineObject238 {
	return v.value
}

func (v *NullableInlineObject238) Set(val *InlineObject238) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject238) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject238) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject238(val *InlineObject238) *NullableInlineObject238 {
	return &NullableInlineObject238{value: val, isSet: true}
}

func (v NullableInlineObject238) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject238) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


