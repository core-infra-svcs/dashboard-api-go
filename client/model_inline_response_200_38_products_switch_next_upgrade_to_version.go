/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20038ProductsSwitchNextUpgradeToVersion Details of the version the device will upgrade to
type InlineResponse20038ProductsSwitchNextUpgradeToVersion struct {
	// Id of the Version being upgraded to
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
}

// NewInlineResponse20038ProductsSwitchNextUpgradeToVersion instantiates a new InlineResponse20038ProductsSwitchNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20038ProductsSwitchNextUpgradeToVersion() *InlineResponse20038ProductsSwitchNextUpgradeToVersion {
	this := InlineResponse20038ProductsSwitchNextUpgradeToVersion{}
	return &this
}

// NewInlineResponse20038ProductsSwitchNextUpgradeToVersionWithDefaults instantiates a new InlineResponse20038ProductsSwitchNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20038ProductsSwitchNextUpgradeToVersionWithDefaults() *InlineResponse20038ProductsSwitchNextUpgradeToVersion {
	this := InlineResponse20038ProductsSwitchNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20038ProductsSwitchNextUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

func (o InlineResponse20038ProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion struct {
	value *InlineResponse20038ProductsSwitchNextUpgradeToVersion
	isSet bool
}

func (v NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion) Get() *InlineResponse20038ProductsSwitchNextUpgradeToVersion {
	return v.value
}

func (v *NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion) Set(val *InlineResponse20038ProductsSwitchNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20038ProductsSwitchNextUpgradeToVersion(val *InlineResponse20038ProductsSwitchNextUpgradeToVersion) *NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion {
	return &NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20038ProductsSwitchNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

