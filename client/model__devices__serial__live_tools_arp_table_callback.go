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

// DevicesSerialLiveToolsArpTableCallback Details for the callback. Please include either an httpServerId OR url and sharedSecret
type DevicesSerialLiveToolsArpTableCallback struct {
	// The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
	Url *string `json:"url,omitempty"`
	// A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	HttpServer *DevicesSerialLiveToolsArpTableCallbackHttpServer `json:"httpServer,omitempty"`
	PayloadTemplate *DevicesSerialLiveToolsArpTableCallbackPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewDevicesSerialLiveToolsArpTableCallback instantiates a new DevicesSerialLiveToolsArpTableCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsArpTableCallback() *DevicesSerialLiveToolsArpTableCallback {
	this := DevicesSerialLiveToolsArpTableCallback{}
	return &this
}

// NewDevicesSerialLiveToolsArpTableCallbackWithDefaults instantiates a new DevicesSerialLiveToolsArpTableCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsArpTableCallbackWithDefaults() *DevicesSerialLiveToolsArpTableCallback {
	this := DevicesSerialLiveToolsArpTableCallback{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsArpTableCallback) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DevicesSerialLiveToolsArpTableCallback) SetUrl(v string) {
	o.Url = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsArpTableCallback) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *DevicesSerialLiveToolsArpTableCallback) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsArpTableCallback) GetHttpServer() DevicesSerialLiveToolsArpTableCallbackHttpServer {
	if o == nil || isNil(o.HttpServer) {
		var ret DevicesSerialLiveToolsArpTableCallbackHttpServer
		return ret
	}
	return *o.HttpServer
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) GetHttpServerOk() (*DevicesSerialLiveToolsArpTableCallbackHttpServer, bool) {
	if o == nil || isNil(o.HttpServer) {
    return nil, false
	}
	return o.HttpServer, true
}

// HasHttpServer returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) HasHttpServer() bool {
	if o != nil && !isNil(o.HttpServer) {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given DevicesSerialLiveToolsArpTableCallbackHttpServer and assigns it to the HttpServer field.
func (o *DevicesSerialLiveToolsArpTableCallback) SetHttpServer(v DevicesSerialLiveToolsArpTableCallbackHttpServer) {
	o.HttpServer = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsArpTableCallback) GetPayloadTemplate() DevicesSerialLiveToolsArpTableCallbackPayloadTemplate {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret DevicesSerialLiveToolsArpTableCallbackPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) GetPayloadTemplateOk() (*DevicesSerialLiveToolsArpTableCallbackPayloadTemplate, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsArpTableCallback) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given DevicesSerialLiveToolsArpTableCallbackPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *DevicesSerialLiveToolsArpTableCallback) SetPayloadTemplate(v DevicesSerialLiveToolsArpTableCallbackPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o DevicesSerialLiveToolsArpTableCallback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.HttpServer) {
		toSerialize["httpServer"] = o.HttpServer
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialLiveToolsArpTableCallback struct {
	value *DevicesSerialLiveToolsArpTableCallback
	isSet bool
}

func (v NullableDevicesSerialLiveToolsArpTableCallback) Get() *DevicesSerialLiveToolsArpTableCallback {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsArpTableCallback) Set(val *DevicesSerialLiveToolsArpTableCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsArpTableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsArpTableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsArpTableCallback(val *DevicesSerialLiveToolsArpTableCallback) *NullableDevicesSerialLiveToolsArpTableCallback {
	return &NullableDevicesSerialLiveToolsArpTableCallback{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsArpTableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsArpTableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


