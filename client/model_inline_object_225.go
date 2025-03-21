/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject225 struct for InlineObject225
type InlineObject225 struct {
	// List containing the network ID and Profile ID
	Items []OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems `json:"items"`
}

// NewInlineObject225 instantiates a new InlineObject225 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject225(items []OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) *InlineObject225 {
	this := InlineObject225{}
	this.Items = items
	return &this
}

// NewInlineObject225WithDefaults instantiates a new InlineObject225 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject225WithDefaults() *InlineObject225 {
	this := InlineObject225{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineObject225) GetItems() []OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineObject225) GetItemsOk() ([]OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InlineObject225) SetItems(v []OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems) {
	o.Items = v
}

func (o InlineObject225) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject225 struct {
	value *InlineObject225
	isSet bool
}

func (v NullableInlineObject225) Get() *InlineObject225 {
	return v.value
}

func (v *NullableInlineObject225) Set(val *InlineObject225) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject225) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject225) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject225(val *InlineObject225) *NullableInlineObject225 {
	return &NullableInlineObject225{value: val, isSet: true}
}

func (v NullableInlineObject225) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject225) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


