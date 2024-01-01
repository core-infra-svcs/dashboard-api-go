/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200137 struct for InlineResponse200137
type InlineResponse200137 struct {
	// Switch template id
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// Switch template name
	Name *string `json:"name,omitempty"`
	// Switch model
	Model *string `json:"model,omitempty"`
}

// NewInlineResponse200137 instantiates a new InlineResponse200137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200137() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// NewInlineResponse200137WithDefaults instantiates a new InlineResponse200137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200137WithDefaults() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *InlineResponse200137) GetSwitchProfileId() string {
	if o == nil || isNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.SwitchProfileId) {
    return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *InlineResponse200137) HasSwitchProfileId() bool {
	if o != nil && !isNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *InlineResponse200137) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200137) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200137) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200137) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200137) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200137) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200137) SetModel(v string) {
	o.Model = &v
}

func (o InlineResponse200137) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200137 struct {
	value *InlineResponse200137
	isSet bool
}

func (v NullableInlineResponse200137) Get() *InlineResponse200137 {
	return v.value
}

func (v *NullableInlineResponse200137) Set(val *InlineResponse200137) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200137) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200137(val *InlineResponse200137) *NullableInlineResponse200137 {
	return &NullableInlineResponse200137{value: val, isSet: true}
}

func (v NullableInlineResponse200137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


