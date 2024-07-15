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

// InlineResponse200143Alerts Email alert settings for DHCP servers
type InlineResponse200143Alerts struct {
	Email *InlineResponse200143AlertsEmail `json:"email,omitempty"`
}

// NewInlineResponse200143Alerts instantiates a new InlineResponse200143Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200143Alerts() *InlineResponse200143Alerts {
	this := InlineResponse200143Alerts{}
	return &this
}

// NewInlineResponse200143AlertsWithDefaults instantiates a new InlineResponse200143Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200143AlertsWithDefaults() *InlineResponse200143Alerts {
	this := InlineResponse200143Alerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200143Alerts) GetEmail() InlineResponse200143AlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret InlineResponse200143AlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143Alerts) GetEmailOk() (*InlineResponse200143AlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200143Alerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given InlineResponse200143AlertsEmail and assigns it to the Email field.
func (o *InlineResponse200143Alerts) SetEmail(v InlineResponse200143AlertsEmail) {
	o.Email = &v
}

func (o InlineResponse200143Alerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200143Alerts struct {
	value *InlineResponse200143Alerts
	isSet bool
}

func (v NullableInlineResponse200143Alerts) Get() *InlineResponse200143Alerts {
	return v.value
}

func (v *NullableInlineResponse200143Alerts) Set(val *InlineResponse200143Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200143Alerts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200143Alerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200143Alerts(val *InlineResponse200143Alerts) *NullableInlineResponse200143Alerts {
	return &NullableInlineResponse200143Alerts{value: val, isSet: true}
}

func (v NullableInlineResponse200143Alerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200143Alerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


