/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20090ProductsWireless The network device to be updated
type InlineResponse20090ProductsWireless struct {
	CurrentVersion *InlineResponse20090ProductsWirelessCurrentVersion `json:"currentVersion,omitempty"`
	LastUpgrade *InlineResponse20090ProductsWirelessLastUpgrade `json:"lastUpgrade,omitempty"`
	NextUpgrade *InlineResponse20090ProductsWirelessNextUpgrade `json:"nextUpgrade,omitempty"`
	// Firmware versions available for upgrade
	AvailableVersions []InlineResponse20090ProductsWirelessAvailableVersions `json:"availableVersions,omitempty"`
	// Whether or not the network wants beta firmware
	ParticipateInNextBetaRelease *bool `json:"participateInNextBetaRelease,omitempty"`
}

// NewInlineResponse20090ProductsWireless instantiates a new InlineResponse20090ProductsWireless object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090ProductsWireless() *InlineResponse20090ProductsWireless {
	this := InlineResponse20090ProductsWireless{}
	return &this
}

// NewInlineResponse20090ProductsWirelessWithDefaults instantiates a new InlineResponse20090ProductsWireless object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090ProductsWirelessWithDefaults() *InlineResponse20090ProductsWireless {
	this := InlineResponse20090ProductsWireless{}
	return &this
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *InlineResponse20090ProductsWireless) GetCurrentVersion() InlineResponse20090ProductsWirelessCurrentVersion {
	if o == nil || isNil(o.CurrentVersion) {
		var ret InlineResponse20090ProductsWirelessCurrentVersion
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090ProductsWireless) GetCurrentVersionOk() (*InlineResponse20090ProductsWirelessCurrentVersion, bool) {
	if o == nil || isNil(o.CurrentVersion) {
    return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *InlineResponse20090ProductsWireless) HasCurrentVersion() bool {
	if o != nil && !isNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given InlineResponse20090ProductsWirelessCurrentVersion and assigns it to the CurrentVersion field.
func (o *InlineResponse20090ProductsWireless) SetCurrentVersion(v InlineResponse20090ProductsWirelessCurrentVersion) {
	o.CurrentVersion = &v
}

// GetLastUpgrade returns the LastUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20090ProductsWireless) GetLastUpgrade() InlineResponse20090ProductsWirelessLastUpgrade {
	if o == nil || isNil(o.LastUpgrade) {
		var ret InlineResponse20090ProductsWirelessLastUpgrade
		return ret
	}
	return *o.LastUpgrade
}

// GetLastUpgradeOk returns a tuple with the LastUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090ProductsWireless) GetLastUpgradeOk() (*InlineResponse20090ProductsWirelessLastUpgrade, bool) {
	if o == nil || isNil(o.LastUpgrade) {
    return nil, false
	}
	return o.LastUpgrade, true
}

// HasLastUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20090ProductsWireless) HasLastUpgrade() bool {
	if o != nil && !isNil(o.LastUpgrade) {
		return true
	}

	return false
}

// SetLastUpgrade gets a reference to the given InlineResponse20090ProductsWirelessLastUpgrade and assigns it to the LastUpgrade field.
func (o *InlineResponse20090ProductsWireless) SetLastUpgrade(v InlineResponse20090ProductsWirelessLastUpgrade) {
	o.LastUpgrade = &v
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20090ProductsWireless) GetNextUpgrade() InlineResponse20090ProductsWirelessNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret InlineResponse20090ProductsWirelessNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090ProductsWireless) GetNextUpgradeOk() (*InlineResponse20090ProductsWirelessNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20090ProductsWireless) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given InlineResponse20090ProductsWirelessNextUpgrade and assigns it to the NextUpgrade field.
func (o *InlineResponse20090ProductsWireless) SetNextUpgrade(v InlineResponse20090ProductsWirelessNextUpgrade) {
	o.NextUpgrade = &v
}

// GetAvailableVersions returns the AvailableVersions field value if set, zero value otherwise.
func (o *InlineResponse20090ProductsWireless) GetAvailableVersions() []InlineResponse20090ProductsWirelessAvailableVersions {
	if o == nil || isNil(o.AvailableVersions) {
		var ret []InlineResponse20090ProductsWirelessAvailableVersions
		return ret
	}
	return o.AvailableVersions
}

// GetAvailableVersionsOk returns a tuple with the AvailableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090ProductsWireless) GetAvailableVersionsOk() ([]InlineResponse20090ProductsWirelessAvailableVersions, bool) {
	if o == nil || isNil(o.AvailableVersions) {
    return nil, false
	}
	return o.AvailableVersions, true
}

// HasAvailableVersions returns a boolean if a field has been set.
func (o *InlineResponse20090ProductsWireless) HasAvailableVersions() bool {
	if o != nil && !isNil(o.AvailableVersions) {
		return true
	}

	return false
}

// SetAvailableVersions gets a reference to the given []InlineResponse20090ProductsWirelessAvailableVersions and assigns it to the AvailableVersions field.
func (o *InlineResponse20090ProductsWireless) SetAvailableVersions(v []InlineResponse20090ProductsWirelessAvailableVersions) {
	o.AvailableVersions = v
}

// GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field value if set, zero value otherwise.
func (o *InlineResponse20090ProductsWireless) GetParticipateInNextBetaRelease() bool {
	if o == nil || isNil(o.ParticipateInNextBetaRelease) {
		var ret bool
		return ret
	}
	return *o.ParticipateInNextBetaRelease
}

// GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090ProductsWireless) GetParticipateInNextBetaReleaseOk() (*bool, bool) {
	if o == nil || isNil(o.ParticipateInNextBetaRelease) {
    return nil, false
	}
	return o.ParticipateInNextBetaRelease, true
}

// HasParticipateInNextBetaRelease returns a boolean if a field has been set.
func (o *InlineResponse20090ProductsWireless) HasParticipateInNextBetaRelease() bool {
	if o != nil && !isNil(o.ParticipateInNextBetaRelease) {
		return true
	}

	return false
}

// SetParticipateInNextBetaRelease gets a reference to the given bool and assigns it to the ParticipateInNextBetaRelease field.
func (o *InlineResponse20090ProductsWireless) SetParticipateInNextBetaRelease(v bool) {
	o.ParticipateInNextBetaRelease = &v
}

func (o InlineResponse20090ProductsWireless) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20090ProductsWireless struct {
	value *InlineResponse20090ProductsWireless
	isSet bool
}

func (v NullableInlineResponse20090ProductsWireless) Get() *InlineResponse20090ProductsWireless {
	return v.value
}

func (v *NullableInlineResponse20090ProductsWireless) Set(val *InlineResponse20090ProductsWireless) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090ProductsWireless) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090ProductsWireless) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090ProductsWireless(val *InlineResponse20090ProductsWireless) *NullableInlineResponse20090ProductsWireless {
	return &NullableInlineResponse20090ProductsWireless{value: val, isSet: true}
}

func (v NullableInlineResponse20090ProductsWireless) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090ProductsWireless) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


