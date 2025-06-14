/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApplianceVpnStatsMosSummaries struct for OrganizationsOrganizationIdApplianceVpnStatsMosSummaries
type OrganizationsOrganizationIdApplianceVpnStatsMosSummaries struct {
	// Sender uplink interface
	SenderUplink *string `json:"senderUplink,omitempty"`
	// Receiver uplink interface
	ReceiverUplink *string `json:"receiverUplink,omitempty"`
	// Average MOS (Mean opinion score)
	AvgMos *float32 `json:"avgMos,omitempty"`
	// Minimum MOS (Mean opinion score
	MinMos *float32 `json:"minMos,omitempty"`
	// Maximum MOS (Mean opinion score
	MaxMos *float32 `json:"maxMos,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceVpnStatsMosSummaries instantiates a new OrganizationsOrganizationIdApplianceVpnStatsMosSummaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceVpnStatsMosSummaries() *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries {
	this := OrganizationsOrganizationIdApplianceVpnStatsMosSummaries{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceVpnStatsMosSummariesWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnStatsMosSummaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceVpnStatsMosSummariesWithDefaults() *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries {
	this := OrganizationsOrganizationIdApplianceVpnStatsMosSummaries{}
	return &this
}

// GetSenderUplink returns the SenderUplink field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetSenderUplink() string {
	if o == nil || isNil(o.SenderUplink) {
		var ret string
		return ret
	}
	return *o.SenderUplink
}

// GetSenderUplinkOk returns a tuple with the SenderUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetSenderUplinkOk() (*string, bool) {
	if o == nil || isNil(o.SenderUplink) {
    return nil, false
	}
	return o.SenderUplink, true
}

// HasSenderUplink returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) HasSenderUplink() bool {
	if o != nil && !isNil(o.SenderUplink) {
		return true
	}

	return false
}

// SetSenderUplink gets a reference to the given string and assigns it to the SenderUplink field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) SetSenderUplink(v string) {
	o.SenderUplink = &v
}

// GetReceiverUplink returns the ReceiverUplink field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetReceiverUplink() string {
	if o == nil || isNil(o.ReceiverUplink) {
		var ret string
		return ret
	}
	return *o.ReceiverUplink
}

// GetReceiverUplinkOk returns a tuple with the ReceiverUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetReceiverUplinkOk() (*string, bool) {
	if o == nil || isNil(o.ReceiverUplink) {
    return nil, false
	}
	return o.ReceiverUplink, true
}

// HasReceiverUplink returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) HasReceiverUplink() bool {
	if o != nil && !isNil(o.ReceiverUplink) {
		return true
	}

	return false
}

// SetReceiverUplink gets a reference to the given string and assigns it to the ReceiverUplink field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) SetReceiverUplink(v string) {
	o.ReceiverUplink = &v
}

// GetAvgMos returns the AvgMos field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetAvgMos() float32 {
	if o == nil || isNil(o.AvgMos) {
		var ret float32
		return ret
	}
	return *o.AvgMos
}

// GetAvgMosOk returns a tuple with the AvgMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetAvgMosOk() (*float32, bool) {
	if o == nil || isNil(o.AvgMos) {
    return nil, false
	}
	return o.AvgMos, true
}

// HasAvgMos returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) HasAvgMos() bool {
	if o != nil && !isNil(o.AvgMos) {
		return true
	}

	return false
}

// SetAvgMos gets a reference to the given float32 and assigns it to the AvgMos field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) SetAvgMos(v float32) {
	o.AvgMos = &v
}

// GetMinMos returns the MinMos field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetMinMos() float32 {
	if o == nil || isNil(o.MinMos) {
		var ret float32
		return ret
	}
	return *o.MinMos
}

// GetMinMosOk returns a tuple with the MinMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetMinMosOk() (*float32, bool) {
	if o == nil || isNil(o.MinMos) {
    return nil, false
	}
	return o.MinMos, true
}

// HasMinMos returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) HasMinMos() bool {
	if o != nil && !isNil(o.MinMos) {
		return true
	}

	return false
}

// SetMinMos gets a reference to the given float32 and assigns it to the MinMos field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) SetMinMos(v float32) {
	o.MinMos = &v
}

// GetMaxMos returns the MaxMos field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetMaxMos() float32 {
	if o == nil || isNil(o.MaxMos) {
		var ret float32
		return ret
	}
	return *o.MaxMos
}

// GetMaxMosOk returns a tuple with the MaxMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) GetMaxMosOk() (*float32, bool) {
	if o == nil || isNil(o.MaxMos) {
    return nil, false
	}
	return o.MaxMos, true
}

// HasMaxMos returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) HasMaxMos() bool {
	if o != nil && !isNil(o.MaxMos) {
		return true
	}

	return false
}

// SetMaxMos gets a reference to the given float32 and assigns it to the MaxMos field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) SetMaxMos(v float32) {
	o.MaxMos = &v
}

func (o OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SenderUplink) {
		toSerialize["senderUplink"] = o.SenderUplink
	}
	if !isNil(o.ReceiverUplink) {
		toSerialize["receiverUplink"] = o.ReceiverUplink
	}
	if !isNil(o.AvgMos) {
		toSerialize["avgMos"] = o.AvgMos
	}
	if !isNil(o.MinMos) {
		toSerialize["minMos"] = o.MinMos
	}
	if !isNil(o.MaxMos) {
		toSerialize["maxMos"] = o.MaxMos
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries struct {
	value *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries) Get() *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries) Set(val *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries(val *OrganizationsOrganizationIdApplianceVpnStatsMosSummaries) *NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries {
	return &NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsMosSummaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


