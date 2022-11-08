/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20028 struct for InlineResponse20028
type InlineResponse20028 struct {
	// The Meraki Id of the device record.
	Id *string `json:"id,omitempty"`
	// The name of the device.
	Name *string `json:"name,omitempty"`
	// An array of tags associated with the device.
	Tags []string `json:"tags,omitempty"`
	// The name of the SSID the device was last connected to.
	Ssid *string `json:"ssid,omitempty"`
	// The MAC of the device.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The name of the device OS.
	OsName *string `json:"osName,omitempty"`
	// The device model.
	SystemModel *string `json:"systemModel,omitempty"`
	// The UUID of the device.
	Uuid *string `json:"uuid,omitempty"`
	// The device serial number.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The device serial.
	Serial *string `json:"serial,omitempty"`
	// The IP address of the device.
	Ip *string `json:"ip,omitempty"`
	// Notes associated with the device.
	Notes *string `json:"notes,omitempty"`
}

// NewInlineResponse20028 instantiates a new InlineResponse20028 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20028() *InlineResponse20028 {
	this := InlineResponse20028{}
	return &this
}

// NewInlineResponse20028WithDefaults instantiates a new InlineResponse20028 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20028WithDefaults() *InlineResponse20028 {
	this := InlineResponse20028{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20028) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20028) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20028) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20028) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20028) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20028) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20028) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20028) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse20028) SetTags(v []string) {
	o.Tags = v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse20028) GetSsid() string {
	if o == nil || isNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetSsidOk() (*string, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse20028) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *InlineResponse20028) SetSsid(v string) {
	o.Ssid = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *InlineResponse20028) GetWifiMac() string {
	if o == nil || isNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetWifiMacOk() (*string, bool) {
	if o == nil || isNil(o.WifiMac) {
    return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *InlineResponse20028) HasWifiMac() bool {
	if o != nil && !isNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *InlineResponse20028) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *InlineResponse20028) GetOsName() string {
	if o == nil || isNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetOsNameOk() (*string, bool) {
	if o == nil || isNil(o.OsName) {
    return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *InlineResponse20028) HasOsName() bool {
	if o != nil && !isNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *InlineResponse20028) SetOsName(v string) {
	o.OsName = &v
}

// GetSystemModel returns the SystemModel field value if set, zero value otherwise.
func (o *InlineResponse20028) GetSystemModel() string {
	if o == nil || isNil(o.SystemModel) {
		var ret string
		return ret
	}
	return *o.SystemModel
}

// GetSystemModelOk returns a tuple with the SystemModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetSystemModelOk() (*string, bool) {
	if o == nil || isNil(o.SystemModel) {
    return nil, false
	}
	return o.SystemModel, true
}

// HasSystemModel returns a boolean if a field has been set.
func (o *InlineResponse20028) HasSystemModel() bool {
	if o != nil && !isNil(o.SystemModel) {
		return true
	}

	return false
}

// SetSystemModel gets a reference to the given string and assigns it to the SystemModel field.
func (o *InlineResponse20028) SetSystemModel(v string) {
	o.SystemModel = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InlineResponse20028) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InlineResponse20028) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InlineResponse20028) SetUuid(v string) {
	o.Uuid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *InlineResponse20028) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetSerialNumberOk() (*string, bool) {
	if o == nil || isNil(o.SerialNumber) {
    return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InlineResponse20028) HasSerialNumber() bool {
	if o != nil && !isNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *InlineResponse20028) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20028) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20028) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20028) SetSerial(v string) {
	o.Serial = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20028) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20028) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20028) SetIp(v string) {
	o.Ip = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse20028) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse20028) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse20028) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineResponse20028) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !isNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !isNil(o.SystemModel) {
		toSerialize["systemModel"] = o.SystemModel
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20028 struct {
	value *InlineResponse20028
	isSet bool
}

func (v NullableInlineResponse20028) Get() *InlineResponse20028 {
	return v.value
}

func (v *NullableInlineResponse20028) Set(val *InlineResponse20028) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20028) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20028) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20028(val *InlineResponse20028) *NullableInlineResponse20028 {
	return &NullableInlineResponse20028{value: val, isSet: true}
}

func (v NullableInlineResponse20028) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20028) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


