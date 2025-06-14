/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject284 struct for InlineObject284
type InlineObject284 struct {
	// The name of the combined network
	Name string `json:"name"`
	// A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
	NetworkIds []string `json:"networkIds"`
	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	EnrollmentString *string `json:"enrollmentString,omitempty"`
}

// NewInlineObject284 instantiates a new InlineObject284 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject284(name string, networkIds []string) *InlineObject284 {
	this := InlineObject284{}
	this.Name = name
	this.NetworkIds = networkIds
	return &this
}

// NewInlineObject284WithDefaults instantiates a new InlineObject284 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject284WithDefaults() *InlineObject284 {
	this := InlineObject284{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject284) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject284) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject284) SetName(v string) {
	o.Name = v
}

// GetNetworkIds returns the NetworkIds field value
func (o *InlineObject284) GetNetworkIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject284) GetNetworkIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NetworkIds, true
}

// SetNetworkIds sets field value
func (o *InlineObject284) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *InlineObject284) GetEnrollmentString() string {
	if o == nil || isNil(o.EnrollmentString) {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject284) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || isNil(o.EnrollmentString) {
    return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *InlineObject284) HasEnrollmentString() bool {
	if o != nil && !isNil(o.EnrollmentString) {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *InlineObject284) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

func (o InlineObject284) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["networkIds"] = o.NetworkIds
	}
	if !isNil(o.EnrollmentString) {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject284 struct {
	value *InlineObject284
	isSet bool
}

func (v NullableInlineObject284) Get() *InlineObject284 {
	return v.value
}

func (v *NullableInlineObject284) Set(val *InlineObject284) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject284) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject284) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject284(val *InlineObject284) *NullableInlineObject284 {
	return &NullableInlineObject284{value: val, isSet: true}
}

func (v NullableInlineObject284) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject284) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


