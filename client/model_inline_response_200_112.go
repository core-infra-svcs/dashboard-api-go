/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200112 List of pii keys
type InlineResponse200112 struct {
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

// NewInlineResponse200112 instantiates a new InlineResponse200112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200112() *InlineResponse200112 {
	this := InlineResponse200112{}
	return &this
}

// NewInlineResponse200112WithDefaults instantiates a new InlineResponse200112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200112WithDefaults() *InlineResponse200112 {
	this := InlineResponse200112{}
	return &this
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *InlineResponse200112) GetMacs() []string {
	if o == nil || isNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetMacsOk() ([]string, bool) {
	if o == nil || isNil(o.Macs) {
    return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *InlineResponse200112) HasMacs() bool {
	if o != nil && !isNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *InlineResponse200112) SetMacs(v []string) {
	o.Macs = v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *InlineResponse200112) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
    return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *InlineResponse200112) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *InlineResponse200112) SetEmails(v []string) {
	o.Emails = v
}

// GetUsernames returns the Usernames field value if set, zero value otherwise.
func (o *InlineResponse200112) GetUsernames() []string {
	if o == nil || isNil(o.Usernames) {
		var ret []string
		return ret
	}
	return o.Usernames
}

// GetUsernamesOk returns a tuple with the Usernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetUsernamesOk() ([]string, bool) {
	if o == nil || isNil(o.Usernames) {
    return nil, false
	}
	return o.Usernames, true
}

// HasUsernames returns a boolean if a field has been set.
func (o *InlineResponse200112) HasUsernames() bool {
	if o != nil && !isNil(o.Usernames) {
		return true
	}

	return false
}

// SetUsernames gets a reference to the given []string and assigns it to the Usernames field.
func (o *InlineResponse200112) SetUsernames(v []string) {
	o.Usernames = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200112) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200112) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200112) SetSerials(v []string) {
	o.Serials = v
}

// GetImeis returns the Imeis field value if set, zero value otherwise.
func (o *InlineResponse200112) GetImeis() []string {
	if o == nil || isNil(o.Imeis) {
		var ret []string
		return ret
	}
	return o.Imeis
}

// GetImeisOk returns a tuple with the Imeis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetImeisOk() ([]string, bool) {
	if o == nil || isNil(o.Imeis) {
    return nil, false
	}
	return o.Imeis, true
}

// HasImeis returns a boolean if a field has been set.
func (o *InlineResponse200112) HasImeis() bool {
	if o != nil && !isNil(o.Imeis) {
		return true
	}

	return false
}

// SetImeis gets a reference to the given []string and assigns it to the Imeis field.
func (o *InlineResponse200112) SetImeis(v []string) {
	o.Imeis = v
}

// GetBluetoothMacs returns the BluetoothMacs field value if set, zero value otherwise.
func (o *InlineResponse200112) GetBluetoothMacs() []string {
	if o == nil || isNil(o.BluetoothMacs) {
		var ret []string
		return ret
	}
	return o.BluetoothMacs
}

// GetBluetoothMacsOk returns a tuple with the BluetoothMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetBluetoothMacsOk() ([]string, bool) {
	if o == nil || isNil(o.BluetoothMacs) {
    return nil, false
	}
	return o.BluetoothMacs, true
}

// HasBluetoothMacs returns a boolean if a field has been set.
func (o *InlineResponse200112) HasBluetoothMacs() bool {
	if o != nil && !isNil(o.BluetoothMacs) {
		return true
	}

	return false
}

// SetBluetoothMacs gets a reference to the given []string and assigns it to the BluetoothMacs field.
func (o *InlineResponse200112) SetBluetoothMacs(v []string) {
	o.BluetoothMacs = v
}

func (o InlineResponse200112) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200112 struct {
	value *InlineResponse200112
	isSet bool
}

func (v NullableInlineResponse200112) Get() *InlineResponse200112 {
	return v.value
}

func (v *NullableInlineResponse200112) Set(val *InlineResponse200112) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200112) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200112(val *InlineResponse200112) *NullableInlineResponse200112 {
	return &NullableInlineResponse200112{value: val, isSet: true}
}

func (v NullableInlineResponse200112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


