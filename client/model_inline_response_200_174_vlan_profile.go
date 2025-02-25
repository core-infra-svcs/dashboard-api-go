/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200174VlanProfile The VLAN Profile
type InlineResponse200174VlanProfile struct {
	// IName of the VLAN Profile
	Iname *string `json:"iname,omitempty"`
	// Name of the VLAN Profile
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200174VlanProfile instantiates a new InlineResponse200174VlanProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200174VlanProfile() *InlineResponse200174VlanProfile {
	this := InlineResponse200174VlanProfile{}
	return &this
}

// NewInlineResponse200174VlanProfileWithDefaults instantiates a new InlineResponse200174VlanProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200174VlanProfileWithDefaults() *InlineResponse200174VlanProfile {
	this := InlineResponse200174VlanProfile{}
	return &this
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *InlineResponse200174VlanProfile) GetIname() string {
	if o == nil || isNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200174VlanProfile) GetInameOk() (*string, bool) {
	if o == nil || isNil(o.Iname) {
    return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *InlineResponse200174VlanProfile) HasIname() bool {
	if o != nil && !isNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *InlineResponse200174VlanProfile) SetIname(v string) {
	o.Iname = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200174VlanProfile) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200174VlanProfile) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200174VlanProfile) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200174VlanProfile) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200174VlanProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200174VlanProfile struct {
	value *InlineResponse200174VlanProfile
	isSet bool
}

func (v NullableInlineResponse200174VlanProfile) Get() *InlineResponse200174VlanProfile {
	return v.value
}

func (v *NullableInlineResponse200174VlanProfile) Set(val *InlineResponse200174VlanProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200174VlanProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200174VlanProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200174VlanProfile(val *InlineResponse200174VlanProfile) *NullableInlineResponse200174VlanProfile {
	return &NullableInlineResponse200174VlanProfile{value: val, isSet: true}
}

func (v NullableInlineResponse200174VlanProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200174VlanProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


