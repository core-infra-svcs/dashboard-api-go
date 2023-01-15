/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023Products The network devices to be updated
type InlineResponse20023Products struct {
	Switch *InlineResponse20023ProductsSwitch `json:"switch,omitempty"`
}

// NewInlineResponse20023Products instantiates a new InlineResponse20023Products object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023Products() *InlineResponse20023Products {
	this := InlineResponse20023Products{}
	return &this
}

// NewInlineResponse20023ProductsWithDefaults instantiates a new InlineResponse20023Products object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023ProductsWithDefaults() *InlineResponse20023Products {
	this := InlineResponse20023Products{}
	return &this
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *InlineResponse20023Products) GetSwitch() InlineResponse20023ProductsSwitch {
	if o == nil || isNil(o.Switch) {
		var ret InlineResponse20023ProductsSwitch
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023Products) GetSwitchOk() (*InlineResponse20023ProductsSwitch, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *InlineResponse20023Products) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given InlineResponse20023ProductsSwitch and assigns it to the Switch field.
func (o *InlineResponse20023Products) SetSwitch(v InlineResponse20023ProductsSwitch) {
	o.Switch = &v
}

func (o InlineResponse20023Products) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023Products struct {
	value *InlineResponse20023Products
	isSet bool
}

func (v NullableInlineResponse20023Products) Get() *InlineResponse20023Products {
	return v.value
}

func (v *NullableInlineResponse20023Products) Set(val *InlineResponse20023Products) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023Products) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023Products) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023Products(val *InlineResponse20023Products) *NullableInlineResponse20023Products {
	return &NullableInlineResponse20023Products{value: val, isSet: true}
}

func (v NullableInlineResponse20023Products) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023Products) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


