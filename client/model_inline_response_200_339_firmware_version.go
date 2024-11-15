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

// InlineResponse200339FirmwareVersion Wireless LAN controller firmware version
type InlineResponse200339FirmwareVersion struct {
	// Wireless LAN controller firmware version short name
	ShortName *string `json:"shortName,omitempty"`
}

// NewInlineResponse200339FirmwareVersion instantiates a new InlineResponse200339FirmwareVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339FirmwareVersion() *InlineResponse200339FirmwareVersion {
	this := InlineResponse200339FirmwareVersion{}
	return &this
}

// NewInlineResponse200339FirmwareVersionWithDefaults instantiates a new InlineResponse200339FirmwareVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339FirmwareVersionWithDefaults() *InlineResponse200339FirmwareVersion {
	this := InlineResponse200339FirmwareVersion{}
	return &this
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse200339FirmwareVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339FirmwareVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse200339FirmwareVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse200339FirmwareVersion) SetShortName(v string) {
	o.ShortName = &v
}

func (o InlineResponse200339FirmwareVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339FirmwareVersion struct {
	value *InlineResponse200339FirmwareVersion
	isSet bool
}

func (v NullableInlineResponse200339FirmwareVersion) Get() *InlineResponse200339FirmwareVersion {
	return v.value
}

func (v *NullableInlineResponse200339FirmwareVersion) Set(val *InlineResponse200339FirmwareVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339FirmwareVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339FirmwareVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339FirmwareVersion(val *InlineResponse200339FirmwareVersion) *NullableInlineResponse200339FirmwareVersion {
	return &NullableInlineResponse200339FirmwareVersion{value: val, isSet: true}
}

func (v NullableInlineResponse200339FirmwareVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339FirmwareVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

