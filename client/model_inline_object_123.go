/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject123 struct for InlineObject123
type InlineObject123 struct {
	// ids of applications to be installed
	AppIds []string `json:"appIds"`
	// By default, installation of an app which is believed to already be present on the device will be skipped. If you'd like to force the installation of the app, set this parameter to true.
	Force *bool `json:"force,omitempty"`
}

// NewInlineObject123 instantiates a new InlineObject123 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject123(appIds []string) *InlineObject123 {
	this := InlineObject123{}
	this.AppIds = appIds
	return &this
}

// NewInlineObject123WithDefaults instantiates a new InlineObject123 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject123WithDefaults() *InlineObject123 {
	this := InlineObject123{}
	return &this
}

// GetAppIds returns the AppIds field value
func (o *InlineObject123) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject123) GetAppIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppIds, true
}

// SetAppIds sets field value
func (o *InlineObject123) SetAppIds(v []string) {
	o.AppIds = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *InlineObject123) GetForce() bool {
	if o == nil || isNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject123) GetForceOk() (*bool, bool) {
	if o == nil || isNil(o.Force) {
    return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *InlineObject123) HasForce() bool {
	if o != nil && !isNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *InlineObject123) SetForce(v bool) {
	o.Force = &v
}

func (o InlineObject123) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject123 struct {
	value *InlineObject123
	isSet bool
}

func (v NullableInlineObject123) Get() *InlineObject123 {
	return v.value
}

func (v *NullableInlineObject123) Set(val *InlineObject123) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject123) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject123) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject123(val *InlineObject123) *NullableInlineObject123 {
	return &NullableInlineObject123{value: val, isSet: true}
}

func (v NullableInlineObject123) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject123) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


