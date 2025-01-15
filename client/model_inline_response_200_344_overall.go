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

// InlineResponse200344Overall The overall CPU usage of the wireless LAN controller
type InlineResponse200344Overall struct {
	Usage *InlineResponse200344OverallUsage `json:"usage,omitempty"`
}

// NewInlineResponse200344Overall instantiates a new InlineResponse200344Overall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200344Overall() *InlineResponse200344Overall {
	this := InlineResponse200344Overall{}
	return &this
}

// NewInlineResponse200344OverallWithDefaults instantiates a new InlineResponse200344Overall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200344OverallWithDefaults() *InlineResponse200344Overall {
	this := InlineResponse200344Overall{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200344Overall) GetUsage() InlineResponse200344OverallUsage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200344OverallUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200344Overall) GetUsageOk() (*InlineResponse200344OverallUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200344Overall) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200344OverallUsage and assigns it to the Usage field.
func (o *InlineResponse200344Overall) SetUsage(v InlineResponse200344OverallUsage) {
	o.Usage = &v
}

func (o InlineResponse200344Overall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200344Overall struct {
	value *InlineResponse200344Overall
	isSet bool
}

func (v NullableInlineResponse200344Overall) Get() *InlineResponse200344Overall {
	return v.value
}

func (v *NullableInlineResponse200344Overall) Set(val *InlineResponse200344Overall) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200344Overall) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200344Overall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200344Overall(val *InlineResponse200344Overall) *NullableInlineResponse200344Overall {
	return &NullableInlineResponse200344Overall{value: val, isSet: true}
}

func (v NullableInlineResponse200344Overall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200344Overall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


