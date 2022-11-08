/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20017ProductsWireless The network device to be updated
type InlineResponse20017ProductsWireless struct {
	CurrentVersion *InlineResponse20017ProductsWirelessCurrentVersion `json:"currentVersion,omitempty"`
	LastUpgrade *InlineResponse20017ProductsWirelessLastUpgrade `json:"lastUpgrade,omitempty"`
	NextUpgrade *InlineResponse20017ProductsWirelessNextUpgrade `json:"nextUpgrade,omitempty"`
	// Firmware versions available for upgrade
	AvailableVersions []InlineResponse20017ProductsWirelessAvailableVersions `json:"availableVersions,omitempty"`
	// Whether or not the network wants beta firmware
	ParticipateInNextBetaRelease *bool `json:"participateInNextBetaRelease,omitempty"`
}

// NewInlineResponse20017ProductsWireless instantiates a new InlineResponse20017ProductsWireless object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017ProductsWireless() *InlineResponse20017ProductsWireless {
	this := InlineResponse20017ProductsWireless{}
	return &this
}

// NewInlineResponse20017ProductsWirelessWithDefaults instantiates a new InlineResponse20017ProductsWireless object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017ProductsWirelessWithDefaults() *InlineResponse20017ProductsWireless {
	this := InlineResponse20017ProductsWireless{}
	return &this
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *InlineResponse20017ProductsWireless) GetCurrentVersion() InlineResponse20017ProductsWirelessCurrentVersion {
	if o == nil || isNil(o.CurrentVersion) {
		var ret InlineResponse20017ProductsWirelessCurrentVersion
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ProductsWireless) GetCurrentVersionOk() (*InlineResponse20017ProductsWirelessCurrentVersion, bool) {
	if o == nil || isNil(o.CurrentVersion) {
    return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *InlineResponse20017ProductsWireless) HasCurrentVersion() bool {
	if o != nil && !isNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given InlineResponse20017ProductsWirelessCurrentVersion and assigns it to the CurrentVersion field.
func (o *InlineResponse20017ProductsWireless) SetCurrentVersion(v InlineResponse20017ProductsWirelessCurrentVersion) {
	o.CurrentVersion = &v
}

// GetLastUpgrade returns the LastUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20017ProductsWireless) GetLastUpgrade() InlineResponse20017ProductsWirelessLastUpgrade {
	if o == nil || isNil(o.LastUpgrade) {
		var ret InlineResponse20017ProductsWirelessLastUpgrade
		return ret
	}
	return *o.LastUpgrade
}

// GetLastUpgradeOk returns a tuple with the LastUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ProductsWireless) GetLastUpgradeOk() (*InlineResponse20017ProductsWirelessLastUpgrade, bool) {
	if o == nil || isNil(o.LastUpgrade) {
    return nil, false
	}
	return o.LastUpgrade, true
}

// HasLastUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20017ProductsWireless) HasLastUpgrade() bool {
	if o != nil && !isNil(o.LastUpgrade) {
		return true
	}

	return false
}

// SetLastUpgrade gets a reference to the given InlineResponse20017ProductsWirelessLastUpgrade and assigns it to the LastUpgrade field.
func (o *InlineResponse20017ProductsWireless) SetLastUpgrade(v InlineResponse20017ProductsWirelessLastUpgrade) {
	o.LastUpgrade = &v
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20017ProductsWireless) GetNextUpgrade() InlineResponse20017ProductsWirelessNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret InlineResponse20017ProductsWirelessNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ProductsWireless) GetNextUpgradeOk() (*InlineResponse20017ProductsWirelessNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20017ProductsWireless) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given InlineResponse20017ProductsWirelessNextUpgrade and assigns it to the NextUpgrade field.
func (o *InlineResponse20017ProductsWireless) SetNextUpgrade(v InlineResponse20017ProductsWirelessNextUpgrade) {
	o.NextUpgrade = &v
}

// GetAvailableVersions returns the AvailableVersions field value if set, zero value otherwise.
func (o *InlineResponse20017ProductsWireless) GetAvailableVersions() []InlineResponse20017ProductsWirelessAvailableVersions {
	if o == nil || isNil(o.AvailableVersions) {
		var ret []InlineResponse20017ProductsWirelessAvailableVersions
		return ret
	}
	return o.AvailableVersions
}

// GetAvailableVersionsOk returns a tuple with the AvailableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ProductsWireless) GetAvailableVersionsOk() ([]InlineResponse20017ProductsWirelessAvailableVersions, bool) {
	if o == nil || isNil(o.AvailableVersions) {
    return nil, false
	}
	return o.AvailableVersions, true
}

// HasAvailableVersions returns a boolean if a field has been set.
func (o *InlineResponse20017ProductsWireless) HasAvailableVersions() bool {
	if o != nil && !isNil(o.AvailableVersions) {
		return true
	}

	return false
}

// SetAvailableVersions gets a reference to the given []InlineResponse20017ProductsWirelessAvailableVersions and assigns it to the AvailableVersions field.
func (o *InlineResponse20017ProductsWireless) SetAvailableVersions(v []InlineResponse20017ProductsWirelessAvailableVersions) {
	o.AvailableVersions = v
}

// GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field value if set, zero value otherwise.
func (o *InlineResponse20017ProductsWireless) GetParticipateInNextBetaRelease() bool {
	if o == nil || isNil(o.ParticipateInNextBetaRelease) {
		var ret bool
		return ret
	}
	return *o.ParticipateInNextBetaRelease
}

// GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017ProductsWireless) GetParticipateInNextBetaReleaseOk() (*bool, bool) {
	if o == nil || isNil(o.ParticipateInNextBetaRelease) {
    return nil, false
	}
	return o.ParticipateInNextBetaRelease, true
}

// HasParticipateInNextBetaRelease returns a boolean if a field has been set.
func (o *InlineResponse20017ProductsWireless) HasParticipateInNextBetaRelease() bool {
	if o != nil && !isNil(o.ParticipateInNextBetaRelease) {
		return true
	}

	return false
}

// SetParticipateInNextBetaRelease gets a reference to the given bool and assigns it to the ParticipateInNextBetaRelease field.
func (o *InlineResponse20017ProductsWireless) SetParticipateInNextBetaRelease(v bool) {
	o.ParticipateInNextBetaRelease = &v
}

func (o InlineResponse20017ProductsWireless) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentVersion) {
		toSerialize["currentVersion"] = o.CurrentVersion
	}
	if !isNil(o.LastUpgrade) {
		toSerialize["lastUpgrade"] = o.LastUpgrade
	}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	if !isNil(o.AvailableVersions) {
		toSerialize["availableVersions"] = o.AvailableVersions
	}
	if !isNil(o.ParticipateInNextBetaRelease) {
		toSerialize["participateInNextBetaRelease"] = o.ParticipateInNextBetaRelease
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017ProductsWireless struct {
	value *InlineResponse20017ProductsWireless
	isSet bool
}

func (v NullableInlineResponse20017ProductsWireless) Get() *InlineResponse20017ProductsWireless {
	return v.value
}

func (v *NullableInlineResponse20017ProductsWireless) Set(val *InlineResponse20017ProductsWireless) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017ProductsWireless) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017ProductsWireless) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017ProductsWireless(val *InlineResponse20017ProductsWireless) *NullableInlineResponse20017ProductsWireless {
	return &NullableInlineResponse20017ProductsWireless{value: val, isSet: true}
}

func (v NullableInlineResponse20017ProductsWireless) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017ProductsWireless) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


