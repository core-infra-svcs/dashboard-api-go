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

// InlineResponse200263 struct for InlineResponse200263
type InlineResponse200263 struct {
	// Items in the paginated dataset
	Items []OrganizationsOrganizationIdFloorPlansAutoLocateStatusesItems `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta `json:"meta,omitempty"`
}

// NewInlineResponse200263 instantiates a new InlineResponse200263 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200263() *InlineResponse200263 {
	this := InlineResponse200263{}
	return &this
}

// NewInlineResponse200263WithDefaults instantiates a new InlineResponse200263 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200263WithDefaults() *InlineResponse200263 {
	this := InlineResponse200263{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200263) GetItems() []OrganizationsOrganizationIdFloorPlansAutoLocateStatusesItems {
	if o == nil || isNil(o.Items) {
		var ret []OrganizationsOrganizationIdFloorPlansAutoLocateStatusesItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetItemsOk() ([]OrganizationsOrganizationIdFloorPlansAutoLocateStatusesItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200263) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrganizationsOrganizationIdFloorPlansAutoLocateStatusesItems and assigns it to the Items field.
func (o *InlineResponse200263) SetItems(v []OrganizationsOrganizationIdFloorPlansAutoLocateStatusesItems) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200263) GetMeta() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200263) GetMetaOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200263) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta and assigns it to the Meta field.
func (o *InlineResponse200263) SetMeta(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta) {
	o.Meta = &v
}

func (o InlineResponse200263) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200263 struct {
	value *InlineResponse200263
	isSet bool
}

func (v NullableInlineResponse200263) Get() *InlineResponse200263 {
	return v.value
}

func (v *NullableInlineResponse200263) Set(val *InlineResponse200263) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200263) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200263) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200263(val *InlineResponse200263) *NullableInlineResponse200263 {
	return &NullableInlineResponse200263{value: val, isSet: true}
}

func (v NullableInlineResponse200263) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200263) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


