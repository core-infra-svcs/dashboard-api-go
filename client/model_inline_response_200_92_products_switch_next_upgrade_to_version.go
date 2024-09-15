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

// InlineResponse20092ProductsSwitchNextUpgradeToVersion Details of the version the device will upgrade to
type InlineResponse20092ProductsSwitchNextUpgradeToVersion struct {
	// Id of the Version being upgraded to
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
}

// NewInlineResponse20092ProductsSwitchNextUpgradeToVersion instantiates a new InlineResponse20092ProductsSwitchNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092ProductsSwitchNextUpgradeToVersion() *InlineResponse20092ProductsSwitchNextUpgradeToVersion {
	this := InlineResponse20092ProductsSwitchNextUpgradeToVersion{}
	return &this
}

// NewInlineResponse20092ProductsSwitchNextUpgradeToVersionWithDefaults instantiates a new InlineResponse20092ProductsSwitchNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092ProductsSwitchNextUpgradeToVersionWithDefaults() *InlineResponse20092ProductsSwitchNextUpgradeToVersion {
	this := InlineResponse20092ProductsSwitchNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20092ProductsSwitchNextUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

func (o InlineResponse20092ProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion struct {
	value *InlineResponse20092ProductsSwitchNextUpgradeToVersion
	isSet bool
}

func (v NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion) Get() *InlineResponse20092ProductsSwitchNextUpgradeToVersion {
	return v.value
}

func (v *NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion) Set(val *InlineResponse20092ProductsSwitchNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092ProductsSwitchNextUpgradeToVersion(val *InlineResponse20092ProductsSwitchNextUpgradeToVersion) *NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion {
	return &NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092ProductsSwitchNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


