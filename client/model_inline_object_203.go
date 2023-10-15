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

// InlineObject203 struct for InlineObject203
type InlineObject203 struct {
	// Serial of camera
	Serial *string `json:"serial,omitempty"`
	// Note whether credentials were sent successfully
	WirelessCredentialsSent *bool `json:"wirelessCredentialsSent,omitempty"`
}

// NewInlineObject203 instantiates a new InlineObject203 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject203() *InlineObject203 {
	this := InlineObject203{}
	return &this
}

// NewInlineObject203WithDefaults instantiates a new InlineObject203 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject203WithDefaults() *InlineObject203 {
	this := InlineObject203{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineObject203) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineObject203) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineObject203) SetSerial(v string) {
	o.Serial = &v
}

// GetWirelessCredentialsSent returns the WirelessCredentialsSent field value if set, zero value otherwise.
func (o *InlineObject203) GetWirelessCredentialsSent() bool {
	if o == nil || isNil(o.WirelessCredentialsSent) {
		var ret bool
		return ret
	}
	return *o.WirelessCredentialsSent
}

// GetWirelessCredentialsSentOk returns a tuple with the WirelessCredentialsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetWirelessCredentialsSentOk() (*bool, bool) {
	if o == nil || isNil(o.WirelessCredentialsSent) {
    return nil, false
	}
	return o.WirelessCredentialsSent, true
}

// HasWirelessCredentialsSent returns a boolean if a field has been set.
func (o *InlineObject203) HasWirelessCredentialsSent() bool {
	if o != nil && !isNil(o.WirelessCredentialsSent) {
		return true
	}

	return false
}

// SetWirelessCredentialsSent gets a reference to the given bool and assigns it to the WirelessCredentialsSent field.
func (o *InlineObject203) SetWirelessCredentialsSent(v bool) {
	o.WirelessCredentialsSent = &v
}

func (o InlineObject203) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.WirelessCredentialsSent) {
		toSerialize["wirelessCredentialsSent"] = o.WirelessCredentialsSent
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject203 struct {
	value *InlineObject203
	isSet bool
}

func (v NullableInlineObject203) Get() *InlineObject203 {
	return v.value
}

func (v *NullableInlineObject203) Set(val *InlineObject203) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject203) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject203) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject203(val *InlineObject203) *NullableInlineObject203 {
	return &NullableInlineObject203{value: val, isSet: true}
}

func (v NullableInlineObject203) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject203) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


