/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject221 struct for InlineObject221
type InlineObject221 struct {
	// Array of alert IDs to restore
	AlertIds []string `json:"alertIds"`
}

// NewInlineObject221 instantiates a new InlineObject221 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject221(alertIds []string) *InlineObject221 {
	this := InlineObject221{}
	this.AlertIds = alertIds
	return &this
}

// NewInlineObject221WithDefaults instantiates a new InlineObject221 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject221WithDefaults() *InlineObject221 {
	this := InlineObject221{}
	return &this
}

// GetAlertIds returns the AlertIds field value
func (o *InlineObject221) GetAlertIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertIds
}

// GetAlertIdsOk returns a tuple with the AlertIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject221) GetAlertIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AlertIds, true
}

// SetAlertIds sets field value
func (o *InlineObject221) SetAlertIds(v []string) {
	o.AlertIds = v
}

func (o InlineObject221) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alertIds"] = o.AlertIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject221 struct {
	value *InlineObject221
	isSet bool
}

func (v NullableInlineObject221) Get() *InlineObject221 {
	return v.value
}

func (v *NullableInlineObject221) Set(val *InlineObject221) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject221) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject221) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject221(val *InlineObject221) *NullableInlineObject221 {
	return &NullableInlineObject221{value: val, isSet: true}
}

func (v NullableInlineObject221) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject221) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


