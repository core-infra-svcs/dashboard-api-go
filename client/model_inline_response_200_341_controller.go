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

// InlineResponse200341Controller Associated wireless controller
type InlineResponse200341Controller struct {
	// Associated wireless controller cloud ID
	Serial *string `json:"serial,omitempty"`
}

// NewInlineResponse200341Controller instantiates a new InlineResponse200341Controller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200341Controller() *InlineResponse200341Controller {
	this := InlineResponse200341Controller{}
	return &this
}

// NewInlineResponse200341ControllerWithDefaults instantiates a new InlineResponse200341Controller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200341ControllerWithDefaults() *InlineResponse200341Controller {
	this := InlineResponse200341Controller{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200341Controller) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Controller) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200341Controller) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200341Controller) SetSerial(v string) {
	o.Serial = &v
}

func (o InlineResponse200341Controller) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200341Controller struct {
	value *InlineResponse200341Controller
	isSet bool
}

func (v NullableInlineResponse200341Controller) Get() *InlineResponse200341Controller {
	return v.value
}

func (v *NullableInlineResponse200341Controller) Set(val *InlineResponse200341Controller) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200341Controller) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200341Controller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200341Controller(val *InlineResponse200341Controller) *NullableInlineResponse200341Controller {
	return &NullableInlineResponse200341Controller{value: val, isSet: true}
}

func (v NullableInlineResponse200341Controller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200341Controller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


