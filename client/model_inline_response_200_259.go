/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200259 struct for InlineResponse200259
type InlineResponse200259 struct {
	// The upgrade
	UpgradeId *string `json:"upgradeId,omitempty"`
	// The upgrade batch
	UpgradeBatchId *string `json:"upgradeBatchId,omitempty"`
	Network *OrganizationsOrganizationIdFirmwareUpgradesNetwork `json:"network,omitempty"`
	// Status of upgrade event: [Cancelled, Completed]
	Status *string `json:"status,omitempty"`
	// Scheduled start time
	Time *time.Time `json:"time,omitempty"`
	// Timestamp when upgrade completed. Null if status pending.
	CompletedAt *string `json:"completedAt,omitempty"`
	// product upgraded [wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor]
	ProductTypes *string `json:"productTypes,omitempty"`
	ToVersion *OrganizationsOrganizationIdFirmwareUpgradesToVersion `json:"toVersion,omitempty"`
	FromVersion *OrganizationsOrganizationIdFirmwareUpgradesFromVersion `json:"fromVersion,omitempty"`
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

// GetUpgradeId returns the UpgradeId field value if set, zero value otherwise.
func (o *InlineResponse200259) GetUpgradeId() string {
	if o == nil || isNil(o.UpgradeId) {
		var ret string
		return ret
	}
	return *o.UpgradeId
}

// GetUpgradeIdOk returns a tuple with the UpgradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetUpgradeIdOk() (*string, bool) {
	if o == nil || isNil(o.UpgradeId) {
    return nil, false
	}
	return o.UpgradeId, true
}

// HasUpgradeId returns a boolean if a field has been set.
func (o *InlineResponse200259) HasUpgradeId() bool {
	if o != nil && !isNil(o.UpgradeId) {
		return true
	}

	return false
}

// SetUpgradeId gets a reference to the given string and assigns it to the UpgradeId field.
func (o *InlineResponse200259) SetUpgradeId(v string) {
	o.UpgradeId = &v
}

// GetUpgradeBatchId returns the UpgradeBatchId field value if set, zero value otherwise.
func (o *InlineResponse200259) GetUpgradeBatchId() string {
	if o == nil || isNil(o.UpgradeBatchId) {
		var ret string
		return ret
	}
	return *o.UpgradeBatchId
}

// GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetUpgradeBatchIdOk() (*string, bool) {
	if o == nil || isNil(o.UpgradeBatchId) {
    return nil, false
	}
	return o.UpgradeBatchId, true
}

// HasUpgradeBatchId returns a boolean if a field has been set.
func (o *InlineResponse200259) HasUpgradeBatchId() bool {
	if o != nil && !isNil(o.UpgradeBatchId) {
		return true
	}

	return false
}

// SetUpgradeBatchId gets a reference to the given string and assigns it to the UpgradeBatchId field.
func (o *InlineResponse200259) SetUpgradeBatchId(v string) {
	o.UpgradeBatchId = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200259) GetNetwork() OrganizationsOrganizationIdFirmwareUpgradesNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetNetworkOk() (*OrganizationsOrganizationIdFirmwareUpgradesNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200259) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesNetwork and assigns it to the Network field.
func (o *InlineResponse200259) SetNetwork(v OrganizationsOrganizationIdFirmwareUpgradesNetwork) {
	o.Network = &v
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

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse200259) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse200259) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineResponse200259) SetTime(v time.Time) {
	o.Time = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *InlineResponse200259) GetCompletedAt() string {
	if o == nil || isNil(o.CompletedAt) {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetCompletedAtOk() (*string, bool) {
	if o == nil || isNil(o.CompletedAt) {
    return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse200259) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *InlineResponse200259) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse200259) GetProductTypes() string {
	if o == nil || isNil(o.ProductTypes) {
		var ret string
		return ret
	}
	return *o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetProductTypesOk() (*string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse200259) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given string and assigns it to the ProductTypes field.
func (o *InlineResponse200259) SetProductTypes(v string) {
	o.ProductTypes = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse200259) GetToVersion() OrganizationsOrganizationIdFirmwareUpgradesToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetToVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse200259) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesToVersion and assigns it to the ToVersion field.
func (o *InlineResponse200259) SetToVersion(v OrganizationsOrganizationIdFirmwareUpgradesToVersion) {
	o.ToVersion = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *InlineResponse200259) GetFromVersion() OrganizationsOrganizationIdFirmwareUpgradesFromVersion {
	if o == nil || isNil(o.FromVersion) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesFromVersion
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259) GetFromVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesFromVersion, bool) {
	if o == nil || isNil(o.FromVersion) {
    return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *InlineResponse200259) HasFromVersion() bool {
	if o != nil && !isNil(o.FromVersion) {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesFromVersion and assigns it to the FromVersion field.
func (o *InlineResponse200259) SetFromVersion(v OrganizationsOrganizationIdFirmwareUpgradesFromVersion) {
	o.FromVersion = &v
}

func (o InlineResponse200259) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UpgradeId) {
		toSerialize["upgradeId"] = o.UpgradeId
	}
	if !isNil(o.UpgradeBatchId) {
		toSerialize["upgradeBatchId"] = o.UpgradeBatchId
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	if !isNil(o.FromVersion) {
		toSerialize["fromVersion"] = o.FromVersion
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


