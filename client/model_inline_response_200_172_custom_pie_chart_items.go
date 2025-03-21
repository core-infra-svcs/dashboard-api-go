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

// InlineResponse200172CustomPieChartItems struct for InlineResponse200172CustomPieChartItems
type InlineResponse200172CustomPieChartItems struct {
	// The name of the custom pie chart item.
	Name string `json:"name"`
	//     The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'. 
	Type string `json:"type"`
	//     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details). 
	Value string `json:"value"`
}

// NewInlineResponse200172CustomPieChartItems instantiates a new InlineResponse200172CustomPieChartItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200172CustomPieChartItems(name string, type_ string, value string) *InlineResponse200172CustomPieChartItems {
	this := InlineResponse200172CustomPieChartItems{}
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewInlineResponse200172CustomPieChartItemsWithDefaults instantiates a new InlineResponse200172CustomPieChartItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200172CustomPieChartItemsWithDefaults() *InlineResponse200172CustomPieChartItems {
	this := InlineResponse200172CustomPieChartItems{}
	return &this
}

// GetName returns the Name field value
func (o *InlineResponse200172CustomPieChartItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200172CustomPieChartItems) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse200172CustomPieChartItems) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *InlineResponse200172CustomPieChartItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200172CustomPieChartItems) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200172CustomPieChartItems) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InlineResponse200172CustomPieChartItems) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200172CustomPieChartItems) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InlineResponse200172CustomPieChartItems) SetValue(v string) {
	o.Value = v
}

func (o InlineResponse200172CustomPieChartItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200172CustomPieChartItems struct {
	value *InlineResponse200172CustomPieChartItems
	isSet bool
}

func (v NullableInlineResponse200172CustomPieChartItems) Get() *InlineResponse200172CustomPieChartItems {
	return v.value
}

func (v *NullableInlineResponse200172CustomPieChartItems) Set(val *InlineResponse200172CustomPieChartItems) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200172CustomPieChartItems) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200172CustomPieChartItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200172CustomPieChartItems(val *InlineResponse200172CustomPieChartItems) *NullableInlineResponse200172CustomPieChartItems {
	return &NullableInlineResponse200172CustomPieChartItems{value: val, isSet: true}
}

func (v NullableInlineResponse200172CustomPieChartItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200172CustomPieChartItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


