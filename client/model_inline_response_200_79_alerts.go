/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20079Alerts Email alert settings for DHCP servers
type InlineResponse20079Alerts struct {
	Email *InlineResponse20079AlertsEmail `json:"email,omitempty"`
}

// NewInlineResponse20079Alerts instantiates a new InlineResponse20079Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079Alerts() *InlineResponse20079Alerts {
	this := InlineResponse20079Alerts{}
	return &this
}

// NewInlineResponse20079AlertsWithDefaults instantiates a new InlineResponse20079Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079AlertsWithDefaults() *InlineResponse20079Alerts {
	this := InlineResponse20079Alerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse20079Alerts) GetEmail() InlineResponse20079AlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret InlineResponse20079AlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079Alerts) GetEmailOk() (*InlineResponse20079AlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse20079Alerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given InlineResponse20079AlertsEmail and assigns it to the Email field.
func (o *InlineResponse20079Alerts) SetEmail(v InlineResponse20079AlertsEmail) {
	o.Email = &v
}

func (o InlineResponse20079Alerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079Alerts struct {
	value *InlineResponse20079Alerts
	isSet bool
}

func (v NullableInlineResponse20079Alerts) Get() *InlineResponse20079Alerts {
	return v.value
}

func (v *NullableInlineResponse20079Alerts) Set(val *InlineResponse20079Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079Alerts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079Alerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079Alerts(val *InlineResponse20079Alerts) *NullableInlineResponse20079Alerts {
	return &NullableInlineResponse20079Alerts{value: val, isSet: true}
}

func (v NullableInlineResponse20079Alerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079Alerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

