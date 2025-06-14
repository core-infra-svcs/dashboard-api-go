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

// OrganizationsOrganizationIdApplianceVpnStatsUsageSummary Usage summary
type OrganizationsOrganizationIdApplianceVpnStatsUsageSummary struct {
	// The amount of data received (in kilobytes)
	ReceivedInKilobytes *int32 `json:"receivedInKilobytes,omitempty"`
	// The amount of data sent (in kilobytes)
	SentInKilobytes *int32 `json:"sentInKilobytes,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceVpnStatsUsageSummary instantiates a new OrganizationsOrganizationIdApplianceVpnStatsUsageSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceVpnStatsUsageSummary() *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary {
	this := OrganizationsOrganizationIdApplianceVpnStatsUsageSummary{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceVpnStatsUsageSummaryWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnStatsUsageSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceVpnStatsUsageSummaryWithDefaults() *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary {
	this := OrganizationsOrganizationIdApplianceVpnStatsUsageSummary{}
	return &this
}

// GetReceivedInKilobytes returns the ReceivedInKilobytes field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) GetReceivedInKilobytes() int32 {
	if o == nil || isNil(o.ReceivedInKilobytes) {
		var ret int32
		return ret
	}
	return *o.ReceivedInKilobytes
}

// GetReceivedInKilobytesOk returns a tuple with the ReceivedInKilobytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) GetReceivedInKilobytesOk() (*int32, bool) {
	if o == nil || isNil(o.ReceivedInKilobytes) {
    return nil, false
	}
	return o.ReceivedInKilobytes, true
}

// HasReceivedInKilobytes returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) HasReceivedInKilobytes() bool {
	if o != nil && !isNil(o.ReceivedInKilobytes) {
		return true
	}

	return false
}

// SetReceivedInKilobytes gets a reference to the given int32 and assigns it to the ReceivedInKilobytes field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) SetReceivedInKilobytes(v int32) {
	o.ReceivedInKilobytes = &v
}

// GetSentInKilobytes returns the SentInKilobytes field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) GetSentInKilobytes() int32 {
	if o == nil || isNil(o.SentInKilobytes) {
		var ret int32
		return ret
	}
	return *o.SentInKilobytes
}

// GetSentInKilobytesOk returns a tuple with the SentInKilobytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) GetSentInKilobytesOk() (*int32, bool) {
	if o == nil || isNil(o.SentInKilobytes) {
    return nil, false
	}
	return o.SentInKilobytes, true
}

// HasSentInKilobytes returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) HasSentInKilobytes() bool {
	if o != nil && !isNil(o.SentInKilobytes) {
		return true
	}

	return false
}

// SetSentInKilobytes gets a reference to the given int32 and assigns it to the SentInKilobytes field.
func (o *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) SetSentInKilobytes(v int32) {
	o.SentInKilobytes = &v
}

func (o OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReceivedInKilobytes) {
		toSerialize["receivedInKilobytes"] = o.ReceivedInKilobytes
	}
	if !isNil(o.SentInKilobytes) {
		toSerialize["sentInKilobytes"] = o.SentInKilobytes
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary struct {
	value *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary) Get() *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary) Set(val *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary(val *OrganizationsOrganizationIdApplianceVpnStatsUsageSummary) *NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary {
	return &NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceVpnStatsUsageSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


