/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200289 struct for InlineResponse200289
type InlineResponse200289 struct {
	// Boolean indicating whether SNMP version 2c is enabled for the organization.
	V2cEnabled *bool `json:"v2cEnabled,omitempty"`
	// The community string for SNMP version 2c, if enabled.
	V2CommunityString *string `json:"v2CommunityString,omitempty"`
	// Boolean indicating whether SNMP version 3 is enabled for the organization.
	V3Enabled *bool `json:"v3Enabled,omitempty"`
	// The user for SNMP version 3, if enabled.
	V3User *string `json:"v3User,omitempty"`
	// The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
	V3AuthMode *string `json:"v3AuthMode,omitempty"`
	// The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
	V3PrivMode *string `json:"v3PrivMode,omitempty"`
	// The list of IPv4 addresses that are allowed to access the SNMP server.
	PeerIps []string `json:"peerIps,omitempty"`
	// The hostname of the SNMP server.
	Hostname *string `json:"hostname,omitempty"`
	// The port of the SNMP server.
	Port *int32 `json:"port,omitempty"`
}

// NewInlineResponse200289 instantiates a new InlineResponse200289 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200289() *InlineResponse200289 {
	this := InlineResponse200289{}
	return &this
}

// NewInlineResponse200289WithDefaults instantiates a new InlineResponse200289 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200289WithDefaults() *InlineResponse200289 {
	this := InlineResponse200289{}
	return &this
}

// GetV2cEnabled returns the V2cEnabled field value if set, zero value otherwise.
func (o *InlineResponse200289) GetV2cEnabled() bool {
	if o == nil || isNil(o.V2cEnabled) {
		var ret bool
		return ret
	}
	return *o.V2cEnabled
}

// GetV2cEnabledOk returns a tuple with the V2cEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetV2cEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.V2cEnabled) {
    return nil, false
	}
	return o.V2cEnabled, true
}

// HasV2cEnabled returns a boolean if a field has been set.
func (o *InlineResponse200289) HasV2cEnabled() bool {
	if o != nil && !isNil(o.V2cEnabled) {
		return true
	}

	return false
}

// SetV2cEnabled gets a reference to the given bool and assigns it to the V2cEnabled field.
func (o *InlineResponse200289) SetV2cEnabled(v bool) {
	o.V2cEnabled = &v
}

// GetV2CommunityString returns the V2CommunityString field value if set, zero value otherwise.
func (o *InlineResponse200289) GetV2CommunityString() string {
	if o == nil || isNil(o.V2CommunityString) {
		var ret string
		return ret
	}
	return *o.V2CommunityString
}

// GetV2CommunityStringOk returns a tuple with the V2CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetV2CommunityStringOk() (*string, bool) {
	if o == nil || isNil(o.V2CommunityString) {
    return nil, false
	}
	return o.V2CommunityString, true
}

// HasV2CommunityString returns a boolean if a field has been set.
func (o *InlineResponse200289) HasV2CommunityString() bool {
	if o != nil && !isNil(o.V2CommunityString) {
		return true
	}

	return false
}

// SetV2CommunityString gets a reference to the given string and assigns it to the V2CommunityString field.
func (o *InlineResponse200289) SetV2CommunityString(v string) {
	o.V2CommunityString = &v
}

// GetV3Enabled returns the V3Enabled field value if set, zero value otherwise.
func (o *InlineResponse200289) GetV3Enabled() bool {
	if o == nil || isNil(o.V3Enabled) {
		var ret bool
		return ret
	}
	return *o.V3Enabled
}

// GetV3EnabledOk returns a tuple with the V3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetV3EnabledOk() (*bool, bool) {
	if o == nil || isNil(o.V3Enabled) {
    return nil, false
	}
	return o.V3Enabled, true
}

// HasV3Enabled returns a boolean if a field has been set.
func (o *InlineResponse200289) HasV3Enabled() bool {
	if o != nil && !isNil(o.V3Enabled) {
		return true
	}

	return false
}

// SetV3Enabled gets a reference to the given bool and assigns it to the V3Enabled field.
func (o *InlineResponse200289) SetV3Enabled(v bool) {
	o.V3Enabled = &v
}

// GetV3User returns the V3User field value if set, zero value otherwise.
func (o *InlineResponse200289) GetV3User() string {
	if o == nil || isNil(o.V3User) {
		var ret string
		return ret
	}
	return *o.V3User
}

// GetV3UserOk returns a tuple with the V3User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetV3UserOk() (*string, bool) {
	if o == nil || isNil(o.V3User) {
    return nil, false
	}
	return o.V3User, true
}

// HasV3User returns a boolean if a field has been set.
func (o *InlineResponse200289) HasV3User() bool {
	if o != nil && !isNil(o.V3User) {
		return true
	}

	return false
}

// SetV3User gets a reference to the given string and assigns it to the V3User field.
func (o *InlineResponse200289) SetV3User(v string) {
	o.V3User = &v
}

// GetV3AuthMode returns the V3AuthMode field value if set, zero value otherwise.
func (o *InlineResponse200289) GetV3AuthMode() string {
	if o == nil || isNil(o.V3AuthMode) {
		var ret string
		return ret
	}
	return *o.V3AuthMode
}

// GetV3AuthModeOk returns a tuple with the V3AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetV3AuthModeOk() (*string, bool) {
	if o == nil || isNil(o.V3AuthMode) {
    return nil, false
	}
	return o.V3AuthMode, true
}

// HasV3AuthMode returns a boolean if a field has been set.
func (o *InlineResponse200289) HasV3AuthMode() bool {
	if o != nil && !isNil(o.V3AuthMode) {
		return true
	}

	return false
}

// SetV3AuthMode gets a reference to the given string and assigns it to the V3AuthMode field.
func (o *InlineResponse200289) SetV3AuthMode(v string) {
	o.V3AuthMode = &v
}

// GetV3PrivMode returns the V3PrivMode field value if set, zero value otherwise.
func (o *InlineResponse200289) GetV3PrivMode() string {
	if o == nil || isNil(o.V3PrivMode) {
		var ret string
		return ret
	}
	return *o.V3PrivMode
}

// GetV3PrivModeOk returns a tuple with the V3PrivMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetV3PrivModeOk() (*string, bool) {
	if o == nil || isNil(o.V3PrivMode) {
    return nil, false
	}
	return o.V3PrivMode, true
}

// HasV3PrivMode returns a boolean if a field has been set.
func (o *InlineResponse200289) HasV3PrivMode() bool {
	if o != nil && !isNil(o.V3PrivMode) {
		return true
	}

	return false
}

// SetV3PrivMode gets a reference to the given string and assigns it to the V3PrivMode field.
func (o *InlineResponse200289) SetV3PrivMode(v string) {
	o.V3PrivMode = &v
}

// GetPeerIps returns the PeerIps field value if set, zero value otherwise.
func (o *InlineResponse200289) GetPeerIps() []string {
	if o == nil || isNil(o.PeerIps) {
		var ret []string
		return ret
	}
	return o.PeerIps
}

// GetPeerIpsOk returns a tuple with the PeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetPeerIpsOk() ([]string, bool) {
	if o == nil || isNil(o.PeerIps) {
    return nil, false
	}
	return o.PeerIps, true
}

// HasPeerIps returns a boolean if a field has been set.
func (o *InlineResponse200289) HasPeerIps() bool {
	if o != nil && !isNil(o.PeerIps) {
		return true
	}

	return false
}

// SetPeerIps gets a reference to the given []string and assigns it to the PeerIps field.
func (o *InlineResponse200289) SetPeerIps(v []string) {
	o.PeerIps = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *InlineResponse200289) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *InlineResponse200289) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *InlineResponse200289) SetHostname(v string) {
	o.Hostname = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse200289) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse200289) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InlineResponse200289) SetPort(v int32) {
	o.Port = &v
}

func (o InlineResponse200289) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.V2cEnabled) {
		toSerialize["v2cEnabled"] = o.V2cEnabled
	}
	if !isNil(o.V2CommunityString) {
		toSerialize["v2CommunityString"] = o.V2CommunityString
	}
	if !isNil(o.V3Enabled) {
		toSerialize["v3Enabled"] = o.V3Enabled
	}
	if !isNil(o.V3User) {
		toSerialize["v3User"] = o.V3User
	}
	if !isNil(o.V3AuthMode) {
		toSerialize["v3AuthMode"] = o.V3AuthMode
	}
	if !isNil(o.V3PrivMode) {
		toSerialize["v3PrivMode"] = o.V3PrivMode
	}
	if !isNil(o.PeerIps) {
		toSerialize["peerIps"] = o.PeerIps
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200289 struct {
	value *InlineResponse200289
	isSet bool
}

func (v NullableInlineResponse200289) Get() *InlineResponse200289 {
	return v.value
}

func (v *NullableInlineResponse200289) Set(val *InlineResponse200289) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200289) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200289) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200289(val *InlineResponse200289) *NullableInlineResponse200289 {
	return &NullableInlineResponse200289{value: val, isSet: true}
}

func (v NullableInlineResponse200289) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200289) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


