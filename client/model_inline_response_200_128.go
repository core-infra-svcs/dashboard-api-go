/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200128 struct for InlineResponse200128
type InlineResponse200128 struct {
	// Boolean indicating if the device is rooted.
	IsRooted *bool `json:"isRooted,omitempty"`
	// Boolean indicating if the device has Antivirus.
	HasAntiVirus *bool `json:"hasAntiVirus,omitempty"`
	// The name of the Antivirus.
	AntiVirusName *string `json:"antiVirusName,omitempty"`
	// Boolean indicating if the device has a Firewall enabled.
	IsFireWallEnabled *bool `json:"isFireWallEnabled,omitempty"`
	// Boolean indicating if the device has a Firewall installed.
	HasFireWallInstalled *bool `json:"hasFireWallInstalled,omitempty"`
	// The name of the Firewall.
	FireWallName *string `json:"fireWallName,omitempty"`
	// Boolean indicating if the device has disk encryption.
	IsDiskEncrypted *bool `json:"isDiskEncrypted,omitempty"`
	// Boolean indicating if the device has auto login disabled.
	IsAutoLoginDisabled *bool `json:"isAutoLoginDisabled,omitempty"`
	// The Meraki identifier for the security center record.
	Id *string `json:"id,omitempty"`
	// A comma seperated list of procs running on the device.
	RunningProcs *string `json:"runningProcs,omitempty"`
}

// NewInlineResponse200128 instantiates a new InlineResponse200128 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200128() *InlineResponse200128 {
	this := InlineResponse200128{}
	return &this
}

// NewInlineResponse200128WithDefaults instantiates a new InlineResponse200128 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200128WithDefaults() *InlineResponse200128 {
	this := InlineResponse200128{}
	return &this
}

// GetIsRooted returns the IsRooted field value if set, zero value otherwise.
func (o *InlineResponse200128) GetIsRooted() bool {
	if o == nil || isNil(o.IsRooted) {
		var ret bool
		return ret
	}
	return *o.IsRooted
}

// GetIsRootedOk returns a tuple with the IsRooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetIsRootedOk() (*bool, bool) {
	if o == nil || isNil(o.IsRooted) {
    return nil, false
	}
	return o.IsRooted, true
}

// HasIsRooted returns a boolean if a field has been set.
func (o *InlineResponse200128) HasIsRooted() bool {
	if o != nil && !isNil(o.IsRooted) {
		return true
	}

	return false
}

// SetIsRooted gets a reference to the given bool and assigns it to the IsRooted field.
func (o *InlineResponse200128) SetIsRooted(v bool) {
	o.IsRooted = &v
}

// GetHasAntiVirus returns the HasAntiVirus field value if set, zero value otherwise.
func (o *InlineResponse200128) GetHasAntiVirus() bool {
	if o == nil || isNil(o.HasAntiVirus) {
		var ret bool
		return ret
	}
	return *o.HasAntiVirus
}

// GetHasAntiVirusOk returns a tuple with the HasAntiVirus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetHasAntiVirusOk() (*bool, bool) {
	if o == nil || isNil(o.HasAntiVirus) {
    return nil, false
	}
	return o.HasAntiVirus, true
}

// HasHasAntiVirus returns a boolean if a field has been set.
func (o *InlineResponse200128) HasHasAntiVirus() bool {
	if o != nil && !isNil(o.HasAntiVirus) {
		return true
	}

	return false
}

// SetHasAntiVirus gets a reference to the given bool and assigns it to the HasAntiVirus field.
func (o *InlineResponse200128) SetHasAntiVirus(v bool) {
	o.HasAntiVirus = &v
}

// GetAntiVirusName returns the AntiVirusName field value if set, zero value otherwise.
func (o *InlineResponse200128) GetAntiVirusName() string {
	if o == nil || isNil(o.AntiVirusName) {
		var ret string
		return ret
	}
	return *o.AntiVirusName
}

// GetAntiVirusNameOk returns a tuple with the AntiVirusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetAntiVirusNameOk() (*string, bool) {
	if o == nil || isNil(o.AntiVirusName) {
    return nil, false
	}
	return o.AntiVirusName, true
}

// HasAntiVirusName returns a boolean if a field has been set.
func (o *InlineResponse200128) HasAntiVirusName() bool {
	if o != nil && !isNil(o.AntiVirusName) {
		return true
	}

	return false
}

// SetAntiVirusName gets a reference to the given string and assigns it to the AntiVirusName field.
func (o *InlineResponse200128) SetAntiVirusName(v string) {
	o.AntiVirusName = &v
}

// GetIsFireWallEnabled returns the IsFireWallEnabled field value if set, zero value otherwise.
func (o *InlineResponse200128) GetIsFireWallEnabled() bool {
	if o == nil || isNil(o.IsFireWallEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFireWallEnabled
}

// GetIsFireWallEnabledOk returns a tuple with the IsFireWallEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetIsFireWallEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsFireWallEnabled) {
    return nil, false
	}
	return o.IsFireWallEnabled, true
}

// HasIsFireWallEnabled returns a boolean if a field has been set.
func (o *InlineResponse200128) HasIsFireWallEnabled() bool {
	if o != nil && !isNil(o.IsFireWallEnabled) {
		return true
	}

	return false
}

// SetIsFireWallEnabled gets a reference to the given bool and assigns it to the IsFireWallEnabled field.
func (o *InlineResponse200128) SetIsFireWallEnabled(v bool) {
	o.IsFireWallEnabled = &v
}

// GetHasFireWallInstalled returns the HasFireWallInstalled field value if set, zero value otherwise.
func (o *InlineResponse200128) GetHasFireWallInstalled() bool {
	if o == nil || isNil(o.HasFireWallInstalled) {
		var ret bool
		return ret
	}
	return *o.HasFireWallInstalled
}

// GetHasFireWallInstalledOk returns a tuple with the HasFireWallInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetHasFireWallInstalledOk() (*bool, bool) {
	if o == nil || isNil(o.HasFireWallInstalled) {
    return nil, false
	}
	return o.HasFireWallInstalled, true
}

// HasHasFireWallInstalled returns a boolean if a field has been set.
func (o *InlineResponse200128) HasHasFireWallInstalled() bool {
	if o != nil && !isNil(o.HasFireWallInstalled) {
		return true
	}

	return false
}

// SetHasFireWallInstalled gets a reference to the given bool and assigns it to the HasFireWallInstalled field.
func (o *InlineResponse200128) SetHasFireWallInstalled(v bool) {
	o.HasFireWallInstalled = &v
}

// GetFireWallName returns the FireWallName field value if set, zero value otherwise.
func (o *InlineResponse200128) GetFireWallName() string {
	if o == nil || isNil(o.FireWallName) {
		var ret string
		return ret
	}
	return *o.FireWallName
}

// GetFireWallNameOk returns a tuple with the FireWallName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetFireWallNameOk() (*string, bool) {
	if o == nil || isNil(o.FireWallName) {
    return nil, false
	}
	return o.FireWallName, true
}

// HasFireWallName returns a boolean if a field has been set.
func (o *InlineResponse200128) HasFireWallName() bool {
	if o != nil && !isNil(o.FireWallName) {
		return true
	}

	return false
}

// SetFireWallName gets a reference to the given string and assigns it to the FireWallName field.
func (o *InlineResponse200128) SetFireWallName(v string) {
	o.FireWallName = &v
}

// GetIsDiskEncrypted returns the IsDiskEncrypted field value if set, zero value otherwise.
func (o *InlineResponse200128) GetIsDiskEncrypted() bool {
	if o == nil || isNil(o.IsDiskEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsDiskEncrypted
}

// GetIsDiskEncryptedOk returns a tuple with the IsDiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetIsDiskEncryptedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDiskEncrypted) {
    return nil, false
	}
	return o.IsDiskEncrypted, true
}

// HasIsDiskEncrypted returns a boolean if a field has been set.
func (o *InlineResponse200128) HasIsDiskEncrypted() bool {
	if o != nil && !isNil(o.IsDiskEncrypted) {
		return true
	}

	return false
}

// SetIsDiskEncrypted gets a reference to the given bool and assigns it to the IsDiskEncrypted field.
func (o *InlineResponse200128) SetIsDiskEncrypted(v bool) {
	o.IsDiskEncrypted = &v
}

// GetIsAutoLoginDisabled returns the IsAutoLoginDisabled field value if set, zero value otherwise.
func (o *InlineResponse200128) GetIsAutoLoginDisabled() bool {
	if o == nil || isNil(o.IsAutoLoginDisabled) {
		var ret bool
		return ret
	}
	return *o.IsAutoLoginDisabled
}

// GetIsAutoLoginDisabledOk returns a tuple with the IsAutoLoginDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetIsAutoLoginDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsAutoLoginDisabled) {
    return nil, false
	}
	return o.IsAutoLoginDisabled, true
}

// HasIsAutoLoginDisabled returns a boolean if a field has been set.
func (o *InlineResponse200128) HasIsAutoLoginDisabled() bool {
	if o != nil && !isNil(o.IsAutoLoginDisabled) {
		return true
	}

	return false
}

// SetIsAutoLoginDisabled gets a reference to the given bool and assigns it to the IsAutoLoginDisabled field.
func (o *InlineResponse200128) SetIsAutoLoginDisabled(v bool) {
	o.IsAutoLoginDisabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200128) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200128) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200128) SetId(v string) {
	o.Id = &v
}

// GetRunningProcs returns the RunningProcs field value if set, zero value otherwise.
func (o *InlineResponse200128) GetRunningProcs() string {
	if o == nil || isNil(o.RunningProcs) {
		var ret string
		return ret
	}
	return *o.RunningProcs
}

// GetRunningProcsOk returns a tuple with the RunningProcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetRunningProcsOk() (*string, bool) {
	if o == nil || isNil(o.RunningProcs) {
    return nil, false
	}
	return o.RunningProcs, true
}

// HasRunningProcs returns a boolean if a field has been set.
func (o *InlineResponse200128) HasRunningProcs() bool {
	if o != nil && !isNil(o.RunningProcs) {
		return true
	}

	return false
}

// SetRunningProcs gets a reference to the given string and assigns it to the RunningProcs field.
func (o *InlineResponse200128) SetRunningProcs(v string) {
	o.RunningProcs = &v
}

func (o InlineResponse200128) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsRooted) {
		toSerialize["isRooted"] = o.IsRooted
	}
	if !isNil(o.HasAntiVirus) {
		toSerialize["hasAntiVirus"] = o.HasAntiVirus
	}
	if !isNil(o.AntiVirusName) {
		toSerialize["antiVirusName"] = o.AntiVirusName
	}
	if !isNil(o.IsFireWallEnabled) {
		toSerialize["isFireWallEnabled"] = o.IsFireWallEnabled
	}
	if !isNil(o.HasFireWallInstalled) {
		toSerialize["hasFireWallInstalled"] = o.HasFireWallInstalled
	}
	if !isNil(o.FireWallName) {
		toSerialize["fireWallName"] = o.FireWallName
	}
	if !isNil(o.IsDiskEncrypted) {
		toSerialize["isDiskEncrypted"] = o.IsDiskEncrypted
	}
	if !isNil(o.IsAutoLoginDisabled) {
		toSerialize["isAutoLoginDisabled"] = o.IsAutoLoginDisabled
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RunningProcs) {
		toSerialize["runningProcs"] = o.RunningProcs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200128 struct {
	value *InlineResponse200128
	isSet bool
}

func (v NullableInlineResponse200128) Get() *InlineResponse200128 {
	return v.value
}

func (v *NullableInlineResponse200128) Set(val *InlineResponse200128) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200128) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200128) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200128(val *InlineResponse200128) *NullableInlineResponse200128 {
	return &NullableInlineResponse200128{value: val, isSet: true}
}

func (v NullableInlineResponse200128) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200128) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


