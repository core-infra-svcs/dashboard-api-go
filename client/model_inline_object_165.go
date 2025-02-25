/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject165 struct for InlineObject165
type InlineObject165 struct {
	//     The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),     'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames). 
	Mode *string `json:"mode,omitempty"`
	// The list of items that make up the custom pie chart for traffic reporting.
	CustomPieChartItems []InlineResponse200171CustomPieChartItems `json:"customPieChartItems,omitempty"`
}

// NewInlineObject165 instantiates a new InlineObject165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject165() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// NewInlineObject165WithDefaults instantiates a new InlineObject165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject165WithDefaults() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineObject165) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineObject165) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineObject165) SetMode(v string) {
	o.Mode = &v
}

// GetCustomPieChartItems returns the CustomPieChartItems field value if set, zero value otherwise.
func (o *InlineObject165) GetCustomPieChartItems() []InlineResponse200171CustomPieChartItems {
	if o == nil || isNil(o.CustomPieChartItems) {
		var ret []InlineResponse200171CustomPieChartItems
		return ret
	}
	return o.CustomPieChartItems
}

// GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetCustomPieChartItemsOk() ([]InlineResponse200171CustomPieChartItems, bool) {
	if o == nil || isNil(o.CustomPieChartItems) {
    return nil, false
	}
	return o.CustomPieChartItems, true
}

// HasCustomPieChartItems returns a boolean if a field has been set.
func (o *InlineObject165) HasCustomPieChartItems() bool {
	if o != nil && !isNil(o.CustomPieChartItems) {
		return true
	}

	return false
}

// SetCustomPieChartItems gets a reference to the given []InlineResponse200171CustomPieChartItems and assigns it to the CustomPieChartItems field.
func (o *InlineObject165) SetCustomPieChartItems(v []InlineResponse200171CustomPieChartItems) {
	o.CustomPieChartItems = v
}

func (o InlineObject165) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.CustomPieChartItems) {
		toSerialize["customPieChartItems"] = o.CustomPieChartItems
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject165 struct {
	value *InlineObject165
	isSet bool
}

func (v NullableInlineObject165) Get() *InlineObject165 {
	return v.value
}

func (v *NullableInlineObject165) Set(val *InlineObject165) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject165) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject165(val *InlineObject165) *NullableInlineObject165 {
	return &NullableInlineObject165{value: val, isSet: true}
}

func (v NullableInlineObject165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


