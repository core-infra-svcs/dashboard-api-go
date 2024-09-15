/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	Interfaces DevicesSerialApplianceUplinksSettingsInterfaces `json:"interfaces"`
}

// NewInlineObject5 instantiates a new InlineObject5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject5(interfaces DevicesSerialApplianceUplinksSettingsInterfaces) *InlineObject5 {
	this := InlineObject5{}
	this.Interfaces = interfaces
	return &this
}

// NewInlineObject5WithDefaults instantiates a new InlineObject5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject5WithDefaults() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// GetInterfaces returns the Interfaces field value
func (o *InlineObject5) GetInterfaces() DevicesSerialApplianceUplinksSettingsInterfaces {
	if o == nil {
		var ret DevicesSerialApplianceUplinksSettingsInterfaces
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetInterfacesOk() (*DevicesSerialApplianceUplinksSettingsInterfaces, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *InlineObject5) SetInterfaces(v DevicesSerialApplianceUplinksSettingsInterfaces) {
	o.Interfaces = v
}

func (o InlineObject5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject5 struct {
	value *InlineObject5
	isSet bool
}

func (v NullableInlineObject5) Get() *InlineObject5 {
	return v.value
}

func (v *NullableInlineObject5) Set(val *InlineObject5) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject5) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject5(val *InlineObject5) *NullableInlineObject5 {
	return &NullableInlineObject5{value: val, isSet: true}
}

func (v NullableInlineObject5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


