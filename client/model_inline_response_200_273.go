/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200273 struct for InlineResponse200273
type InlineResponse200273 struct {
	// Sentry Group Policies for the Organization keyed by Network Id
	Items []InlineResponse200273Items `json:"items,omitempty"`
}

// NewInlineResponse200273 instantiates a new InlineResponse200273 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200273() *InlineResponse200273 {
	this := InlineResponse200273{}
	return &this
}

// NewInlineResponse200273WithDefaults instantiates a new InlineResponse200273 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200273WithDefaults() *InlineResponse200273 {
	this := InlineResponse200273{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200273) GetItems() []InlineResponse200273Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200273Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200273) GetItemsOk() ([]InlineResponse200273Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200273) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200273Items and assigns it to the Items field.
func (o *InlineResponse200273) SetItems(v []InlineResponse200273Items) {
	o.Items = v
}

func (o InlineResponse200273) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200273 struct {
	value *InlineResponse200273
	isSet bool
}

func (v NullableInlineResponse200273) Get() *InlineResponse200273 {
	return v.value
}

func (v *NullableInlineResponse200273) Set(val *InlineResponse200273) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200273) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200273) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200273(val *InlineResponse200273) *NullableInlineResponse200273 {
	return &NullableInlineResponse200273{value: val, isSet: true}
}

func (v NullableInlineResponse200273) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200273) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


