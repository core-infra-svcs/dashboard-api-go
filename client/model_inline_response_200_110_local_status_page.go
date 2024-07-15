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

// InlineResponse200110LocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type InlineResponse200110LocalStatusPage struct {
	Authentication *InlineResponse200110LocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse200110LocalStatusPage instantiates a new InlineResponse200110LocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110LocalStatusPage() *InlineResponse200110LocalStatusPage {
	this := InlineResponse200110LocalStatusPage{}
	return &this
}

// NewInlineResponse200110LocalStatusPageWithDefaults instantiates a new InlineResponse200110LocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110LocalStatusPageWithDefaults() *InlineResponse200110LocalStatusPage {
	this := InlineResponse200110LocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse200110LocalStatusPage) GetAuthentication() InlineResponse200110LocalStatusPageAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse200110LocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110LocalStatusPage) GetAuthenticationOk() (*InlineResponse200110LocalStatusPageAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse200110LocalStatusPage) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse200110LocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *InlineResponse200110LocalStatusPage) SetAuthentication(v InlineResponse200110LocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse200110LocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110LocalStatusPage struct {
	value *InlineResponse200110LocalStatusPage
	isSet bool
}

func (v NullableInlineResponse200110LocalStatusPage) Get() *InlineResponse200110LocalStatusPage {
	return v.value
}

func (v *NullableInlineResponse200110LocalStatusPage) Set(val *InlineResponse200110LocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110LocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110LocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110LocalStatusPage(val *InlineResponse200110LocalStatusPage) *NullableInlineResponse200110LocalStatusPage {
	return &NullableInlineResponse200110LocalStatusPage{value: val, isSet: true}
}

func (v NullableInlineResponse200110LocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110LocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


