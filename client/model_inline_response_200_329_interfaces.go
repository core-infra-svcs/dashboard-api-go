/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200329Interfaces struct for InlineResponse200329Interfaces
type InlineResponse200329Interfaces struct {
	// The name of the wireless LAN controller interface
	Name *string `json:"name,omitempty"`
	// The MAC address of the wireless LAN controller interface
	Mac *string `json:"mac,omitempty"`
	// The statuses of layer 2 interfaces of the wireless LAN controller
	Changes []InlineResponse200329Changes `json:"changes,omitempty"`
}

// NewInlineResponse200329Interfaces instantiates a new InlineResponse200329Interfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200329Interfaces() *InlineResponse200329Interfaces {
	this := InlineResponse200329Interfaces{}
	return &this
}

// NewInlineResponse200329InterfacesWithDefaults instantiates a new InlineResponse200329Interfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200329InterfacesWithDefaults() *InlineResponse200329Interfaces {
	this := InlineResponse200329Interfaces{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200329Interfaces) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329Interfaces) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200329Interfaces) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200329Interfaces) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200329Interfaces) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329Interfaces) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200329Interfaces) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200329Interfaces) SetMac(v string) {
	o.Mac = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *InlineResponse200329Interfaces) GetChanges() []InlineResponse200329Changes {
	if o == nil || isNil(o.Changes) {
		var ret []InlineResponse200329Changes
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200329Interfaces) GetChangesOk() ([]InlineResponse200329Changes, bool) {
	if o == nil || isNil(o.Changes) {
    return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *InlineResponse200329Interfaces) HasChanges() bool {
	if o != nil && !isNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []InlineResponse200329Changes and assigns it to the Changes field.
func (o *InlineResponse200329Interfaces) SetChanges(v []InlineResponse200329Changes) {
	o.Changes = v
}

func (o InlineResponse200329Interfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Changes) {
		toSerialize["changes"] = o.Changes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200329Interfaces struct {
	value *InlineResponse200329Interfaces
	isSet bool
}

func (v NullableInlineResponse200329Interfaces) Get() *InlineResponse200329Interfaces {
	return v.value
}

func (v *NullableInlineResponse200329Interfaces) Set(val *InlineResponse200329Interfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200329Interfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200329Interfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200329Interfaces(val *InlineResponse200329Interfaces) *NullableInlineResponse200329Interfaces {
	return &NullableInlineResponse200329Interfaces{value: val, isSet: true}
}

func (v NullableInlineResponse200329Interfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200329Interfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


