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

// InlineResponse200343 struct for InlineResponse200343
type InlineResponse200343 struct {
	// Wireless LAN controller overview
	Items []InlineResponse200343Items `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta `json:"meta,omitempty"`
}

// NewInlineResponse200343 instantiates a new InlineResponse200343 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200343() *InlineResponse200343 {
	this := InlineResponse200343{}
	return &this
}

// NewInlineResponse200343WithDefaults instantiates a new InlineResponse200343 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200343WithDefaults() *InlineResponse200343 {
	this := InlineResponse200343{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200343) GetItems() []InlineResponse200343Items {
	if o == nil || isNil(o.Items) {
		var ret []InlineResponse200343Items
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343) GetItemsOk() ([]InlineResponse200343Items, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200343) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InlineResponse200343Items and assigns it to the Items field.
func (o *InlineResponse200343) SetItems(v []InlineResponse200343Items) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200343) GetMeta() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343) GetMetaOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200343) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta and assigns it to the Meta field.
func (o *InlineResponse200343) SetMeta(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta) {
	o.Meta = &v
}

func (o InlineResponse200343) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200343 struct {
	value *InlineResponse200343
	isSet bool
}

func (v NullableInlineResponse200343) Get() *InlineResponse200343 {
	return v.value
}

func (v *NullableInlineResponse200343) Set(val *InlineResponse200343) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200343) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200343) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200343(val *InlineResponse200343) *NullableInlineResponse200343 {
	return &NullableInlineResponse200343{value: val, isSet: true}
}

func (v NullableInlineResponse200343) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200343) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


