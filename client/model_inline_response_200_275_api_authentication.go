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

// InlineResponse200275ApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type InlineResponse200275ApiAuthentication struct {
	IpRestrictionsForKeys *InlineResponse200275ApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewInlineResponse200275ApiAuthentication instantiates a new InlineResponse200275ApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200275ApiAuthentication() *InlineResponse200275ApiAuthentication {
	this := InlineResponse200275ApiAuthentication{}
	return &this
}

// NewInlineResponse200275ApiAuthenticationWithDefaults instantiates a new InlineResponse200275ApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200275ApiAuthenticationWithDefaults() *InlineResponse200275ApiAuthentication {
	this := InlineResponse200275ApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *InlineResponse200275ApiAuthentication) GetIpRestrictionsForKeys() InlineResponse200275ApiAuthenticationIpRestrictionsForKeys {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
		var ret InlineResponse200275ApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200275ApiAuthentication) GetIpRestrictionsForKeysOk() (*InlineResponse200275ApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
    return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *InlineResponse200275ApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && !isNil(o.IpRestrictionsForKeys) {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given InlineResponse200275ApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *InlineResponse200275ApiAuthentication) SetIpRestrictionsForKeys(v InlineResponse200275ApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o InlineResponse200275ApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpRestrictionsForKeys) {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200275ApiAuthentication struct {
	value *InlineResponse200275ApiAuthentication
	isSet bool
}

func (v NullableInlineResponse200275ApiAuthentication) Get() *InlineResponse200275ApiAuthentication {
	return v.value
}

func (v *NullableInlineResponse200275ApiAuthentication) Set(val *InlineResponse200275ApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200275ApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200275ApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200275ApiAuthentication(val *InlineResponse200275ApiAuthentication) *NullableInlineResponse200275ApiAuthentication {
	return &NullableInlineResponse200275ApiAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse200275ApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200275ApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


