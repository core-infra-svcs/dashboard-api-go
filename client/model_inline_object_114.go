/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject114 struct for InlineObject114
type InlineObject114 struct {
	// ids of applications to be installed
	AppIds []string `json:"appIds"`
	// By default, installation of an app which is believed to already be present on the device will be skipped. If you'd like to force the installation of the app, set this parameter to true.
	Force *bool `json:"force,omitempty"`
}

// NewInlineObject114 instantiates a new InlineObject114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject114(appIds []string) *InlineObject114 {
	this := InlineObject114{}
	this.AppIds = appIds
	return &this
}

// NewInlineObject114WithDefaults instantiates a new InlineObject114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject114WithDefaults() *InlineObject114 {
	this := InlineObject114{}
	return &this
}

// GetAppIds returns the AppIds field value
func (o *InlineObject114) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetAppIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppIds, true
}

// SetAppIds sets field value
func (o *InlineObject114) SetAppIds(v []string) {
	o.AppIds = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *InlineObject114) GetForce() bool {
	if o == nil || isNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetForceOk() (*bool, bool) {
	if o == nil || isNil(o.Force) {
    return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *InlineObject114) HasForce() bool {
	if o != nil && !isNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *InlineObject114) SetForce(v bool) {
	o.Force = &v
}

func (o InlineObject114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject114 struct {
	value *InlineObject114
	isSet bool
}

func (v NullableInlineObject114) Get() *InlineObject114 {
	return v.value
}

func (v *NullableInlineObject114) Set(val *InlineObject114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject114(val *InlineObject114) *NullableInlineObject114 {
	return &NullableInlineObject114{value: val, isSet: true}
}

func (v NullableInlineObject114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


