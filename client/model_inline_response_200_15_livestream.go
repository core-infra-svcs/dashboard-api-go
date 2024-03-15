/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20015Livestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type InlineResponse20015Livestream struct {
	// An array of the related devices for the role
	RelatedDevices []InlineResponse20015LivestreamRelatedDevices `json:"relatedDevices,omitempty"`
}

// NewInlineResponse20015Livestream instantiates a new InlineResponse20015Livestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015Livestream() *InlineResponse20015Livestream {
	this := InlineResponse20015Livestream{}
	return &this
}

// NewInlineResponse20015LivestreamWithDefaults instantiates a new InlineResponse20015Livestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015LivestreamWithDefaults() *InlineResponse20015Livestream {
	this := InlineResponse20015Livestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *InlineResponse20015Livestream) GetRelatedDevices() []InlineResponse20015LivestreamRelatedDevices {
	if o == nil || isNil(o.RelatedDevices) {
		var ret []InlineResponse20015LivestreamRelatedDevices
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015Livestream) GetRelatedDevicesOk() ([]InlineResponse20015LivestreamRelatedDevices, bool) {
	if o == nil || isNil(o.RelatedDevices) {
    return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *InlineResponse20015Livestream) HasRelatedDevices() bool {
	if o != nil && !isNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []InlineResponse20015LivestreamRelatedDevices and assigns it to the RelatedDevices field.
func (o *InlineResponse20015Livestream) SetRelatedDevices(v []InlineResponse20015LivestreamRelatedDevices) {
	o.RelatedDevices = v
}

func (o InlineResponse20015Livestream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20015Livestream struct {
	value *InlineResponse20015Livestream
	isSet bool
}

func (v NullableInlineResponse20015Livestream) Get() *InlineResponse20015Livestream {
	return v.value
}

func (v *NullableInlineResponse20015Livestream) Set(val *InlineResponse20015Livestream) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20015Livestream) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20015Livestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20015Livestream(val *InlineResponse20015Livestream) *NullableInlineResponse20015Livestream {
	return &NullableInlineResponse20015Livestream{value: val, isSet: true}
}

func (v NullableInlineResponse20015Livestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20015Livestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


