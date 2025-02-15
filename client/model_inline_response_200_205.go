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

// InlineResponse200205 struct for InlineResponse200205
type InlineResponse200205 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Total usage in kilobytes-per-second
	TotalKbps *int32 `json:"totalKbps,omitempty"`
	// Sent kilobytes-per-second
	SentKbps *int32 `json:"sentKbps,omitempty"`
	// Received kilobytes-per-second
	ReceivedKbps *int32 `json:"receivedKbps,omitempty"`
}

// NewInlineResponse200205 instantiates a new InlineResponse200205 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200205() *InlineResponse200205 {
	this := InlineResponse200205{}
	return &this
}

// NewInlineResponse200205WithDefaults instantiates a new InlineResponse200205 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200205WithDefaults() *InlineResponse200205 {
	this := InlineResponse200205{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200205) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200205) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200205) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200205) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200205) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200205) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetTotalKbps returns the TotalKbps field value if set, zero value otherwise.
func (o *InlineResponse200205) GetTotalKbps() int32 {
	if o == nil || isNil(o.TotalKbps) {
		var ret int32
		return ret
	}
	return *o.TotalKbps
}

// GetTotalKbpsOk returns a tuple with the TotalKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetTotalKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalKbps) {
    return nil, false
	}
	return o.TotalKbps, true
}

// HasTotalKbps returns a boolean if a field has been set.
func (o *InlineResponse200205) HasTotalKbps() bool {
	if o != nil && !isNil(o.TotalKbps) {
		return true
	}

	return false
}

// SetTotalKbps gets a reference to the given int32 and assigns it to the TotalKbps field.
func (o *InlineResponse200205) SetTotalKbps(v int32) {
	o.TotalKbps = &v
}

// GetSentKbps returns the SentKbps field value if set, zero value otherwise.
func (o *InlineResponse200205) GetSentKbps() int32 {
	if o == nil || isNil(o.SentKbps) {
		var ret int32
		return ret
	}
	return *o.SentKbps
}

// GetSentKbpsOk returns a tuple with the SentKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetSentKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.SentKbps) {
    return nil, false
	}
	return o.SentKbps, true
}

// HasSentKbps returns a boolean if a field has been set.
func (o *InlineResponse200205) HasSentKbps() bool {
	if o != nil && !isNil(o.SentKbps) {
		return true
	}

	return false
}

// SetSentKbps gets a reference to the given int32 and assigns it to the SentKbps field.
func (o *InlineResponse200205) SetSentKbps(v int32) {
	o.SentKbps = &v
}

// GetReceivedKbps returns the ReceivedKbps field value if set, zero value otherwise.
func (o *InlineResponse200205) GetReceivedKbps() int32 {
	if o == nil || isNil(o.ReceivedKbps) {
		var ret int32
		return ret
	}
	return *o.ReceivedKbps
}

// GetReceivedKbpsOk returns a tuple with the ReceivedKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetReceivedKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.ReceivedKbps) {
    return nil, false
	}
	return o.ReceivedKbps, true
}

// HasReceivedKbps returns a boolean if a field has been set.
func (o *InlineResponse200205) HasReceivedKbps() bool {
	if o != nil && !isNil(o.ReceivedKbps) {
		return true
	}

	return false
}

// SetReceivedKbps gets a reference to the given int32 and assigns it to the ReceivedKbps field.
func (o *InlineResponse200205) SetReceivedKbps(v int32) {
	o.ReceivedKbps = &v
}

func (o InlineResponse200205) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.TotalKbps) {
		toSerialize["totalKbps"] = o.TotalKbps
	}
	if !isNil(o.SentKbps) {
		toSerialize["sentKbps"] = o.SentKbps
	}
	if !isNil(o.ReceivedKbps) {
		toSerialize["receivedKbps"] = o.ReceivedKbps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200205 struct {
	value *InlineResponse200205
	isSet bool
}

func (v NullableInlineResponse200205) Get() *InlineResponse200205 {
	return v.value
}

func (v *NullableInlineResponse200205) Set(val *InlineResponse200205) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200205) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200205) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200205(val *InlineResponse200205) *NullableInlineResponse200205 {
	return &NullableInlineResponse200205{value: val, isSet: true}
}

func (v NullableInlineResponse200205) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200205) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


