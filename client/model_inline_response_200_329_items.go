/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200329Items struct for InlineResponse200329Items
type InlineResponse200329Items struct {
	// Wireless LAN controller cloud ID
	Serial *string `json:"serial,omitempty"`
	// Connectivity information of a wireless LAN controller
	Changes []InlineResponse200329Changes `json:"changes,omitempty"`
}

// NewInlineResponse200329Items instantiates a new InlineResponse200329Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200329Items() *InlineResponse200329Items {
	this := InlineResponse200329Items{}
	return &this
}

// NewInlineResponse200329ItemsWithDefaults instantiates a new InlineResponse200329Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200329ItemsWithDefaults() *InlineResponse200329Items {
	this := InlineResponse200329Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200329Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200329Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200329Items) SetSerial(v string) {
	o.Serial = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *InlineResponse200329Items) GetChanges() []InlineResponse200329Changes {
	if o == nil || isNil(o.Changes) {
		var ret []InlineResponse200329Changes
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329Items) GetChangesOk() ([]InlineResponse200329Changes, bool) {
	if o == nil || isNil(o.Changes) {
    return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *InlineResponse200329Items) HasChanges() bool {
	if o != nil && !isNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []InlineResponse200329Changes and assigns it to the Changes field.
func (o *InlineResponse200329Items) SetChanges(v []InlineResponse200329Changes) {
	o.Changes = v
}

func (o InlineResponse200329Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Changes) {
		toSerialize["changes"] = o.Changes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200329Items struct {
	value *InlineResponse200329Items
	isSet bool
}

func (v NullableInlineResponse200329Items) Get() *InlineResponse200329Items {
	return v.value
}

func (v *NullableInlineResponse200329Items) Set(val *InlineResponse200329Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200329Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200329Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200329Items(val *InlineResponse200329Items) *NullableInlineResponse200329Items {
	return &NullableInlineResponse200329Items{value: val, isSet: true}
}

func (v NullableInlineResponse200329Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200329Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


