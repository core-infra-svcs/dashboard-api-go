/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200186 struct for InlineResponse200186
type InlineResponse200186 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Signal to noise ratio
	Snr *int32 `json:"snr,omitempty"`
	// Received signal strength indicator
	Rssi *int32 `json:"rssi,omitempty"`
}

// NewInlineResponse200186 instantiates a new InlineResponse200186 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200186() *InlineResponse200186 {
	this := InlineResponse200186{}
	return &this
}

// NewInlineResponse200186WithDefaults instantiates a new InlineResponse200186 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200186WithDefaults() *InlineResponse200186 {
	this := InlineResponse200186{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200186) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200186) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200186) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200186) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200186) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200186) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200186) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200186) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetSnr returns the Snr field value if set, zero value otherwise.
func (o *InlineResponse200186) GetSnr() int32 {
	if o == nil || isNil(o.Snr) {
		var ret int32
		return ret
	}
	return *o.Snr
}

// GetSnrOk returns a tuple with the Snr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200186) GetSnrOk() (*int32, bool) {
	if o == nil || isNil(o.Snr) {
    return nil, false
	}
	return o.Snr, true
}

// HasSnr returns a boolean if a field has been set.
func (o *InlineResponse200186) HasSnr() bool {
	if o != nil && !isNil(o.Snr) {
		return true
	}

	return false
}

// SetSnr gets a reference to the given int32 and assigns it to the Snr field.
func (o *InlineResponse200186) SetSnr(v int32) {
	o.Snr = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *InlineResponse200186) GetRssi() int32 {
	if o == nil || isNil(o.Rssi) {
		var ret int32
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200186) GetRssiOk() (*int32, bool) {
	if o == nil || isNil(o.Rssi) {
    return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *InlineResponse200186) HasRssi() bool {
	if o != nil && !isNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given int32 and assigns it to the Rssi field.
func (o *InlineResponse200186) SetRssi(v int32) {
	o.Rssi = &v
}

func (o InlineResponse200186) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.Snr) {
		toSerialize["snr"] = o.Snr
	}
	if !isNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200186 struct {
	value *InlineResponse200186
	isSet bool
}

func (v NullableInlineResponse200186) Get() *InlineResponse200186 {
	return v.value
}

func (v *NullableInlineResponse200186) Set(val *InlineResponse200186) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200186) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200186) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200186(val *InlineResponse200186) *NullableInlineResponse200186 {
	return &NullableInlineResponse200186{value: val, isSet: true}
}

func (v NullableInlineResponse200186) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200186) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


