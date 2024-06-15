/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20091Products The network devices to be updated
type InlineResponse20091Products struct {
	Switch *InlineResponse20091ProductsSwitch `json:"switch,omitempty"`
}

// NewInlineResponse20091Products instantiates a new InlineResponse20091Products object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091Products() *InlineResponse20091Products {
	this := InlineResponse20091Products{}
	return &this
}

// NewInlineResponse20091ProductsWithDefaults instantiates a new InlineResponse20091Products object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091ProductsWithDefaults() *InlineResponse20091Products {
	this := InlineResponse20091Products{}
	return &this
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *InlineResponse20091Products) GetSwitch() InlineResponse20091ProductsSwitch {
	if o == nil || isNil(o.Switch) {
		var ret InlineResponse20091ProductsSwitch
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091Products) GetSwitchOk() (*InlineResponse20091ProductsSwitch, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *InlineResponse20091Products) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given InlineResponse20091ProductsSwitch and assigns it to the Switch field.
func (o *InlineResponse20091Products) SetSwitch(v InlineResponse20091ProductsSwitch) {
	o.Switch = &v
}

func (o InlineResponse20091Products) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20091Products struct {
	value *InlineResponse20091Products
	isSet bool
}

func (v NullableInlineResponse20091Products) Get() *InlineResponse20091Products {
	return v.value
}

func (v *NullableInlineResponse20091Products) Set(val *InlineResponse20091Products) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091Products) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091Products) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091Products(val *InlineResponse20091Products) *NullableInlineResponse20091Products {
	return &NullableInlineResponse20091Products{value: val, isSet: true}
}

func (v NullableInlineResponse20091Products) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091Products) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


