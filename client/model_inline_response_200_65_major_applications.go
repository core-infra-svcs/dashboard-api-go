/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20065MajorApplications struct for InlineResponse20065MajorApplications
type InlineResponse20065MajorApplications struct {
	// Application's Meraki ID.
	Id string `json:"id"`
	// Application's name.
	Name string `json:"name"`
}

// NewInlineResponse20065MajorApplications instantiates a new InlineResponse20065MajorApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20065MajorApplications(id string, name string) *InlineResponse20065MajorApplications {
	this := InlineResponse20065MajorApplications{}
	this.Id = id
	this.Name = name
	return &this
}

// NewInlineResponse20065MajorApplicationsWithDefaults instantiates a new InlineResponse20065MajorApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20065MajorApplicationsWithDefaults() *InlineResponse20065MajorApplications {
	this := InlineResponse20065MajorApplications{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse20065MajorApplications) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20065MajorApplications) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20065MajorApplications) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InlineResponse20065MajorApplications) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20065MajorApplications) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20065MajorApplications) SetName(v string) {
	o.Name = v
}

func (o InlineResponse20065MajorApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20065MajorApplications struct {
	value *InlineResponse20065MajorApplications
	isSet bool
}

func (v NullableInlineResponse20065MajorApplications) Get() *InlineResponse20065MajorApplications {
	return v.value
}

func (v *NullableInlineResponse20065MajorApplications) Set(val *InlineResponse20065MajorApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20065MajorApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20065MajorApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20065MajorApplications(val *InlineResponse20065MajorApplications) *NullableInlineResponse20065MajorApplications {
	return &NullableInlineResponse20065MajorApplications{value: val, isSet: true}
}

func (v NullableInlineResponse20065MajorApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20065MajorApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

