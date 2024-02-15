/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20011Livestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type InlineResponse20011Livestream struct {
	// An array of the related devices for the role
	RelatedDevices []InlineResponse20011LivestreamRelatedDevices `json:"relatedDevices,omitempty"`
}

// NewInlineResponse20011Livestream instantiates a new InlineResponse20011Livestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011Livestream() *InlineResponse20011Livestream {
	this := InlineResponse20011Livestream{}
	return &this
}

// NewInlineResponse20011LivestreamWithDefaults instantiates a new InlineResponse20011Livestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011LivestreamWithDefaults() *InlineResponse20011Livestream {
	this := InlineResponse20011Livestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *InlineResponse20011Livestream) GetRelatedDevices() []InlineResponse20011LivestreamRelatedDevices {
	if o == nil || isNil(o.RelatedDevices) {
		var ret []InlineResponse20011LivestreamRelatedDevices
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011Livestream) GetRelatedDevicesOk() ([]InlineResponse20011LivestreamRelatedDevices, bool) {
	if o == nil || isNil(o.RelatedDevices) {
    return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *InlineResponse20011Livestream) HasRelatedDevices() bool {
	if o != nil && !isNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []InlineResponse20011LivestreamRelatedDevices and assigns it to the RelatedDevices field.
func (o *InlineResponse20011Livestream) SetRelatedDevices(v []InlineResponse20011LivestreamRelatedDevices) {
	o.RelatedDevices = v
}

func (o InlineResponse20011Livestream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011Livestream struct {
	value *InlineResponse20011Livestream
	isSet bool
}

func (v NullableInlineResponse20011Livestream) Get() *InlineResponse20011Livestream {
	return v.value
}

func (v *NullableInlineResponse20011Livestream) Set(val *InlineResponse20011Livestream) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011Livestream) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011Livestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011Livestream(val *InlineResponse20011Livestream) *NullableInlineResponse20011Livestream {
	return &NullableInlineResponse20011Livestream{value: val, isSet: true}
}

func (v NullableInlineResponse20011Livestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011Livestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


