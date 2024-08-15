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

// InlineResponse200157 struct for InlineResponse200157
type InlineResponse200157 struct {
	// ID of the Switch stack
	Id *string `json:"id,omitempty"`
	// Name of the Switch stack
	Name *string `json:"name,omitempty"`
	// Serials of the switches in the switch stack
	Serials []string `json:"serials,omitempty"`
}

// NewInlineResponse200157 instantiates a new InlineResponse200157 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200157() *InlineResponse200157 {
	this := InlineResponse200157{}
	return &this
}

// NewInlineResponse200157WithDefaults instantiates a new InlineResponse200157 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200157WithDefaults() *InlineResponse200157 {
	this := InlineResponse200157{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200157) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200157) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200157) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200157) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200157) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200157) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200157) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200157) SetName(v string) {
	o.Name = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200157) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200157) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200157) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200157) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse200157) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200157 struct {
	value *InlineResponse200157
	isSet bool
}

func (v NullableInlineResponse200157) Get() *InlineResponse200157 {
	return v.value
}

func (v *NullableInlineResponse200157) Set(val *InlineResponse200157) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200157) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200157) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200157(val *InlineResponse200157) *NullableInlineResponse200157 {
	return &NullableInlineResponse200157{value: val, isSet: true}
}

func (v NullableInlineResponse200157) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200157) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


