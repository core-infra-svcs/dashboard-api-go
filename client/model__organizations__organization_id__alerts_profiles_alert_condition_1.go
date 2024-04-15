/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAlertsProfilesAlertCondition1 The conditions that determine if the alert triggers
type OrganizationsOrganizationIdAlertsProfilesAlertCondition1 struct {
	// The total duration in seconds that the threshold should be crossed before alerting
	Duration *int32 `json:"duration,omitempty"`
	// The look back period in seconds for sensing the alert
	Window *int32 `json:"window,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.
	BitRateBps *int32 `json:"bit_rate_bps,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.
	LossRatio *float32 `json:"loss_ratio,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.
	LatencyMs *int32 `json:"latency_ms,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.
	JitterMs *int32 `json:"jitter_ms,omitempty"`
	// The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.
	Mos *float32 `json:"mos,omitempty"`
	// The uplink observed for the alert.  interface must be one of the following: wan1, wan2, wan3, cellular
	Interface *string `json:"interface,omitempty"`
}

// NewOrganizationsOrganizationIdAlertsProfilesAlertCondition1 instantiates a new OrganizationsOrganizationIdAlertsProfilesAlertCondition1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAlertsProfilesAlertCondition1() *OrganizationsOrganizationIdAlertsProfilesAlertCondition1 {
	this := OrganizationsOrganizationIdAlertsProfilesAlertCondition1{}
	return &this
}

// NewOrganizationsOrganizationIdAlertsProfilesAlertCondition1WithDefaults instantiates a new OrganizationsOrganizationIdAlertsProfilesAlertCondition1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAlertsProfilesAlertCondition1WithDefaults() *OrganizationsOrganizationIdAlertsProfilesAlertCondition1 {
	this := OrganizationsOrganizationIdAlertsProfilesAlertCondition1{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetDuration(v int32) {
	o.Duration = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetWindow() int32 {
	if o == nil || isNil(o.Window) {
		var ret int32
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetWindowOk() (*int32, bool) {
	if o == nil || isNil(o.Window) {
    return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasWindow() bool {
	if o != nil && !isNil(o.Window) {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int32 and assigns it to the Window field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetWindow(v int32) {
	o.Window = &v
}

// GetBitRateBps returns the BitRateBps field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetBitRateBps() int32 {
	if o == nil || isNil(o.BitRateBps) {
		var ret int32
		return ret
	}
	return *o.BitRateBps
}

// GetBitRateBpsOk returns a tuple with the BitRateBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetBitRateBpsOk() (*int32, bool) {
	if o == nil || isNil(o.BitRateBps) {
    return nil, false
	}
	return o.BitRateBps, true
}

// HasBitRateBps returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasBitRateBps() bool {
	if o != nil && !isNil(o.BitRateBps) {
		return true
	}

	return false
}

// SetBitRateBps gets a reference to the given int32 and assigns it to the BitRateBps field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetBitRateBps(v int32) {
	o.BitRateBps = &v
}

// GetLossRatio returns the LossRatio field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetLossRatio() float32 {
	if o == nil || isNil(o.LossRatio) {
		var ret float32
		return ret
	}
	return *o.LossRatio
}

// GetLossRatioOk returns a tuple with the LossRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetLossRatioOk() (*float32, bool) {
	if o == nil || isNil(o.LossRatio) {
    return nil, false
	}
	return o.LossRatio, true
}

// HasLossRatio returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasLossRatio() bool {
	if o != nil && !isNil(o.LossRatio) {
		return true
	}

	return false
}

// SetLossRatio gets a reference to the given float32 and assigns it to the LossRatio field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetLossRatio(v float32) {
	o.LossRatio = &v
}

// GetLatencyMs returns the LatencyMs field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetLatencyMs() int32 {
	if o == nil || isNil(o.LatencyMs) {
		var ret int32
		return ret
	}
	return *o.LatencyMs
}

// GetLatencyMsOk returns a tuple with the LatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetLatencyMsOk() (*int32, bool) {
	if o == nil || isNil(o.LatencyMs) {
    return nil, false
	}
	return o.LatencyMs, true
}

// HasLatencyMs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasLatencyMs() bool {
	if o != nil && !isNil(o.LatencyMs) {
		return true
	}

	return false
}

// SetLatencyMs gets a reference to the given int32 and assigns it to the LatencyMs field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetLatencyMs(v int32) {
	o.LatencyMs = &v
}

// GetJitterMs returns the JitterMs field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetJitterMs() int32 {
	if o == nil || isNil(o.JitterMs) {
		var ret int32
		return ret
	}
	return *o.JitterMs
}

// GetJitterMsOk returns a tuple with the JitterMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetJitterMsOk() (*int32, bool) {
	if o == nil || isNil(o.JitterMs) {
    return nil, false
	}
	return o.JitterMs, true
}

// HasJitterMs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasJitterMs() bool {
	if o != nil && !isNil(o.JitterMs) {
		return true
	}

	return false
}

// SetJitterMs gets a reference to the given int32 and assigns it to the JitterMs field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetJitterMs(v int32) {
	o.JitterMs = &v
}

// GetMos returns the Mos field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetMos() float32 {
	if o == nil || isNil(o.Mos) {
		var ret float32
		return ret
	}
	return *o.Mos
}

// GetMosOk returns a tuple with the Mos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetMosOk() (*float32, bool) {
	if o == nil || isNil(o.Mos) {
    return nil, false
	}
	return o.Mos, true
}

// HasMos returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasMos() bool {
	if o != nil && !isNil(o.Mos) {
		return true
	}

	return false
}

// SetMos gets a reference to the given float32 and assigns it to the Mos field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetMos(v float32) {
	o.Mos = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) SetInterface(v string) {
	o.Interface = &v
}

func (o OrganizationsOrganizationIdAlertsProfilesAlertCondition1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Window) {
		toSerialize["window"] = o.Window
	}
	if !isNil(o.BitRateBps) {
		toSerialize["bit_rate_bps"] = o.BitRateBps
	}
	if !isNil(o.LossRatio) {
		toSerialize["loss_ratio"] = o.LossRatio
	}
	if !isNil(o.LatencyMs) {
		toSerialize["latency_ms"] = o.LatencyMs
	}
	if !isNil(o.JitterMs) {
		toSerialize["jitter_ms"] = o.JitterMs
	}
	if !isNil(o.Mos) {
		toSerialize["mos"] = o.Mos
	}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1 struct {
	value *OrganizationsOrganizationIdAlertsProfilesAlertCondition1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1) Get() *OrganizationsOrganizationIdAlertsProfilesAlertCondition1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1) Set(val *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1(val *OrganizationsOrganizationIdAlertsProfilesAlertCondition1) *NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1 {
	return &NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAlertsProfilesAlertCondition1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


