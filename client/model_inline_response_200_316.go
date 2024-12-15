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

// InlineResponse200316 struct for InlineResponse200316
type InlineResponse200316 struct {
	// Access point client count
	Items []InlineResponse200316Items `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta `json:"meta,omitempty"`
}

// NewInlineResponse200316 instantiates a new InlineResponse200316 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200316() *InlineResponse200316 {
	this := InlineResponse200316{}
	return &this
}

// NewInlineResponse200316WithDefaults instantiates a new InlineResponse200316 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200316WithDefaults() *InlineResponse200316 {
	this := InlineResponse200316{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200316) GetItems() []InlineResponse200316Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200316Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200316) GetItemsOk() ([]InlineResponse200316Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200316) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200316Items and assigns it to the Items field.
func (o *InlineResponse200316) SetItems(v []InlineResponse200316Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200316) GetMeta() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200316) GetMetaOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200316) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta and assigns it to the Meta field.
func (o *InlineResponse200316) SetMeta(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta) {
	o.Meta = &v
}

func (o InlineResponse200316) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200316 struct {
	value *InlineResponse200316
	isSet bool
}

func (v NullableInlineResponse200316) Get() *InlineResponse200316 {
	return v.value
}

func (v *NullableInlineResponse200316) Set(val *InlineResponse200316) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200316) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200316) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200316(val *InlineResponse200316) *NullableInlineResponse200316 {
	return &NullableInlineResponse200316{value: val, isSet: true}
}

func (v NullableInlineResponse200316) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200316) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


