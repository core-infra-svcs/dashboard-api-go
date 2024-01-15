/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20057 struct for InlineResponse20057
type InlineResponse20057 struct {
	// The Meraki Id of the device record.
	Id *string `json:"id,omitempty"`
	// The name of the device.
	Name *string `json:"name,omitempty"`
	// The MAC of the device.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The device serial.
	Serial *string `json:"serial,omitempty"`
	// Notes associated with the device.
	Notes *string `json:"notes,omitempty"`
}

// NewInlineResponse20057 instantiates a new InlineResponse20057 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20057() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// NewInlineResponse20057WithDefaults instantiates a new InlineResponse20057 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20057WithDefaults() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20057) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20057) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20057) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20057) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20057) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20057) SetName(v string) {
	o.Name = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *InlineResponse20057) GetWifiMac() string {
	if o == nil || isNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetWifiMacOk() (*string, bool) {
	if o == nil || isNil(o.WifiMac) {
    return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *InlineResponse20057) HasWifiMac() bool {
	if o != nil && !isNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *InlineResponse20057) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20057) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20057) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20057) SetSerial(v string) {
	o.Serial = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse20057) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse20057) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse20057) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineResponse20057) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20057 struct {
	value *InlineResponse20057
	isSet bool
}

func (v NullableInlineResponse20057) Get() *InlineResponse20057 {
	return v.value
}

func (v *NullableInlineResponse20057) Set(val *InlineResponse20057) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20057) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20057) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20057(val *InlineResponse20057) *NullableInlineResponse20057 {
	return &NullableInlineResponse20057{value: val, isSet: true}
}

func (v NullableInlineResponse20057) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20057) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


