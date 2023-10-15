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

// InlineResponse20037ProductsSwitch The Switch network to be updated
type InlineResponse20037ProductsSwitch struct {
	NextUpgrade *InlineResponse20037ProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewInlineResponse20037ProductsSwitch instantiates a new InlineResponse20037ProductsSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20037ProductsSwitch() *InlineResponse20037ProductsSwitch {
	this := InlineResponse20037ProductsSwitch{}
	return &this
}

// NewInlineResponse20037ProductsSwitchWithDefaults instantiates a new InlineResponse20037ProductsSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20037ProductsSwitchWithDefaults() *InlineResponse20037ProductsSwitch {
	this := InlineResponse20037ProductsSwitch{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20037ProductsSwitch) GetNextUpgrade() InlineResponse20037ProductsSwitchNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret InlineResponse20037ProductsSwitchNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20037ProductsSwitch) GetNextUpgradeOk() (*InlineResponse20037ProductsSwitchNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20037ProductsSwitch) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given InlineResponse20037ProductsSwitchNextUpgrade and assigns it to the NextUpgrade field.
func (o *InlineResponse20037ProductsSwitch) SetNextUpgrade(v InlineResponse20037ProductsSwitchNextUpgrade) {
	o.NextUpgrade = &v
}

func (o InlineResponse20037ProductsSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20037ProductsSwitch struct {
	value *InlineResponse20037ProductsSwitch
	isSet bool
}

func (v NullableInlineResponse20037ProductsSwitch) Get() *InlineResponse20037ProductsSwitch {
	return v.value
}

func (v *NullableInlineResponse20037ProductsSwitch) Set(val *InlineResponse20037ProductsSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20037ProductsSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20037ProductsSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20037ProductsSwitch(val *InlineResponse20037ProductsSwitch) *NullableInlineResponse20037ProductsSwitch {
	return &NullableInlineResponse20037ProductsSwitch{value: val, isSet: true}
}

func (v NullableInlineResponse20037ProductsSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20037ProductsSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


