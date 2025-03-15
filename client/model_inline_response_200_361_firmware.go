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

// InlineResponse200361Firmware Wireless LAN controller device firmware information
type InlineResponse200361Firmware struct {
	Version *InlineResponse200361FirmwareVersion `json:"version,omitempty"`
}

// NewInlineResponse200361Firmware instantiates a new InlineResponse200361Firmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200361Firmware() *InlineResponse200361Firmware {
	this := InlineResponse200361Firmware{}
	return &this
}

// NewInlineResponse200361FirmwareWithDefaults instantiates a new InlineResponse200361Firmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200361FirmwareWithDefaults() *InlineResponse200361Firmware {
	this := InlineResponse200361Firmware{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse200361Firmware) GetVersion() InlineResponse200361FirmwareVersion {
	if o == nil || isNil(o.Version) {
		var ret InlineResponse200361FirmwareVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200361Firmware) GetVersionOk() (*InlineResponse200361FirmwareVersion, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse200361Firmware) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given InlineResponse200361FirmwareVersion and assigns it to the Version field.
func (o *InlineResponse200361Firmware) SetVersion(v InlineResponse200361FirmwareVersion) {
	o.Version = &v
}

func (o InlineResponse200361Firmware) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200361Firmware struct {
	value *InlineResponse200361Firmware
	isSet bool
}

func (v NullableInlineResponse200361Firmware) Get() *InlineResponse200361Firmware {
	return v.value
}

func (v *NullableInlineResponse200361Firmware) Set(val *InlineResponse200361Firmware) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200361Firmware) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200361Firmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200361Firmware(val *InlineResponse200361Firmware) *NullableInlineResponse200361Firmware {
	return &NullableInlineResponse200361Firmware{value: val, isSet: true}
}

func (v NullableInlineResponse200361Firmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200361Firmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


