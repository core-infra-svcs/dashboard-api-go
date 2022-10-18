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

// InlineResponse20015ProductsWirelessNextUpgradeToVersion Details of the version the device will upgrade to if it exists
type InlineResponse20015ProductsWirelessNextUpgradeToVersion struct {
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

// NewInlineResponse20015ProductsWirelessNextUpgradeToVersion instantiates a new InlineResponse20015ProductsWirelessNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015ProductsWirelessNextUpgradeToVersion() *InlineResponse20015ProductsWirelessNextUpgradeToVersion {
	this := InlineResponse20015ProductsWirelessNextUpgradeToVersion{}
	return &this
}

// NewInlineResponse20015ProductsWirelessNextUpgradeToVersionWithDefaults instantiates a new InlineResponse20015ProductsWirelessNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015ProductsWirelessNextUpgradeToVersionWithDefaults() *InlineResponse20015ProductsWirelessNextUpgradeToVersion {
	this := InlineResponse20015ProductsWirelessNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetFirmwareOk() (*string, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) SetFirmware(v string) {
	o.Firmware = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetReleaseType() string {
	if o == nil || o.ReleaseType == nil {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || o.ReleaseType == nil {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) HasReleaseType() bool {
	if o != nil && o.ReleaseType != nil {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetReleaseDate() time.Time {
	if o == nil || o.ReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *InlineResponse20015ProductsWirelessNextUpgradeToVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o InlineResponse20015ProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion struct {
	value *InlineResponse20015ProductsWirelessNextUpgradeToVersion
	isSet bool
}

func (v NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion) Get() *InlineResponse20015ProductsWirelessNextUpgradeToVersion {
	return v.value
}

func (v *NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion) Set(val *InlineResponse20015ProductsWirelessNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20015ProductsWirelessNextUpgradeToVersion(val *InlineResponse20015ProductsWirelessNextUpgradeToVersion) *NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion {
	return &NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20015ProductsWirelessNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


