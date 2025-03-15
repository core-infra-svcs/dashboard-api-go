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

// InlineObject129 struct for InlineObject129
type InlineObject129 struct {
	// ids of applications to be installed
	AppIds []string `json:"appIds"`
	// By default, installation of an app which is believed to already be present on the device will be skipped. If you'd like to force the installation of the app, set this parameter to true.
	Force *bool `json:"force,omitempty"`
}

// NewInlineObject129 instantiates a new InlineObject129 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject129(appIds []string) *InlineObject129 {
	this := InlineObject129{}
	this.AppIds = appIds
	return &this
}

// NewInlineObject129WithDefaults instantiates a new InlineObject129 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject129WithDefaults() *InlineObject129 {
	this := InlineObject129{}
	return &this
}

// GetAppIds returns the AppIds field value
func (o *InlineObject129) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject129) GetAppIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppIds, true
}

// SetAppIds sets field value
func (o *InlineObject129) SetAppIds(v []string) {
	o.AppIds = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *InlineObject129) GetForce() bool {
	if o == nil || isNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject129) GetForceOk() (*bool, bool) {
	if o == nil || isNil(o.Force) {
    return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *InlineObject129) HasForce() bool {
	if o != nil && !isNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *InlineObject129) SetForce(v bool) {
	o.Force = &v
}

func (o InlineObject129) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject129 struct {
	value *InlineObject129
	isSet bool
}

func (v NullableInlineObject129) Get() *InlineObject129 {
	return v.value
}

func (v *NullableInlineObject129) Set(val *InlineObject129) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject129) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject129) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject129(val *InlineObject129) *NullableInlineObject129 {
	return &NullableInlineObject129{value: val, isSet: true}
}

func (v NullableInlineObject129) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject129) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


