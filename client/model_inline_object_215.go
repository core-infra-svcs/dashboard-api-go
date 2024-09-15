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

// InlineObject215 struct for InlineObject215
type InlineObject215 struct {
	// Array of alert IDs to dismiss
	AlertIds []string `json:"alertIds"`
}

// NewInlineObject215 instantiates a new InlineObject215 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject215(alertIds []string) *InlineObject215 {
	this := InlineObject215{}
	this.AlertIds = alertIds
	return &this
}

// NewInlineObject215WithDefaults instantiates a new InlineObject215 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject215WithDefaults() *InlineObject215 {
	this := InlineObject215{}
	return &this
}

// GetAlertIds returns the AlertIds field value
func (o *InlineObject215) GetAlertIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertIds
}

// GetAlertIdsOk returns a tuple with the AlertIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject215) GetAlertIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AlertIds, true
}

// SetAlertIds sets field value
func (o *InlineObject215) SetAlertIds(v []string) {
	o.AlertIds = v
}

func (o InlineObject215) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alertIds"] = o.AlertIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject215 struct {
	value *InlineObject215
	isSet bool
}

func (v NullableInlineObject215) Get() *InlineObject215 {
	return v.value
}

func (v *NullableInlineObject215) Set(val *InlineObject215) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject215) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject215) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject215(val *InlineObject215) *NullableInlineObject215 {
	return &NullableInlineObject215{value: val, isSet: true}
}

func (v NullableInlineObject215) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject215) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


