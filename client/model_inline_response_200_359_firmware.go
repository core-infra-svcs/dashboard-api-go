/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200359Firmware Wireless LAN controller device firmware information
type InlineResponse200359Firmware struct {
	Version *InlineResponse200359FirmwareVersion `json:"version,omitempty"`
}

// NewInlineResponse200359Firmware instantiates a new InlineResponse200359Firmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200359Firmware() *InlineResponse200359Firmware {
	this := InlineResponse200359Firmware{}
	return &this
}

// NewInlineResponse200359FirmwareWithDefaults instantiates a new InlineResponse200359Firmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200359FirmwareWithDefaults() *InlineResponse200359Firmware {
	this := InlineResponse200359Firmware{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse200359Firmware) GetVersion() InlineResponse200359FirmwareVersion {
	if o == nil || isNil(o.Version) {
		var ret InlineResponse200359FirmwareVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359Firmware) GetVersionOk() (*InlineResponse200359FirmwareVersion, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse200359Firmware) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given InlineResponse200359FirmwareVersion and assigns it to the Version field.
func (o *InlineResponse200359Firmware) SetVersion(v InlineResponse200359FirmwareVersion) {
	o.Version = &v
}

func (o InlineResponse200359Firmware) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200359Firmware struct {
	value *InlineResponse200359Firmware
	isSet bool
}

func (v NullableInlineResponse200359Firmware) Get() *InlineResponse200359Firmware {
	return v.value
}

func (v *NullableInlineResponse200359Firmware) Set(val *InlineResponse200359Firmware) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200359Firmware) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200359Firmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200359Firmware(val *InlineResponse200359Firmware) *NullableInlineResponse200359Firmware {
	return &NullableInlineResponse200359Firmware{value: val, isSet: true}
}

func (v NullableInlineResponse200359Firmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200359Firmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


