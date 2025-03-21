/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200120 struct for InlineResponse200120
type InlineResponse200120 struct {
	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled *bool `json:"localStatusPageEnabled,omitempty"`
	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled *bool `json:"remoteStatusPageEnabled,omitempty"`
	LocalStatusPage *InlineResponse200120LocalStatusPage `json:"localStatusPage,omitempty"`
	SecurePort *InlineResponse200120SecurePort `json:"securePort,omitempty"`
	Fips *InlineResponse200120Fips `json:"fips,omitempty"`
	NamedVlans *InlineResponse200120NamedVlans `json:"namedVlans,omitempty"`
}

// NewInlineResponse200120 instantiates a new InlineResponse200120 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200120() *InlineResponse200120 {
	this := InlineResponse200120{}
	return &this
}

// NewInlineResponse200120WithDefaults instantiates a new InlineResponse200120 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200120WithDefaults() *InlineResponse200120 {
	this := InlineResponse200120{}
	return &this
}

// GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineResponse200120) GetLocalStatusPageEnabled() bool {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.LocalStatusPageEnabled
}

// GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetLocalStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
    return nil, false
	}
	return o.LocalStatusPageEnabled, true
}

// HasLocalStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineResponse200120) HasLocalStatusPageEnabled() bool {
	if o != nil && !isNil(o.LocalStatusPageEnabled) {
		return true
	}

	return false
}

// SetLocalStatusPageEnabled gets a reference to the given bool and assigns it to the LocalStatusPageEnabled field.
func (o *InlineResponse200120) SetLocalStatusPageEnabled(v bool) {
	o.LocalStatusPageEnabled = &v
}

// GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineResponse200120) GetRemoteStatusPageEnabled() bool {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.RemoteStatusPageEnabled
}

// GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetRemoteStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
    return nil, false
	}
	return o.RemoteStatusPageEnabled, true
}

// HasRemoteStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineResponse200120) HasRemoteStatusPageEnabled() bool {
	if o != nil && !isNil(o.RemoteStatusPageEnabled) {
		return true
	}

	return false
}

// SetRemoteStatusPageEnabled gets a reference to the given bool and assigns it to the RemoteStatusPageEnabled field.
func (o *InlineResponse200120) SetRemoteStatusPageEnabled(v bool) {
	o.RemoteStatusPageEnabled = &v
}

// GetLocalStatusPage returns the LocalStatusPage field value if set, zero value otherwise.
func (o *InlineResponse200120) GetLocalStatusPage() InlineResponse200120LocalStatusPage {
	if o == nil || isNil(o.LocalStatusPage) {
		var ret InlineResponse200120LocalStatusPage
		return ret
	}
	return *o.LocalStatusPage
}

// GetLocalStatusPageOk returns a tuple with the LocalStatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetLocalStatusPageOk() (*InlineResponse200120LocalStatusPage, bool) {
	if o == nil || isNil(o.LocalStatusPage) {
    return nil, false
	}
	return o.LocalStatusPage, true
}

// HasLocalStatusPage returns a boolean if a field has been set.
func (o *InlineResponse200120) HasLocalStatusPage() bool {
	if o != nil && !isNil(o.LocalStatusPage) {
		return true
	}

	return false
}

// SetLocalStatusPage gets a reference to the given InlineResponse200120LocalStatusPage and assigns it to the LocalStatusPage field.
func (o *InlineResponse200120) SetLocalStatusPage(v InlineResponse200120LocalStatusPage) {
	o.LocalStatusPage = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *InlineResponse200120) GetSecurePort() InlineResponse200120SecurePort {
	if o == nil || isNil(o.SecurePort) {
		var ret InlineResponse200120SecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetSecurePortOk() (*InlineResponse200120SecurePort, bool) {
	if o == nil || isNil(o.SecurePort) {
    return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *InlineResponse200120) HasSecurePort() bool {
	if o != nil && !isNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given InlineResponse200120SecurePort and assigns it to the SecurePort field.
func (o *InlineResponse200120) SetSecurePort(v InlineResponse200120SecurePort) {
	o.SecurePort = &v
}

// GetFips returns the Fips field value if set, zero value otherwise.
func (o *InlineResponse200120) GetFips() InlineResponse200120Fips {
	if o == nil || isNil(o.Fips) {
		var ret InlineResponse200120Fips
		return ret
	}
	return *o.Fips
}

// GetFipsOk returns a tuple with the Fips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetFipsOk() (*InlineResponse200120Fips, bool) {
	if o == nil || isNil(o.Fips) {
    return nil, false
	}
	return o.Fips, true
}

// HasFips returns a boolean if a field has been set.
func (o *InlineResponse200120) HasFips() bool {
	if o != nil && !isNil(o.Fips) {
		return true
	}

	return false
}

// SetFips gets a reference to the given InlineResponse200120Fips and assigns it to the Fips field.
func (o *InlineResponse200120) SetFips(v InlineResponse200120Fips) {
	o.Fips = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *InlineResponse200120) GetNamedVlans() InlineResponse200120NamedVlans {
	if o == nil || isNil(o.NamedVlans) {
		var ret InlineResponse200120NamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetNamedVlansOk() (*InlineResponse200120NamedVlans, bool) {
	if o == nil || isNil(o.NamedVlans) {
    return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *InlineResponse200120) HasNamedVlans() bool {
	if o != nil && !isNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given InlineResponse200120NamedVlans and assigns it to the NamedVlans field.
func (o *InlineResponse200120) SetNamedVlans(v InlineResponse200120NamedVlans) {
	o.NamedVlans = &v
}

func (o InlineResponse200120) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200120 struct {
	value *InlineResponse200120
	isSet bool
}

func (v NullableInlineResponse200120) Get() *InlineResponse200120 {
	return v.value
}

func (v *NullableInlineResponse200120) Set(val *InlineResponse200120) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200120) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200120) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200120(val *InlineResponse200120) *NullableInlineResponse200120 {
	return &NullableInlineResponse200120{value: val, isSet: true}
}

func (v NullableInlineResponse200120) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200120) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


