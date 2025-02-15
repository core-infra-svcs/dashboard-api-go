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

// InlineResponse200339 struct for InlineResponse200339
type InlineResponse200339 struct {
	// The top-level property containing all power mode data.
	Items []InlineResponse200339Items `json:"items,omitempty"`
}

// NewInlineResponse200339 instantiates a new InlineResponse200339 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339() *InlineResponse200339 {
	this := InlineResponse200339{}
	return &this
}

// NewInlineResponse200339WithDefaults instantiates a new InlineResponse200339 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339WithDefaults() *InlineResponse200339 {
	this := InlineResponse200339{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200339) GetItems() []InlineResponse200339Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200339Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339) GetItemsOk() ([]InlineResponse200339Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200339) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200339Items and assigns it to the Items field.
func (o *InlineResponse200339) SetItems(v []InlineResponse200339Items) {
	o.Items = v
}

func (o InlineResponse200339) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339 struct {
	value *InlineResponse200339
	isSet bool
}

func (v NullableInlineResponse200339) Get() *InlineResponse200339 {
	return v.value
}

func (v *NullableInlineResponse200339) Set(val *InlineResponse200339) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339(val *InlineResponse200339) *NullableInlineResponse200339 {
	return &NullableInlineResponse200339{value: val, isSet: true}
}

func (v NullableInlineResponse200339) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


