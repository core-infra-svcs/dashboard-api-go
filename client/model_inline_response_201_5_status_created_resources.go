/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2015StatusCreatedResources struct for InlineResponse2015StatusCreatedResources
type InlineResponse2015StatusCreatedResources struct {
	// ID of the created resource
	Id *string `json:"id,omitempty"`
	// URI, not including base, of the created resource
	Uri *string `json:"uri,omitempty"`
}

// NewInlineResponse2015StatusCreatedResources instantiates a new InlineResponse2015StatusCreatedResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2015StatusCreatedResources() *InlineResponse2015StatusCreatedResources {
	this := InlineResponse2015StatusCreatedResources{}
	return &this
}

// NewInlineResponse2015StatusCreatedResourcesWithDefaults instantiates a new InlineResponse2015StatusCreatedResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2015StatusCreatedResourcesWithDefaults() *InlineResponse2015StatusCreatedResources {
	this := InlineResponse2015StatusCreatedResources{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2015StatusCreatedResources) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015StatusCreatedResources) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2015StatusCreatedResources) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2015StatusCreatedResources) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineResponse2015StatusCreatedResources) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015StatusCreatedResources) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
    return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineResponse2015StatusCreatedResources) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineResponse2015StatusCreatedResources) SetUri(v string) {
	o.Uri = &v
}

func (o InlineResponse2015StatusCreatedResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2015StatusCreatedResources struct {
	value *InlineResponse2015StatusCreatedResources
	isSet bool
}

func (v NullableInlineResponse2015StatusCreatedResources) Get() *InlineResponse2015StatusCreatedResources {
	return v.value
}

func (v *NullableInlineResponse2015StatusCreatedResources) Set(val *InlineResponse2015StatusCreatedResources) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2015StatusCreatedResources) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2015StatusCreatedResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2015StatusCreatedResources(val *InlineResponse2015StatusCreatedResources) *NullableInlineResponse2015StatusCreatedResources {
	return &NullableInlineResponse2015StatusCreatedResources{value: val, isSet: true}
}

func (v NullableInlineResponse2015StatusCreatedResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2015StatusCreatedResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


