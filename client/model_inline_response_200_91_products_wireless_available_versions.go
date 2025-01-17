/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20091ProductsWirelessAvailableVersions struct for InlineResponse20091ProductsWirelessAvailableVersions
type InlineResponse20091ProductsWirelessAvailableVersions struct {
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

// NewInlineResponse20091ProductsWirelessAvailableVersions instantiates a new InlineResponse20091ProductsWirelessAvailableVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091ProductsWirelessAvailableVersions() *InlineResponse20091ProductsWirelessAvailableVersions {
	this := InlineResponse20091ProductsWirelessAvailableVersions{}
	return &this
}

// NewInlineResponse20091ProductsWirelessAvailableVersionsWithDefaults instantiates a new InlineResponse20091ProductsWirelessAvailableVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091ProductsWirelessAvailableVersionsWithDefaults() *InlineResponse20091ProductsWirelessAvailableVersions {
	this := InlineResponse20091ProductsWirelessAvailableVersions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) SetId(v string) {
	o.Id = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetFirmware() string {
	if o == nil || isNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetFirmwareOk() (*string, bool) {
	if o == nil || isNil(o.Firmware) {
    return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) HasFirmware() bool {
	if o != nil && !isNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) SetFirmware(v string) {
	o.Firmware = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetReleaseType() string {
	if o == nil || isNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetReleaseTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseType) {
    return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) HasReleaseType() bool {
	if o != nil && !isNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetReleaseDate() time.Time {
	if o == nil || isNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ReleaseDate) {
    return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) HasReleaseDate() bool {
	if o != nil && !isNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *InlineResponse20091ProductsWirelessAvailableVersions) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o InlineResponse20091ProductsWirelessAvailableVersions) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20091ProductsWirelessAvailableVersions struct {
	value *InlineResponse20091ProductsWirelessAvailableVersions
	isSet bool
}

func (v NullableInlineResponse20091ProductsWirelessAvailableVersions) Get() *InlineResponse20091ProductsWirelessAvailableVersions {
	return v.value
}

func (v *NullableInlineResponse20091ProductsWirelessAvailableVersions) Set(val *InlineResponse20091ProductsWirelessAvailableVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091ProductsWirelessAvailableVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091ProductsWirelessAvailableVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091ProductsWirelessAvailableVersions(val *InlineResponse20091ProductsWirelessAvailableVersions) *NullableInlineResponse20091ProductsWirelessAvailableVersions {
	return &NullableInlineResponse20091ProductsWirelessAvailableVersions{value: val, isSet: true}
}

func (v NullableInlineResponse20091ProductsWirelessAvailableVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091ProductsWirelessAvailableVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


