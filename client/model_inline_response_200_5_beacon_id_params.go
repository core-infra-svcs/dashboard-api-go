/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2005BeaconIdParams Beacon Id parameters with an identifier and major and minor versions
type InlineResponse2005BeaconIdParams struct {
	// The UUID to be used in the beacon identifier
	Uuid *string `json:"uuid,omitempty"`
	// The major number to be used in the beacon identifier
	Major *int32 `json:"major,omitempty"`
	// The minor number to be used in the beacon identifier
	Minor *int32 `json:"minor,omitempty"`
}

// NewInlineResponse2005BeaconIdParams instantiates a new InlineResponse2005BeaconIdParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005BeaconIdParams() *InlineResponse2005BeaconIdParams {
	this := InlineResponse2005BeaconIdParams{}
	return &this
}

// NewInlineResponse2005BeaconIdParamsWithDefaults instantiates a new InlineResponse2005BeaconIdParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005BeaconIdParamsWithDefaults() *InlineResponse2005BeaconIdParams {
	this := InlineResponse2005BeaconIdParams{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InlineResponse2005BeaconIdParams) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005BeaconIdParams) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InlineResponse2005BeaconIdParams) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InlineResponse2005BeaconIdParams) SetUuid(v string) {
	o.Uuid = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *InlineResponse2005BeaconIdParams) GetMajor() int32 {
	if o == nil || isNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005BeaconIdParams) GetMajorOk() (*int32, bool) {
	if o == nil || isNil(o.Major) {
    return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *InlineResponse2005BeaconIdParams) HasMajor() bool {
	if o != nil && !isNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *InlineResponse2005BeaconIdParams) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *InlineResponse2005BeaconIdParams) GetMinor() int32 {
	if o == nil || isNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005BeaconIdParams) GetMinorOk() (*int32, bool) {
	if o == nil || isNil(o.Minor) {
    return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *InlineResponse2005BeaconIdParams) HasMinor() bool {
	if o != nil && !isNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *InlineResponse2005BeaconIdParams) SetMinor(v int32) {
	o.Minor = &v
}

func (o InlineResponse2005BeaconIdParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !isNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005BeaconIdParams struct {
	value *InlineResponse2005BeaconIdParams
	isSet bool
}

func (v NullableInlineResponse2005BeaconIdParams) Get() *InlineResponse2005BeaconIdParams {
	return v.value
}

func (v *NullableInlineResponse2005BeaconIdParams) Set(val *InlineResponse2005BeaconIdParams) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005BeaconIdParams) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005BeaconIdParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005BeaconIdParams(val *InlineResponse2005BeaconIdParams) *NullableInlineResponse2005BeaconIdParams {
	return &NullableInlineResponse2005BeaconIdParams{value: val, isSet: true}
}

func (v NullableInlineResponse2005BeaconIdParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005BeaconIdParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


