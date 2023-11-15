/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2005Livestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type InlineResponse2005Livestream struct {
	// An array of the related devices for the role
	RelatedDevices []InlineResponse2005LivestreamRelatedDevices `json:"relatedDevices,omitempty"`
}

// NewInlineResponse2005Livestream instantiates a new InlineResponse2005Livestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005Livestream() *InlineResponse2005Livestream {
	this := InlineResponse2005Livestream{}
	return &this
}

// NewInlineResponse2005LivestreamWithDefaults instantiates a new InlineResponse2005Livestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005LivestreamWithDefaults() *InlineResponse2005Livestream {
	this := InlineResponse2005Livestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *InlineResponse2005Livestream) GetRelatedDevices() []InlineResponse2005LivestreamRelatedDevices {
	if o == nil || isNil(o.RelatedDevices) {
		var ret []InlineResponse2005LivestreamRelatedDevices
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005Livestream) GetRelatedDevicesOk() ([]InlineResponse2005LivestreamRelatedDevices, bool) {
	if o == nil || isNil(o.RelatedDevices) {
    return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *InlineResponse2005Livestream) HasRelatedDevices() bool {
	if o != nil && !isNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []InlineResponse2005LivestreamRelatedDevices and assigns it to the RelatedDevices field.
func (o *InlineResponse2005Livestream) SetRelatedDevices(v []InlineResponse2005LivestreamRelatedDevices) {
	o.RelatedDevices = v
}

func (o InlineResponse2005Livestream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005Livestream struct {
	value *InlineResponse2005Livestream
	isSet bool
}

func (v NullableInlineResponse2005Livestream) Get() *InlineResponse2005Livestream {
	return v.value
}

func (v *NullableInlineResponse2005Livestream) Set(val *InlineResponse2005Livestream) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005Livestream) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005Livestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005Livestream(val *InlineResponse2005Livestream) *NullableInlineResponse2005Livestream {
	return &NullableInlineResponse2005Livestream{value: val, isSet: true}
}

func (v NullableInlineResponse2005Livestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005Livestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


