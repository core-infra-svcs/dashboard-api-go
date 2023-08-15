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

// InlineResponse200144ApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type InlineResponse200144ApiAuthentication struct {
	IpRestrictionsForKeys *InlineResponse200144ApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewInlineResponse200144ApiAuthentication instantiates a new InlineResponse200144ApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200144ApiAuthentication() *InlineResponse200144ApiAuthentication {
	this := InlineResponse200144ApiAuthentication{}
	return &this
}

// NewInlineResponse200144ApiAuthenticationWithDefaults instantiates a new InlineResponse200144ApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200144ApiAuthenticationWithDefaults() *InlineResponse200144ApiAuthentication {
	this := InlineResponse200144ApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *InlineResponse200144ApiAuthentication) GetIpRestrictionsForKeys() InlineResponse200144ApiAuthenticationIpRestrictionsForKeys {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
		var ret InlineResponse200144ApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200144ApiAuthentication) GetIpRestrictionsForKeysOk() (*InlineResponse200144ApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
    return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *InlineResponse200144ApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && !isNil(o.IpRestrictionsForKeys) {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given InlineResponse200144ApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *InlineResponse200144ApiAuthentication) SetIpRestrictionsForKeys(v InlineResponse200144ApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o InlineResponse200144ApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpRestrictionsForKeys) {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200144ApiAuthentication struct {
	value *InlineResponse200144ApiAuthentication
	isSet bool
}

func (v NullableInlineResponse200144ApiAuthentication) Get() *InlineResponse200144ApiAuthentication {
	return v.value
}

func (v *NullableInlineResponse200144ApiAuthentication) Set(val *InlineResponse200144ApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200144ApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200144ApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200144ApiAuthentication(val *InlineResponse200144ApiAuthentication) *NullableInlineResponse200144ApiAuthentication {
	return &NullableInlineResponse200144ApiAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse200144ApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200144ApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


