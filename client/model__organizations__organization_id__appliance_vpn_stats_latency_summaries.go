/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries struct for OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries
type OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries struct {
	// Sender uplink interface
	SenderUplink *string `json:"senderUplink,omitempty"`
	// Receiver uplink interface
	ReceiverUplink *string `json:"receiverUplink,omitempty"`
	// Average latency in milliseconds
	AvgLatencyMs *int32 `json:"avgLatencyMs,omitempty"`
	// Minimum latency in milliseconds
	MinLatencyMs *int32 `json:"minLatencyMs,omitempty"`
	// Maximum latency in milliseconds
	MaxLatencyMs *int32 `json:"maxLatencyMs,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries instantiates a new OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries() *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries {
	this := OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceVpnStatsLatencySummariesWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceVpnStatsLatencySummariesWithDefaults() *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries {
	this := OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries{}
	return &this
}

// GetSenderUplink returns the SenderUplink field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetSenderUplink() string {
	if o == nil || isNil(o.SenderUplink) {
		var ret string
		return ret
	}
	return *o.SenderUplink
}

// GetSenderUplinkOk returns a tuple with the SenderUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetSenderUplinkOk() (*string, bool) {
	if o == nil || isNil(o.SenderUplink) {
    return nil, false
	}
	return o.SenderUplink, true
}

// HasSenderUplink returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) HasSenderUplink() bool {
	if o != nil && !isNil(o.SenderUplink) {
		return true
	}

	return false
}

// SetSenderUplink gets a reference to the given string and assigns it to the SenderUplink field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) SetSenderUplink(v string) {
	o.SenderUplink = &v
}

// GetReceiverUplink returns the ReceiverUplink field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetReceiverUplink() string {
	if o == nil || isNil(o.ReceiverUplink) {
		var ret string
		return ret
	}
	return *o.ReceiverUplink
}

// GetReceiverUplinkOk returns a tuple with the ReceiverUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetReceiverUplinkOk() (*string, bool) {
	if o == nil || isNil(o.ReceiverUplink) {
    return nil, false
	}
	return o.ReceiverUplink, true
}

// HasReceiverUplink returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) HasReceiverUplink() bool {
	if o != nil && !isNil(o.ReceiverUplink) {
		return true
	}

	return false
}

// SetReceiverUplink gets a reference to the given string and assigns it to the ReceiverUplink field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) SetReceiverUplink(v string) {
	o.ReceiverUplink = &v
}

// GetAvgLatencyMs returns the AvgLatencyMs field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetAvgLatencyMs() int32 {
	if o == nil || isNil(o.AvgLatencyMs) {
		var ret int32
		return ret
	}
	return *o.AvgLatencyMs
}

// GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetAvgLatencyMsOk() (*int32, bool) {
	if o == nil || isNil(o.AvgLatencyMs) {
    return nil, false
	}
	return o.AvgLatencyMs, true
}

// HasAvgLatencyMs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) HasAvgLatencyMs() bool {
	if o != nil && !isNil(o.AvgLatencyMs) {
		return true
	}

	return false
}

// SetAvgLatencyMs gets a reference to the given int32 and assigns it to the AvgLatencyMs field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) SetAvgLatencyMs(v int32) {
	o.AvgLatencyMs = &v
}

// GetMinLatencyMs returns the MinLatencyMs field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetMinLatencyMs() int32 {
	if o == nil || isNil(o.MinLatencyMs) {
		var ret int32
		return ret
	}
	return *o.MinLatencyMs
}

// GetMinLatencyMsOk returns a tuple with the MinLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetMinLatencyMsOk() (*int32, bool) {
	if o == nil || isNil(o.MinLatencyMs) {
    return nil, false
	}
	return o.MinLatencyMs, true
}

// HasMinLatencyMs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) HasMinLatencyMs() bool {
	if o != nil && !isNil(o.MinLatencyMs) {
		return true
	}

	return false
}

// SetMinLatencyMs gets a reference to the given int32 and assigns it to the MinLatencyMs field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) SetMinLatencyMs(v int32) {
	o.MinLatencyMs = &v
}

// GetMaxLatencyMs returns the MaxLatencyMs field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetMaxLatencyMs() int32 {
	if o == nil || isNil(o.MaxLatencyMs) {
		var ret int32
		return ret
	}
	return *o.MaxLatencyMs
}

// GetMaxLatencyMsOk returns a tuple with the MaxLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) GetMaxLatencyMsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLatencyMs) {
    return nil, false
	}
	return o.MaxLatencyMs, true
}

// HasMaxLatencyMs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) HasMaxLatencyMs() bool {
	if o != nil && !isNil(o.MaxLatencyMs) {
		return true
	}

	return false
}

// SetMaxLatencyMs gets a reference to the given int32 and assigns it to the MaxLatencyMs field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) SetMaxLatencyMs(v int32) {
	o.MaxLatencyMs = &v
}

func (o OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SenderUplink) {
		toSerialize["senderUplink"] = o.SenderUplink
	}
	if !isNil(o.ReceiverUplink) {
		toSerialize["receiverUplink"] = o.ReceiverUplink
	}
	if !isNil(o.AvgLatencyMs) {
		toSerialize["avgLatencyMs"] = o.AvgLatencyMs
	}
	if !isNil(o.MinLatencyMs) {
		toSerialize["minLatencyMs"] = o.MinLatencyMs
	}
	if !isNil(o.MaxLatencyMs) {
		toSerialize["maxLatencyMs"] = o.MaxLatencyMs
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries struct {
	value *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) Get() *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) Set(val *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries(val *OrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) *NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries {
	return &NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsLatencySummaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


