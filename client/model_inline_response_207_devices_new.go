/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse207DevicesNew The device that will have settings cloned to
type InlineResponse207DevicesNew struct {
	// MAC address of the device
	Mac string `json:"mac"`
	// Serial number of device
	Serial string `json:"serial"`
	// Model name for device
	Model string `json:"model"`
	// Customized name for device, or MAC address
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse207DevicesNew instantiates a new InlineResponse207DevicesNew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207DevicesNew(mac string, serial string, model string) *InlineResponse207DevicesNew {
	this := InlineResponse207DevicesNew{}
	this.Mac = mac
	this.Serial = serial
	this.Model = model
	return &this
}

// NewInlineResponse207DevicesNewWithDefaults instantiates a new InlineResponse207DevicesNew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207DevicesNewWithDefaults() *InlineResponse207DevicesNew {
	this := InlineResponse207DevicesNew{}
	return &this
}

// GetMac returns the Mac field value
func (o *InlineResponse207DevicesNew) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207DevicesNew) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *InlineResponse207DevicesNew) SetMac(v string) {
	o.Mac = v
}

// GetSerial returns the Serial field value
func (o *InlineResponse207DevicesNew) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207DevicesNew) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineResponse207DevicesNew) SetSerial(v string) {
	o.Serial = v
}

// GetModel returns the Model field value
func (o *InlineResponse207DevicesNew) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207DevicesNew) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *InlineResponse207DevicesNew) SetModel(v string) {
	o.Model = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse207DevicesNew) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207DevicesNew) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse207DevicesNew) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse207DevicesNew) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse207DevicesNew) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse207DevicesNew struct {
	value *InlineResponse207DevicesNew
	isSet bool
}

func (v NullableInlineResponse207DevicesNew) Get() *InlineResponse207DevicesNew {
	return v.value
}

func (v *NullableInlineResponse207DevicesNew) Set(val *InlineResponse207DevicesNew) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse207DevicesNew) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse207DevicesNew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse207DevicesNew(val *InlineResponse207DevicesNew) *NullableInlineResponse207DevicesNew {
	return &NullableInlineResponse207DevicesNew{value: val, isSet: true}
}

func (v NullableInlineResponse207DevicesNew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse207DevicesNew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


