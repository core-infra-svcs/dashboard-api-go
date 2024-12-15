/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200168 struct for InlineResponse200168
type InlineResponse200168 struct {
	//     The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),     'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames). 
	Mode *string `json:"mode,omitempty"`
	// The list of items that make up the custom pie chart for traffic reporting.
	CustomPieChartItems []InlineResponse200168CustomPieChartItems `json:"customPieChartItems,omitempty"`
}

// NewInlineResponse200168 instantiates a new InlineResponse200168 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200168() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// NewInlineResponse200168WithDefaults instantiates a new InlineResponse200168 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200168WithDefaults() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse200168) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse200168) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse200168) SetMode(v string) {
	o.Mode = &v
}

// GetCustomPieChartItems returns the CustomPieChartItems field value if set, zero value otherwise.
func (o *InlineResponse200168) GetCustomPieChartItems() []InlineResponse200168CustomPieChartItems {
	if o == nil || isNil(o.CustomPieChartItems) {
		var ret []InlineResponse200168CustomPieChartItems
		return ret
	}
	return o.CustomPieChartItems
}

// GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetCustomPieChartItemsOk() ([]InlineResponse200168CustomPieChartItems, bool) {
	if o == nil || isNil(o.CustomPieChartItems) {
    return nil, false
	}
	return o.CustomPieChartItems, true
}

// HasCustomPieChartItems returns a boolean if a field has been set.
func (o *InlineResponse200168) HasCustomPieChartItems() bool {
	if o != nil && !isNil(o.CustomPieChartItems) {
		return true
	}

	return false
}

// SetCustomPieChartItems gets a reference to the given []InlineResponse200168CustomPieChartItems and assigns it to the CustomPieChartItems field.
func (o *InlineResponse200168) SetCustomPieChartItems(v []InlineResponse200168CustomPieChartItems) {
	o.CustomPieChartItems = v
}

func (o InlineResponse200168) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.CustomPieChartItems) {
		toSerialize["customPieChartItems"] = o.CustomPieChartItems
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200168 struct {
	value *InlineResponse200168
	isSet bool
}

func (v NullableInlineResponse200168) Get() *InlineResponse200168 {
	return v.value
}

func (v *NullableInlineResponse200168) Set(val *InlineResponse200168) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200168) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200168) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200168(val *InlineResponse200168) *NullableInlineResponse200168 {
	return &NullableInlineResponse200168{value: val, isSet: true}
}

func (v NullableInlineResponse200168) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200168) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


