/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialLiveToolsPingCallbackHttpServer The webhook receiver used for the callback webhook.
type DevicesSerialLiveToolsPingCallbackHttpServer struct {
	// The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
	Id *string `json:"id,omitempty"`
}

// NewDevicesSerialLiveToolsPingCallbackHttpServer instantiates a new DevicesSerialLiveToolsPingCallbackHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsPingCallbackHttpServer() *DevicesSerialLiveToolsPingCallbackHttpServer {
	this := DevicesSerialLiveToolsPingCallbackHttpServer{}
	return &this
}

// NewDevicesSerialLiveToolsPingCallbackHttpServerWithDefaults instantiates a new DevicesSerialLiveToolsPingCallbackHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsPingCallbackHttpServerWithDefaults() *DevicesSerialLiveToolsPingCallbackHttpServer {
	this := DevicesSerialLiveToolsPingCallbackHttpServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsPingCallbackHttpServer) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsPingCallbackHttpServer) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsPingCallbackHttpServer) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevicesSerialLiveToolsPingCallbackHttpServer) SetId(v string) {
	o.Id = &v
}

func (o DevicesSerialLiveToolsPingCallbackHttpServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialLiveToolsPingCallbackHttpServer struct {
	value *DevicesSerialLiveToolsPingCallbackHttpServer
	isSet bool
}

func (v NullableDevicesSerialLiveToolsPingCallbackHttpServer) Get() *DevicesSerialLiveToolsPingCallbackHttpServer {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsPingCallbackHttpServer) Set(val *DevicesSerialLiveToolsPingCallbackHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsPingCallbackHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsPingCallbackHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsPingCallbackHttpServer(val *DevicesSerialLiveToolsPingCallbackHttpServer) *NullableDevicesSerialLiveToolsPingCallbackHttpServer {
	return &NullableDevicesSerialLiveToolsPingCallbackHttpServer{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsPingCallbackHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsPingCallbackHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


