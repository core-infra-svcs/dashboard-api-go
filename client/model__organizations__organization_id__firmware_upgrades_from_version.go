/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// OrganizationsOrganizationIdFirmwareUpgradesFromVersion ID of the upgrade's starting version
type OrganizationsOrganizationIdFirmwareUpgradesFromVersion struct {
	// Firmware version ID
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Firmware name
	Firmware *string `json:"firmware,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
}

// NewOrganizationsOrganizationIdFirmwareUpgradesFromVersion instantiates a new OrganizationsOrganizationIdFirmwareUpgradesFromVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFirmwareUpgradesFromVersion() *OrganizationsOrganizationIdFirmwareUpgradesFromVersion {
	this := OrganizationsOrganizationIdFirmwareUpgradesFromVersion{}
	return &this
}

// NewOrganizationsOrganizationIdFirmwareUpgradesFromVersionWithDefaults instantiates a new OrganizationsOrganizationIdFirmwareUpgradesFromVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFirmwareUpgradesFromVersionWithDefaults() *OrganizationsOrganizationIdFirmwareUpgradesFromVersion {
	this := OrganizationsOrganizationIdFirmwareUpgradesFromVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetFirmware() string {
	if o == nil || isNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetFirmwareOk() (*string, bool) {
	if o == nil || isNil(o.Firmware) {
    return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) HasFirmware() bool {
	if o != nil && !isNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) SetFirmware(v string) {
	o.Firmware = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetReleaseType() string {
	if o == nil || isNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseType) {
    return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) HasReleaseType() bool {
	if o != nil && !isNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetReleaseDate() time.Time {
	if o == nil || isNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ReleaseDate) {
    return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) HasReleaseDate() bool {
	if o != nil && !isNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o OrganizationsOrganizationIdFirmwareUpgradesFromVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !isNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !isNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion struct {
	value *OrganizationsOrganizationIdFirmwareUpgradesFromVersion
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion) Get() *OrganizationsOrganizationIdFirmwareUpgradesFromVersion {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion) Set(val *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion(val *OrganizationsOrganizationIdFirmwareUpgradesFromVersion) *NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion {
	return &NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesFromVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


