/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200116 struct for InlineResponse200116
type InlineResponse200116 struct {
	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled *bool `json:"localStatusPageEnabled,omitempty"`
	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled *bool `json:"remoteStatusPageEnabled,omitempty"`
	LocalStatusPage *InlineResponse200116LocalStatusPage `json:"localStatusPage,omitempty"`
	SecurePort *InlineResponse200116SecurePort `json:"securePort,omitempty"`
	Fips *InlineResponse200116Fips `json:"fips,omitempty"`
	NamedVlans *InlineResponse200116NamedVlans `json:"namedVlans,omitempty"`
}

// NewInlineResponse200116 instantiates a new InlineResponse200116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200116() *InlineResponse200116 {
	this := InlineResponse200116{}
	return &this
}

// NewInlineResponse200116WithDefaults instantiates a new InlineResponse200116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200116WithDefaults() *InlineResponse200116 {
	this := InlineResponse200116{}
	return &this
}

// GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineResponse200116) GetLocalStatusPageEnabled() bool {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.LocalStatusPageEnabled
}

// GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetLocalStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
    return nil, false
	}
	return o.LocalStatusPageEnabled, true
}

// HasLocalStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineResponse200116) HasLocalStatusPageEnabled() bool {
	if o != nil && !isNil(o.LocalStatusPageEnabled) {
		return true
	}

	return false
}

// SetLocalStatusPageEnabled gets a reference to the given bool and assigns it to the LocalStatusPageEnabled field.
func (o *InlineResponse200116) SetLocalStatusPageEnabled(v bool) {
	o.LocalStatusPageEnabled = &v
}

// GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineResponse200116) GetRemoteStatusPageEnabled() bool {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.RemoteStatusPageEnabled
}

// GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetRemoteStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
    return nil, false
	}
	return o.RemoteStatusPageEnabled, true
}

// HasRemoteStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineResponse200116) HasRemoteStatusPageEnabled() bool {
	if o != nil && !isNil(o.RemoteStatusPageEnabled) {
		return true
	}

	return false
}

// SetRemoteStatusPageEnabled gets a reference to the given bool and assigns it to the RemoteStatusPageEnabled field.
func (o *InlineResponse200116) SetRemoteStatusPageEnabled(v bool) {
	o.RemoteStatusPageEnabled = &v
}

// GetLocalStatusPage returns the LocalStatusPage field value if set, zero value otherwise.
func (o *InlineResponse200116) GetLocalStatusPage() InlineResponse200116LocalStatusPage {
	if o == nil || isNil(o.LocalStatusPage) {
		var ret InlineResponse200116LocalStatusPage
		return ret
	}
	return *o.LocalStatusPage
}

// GetLocalStatusPageOk returns a tuple with the LocalStatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetLocalStatusPageOk() (*InlineResponse200116LocalStatusPage, bool) {
	if o == nil || isNil(o.LocalStatusPage) {
    return nil, false
	}
	return o.LocalStatusPage, true
}

// HasLocalStatusPage returns a boolean if a field has been set.
func (o *InlineResponse200116) HasLocalStatusPage() bool {
	if o != nil && !isNil(o.LocalStatusPage) {
		return true
	}

	return false
}

// SetLocalStatusPage gets a reference to the given InlineResponse200116LocalStatusPage and assigns it to the LocalStatusPage field.
func (o *InlineResponse200116) SetLocalStatusPage(v InlineResponse200116LocalStatusPage) {
	o.LocalStatusPage = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *InlineResponse200116) GetSecurePort() InlineResponse200116SecurePort {
	if o == nil || isNil(o.SecurePort) {
		var ret InlineResponse200116SecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetSecurePortOk() (*InlineResponse200116SecurePort, bool) {
	if o == nil || isNil(o.SecurePort) {
    return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *InlineResponse200116) HasSecurePort() bool {
	if o != nil && !isNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given InlineResponse200116SecurePort and assigns it to the SecurePort field.
func (o *InlineResponse200116) SetSecurePort(v InlineResponse200116SecurePort) {
	o.SecurePort = &v
}

// GetFips returns the Fips field value if set, zero value otherwise.
func (o *InlineResponse200116) GetFips() InlineResponse200116Fips {
	if o == nil || isNil(o.Fips) {
		var ret InlineResponse200116Fips
		return ret
	}
	return *o.Fips
}

// GetFipsOk returns a tuple with the Fips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetFipsOk() (*InlineResponse200116Fips, bool) {
	if o == nil || isNil(o.Fips) {
    return nil, false
	}
	return o.Fips, true
}

// HasFips returns a boolean if a field has been set.
func (o *InlineResponse200116) HasFips() bool {
	if o != nil && !isNil(o.Fips) {
		return true
	}

	return false
}

// SetFips gets a reference to the given InlineResponse200116Fips and assigns it to the Fips field.
func (o *InlineResponse200116) SetFips(v InlineResponse200116Fips) {
	o.Fips = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *InlineResponse200116) GetNamedVlans() InlineResponse200116NamedVlans {
	if o == nil || isNil(o.NamedVlans) {
		var ret InlineResponse200116NamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetNamedVlansOk() (*InlineResponse200116NamedVlans, bool) {
	if o == nil || isNil(o.NamedVlans) {
    return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *InlineResponse200116) HasNamedVlans() bool {
	if o != nil && !isNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given InlineResponse200116NamedVlans and assigns it to the NamedVlans field.
func (o *InlineResponse200116) SetNamedVlans(v InlineResponse200116NamedVlans) {
	o.NamedVlans = &v
}

func (o InlineResponse200116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocalStatusPageEnabled) {
		toSerialize["localStatusPageEnabled"] = o.LocalStatusPageEnabled
	}
	if !isNil(o.RemoteStatusPageEnabled) {
		toSerialize["remoteStatusPageEnabled"] = o.RemoteStatusPageEnabled
	}
	if !isNil(o.LocalStatusPage) {
		toSerialize["localStatusPage"] = o.LocalStatusPage
	}
	if !isNil(o.SecurePort) {
		toSerialize["securePort"] = o.SecurePort
	}
	if !isNil(o.Fips) {
		toSerialize["fips"] = o.Fips
	}
	if !isNil(o.NamedVlans) {
		toSerialize["namedVlans"] = o.NamedVlans
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200116 struct {
	value *InlineResponse200116
	isSet bool
}

func (v NullableInlineResponse200116) Get() *InlineResponse200116 {
	return v.value
}

func (v *NullableInlineResponse200116) Set(val *InlineResponse200116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200116(val *InlineResponse200116) *NullableInlineResponse200116 {
	return &NullableInlineResponse200116{value: val, isSet: true}
}

func (v NullableInlineResponse200116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


