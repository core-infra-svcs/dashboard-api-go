/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200131 struct for InlineResponse200131
type InlineResponse200131 struct {
	// Timestamp, in iso8601 format, at which the event happened
	Ts *time.Time `json:"ts,omitempty"`
	Device *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice `json:"device,omitempty"`
	Details *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails `json:"details,omitempty"`
	Network *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork `json:"network,omitempty"`
}

// NewInlineResponse200131 instantiates a new InlineResponse200131 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200131() *InlineResponse200131 {
	this := InlineResponse200131{}
	return &this
}

// NewInlineResponse200131WithDefaults instantiates a new InlineResponse200131 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200131WithDefaults() *InlineResponse200131 {
	this := InlineResponse200131{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200131) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200131) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200131) SetTs(v time.Time) {
	o.Ts = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200131) GetDevice() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetDeviceOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200131) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice and assigns it to the Device field.
func (o *InlineResponse200131) SetDevice(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) {
	o.Device = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse200131) GetDetails() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails {
	if o == nil || isNil(o.Details) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetDetailsOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse200131) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails and assigns it to the Details field.
func (o *InlineResponse200131) SetDetails(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) {
	o.Details = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200131) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200131) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork and assigns it to the Network field.
func (o *InlineResponse200131) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork) {
	o.Network = &v
}

func (o InlineResponse200131) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200131 struct {
	value *InlineResponse200131
	isSet bool
}

func (v NullableInlineResponse200131) Get() *InlineResponse200131 {
	return v.value
}

func (v *NullableInlineResponse200131) Set(val *InlineResponse200131) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200131) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200131) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200131(val *InlineResponse200131) *NullableInlineResponse200131 {
	return &NullableInlineResponse200131{value: val, isSet: true}
}

func (v NullableInlineResponse200131) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200131) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


