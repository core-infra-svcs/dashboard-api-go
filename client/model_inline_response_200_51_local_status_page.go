/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20051LocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type InlineResponse20051LocalStatusPage struct {
	Authentication *InlineResponse20051LocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse20051LocalStatusPage instantiates a new InlineResponse20051LocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20051LocalStatusPage() *InlineResponse20051LocalStatusPage {
	this := InlineResponse20051LocalStatusPage{}
	return &this
}

// NewInlineResponse20051LocalStatusPageWithDefaults instantiates a new InlineResponse20051LocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20051LocalStatusPageWithDefaults() *InlineResponse20051LocalStatusPage {
	this := InlineResponse20051LocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse20051LocalStatusPage) GetAuthentication() InlineResponse20051LocalStatusPageAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse20051LocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051LocalStatusPage) GetAuthenticationOk() (*InlineResponse20051LocalStatusPageAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse20051LocalStatusPage) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse20051LocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *InlineResponse20051LocalStatusPage) SetAuthentication(v InlineResponse20051LocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse20051LocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20051LocalStatusPage struct {
	value *InlineResponse20051LocalStatusPage
	isSet bool
}

func (v NullableInlineResponse20051LocalStatusPage) Get() *InlineResponse20051LocalStatusPage {
	return v.value
}

func (v *NullableInlineResponse20051LocalStatusPage) Set(val *InlineResponse20051LocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20051LocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20051LocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20051LocalStatusPage(val *InlineResponse20051LocalStatusPage) *NullableInlineResponse20051LocalStatusPage {
	return &NullableInlineResponse20051LocalStatusPage{value: val, isSet: true}
}

func (v NullableInlineResponse20051LocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20051LocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


