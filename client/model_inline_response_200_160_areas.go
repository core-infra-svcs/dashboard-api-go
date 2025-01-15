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

// InlineResponse200160Areas struct for InlineResponse200160Areas
type InlineResponse200160Areas struct {
	// OSPF area ID
	AreaId *string `json:"areaId,omitempty"`
	// Name of the OSPF area
	AreaName *string `json:"areaName,omitempty"`
	// Area types in OSPF. Must be one of: [\"normal\", \"stub\", \"nssa\"]
	AreaType *string `json:"areaType,omitempty"`
}

// NewInlineResponse200160Areas instantiates a new InlineResponse200160Areas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200160Areas() *InlineResponse200160Areas {
	this := InlineResponse200160Areas{}
	return &this
}

// NewInlineResponse200160AreasWithDefaults instantiates a new InlineResponse200160Areas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200160AreasWithDefaults() *InlineResponse200160Areas {
	this := InlineResponse200160Areas{}
	return &this
}

// GetAreaId returns the AreaId field value if set, zero value otherwise.
func (o *InlineResponse200160Areas) GetAreaId() string {
	if o == nil || isNil(o.AreaId) {
		var ret string
		return ret
	}
	return *o.AreaId
}

// GetAreaIdOk returns a tuple with the AreaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160Areas) GetAreaIdOk() (*string, bool) {
	if o == nil || isNil(o.AreaId) {
    return nil, false
	}
	return o.AreaId, true
}

// HasAreaId returns a boolean if a field has been set.
func (o *InlineResponse200160Areas) HasAreaId() bool {
	if o != nil && !isNil(o.AreaId) {
		return true
	}

	return false
}

// SetAreaId gets a reference to the given string and assigns it to the AreaId field.
func (o *InlineResponse200160Areas) SetAreaId(v string) {
	o.AreaId = &v
}

// GetAreaName returns the AreaName field value if set, zero value otherwise.
func (o *InlineResponse200160Areas) GetAreaName() string {
	if o == nil || isNil(o.AreaName) {
		var ret string
		return ret
	}
	return *o.AreaName
}

// GetAreaNameOk returns a tuple with the AreaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160Areas) GetAreaNameOk() (*string, bool) {
	if o == nil || isNil(o.AreaName) {
    return nil, false
	}
	return o.AreaName, true
}

// HasAreaName returns a boolean if a field has been set.
func (o *InlineResponse200160Areas) HasAreaName() bool {
	if o != nil && !isNil(o.AreaName) {
		return true
	}

	return false
}

// SetAreaName gets a reference to the given string and assigns it to the AreaName field.
func (o *InlineResponse200160Areas) SetAreaName(v string) {
	o.AreaName = &v
}

// GetAreaType returns the AreaType field value if set, zero value otherwise.
func (o *InlineResponse200160Areas) GetAreaType() string {
	if o == nil || isNil(o.AreaType) {
		var ret string
		return ret
	}
	return *o.AreaType
}

// GetAreaTypeOk returns a tuple with the AreaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200160Areas) GetAreaTypeOk() (*string, bool) {
	if o == nil || isNil(o.AreaType) {
    return nil, false
	}
	return o.AreaType, true
}

// HasAreaType returns a boolean if a field has been set.
func (o *InlineResponse200160Areas) HasAreaType() bool {
	if o != nil && !isNil(o.AreaType) {
		return true
	}

	return false
}

// SetAreaType gets a reference to the given string and assigns it to the AreaType field.
func (o *InlineResponse200160Areas) SetAreaType(v string) {
	o.AreaType = &v
}

func (o InlineResponse200160Areas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AreaId) {
		toSerialize["areaId"] = o.AreaId
	}
	if !isNil(o.AreaName) {
		toSerialize["areaName"] = o.AreaName
	}
	if !isNil(o.AreaType) {
		toSerialize["areaType"] = o.AreaType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200160Areas struct {
	value *InlineResponse200160Areas
	isSet bool
}

func (v NullableInlineResponse200160Areas) Get() *InlineResponse200160Areas {
	return v.value
}

func (v *NullableInlineResponse200160Areas) Set(val *InlineResponse200160Areas) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200160Areas) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200160Areas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200160Areas(val *InlineResponse200160Areas) *NullableInlineResponse200160Areas {
	return &NullableInlineResponse200160Areas{value: val, isSet: true}
}

func (v NullableInlineResponse200160Areas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200160Areas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


