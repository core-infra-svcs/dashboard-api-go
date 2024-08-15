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

// InlineResponse20066MajorApplications struct for InlineResponse20066MajorApplications
type InlineResponse20066MajorApplications struct {
	// Application's Meraki ID.
	Id string `json:"id"`
	// Application's name.
	Name string `json:"name"`
}

// NewInlineResponse20066MajorApplications instantiates a new InlineResponse20066MajorApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066MajorApplications(id string, name string) *InlineResponse20066MajorApplications {
	this := InlineResponse20066MajorApplications{}
	this.Id = id
	this.Name = name
	return &this
}

// NewInlineResponse20066MajorApplicationsWithDefaults instantiates a new InlineResponse20066MajorApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066MajorApplicationsWithDefaults() *InlineResponse20066MajorApplications {
	this := InlineResponse20066MajorApplications{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse20066MajorApplications) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066MajorApplications) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20066MajorApplications) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InlineResponse20066MajorApplications) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066MajorApplications) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20066MajorApplications) SetName(v string) {
	o.Name = v
}

func (o InlineResponse20066MajorApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20066MajorApplications struct {
	value *InlineResponse20066MajorApplications
	isSet bool
}

func (v NullableInlineResponse20066MajorApplications) Get() *InlineResponse20066MajorApplications {
	return v.value
}

func (v *NullableInlineResponse20066MajorApplications) Set(val *InlineResponse20066MajorApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066MajorApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066MajorApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066MajorApplications(val *InlineResponse20066MajorApplications) *NullableInlineResponse20066MajorApplications {
	return &NullableInlineResponse20066MajorApplications{value: val, isSet: true}
}

func (v NullableInlineResponse20066MajorApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066MajorApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


