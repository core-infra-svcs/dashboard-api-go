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

// InlineResponse200342 struct for InlineResponse200342
type InlineResponse200342 struct {
	// The type ID of Meraki alert
	AlertTypeId *string `json:"alertTypeId,omitempty"`
	// The type of Meraki alert
	AlertType *string `json:"alertType,omitempty"`
	Example *InlineResponse200342Example `json:"example,omitempty"`
}

// NewInlineResponse200342 instantiates a new InlineResponse200342 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200342() *InlineResponse200342 {
	this := InlineResponse200342{}
	return &this
}

// NewInlineResponse200342WithDefaults instantiates a new InlineResponse200342 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200342WithDefaults() *InlineResponse200342 {
	this := InlineResponse200342{}
	return &this
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *InlineResponse200342) GetAlertTypeId() string {
	if o == nil || isNil(o.AlertTypeId) {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200342) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || isNil(o.AlertTypeId) {
    return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *InlineResponse200342) HasAlertTypeId() bool {
	if o != nil && !isNil(o.AlertTypeId) {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *InlineResponse200342) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *InlineResponse200342) GetAlertType() string {
	if o == nil || isNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200342) GetAlertTypeOk() (*string, bool) {
	if o == nil || isNil(o.AlertType) {
    return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *InlineResponse200342) HasAlertType() bool {
	if o != nil && !isNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *InlineResponse200342) SetAlertType(v string) {
	o.AlertType = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *InlineResponse200342) GetExample() InlineResponse200342Example {
	if o == nil || isNil(o.Example) {
		var ret InlineResponse200342Example
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200342) GetExampleOk() (*InlineResponse200342Example, bool) {
	if o == nil || isNil(o.Example) {
    return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *InlineResponse200342) HasExample() bool {
	if o != nil && !isNil(o.Example) {
		return true
	}

	return false
}

// SetExample gets a reference to the given InlineResponse200342Example and assigns it to the Example field.
func (o *InlineResponse200342) SetExample(v InlineResponse200342Example) {
	o.Example = &v
}

func (o InlineResponse200342) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlertTypeId) {
		toSerialize["alertTypeId"] = o.AlertTypeId
	}
	if !isNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !isNil(o.Example) {
		toSerialize["example"] = o.Example
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200342 struct {
	value *InlineResponse200342
	isSet bool
}

func (v NullableInlineResponse200342) Get() *InlineResponse200342 {
	return v.value
}

func (v *NullableInlineResponse200342) Set(val *InlineResponse200342) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200342) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200342) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200342(val *InlineResponse200342) *NullableInlineResponse200342 {
	return &NullableInlineResponse200342{value: val, isSet: true}
}

func (v NullableInlineResponse200342) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200342) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


