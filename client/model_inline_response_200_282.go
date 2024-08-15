/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200282 struct for InlineResponse200282
type InlineResponse200282 struct {
	// Name of the Application Category
	Category *string `json:"category,omitempty"`
	// Total usage of the Application Category, in megabytes
	Total *float32 `json:"total,omitempty"`
	// Downstream usage of the Application Category, in megabytes
	Downstream *float32 `json:"downstream,omitempty"`
	// Upstream usage of the Application Category, in megabytes
	Upstream *float32 `json:"upstream,omitempty"`
	// Percent usage of the Application Category
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewInlineResponse200282 instantiates a new InlineResponse200282 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200282() *InlineResponse200282 {
	this := InlineResponse200282{}
	return &this
}

// NewInlineResponse200282WithDefaults instantiates a new InlineResponse200282 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200282WithDefaults() *InlineResponse200282 {
	this := InlineResponse200282{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse200282) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse200282) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse200282) SetCategory(v string) {
	o.Category = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200282) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200282) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *InlineResponse200282) SetTotal(v float32) {
	o.Total = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse200282) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse200282) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *InlineResponse200282) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse200282) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse200282) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *InlineResponse200282) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse200282) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse200282) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *InlineResponse200282) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o InlineResponse200282) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200282 struct {
	value *InlineResponse200282
	isSet bool
}

func (v NullableInlineResponse200282) Get() *InlineResponse200282 {
	return v.value
}

func (v *NullableInlineResponse200282) Set(val *InlineResponse200282) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200282) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200282) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200282(val *InlineResponse200282) *NullableInlineResponse200282 {
	return &NullableInlineResponse200282{value: val, isSet: true}
}

func (v NullableInlineResponse200282) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200282) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


