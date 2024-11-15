/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject42 struct for InlineObject42
type InlineObject42 struct {
	// A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are \"blocked\" (no remote IPs can access the service), \"restricted\" (only allowed IPs can access the service), and \"unrestriced\" (any remote IP can access the service). This field is required
	Access string `json:"access"`
	// An array of allowed IPs that can access the service. This field is required if \"access\" is set to \"restricted\". Otherwise this field is ignored
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewInlineObject42 instantiates a new InlineObject42 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject42(access string) *InlineObject42 {
	this := InlineObject42{}
	this.Access = access
	return &this
}

// NewInlineObject42WithDefaults instantiates a new InlineObject42 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject42WithDefaults() *InlineObject42 {
	this := InlineObject42{}
	return &this
}

// GetAccess returns the Access field value
func (o *InlineObject42) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *InlineObject42) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *InlineObject42) SetAccess(v string) {
	o.Access = v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *InlineObject42) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject42) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *InlineObject42) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *InlineObject42) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o InlineObject42) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject42 struct {
	value *InlineObject42
	isSet bool
}

func (v NullableInlineObject42) Get() *InlineObject42 {
	return v.value
}

func (v *NullableInlineObject42) Set(val *InlineObject42) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject42) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject42) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject42(val *InlineObject42) *NullableInlineObject42 {
	return &NullableInlineObject42{value: val, isSet: true}
}

func (v NullableInlineObject42) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject42) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


