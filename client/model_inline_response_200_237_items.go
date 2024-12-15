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

// InlineResponse200237Items struct for InlineResponse200237Items
type InlineResponse200237Items struct {
	// Account ID of plans to be fetched
	AccountId *string `json:"accountId,omitempty"`
	// Rate plan name
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200237Items instantiates a new InlineResponse200237Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200237Items() *InlineResponse200237Items {
	this := InlineResponse200237Items{}
	return &this
}

// NewInlineResponse200237ItemsWithDefaults instantiates a new InlineResponse200237Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200237ItemsWithDefaults() *InlineResponse200237Items {
	this := InlineResponse200237Items{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *InlineResponse200237Items) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237Items) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
    return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *InlineResponse200237Items) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *InlineResponse200237Items) SetAccountId(v string) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200237Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200237Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200237Items) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200237Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200237Items struct {
	value *InlineResponse200237Items
	isSet bool
}

func (v NullableInlineResponse200237Items) Get() *InlineResponse200237Items {
	return v.value
}

func (v *NullableInlineResponse200237Items) Set(val *InlineResponse200237Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200237Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200237Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200237Items(val *InlineResponse200237Items) *NullableInlineResponse200237Items {
	return &NullableInlineResponse200237Items{value: val, isSet: true}
}

func (v NullableInlineResponse200237Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200237Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


