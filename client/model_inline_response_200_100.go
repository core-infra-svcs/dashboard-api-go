/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200100 struct for InlineResponse200100
type InlineResponse200100 struct {
	// Device serial
	Serial *string `json:"serial,omitempty"`
	// Device model.
	Model *string `json:"model,omitempty"`
	// Device tags.
	Tags *string `json:"tags,omitempty"`
	// Channel utilization for first wifi radio of device.
	Wifi0 []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0 `json:"wifi0,omitempty"`
	// Channel utilization for second wifi radio of device.
	Wifi1 []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0 `json:"wifi1,omitempty"`
}

// NewInlineResponse200100 instantiates a new InlineResponse200100 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200100() *InlineResponse200100 {
	this := InlineResponse200100{}
	return &this
}

// NewInlineResponse200100WithDefaults instantiates a new InlineResponse200100 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200100WithDefaults() *InlineResponse200100 {
	this := InlineResponse200100{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200100) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200100) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200100) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200100) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200100) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200100) SetModel(v string) {
	o.Model = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200100) GetTags() string {
	if o == nil || isNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetTagsOk() (*string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200100) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *InlineResponse200100) SetTags(v string) {
	o.Tags = &v
}

// GetWifi0 returns the Wifi0 field value if set, zero value otherwise.
func (o *InlineResponse200100) GetWifi0() []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0 {
	if o == nil || isNil(o.Wifi0) {
		var ret []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0
		return ret
	}
	return o.Wifi0
}

// GetWifi0Ok returns a tuple with the Wifi0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetWifi0Ok() ([]NetworksNetworkIdNetworkHealthChannelUtilizationWifi0, bool) {
	if o == nil || isNil(o.Wifi0) {
    return nil, false
	}
	return o.Wifi0, true
}

// HasWifi0 returns a boolean if a field has been set.
func (o *InlineResponse200100) HasWifi0() bool {
	if o != nil && !isNil(o.Wifi0) {
		return true
	}

	return false
}

// SetWifi0 gets a reference to the given []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0 and assigns it to the Wifi0 field.
func (o *InlineResponse200100) SetWifi0(v []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0) {
	o.Wifi0 = v
}

// GetWifi1 returns the Wifi1 field value if set, zero value otherwise.
func (o *InlineResponse200100) GetWifi1() []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0 {
	if o == nil || isNil(o.Wifi1) {
		var ret []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0
		return ret
	}
	return o.Wifi1
}

// GetWifi1Ok returns a tuple with the Wifi1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetWifi1Ok() ([]NetworksNetworkIdNetworkHealthChannelUtilizationWifi0, bool) {
	if o == nil || isNil(o.Wifi1) {
    return nil, false
	}
	return o.Wifi1, true
}

// HasWifi1 returns a boolean if a field has been set.
func (o *InlineResponse200100) HasWifi1() bool {
	if o != nil && !isNil(o.Wifi1) {
		return true
	}

	return false
}

// SetWifi1 gets a reference to the given []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0 and assigns it to the Wifi1 field.
func (o *InlineResponse200100) SetWifi1(v []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0) {
	o.Wifi1 = v
}

func (o InlineResponse200100) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Wifi0) {
		toSerialize["wifi0"] = o.Wifi0
	}
	if !isNil(o.Wifi1) {
		toSerialize["wifi1"] = o.Wifi1
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200100 struct {
	value *InlineResponse200100
	isSet bool
}

func (v NullableInlineResponse200100) Get() *InlineResponse200100 {
	return v.value
}

func (v *NullableInlineResponse200100) Set(val *InlineResponse200100) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200100) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200100) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200100(val *InlineResponse200100) *NullableInlineResponse200100 {
	return &NullableInlineResponse200100{value: val, isSet: true}
}

func (v NullableInlineResponse200100) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200100) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


