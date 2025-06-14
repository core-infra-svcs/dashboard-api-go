/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200121LocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type InlineResponse200121LocalStatusPage struct {
	Authentication *InlineResponse200121LocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse200121LocalStatusPage instantiates a new InlineResponse200121LocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200121LocalStatusPage() *InlineResponse200121LocalStatusPage {
	this := InlineResponse200121LocalStatusPage{}
	return &this
}

// NewInlineResponse200121LocalStatusPageWithDefaults instantiates a new InlineResponse200121LocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200121LocalStatusPageWithDefaults() *InlineResponse200121LocalStatusPage {
	this := InlineResponse200121LocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse200121LocalStatusPage) GetAuthentication() InlineResponse200121LocalStatusPageAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse200121LocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121LocalStatusPage) GetAuthenticationOk() (*InlineResponse200121LocalStatusPageAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse200121LocalStatusPage) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse200121LocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *InlineResponse200121LocalStatusPage) SetAuthentication(v InlineResponse200121LocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse200121LocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200121LocalStatusPage struct {
	value *InlineResponse200121LocalStatusPage
	isSet bool
}

func (v NullableInlineResponse200121LocalStatusPage) Get() *InlineResponse200121LocalStatusPage {
	return v.value
}

func (v *NullableInlineResponse200121LocalStatusPage) Set(val *InlineResponse200121LocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200121LocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200121LocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200121LocalStatusPage(val *InlineResponse200121LocalStatusPage) *NullableInlineResponse200121LocalStatusPage {
	return &NullableInlineResponse200121LocalStatusPage{value: val, isSet: true}
}

func (v NullableInlineResponse200121LocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200121LocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


