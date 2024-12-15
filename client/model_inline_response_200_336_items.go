/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200336Items struct for InlineResponse200336Items
type InlineResponse200336Items struct {
	// The cloud ID of the wireless LAN controller
	Serial *string `json:"serial,omitempty"`
	// layer 3 interfaces belongs to the wireless LAN controller
	Interfaces []InlineResponse200336Interfaces `json:"interfaces,omitempty"`
}

// NewInlineResponse200336Items instantiates a new InlineResponse200336Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200336Items() *InlineResponse200336Items {
	this := InlineResponse200336Items{}
	return &this
}

// NewInlineResponse200336ItemsWithDefaults instantiates a new InlineResponse200336Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200336ItemsWithDefaults() *InlineResponse200336Items {
	this := InlineResponse200336Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200336Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200336Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200336Items) SetSerial(v string) {
	o.Serial = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *InlineResponse200336Items) GetInterfaces() []InlineResponse200336Interfaces {
	if o == nil || isNil(o.Interfaces) {
		var ret []InlineResponse200336Interfaces
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336Items) GetInterfacesOk() ([]InlineResponse200336Interfaces, bool) {
	if o == nil || isNil(o.Interfaces) {
    return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *InlineResponse200336Items) HasInterfaces() bool {
	if o != nil && !isNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []InlineResponse200336Interfaces and assigns it to the Interfaces field.
func (o *InlineResponse200336Items) SetInterfaces(v []InlineResponse200336Interfaces) {
	o.Interfaces = v
}

func (o InlineResponse200336Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200336Items struct {
	value *InlineResponse200336Items
	isSet bool
}

func (v NullableInlineResponse200336Items) Get() *InlineResponse200336Items {
	return v.value
}

func (v *NullableInlineResponse200336Items) Set(val *InlineResponse200336Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200336Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200336Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200336Items(val *InlineResponse200336Items) *NullableInlineResponse200336Items {
	return &NullableInlineResponse200336Items{value: val, isSet: true}
}

func (v NullableInlineResponse200336Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200336Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


