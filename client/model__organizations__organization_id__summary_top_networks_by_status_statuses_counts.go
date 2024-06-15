/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts Counts of devices by status
type OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts struct {
	// Count of online devices
	Online *int32 `json:"online,omitempty"`
	// Count of offline devices
	Offline *int32 `json:"offline,omitempty"`
	// Count of alerting devices
	Alerting *int32 `json:"alerting,omitempty"`
	// Count of dormant devices
	Dormant *int32 `json:"dormant,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCountsWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCountsWithDefaults() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts {
	this := OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts{}
	return &this
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetOnline() int32 {
	if o == nil || isNil(o.Online) {
		var ret int32
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetOnlineOk() (*int32, bool) {
	if o == nil || isNil(o.Online) {
    return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) HasOnline() bool {
	if o != nil && !isNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given int32 and assigns it to the Online field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) SetOnline(v int32) {
	o.Online = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetOffline() int32 {
	if o == nil || isNil(o.Offline) {
		var ret int32
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetOfflineOk() (*int32, bool) {
	if o == nil || isNil(o.Offline) {
    return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) HasOffline() bool {
	if o != nil && !isNil(o.Offline) {
		return true
	}

	return false
}

// SetOffline gets a reference to the given int32 and assigns it to the Offline field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) SetOffline(v int32) {
	o.Offline = &v
}

// GetAlerting returns the Alerting field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetAlerting() int32 {
	if o == nil || isNil(o.Alerting) {
		var ret int32
		return ret
	}
	return *o.Alerting
}

// GetAlertingOk returns a tuple with the Alerting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetAlertingOk() (*int32, bool) {
	if o == nil || isNil(o.Alerting) {
    return nil, false
	}
	return o.Alerting, true
}

// HasAlerting returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) HasAlerting() bool {
	if o != nil && !isNil(o.Alerting) {
		return true
	}

	return false
}

// SetAlerting gets a reference to the given int32 and assigns it to the Alerting field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) SetAlerting(v int32) {
	o.Alerting = &v
}

// GetDormant returns the Dormant field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetDormant() int32 {
	if o == nil || isNil(o.Dormant) {
		var ret int32
		return ret
	}
	return *o.Dormant
}

// GetDormantOk returns a tuple with the Dormant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) GetDormantOk() (*int32, bool) {
	if o == nil || isNil(o.Dormant) {
    return nil, false
	}
	return o.Dormant, true
}

// HasDormant returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) HasDormant() bool {
	if o != nil && !isNil(o.Dormant) {
		return true
	}

	return false
}

// SetDormant gets a reference to the given int32 and assigns it to the Dormant field.
func (o *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) SetDormant(v int32) {
	o.Dormant = &v
}

func (o OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !isNil(o.Offline) {
		toSerialize["offline"] = o.Offline
	}
	if !isNil(o.Alerting) {
		toSerialize["alerting"] = o.Alerting
	}
	if !isNil(o.Dormant) {
		toSerialize["dormant"] = o.Dormant
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts struct {
	value *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) Get() *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) Set(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts(val *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts {
	return &NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopNetworksByStatusStatusesCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


