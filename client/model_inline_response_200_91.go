/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20091 struct for InlineResponse20091
type InlineResponse20091 struct {
	Products *InlineResponse20091Products `json:"products,omitempty"`
	// The ordered stages in the network
	Stages []InlineResponse20091Stages `json:"stages,omitempty"`
	// Reasons for the rollback
	Reasons []InlineResponse20090Reasons `json:"reasons,omitempty"`
}

// NewInlineResponse20091 instantiates a new InlineResponse20091 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// NewInlineResponse20091WithDefaults instantiates a new InlineResponse20091 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091WithDefaults() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineResponse20091) GetProducts() InlineResponse20091Products {
	if o == nil || isNil(o.Products) {
		var ret InlineResponse20091Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetProductsOk() (*InlineResponse20091Products, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineResponse20091) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given InlineResponse20091Products and assigns it to the Products field.
func (o *InlineResponse20091) SetProducts(v InlineResponse20091Products) {
	o.Products = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *InlineResponse20091) GetStages() []InlineResponse20091Stages {
	if o == nil || isNil(o.Stages) {
		var ret []InlineResponse20091Stages
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetStagesOk() ([]InlineResponse20091Stages, bool) {
	if o == nil || isNil(o.Stages) {
    return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *InlineResponse20091) HasStages() bool {
	if o != nil && !isNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []InlineResponse20091Stages and assigns it to the Stages field.
func (o *InlineResponse20091) SetStages(v []InlineResponse20091Stages) {
	o.Stages = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *InlineResponse20091) GetReasons() []InlineResponse20090Reasons {
	if o == nil || isNil(o.Reasons) {
		var ret []InlineResponse20090Reasons
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetReasonsOk() ([]InlineResponse20090Reasons, bool) {
	if o == nil || isNil(o.Reasons) {
    return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *InlineResponse20091) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []InlineResponse20090Reasons and assigns it to the Reasons field.
func (o *InlineResponse20091) SetReasons(v []InlineResponse20090Reasons) {
	o.Reasons = v
}

func (o InlineResponse20091) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20091 struct {
	value *InlineResponse20091
	isSet bool
}

func (v NullableInlineResponse20091) Get() *InlineResponse20091 {
	return v.value
}

func (v *NullableInlineResponse20091) Set(val *InlineResponse20091) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091(val *InlineResponse20091) *NullableInlineResponse20091 {
	return &NullableInlineResponse20091{value: val, isSet: true}
}

func (v NullableInlineResponse20091) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


