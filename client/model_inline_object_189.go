/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject189 struct for InlineObject189
type InlineObject189 struct {
	//       An ordered list of branding policy IDs that determines the priority order of how to apply the policies 
	BrandingPolicyIds []string `json:"brandingPolicyIds,omitempty"`
}

// NewInlineObject189 instantiates a new InlineObject189 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject189() *InlineObject189 {
	this := InlineObject189{}
	return &this
}

// NewInlineObject189WithDefaults instantiates a new InlineObject189 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject189WithDefaults() *InlineObject189 {
	this := InlineObject189{}
	return &this
}

// GetBrandingPolicyIds returns the BrandingPolicyIds field value if set, zero value otherwise.
func (o *InlineObject189) GetBrandingPolicyIds() []string {
	if o == nil || isNil(o.BrandingPolicyIds) {
		var ret []string
		return ret
	}
	return o.BrandingPolicyIds
}

// GetBrandingPolicyIdsOk returns a tuple with the BrandingPolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetBrandingPolicyIdsOk() ([]string, bool) {
	if o == nil || isNil(o.BrandingPolicyIds) {
    return nil, false
	}
	return o.BrandingPolicyIds, true
}

// HasBrandingPolicyIds returns a boolean if a field has been set.
func (o *InlineObject189) HasBrandingPolicyIds() bool {
	if o != nil && !isNil(o.BrandingPolicyIds) {
		return true
	}

	return false
}

// SetBrandingPolicyIds gets a reference to the given []string and assigns it to the BrandingPolicyIds field.
func (o *InlineObject189) SetBrandingPolicyIds(v []string) {
	o.BrandingPolicyIds = v
}

func (o InlineObject189) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BrandingPolicyIds) {
		toSerialize["brandingPolicyIds"] = o.BrandingPolicyIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject189 struct {
	value *InlineObject189
	isSet bool
}

func (v NullableInlineObject189) Get() *InlineObject189 {
	return v.value
}

func (v *NullableInlineObject189) Set(val *InlineObject189) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject189) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject189) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject189(val *InlineObject189) *NullableInlineObject189 {
	return &NullableInlineObject189{value: val, isSet: true}
}

func (v NullableInlineObject189) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject189) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


