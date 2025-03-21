/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200195RegulatoryDomain Regulatory domain information for this network.
type InlineResponse200195RegulatoryDomain struct {
	// The name of the regulatory domain for this network.
	Name *string `json:"name,omitempty"`
	// The country code of the regulatory domain.
	CountryCode *string `json:"countryCode,omitempty"`
	// Whether or not the regulatory domain for this network permits Wifi 6E.
	Permits6e *bool `json:"permits6e,omitempty"`
}

// NewInlineResponse200195RegulatoryDomain instantiates a new InlineResponse200195RegulatoryDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200195RegulatoryDomain() *InlineResponse200195RegulatoryDomain {
	this := InlineResponse200195RegulatoryDomain{}
	return &this
}

// NewInlineResponse200195RegulatoryDomainWithDefaults instantiates a new InlineResponse200195RegulatoryDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200195RegulatoryDomainWithDefaults() *InlineResponse200195RegulatoryDomain {
	this := InlineResponse200195RegulatoryDomain{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200195RegulatoryDomain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195RegulatoryDomain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200195RegulatoryDomain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200195RegulatoryDomain) SetName(v string) {
	o.Name = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *InlineResponse200195RegulatoryDomain) GetCountryCode() string {
	if o == nil || isNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195RegulatoryDomain) GetCountryCodeOk() (*string, bool) {
	if o == nil || isNil(o.CountryCode) {
    return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *InlineResponse200195RegulatoryDomain) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *InlineResponse200195RegulatoryDomain) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPermits6e returns the Permits6e field value if set, zero value otherwise.
func (o *InlineResponse200195RegulatoryDomain) GetPermits6e() bool {
	if o == nil || isNil(o.Permits6e) {
		var ret bool
		return ret
	}
	return *o.Permits6e
}

// GetPermits6eOk returns a tuple with the Permits6e field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200195RegulatoryDomain) GetPermits6eOk() (*bool, bool) {
	if o == nil || isNil(o.Permits6e) {
    return nil, false
	}
	return o.Permits6e, true
}

// HasPermits6e returns a boolean if a field has been set.
func (o *InlineResponse200195RegulatoryDomain) HasPermits6e() bool {
	if o != nil && !isNil(o.Permits6e) {
		return true
	}

	return false
}

// SetPermits6e gets a reference to the given bool and assigns it to the Permits6e field.
func (o *InlineResponse200195RegulatoryDomain) SetPermits6e(v bool) {
	o.Permits6e = &v
}

func (o InlineResponse200195RegulatoryDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !isNil(o.Permits6e) {
		toSerialize["permits6e"] = o.Permits6e
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200195RegulatoryDomain struct {
	value *InlineResponse200195RegulatoryDomain
	isSet bool
}

func (v NullableInlineResponse200195RegulatoryDomain) Get() *InlineResponse200195RegulatoryDomain {
	return v.value
}

func (v *NullableInlineResponse200195RegulatoryDomain) Set(val *InlineResponse200195RegulatoryDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200195RegulatoryDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200195RegulatoryDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200195RegulatoryDomain(val *InlineResponse200195RegulatoryDomain) *NullableInlineResponse200195RegulatoryDomain {
	return &NullableInlineResponse200195RegulatoryDomain{value: val, isSet: true}
}

func (v NullableInlineResponse200195RegulatoryDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200195RegulatoryDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


