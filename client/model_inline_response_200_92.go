/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20092 struct for InlineResponse20092
type InlineResponse20092 struct {
	Products *InlineResponse20092Products `json:"products,omitempty"`
	// The ordered stages in the network
	Stages []InlineResponse20092Stages `json:"stages,omitempty"`
	// Reasons for the rollback
	Reasons []InlineResponse20091Reasons `json:"reasons,omitempty"`
}

// NewInlineResponse20092 instantiates a new InlineResponse20092 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092WithDefaults() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineResponse20092) GetProducts() InlineResponse20092Products {
	if o == nil || isNil(o.Products) {
		var ret InlineResponse20092Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetProductsOk() (*InlineResponse20092Products, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineResponse20092) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given InlineResponse20092Products and assigns it to the Products field.
func (o *InlineResponse20092) SetProducts(v InlineResponse20092Products) {
	o.Products = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *InlineResponse20092) GetStages() []InlineResponse20092Stages {
	if o == nil || isNil(o.Stages) {
		var ret []InlineResponse20092Stages
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetStagesOk() ([]InlineResponse20092Stages, bool) {
	if o == nil || isNil(o.Stages) {
    return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *InlineResponse20092) HasStages() bool {
	if o != nil && !isNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []InlineResponse20092Stages and assigns it to the Stages field.
func (o *InlineResponse20092) SetStages(v []InlineResponse20092Stages) {
	o.Stages = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *InlineResponse20092) GetReasons() []InlineResponse20091Reasons {
	if o == nil || isNil(o.Reasons) {
		var ret []InlineResponse20091Reasons
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetReasonsOk() ([]InlineResponse20091Reasons, bool) {
	if o == nil || isNil(o.Reasons) {
    return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *InlineResponse20092) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []InlineResponse20091Reasons and assigns it to the Reasons field.
func (o *InlineResponse20092) SetReasons(v []InlineResponse20091Reasons) {
	o.Reasons = v
}

func (o InlineResponse20092) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !isNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	if !isNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092 struct {
	value *InlineResponse20092
	isSet bool
}

func (v NullableInlineResponse20092) Get() *InlineResponse20092 {
	return v.value
}

func (v *NullableInlineResponse20092) Set(val *InlineResponse20092) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092(val *InlineResponse20092) *NullableInlineResponse20092 {
	return &NullableInlineResponse20092{value: val, isSet: true}
}

func (v NullableInlineResponse20092) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


