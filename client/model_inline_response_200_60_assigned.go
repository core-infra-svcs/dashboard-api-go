/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20060Assigned struct for InlineResponse20060Assigned
type InlineResponse20060Assigned struct {
	// ID of the RF Profile.
	Id *string `json:"id,omitempty"`
	// ID of network this RF Profile belongs in.
	NetworkId *string `json:"networkId,omitempty"`
	// The name of the profile.
	Name *string `json:"name,omitempty"`
	TwoFourGhzSettings *InlineResponse20060TwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *InlineResponse20060FiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *InlineResponse20060PerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineResponse20060Assigned instantiates a new InlineResponse20060Assigned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060Assigned() *InlineResponse20060Assigned {
	this := InlineResponse20060Assigned{}
	return &this
}

// NewInlineResponse20060AssignedWithDefaults instantiates a new InlineResponse20060Assigned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060AssignedWithDefaults() *InlineResponse20060Assigned {
	this := InlineResponse20060Assigned{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20060Assigned) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assigned) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20060Assigned) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20060Assigned) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20060Assigned) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assigned) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20060Assigned) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20060Assigned) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20060Assigned) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assigned) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20060Assigned) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20060Assigned) SetName(v string) {
	o.Name = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse20060Assigned) GetTwoFourGhzSettings() InlineResponse20060TwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret InlineResponse20060TwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assigned) GetTwoFourGhzSettingsOk() (*InlineResponse20060TwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse20060Assigned) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given InlineResponse20060TwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineResponse20060Assigned) SetTwoFourGhzSettings(v InlineResponse20060TwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse20060Assigned) GetFiveGhzSettings() InlineResponse20060FiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret InlineResponse20060FiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assigned) GetFiveGhzSettingsOk() (*InlineResponse20060FiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse20060Assigned) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given InlineResponse20060FiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineResponse20060Assigned) SetFiveGhzSettings(v InlineResponse20060FiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineResponse20060Assigned) GetPerSsidSettings() InlineResponse20060PerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret InlineResponse20060PerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assigned) GetPerSsidSettingsOk() (*InlineResponse20060PerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineResponse20060Assigned) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given InlineResponse20060PerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineResponse20060Assigned) SetPerSsidSettings(v InlineResponse20060PerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineResponse20060Assigned) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !isNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20060Assigned struct {
	value *InlineResponse20060Assigned
	isSet bool
}

func (v NullableInlineResponse20060Assigned) Get() *InlineResponse20060Assigned {
	return v.value
}

func (v *NullableInlineResponse20060Assigned) Set(val *InlineResponse20060Assigned) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060Assigned) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060Assigned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060Assigned(val *InlineResponse20060Assigned) *NullableInlineResponse20060Assigned {
	return &NullableInlineResponse20060Assigned{value: val, isSet: true}
}

func (v NullableInlineResponse20060Assigned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060Assigned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


