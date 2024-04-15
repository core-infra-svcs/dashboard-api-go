/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200108LocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type InlineResponse200108LocalStatusPage struct {
	Authentication *InlineResponse200108LocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse200108LocalStatusPage instantiates a new InlineResponse200108LocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200108LocalStatusPage() *InlineResponse200108LocalStatusPage {
	this := InlineResponse200108LocalStatusPage{}
	return &this
}

// NewInlineResponse200108LocalStatusPageWithDefaults instantiates a new InlineResponse200108LocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200108LocalStatusPageWithDefaults() *InlineResponse200108LocalStatusPage {
	this := InlineResponse200108LocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse200108LocalStatusPage) GetAuthentication() InlineResponse200108LocalStatusPageAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse200108LocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108LocalStatusPage) GetAuthenticationOk() (*InlineResponse200108LocalStatusPageAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse200108LocalStatusPage) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse200108LocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *InlineResponse200108LocalStatusPage) SetAuthentication(v InlineResponse200108LocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse200108LocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200108LocalStatusPage struct {
	value *InlineResponse200108LocalStatusPage
	isSet bool
}

func (v NullableInlineResponse200108LocalStatusPage) Get() *InlineResponse200108LocalStatusPage {
	return v.value
}

func (v *NullableInlineResponse200108LocalStatusPage) Set(val *InlineResponse200108LocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200108LocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200108LocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200108LocalStatusPage(val *InlineResponse200108LocalStatusPage) *NullableInlineResponse200108LocalStatusPage {
	return &NullableInlineResponse200108LocalStatusPage{value: val, isSet: true}
}

func (v NullableInlineResponse200108LocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200108LocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


