/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject185 struct for InlineObject185
type InlineObject185 struct {
	// The name of the combined network
	Name string `json:"name"`
	// A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
	NetworkIds []string `json:"networkIds"`
	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	EnrollmentString *string `json:"enrollmentString,omitempty"`
}

// NewInlineObject185 instantiates a new InlineObject185 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject185(name string, networkIds []string) *InlineObject185 {
	this := InlineObject185{}
	this.Name = name
	this.NetworkIds = networkIds
	return &this
}

// NewInlineObject185WithDefaults instantiates a new InlineObject185 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject185WithDefaults() *InlineObject185 {
	this := InlineObject185{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject185) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject185) SetName(v string) {
	o.Name = v
}

// GetNetworkIds returns the NetworkIds field value
func (o *InlineObject185) GetNetworkIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetNetworkIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkIds, true
}

// SetNetworkIds sets field value
func (o *InlineObject185) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *InlineObject185) GetEnrollmentString() string {
	if o == nil || o.EnrollmentString == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || o.EnrollmentString == nil {
		return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *InlineObject185) HasEnrollmentString() bool {
	if o != nil && o.EnrollmentString != nil {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *InlineObject185) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

func (o InlineObject185) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["networkIds"] = o.NetworkIds
	}
	if o.EnrollmentString != nil {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject185 struct {
	value *InlineObject185
	isSet bool
}

func (v NullableInlineObject185) Get() *InlineObject185 {
	return v.value
}

func (v *NullableInlineObject185) Set(val *InlineObject185) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject185) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject185) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject185(val *InlineObject185) *NullableInlineObject185 {
	return &NullableInlineObject185{value: val, isSet: true}
}

func (v NullableInlineObject185) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject185) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


