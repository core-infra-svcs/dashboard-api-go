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

// InlineResponse200354Interfaces struct for InlineResponse200354Interfaces
type InlineResponse200354Interfaces struct {
	// The name of the wireless LAN controller interface
	Name *string `json:"name,omitempty"`
	// The MAC address of the wireless LAN controller interface
	Mac *string `json:"mac,omitempty"`
	// The statuses of layer 3 interfaces of the wireless LAN controller
	Changes []InlineResponse200351Changes `json:"changes,omitempty"`
}

// NewInlineResponse200354Interfaces instantiates a new InlineResponse200354Interfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200354Interfaces() *InlineResponse200354Interfaces {
	this := InlineResponse200354Interfaces{}
	return &this
}

// NewInlineResponse200354InterfacesWithDefaults instantiates a new InlineResponse200354Interfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200354InterfacesWithDefaults() *InlineResponse200354Interfaces {
	this := InlineResponse200354Interfaces{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200354Interfaces) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354Interfaces) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200354Interfaces) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200354Interfaces) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200354Interfaces) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354Interfaces) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200354Interfaces) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200354Interfaces) SetMac(v string) {
	o.Mac = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *InlineResponse200354Interfaces) GetChanges() []InlineResponse200351Changes {
	if o == nil || isNil(o.Changes) {
		var ret []InlineResponse200351Changes
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200354Interfaces) GetChangesOk() ([]InlineResponse200351Changes, bool) {
	if o == nil || isNil(o.Changes) {
    return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *InlineResponse200354Interfaces) HasChanges() bool {
	if o != nil && !isNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []InlineResponse200351Changes and assigns it to the Changes field.
func (o *InlineResponse200354Interfaces) SetChanges(v []InlineResponse200351Changes) {
	o.Changes = v
}

func (o InlineResponse200354Interfaces) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200354Interfaces struct {
	value *InlineResponse200354Interfaces
	isSet bool
}

func (v NullableInlineResponse200354Interfaces) Get() *InlineResponse200354Interfaces {
	return v.value
}

func (v *NullableInlineResponse200354Interfaces) Set(val *InlineResponse200354Interfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200354Interfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200354Interfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200354Interfaces(val *InlineResponse200354Interfaces) *NullableInlineResponse200354Interfaces {
	return &NullableInlineResponse200354Interfaces{value: val, isSet: true}
}

func (v NullableInlineResponse200354Interfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200354Interfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


