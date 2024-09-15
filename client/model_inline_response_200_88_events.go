/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20088Events struct for InlineResponse20088Events
type InlineResponse20088Events struct {
	// An UTC ISO8601 string of the time the event occurred at.
	OccurredAt *string `json:"occurredAt,omitempty"`
	// The ID of the network.
	NetworkId *string `json:"networkId,omitempty"`
	// The type of event being listed.
	Type *string `json:"type,omitempty"`
	// A description of the event the happened.
	Description *string `json:"description,omitempty"`
	// The category that the event type belongs to
	Category *string `json:"category,omitempty"`
	// A string identifying the client. This could be a client's MAC or IP address
	ClientId *string `json:"clientId,omitempty"`
	// A description of the client. This is usually the client's device name.
	ClientDescription *string `json:"clientDescription,omitempty"`
	// The client's MAC address.
	ClientMac *string `json:"clientMac,omitempty"`
	// The serial number of the device. Only shown if the device is an access point.
	DeviceSerial *string `json:"deviceSerial,omitempty"`
	// The name of the device. Only shown if the device is an access point.
	DeviceName *string `json:"deviceName,omitempty"`
	// The SSID number of the device.
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	EventData *InlineResponse20088EventData `json:"eventData,omitempty"`
}

// NewInlineResponse20088Events instantiates a new InlineResponse20088Events object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20088Events() *InlineResponse20088Events {
	this := InlineResponse20088Events{}
	return &this
}

// NewInlineResponse20088EventsWithDefaults instantiates a new InlineResponse20088Events object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20088EventsWithDefaults() *InlineResponse20088Events {
	this := InlineResponse20088Events{}
	return &this
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetOccurredAt() string {
	if o == nil || isNil(o.OccurredAt) {
		var ret string
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetOccurredAtOk() (*string, bool) {
	if o == nil || isNil(o.OccurredAt) {
    return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasOccurredAt() bool {
	if o != nil && !isNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given string and assigns it to the OccurredAt field.
func (o *InlineResponse20088Events) SetOccurredAt(v string) {
	o.OccurredAt = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20088Events) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20088Events) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20088Events) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse20088Events) SetCategory(v string) {
	o.Category = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineResponse20088Events) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientDescription returns the ClientDescription field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetClientDescription() string {
	if o == nil || isNil(o.ClientDescription) {
		var ret string
		return ret
	}
	return *o.ClientDescription
}

// GetClientDescriptionOk returns a tuple with the ClientDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetClientDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ClientDescription) {
    return nil, false
	}
	return o.ClientDescription, true
}

// HasClientDescription returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasClientDescription() bool {
	if o != nil && !isNil(o.ClientDescription) {
		return true
	}

	return false
}

// SetClientDescription gets a reference to the given string and assigns it to the ClientDescription field.
func (o *InlineResponse20088Events) SetClientDescription(v string) {
	o.ClientDescription = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetClientMac() string {
	if o == nil || isNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetClientMacOk() (*string, bool) {
	if o == nil || isNil(o.ClientMac) {
    return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasClientMac() bool {
	if o != nil && !isNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *InlineResponse20088Events) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *InlineResponse20088Events) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetDeviceName() string {
	if o == nil || isNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetDeviceNameOk() (*string, bool) {
	if o == nil || isNil(o.DeviceName) {
    return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasDeviceName() bool {
	if o != nil && !isNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *InlineResponse20088Events) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse20088Events) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetEventData returns the EventData field value if set, zero value otherwise.
func (o *InlineResponse20088Events) GetEventData() InlineResponse20088EventData {
	if o == nil || isNil(o.EventData) {
		var ret InlineResponse20088EventData
		return ret
	}
	return *o.EventData
}

// GetEventDataOk returns a tuple with the EventData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088Events) GetEventDataOk() (*InlineResponse20088EventData, bool) {
	if o == nil || isNil(o.EventData) {
    return nil, false
	}
	return o.EventData, true
}

// HasEventData returns a boolean if a field has been set.
func (o *InlineResponse20088Events) HasEventData() bool {
	if o != nil && !isNil(o.EventData) {
		return true
	}

	return false
}

// SetEventData gets a reference to the given InlineResponse20088EventData and assigns it to the EventData field.
func (o *InlineResponse20088Events) SetEventData(v InlineResponse20088EventData) {
	o.EventData = &v
}

func (o InlineResponse20088Events) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OccurredAt) {
		toSerialize["occurredAt"] = o.OccurredAt
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientDescription) {
		toSerialize["clientDescription"] = o.ClientDescription
	}
	if !isNil(o.ClientMac) {
		toSerialize["clientMac"] = o.ClientMac
	}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	if !isNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.EventData) {
		toSerialize["eventData"] = o.EventData
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20088Events struct {
	value *InlineResponse20088Events
	isSet bool
}

func (v NullableInlineResponse20088Events) Get() *InlineResponse20088Events {
	return v.value
}

func (v *NullableInlineResponse20088Events) Set(val *InlineResponse20088Events) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20088Events) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20088Events) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20088Events(val *InlineResponse20088Events) *NullableInlineResponse20088Events {
	return &NullableInlineResponse20088Events{value: val, isSet: true}
}

func (v NullableInlineResponse20088Events) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20088Events) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


