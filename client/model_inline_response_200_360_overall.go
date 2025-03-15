/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200360Overall The overall CPU usage of the wireless LAN controller
type InlineResponse200360Overall struct {
	Usage *InlineResponse200360OverallUsage `json:"usage,omitempty"`
}

// NewInlineResponse200360Overall instantiates a new InlineResponse200360Overall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200360Overall() *InlineResponse200360Overall {
	this := InlineResponse200360Overall{}
	return &this
}

// NewInlineResponse200360OverallWithDefaults instantiates a new InlineResponse200360Overall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200360OverallWithDefaults() *InlineResponse200360Overall {
	this := InlineResponse200360Overall{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200360Overall) GetUsage() InlineResponse200360OverallUsage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200360OverallUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200360Overall) GetUsageOk() (*InlineResponse200360OverallUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200360Overall) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200360OverallUsage and assigns it to the Usage field.
func (o *InlineResponse200360Overall) SetUsage(v InlineResponse200360OverallUsage) {
	o.Usage = &v
}

func (o InlineResponse200360Overall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200360Overall struct {
	value *InlineResponse200360Overall
	isSet bool
}

func (v NullableInlineResponse200360Overall) Get() *InlineResponse200360Overall {
	return v.value
}

func (v *NullableInlineResponse200360Overall) Set(val *InlineResponse200360Overall) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200360Overall) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200360Overall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200360Overall(val *InlineResponse200360Overall) *NullableInlineResponse200360Overall {
	return &NullableInlineResponse200360Overall{value: val, isSet: true}
}

func (v NullableInlineResponse200360Overall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200360Overall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


