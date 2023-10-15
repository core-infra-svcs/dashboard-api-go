/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026MajorApplications struct for InlineResponse20026MajorApplications
type InlineResponse20026MajorApplications struct {
	// Application's Meraki ID.
	Id string `json:"id"`
	// Application's name.
	Name string `json:"name"`
}

// NewInlineResponse20026MajorApplications instantiates a new InlineResponse20026MajorApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026MajorApplications(id string, name string) *InlineResponse20026MajorApplications {
	this := InlineResponse20026MajorApplications{}
	this.Id = id
	this.Name = name
	return &this
}

// NewInlineResponse20026MajorApplicationsWithDefaults instantiates a new InlineResponse20026MajorApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026MajorApplicationsWithDefaults() *InlineResponse20026MajorApplications {
	this := InlineResponse20026MajorApplications{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse20026MajorApplications) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026MajorApplications) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20026MajorApplications) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InlineResponse20026MajorApplications) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20026MajorApplications) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20026MajorApplications) SetName(v string) {
	o.Name = v
}

func (o InlineResponse20026MajorApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026MajorApplications struct {
	value *InlineResponse20026MajorApplications
	isSet bool
}

func (v NullableInlineResponse20026MajorApplications) Get() *InlineResponse20026MajorApplications {
	return v.value
}

func (v *NullableInlineResponse20026MajorApplications) Set(val *InlineResponse20026MajorApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026MajorApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026MajorApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026MajorApplications(val *InlineResponse20026MajorApplications) *NullableInlineResponse20026MajorApplications {
	return &NullableInlineResponse20026MajorApplications{value: val, isSet: true}
}

func (v NullableInlineResponse20026MajorApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026MajorApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


