/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200111RegulatoryDomain Regulatory domain information for this network.
type InlineResponse200111RegulatoryDomain struct {
	// The name of the regulatory domain for this network.
	Name *string `json:"name,omitempty"`
	// Whether or not the regulatory domain for this network permits Wifi 6E.
	Permits6e *bool `json:"permits6e,omitempty"`
}

// NewInlineResponse200111RegulatoryDomain instantiates a new InlineResponse200111RegulatoryDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200111RegulatoryDomain() *InlineResponse200111RegulatoryDomain {
	this := InlineResponse200111RegulatoryDomain{}
	return &this
}

// NewInlineResponse200111RegulatoryDomainWithDefaults instantiates a new InlineResponse200111RegulatoryDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200111RegulatoryDomainWithDefaults() *InlineResponse200111RegulatoryDomain {
	this := InlineResponse200111RegulatoryDomain{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200111RegulatoryDomain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111RegulatoryDomain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200111RegulatoryDomain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200111RegulatoryDomain) SetName(v string) {
	o.Name = &v
}

// GetPermits6e returns the Permits6e field value if set, zero value otherwise.
func (o *InlineResponse200111RegulatoryDomain) GetPermits6e() bool {
	if o == nil || isNil(o.Permits6e) {
		var ret bool
		return ret
	}
	return *o.Permits6e
}

// GetPermits6eOk returns a tuple with the Permits6e field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111RegulatoryDomain) GetPermits6eOk() (*bool, bool) {
	if o == nil || isNil(o.Permits6e) {
    return nil, false
	}
	return o.Permits6e, true
}

// HasPermits6e returns a boolean if a field has been set.
func (o *InlineResponse200111RegulatoryDomain) HasPermits6e() bool {
	if o != nil && !isNil(o.Permits6e) {
		return true
	}

	return false
}

// SetPermits6e gets a reference to the given bool and assigns it to the Permits6e field.
func (o *InlineResponse200111RegulatoryDomain) SetPermits6e(v bool) {
	o.Permits6e = &v
}

func (o InlineResponse200111RegulatoryDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Permits6e) {
		toSerialize["permits6e"] = o.Permits6e
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200111RegulatoryDomain struct {
	value *InlineResponse200111RegulatoryDomain
	isSet bool
}

func (v NullableInlineResponse200111RegulatoryDomain) Get() *InlineResponse200111RegulatoryDomain {
	return v.value
}

func (v *NullableInlineResponse200111RegulatoryDomain) Set(val *InlineResponse200111RegulatoryDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200111RegulatoryDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200111RegulatoryDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200111RegulatoryDomain(val *InlineResponse200111RegulatoryDomain) *NullableInlineResponse200111RegulatoryDomain {
	return &NullableInlineResponse200111RegulatoryDomain{value: val, isSet: true}
}

func (v NullableInlineResponse200111RegulatoryDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200111RegulatoryDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


