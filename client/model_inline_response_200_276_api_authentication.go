/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200276ApiAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
type InlineResponse200276ApiAuthentication struct {
	IpRestrictionsForKeys *InlineResponse200276ApiAuthenticationIpRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// NewInlineResponse200276ApiAuthentication instantiates a new InlineResponse200276ApiAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200276ApiAuthentication() *InlineResponse200276ApiAuthentication {
	this := InlineResponse200276ApiAuthentication{}
	return &this
}

// NewInlineResponse200276ApiAuthenticationWithDefaults instantiates a new InlineResponse200276ApiAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200276ApiAuthenticationWithDefaults() *InlineResponse200276ApiAuthentication {
	this := InlineResponse200276ApiAuthentication{}
	return &this
}

// GetIpRestrictionsForKeys returns the IpRestrictionsForKeys field value if set, zero value otherwise.
func (o *InlineResponse200276ApiAuthentication) GetIpRestrictionsForKeys() InlineResponse200276ApiAuthenticationIpRestrictionsForKeys {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
		var ret InlineResponse200276ApiAuthenticationIpRestrictionsForKeys
		return ret
	}
	return *o.IpRestrictionsForKeys
}

// GetIpRestrictionsForKeysOk returns a tuple with the IpRestrictionsForKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276ApiAuthentication) GetIpRestrictionsForKeysOk() (*InlineResponse200276ApiAuthenticationIpRestrictionsForKeys, bool) {
	if o == nil || isNil(o.IpRestrictionsForKeys) {
    return nil, false
	}
	return o.IpRestrictionsForKeys, true
}

// HasIpRestrictionsForKeys returns a boolean if a field has been set.
func (o *InlineResponse200276ApiAuthentication) HasIpRestrictionsForKeys() bool {
	if o != nil && !isNil(o.IpRestrictionsForKeys) {
		return true
	}

	return false
}

// SetIpRestrictionsForKeys gets a reference to the given InlineResponse200276ApiAuthenticationIpRestrictionsForKeys and assigns it to the IpRestrictionsForKeys field.
func (o *InlineResponse200276ApiAuthentication) SetIpRestrictionsForKeys(v InlineResponse200276ApiAuthenticationIpRestrictionsForKeys) {
	o.IpRestrictionsForKeys = &v
}

func (o InlineResponse200276ApiAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpRestrictionsForKeys) {
		toSerialize["ipRestrictionsForKeys"] = o.IpRestrictionsForKeys
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200276ApiAuthentication struct {
	value *InlineResponse200276ApiAuthentication
	isSet bool
}

func (v NullableInlineResponse200276ApiAuthentication) Get() *InlineResponse200276ApiAuthentication {
	return v.value
}

func (v *NullableInlineResponse200276ApiAuthentication) Set(val *InlineResponse200276ApiAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200276ApiAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200276ApiAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200276ApiAuthentication(val *InlineResponse200276ApiAuthentication) *NullableInlineResponse200276ApiAuthentication {
	return &NullableInlineResponse200276ApiAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse200276ApiAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200276ApiAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


