/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200327Tags struct for InlineResponse200327Tags
type InlineResponse200327Tags struct {
	// Policy tag
	Policy *string `json:"policy,omitempty"`
	// Site tag
	Site *string `json:"site,omitempty"`
	// RF tag
	Rf *string `json:"rf,omitempty"`
}

// NewInlineResponse200327Tags instantiates a new InlineResponse200327Tags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200327Tags() *InlineResponse200327Tags {
	this := InlineResponse200327Tags{}
	return &this
}

// NewInlineResponse200327TagsWithDefaults instantiates a new InlineResponse200327Tags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200327TagsWithDefaults() *InlineResponse200327Tags {
	this := InlineResponse200327Tags{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *InlineResponse200327Tags) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200327Tags) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *InlineResponse200327Tags) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *InlineResponse200327Tags) SetPolicy(v string) {
	o.Policy = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *InlineResponse200327Tags) GetSite() string {
	if o == nil || isNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200327Tags) GetSiteOk() (*string, bool) {
	if o == nil || isNil(o.Site) {
    return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *InlineResponse200327Tags) HasSite() bool {
	if o != nil && !isNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *InlineResponse200327Tags) SetSite(v string) {
	o.Site = &v
}

// GetRf returns the Rf field value if set, zero value otherwise.
func (o *InlineResponse200327Tags) GetRf() string {
	if o == nil || isNil(o.Rf) {
		var ret string
		return ret
	}
	return *o.Rf
}

// GetRfOk returns a tuple with the Rf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200327Tags) GetRfOk() (*string, bool) {
	if o == nil || isNil(o.Rf) {
    return nil, false
	}
	return o.Rf, true
}

// HasRf returns a boolean if a field has been set.
func (o *InlineResponse200327Tags) HasRf() bool {
	if o != nil && !isNil(o.Rf) {
		return true
	}

	return false
}

// SetRf gets a reference to the given string and assigns it to the Rf field.
func (o *InlineResponse200327Tags) SetRf(v string) {
	o.Rf = &v
}

func (o InlineResponse200327Tags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !isNil(o.Rf) {
		toSerialize["rf"] = o.Rf
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200327Tags struct {
	value *InlineResponse200327Tags
	isSet bool
}

func (v NullableInlineResponse200327Tags) Get() *InlineResponse200327Tags {
	return v.value
}

func (v *NullableInlineResponse200327Tags) Set(val *InlineResponse200327Tags) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200327Tags) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200327Tags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200327Tags(val *InlineResponse200327Tags) *NullableInlineResponse200327Tags {
	return &NullableInlineResponse200327Tags{value: val, isSet: true}
}

func (v NullableInlineResponse200327Tags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200327Tags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


