/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200137 struct for InlineResponse200137
type InlineResponse200137 struct {
	// User name
	Name *string `json:"name,omitempty"`
	// User login identifier
	Login *string `json:"login,omitempty"`
	// SSID name
	Ssid *string `json:"ssid,omitempty"`
	// Login timestamp
	LoginAt *time.Time `json:"loginAt,omitempty"`
	// Gateway device mac address
	GatewayDeviceMac *string `json:"gatewayDeviceMac,omitempty"`
	// Client mac address
	ClientMac *string `json:"clientMac,omitempty"`
	// Client ID
	ClientId *string `json:"clientId,omitempty"`
	// Authorization status
	Authorization *string `json:"authorization,omitempty"`
}

// NewInlineResponse200137 instantiates a new InlineResponse200137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200137() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// NewInlineResponse200137WithDefaults instantiates a new InlineResponse200137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200137WithDefaults() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200137) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200137) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200137) SetName(v string) {
	o.Name = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *InlineResponse200137) GetLogin() string {
	if o == nil || isNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetLoginOk() (*string, bool) {
	if o == nil || isNil(o.Login) {
    return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *InlineResponse200137) HasLogin() bool {
	if o != nil && !isNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *InlineResponse200137) SetLogin(v string) {
	o.Login = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *InlineResponse200137) GetSsid() string {
	if o == nil || isNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetSsidOk() (*string, bool) {
	if o == nil || isNil(o.Ssid) {
    return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *InlineResponse200137) HasSsid() bool {
	if o != nil && !isNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *InlineResponse200137) SetSsid(v string) {
	o.Ssid = &v
}

// GetLoginAt returns the LoginAt field value if set, zero value otherwise.
func (o *InlineResponse200137) GetLoginAt() time.Time {
	if o == nil || isNil(o.LoginAt) {
		var ret time.Time
		return ret
	}
	return *o.LoginAt
}

// GetLoginAtOk returns a tuple with the LoginAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetLoginAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LoginAt) {
    return nil, false
	}
	return o.LoginAt, true
}

// HasLoginAt returns a boolean if a field has been set.
func (o *InlineResponse200137) HasLoginAt() bool {
	if o != nil && !isNil(o.LoginAt) {
		return true
	}

	return false
}

// SetLoginAt gets a reference to the given time.Time and assigns it to the LoginAt field.
func (o *InlineResponse200137) SetLoginAt(v time.Time) {
	o.LoginAt = &v
}

// GetGatewayDeviceMac returns the GatewayDeviceMac field value if set, zero value otherwise.
func (o *InlineResponse200137) GetGatewayDeviceMac() string {
	if o == nil || isNil(o.GatewayDeviceMac) {
		var ret string
		return ret
	}
	return *o.GatewayDeviceMac
}

// GetGatewayDeviceMacOk returns a tuple with the GatewayDeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetGatewayDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.GatewayDeviceMac) {
    return nil, false
	}
	return o.GatewayDeviceMac, true
}

// HasGatewayDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse200137) HasGatewayDeviceMac() bool {
	if o != nil && !isNil(o.GatewayDeviceMac) {
		return true
	}

	return false
}

// SetGatewayDeviceMac gets a reference to the given string and assigns it to the GatewayDeviceMac field.
func (o *InlineResponse200137) SetGatewayDeviceMac(v string) {
	o.GatewayDeviceMac = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *InlineResponse200137) GetClientMac() string {
	if o == nil || isNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetClientMacOk() (*string, bool) {
	if o == nil || isNil(o.ClientMac) {
    return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *InlineResponse200137) HasClientMac() bool {
	if o != nil && !isNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *InlineResponse200137) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineResponse200137) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineResponse200137) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineResponse200137) SetClientId(v string) {
	o.ClientId = &v
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *InlineResponse200137) GetAuthorization() string {
	if o == nil || isNil(o.Authorization) {
		var ret string
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetAuthorizationOk() (*string, bool) {
	if o == nil || isNil(o.Authorization) {
    return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *InlineResponse200137) HasAuthorization() bool {
	if o != nil && !isNil(o.Authorization) {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given string and assigns it to the Authorization field.
func (o *InlineResponse200137) SetAuthorization(v string) {
	o.Authorization = &v
}

func (o InlineResponse200137) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !isNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !isNil(o.LoginAt) {
		toSerialize["loginAt"] = o.LoginAt
	}
	if !isNil(o.GatewayDeviceMac) {
		toSerialize["gatewayDeviceMac"] = o.GatewayDeviceMac
	}
	if !isNil(o.ClientMac) {
		toSerialize["clientMac"] = o.ClientMac
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200137 struct {
	value *InlineResponse200137
	isSet bool
}

func (v NullableInlineResponse200137) Get() *InlineResponse200137 {
	return v.value
}

func (v *NullableInlineResponse200137) Set(val *InlineResponse200137) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200137) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200137(val *InlineResponse200137) *NullableInlineResponse200137 {
	return &NullableInlineResponse200137{value: val, isSet: true}
}

func (v NullableInlineResponse200137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


