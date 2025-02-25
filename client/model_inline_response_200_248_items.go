/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200248Items struct for InlineResponse200248Items
type InlineResponse200248Items struct {
	// Account ID of plans to be fetched
	AccountId *string `json:"accountId,omitempty"`
	// Rate plan name
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200248Items instantiates a new InlineResponse200248Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200248Items() *InlineResponse200248Items {
	this := InlineResponse200248Items{}
	return &this
}

// NewInlineResponse200248ItemsWithDefaults instantiates a new InlineResponse200248Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200248ItemsWithDefaults() *InlineResponse200248Items {
	this := InlineResponse200248Items{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *InlineResponse200248Items) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248Items) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
    return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *InlineResponse200248Items) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *InlineResponse200248Items) SetAccountId(v string) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200248Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200248Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200248Items) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200248Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200248Items struct {
	value *InlineResponse200248Items
	isSet bool
}

func (v NullableInlineResponse200248Items) Get() *InlineResponse200248Items {
	return v.value
}

func (v *NullableInlineResponse200248Items) Set(val *InlineResponse200248Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200248Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200248Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200248Items(val *InlineResponse200248Items) *NullableInlineResponse200248Items {
	return &NullableInlineResponse200248Items{value: val, isSet: true}
}

func (v NullableInlineResponse200248Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200248Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


