/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200108 List of pii keys
type InlineResponse200108 struct {
	// List of mac addresses
	Macs []string `json:"macs,omitempty"`
	// List of email addresses
	Emails []string `json:"emails,omitempty"`
	// List of usernames
	Usernames []string `json:"usernames,omitempty"`
	// List of device serials
	Serials []string `json:"serials,omitempty"`
	// List of IMEIs
	Imeis []string `json:"imeis,omitempty"`
	// List of bluetooth mac addresses
	BluetoothMacs []string `json:"bluetoothMacs,omitempty"`
}

// NewInlineResponse200108 instantiates a new InlineResponse200108 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200108() *InlineResponse200108 {
	this := InlineResponse200108{}
	return &this
}

// NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200108WithDefaults() *InlineResponse200108 {
	this := InlineResponse200108{}
	return &this
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *InlineResponse200108) GetMacs() []string {
	if o == nil || isNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetMacsOk() ([]string, bool) {
	if o == nil || isNil(o.Macs) {
    return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *InlineResponse200108) HasMacs() bool {
	if o != nil && !isNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *InlineResponse200108) SetMacs(v []string) {
	o.Macs = v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *InlineResponse200108) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
    return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *InlineResponse200108) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *InlineResponse200108) SetEmails(v []string) {
	o.Emails = v
}

// GetUsernames returns the Usernames field value if set, zero value otherwise.
func (o *InlineResponse200108) GetUsernames() []string {
	if o == nil || isNil(o.Usernames) {
		var ret []string
		return ret
	}
	return o.Usernames
}

// GetUsernamesOk returns a tuple with the Usernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetUsernamesOk() ([]string, bool) {
	if o == nil || isNil(o.Usernames) {
    return nil, false
	}
	return o.Usernames, true
}

// HasUsernames returns a boolean if a field has been set.
func (o *InlineResponse200108) HasUsernames() bool {
	if o != nil && !isNil(o.Usernames) {
		return true
	}

	return false
}

// SetUsernames gets a reference to the given []string and assigns it to the Usernames field.
func (o *InlineResponse200108) SetUsernames(v []string) {
	o.Usernames = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200108) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200108) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200108) SetSerials(v []string) {
	o.Serials = v
}

// GetImeis returns the Imeis field value if set, zero value otherwise.
func (o *InlineResponse200108) GetImeis() []string {
	if o == nil || isNil(o.Imeis) {
		var ret []string
		return ret
	}
	return o.Imeis
}

// GetImeisOk returns a tuple with the Imeis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetImeisOk() ([]string, bool) {
	if o == nil || isNil(o.Imeis) {
    return nil, false
	}
	return o.Imeis, true
}

// HasImeis returns a boolean if a field has been set.
func (o *InlineResponse200108) HasImeis() bool {
	if o != nil && !isNil(o.Imeis) {
		return true
	}

	return false
}

// SetImeis gets a reference to the given []string and assigns it to the Imeis field.
func (o *InlineResponse200108) SetImeis(v []string) {
	o.Imeis = v
}

// GetBluetoothMacs returns the BluetoothMacs field value if set, zero value otherwise.
func (o *InlineResponse200108) GetBluetoothMacs() []string {
	if o == nil || isNil(o.BluetoothMacs) {
		var ret []string
		return ret
	}
	return o.BluetoothMacs
}

// GetBluetoothMacsOk returns a tuple with the BluetoothMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetBluetoothMacsOk() ([]string, bool) {
	if o == nil || isNil(o.BluetoothMacs) {
    return nil, false
	}
	return o.BluetoothMacs, true
}

// HasBluetoothMacs returns a boolean if a field has been set.
func (o *InlineResponse200108) HasBluetoothMacs() bool {
	if o != nil && !isNil(o.BluetoothMacs) {
		return true
	}

	return false
}

// SetBluetoothMacs gets a reference to the given []string and assigns it to the BluetoothMacs field.
func (o *InlineResponse200108) SetBluetoothMacs(v []string) {
	o.BluetoothMacs = v
}

func (o InlineResponse200108) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}
	if !isNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !isNil(o.Usernames) {
		toSerialize["usernames"] = o.Usernames
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.Imeis) {
		toSerialize["imeis"] = o.Imeis
	}
	if !isNil(o.BluetoothMacs) {
		toSerialize["bluetoothMacs"] = o.BluetoothMacs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200108 struct {
	value *InlineResponse200108
	isSet bool
}

func (v NullableInlineResponse200108) Get() *InlineResponse200108 {
	return v.value
}

func (v *NullableInlineResponse200108) Set(val *InlineResponse200108) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200108) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200108) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200108(val *InlineResponse200108) *NullableInlineResponse200108 {
	return &NullableInlineResponse200108{value: val, isSet: true}
}

func (v NullableInlineResponse200108) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200108) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


