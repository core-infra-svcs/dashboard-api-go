/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade The devices upgrade details and status
type OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade struct {
	// Start time of the upgrade
	Time *string `json:"time,omitempty"`
	FromVersion *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion `json:"fromVersion,omitempty"`
	ToVersion *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion `json:"toVersion,omitempty"`
	// Status of the upgrade
	Status *string `json:"status,omitempty"`
	// ID of the upgrade
	Id *string `json:"id,omitempty"`
	// ID of the upgrade batch
	UpgradeBatchId *string `json:"upgradeBatchId,omitempty"`
	Staged *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged `json:"staged,omitempty"`
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade{}
	return &this
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeWithDefaults instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeWithDefaults() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetTime() string {
	if o == nil || isNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetTimeOk() (*string, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetTime(v string) {
	o.Time = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetFromVersion() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion {
	if o == nil || isNil(o.FromVersion) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetFromVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion, bool) {
	if o == nil || isNil(o.FromVersion) {
    return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasFromVersion() bool {
	if o != nil && !isNil(o.FromVersion) {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion and assigns it to the FromVersion field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetFromVersion(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) {
	o.FromVersion = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetToVersion() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetToVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion and assigns it to the ToVersion field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetToVersion(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeToVersion) {
	o.ToVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetStatus(v string) {
	o.Status = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetId(v string) {
	o.Id = &v
}

// GetUpgradeBatchId returns the UpgradeBatchId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetUpgradeBatchId() string {
	if o == nil || isNil(o.UpgradeBatchId) {
		var ret string
		return ret
	}
	return *o.UpgradeBatchId
}

// GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetUpgradeBatchIdOk() (*string, bool) {
	if o == nil || isNil(o.UpgradeBatchId) {
    return nil, false
	}
	return o.UpgradeBatchId, true
}

// HasUpgradeBatchId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasUpgradeBatchId() bool {
	if o != nil && !isNil(o.UpgradeBatchId) {
		return true
	}

	return false
}

// SetUpgradeBatchId gets a reference to the given string and assigns it to the UpgradeBatchId field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetUpgradeBatchId(v string) {
	o.UpgradeBatchId = &v
}

// GetStaged returns the Staged field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStaged() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged {
	if o == nil || isNil(o.Staged) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged
		return ret
	}
	return *o.Staged
}

// GetStagedOk returns a tuple with the Staged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) GetStagedOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged, bool) {
	if o == nil || isNil(o.Staged) {
    return nil, false
	}
	return o.Staged, true
}

// HasStaged returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) HasStaged() bool {
	if o != nil && !isNil(o.Staged) {
		return true
	}

	return false
}

// SetStaged gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged and assigns it to the Staged field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) SetStaged(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) {
	o.Staged = &v
}

func (o OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.FromVersion) {
		toSerialize["fromVersion"] = o.FromVersion
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.UpgradeBatchId) {
		toSerialize["upgradeBatchId"] = o.UpgradeBatchId
	}
	if !isNil(o.Staged) {
		toSerialize["staged"] = o.Staged
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade struct {
	value *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) Get() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) Set(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade {
	return &NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


