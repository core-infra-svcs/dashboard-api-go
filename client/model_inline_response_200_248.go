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

// InlineResponse200248 struct for InlineResponse200248
type InlineResponse200248 struct {
	// Timestamp, in iso8601 format, at which the event happened
	Ts *time.Time `json:"ts,omitempty"`
	Device *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice `json:"device,omitempty"`
	Details *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails `json:"details,omitempty"`
	Network *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork `json:"network,omitempty"`
}

// NewInlineResponse200248 instantiates a new InlineResponse200248 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200248() *InlineResponse200248 {
	this := InlineResponse200248{}
	return &this
}

// NewInlineResponse200248WithDefaults instantiates a new InlineResponse200248 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200248WithDefaults() *InlineResponse200248 {
	this := InlineResponse200248{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200248) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200248) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200248) SetTs(v time.Time) {
	o.Ts = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200248) GetDevice() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248) GetDeviceOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200248) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice and assigns it to the Device field.
func (o *InlineResponse200248) SetDevice(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice) {
	o.Device = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse200248) GetDetails() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails {
	if o == nil || isNil(o.Details) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248) GetDetailsOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse200248) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails and assigns it to the Details field.
func (o *InlineResponse200248) SetDetails(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) {
	o.Details = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200248) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200248) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200248) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork and assigns it to the Network field.
func (o *InlineResponse200248) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork) {
	o.Network = &v
}

func (o InlineResponse200248) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200248 struct {
	value *InlineResponse200248
	isSet bool
}

func (v NullableInlineResponse200248) Get() *InlineResponse200248 {
	return v.value
}

func (v *NullableInlineResponse200248) Set(val *InlineResponse200248) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200248) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200248) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200248(val *InlineResponse200248) *NullableInlineResponse200248 {
	return &NullableInlineResponse200248{value: val, isSet: true}
}

func (v NullableInlineResponse200248) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200248) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


