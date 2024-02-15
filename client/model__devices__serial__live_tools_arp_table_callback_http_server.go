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

// DevicesSerialLiveToolsArpTableCallbackHttpServer The webhook receiver used for the callback webhook.
type DevicesSerialLiveToolsArpTableCallbackHttpServer struct {
	// The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.
	Id *string `json:"id,omitempty"`
}

// NewDevicesSerialLiveToolsArpTableCallbackHttpServer instantiates a new DevicesSerialLiveToolsArpTableCallbackHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsArpTableCallbackHttpServer() *DevicesSerialLiveToolsArpTableCallbackHttpServer {
	this := DevicesSerialLiveToolsArpTableCallbackHttpServer{}
	return &this
}

// NewDevicesSerialLiveToolsArpTableCallbackHttpServerWithDefaults instantiates a new DevicesSerialLiveToolsArpTableCallbackHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsArpTableCallbackHttpServerWithDefaults() *DevicesSerialLiveToolsArpTableCallbackHttpServer {
	this := DevicesSerialLiveToolsArpTableCallbackHttpServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsArpTableCallbackHttpServer) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsArpTableCallbackHttpServer) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsArpTableCallbackHttpServer) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevicesSerialLiveToolsArpTableCallbackHttpServer) SetId(v string) {
	o.Id = &v
}

func (o DevicesSerialLiveToolsArpTableCallbackHttpServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialLiveToolsArpTableCallbackHttpServer struct {
	value *DevicesSerialLiveToolsArpTableCallbackHttpServer
	isSet bool
}

func (v NullableDevicesSerialLiveToolsArpTableCallbackHttpServer) Get() *DevicesSerialLiveToolsArpTableCallbackHttpServer {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsArpTableCallbackHttpServer) Set(val *DevicesSerialLiveToolsArpTableCallbackHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsArpTableCallbackHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsArpTableCallbackHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsArpTableCallbackHttpServer(val *DevicesSerialLiveToolsArpTableCallbackHttpServer) *NullableDevicesSerialLiveToolsArpTableCallbackHttpServer {
	return &NullableDevicesSerialLiveToolsArpTableCallbackHttpServer{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsArpTableCallbackHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsArpTableCallbackHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


