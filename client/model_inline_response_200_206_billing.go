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

// InlineResponse200206Billing Details associated with billing splash
type InlineResponse200206Billing struct {
	FreeAccess *InlineResponse200206BillingFreeAccess `json:"freeAccess,omitempty"`
	// Whether or not billing uses the fast login prepaid access option.
	PrepaidAccessFastLoginEnabled *bool `json:"prepaidAccessFastLoginEnabled,omitempty"`
	// The email address that reeceives replies from clients
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty"`
}

// NewInlineResponse200206Billing instantiates a new InlineResponse200206Billing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200206Billing() *InlineResponse200206Billing {
	this := InlineResponse200206Billing{}
	return &this
}

// NewInlineResponse200206BillingWithDefaults instantiates a new InlineResponse200206Billing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200206BillingWithDefaults() *InlineResponse200206Billing {
	this := InlineResponse200206Billing{}
	return &this
}

// GetFreeAccess returns the FreeAccess field value if set, zero value otherwise.
func (o *InlineResponse200206Billing) GetFreeAccess() InlineResponse200206BillingFreeAccess {
	if o == nil || isNil(o.FreeAccess) {
		var ret InlineResponse200206BillingFreeAccess
		return ret
	}
	return *o.FreeAccess
}

// GetFreeAccessOk returns a tuple with the FreeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200206Billing) GetFreeAccessOk() (*InlineResponse200206BillingFreeAccess, bool) {
	if o == nil || isNil(o.FreeAccess) {
    return nil, false
	}
	return o.FreeAccess, true
}

// HasFreeAccess returns a boolean if a field has been set.
func (o *InlineResponse200206Billing) HasFreeAccess() bool {
	if o != nil && !isNil(o.FreeAccess) {
		return true
	}

	return false
}

// SetFreeAccess gets a reference to the given InlineResponse200206BillingFreeAccess and assigns it to the FreeAccess field.
func (o *InlineResponse200206Billing) SetFreeAccess(v InlineResponse200206BillingFreeAccess) {
	o.FreeAccess = &v
}

// GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field value if set, zero value otherwise.
func (o *InlineResponse200206Billing) GetPrepaidAccessFastLoginEnabled() bool {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.PrepaidAccessFastLoginEnabled
}

// GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200206Billing) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
    return nil, false
	}
	return o.PrepaidAccessFastLoginEnabled, true
}

// HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.
func (o *InlineResponse200206Billing) HasPrepaidAccessFastLoginEnabled() bool {
	if o != nil && !isNil(o.PrepaidAccessFastLoginEnabled) {
		return true
	}

	return false
}

// SetPrepaidAccessFastLoginEnabled gets a reference to the given bool and assigns it to the PrepaidAccessFastLoginEnabled field.
func (o *InlineResponse200206Billing) SetPrepaidAccessFastLoginEnabled(v bool) {
	o.PrepaidAccessFastLoginEnabled = &v
}

// GetReplyToEmailAddress returns the ReplyToEmailAddress field value if set, zero value otherwise.
func (o *InlineResponse200206Billing) GetReplyToEmailAddress() string {
	if o == nil || isNil(o.ReplyToEmailAddress) {
		var ret string
		return ret
	}
	return *o.ReplyToEmailAddress
}

// GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200206Billing) GetReplyToEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.ReplyToEmailAddress) {
    return nil, false
	}
	return o.ReplyToEmailAddress, true
}

// HasReplyToEmailAddress returns a boolean if a field has been set.
func (o *InlineResponse200206Billing) HasReplyToEmailAddress() bool {
	if o != nil && !isNil(o.ReplyToEmailAddress) {
		return true
	}

	return false
}

// SetReplyToEmailAddress gets a reference to the given string and assigns it to the ReplyToEmailAddress field.
func (o *InlineResponse200206Billing) SetReplyToEmailAddress(v string) {
	o.ReplyToEmailAddress = &v
}

func (o InlineResponse200206Billing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FreeAccess) {
		toSerialize["freeAccess"] = o.FreeAccess
	}
	if !isNil(o.PrepaidAccessFastLoginEnabled) {
		toSerialize["prepaidAccessFastLoginEnabled"] = o.PrepaidAccessFastLoginEnabled
	}
	if !isNil(o.ReplyToEmailAddress) {
		toSerialize["replyToEmailAddress"] = o.ReplyToEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200206Billing struct {
	value *InlineResponse200206Billing
	isSet bool
}

func (v NullableInlineResponse200206Billing) Get() *InlineResponse200206Billing {
	return v.value
}

func (v *NullableInlineResponse200206Billing) Set(val *InlineResponse200206Billing) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200206Billing) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200206Billing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200206Billing(val *InlineResponse200206Billing) *NullableInlineResponse200206Billing {
	return &NullableInlineResponse200206Billing{value: val, isSet: true}
}

func (v NullableInlineResponse200206Billing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200206Billing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


