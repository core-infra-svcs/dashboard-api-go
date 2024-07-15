/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20027DdnsHostnames Dynamic DNS hostnames.
type InlineResponse20027DdnsHostnames struct {
	// Active dynamic DNS hostname.
	ActiveDdnsHostname *string `json:"activeDdnsHostname,omitempty"`
	// WAN 1 dynamic DNS hostname.
	DdnsHostnameWan1 *string `json:"ddnsHostnameWan1,omitempty"`
	// WAN 2 dynamic DNS hostname.
	DdnsHostnameWan2 *string `json:"ddnsHostnameWan2,omitempty"`
}

// NewInlineResponse20027DdnsHostnames instantiates a new InlineResponse20027DdnsHostnames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027DdnsHostnames() *InlineResponse20027DdnsHostnames {
	this := InlineResponse20027DdnsHostnames{}
	return &this
}

// NewInlineResponse20027DdnsHostnamesWithDefaults instantiates a new InlineResponse20027DdnsHostnames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027DdnsHostnamesWithDefaults() *InlineResponse20027DdnsHostnames {
	this := InlineResponse20027DdnsHostnames{}
	return &this
}

// GetActiveDdnsHostname returns the ActiveDdnsHostname field value if set, zero value otherwise.
func (o *InlineResponse20027DdnsHostnames) GetActiveDdnsHostname() string {
	if o == nil || isNil(o.ActiveDdnsHostname) {
		var ret string
		return ret
	}
	return *o.ActiveDdnsHostname
}

// GetActiveDdnsHostnameOk returns a tuple with the ActiveDdnsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027DdnsHostnames) GetActiveDdnsHostnameOk() (*string, bool) {
	if o == nil || isNil(o.ActiveDdnsHostname) {
    return nil, false
	}
	return o.ActiveDdnsHostname, true
}

// HasActiveDdnsHostname returns a boolean if a field has been set.
func (o *InlineResponse20027DdnsHostnames) HasActiveDdnsHostname() bool {
	if o != nil && !isNil(o.ActiveDdnsHostname) {
		return true
	}

	return false
}

// SetActiveDdnsHostname gets a reference to the given string and assigns it to the ActiveDdnsHostname field.
func (o *InlineResponse20027DdnsHostnames) SetActiveDdnsHostname(v string) {
	o.ActiveDdnsHostname = &v
}

// GetDdnsHostnameWan1 returns the DdnsHostnameWan1 field value if set, zero value otherwise.
func (o *InlineResponse20027DdnsHostnames) GetDdnsHostnameWan1() string {
	if o == nil || isNil(o.DdnsHostnameWan1) {
		var ret string
		return ret
	}
	return *o.DdnsHostnameWan1
}

// GetDdnsHostnameWan1Ok returns a tuple with the DdnsHostnameWan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027DdnsHostnames) GetDdnsHostnameWan1Ok() (*string, bool) {
	if o == nil || isNil(o.DdnsHostnameWan1) {
    return nil, false
	}
	return o.DdnsHostnameWan1, true
}

// HasDdnsHostnameWan1 returns a boolean if a field has been set.
func (o *InlineResponse20027DdnsHostnames) HasDdnsHostnameWan1() bool {
	if o != nil && !isNil(o.DdnsHostnameWan1) {
		return true
	}

	return false
}

// SetDdnsHostnameWan1 gets a reference to the given string and assigns it to the DdnsHostnameWan1 field.
func (o *InlineResponse20027DdnsHostnames) SetDdnsHostnameWan1(v string) {
	o.DdnsHostnameWan1 = &v
}

// GetDdnsHostnameWan2 returns the DdnsHostnameWan2 field value if set, zero value otherwise.
func (o *InlineResponse20027DdnsHostnames) GetDdnsHostnameWan2() string {
	if o == nil || isNil(o.DdnsHostnameWan2) {
		var ret string
		return ret
	}
	return *o.DdnsHostnameWan2
}

// GetDdnsHostnameWan2Ok returns a tuple with the DdnsHostnameWan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027DdnsHostnames) GetDdnsHostnameWan2Ok() (*string, bool) {
	if o == nil || isNil(o.DdnsHostnameWan2) {
    return nil, false
	}
	return o.DdnsHostnameWan2, true
}

// HasDdnsHostnameWan2 returns a boolean if a field has been set.
func (o *InlineResponse20027DdnsHostnames) HasDdnsHostnameWan2() bool {
	if o != nil && !isNil(o.DdnsHostnameWan2) {
		return true
	}

	return false
}

// SetDdnsHostnameWan2 gets a reference to the given string and assigns it to the DdnsHostnameWan2 field.
func (o *InlineResponse20027DdnsHostnames) SetDdnsHostnameWan2(v string) {
	o.DdnsHostnameWan2 = &v
}

func (o InlineResponse20027DdnsHostnames) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActiveDdnsHostname) {
		toSerialize["activeDdnsHostname"] = o.ActiveDdnsHostname
	}
	if !isNil(o.DdnsHostnameWan1) {
		toSerialize["ddnsHostnameWan1"] = o.DdnsHostnameWan1
	}
	if !isNil(o.DdnsHostnameWan2) {
		toSerialize["ddnsHostnameWan2"] = o.DdnsHostnameWan2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027DdnsHostnames struct {
	value *InlineResponse20027DdnsHostnames
	isSet bool
}

func (v NullableInlineResponse20027DdnsHostnames) Get() *InlineResponse20027DdnsHostnames {
	return v.value
}

func (v *NullableInlineResponse20027DdnsHostnames) Set(val *InlineResponse20027DdnsHostnames) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027DdnsHostnames) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027DdnsHostnames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027DdnsHostnames(val *InlineResponse20027DdnsHostnames) *NullableInlineResponse20027DdnsHostnames {
	return &NullableInlineResponse20027DdnsHostnames{value: val, isSet: true}
}

func (v NullableInlineResponse20027DdnsHostnames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027DdnsHostnames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


