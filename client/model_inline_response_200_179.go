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

// InlineResponse200179 struct for InlineResponse200179
type InlineResponse200179 struct {
	// Timestamp at which the event occurred
	OccurredAt *time.Time `json:"occurredAt,omitempty"`
	// Wireless band the event occurred on
	Band *string `json:"band,omitempty"`
	// Number of the SSID the event occurred in
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// Event type
	Type *string `json:"type,omitempty"`
	// Event subtype
	Subtype *string `json:"subtype,omitempty"`
	// Event severity
	Severity *string `json:"severity,omitempty"`
	// Duration of the event in milliseconds
	DurationMs *int32 `json:"durationMs,omitempty"`
	// Wireless channel the event occurred over
	Channel *int32 `json:"channel,omitempty"`
	// RSSI recorded at the time of the event
	Rssi *int32 `json:"rssi,omitempty"`
	// Additional information relevant to the given event. Properties vary based on event type.
	EventData map[string]interface{} `json:"eventData,omitempty"`
	// Serial number of the device the event occurred for
	DeviceSerial *string `json:"deviceSerial,omitempty"`
	// Id of the packet capture triggered for the event, if any
	CaptureId *string `json:"captureId,omitempty"`
}

// NewInlineResponse200179 instantiates a new InlineResponse200179 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200179() *InlineResponse200179 {
	this := InlineResponse200179{}
	return &this
}

// NewInlineResponse200179WithDefaults instantiates a new InlineResponse200179 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200179WithDefaults() *InlineResponse200179 {
	this := InlineResponse200179{}
	return &this
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *InlineResponse200179) GetOccurredAt() time.Time {
	if o == nil || isNil(o.OccurredAt) {
		var ret time.Time
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.OccurredAt) {
    return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *InlineResponse200179) HasOccurredAt() bool {
	if o != nil && !isNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given time.Time and assigns it to the OccurredAt field.
func (o *InlineResponse200179) SetOccurredAt(v time.Time) {
	o.OccurredAt = &v
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *InlineResponse200179) GetBand() string {
	if o == nil || isNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetBandOk() (*string, bool) {
	if o == nil || isNil(o.Band) {
    return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *InlineResponse200179) HasBand() bool {
	if o != nil && !isNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *InlineResponse200179) SetBand(v string) {
	o.Band = &v
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse200179) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse200179) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse200179) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200179) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200179) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200179) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *InlineResponse200179) GetSubtype() string {
	if o == nil || isNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetSubtypeOk() (*string, bool) {
	if o == nil || isNil(o.Subtype) {
    return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *InlineResponse200179) HasSubtype() bool {
	if o != nil && !isNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *InlineResponse200179) SetSubtype(v string) {
	o.Subtype = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InlineResponse200179) GetSeverity() string {
	if o == nil || isNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetSeverityOk() (*string, bool) {
	if o == nil || isNil(o.Severity) {
    return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InlineResponse200179) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *InlineResponse200179) SetSeverity(v string) {
	o.Severity = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *InlineResponse200179) GetDurationMs() int32 {
	if o == nil || isNil(o.DurationMs) {
		var ret int32
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetDurationMsOk() (*int32, bool) {
	if o == nil || isNil(o.DurationMs) {
    return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *InlineResponse200179) HasDurationMs() bool {
	if o != nil && !isNil(o.DurationMs) {
		return true
	}

	return false
}

// SetDurationMs gets a reference to the given int32 and assigns it to the DurationMs field.
func (o *InlineResponse200179) SetDurationMs(v int32) {
	o.DurationMs = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineResponse200179) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineResponse200179) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *InlineResponse200179) SetChannel(v int32) {
	o.Channel = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *InlineResponse200179) GetRssi() int32 {
	if o == nil || isNil(o.Rssi) {
		var ret int32
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetRssiOk() (*int32, bool) {
	if o == nil || isNil(o.Rssi) {
    return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *InlineResponse200179) HasRssi() bool {
	if o != nil && !isNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given int32 and assigns it to the Rssi field.
func (o *InlineResponse200179) SetRssi(v int32) {
	o.Rssi = &v
}

// GetEventData returns the EventData field value if set, zero value otherwise.
func (o *InlineResponse200179) GetEventData() map[string]interface{} {
	if o == nil || isNil(o.EventData) {
		var ret map[string]interface{}
		return ret
	}
	return o.EventData
}

// GetEventDataOk returns a tuple with the EventData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetEventDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.EventData) {
    return map[string]interface{}{}, false
	}
	return o.EventData, true
}

// HasEventData returns a boolean if a field has been set.
func (o *InlineResponse200179) HasEventData() bool {
	if o != nil && !isNil(o.EventData) {
		return true
	}

	return false
}

// SetEventData gets a reference to the given map[string]interface{} and assigns it to the EventData field.
func (o *InlineResponse200179) SetEventData(v map[string]interface{}) {
	o.EventData = v
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *InlineResponse200179) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *InlineResponse200179) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *InlineResponse200179) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

// GetCaptureId returns the CaptureId field value if set, zero value otherwise.
func (o *InlineResponse200179) GetCaptureId() string {
	if o == nil || isNil(o.CaptureId) {
		var ret string
		return ret
	}
	return *o.CaptureId
}

// GetCaptureIdOk returns a tuple with the CaptureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200179) GetCaptureIdOk() (*string, bool) {
	if o == nil || isNil(o.CaptureId) {
    return nil, false
	}
	return o.CaptureId, true
}

// HasCaptureId returns a boolean if a field has been set.
func (o *InlineResponse200179) HasCaptureId() bool {
	if o != nil && !isNil(o.CaptureId) {
		return true
	}

	return false
}

// SetCaptureId gets a reference to the given string and assigns it to the CaptureId field.
func (o *InlineResponse200179) SetCaptureId(v string) {
	o.CaptureId = &v
}

func (o InlineResponse200179) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OccurredAt) {
		toSerialize["occurredAt"] = o.OccurredAt
	}
	if !isNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !isNil(o.DurationMs) {
		toSerialize["durationMs"] = o.DurationMs
	}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	if !isNil(o.EventData) {
		toSerialize["eventData"] = o.EventData
	}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	if !isNil(o.CaptureId) {
		toSerialize["captureId"] = o.CaptureId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200179 struct {
	value *InlineResponse200179
	isSet bool
}

func (v NullableInlineResponse200179) Get() *InlineResponse200179 {
	return v.value
}

func (v *NullableInlineResponse200179) Set(val *InlineResponse200179) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200179) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200179) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200179(val *InlineResponse200179) *NullableInlineResponse200179 {
	return &NullableInlineResponse200179{value: val, isSet: true}
}

func (v NullableInlineResponse200179) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200179) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


