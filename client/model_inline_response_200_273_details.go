/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200273Details struct for InlineResponse200273Details
type InlineResponse200273Details struct {
	// Name of detail
	Name *string `json:"name,omitempty"`
	// Value of the named detail
	Value *string `json:"value,omitempty"`
	// Product type this set of details belongs to
	ProductType *string `json:"productType,omitempty"`
}

// NewInlineResponse200273Details instantiates a new InlineResponse200273Details object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200273Details() *InlineResponse200273Details {
	this := InlineResponse200273Details{}
	return &this
}

// NewInlineResponse200273DetailsWithDefaults instantiates a new InlineResponse200273Details object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200273DetailsWithDefaults() *InlineResponse200273Details {
	this := InlineResponse200273Details{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200273Details) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200273Details) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200273Details) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200273Details) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse200273Details) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200273Details) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse200273Details) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineResponse200273Details) SetValue(v string) {
	o.Value = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse200273Details) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200273Details) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse200273Details) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse200273Details) SetProductType(v string) {
	o.ProductType = &v
}

func (o InlineResponse200273Details) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200273Details struct {
	value *InlineResponse200273Details
	isSet bool
}

func (v NullableInlineResponse200273Details) Get() *InlineResponse200273Details {
	return v.value
}

func (v *NullableInlineResponse200273Details) Set(val *InlineResponse200273Details) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200273Details) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200273Details) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200273Details(val *InlineResponse200273Details) *NullableInlineResponse200273Details {
	return &NullableInlineResponse200273Details{value: val, isSet: true}
}

func (v NullableInlineResponse200273Details) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200273Details) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


