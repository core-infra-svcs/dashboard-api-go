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

// InlineResponse200348Items struct for InlineResponse200348Items
type InlineResponse200348Items struct {
	// The cloud ID of the wireless LAN controller
	Serial *string `json:"serial,omitempty"`
	// Layer 2 interfaces belongs to the wireless LAN controller
	Interfaces []InlineResponse200348Interfaces `json:"interfaces,omitempty"`
}

// NewInlineResponse200348Items instantiates a new InlineResponse200348Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200348Items() *InlineResponse200348Items {
	this := InlineResponse200348Items{}
	return &this
}

// NewInlineResponse200348ItemsWithDefaults instantiates a new InlineResponse200348Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200348ItemsWithDefaults() *InlineResponse200348Items {
	this := InlineResponse200348Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200348Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200348Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200348Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200348Items) SetSerial(v string) {
	o.Serial = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *InlineResponse200348Items) GetInterfaces() []InlineResponse200348Interfaces {
	if o == nil || isNil(o.Interfaces) {
		var ret []InlineResponse200348Interfaces
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200348Items) GetInterfacesOk() ([]InlineResponse200348Interfaces, bool) {
	if o == nil || isNil(o.Interfaces) {
    return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *InlineResponse200348Items) HasInterfaces() bool {
	if o != nil && !isNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []InlineResponse200348Interfaces and assigns it to the Interfaces field.
func (o *InlineResponse200348Items) SetInterfaces(v []InlineResponse200348Interfaces) {
	o.Interfaces = v
}

func (o InlineResponse200348Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200348Items struct {
	value *InlineResponse200348Items
	isSet bool
}

func (v NullableInlineResponse200348Items) Get() *InlineResponse200348Items {
	return v.value
}

func (v *NullableInlineResponse200348Items) Set(val *InlineResponse200348Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200348Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200348Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200348Items(val *InlineResponse200348Items) *NullableInlineResponse200348Items {
	return &NullableInlineResponse200348Items{value: val, isSet: true}
}

func (v NullableInlineResponse200348Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200348Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


