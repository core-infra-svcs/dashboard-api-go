/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20066LocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type InlineResponse20066LocalStatusPage struct {
	Authentication *InlineResponse20066LocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse20066LocalStatusPage instantiates a new InlineResponse20066LocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066LocalStatusPage() *InlineResponse20066LocalStatusPage {
	this := InlineResponse20066LocalStatusPage{}
	return &this
}

// NewInlineResponse20066LocalStatusPageWithDefaults instantiates a new InlineResponse20066LocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066LocalStatusPageWithDefaults() *InlineResponse20066LocalStatusPage {
	this := InlineResponse20066LocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse20066LocalStatusPage) GetAuthentication() InlineResponse20066LocalStatusPageAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse20066LocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066LocalStatusPage) GetAuthenticationOk() (*InlineResponse20066LocalStatusPageAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse20066LocalStatusPage) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse20066LocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *InlineResponse20066LocalStatusPage) SetAuthentication(v InlineResponse20066LocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse20066LocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20066LocalStatusPage struct {
	value *InlineResponse20066LocalStatusPage
	isSet bool
}

func (v NullableInlineResponse20066LocalStatusPage) Get() *InlineResponse20066LocalStatusPage {
	return v.value
}

func (v *NullableInlineResponse20066LocalStatusPage) Set(val *InlineResponse20066LocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066LocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066LocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066LocalStatusPage(val *InlineResponse20066LocalStatusPage) *NullableInlineResponse20066LocalStatusPage {
	return &NullableInlineResponse20066LocalStatusPage{value: val, isSet: true}
}

func (v NullableInlineResponse20066LocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066LocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


