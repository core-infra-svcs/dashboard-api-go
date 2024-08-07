/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200220 struct for InlineResponse200220
type InlineResponse200220 struct {
	//       An ordered list of branding policy IDs that determines the priority order of how to apply the policies 
	BrandingPolicyIds []string `json:"brandingPolicyIds,omitempty"`
}

// NewInlineResponse200220 instantiates a new InlineResponse200220 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200220() *InlineResponse200220 {
	this := InlineResponse200220{}
	return &this
}

// NewInlineResponse200220WithDefaults instantiates a new InlineResponse200220 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200220WithDefaults() *InlineResponse200220 {
	this := InlineResponse200220{}
	return &this
}

// GetBrandingPolicyIds returns the BrandingPolicyIds field value if set, zero value otherwise.
func (o *InlineResponse200220) GetBrandingPolicyIds() []string {
	if o == nil || isNil(o.BrandingPolicyIds) {
		var ret []string
		return ret
	}
	return o.BrandingPolicyIds
}

// GetBrandingPolicyIdsOk returns a tuple with the BrandingPolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200220) GetBrandingPolicyIdsOk() ([]string, bool) {
	if o == nil || isNil(o.BrandingPolicyIds) {
    return nil, false
	}
	return o.BrandingPolicyIds, true
}

// HasBrandingPolicyIds returns a boolean if a field has been set.
func (o *InlineResponse200220) HasBrandingPolicyIds() bool {
	if o != nil && !isNil(o.BrandingPolicyIds) {
		return true
	}

	return false
}

// SetBrandingPolicyIds gets a reference to the given []string and assigns it to the BrandingPolicyIds field.
func (o *InlineResponse200220) SetBrandingPolicyIds(v []string) {
	o.BrandingPolicyIds = v
}

func (o InlineResponse200220) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BrandingPolicyIds) {
		toSerialize["brandingPolicyIds"] = o.BrandingPolicyIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200220 struct {
	value *InlineResponse200220
	isSet bool
}

func (v NullableInlineResponse200220) Get() *InlineResponse200220 {
	return v.value
}

func (v *NullableInlineResponse200220) Set(val *InlineResponse200220) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200220) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200220) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200220(val *InlineResponse200220) *NullableInlineResponse200220 {
	return &NullableInlineResponse200220{value: val, isSet: true}
}

func (v NullableInlineResponse200220) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200220) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


