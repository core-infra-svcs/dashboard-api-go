/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200289ApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type InlineResponse200289ApiAuthentication struct {
	IpRestrictionsForKeys *InlineResponse200289ApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewInlineResponse200289ApiAuthentication instantiates a new InlineResponse200289ApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200289ApiAuthentication() *InlineResponse200289ApiAuthentication {
	this := InlineResponse200289ApiAuthentication{}
	return &this
}

// NewInlineResponse200289ApiAuthenticationWithDefaults instantiates a new InlineResponse200289ApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200289ApiAuthenticationWithDefaults() *InlineResponse200289ApiAuthentication {
	this := InlineResponse200289ApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *InlineResponse200289ApiAuthentication) GetIpRestrictionsForKeys() InlineResponse200289ApiAuthenticationIpRestrictionsForKeys {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
		var ret InlineResponse200289ApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289ApiAuthentication) GetIpRestrictionsForKeysOk() (*InlineResponse200289ApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
    return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *InlineResponse200289ApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && !isNil(o.IpRestrictionsForKeys) {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given InlineResponse200289ApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *InlineResponse200289ApiAuthentication) SetIpRestrictionsForKeys(v InlineResponse200289ApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o InlineResponse200289ApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpRestrictionsForKeys) {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200289ApiAuthentication struct {
	value *InlineResponse200289ApiAuthentication
	isSet bool
}

func (v NullableInlineResponse200289ApiAuthentication) Get() *InlineResponse200289ApiAuthentication {
	return v.value
}

func (v *NullableInlineResponse200289ApiAuthentication) Set(val *InlineResponse200289ApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200289ApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200289ApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200289ApiAuthentication(val *InlineResponse200289ApiAuthentication) *NullableInlineResponse200289ApiAuthentication {
	return &NullableInlineResponse200289ApiAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse200289ApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200289ApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


