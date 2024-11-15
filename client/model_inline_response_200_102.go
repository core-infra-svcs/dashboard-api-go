/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200102 struct for InlineResponse200102
type InlineResponse200102 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// WAN goodput (Number of useful information bits delivered over a WAN per unit of time)
	WanGoodput *int32 `json:"wanGoodput,omitempty"`
	// LAN goodput (Number of useful information bits delivered over a LAN per unit of time)
	LanGoodput *int32 `json:"lanGoodput,omitempty"`
	// WAN latency in milliseconds
	WanLatencyMs *float32 `json:"wanLatencyMs,omitempty"`
	// LAN latency in milliseconds
	LanLatencyMs *float32 `json:"lanLatencyMs,omitempty"`
	// WAN loss percentage
	WanLossPercent *float32 `json:"wanLossPercent,omitempty"`
	// LAN loss percentage
	LanLossPercent *float32 `json:"lanLossPercent,omitempty"`
	// Duration of the response, in milliseconds
	ResponseDuration *int32 `json:"responseDuration,omitempty"`
	// Sent kilobytes-per-second
	Sent *int32 `json:"sent,omitempty"`
	// Received kilobytes-per-second
	Recv *int32 `json:"recv,omitempty"`
	// Number of clients
	NumClients *int32 `json:"numClients,omitempty"`
}

// NewInlineResponse200102 instantiates a new InlineResponse200102 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200102() *InlineResponse200102 {
	this := InlineResponse200102{}
	return &this
}

// NewInlineResponse200102WithDefaults instantiates a new InlineResponse200102 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200102WithDefaults() *InlineResponse200102 {
	this := InlineResponse200102{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200102) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200102) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200102) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200102) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200102) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200102) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetWanGoodput returns the WanGoodput field value if set, zero value otherwise.
func (o *InlineResponse200102) GetWanGoodput() int32 {
	if o == nil || isNil(o.WanGoodput) {
		var ret int32
		return ret
	}
	return *o.WanGoodput
}

// GetWanGoodputOk returns a tuple with the WanGoodput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetWanGoodputOk() (*int32, bool) {
	if o == nil || isNil(o.WanGoodput) {
    return nil, false
	}
	return o.WanGoodput, true
}

// HasWanGoodput returns a boolean if a field has been set.
func (o *InlineResponse200102) HasWanGoodput() bool {
	if o != nil && !isNil(o.WanGoodput) {
		return true
	}

	return false
}

// SetWanGoodput gets a reference to the given int32 and assigns it to the WanGoodput field.
func (o *InlineResponse200102) SetWanGoodput(v int32) {
	o.WanGoodput = &v
}

// GetLanGoodput returns the LanGoodput field value if set, zero value otherwise.
func (o *InlineResponse200102) GetLanGoodput() int32 {
	if o == nil || isNil(o.LanGoodput) {
		var ret int32
		return ret
	}
	return *o.LanGoodput
}

// GetLanGoodputOk returns a tuple with the LanGoodput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetLanGoodputOk() (*int32, bool) {
	if o == nil || isNil(o.LanGoodput) {
    return nil, false
	}
	return o.LanGoodput, true
}

// HasLanGoodput returns a boolean if a field has been set.
func (o *InlineResponse200102) HasLanGoodput() bool {
	if o != nil && !isNil(o.LanGoodput) {
		return true
	}

	return false
}

// SetLanGoodput gets a reference to the given int32 and assigns it to the LanGoodput field.
func (o *InlineResponse200102) SetLanGoodput(v int32) {
	o.LanGoodput = &v
}

// GetWanLatencyMs returns the WanLatencyMs field value if set, zero value otherwise.
func (o *InlineResponse200102) GetWanLatencyMs() float32 {
	if o == nil || isNil(o.WanLatencyMs) {
		var ret float32
		return ret
	}
	return *o.WanLatencyMs
}

// GetWanLatencyMsOk returns a tuple with the WanLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetWanLatencyMsOk() (*float32, bool) {
	if o == nil || isNil(o.WanLatencyMs) {
    return nil, false
	}
	return o.WanLatencyMs, true
}

// HasWanLatencyMs returns a boolean if a field has been set.
func (o *InlineResponse200102) HasWanLatencyMs() bool {
	if o != nil && !isNil(o.WanLatencyMs) {
		return true
	}

	return false
}

// SetWanLatencyMs gets a reference to the given float32 and assigns it to the WanLatencyMs field.
func (o *InlineResponse200102) SetWanLatencyMs(v float32) {
	o.WanLatencyMs = &v
}

// GetLanLatencyMs returns the LanLatencyMs field value if set, zero value otherwise.
func (o *InlineResponse200102) GetLanLatencyMs() float32 {
	if o == nil || isNil(o.LanLatencyMs) {
		var ret float32
		return ret
	}
	return *o.LanLatencyMs
}

// GetLanLatencyMsOk returns a tuple with the LanLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetLanLatencyMsOk() (*float32, bool) {
	if o == nil || isNil(o.LanLatencyMs) {
    return nil, false
	}
	return o.LanLatencyMs, true
}

// HasLanLatencyMs returns a boolean if a field has been set.
func (o *InlineResponse200102) HasLanLatencyMs() bool {
	if o != nil && !isNil(o.LanLatencyMs) {
		return true
	}

	return false
}

// SetLanLatencyMs gets a reference to the given float32 and assigns it to the LanLatencyMs field.
func (o *InlineResponse200102) SetLanLatencyMs(v float32) {
	o.LanLatencyMs = &v
}

// GetWanLossPercent returns the WanLossPercent field value if set, zero value otherwise.
func (o *InlineResponse200102) GetWanLossPercent() float32 {
	if o == nil || isNil(o.WanLossPercent) {
		var ret float32
		return ret
	}
	return *o.WanLossPercent
}

// GetWanLossPercentOk returns a tuple with the WanLossPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetWanLossPercentOk() (*float32, bool) {
	if o == nil || isNil(o.WanLossPercent) {
    return nil, false
	}
	return o.WanLossPercent, true
}

// HasWanLossPercent returns a boolean if a field has been set.
func (o *InlineResponse200102) HasWanLossPercent() bool {
	if o != nil && !isNil(o.WanLossPercent) {
		return true
	}

	return false
}

// SetWanLossPercent gets a reference to the given float32 and assigns it to the WanLossPercent field.
func (o *InlineResponse200102) SetWanLossPercent(v float32) {
	o.WanLossPercent = &v
}

// GetLanLossPercent returns the LanLossPercent field value if set, zero value otherwise.
func (o *InlineResponse200102) GetLanLossPercent() float32 {
	if o == nil || isNil(o.LanLossPercent) {
		var ret float32
		return ret
	}
	return *o.LanLossPercent
}

// GetLanLossPercentOk returns a tuple with the LanLossPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetLanLossPercentOk() (*float32, bool) {
	if o == nil || isNil(o.LanLossPercent) {
    return nil, false
	}
	return o.LanLossPercent, true
}

// HasLanLossPercent returns a boolean if a field has been set.
func (o *InlineResponse200102) HasLanLossPercent() bool {
	if o != nil && !isNil(o.LanLossPercent) {
		return true
	}

	return false
}

// SetLanLossPercent gets a reference to the given float32 and assigns it to the LanLossPercent field.
func (o *InlineResponse200102) SetLanLossPercent(v float32) {
	o.LanLossPercent = &v
}

// GetResponseDuration returns the ResponseDuration field value if set, zero value otherwise.
func (o *InlineResponse200102) GetResponseDuration() int32 {
	if o == nil || isNil(o.ResponseDuration) {
		var ret int32
		return ret
	}
	return *o.ResponseDuration
}

// GetResponseDurationOk returns a tuple with the ResponseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetResponseDurationOk() (*int32, bool) {
	if o == nil || isNil(o.ResponseDuration) {
    return nil, false
	}
	return o.ResponseDuration, true
}

// HasResponseDuration returns a boolean if a field has been set.
func (o *InlineResponse200102) HasResponseDuration() bool {
	if o != nil && !isNil(o.ResponseDuration) {
		return true
	}

	return false
}

// SetResponseDuration gets a reference to the given int32 and assigns it to the ResponseDuration field.
func (o *InlineResponse200102) SetResponseDuration(v int32) {
	o.ResponseDuration = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse200102) GetSent() int32 {
	if o == nil || isNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetSentOk() (*int32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse200102) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *InlineResponse200102) SetSent(v int32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *InlineResponse200102) GetRecv() int32 {
	if o == nil || isNil(o.Recv) {
		var ret int32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetRecvOk() (*int32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *InlineResponse200102) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given int32 and assigns it to the Recv field.
func (o *InlineResponse200102) SetRecv(v int32) {
	o.Recv = &v
}

// GetNumClients returns the NumClients field value if set, zero value otherwise.
func (o *InlineResponse200102) GetNumClients() int32 {
	if o == nil || isNil(o.NumClients) {
		var ret int32
		return ret
	}
	return *o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102) GetNumClientsOk() (*int32, bool) {
	if o == nil || isNil(o.NumClients) {
    return nil, false
	}
	return o.NumClients, true
}

// HasNumClients returns a boolean if a field has been set.
func (o *InlineResponse200102) HasNumClients() bool {
	if o != nil && !isNil(o.NumClients) {
		return true
	}

	return false
}

// SetNumClients gets a reference to the given int32 and assigns it to the NumClients field.
func (o *InlineResponse200102) SetNumClients(v int32) {
	o.NumClients = &v
}

func (o InlineResponse200102) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.WanGoodput) {
		toSerialize["wanGoodput"] = o.WanGoodput
	}
	if !isNil(o.LanGoodput) {
		toSerialize["lanGoodput"] = o.LanGoodput
	}
	if !isNil(o.WanLatencyMs) {
		toSerialize["wanLatencyMs"] = o.WanLatencyMs
	}
	if !isNil(o.LanLatencyMs) {
		toSerialize["lanLatencyMs"] = o.LanLatencyMs
	}
	if !isNil(o.WanLossPercent) {
		toSerialize["wanLossPercent"] = o.WanLossPercent
	}
	if !isNil(o.LanLossPercent) {
		toSerialize["lanLossPercent"] = o.LanLossPercent
	}
	if !isNil(o.ResponseDuration) {
		toSerialize["responseDuration"] = o.ResponseDuration
	}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	if !isNil(o.NumClients) {
		toSerialize["numClients"] = o.NumClients
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200102 struct {
	value *InlineResponse200102
	isSet bool
}

func (v NullableInlineResponse200102) Get() *InlineResponse200102 {
	return v.value
}

func (v *NullableInlineResponse200102) Set(val *InlineResponse200102) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200102) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200102) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200102(val *InlineResponse200102) *NullableInlineResponse200102 {
	return &NullableInlineResponse200102{value: val, isSet: true}
}

func (v NullableInlineResponse200102) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200102) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


