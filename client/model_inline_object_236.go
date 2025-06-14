/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject236 struct for InlineObject236
type InlineObject236 struct {
	// Array of alert IDs to dismiss
	AlertIds []string `json:"alertIds"`
}

// NewInlineObject236 instantiates a new InlineObject236 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject236(alertIds []string) *InlineObject236 {
	this := InlineObject236{}
	this.AlertIds = alertIds
	return &this
}

// NewInlineObject236WithDefaults instantiates a new InlineObject236 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject236WithDefaults() *InlineObject236 {
	this := InlineObject236{}
	return &this
}

// GetAlertIds returns the AlertIds field value
func (o *InlineObject236) GetAlertIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertIds
}

// GetAlertIdsOk returns a tuple with the AlertIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject236) GetAlertIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AlertIds, true
}

// SetAlertIds sets field value
func (o *InlineObject236) SetAlertIds(v []string) {
	o.AlertIds = v
}

func (o InlineObject236) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alertIds"] = o.AlertIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject236 struct {
	value *InlineObject236
	isSet bool
}

func (v NullableInlineObject236) Get() *InlineObject236 {
	return v.value
}

func (v *NullableInlineObject236) Set(val *InlineObject236) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject236) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject236) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject236(val *InlineObject236) *NullableInlineObject236 {
	return &NullableInlineObject236{value: val, isSet: true}
}

func (v NullableInlineObject236) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject236) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


