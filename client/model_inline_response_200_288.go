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

// InlineResponse200288 struct for InlineResponse200288
type InlineResponse200288 struct {
	// Items in the paginated dataset
	Items []OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems `json:"items,omitempty"`
	Meta *InlineResponse200222Meta `json:"meta,omitempty"`
}

// NewInlineResponse200288 instantiates a new InlineResponse200288 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200288() *InlineResponse200288 {
	this := InlineResponse200288{}
	return &this
}

// NewInlineResponse200288WithDefaults instantiates a new InlineResponse200288 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200288WithDefaults() *InlineResponse200288 {
	this := InlineResponse200288{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200288) GetItems() []OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems {
	if o == nil || isNil(o.Items) {
		var ret []OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288) GetItemsOk() ([]OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200288) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems and assigns it to the Items field.
func (o *InlineResponse200288) SetItems(v []OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200288) GetMeta() InlineResponse200222Meta {
	if o == nil || isNil(o.Meta) {
		var ret InlineResponse200222Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288) GetMetaOk() (*InlineResponse200222Meta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200288) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse200222Meta and assigns it to the Meta field.
func (o *InlineResponse200288) SetMeta(v InlineResponse200222Meta) {
	o.Meta = &v
}

func (o InlineResponse200288) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200288 struct {
	value *InlineResponse200288
	isSet bool
}

func (v NullableInlineResponse200288) Get() *InlineResponse200288 {
	return v.value
}

func (v *NullableInlineResponse200288) Set(val *InlineResponse200288) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200288) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200288) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200288(val *InlineResponse200288) *NullableInlineResponse200288 {
	return &NullableInlineResponse200288{value: val, isSet: true}
}

func (v NullableInlineResponse200288) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200288) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


