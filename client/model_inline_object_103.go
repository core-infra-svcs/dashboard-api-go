/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject103 struct for InlineObject103
type InlineObject103 struct {
	// The list of devices to publish positions for
	Devices []NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices `json:"devices,omitempty"`
}

// NewInlineObject103 instantiates a new InlineObject103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject103() *InlineObject103 {
	this := InlineObject103{}
	return &this
}

// NewInlineObject103WithDefaults instantiates a new InlineObject103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject103WithDefaults() *InlineObject103 {
	this := InlineObject103{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *InlineObject103) GetDevices() []NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices {
	if o == nil || isNil(o.Devices) {
		var ret []NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject103) GetDevicesOk() ([]NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *InlineObject103) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices and assigns it to the Devices field.
func (o *InlineObject103) SetDevices(v []NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) {
	o.Devices = v
}

func (o InlineObject103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject103 struct {
	value *InlineObject103
	isSet bool
}

func (v NullableInlineObject103) Get() *InlineObject103 {
	return v.value
}

func (v *NullableInlineObject103) Set(val *InlineObject103) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject103) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject103(val *InlineObject103) *NullableInlineObject103 {
	return &NullableInlineObject103{value: val, isSet: true}
}

func (v NullableInlineObject103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


