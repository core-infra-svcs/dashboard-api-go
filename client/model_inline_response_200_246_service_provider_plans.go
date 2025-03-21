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

// InlineResponse200246ServiceProviderPlans struct for InlineResponse200246ServiceProviderPlans
type InlineResponse200246ServiceProviderPlans struct {
	// Plan name
	Name *string `json:"name,omitempty"`
	// Plan type (communication, rate)
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse200246ServiceProviderPlans instantiates a new InlineResponse200246ServiceProviderPlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200246ServiceProviderPlans() *InlineResponse200246ServiceProviderPlans {
	this := InlineResponse200246ServiceProviderPlans{}
	return &this
}

// NewInlineResponse200246ServiceProviderPlansWithDefaults instantiates a new InlineResponse200246ServiceProviderPlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200246ServiceProviderPlansWithDefaults() *InlineResponse200246ServiceProviderPlans {
	this := InlineResponse200246ServiceProviderPlans{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200246ServiceProviderPlans) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246ServiceProviderPlans) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200246ServiceProviderPlans) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200246ServiceProviderPlans) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200246ServiceProviderPlans) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246ServiceProviderPlans) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200246ServiceProviderPlans) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200246ServiceProviderPlans) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse200246ServiceProviderPlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200246ServiceProviderPlans struct {
	value *InlineResponse200246ServiceProviderPlans
	isSet bool
}

func (v NullableInlineResponse200246ServiceProviderPlans) Get() *InlineResponse200246ServiceProviderPlans {
	return v.value
}

func (v *NullableInlineResponse200246ServiceProviderPlans) Set(val *InlineResponse200246ServiceProviderPlans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200246ServiceProviderPlans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200246ServiceProviderPlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200246ServiceProviderPlans(val *InlineResponse200246ServiceProviderPlans) *NullableInlineResponse200246ServiceProviderPlans {
	return &NullableInlineResponse200246ServiceProviderPlans{value: val, isSet: true}
}

func (v NullableInlineResponse200246ServiceProviderPlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200246ServiceProviderPlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


