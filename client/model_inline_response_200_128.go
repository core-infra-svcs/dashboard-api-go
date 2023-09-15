/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200128 struct for InlineResponse200128
type InlineResponse200128 struct {
	// Switch template id
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// Switch template name
	Name *string `json:"name,omitempty"`
	// Switch model
	Model *string `json:"model,omitempty"`
}

// NewInlineResponse200128 instantiates a new InlineResponse200128 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200128() *InlineResponse200128 {
	this := InlineResponse200128{}
	return &this
}

// NewInlineResponse200128WithDefaults instantiates a new InlineResponse200128 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200128WithDefaults() *InlineResponse200128 {
	this := InlineResponse200128{}
	return &this
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *InlineResponse200128) GetSwitchProfileId() string {
	if o == nil || isNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.SwitchProfileId) {
    return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *InlineResponse200128) HasSwitchProfileId() bool {
	if o != nil && !isNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *InlineResponse200128) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200128) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200128) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200128) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200128) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200128) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200128) SetModel(v string) {
	o.Model = &v
}

func (o InlineResponse200128) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SwitchProfileId) {
		toSerialize["switchProfileId"] = o.SwitchProfileId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200128 struct {
	value *InlineResponse200128
	isSet bool
}

func (v NullableInlineResponse200128) Get() *InlineResponse200128 {
	return v.value
}

func (v *NullableInlineResponse200128) Set(val *InlineResponse200128) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200128) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200128) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200128(val *InlineResponse200128) *NullableInlineResponse200128 {
	return &NullableInlineResponse200128{value: val, isSet: true}
}

func (v NullableInlineResponse200128) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200128) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


