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

// InlineResponse200119 struct for InlineResponse200119
type InlineResponse200119 struct {
	// The amount of cellular data received by the device.
	Received *float32 `json:"received,omitempty"`
	// The amount of cellular sent received by the device.
	Sent *float32 `json:"sent,omitempty"`
	// When the cellular usage data was collected.
	Ts *string `json:"ts,omitempty"`
}

// NewInlineResponse200119 instantiates a new InlineResponse200119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200119() *InlineResponse200119 {
	this := InlineResponse200119{}
	return &this
}

// NewInlineResponse200119WithDefaults instantiates a new InlineResponse200119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200119WithDefaults() *InlineResponse200119 {
	this := InlineResponse200119{}
	return &this
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *InlineResponse200119) GetReceived() float32 {
	if o == nil || isNil(o.Received) {
		var ret float32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetReceivedOk() (*float32, bool) {
	if o == nil || isNil(o.Received) {
    return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *InlineResponse200119) HasReceived() bool {
	if o != nil && !isNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given float32 and assigns it to the Received field.
func (o *InlineResponse200119) SetReceived(v float32) {
	o.Received = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse200119) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse200119) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *InlineResponse200119) SetSent(v float32) {
	o.Sent = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200119) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200119) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse200119) SetTs(v string) {
	o.Ts = &v
}

func (o InlineResponse200119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200119 struct {
	value *InlineResponse200119
	isSet bool
}

func (v NullableInlineResponse200119) Get() *InlineResponse200119 {
	return v.value
}

func (v *NullableInlineResponse200119) Set(val *InlineResponse200119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200119(val *InlineResponse200119) *NullableInlineResponse200119 {
	return &NullableInlineResponse200119{value: val, isSet: true}
}

func (v NullableInlineResponse200119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


