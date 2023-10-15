/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200155ApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type InlineResponse200155ApiAuthentication struct {
	IpRestrictionsForKeys *InlineResponse200155ApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewInlineResponse200155ApiAuthentication instantiates a new InlineResponse200155ApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200155ApiAuthentication() *InlineResponse200155ApiAuthentication {
	this := InlineResponse200155ApiAuthentication{}
	return &this
}

// NewInlineResponse200155ApiAuthenticationWithDefaults instantiates a new InlineResponse200155ApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200155ApiAuthenticationWithDefaults() *InlineResponse200155ApiAuthentication {
	this := InlineResponse200155ApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *InlineResponse200155ApiAuthentication) GetIpRestrictionsForKeys() InlineResponse200155ApiAuthenticationIpRestrictionsForKeys {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
		var ret InlineResponse200155ApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200155ApiAuthentication) GetIpRestrictionsForKeysOk() (*InlineResponse200155ApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
    return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *InlineResponse200155ApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && !isNil(o.IpRestrictionsForKeys) {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given InlineResponse200155ApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *InlineResponse200155ApiAuthentication) SetIpRestrictionsForKeys(v InlineResponse200155ApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o InlineResponse200155ApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpRestrictionsForKeys) {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200155ApiAuthentication struct {
	value *InlineResponse200155ApiAuthentication
	isSet bool
}

func (v NullableInlineResponse200155ApiAuthentication) Get() *InlineResponse200155ApiAuthentication {
	return v.value
}

func (v *NullableInlineResponse200155ApiAuthentication) Set(val *InlineResponse200155ApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200155ApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200155ApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200155ApiAuthentication(val *InlineResponse200155ApiAuthentication) *NullableInlineResponse200155ApiAuthentication {
	return &NullableInlineResponse200155ApiAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse200155ApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200155ApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


