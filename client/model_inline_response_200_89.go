/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20089 struct for InlineResponse20089
type InlineResponse20089 struct {
	// Usage received by the client on a given day
	Received *float32 `json:"received,omitempty"`
	// Usage sent by the client on a given day
	Sent *float32 `json:"sent,omitempty"`
	// The day's timestamp
	Ts *time.Time `json:"ts,omitempty"`
}

// NewInlineResponse20089 instantiates a new InlineResponse20089 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// NewInlineResponse20089WithDefaults instantiates a new InlineResponse20089 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089WithDefaults() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *InlineResponse20089) GetReceived() float32 {
	if o == nil || isNil(o.Received) {
		var ret float32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetReceivedOk() (*float32, bool) {
	if o == nil || isNil(o.Received) {
    return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *InlineResponse20089) HasReceived() bool {
	if o != nil && !isNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given float32 and assigns it to the Received field.
func (o *InlineResponse20089) SetReceived(v float32) {
	o.Received = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse20089) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse20089) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *InlineResponse20089) SetSent(v float32) {
	o.Sent = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20089) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20089) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse20089) SetTs(v time.Time) {
	o.Ts = &v
}

func (o InlineResponse20089) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20089 struct {
	value *InlineResponse20089
	isSet bool
}

func (v NullableInlineResponse20089) Get() *InlineResponse20089 {
	return v.value
}

func (v *NullableInlineResponse20089) Set(val *InlineResponse20089) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089(val *InlineResponse20089) *NullableInlineResponse20089 {
	return &NullableInlineResponse20089{value: val, isSet: true}
}

func (v NullableInlineResponse20089) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


