/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20082Alerts Email alert settings for DHCP servers
type InlineResponse20082Alerts struct {
	Email *InlineResponse20082AlertsEmail `json:"email,omitempty"`
}

// NewInlineResponse20082Alerts instantiates a new InlineResponse20082Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20082Alerts() *InlineResponse20082Alerts {
	this := InlineResponse20082Alerts{}
	return &this
}

// NewInlineResponse20082AlertsWithDefaults instantiates a new InlineResponse20082Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20082AlertsWithDefaults() *InlineResponse20082Alerts {
	this := InlineResponse20082Alerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse20082Alerts) GetEmail() InlineResponse20082AlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret InlineResponse20082AlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082Alerts) GetEmailOk() (*InlineResponse20082AlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse20082Alerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given InlineResponse20082AlertsEmail and assigns it to the Email field.
func (o *InlineResponse20082Alerts) SetEmail(v InlineResponse20082AlertsEmail) {
	o.Email = &v
}

func (o InlineResponse20082Alerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20082Alerts struct {
	value *InlineResponse20082Alerts
	isSet bool
}

func (v NullableInlineResponse20082Alerts) Get() *InlineResponse20082Alerts {
	return v.value
}

func (v *NullableInlineResponse20082Alerts) Set(val *InlineResponse20082Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20082Alerts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20082Alerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20082Alerts(val *InlineResponse20082Alerts) *NullableInlineResponse20082Alerts {
	return &NullableInlineResponse20082Alerts{value: val, isSet: true}
}

func (v NullableInlineResponse20082Alerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20082Alerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


