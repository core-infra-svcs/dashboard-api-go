/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20048 struct for InlineResponse20048
type InlineResponse20048 struct {
	// Usage received by the client on a given day
	Received *float32 `json:"received,omitempty"`
	// Usage sent by the client on a given day
	Sent *float32 `json:"sent,omitempty"`
	// The day's timestamp
	Ts *time.Time `json:"ts,omitempty"`
}

// NewInlineResponse20048 instantiates a new InlineResponse20048 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20048() *InlineResponse20048 {
	this := InlineResponse20048{}
	return &this
}

// NewInlineResponse20048WithDefaults instantiates a new InlineResponse20048 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20048WithDefaults() *InlineResponse20048 {
	this := InlineResponse20048{}
	return &this
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *InlineResponse20048) GetReceived() float32 {
	if o == nil || isNil(o.Received) {
		var ret float32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20048) GetReceivedOk() (*float32, bool) {
	if o == nil || isNil(o.Received) {
    return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *InlineResponse20048) HasReceived() bool {
	if o != nil && !isNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given float32 and assigns it to the Received field.
func (o *InlineResponse20048) SetReceived(v float32) {
	o.Received = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse20048) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20048) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse20048) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *InlineResponse20048) SetSent(v float32) {
	o.Sent = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20048) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20048) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20048) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse20048) SetTs(v time.Time) {
	o.Ts = &v
}

func (o InlineResponse20048) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20048 struct {
	value *InlineResponse20048
	isSet bool
}

func (v NullableInlineResponse20048) Get() *InlineResponse20048 {
	return v.value
}

func (v *NullableInlineResponse20048) Set(val *InlineResponse20048) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20048) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20048) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20048(val *InlineResponse20048) *NullableInlineResponse20048 {
	return &NullableInlineResponse20048{value: val, isSet: true}
}

func (v NullableInlineResponse20048) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20048) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


