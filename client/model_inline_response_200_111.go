/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200111 List of pii keys
type InlineResponse200111 struct {
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

// NewInlineResponse200111 instantiates a new InlineResponse200111 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200111() *InlineResponse200111 {
	this := InlineResponse200111{}
	return &this
}

// NewInlineResponse200111WithDefaults instantiates a new InlineResponse200111 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200111WithDefaults() *InlineResponse200111 {
	this := InlineResponse200111{}
	return &this
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *InlineResponse200111) GetMacs() []string {
	if o == nil || isNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetMacsOk() ([]string, bool) {
	if o == nil || isNil(o.Macs) {
    return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *InlineResponse200111) HasMacs() bool {
	if o != nil && !isNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *InlineResponse200111) SetMacs(v []string) {
	o.Macs = v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *InlineResponse200111) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
    return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *InlineResponse200111) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *InlineResponse200111) SetEmails(v []string) {
	o.Emails = v
}

// GetUsernames returns the Usernames field value if set, zero value otherwise.
func (o *InlineResponse200111) GetUsernames() []string {
	if o == nil || isNil(o.Usernames) {
		var ret []string
		return ret
	}
	return o.Usernames
}

// GetUsernamesOk returns a tuple with the Usernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetUsernamesOk() ([]string, bool) {
	if o == nil || isNil(o.Usernames) {
    return nil, false
	}
	return o.Usernames, true
}

// HasUsernames returns a boolean if a field has been set.
func (o *InlineResponse200111) HasUsernames() bool {
	if o != nil && !isNil(o.Usernames) {
		return true
	}

	return false
}

// SetUsernames gets a reference to the given []string and assigns it to the Usernames field.
func (o *InlineResponse200111) SetUsernames(v []string) {
	o.Usernames = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200111) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200111) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200111) SetSerials(v []string) {
	o.Serials = v
}

// GetImeis returns the Imeis field value if set, zero value otherwise.
func (o *InlineResponse200111) GetImeis() []string {
	if o == nil || isNil(o.Imeis) {
		var ret []string
		return ret
	}
	return o.Imeis
}

// GetImeisOk returns a tuple with the Imeis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetImeisOk() ([]string, bool) {
	if o == nil || isNil(o.Imeis) {
    return nil, false
	}
	return o.Imeis, true
}

// HasImeis returns a boolean if a field has been set.
func (o *InlineResponse200111) HasImeis() bool {
	if o != nil && !isNil(o.Imeis) {
		return true
	}

	return false
}

// SetImeis gets a reference to the given []string and assigns it to the Imeis field.
func (o *InlineResponse200111) SetImeis(v []string) {
	o.Imeis = v
}

// GetBluetoothMacs returns the BluetoothMacs field value if set, zero value otherwise.
func (o *InlineResponse200111) GetBluetoothMacs() []string {
	if o == nil || isNil(o.BluetoothMacs) {
		var ret []string
		return ret
	}
	return o.BluetoothMacs
}

// GetBluetoothMacsOk returns a tuple with the BluetoothMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetBluetoothMacsOk() ([]string, bool) {
	if o == nil || isNil(o.BluetoothMacs) {
    return nil, false
	}
	return o.BluetoothMacs, true
}

// HasBluetoothMacs returns a boolean if a field has been set.
func (o *InlineResponse200111) HasBluetoothMacs() bool {
	if o != nil && !isNil(o.BluetoothMacs) {
		return true
	}

	return false
}

// SetBluetoothMacs gets a reference to the given []string and assigns it to the BluetoothMacs field.
func (o *InlineResponse200111) SetBluetoothMacs(v []string) {
	o.BluetoothMacs = v
}

func (o InlineResponse200111) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200111 struct {
	value *InlineResponse200111
	isSet bool
}

func (v NullableInlineResponse200111) Get() *InlineResponse200111 {
	return v.value
}

func (v *NullableInlineResponse200111) Set(val *InlineResponse200111) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200111) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200111) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200111(val *InlineResponse200111) *NullableInlineResponse200111 {
	return &NullableInlineResponse200111{value: val, isSet: true}
}

func (v NullableInlineResponse200111) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200111) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


