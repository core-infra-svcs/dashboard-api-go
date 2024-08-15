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

// InlineResponse200259 struct for InlineResponse200259
type InlineResponse200259 struct {
	// License status (Co-termination licensing only)
	Status *string `json:"status,omitempty"`
	// License expiration date (Co-termination licensing only)
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// License counts (Co-termination licensing only)
	LicensedDeviceCounts *map[string]int32 `json:"licensedDeviceCounts,omitempty"`
	// Total number of licenses (Per-device licensing only)
	LicenseCount *int32 `json:"licenseCount,omitempty"`
	States *InlineResponse200259States `json:"states,omitempty"`
	// Data by license type (Per-device licensing only)
	LicenseTypes []InlineResponse200259LicenseTypes `json:"licenseTypes,omitempty"`
	SystemsManager *InlineResponse200259SystemsManager `json:"systemsManager,omitempty"`
}

// NewInlineResponse200259 instantiates a new InlineResponse200259 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200259() *InlineResponse200259 {
	this := InlineResponse200259{}
	return &this
}

// NewInlineResponse200259WithDefaults instantiates a new InlineResponse200259 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200259WithDefaults() *InlineResponse200259 {
	this := InlineResponse200259{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200259) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200259) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200259) SetStatus(v string) {
	o.Status = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *InlineResponse200259) GetExpirationDate() string {
	if o == nil || isNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetExpirationDateOk() (*string, bool) {
	if o == nil || isNil(o.ExpirationDate) {
    return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *InlineResponse200259) HasExpirationDate() bool {
	if o != nil && !isNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *InlineResponse200259) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetLicensedDeviceCounts returns the LicensedDeviceCounts field value if set, zero value otherwise.
func (o *InlineResponse200259) GetLicensedDeviceCounts() map[string]int32 {
	if o == nil || isNil(o.LicensedDeviceCounts) {
		var ret map[string]int32
		return ret
	}
	return *o.LicensedDeviceCounts
}

// GetLicensedDeviceCountsOk returns a tuple with the LicensedDeviceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetLicensedDeviceCountsOk() (*map[string]int32, bool) {
	if o == nil || isNil(o.LicensedDeviceCounts) {
    return nil, false
	}
	return o.LicensedDeviceCounts, true
}

// HasLicensedDeviceCounts returns a boolean if a field has been set.
func (o *InlineResponse200259) HasLicensedDeviceCounts() bool {
	if o != nil && !isNil(o.LicensedDeviceCounts) {
		return true
	}

	return false
}

// SetLicensedDeviceCounts gets a reference to the given map[string]int32 and assigns it to the LicensedDeviceCounts field.
func (o *InlineResponse200259) SetLicensedDeviceCounts(v map[string]int32) {
	o.LicensedDeviceCounts = &v
}

// GetLicenseCount returns the LicenseCount field value if set, zero value otherwise.
func (o *InlineResponse200259) GetLicenseCount() int32 {
	if o == nil || isNil(o.LicenseCount) {
		var ret int32
		return ret
	}
	return *o.LicenseCount
}

// GetLicenseCountOk returns a tuple with the LicenseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetLicenseCountOk() (*int32, bool) {
	if o == nil || isNil(o.LicenseCount) {
    return nil, false
	}
	return o.LicenseCount, true
}

// HasLicenseCount returns a boolean if a field has been set.
func (o *InlineResponse200259) HasLicenseCount() bool {
	if o != nil && !isNil(o.LicenseCount) {
		return true
	}

	return false
}

// SetLicenseCount gets a reference to the given int32 and assigns it to the LicenseCount field.
func (o *InlineResponse200259) SetLicenseCount(v int32) {
	o.LicenseCount = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *InlineResponse200259) GetStates() InlineResponse200259States {
	if o == nil || isNil(o.States) {
		var ret InlineResponse200259States
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetStatesOk() (*InlineResponse200259States, bool) {
	if o == nil || isNil(o.States) {
    return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *InlineResponse200259) HasStates() bool {
	if o != nil && !isNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given InlineResponse200259States and assigns it to the States field.
func (o *InlineResponse200259) SetStates(v InlineResponse200259States) {
	o.States = &v
}

// GetLicenseTypes returns the LicenseTypes field value if set, zero value otherwise.
func (o *InlineResponse200259) GetLicenseTypes() []InlineResponse200259LicenseTypes {
	if o == nil || isNil(o.LicenseTypes) {
		var ret []InlineResponse200259LicenseTypes
		return ret
	}
	return o.LicenseTypes
}

// GetLicenseTypesOk returns a tuple with the LicenseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetLicenseTypesOk() ([]InlineResponse200259LicenseTypes, bool) {
	if o == nil || isNil(o.LicenseTypes) {
    return nil, false
	}
	return o.LicenseTypes, true
}

// HasLicenseTypes returns a boolean if a field has been set.
func (o *InlineResponse200259) HasLicenseTypes() bool {
	if o != nil && !isNil(o.LicenseTypes) {
		return true
	}

	return false
}

// SetLicenseTypes gets a reference to the given []InlineResponse200259LicenseTypes and assigns it to the LicenseTypes field.
func (o *InlineResponse200259) SetLicenseTypes(v []InlineResponse200259LicenseTypes) {
	o.LicenseTypes = v
}

// GetSystemsManager returns the SystemsManager field value if set, zero value otherwise.
func (o *InlineResponse200259) GetSystemsManager() InlineResponse200259SystemsManager {
	if o == nil || isNil(o.SystemsManager) {
		var ret InlineResponse200259SystemsManager
		return ret
	}
	return *o.SystemsManager
}

// GetSystemsManagerOk returns a tuple with the SystemsManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetSystemsManagerOk() (*InlineResponse200259SystemsManager, bool) {
	if o == nil || isNil(o.SystemsManager) {
    return nil, false
	}
	return o.SystemsManager, true
}

// HasSystemsManager returns a boolean if a field has been set.
func (o *InlineResponse200259) HasSystemsManager() bool {
	if o != nil && !isNil(o.SystemsManager) {
		return true
	}

	return false
}

// SetSystemsManager gets a reference to the given InlineResponse200259SystemsManager and assigns it to the SystemsManager field.
func (o *InlineResponse200259) SetSystemsManager(v InlineResponse200259SystemsManager) {
	o.SystemsManager = &v
}

func (o InlineResponse200259) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !isNil(o.LicensedDeviceCounts) {
		toSerialize["licensedDeviceCounts"] = o.LicensedDeviceCounts
	}
	if !isNil(o.LicenseCount) {
		toSerialize["licenseCount"] = o.LicenseCount
	}
	if !isNil(o.States) {
		toSerialize["states"] = o.States
	}
	if !isNil(o.LicenseTypes) {
		toSerialize["licenseTypes"] = o.LicenseTypes
	}
	if !isNil(o.SystemsManager) {
		toSerialize["systemsManager"] = o.SystemsManager
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200259 struct {
	value *InlineResponse200259
	isSet bool
}

func (v NullableInlineResponse200259) Get() *InlineResponse200259 {
	return v.value
}

func (v *NullableInlineResponse200259) Set(val *InlineResponse200259) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200259) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200259) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200259(val *InlineResponse200259) *NullableInlineResponse200259 {
	return &NullableInlineResponse200259{value: val, isSet: true}
}

func (v NullableInlineResponse200259) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200259) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


