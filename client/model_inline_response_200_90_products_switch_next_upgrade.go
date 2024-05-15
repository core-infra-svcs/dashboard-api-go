/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20090ProductsSwitchNextUpgrade Details of the next firmware upgrade
type InlineResponse20090ProductsSwitchNextUpgrade struct {
	ToVersion *InlineResponse20090ProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewInlineResponse20090ProductsSwitchNextUpgrade instantiates a new InlineResponse20090ProductsSwitchNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090ProductsSwitchNextUpgrade() *InlineResponse20090ProductsSwitchNextUpgrade {
	this := InlineResponse20090ProductsSwitchNextUpgrade{}
	return &this
}

// NewInlineResponse20090ProductsSwitchNextUpgradeWithDefaults instantiates a new InlineResponse20090ProductsSwitchNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090ProductsSwitchNextUpgradeWithDefaults() *InlineResponse20090ProductsSwitchNextUpgrade {
	this := InlineResponse20090ProductsSwitchNextUpgrade{}
	return &this
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse20090ProductsSwitchNextUpgrade) GetToVersion() InlineResponse20090ProductsSwitchNextUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret InlineResponse20090ProductsSwitchNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090ProductsSwitchNextUpgrade) GetToVersionOk() (*InlineResponse20090ProductsSwitchNextUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse20090ProductsSwitchNextUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given InlineResponse20090ProductsSwitchNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *InlineResponse20090ProductsSwitchNextUpgrade) SetToVersion(v InlineResponse20090ProductsSwitchNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o InlineResponse20090ProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20090ProductsSwitchNextUpgrade struct {
	value *InlineResponse20090ProductsSwitchNextUpgrade
	isSet bool
}

func (v NullableInlineResponse20090ProductsSwitchNextUpgrade) Get() *InlineResponse20090ProductsSwitchNextUpgrade {
	return v.value
}

func (v *NullableInlineResponse20090ProductsSwitchNextUpgrade) Set(val *InlineResponse20090ProductsSwitchNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090ProductsSwitchNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090ProductsSwitchNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090ProductsSwitchNextUpgrade(val *InlineResponse20090ProductsSwitchNextUpgrade) *NullableInlineResponse20090ProductsSwitchNextUpgrade {
	return &NullableInlineResponse20090ProductsSwitchNextUpgrade{value: val, isSet: true}
}

func (v NullableInlineResponse20090ProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090ProductsSwitchNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


