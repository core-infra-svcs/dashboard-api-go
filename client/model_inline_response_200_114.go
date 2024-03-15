/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200114 struct for InlineResponse200114
type InlineResponse200114 struct {
	// Whether APs will scan for Bluetooth enabled clients.
	ScanningEnabled *bool `json:"scanningEnabled,omitempty"`
	// Whether APs will advertise beacons.
	AdvertisingEnabled *bool `json:"advertisingEnabled,omitempty"`
	// The UUID to be used in the beacon identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
	MajorMinorAssignmentMode *string `json:"majorMinorAssignmentMode,omitempty"`
	// The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	Major *int32 `json:"major,omitempty"`
	// The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	Minor *int32 `json:"minor,omitempty"`
	// Whether ESL is enabled on this network.
	EslEnabled *bool `json:"eslEnabled,omitempty"`
}

// NewInlineResponse200114 instantiates a new InlineResponse200114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200114() *InlineResponse200114 {
	this := InlineResponse200114{}
	return &this
}

// NewInlineResponse200114WithDefaults instantiates a new InlineResponse200114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200114WithDefaults() *InlineResponse200114 {
	this := InlineResponse200114{}
	return &this
}

// GetScanningEnabled returns the ScanningEnabled field value if set, zero value otherwise.
func (o *InlineResponse200114) GetScanningEnabled() bool {
	if o == nil || isNil(o.ScanningEnabled) {
		var ret bool
		return ret
	}
	return *o.ScanningEnabled
}

// GetScanningEnabledOk returns a tuple with the ScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetScanningEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ScanningEnabled) {
    return nil, false
	}
	return o.ScanningEnabled, true
}

// HasScanningEnabled returns a boolean if a field has been set.
func (o *InlineResponse200114) HasScanningEnabled() bool {
	if o != nil && !isNil(o.ScanningEnabled) {
		return true
	}

	return false
}

// SetScanningEnabled gets a reference to the given bool and assigns it to the ScanningEnabled field.
func (o *InlineResponse200114) SetScanningEnabled(v bool) {
	o.ScanningEnabled = &v
}

// GetAdvertisingEnabled returns the AdvertisingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200114) GetAdvertisingEnabled() bool {
	if o == nil || isNil(o.AdvertisingEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertisingEnabled
}

// GetAdvertisingEnabledOk returns a tuple with the AdvertisingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetAdvertisingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AdvertisingEnabled) {
    return nil, false
	}
	return o.AdvertisingEnabled, true
}

// HasAdvertisingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200114) HasAdvertisingEnabled() bool {
	if o != nil && !isNil(o.AdvertisingEnabled) {
		return true
	}

	return false
}

// SetAdvertisingEnabled gets a reference to the given bool and assigns it to the AdvertisingEnabled field.
func (o *InlineResponse200114) SetAdvertisingEnabled(v bool) {
	o.AdvertisingEnabled = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InlineResponse200114) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InlineResponse200114) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InlineResponse200114) SetUuid(v string) {
	o.Uuid = &v
}

// GetMajorMinorAssignmentMode returns the MajorMinorAssignmentMode field value if set, zero value otherwise.
func (o *InlineResponse200114) GetMajorMinorAssignmentMode() string {
	if o == nil || isNil(o.MajorMinorAssignmentMode) {
		var ret string
		return ret
	}
	return *o.MajorMinorAssignmentMode
}

// GetMajorMinorAssignmentModeOk returns a tuple with the MajorMinorAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetMajorMinorAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.MajorMinorAssignmentMode) {
    return nil, false
	}
	return o.MajorMinorAssignmentMode, true
}

// HasMajorMinorAssignmentMode returns a boolean if a field has been set.
func (o *InlineResponse200114) HasMajorMinorAssignmentMode() bool {
	if o != nil && !isNil(o.MajorMinorAssignmentMode) {
		return true
	}

	return false
}

// SetMajorMinorAssignmentMode gets a reference to the given string and assigns it to the MajorMinorAssignmentMode field.
func (o *InlineResponse200114) SetMajorMinorAssignmentMode(v string) {
	o.MajorMinorAssignmentMode = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *InlineResponse200114) GetMajor() int32 {
	if o == nil || isNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetMajorOk() (*int32, bool) {
	if o == nil || isNil(o.Major) {
    return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *InlineResponse200114) HasMajor() bool {
	if o != nil && !isNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *InlineResponse200114) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *InlineResponse200114) GetMinor() int32 {
	if o == nil || isNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetMinorOk() (*int32, bool) {
	if o == nil || isNil(o.Minor) {
    return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *InlineResponse200114) HasMinor() bool {
	if o != nil && !isNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *InlineResponse200114) SetMinor(v int32) {
	o.Minor = &v
}

// GetEslEnabled returns the EslEnabled field value if set, zero value otherwise.
func (o *InlineResponse200114) GetEslEnabled() bool {
	if o == nil || isNil(o.EslEnabled) {
		var ret bool
		return ret
	}
	return *o.EslEnabled
}

// GetEslEnabledOk returns a tuple with the EslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200114) GetEslEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EslEnabled) {
    return nil, false
	}
	return o.EslEnabled, true
}

// HasEslEnabled returns a boolean if a field has been set.
func (o *InlineResponse200114) HasEslEnabled() bool {
	if o != nil && !isNil(o.EslEnabled) {
		return true
	}

	return false
}

// SetEslEnabled gets a reference to the given bool and assigns it to the EslEnabled field.
func (o *InlineResponse200114) SetEslEnabled(v bool) {
	o.EslEnabled = &v
}

func (o InlineResponse200114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScanningEnabled) {
		toSerialize["scanningEnabled"] = o.ScanningEnabled
	}
	if !isNil(o.AdvertisingEnabled) {
		toSerialize["advertisingEnabled"] = o.AdvertisingEnabled
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.MajorMinorAssignmentMode) {
		toSerialize["majorMinorAssignmentMode"] = o.MajorMinorAssignmentMode
	}
	if !isNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !isNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	if !isNil(o.EslEnabled) {
		toSerialize["eslEnabled"] = o.EslEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200114 struct {
	value *InlineResponse200114
	isSet bool
}

func (v NullableInlineResponse200114) Get() *InlineResponse200114 {
	return v.value
}

func (v *NullableInlineResponse200114) Set(val *InlineResponse200114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200114(val *InlineResponse200114) *NullableInlineResponse200114 {
	return &NullableInlineResponse200114{value: val, isSet: true}
}

func (v NullableInlineResponse200114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


