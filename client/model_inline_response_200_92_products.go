/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20092Products The network devices to be updated
type InlineResponse20092Products struct {
	Switch *InlineResponse20092ProductsSwitch `json:"switch,omitempty"`
}

// NewInlineResponse20092Products instantiates a new InlineResponse20092Products object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092Products() *InlineResponse20092Products {
	this := InlineResponse20092Products{}
	return &this
}

// NewInlineResponse20092ProductsWithDefaults instantiates a new InlineResponse20092Products object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092ProductsWithDefaults() *InlineResponse20092Products {
	this := InlineResponse20092Products{}
	return &this
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *InlineResponse20092Products) GetSwitch() InlineResponse20092ProductsSwitch {
	if o == nil || isNil(o.Switch) {
		var ret InlineResponse20092ProductsSwitch
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092Products) GetSwitchOk() (*InlineResponse20092ProductsSwitch, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *InlineResponse20092Products) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given InlineResponse20092ProductsSwitch and assigns it to the Switch field.
func (o *InlineResponse20092Products) SetSwitch(v InlineResponse20092ProductsSwitch) {
	o.Switch = &v
}

func (o InlineResponse20092Products) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092Products struct {
	value *InlineResponse20092Products
	isSet bool
}

func (v NullableInlineResponse20092Products) Get() *InlineResponse20092Products {
	return v.value
}

func (v *NullableInlineResponse20092Products) Set(val *InlineResponse20092Products) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092Products) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092Products) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092Products(val *InlineResponse20092Products) *NullableInlineResponse20092Products {
	return &NullableInlineResponse20092Products{value: val, isSet: true}
}

func (v NullableInlineResponse20092Products) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092Products) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


