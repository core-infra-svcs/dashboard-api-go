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

// InlineResponse200109 struct for InlineResponse200109
type InlineResponse200109 struct {
	// ID of the MQTT Broker.
	MqttBrokerId *string `json:"mqttBrokerId,omitempty"`
	// Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200109 instantiates a new InlineResponse200109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200109() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// NewInlineResponse200109WithDefaults instantiates a new InlineResponse200109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200109WithDefaults() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// GetMqttBrokerId returns the MqttBrokerId field value if set, zero value otherwise.
func (o *InlineResponse200109) GetMqttBrokerId() string {
	if o == nil || isNil(o.MqttBrokerId) {
		var ret string
		return ret
	}
	return *o.MqttBrokerId
}

// GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetMqttBrokerIdOk() (*string, bool) {
	if o == nil || isNil(o.MqttBrokerId) {
    return nil, false
	}
	return o.MqttBrokerId, true
}

// HasMqttBrokerId returns a boolean if a field has been set.
func (o *InlineResponse200109) HasMqttBrokerId() bool {
	if o != nil && !isNil(o.MqttBrokerId) {
		return true
	}

	return false
}

// SetMqttBrokerId gets a reference to the given string and assigns it to the MqttBrokerId field.
func (o *InlineResponse200109) SetMqttBrokerId(v string) {
	o.MqttBrokerId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200109) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200109) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200109) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MqttBrokerId) {
		toSerialize["mqttBrokerId"] = o.MqttBrokerId
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200109 struct {
	value *InlineResponse200109
	isSet bool
}

func (v NullableInlineResponse200109) Get() *InlineResponse200109 {
	return v.value
}

func (v *NullableInlineResponse200109) Set(val *InlineResponse200109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200109(val *InlineResponse200109) *NullableInlineResponse200109 {
	return &NullableInlineResponse200109{value: val, isSet: true}
}

func (v NullableInlineResponse200109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


