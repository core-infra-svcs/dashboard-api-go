/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20031 struct for InlineResponse20031
type InlineResponse20031 struct {
	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled *bool `json:"localStatusPageEnabled,omitempty"`
	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled *bool `json:"remoteStatusPageEnabled,omitempty"`
	LocalStatusPage *InlineResponse20031LocalStatusPage `json:"localStatusPage,omitempty"`
	SecurePort *InlineResponse20031SecurePort `json:"securePort,omitempty"`
	Fips *InlineResponse20031Fips `json:"fips,omitempty"`
	NamedVlans *InlineResponse20031NamedVlans `json:"namedVlans,omitempty"`
	ClientPrivacy *InlineResponse20031ClientPrivacy `json:"clientPrivacy,omitempty"`
}

// NewInlineResponse20031 instantiates a new InlineResponse20031 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20031() *InlineResponse20031 {
	this := InlineResponse20031{}
	return &this
}

// NewInlineResponse20031WithDefaults instantiates a new InlineResponse20031 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20031WithDefaults() *InlineResponse20031 {
	this := InlineResponse20031{}
	return &this
}

// GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineResponse20031) GetLocalStatusPageEnabled() bool {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.LocalStatusPageEnabled
}

// GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetLocalStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocalStatusPageEnabled) {
    return nil, false
	}
	return o.LocalStatusPageEnabled, true
}

// HasLocalStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineResponse20031) HasLocalStatusPageEnabled() bool {
	if o != nil && !isNil(o.LocalStatusPageEnabled) {
		return true
	}

	return false
}

// SetLocalStatusPageEnabled gets a reference to the given bool and assigns it to the LocalStatusPageEnabled field.
func (o *InlineResponse20031) SetLocalStatusPageEnabled(v bool) {
	o.LocalStatusPageEnabled = &v
}

// GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineResponse20031) GetRemoteStatusPageEnabled() bool {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
		var ret bool
		return ret
	}
	return *o.RemoteStatusPageEnabled
}

// GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetRemoteStatusPageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RemoteStatusPageEnabled) {
    return nil, false
	}
	return o.RemoteStatusPageEnabled, true
}

// HasRemoteStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineResponse20031) HasRemoteStatusPageEnabled() bool {
	if o != nil && !isNil(o.RemoteStatusPageEnabled) {
		return true
	}

	return false
}

// SetRemoteStatusPageEnabled gets a reference to the given bool and assigns it to the RemoteStatusPageEnabled field.
func (o *InlineResponse20031) SetRemoteStatusPageEnabled(v bool) {
	o.RemoteStatusPageEnabled = &v
}

// GetLocalStatusPage returns the LocalStatusPage field value if set, zero value otherwise.
func (o *InlineResponse20031) GetLocalStatusPage() InlineResponse20031LocalStatusPage {
	if o == nil || isNil(o.LocalStatusPage) {
		var ret InlineResponse20031LocalStatusPage
		return ret
	}
	return *o.LocalStatusPage
}

// GetLocalStatusPageOk returns a tuple with the LocalStatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetLocalStatusPageOk() (*InlineResponse20031LocalStatusPage, bool) {
	if o == nil || isNil(o.LocalStatusPage) {
    return nil, false
	}
	return o.LocalStatusPage, true
}

// HasLocalStatusPage returns a boolean if a field has been set.
func (o *InlineResponse20031) HasLocalStatusPage() bool {
	if o != nil && !isNil(o.LocalStatusPage) {
		return true
	}

	return false
}

// SetLocalStatusPage gets a reference to the given InlineResponse20031LocalStatusPage and assigns it to the LocalStatusPage field.
func (o *InlineResponse20031) SetLocalStatusPage(v InlineResponse20031LocalStatusPage) {
	o.LocalStatusPage = &v
}

// GetSecurePort returns the SecurePort field value if set, zero value otherwise.
func (o *InlineResponse20031) GetSecurePort() InlineResponse20031SecurePort {
	if o == nil || isNil(o.SecurePort) {
		var ret InlineResponse20031SecurePort
		return ret
	}
	return *o.SecurePort
}

// GetSecurePortOk returns a tuple with the SecurePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetSecurePortOk() (*InlineResponse20031SecurePort, bool) {
	if o == nil || isNil(o.SecurePort) {
    return nil, false
	}
	return o.SecurePort, true
}

// HasSecurePort returns a boolean if a field has been set.
func (o *InlineResponse20031) HasSecurePort() bool {
	if o != nil && !isNil(o.SecurePort) {
		return true
	}

	return false
}

// SetSecurePort gets a reference to the given InlineResponse20031SecurePort and assigns it to the SecurePort field.
func (o *InlineResponse20031) SetSecurePort(v InlineResponse20031SecurePort) {
	o.SecurePort = &v
}

// GetFips returns the Fips field value if set, zero value otherwise.
func (o *InlineResponse20031) GetFips() InlineResponse20031Fips {
	if o == nil || isNil(o.Fips) {
		var ret InlineResponse20031Fips
		return ret
	}
	return *o.Fips
}

// GetFipsOk returns a tuple with the Fips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetFipsOk() (*InlineResponse20031Fips, bool) {
	if o == nil || isNil(o.Fips) {
    return nil, false
	}
	return o.Fips, true
}

// HasFips returns a boolean if a field has been set.
func (o *InlineResponse20031) HasFips() bool {
	if o != nil && !isNil(o.Fips) {
		return true
	}

	return false
}

// SetFips gets a reference to the given InlineResponse20031Fips and assigns it to the Fips field.
func (o *InlineResponse20031) SetFips(v InlineResponse20031Fips) {
	o.Fips = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *InlineResponse20031) GetNamedVlans() InlineResponse20031NamedVlans {
	if o == nil || isNil(o.NamedVlans) {
		var ret InlineResponse20031NamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetNamedVlansOk() (*InlineResponse20031NamedVlans, bool) {
	if o == nil || isNil(o.NamedVlans) {
    return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *InlineResponse20031) HasNamedVlans() bool {
	if o != nil && !isNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given InlineResponse20031NamedVlans and assigns it to the NamedVlans field.
func (o *InlineResponse20031) SetNamedVlans(v InlineResponse20031NamedVlans) {
	o.NamedVlans = &v
}

// GetClientPrivacy returns the ClientPrivacy field value if set, zero value otherwise.
func (o *InlineResponse20031) GetClientPrivacy() InlineResponse20031ClientPrivacy {
	if o == nil || isNil(o.ClientPrivacy) {
		var ret InlineResponse20031ClientPrivacy
		return ret
	}
	return *o.ClientPrivacy
}

// GetClientPrivacyOk returns a tuple with the ClientPrivacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetClientPrivacyOk() (*InlineResponse20031ClientPrivacy, bool) {
	if o == nil || isNil(o.ClientPrivacy) {
    return nil, false
	}
	return o.ClientPrivacy, true
}

// HasClientPrivacy returns a boolean if a field has been set.
func (o *InlineResponse20031) HasClientPrivacy() bool {
	if o != nil && !isNil(o.ClientPrivacy) {
		return true
	}

	return false
}

// SetClientPrivacy gets a reference to the given InlineResponse20031ClientPrivacy and assigns it to the ClientPrivacy field.
func (o *InlineResponse20031) SetClientPrivacy(v InlineResponse20031ClientPrivacy) {
	o.ClientPrivacy = &v
}

func (o InlineResponse20031) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ClientPrivacy) {
		toSerialize["clientPrivacy"] = o.ClientPrivacy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20031 struct {
	value *InlineResponse20031
	isSet bool
}

func (v NullableInlineResponse20031) Get() *InlineResponse20031 {
	return v.value
}

func (v *NullableInlineResponse20031) Set(val *InlineResponse20031) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20031) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20031) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20031(val *InlineResponse20031) *NullableInlineResponse20031 {
	return &NullableInlineResponse20031{value: val, isSet: true}
}

func (v NullableInlineResponse20031) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20031) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


