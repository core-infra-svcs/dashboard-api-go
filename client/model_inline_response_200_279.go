/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200279 struct for InlineResponse200279
type InlineResponse200279 struct {
	// Sentry Group Policies for the Organization keyed by Network Id
	Items []InlineResponse200279Items `json:"items,omitempty"`
}

// NewInlineResponse200279 instantiates a new InlineResponse200279 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200279() *InlineResponse200279 {
	this := InlineResponse200279{}
	return &this
}

// NewInlineResponse200279WithDefaults instantiates a new InlineResponse200279 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200279WithDefaults() *InlineResponse200279 {
	this := InlineResponse200279{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200279) GetItems() []InlineResponse200279Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200279Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetItemsOk() ([]InlineResponse200279Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200279) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200279Items and assigns it to the Items field.
func (o *InlineResponse200279) SetItems(v []InlineResponse200279Items) {
	o.Items = v
}

func (o InlineResponse200279) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200279 struct {
	value *InlineResponse200279
	isSet bool
}

func (v NullableInlineResponse200279) Get() *InlineResponse200279 {
	return v.value
}

func (v *NullableInlineResponse200279) Set(val *InlineResponse200279) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200279) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200279) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200279(val *InlineResponse200279) *NullableInlineResponse200279 {
	return &NullableInlineResponse200279{value: val, isSet: true}
}

func (v NullableInlineResponse200279) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200279) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


