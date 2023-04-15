/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject41 struct for InlineObject41
type InlineObject41 struct {
	// A static IPv6 prefix
	Prefix string `json:"prefix"`
	Origin NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 `json:"origin"`
	// A name or description for the prefix
	Description *string `json:"description,omitempty"`
}

// NewInlineObject41 instantiates a new InlineObject41 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject41(prefix string, origin NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) *InlineObject41 {
	this := InlineObject41{}
	this.Prefix = prefix
	this.Origin = origin
	return &this
}

// NewInlineObject41WithDefaults instantiates a new InlineObject41 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject41WithDefaults() *InlineObject41 {
	this := InlineObject41{}
	return &this
}

// GetPrefix returns the Prefix field value
func (o *InlineObject41) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetPrefixOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *InlineObject41) SetPrefix(v string) {
	o.Prefix = v
}

// GetOrigin returns the Origin field value
func (o *InlineObject41) GetOrigin() NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 {
	if o == nil {
		var ret NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetOriginOk() (*NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *InlineObject41) SetOrigin(v NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) {
	o.Origin = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject41) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject41) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject41) SetDescription(v string) {
	o.Description = &v
}

func (o InlineObject41) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if true {
		toSerialize["origin"] = o.Origin
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject41 struct {
	value *InlineObject41
	isSet bool
}

func (v NullableInlineObject41) Get() *InlineObject41 {
	return v.value
}

func (v *NullableInlineObject41) Set(val *InlineObject41) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject41) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject41) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject41(val *InlineObject41) *NullableInlineObject41 {
	return &NullableInlineObject41{value: val, isSet: true}
}

func (v NullableInlineObject41) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject41) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


