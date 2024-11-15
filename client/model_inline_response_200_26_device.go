/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026Device Associated device information
type InlineResponse20026Device struct {
	// Dashboard Url for the device
	Url *string `json:"url,omitempty"`
}

// NewInlineResponse20026Device instantiates a new InlineResponse20026Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026Device() *InlineResponse20026Device {
	this := InlineResponse20026Device{}
	return &this
}

// NewInlineResponse20026DeviceWithDefaults instantiates a new InlineResponse20026Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026DeviceWithDefaults() *InlineResponse20026Device {
	this := InlineResponse20026Device{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20026Device) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Device) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20026Device) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20026Device) SetUrl(v string) {
	o.Url = &v
}

func (o InlineResponse20026Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026Device struct {
	value *InlineResponse20026Device
	isSet bool
}

func (v NullableInlineResponse20026Device) Get() *InlineResponse20026Device {
	return v.value
}

func (v *NullableInlineResponse20026Device) Set(val *InlineResponse20026Device) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026Device) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026Device) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026Device(val *InlineResponse20026Device) *NullableInlineResponse20026Device {
	return &NullableInlineResponse20026Device{value: val, isSet: true}
}

func (v NullableInlineResponse20026Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026Device) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


