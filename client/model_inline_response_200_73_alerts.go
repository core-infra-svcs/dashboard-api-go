/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20073Alerts Email alert settings for DHCP servers
type InlineResponse20073Alerts struct {
	Email *InlineResponse20073AlertsEmail `json:"email,omitempty"`
}

// NewInlineResponse20073Alerts instantiates a new InlineResponse20073Alerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20073Alerts() *InlineResponse20073Alerts {
	this := InlineResponse20073Alerts{}
	return &this
}

// NewInlineResponse20073AlertsWithDefaults instantiates a new InlineResponse20073Alerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20073AlertsWithDefaults() *InlineResponse20073Alerts {
	this := InlineResponse20073Alerts{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse20073Alerts) GetEmail() InlineResponse20073AlertsEmail {
	if o == nil || isNil(o.Email) {
		var ret InlineResponse20073AlertsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073Alerts) GetEmailOk() (*InlineResponse20073AlertsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse20073Alerts) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given InlineResponse20073AlertsEmail and assigns it to the Email field.
func (o *InlineResponse20073Alerts) SetEmail(v InlineResponse20073AlertsEmail) {
	o.Email = &v
}

func (o InlineResponse20073Alerts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20073Alerts struct {
	value *InlineResponse20073Alerts
	isSet bool
}

func (v NullableInlineResponse20073Alerts) Get() *InlineResponse20073Alerts {
	return v.value
}

func (v *NullableInlineResponse20073Alerts) Set(val *InlineResponse20073Alerts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20073Alerts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20073Alerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20073Alerts(val *InlineResponse20073Alerts) *NullableInlineResponse20073Alerts {
	return &NullableInlineResponse20073Alerts{value: val, isSet: true}
}

func (v NullableInlineResponse20073Alerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20073Alerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


