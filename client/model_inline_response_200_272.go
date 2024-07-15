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

// InlineResponse200272 struct for InlineResponse200272
type InlineResponse200272 struct {
	// Sentry Group Policies for the Organization keyed by Network Id
	Items []InlineResponse200272Items `json:"items,omitempty"`
}

// NewInlineResponse200272 instantiates a new InlineResponse200272 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200272() *InlineResponse200272 {
	this := InlineResponse200272{}
	return &this
}

// NewInlineResponse200272WithDefaults instantiates a new InlineResponse200272 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200272WithDefaults() *InlineResponse200272 {
	this := InlineResponse200272{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200272) GetItems() []InlineResponse200272Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200272Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272) GetItemsOk() ([]InlineResponse200272Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200272) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200272Items and assigns it to the Items field.
func (o *InlineResponse200272) SetItems(v []InlineResponse200272Items) {
	o.Items = v
}

func (o InlineResponse200272) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200272 struct {
	value *InlineResponse200272
	isSet bool
}

func (v NullableInlineResponse200272) Get() *InlineResponse200272 {
	return v.value
}

func (v *NullableInlineResponse200272) Set(val *InlineResponse200272) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200272) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200272) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200272(val *InlineResponse200272) *NullableInlineResponse200272 {
	return &NullableInlineResponse200272{value: val, isSet: true}
}

func (v NullableInlineResponse200272) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200272) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

