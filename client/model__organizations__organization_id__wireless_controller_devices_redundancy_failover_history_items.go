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

// OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems struct for OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems
type OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems struct {
	// Wireless LAN controller cloud ID
	Serial *string `json:"serial,omitempty"`
	// Failover time
	Ts *time.Time `json:"ts,omitempty"`
	// Failover reason
	Reason *string `json:"reason,omitempty"`
	Failed *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed `json:"failed,omitempty"`
	Active *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive `json:"active,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems {
	this := OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItemsWithDefaults instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItemsWithDefaults() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems {
	this := OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetSerial(v string) {
	o.Serial = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetTs(v time.Time) {
	o.Ts = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetReason(v string) {
	o.Reason = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetFailed() OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed {
	if o == nil || isNil(o.Failed) {
		var ret OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetFailedOk() (*OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed, bool) {
	if o == nil || isNil(o.Failed) {
    return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed and assigns it to the Failed field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetFailed(v OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed) {
	o.Failed = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetActive() OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive {
	if o == nil || isNil(o.Active) {
		var ret OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetActiveOk() (*OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive and assigns it to the Active field.
func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetActive(v OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive) {
	o.Active = &v
}

func (o OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems struct {
	value *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) Get() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) Set(val *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems(val *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems {
	return &NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


