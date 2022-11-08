/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20078Components Components
type InlineResponse20078Components struct {
	// Power Supplies
	PowerSupplies []string `json:"powerSupplies,omitempty"`
}

// NewInlineResponse20078Components instantiates a new InlineResponse20078Components object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20078Components() *InlineResponse20078Components {
	this := InlineResponse20078Components{}
	return &this
}

// NewInlineResponse20078ComponentsWithDefaults instantiates a new InlineResponse20078Components object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20078ComponentsWithDefaults() *InlineResponse20078Components {
	this := InlineResponse20078Components{}
	return &this
}

// GetPowerSupplies returns the PowerSupplies field value if set, zero value otherwise.
func (o *InlineResponse20078Components) GetPowerSupplies() []string {
	if o == nil || isNil(o.PowerSupplies) {
		var ret []string
		return ret
	}
	return o.PowerSupplies
}

// GetPowerSuppliesOk returns a tuple with the PowerSupplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20078Components) GetPowerSuppliesOk() ([]string, bool) {
	if o == nil || isNil(o.PowerSupplies) {
    return nil, false
	}
	return o.PowerSupplies, true
}

// HasPowerSupplies returns a boolean if a field has been set.
func (o *InlineResponse20078Components) HasPowerSupplies() bool {
	if o != nil && !isNil(o.PowerSupplies) {
		return true
	}

	return false
}

// SetPowerSupplies gets a reference to the given []string and assigns it to the PowerSupplies field.
func (o *InlineResponse20078Components) SetPowerSupplies(v []string) {
	o.PowerSupplies = v
}

func (o InlineResponse20078Components) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PowerSupplies) {
		toSerialize["powerSupplies"] = o.PowerSupplies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20078Components struct {
	value *InlineResponse20078Components
	isSet bool
}

func (v NullableInlineResponse20078Components) Get() *InlineResponse20078Components {
	return v.value
}

func (v *NullableInlineResponse20078Components) Set(val *InlineResponse20078Components) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20078Components) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20078Components) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20078Components(val *InlineResponse20078Components) *NullableInlineResponse20078Components {
	return &NullableInlineResponse20078Components{value: val, isSet: true}
}

func (v NullableInlineResponse20078Components) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20078Components) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

