/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200155 struct for InlineResponse200155
type InlineResponse200155 struct {
	// The device model
	Model *string `json:"model,omitempty"`
	// Total number of devices per model
	Count *int32 `json:"count,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage `json:"usage,omitempty"`
}

// NewInlineResponse200155 instantiates a new InlineResponse200155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200155() *InlineResponse200155 {
	this := InlineResponse200155{}
	return &this
}

// NewInlineResponse200155WithDefaults instantiates a new InlineResponse200155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200155WithDefaults() *InlineResponse200155 {
	this := InlineResponse200155{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200155) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200155) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200155) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200155) SetModel(v string) {
	o.Model = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200155) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200155) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200155) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200155) SetCount(v int32) {
	o.Count = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200155) GetUsage() OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage {
	if o == nil || isNil(o.Usage) {
		var ret OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200155) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200155) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse200155) SetUsage(v OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage) {
	o.Usage = &v
}

func (o InlineResponse200155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200155 struct {
	value *InlineResponse200155
	isSet bool
}

func (v NullableInlineResponse200155) Get() *InlineResponse200155 {
	return v.value
}

func (v *NullableInlineResponse200155) Set(val *InlineResponse200155) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200155) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200155(val *InlineResponse200155) *NullableInlineResponse200155 {
	return &NullableInlineResponse200155{value: val, isSet: true}
}

func (v NullableInlineResponse200155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


