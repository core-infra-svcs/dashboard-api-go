/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20029Livestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type InlineResponse20029Livestream struct {
	// An array of the related devices for the role
	RelatedDevices []InlineResponse20029LivestreamRelatedDevices `json:"relatedDevices,omitempty"`
}

// NewInlineResponse20029Livestream instantiates a new InlineResponse20029Livestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029Livestream() *InlineResponse20029Livestream {
	this := InlineResponse20029Livestream{}
	return &this
}

// NewInlineResponse20029LivestreamWithDefaults instantiates a new InlineResponse20029Livestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029LivestreamWithDefaults() *InlineResponse20029Livestream {
	this := InlineResponse20029Livestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *InlineResponse20029Livestream) GetRelatedDevices() []InlineResponse20029LivestreamRelatedDevices {
	if o == nil || isNil(o.RelatedDevices) {
		var ret []InlineResponse20029LivestreamRelatedDevices
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029Livestream) GetRelatedDevicesOk() ([]InlineResponse20029LivestreamRelatedDevices, bool) {
	if o == nil || isNil(o.RelatedDevices) {
    return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *InlineResponse20029Livestream) HasRelatedDevices() bool {
	if o != nil && !isNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []InlineResponse20029LivestreamRelatedDevices and assigns it to the RelatedDevices field.
func (o *InlineResponse20029Livestream) SetRelatedDevices(v []InlineResponse20029LivestreamRelatedDevices) {
	o.RelatedDevices = v
}

func (o InlineResponse20029Livestream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029Livestream struct {
	value *InlineResponse20029Livestream
	isSet bool
}

func (v NullableInlineResponse20029Livestream) Get() *InlineResponse20029Livestream {
	return v.value
}

func (v *NullableInlineResponse20029Livestream) Set(val *InlineResponse20029Livestream) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029Livestream) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029Livestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029Livestream(val *InlineResponse20029Livestream) *NullableInlineResponse20029Livestream {
	return &NullableInlineResponse20029Livestream{value: val, isSet: true}
}

func (v NullableInlineResponse20029Livestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029Livestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


