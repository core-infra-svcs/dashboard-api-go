/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject195 struct for InlineObject195
type InlineObject195 struct {
	// The name of the configuration template
	Name string `json:"name"`
	// The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template
	TimeZone *string `json:"timeZone,omitempty"`
	// The ID of the network or config template to copy configuration from
	CopyFromNetworkId *string `json:"copyFromNetworkId,omitempty"`
}

// NewInlineObject195 instantiates a new InlineObject195 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject195(name string) *InlineObject195 {
	this := InlineObject195{}
	this.Name = name
	return &this
}

// NewInlineObject195WithDefaults instantiates a new InlineObject195 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject195WithDefaults() *InlineObject195 {
	this := InlineObject195{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject195) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject195) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject195) SetName(v string) {
	o.Name = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject195) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject195) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject195) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject195) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCopyFromNetworkId returns the CopyFromNetworkId field value if set, zero value otherwise.
func (o *InlineObject195) GetCopyFromNetworkId() string {
	if o == nil || isNil(o.CopyFromNetworkId) {
		var ret string
		return ret
	}
	return *o.CopyFromNetworkId
}

// GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject195) GetCopyFromNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.CopyFromNetworkId) {
    return nil, false
	}
	return o.CopyFromNetworkId, true
}

// HasCopyFromNetworkId returns a boolean if a field has been set.
func (o *InlineObject195) HasCopyFromNetworkId() bool {
	if o != nil && !isNil(o.CopyFromNetworkId) {
		return true
	}

	return false
}

// SetCopyFromNetworkId gets a reference to the given string and assigns it to the CopyFromNetworkId field.
func (o *InlineObject195) SetCopyFromNetworkId(v string) {
	o.CopyFromNetworkId = &v
}

func (o InlineObject195) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.CopyFromNetworkId) {
		toSerialize["copyFromNetworkId"] = o.CopyFromNetworkId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject195 struct {
	value *InlineObject195
	isSet bool
}

func (v NullableInlineObject195) Get() *InlineObject195 {
	return v.value
}

func (v *NullableInlineObject195) Set(val *InlineObject195) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject195) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject195) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject195(val *InlineObject195) *NullableInlineObject195 {
	return &NullableInlineObject195{value: val, isSet: true}
}

func (v NullableInlineObject195) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject195) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


