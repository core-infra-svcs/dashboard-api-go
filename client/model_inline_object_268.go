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

// InlineObject268 struct for InlineObject268
type InlineObject268 struct {
	// Sentry Group Policies for the Organization keyed by Network Id
	Items []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems `json:"items"`
}

// NewInlineObject268 instantiates a new InlineObject268 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject268(items []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) *InlineObject268 {
	this := InlineObject268{}
	this.Items = items
	return &this
}

// NewInlineObject268WithDefaults instantiates a new InlineObject268 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject268WithDefaults() *InlineObject268 {
	this := InlineObject268{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineObject268) GetItems() []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems {
	if o == nil {
		var ret []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineObject268) GetItemsOk() ([]OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InlineObject268) SetItems(v []OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems) {
	o.Items = v
}

func (o InlineObject268) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject268 struct {
	value *InlineObject268
	isSet bool
}

func (v NullableInlineObject268) Get() *InlineObject268 {
	return v.value
}

func (v *NullableInlineObject268) Set(val *InlineObject268) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject268) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject268) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject268(val *InlineObject268) *NullableInlineObject268 {
	return &NullableInlineObject268{value: val, isSet: true}
}

func (v NullableInlineObject268) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject268) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

