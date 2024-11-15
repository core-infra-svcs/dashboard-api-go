/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200323 struct for InlineResponse200323
type InlineResponse200323 struct {
	// The top-level propery containing all status data.
	Items []OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta `json:"meta,omitempty"`
}

// NewInlineResponse200323 instantiates a new InlineResponse200323 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200323() *InlineResponse200323 {
	this := InlineResponse200323{}
	return &this
}

// NewInlineResponse200323WithDefaults instantiates a new InlineResponse200323 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200323WithDefaults() *InlineResponse200323 {
	this := InlineResponse200323{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200323) GetItems() []OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems {
	if o == nil || isNil(o.Items) {
		var ret []OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200323) GetItemsOk() ([]OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200323) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems and assigns it to the Items field.
func (o *InlineResponse200323) SetItems(v []OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceItems) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200323) GetMeta() OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200323) GetMetaOk() (*OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200323) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta and assigns it to the Meta field.
func (o *InlineResponse200323) SetMeta(v OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) {
	o.Meta = &v
}

func (o InlineResponse200323) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200323 struct {
	value *InlineResponse200323
	isSet bool
}

func (v NullableInlineResponse200323) Get() *InlineResponse200323 {
	return v.value
}

func (v *NullableInlineResponse200323) Set(val *InlineResponse200323) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200323) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200323) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200323(val *InlineResponse200323) *NullableInlineResponse200323 {
	return &NullableInlineResponse200323{value: val, isSet: true}
}

func (v NullableInlineResponse200323) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200323) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


