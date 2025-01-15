/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200277ApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type InlineResponse200277ApiAuthentication struct {
	IpRestrictionsForKeys *InlineResponse200277ApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewInlineResponse200277ApiAuthentication instantiates a new InlineResponse200277ApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200277ApiAuthentication() *InlineResponse200277ApiAuthentication {
	this := InlineResponse200277ApiAuthentication{}
	return &this
}

// NewInlineResponse200277ApiAuthenticationWithDefaults instantiates a new InlineResponse200277ApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200277ApiAuthenticationWithDefaults() *InlineResponse200277ApiAuthentication {
	this := InlineResponse200277ApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *InlineResponse200277ApiAuthentication) GetIpRestrictionsForKeys() InlineResponse200277ApiAuthenticationIpRestrictionsForKeys {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
		var ret InlineResponse200277ApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200277ApiAuthentication) GetIpRestrictionsForKeysOk() (*InlineResponse200277ApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
    return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *InlineResponse200277ApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && !isNil(o.IpRestrictionsForKeys) {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given InlineResponse200277ApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *InlineResponse200277ApiAuthentication) SetIpRestrictionsForKeys(v InlineResponse200277ApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o InlineResponse200277ApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpRestrictionsForKeys) {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200277ApiAuthentication struct {
	value *InlineResponse200277ApiAuthentication
	isSet bool
}

func (v NullableInlineResponse200277ApiAuthentication) Get() *InlineResponse200277ApiAuthentication {
	return v.value
}

func (v *NullableInlineResponse200277ApiAuthentication) Set(val *InlineResponse200277ApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200277ApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200277ApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200277ApiAuthentication(val *InlineResponse200277ApiAuthentication) *NullableInlineResponse200277ApiAuthentication {
	return &NullableInlineResponse200277ApiAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse200277ApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200277ApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

