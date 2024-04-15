/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	// Username
	Name *string `json:"name,omitempty"`
	// User email
	Email *string `json:"email,omitempty"`
	// Last seen active on Dashboard UI
	LastUsedDashboardAt *time.Time `json:"lastUsedDashboardAt,omitempty"`
	Authentication *InlineResponse200Authentication `json:"authentication,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse200) SetEmail(v string) {
	o.Email = &v
}

// GetLastUsedDashboardAt returns the LastUsedDashboardAt field value if set, zero value otherwise.
func (o *InlineResponse200) GetLastUsedDashboardAt() time.Time {
	if o == nil || isNil(o.LastUsedDashboardAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedDashboardAt
}

// GetLastUsedDashboardAtOk returns a tuple with the LastUsedDashboardAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetLastUsedDashboardAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUsedDashboardAt) {
    return nil, false
	}
	return o.LastUsedDashboardAt, true
}

// HasLastUsedDashboardAt returns a boolean if a field has been set.
func (o *InlineResponse200) HasLastUsedDashboardAt() bool {
	if o != nil && !isNil(o.LastUsedDashboardAt) {
		return true
	}

	return false
}

// SetLastUsedDashboardAt gets a reference to the given time.Time and assigns it to the LastUsedDashboardAt field.
func (o *InlineResponse200) SetLastUsedDashboardAt(v time.Time) {
	o.LastUsedDashboardAt = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse200) GetAuthentication() InlineResponse200Authentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse200Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetAuthenticationOk() (*InlineResponse200Authentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse200) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse200Authentication and assigns it to the Authentication field.
func (o *InlineResponse200) SetAuthentication(v InlineResponse200Authentication) {
	o.Authentication = &v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.LastUsedDashboardAt) {
		toSerialize["lastUsedDashboardAt"] = o.LastUsedDashboardAt
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


