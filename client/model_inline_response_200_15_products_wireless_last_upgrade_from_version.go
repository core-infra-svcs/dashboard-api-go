/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20015ProductsWirelessLastUpgradeFromVersion Details of the version the device upgraded from
type InlineResponse20015ProductsWirelessLastUpgradeFromVersion struct {
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

// NewInlineResponse20015ProductsWirelessLastUpgradeFromVersion instantiates a new InlineResponse20015ProductsWirelessLastUpgradeFromVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015ProductsWirelessLastUpgradeFromVersion() *InlineResponse20015ProductsWirelessLastUpgradeFromVersion {
	this := InlineResponse20015ProductsWirelessLastUpgradeFromVersion{}
	return &this
}

// NewInlineResponse20015ProductsWirelessLastUpgradeFromVersionWithDefaults instantiates a new InlineResponse20015ProductsWirelessLastUpgradeFromVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015ProductsWirelessLastUpgradeFromVersionWithDefaults() *InlineResponse20015ProductsWirelessLastUpgradeFromVersion {
	this := InlineResponse20015ProductsWirelessLastUpgradeFromVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) SetId(v string) {
	o.Id = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetFirmwareOk() (*string, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) SetFirmware(v string) {
	o.Firmware = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetReleaseType() string {
	if o == nil || o.ReleaseType == nil {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || o.ReleaseType == nil {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) HasReleaseType() bool {
	if o != nil && o.ReleaseType != nil {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetReleaseDate() time.Time {
	if o == nil || o.ReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o InlineResponse20015ProductsWirelessLastUpgradeFromVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Firmware != nil {
		toSerialize["firmware"] = o.Firmware
	}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.ReleaseType != nil {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if o.ReleaseDate != nil {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion struct {
	value *InlineResponse20015ProductsWirelessLastUpgradeFromVersion
	isSet bool
}

func (v NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion) Get() *InlineResponse20015ProductsWirelessLastUpgradeFromVersion {
	return v.value
}

func (v *NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion) Set(val *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion(val *InlineResponse20015ProductsWirelessLastUpgradeFromVersion) *NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion {
	return &NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion{value: val, isSet: true}
}

func (v NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20015ProductsWirelessLastUpgradeFromVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


