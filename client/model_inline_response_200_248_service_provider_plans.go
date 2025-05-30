/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200248ServiceProviderPlans struct for InlineResponse200248ServiceProviderPlans
type InlineResponse200248ServiceProviderPlans struct {
	// Plan name
	Name *string `json:"name,omitempty"`
	// Plan type (communication, rate)
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse200248ServiceProviderPlans instantiates a new InlineResponse200248ServiceProviderPlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200248ServiceProviderPlans() *InlineResponse200248ServiceProviderPlans {
	this := InlineResponse200248ServiceProviderPlans{}
	return &this
}

// NewInlineResponse200248ServiceProviderPlansWithDefaults instantiates a new InlineResponse200248ServiceProviderPlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200248ServiceProviderPlansWithDefaults() *InlineResponse200248ServiceProviderPlans {
	this := InlineResponse200248ServiceProviderPlans{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200248ServiceProviderPlans) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248ServiceProviderPlans) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200248ServiceProviderPlans) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200248ServiceProviderPlans) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200248ServiceProviderPlans) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248ServiceProviderPlans) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200248ServiceProviderPlans) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200248ServiceProviderPlans) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse200248ServiceProviderPlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200248ServiceProviderPlans struct {
	value *InlineResponse200248ServiceProviderPlans
	isSet bool
}

func (v NullableInlineResponse200248ServiceProviderPlans) Get() *InlineResponse200248ServiceProviderPlans {
	return v.value
}

func (v *NullableInlineResponse200248ServiceProviderPlans) Set(val *InlineResponse200248ServiceProviderPlans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200248ServiceProviderPlans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200248ServiceProviderPlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200248ServiceProviderPlans(val *InlineResponse200248ServiceProviderPlans) *NullableInlineResponse200248ServiceProviderPlans {
	return &NullableInlineResponse200248ServiceProviderPlans{value: val, isSet: true}
}

func (v NullableInlineResponse200248ServiceProviderPlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200248ServiceProviderPlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


