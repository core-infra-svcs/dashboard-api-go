/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSwitchPortsBySwitchPorts struct for OrganizationsOrganizationIdSwitchPortsBySwitchPorts
type OrganizationsOrganizationIdSwitchPortsBySwitchPorts struct {
	// The identifier of the switch port.
	PortId *string `json:"portId,omitempty"`
	// The name of the switch port.
	Name *string `json:"name,omitempty"`
	// The list of tags of the switch port.
	Tags []string `json:"tags,omitempty"`
	// The status of the switch port.
	Enabled *bool `json:"enabled,omitempty"`
	// The PoE status of the switch port.
	PoeEnabled *bool `json:"poeEnabled,omitempty"`
	// The type of the switch port ('trunk', 'access' or 'stack').
	Type *string `json:"type,omitempty"`
	// The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the switch port. Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the switch port. Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The rapid spanning tree protocol status.
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	StpGuard *string `json:"stpGuard,omitempty"`
	// The link speed for the switch port.
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowList []string `json:"stickyMacAllowList,omitempty"`
	// The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int32 `json:"stickyMacAllowListLimit,omitempty"`
}

// NewOrganizationsOrganizationIdSwitchPortsBySwitchPorts instantiates a new OrganizationsOrganizationIdSwitchPortsBySwitchPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSwitchPortsBySwitchPorts() *OrganizationsOrganizationIdSwitchPortsBySwitchPorts {
	this := OrganizationsOrganizationIdSwitchPortsBySwitchPorts{}
	return &this
}

// NewOrganizationsOrganizationIdSwitchPortsBySwitchPortsWithDefaults instantiates a new OrganizationsOrganizationIdSwitchPortsBySwitchPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSwitchPortsBySwitchPortsWithDefaults() *OrganizationsOrganizationIdSwitchPortsBySwitchPorts {
	this := OrganizationsOrganizationIdSwitchPortsBySwitchPorts{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetPortId() string {
	if o == nil || isNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetPortIdOk() (*string, bool) {
	if o == nil || isNil(o.PortId) {
    return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasPortId() bool {
	if o != nil && !isNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetPortId(v string) {
	o.PortId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetTags(v []string) {
	o.Tags = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPoeEnabled returns the PoeEnabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetPoeEnabled() bool {
	if o == nil || isNil(o.PoeEnabled) {
		var ret bool
		return ret
	}
	return *o.PoeEnabled
}

// GetPoeEnabledOk returns a tuple with the PoeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetPoeEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PoeEnabled) {
    return nil, false
	}
	return o.PoeEnabled, true
}

// HasPoeEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasPoeEnabled() bool {
	if o != nil && !isNil(o.PoeEnabled) {
		return true
	}

	return false
}

// SetPoeEnabled gets a reference to the given bool and assigns it to the PoeEnabled field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetPoeEnabled(v bool) {
	o.PoeEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetVoiceVlan() int32 {
	if o == nil || isNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || isNil(o.VoiceVlan) {
    return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasVoiceVlan() bool {
	if o != nil && !isNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetAllowedVlans() string {
	if o == nil || isNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetAllowedVlansOk() (*string, bool) {
	if o == nil || isNil(o.AllowedVlans) {
    return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasAllowedVlans() bool {
	if o != nil && !isNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetRstpEnabled() bool {
	if o == nil || isNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RstpEnabled) {
    return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasRstpEnabled() bool {
	if o != nil && !isNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpGuard returns the StpGuard field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetStpGuard() string {
	if o == nil || isNil(o.StpGuard) {
		var ret string
		return ret
	}
	return *o.StpGuard
}

// GetStpGuardOk returns a tuple with the StpGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetStpGuardOk() (*string, bool) {
	if o == nil || isNil(o.StpGuard) {
    return nil, false
	}
	return o.StpGuard, true
}

// HasStpGuard returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasStpGuard() bool {
	if o != nil && !isNil(o.StpGuard) {
		return true
	}

	return false
}

// SetStpGuard gets a reference to the given string and assigns it to the StpGuard field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetStpGuard(v string) {
	o.StpGuard = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetLinkNegotiation() string {
	if o == nil || isNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || isNil(o.LinkNegotiation) {
    return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasLinkNegotiation() bool {
	if o != nil && !isNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetAccessPolicyType() string {
	if o == nil || isNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessPolicyType) {
    return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasAccessPolicyType() bool {
	if o != nil && !isNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetStickyMacAllowList returns the StickyMacAllowList field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetStickyMacAllowList() []string {
	if o == nil || isNil(o.StickyMacAllowList) {
		var ret []string
		return ret
	}
	return o.StickyMacAllowList
}

// GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetStickyMacAllowListOk() ([]string, bool) {
	if o == nil || isNil(o.StickyMacAllowList) {
    return nil, false
	}
	return o.StickyMacAllowList, true
}

// HasStickyMacAllowList returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasStickyMacAllowList() bool {
	if o != nil && !isNil(o.StickyMacAllowList) {
		return true
	}

	return false
}

// SetStickyMacAllowList gets a reference to the given []string and assigns it to the StickyMacAllowList field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetStickyMacAllowList(v []string) {
	o.StickyMacAllowList = v
}

// GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetStickyMacAllowListLimit() int32 {
	if o == nil || isNil(o.StickyMacAllowListLimit) {
		var ret int32
		return ret
	}
	return *o.StickyMacAllowListLimit
}

// GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) GetStickyMacAllowListLimitOk() (*int32, bool) {
	if o == nil || isNil(o.StickyMacAllowListLimit) {
    return nil, false
	}
	return o.StickyMacAllowListLimit, true
}

// HasStickyMacAllowListLimit returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) HasStickyMacAllowListLimit() bool {
	if o != nil && !isNil(o.StickyMacAllowListLimit) {
		return true
	}

	return false
}

// SetStickyMacAllowListLimit gets a reference to the given int32 and assigns it to the StickyMacAllowListLimit field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) SetStickyMacAllowListLimit(v int32) {
	o.StickyMacAllowListLimit = &v
}

func (o OrganizationsOrganizationIdSwitchPortsBySwitchPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PoeEnabled) {
		toSerialize["poeEnabled"] = o.PoeEnabled
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.VoiceVlan) {
		toSerialize["voiceVlan"] = o.VoiceVlan
	}
	if !isNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if !isNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !isNil(o.StpGuard) {
		toSerialize["stpGuard"] = o.StpGuard
	}
	if !isNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	if !isNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !isNil(o.StickyMacAllowList) {
		toSerialize["stickyMacAllowList"] = o.StickyMacAllowList
	}
	if !isNil(o.StickyMacAllowListLimit) {
		toSerialize["stickyMacAllowListLimit"] = o.StickyMacAllowListLimit
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts struct {
	value *OrganizationsOrganizationIdSwitchPortsBySwitchPorts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts) Get() *OrganizationsOrganizationIdSwitchPortsBySwitchPorts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts) Set(val *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts(val *OrganizationsOrganizationIdSwitchPortsBySwitchPorts) *NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts {
	return &NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSwitchPortsBySwitchPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


