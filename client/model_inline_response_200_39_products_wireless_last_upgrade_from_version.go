/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20039ProductsWirelessLastUpgradeFromVersion Details of the version the device upgraded from
type InlineResponse20039ProductsWirelessLastUpgradeFromVersion struct {
	// Firmware version identifier
	Id *string `json:"id,omitempty"`
	// Name of the firmware version
	Firmware *string `json:"firmware,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
}

// NewInlineResponse20039ProductsWirelessLastUpgradeFromVersion instantiates a new InlineResponse20039ProductsWirelessLastUpgradeFromVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20039ProductsWirelessLastUpgradeFromVersion() *InlineResponse20039ProductsWirelessLastUpgradeFromVersion {
	this := InlineResponse20039ProductsWirelessLastUpgradeFromVersion{}
	return &this
}

// NewInlineResponse20039ProductsWirelessLastUpgradeFromVersionWithDefaults instantiates a new InlineResponse20039ProductsWirelessLastUpgradeFromVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20039ProductsWirelessLastUpgradeFromVersionWithDefaults() *InlineResponse20039ProductsWirelessLastUpgradeFromVersion {
	this := InlineResponse20039ProductsWirelessLastUpgradeFromVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) SetId(v string) {
	o.Id = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetFirmware() string {
	if o == nil || isNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetFirmwareOk() (*string, bool) {
	if o == nil || isNil(o.Firmware) {
    return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) HasFirmware() bool {
	if o != nil && !isNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) SetFirmware(v string) {
	o.Firmware = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetReleaseType() string {
	if o == nil || isNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseType) {
    return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) HasReleaseType() bool {
	if o != nil && !isNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetReleaseDate() time.Time {
	if o == nil || isNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ReleaseDate) {
    return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) HasReleaseDate() bool {
	if o != nil && !isNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o InlineResponse20039ProductsWirelessLastUpgradeFromVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !isNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion struct {
	value *InlineResponse20039ProductsWirelessLastUpgradeFromVersion
	isSet bool
}

func (v NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion) Get() *InlineResponse20039ProductsWirelessLastUpgradeFromVersion {
	return v.value
}

func (v *NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion) Set(val *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion(val *InlineResponse20039ProductsWirelessLastUpgradeFromVersion) *NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion {
	return &NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion{value: val, isSet: true}
}

func (v NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20039ProductsWirelessLastUpgradeFromVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

