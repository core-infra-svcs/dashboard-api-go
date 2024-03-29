/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20053ProductsSwitch The Switch network to be updated
type InlineResponse20053ProductsSwitch struct {
	NextUpgrade *InlineResponse20053ProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewInlineResponse20053ProductsSwitch instantiates a new InlineResponse20053ProductsSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20053ProductsSwitch() *InlineResponse20053ProductsSwitch {
	this := InlineResponse20053ProductsSwitch{}
	return &this
}

// NewInlineResponse20053ProductsSwitchWithDefaults instantiates a new InlineResponse20053ProductsSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20053ProductsSwitchWithDefaults() *InlineResponse20053ProductsSwitch {
	this := InlineResponse20053ProductsSwitch{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20053ProductsSwitch) GetNextUpgrade() InlineResponse20053ProductsSwitchNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret InlineResponse20053ProductsSwitchNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053ProductsSwitch) GetNextUpgradeOk() (*InlineResponse20053ProductsSwitchNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20053ProductsSwitch) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given InlineResponse20053ProductsSwitchNextUpgrade and assigns it to the NextUpgrade field.
func (o *InlineResponse20053ProductsSwitch) SetNextUpgrade(v InlineResponse20053ProductsSwitchNextUpgrade) {
	o.NextUpgrade = &v
}

func (o InlineResponse20053ProductsSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20053ProductsSwitch struct {
	value *InlineResponse20053ProductsSwitch
	isSet bool
}

func (v NullableInlineResponse20053ProductsSwitch) Get() *InlineResponse20053ProductsSwitch {
	return v.value
}

func (v *NullableInlineResponse20053ProductsSwitch) Set(val *InlineResponse20053ProductsSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20053ProductsSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20053ProductsSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20053ProductsSwitch(val *InlineResponse20053ProductsSwitch) *NullableInlineResponse20053ProductsSwitch {
	return &NullableInlineResponse20053ProductsSwitch{value: val, isSet: true}
}

func (v NullableInlineResponse20053ProductsSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20053ProductsSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


