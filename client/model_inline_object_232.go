/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject232 struct for InlineObject232
type InlineObject232 struct {
	// The name of the switch template port.
	Name *string `json:"name,omitempty"`
	// The list of tags of the switch template port.
	Tags []string `json:"tags,omitempty"`
	// The status of the switch template port.
	Enabled *bool `json:"enabled,omitempty"`
	// The PoE status of the switch template port.
	PoeEnabled *bool `json:"poeEnabled,omitempty"`
	// The type of the switch template port ('trunk', 'access' or 'stack').
	Type *string `json:"type,omitempty"`
	// The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the switch template port. Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the switch template port. Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The isolation status of the switch template port.
	IsolationEnabled *bool `json:"isolationEnabled,omitempty"`
	// The rapid spanning tree protocol status.
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	StpGuard *string `json:"stpGuard,omitempty"`
	// The link speed for the switch template port.
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The ID of the port schedule. A value of null will clear the port schedule.
	PortScheduleId *string `json:"portScheduleId,omitempty"`
	// The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	Udld *string `json:"udld,omitempty"`
	// The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyNumber *int32 `json:"accessPolicyNumber,omitempty"`
	// Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	MacAllowList []string `json:"macAllowList,omitempty"`
	// The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowList []string `json:"stickyMacAllowList,omitempty"`
	// The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int32 `json:"stickyMacAllowListLimit,omitempty"`
	// The storm control status of the switch template port.
	StormControlEnabled *bool `json:"stormControlEnabled,omitempty"`
	// For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	FlexibleStackingEnabled *bool `json:"flexibleStackingEnabled,omitempty"`
	// If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	DaiTrusted *bool `json:"daiTrusted,omitempty"`
	Profile *DevicesSerialSwitchPortsProfile `json:"profile,omitempty"`
	Dot3az *OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az `json:"dot3az,omitempty"`
}

// NewInlineObject232 instantiates a new InlineObject232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject232() *InlineObject232 {
	this := InlineObject232{}
	return &this
}

// NewInlineObject232WithDefaults instantiates a new InlineObject232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject232WithDefaults() *InlineObject232 {
	this := InlineObject232{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject232) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject232) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject232) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject232) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject232) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject232) SetTags(v []string) {
	o.Tags = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject232) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject232) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject232) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPoeEnabled returns the PoeEnabled field value if set, zero value otherwise.
func (o *InlineObject232) GetPoeEnabled() bool {
	if o == nil || isNil(o.PoeEnabled) {
		var ret bool
		return ret
	}
	return *o.PoeEnabled
}

// GetPoeEnabledOk returns a tuple with the PoeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetPoeEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PoeEnabled) {
    return nil, false
	}
	return o.PoeEnabled, true
}

// HasPoeEnabled returns a boolean if a field has been set.
func (o *InlineObject232) HasPoeEnabled() bool {
	if o != nil && !isNil(o.PoeEnabled) {
		return true
	}

	return false
}

// SetPoeEnabled gets a reference to the given bool and assigns it to the PoeEnabled field.
func (o *InlineObject232) SetPoeEnabled(v bool) {
	o.PoeEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject232) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject232) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject232) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject232) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject232) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject232) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *InlineObject232) GetVoiceVlan() int32 {
	if o == nil || isNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || isNil(o.VoiceVlan) {
    return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *InlineObject232) HasVoiceVlan() bool {
	if o != nil && !isNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *InlineObject232) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *InlineObject232) GetAllowedVlans() string {
	if o == nil || isNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetAllowedVlansOk() (*string, bool) {
	if o == nil || isNil(o.AllowedVlans) {
    return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *InlineObject232) HasAllowedVlans() bool {
	if o != nil && !isNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *InlineObject232) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetIsolationEnabled returns the IsolationEnabled field value if set, zero value otherwise.
func (o *InlineObject232) GetIsolationEnabled() bool {
	if o == nil || isNil(o.IsolationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsolationEnabled
}

// GetIsolationEnabledOk returns a tuple with the IsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetIsolationEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsolationEnabled) {
    return nil, false
	}
	return o.IsolationEnabled, true
}

// HasIsolationEnabled returns a boolean if a field has been set.
func (o *InlineObject232) HasIsolationEnabled() bool {
	if o != nil && !isNil(o.IsolationEnabled) {
		return true
	}

	return false
}

// SetIsolationEnabled gets a reference to the given bool and assigns it to the IsolationEnabled field.
func (o *InlineObject232) SetIsolationEnabled(v bool) {
	o.IsolationEnabled = &v
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *InlineObject232) GetRstpEnabled() bool {
	if o == nil || isNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RstpEnabled) {
    return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *InlineObject232) HasRstpEnabled() bool {
	if o != nil && !isNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *InlineObject232) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpGuard returns the StpGuard field value if set, zero value otherwise.
func (o *InlineObject232) GetStpGuard() string {
	if o == nil || isNil(o.StpGuard) {
		var ret string
		return ret
	}
	return *o.StpGuard
}

// GetStpGuardOk returns a tuple with the StpGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetStpGuardOk() (*string, bool) {
	if o == nil || isNil(o.StpGuard) {
    return nil, false
	}
	return o.StpGuard, true
}

// HasStpGuard returns a boolean if a field has been set.
func (o *InlineObject232) HasStpGuard() bool {
	if o != nil && !isNil(o.StpGuard) {
		return true
	}

	return false
}

// SetStpGuard gets a reference to the given string and assigns it to the StpGuard field.
func (o *InlineObject232) SetStpGuard(v string) {
	o.StpGuard = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *InlineObject232) GetLinkNegotiation() string {
	if o == nil || isNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || isNil(o.LinkNegotiation) {
    return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *InlineObject232) HasLinkNegotiation() bool {
	if o != nil && !isNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *InlineObject232) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetPortScheduleId returns the PortScheduleId field value if set, zero value otherwise.
func (o *InlineObject232) GetPortScheduleId() string {
	if o == nil || isNil(o.PortScheduleId) {
		var ret string
		return ret
	}
	return *o.PortScheduleId
}

// GetPortScheduleIdOk returns a tuple with the PortScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetPortScheduleIdOk() (*string, bool) {
	if o == nil || isNil(o.PortScheduleId) {
    return nil, false
	}
	return o.PortScheduleId, true
}

// HasPortScheduleId returns a boolean if a field has been set.
func (o *InlineObject232) HasPortScheduleId() bool {
	if o != nil && !isNil(o.PortScheduleId) {
		return true
	}

	return false
}

// SetPortScheduleId gets a reference to the given string and assigns it to the PortScheduleId field.
func (o *InlineObject232) SetPortScheduleId(v string) {
	o.PortScheduleId = &v
}

// GetUdld returns the Udld field value if set, zero value otherwise.
func (o *InlineObject232) GetUdld() string {
	if o == nil || isNil(o.Udld) {
		var ret string
		return ret
	}
	return *o.Udld
}

// GetUdldOk returns a tuple with the Udld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetUdldOk() (*string, bool) {
	if o == nil || isNil(o.Udld) {
    return nil, false
	}
	return o.Udld, true
}

// HasUdld returns a boolean if a field has been set.
func (o *InlineObject232) HasUdld() bool {
	if o != nil && !isNil(o.Udld) {
		return true
	}

	return false
}

// SetUdld gets a reference to the given string and assigns it to the Udld field.
func (o *InlineObject232) SetUdld(v string) {
	o.Udld = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *InlineObject232) GetAccessPolicyType() string {
	if o == nil || isNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessPolicyType) {
    return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *InlineObject232) HasAccessPolicyType() bool {
	if o != nil && !isNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *InlineObject232) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetAccessPolicyNumber returns the AccessPolicyNumber field value if set, zero value otherwise.
func (o *InlineObject232) GetAccessPolicyNumber() int32 {
	if o == nil || isNil(o.AccessPolicyNumber) {
		var ret int32
		return ret
	}
	return *o.AccessPolicyNumber
}

// GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetAccessPolicyNumberOk() (*int32, bool) {
	if o == nil || isNil(o.AccessPolicyNumber) {
    return nil, false
	}
	return o.AccessPolicyNumber, true
}

// HasAccessPolicyNumber returns a boolean if a field has been set.
func (o *InlineObject232) HasAccessPolicyNumber() bool {
	if o != nil && !isNil(o.AccessPolicyNumber) {
		return true
	}

	return false
}

// SetAccessPolicyNumber gets a reference to the given int32 and assigns it to the AccessPolicyNumber field.
func (o *InlineObject232) SetAccessPolicyNumber(v int32) {
	o.AccessPolicyNumber = &v
}

// GetMacAllowList returns the MacAllowList field value if set, zero value otherwise.
func (o *InlineObject232) GetMacAllowList() []string {
	if o == nil || isNil(o.MacAllowList) {
		var ret []string
		return ret
	}
	return o.MacAllowList
}

// GetMacAllowListOk returns a tuple with the MacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetMacAllowListOk() ([]string, bool) {
	if o == nil || isNil(o.MacAllowList) {
    return nil, false
	}
	return o.MacAllowList, true
}

// HasMacAllowList returns a boolean if a field has been set.
func (o *InlineObject232) HasMacAllowList() bool {
	if o != nil && !isNil(o.MacAllowList) {
		return true
	}

	return false
}

// SetMacAllowList gets a reference to the given []string and assigns it to the MacAllowList field.
func (o *InlineObject232) SetMacAllowList(v []string) {
	o.MacAllowList = v
}

// GetStickyMacAllowList returns the StickyMacAllowList field value if set, zero value otherwise.
func (o *InlineObject232) GetStickyMacAllowList() []string {
	if o == nil || isNil(o.StickyMacAllowList) {
		var ret []string
		return ret
	}
	return o.StickyMacAllowList
}

// GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetStickyMacAllowListOk() ([]string, bool) {
	if o == nil || isNil(o.StickyMacAllowList) {
    return nil, false
	}
	return o.StickyMacAllowList, true
}

// HasStickyMacAllowList returns a boolean if a field has been set.
func (o *InlineObject232) HasStickyMacAllowList() bool {
	if o != nil && !isNil(o.StickyMacAllowList) {
		return true
	}

	return false
}

// SetStickyMacAllowList gets a reference to the given []string and assigns it to the StickyMacAllowList field.
func (o *InlineObject232) SetStickyMacAllowList(v []string) {
	o.StickyMacAllowList = v
}

// GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field value if set, zero value otherwise.
func (o *InlineObject232) GetStickyMacAllowListLimit() int32 {
	if o == nil || isNil(o.StickyMacAllowListLimit) {
		var ret int32
		return ret
	}
	return *o.StickyMacAllowListLimit
}

// GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetStickyMacAllowListLimitOk() (*int32, bool) {
	if o == nil || isNil(o.StickyMacAllowListLimit) {
    return nil, false
	}
	return o.StickyMacAllowListLimit, true
}

// HasStickyMacAllowListLimit returns a boolean if a field has been set.
func (o *InlineObject232) HasStickyMacAllowListLimit() bool {
	if o != nil && !isNil(o.StickyMacAllowListLimit) {
		return true
	}

	return false
}

// SetStickyMacAllowListLimit gets a reference to the given int32 and assigns it to the StickyMacAllowListLimit field.
func (o *InlineObject232) SetStickyMacAllowListLimit(v int32) {
	o.StickyMacAllowListLimit = &v
}

// GetStormControlEnabled returns the StormControlEnabled field value if set, zero value otherwise.
func (o *InlineObject232) GetStormControlEnabled() bool {
	if o == nil || isNil(o.StormControlEnabled) {
		var ret bool
		return ret
	}
	return *o.StormControlEnabled
}

// GetStormControlEnabledOk returns a tuple with the StormControlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetStormControlEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.StormControlEnabled) {
    return nil, false
	}
	return o.StormControlEnabled, true
}

// HasStormControlEnabled returns a boolean if a field has been set.
func (o *InlineObject232) HasStormControlEnabled() bool {
	if o != nil && !isNil(o.StormControlEnabled) {
		return true
	}

	return false
}

// SetStormControlEnabled gets a reference to the given bool and assigns it to the StormControlEnabled field.
func (o *InlineObject232) SetStormControlEnabled(v bool) {
	o.StormControlEnabled = &v
}

// GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field value if set, zero value otherwise.
func (o *InlineObject232) GetFlexibleStackingEnabled() bool {
	if o == nil || isNil(o.FlexibleStackingEnabled) {
		var ret bool
		return ret
	}
	return *o.FlexibleStackingEnabled
}

// GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetFlexibleStackingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.FlexibleStackingEnabled) {
    return nil, false
	}
	return o.FlexibleStackingEnabled, true
}

// HasFlexibleStackingEnabled returns a boolean if a field has been set.
func (o *InlineObject232) HasFlexibleStackingEnabled() bool {
	if o != nil && !isNil(o.FlexibleStackingEnabled) {
		return true
	}

	return false
}

// SetFlexibleStackingEnabled gets a reference to the given bool and assigns it to the FlexibleStackingEnabled field.
func (o *InlineObject232) SetFlexibleStackingEnabled(v bool) {
	o.FlexibleStackingEnabled = &v
}

// GetDaiTrusted returns the DaiTrusted field value if set, zero value otherwise.
func (o *InlineObject232) GetDaiTrusted() bool {
	if o == nil || isNil(o.DaiTrusted) {
		var ret bool
		return ret
	}
	return *o.DaiTrusted
}

// GetDaiTrustedOk returns a tuple with the DaiTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetDaiTrustedOk() (*bool, bool) {
	if o == nil || isNil(o.DaiTrusted) {
    return nil, false
	}
	return o.DaiTrusted, true
}

// HasDaiTrusted returns a boolean if a field has been set.
func (o *InlineObject232) HasDaiTrusted() bool {
	if o != nil && !isNil(o.DaiTrusted) {
		return true
	}

	return false
}

// SetDaiTrusted gets a reference to the given bool and assigns it to the DaiTrusted field.
func (o *InlineObject232) SetDaiTrusted(v bool) {
	o.DaiTrusted = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *InlineObject232) GetProfile() DevicesSerialSwitchPortsProfile {
	if o == nil || isNil(o.Profile) {
		var ret DevicesSerialSwitchPortsProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetProfileOk() (*DevicesSerialSwitchPortsProfile, bool) {
	if o == nil || isNil(o.Profile) {
    return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *InlineObject232) HasProfile() bool {
	if o != nil && !isNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DevicesSerialSwitchPortsProfile and assigns it to the Profile field.
func (o *InlineObject232) SetProfile(v DevicesSerialSwitchPortsProfile) {
	o.Profile = &v
}

// GetDot3az returns the Dot3az field value if set, zero value otherwise.
func (o *InlineObject232) GetDot3az() OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az {
	if o == nil || isNil(o.Dot3az) {
		var ret OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az
		return ret
	}
	return *o.Dot3az
}

// GetDot3azOk returns a tuple with the Dot3az field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject232) GetDot3azOk() (*OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az, bool) {
	if o == nil || isNil(o.Dot3az) {
    return nil, false
	}
	return o.Dot3az, true
}

// HasDot3az returns a boolean if a field has been set.
func (o *InlineObject232) HasDot3az() bool {
	if o != nil && !isNil(o.Dot3az) {
		return true
	}

	return false
}

// SetDot3az gets a reference to the given OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az and assigns it to the Dot3az field.
func (o *InlineObject232) SetDot3az(v OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az) {
	o.Dot3az = &v
}

func (o InlineObject232) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if !isNil(o.IsolationEnabled) {
		toSerialize["isolationEnabled"] = o.IsolationEnabled
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
	if !isNil(o.PortScheduleId) {
		toSerialize["portScheduleId"] = o.PortScheduleId
	}
	if !isNil(o.Udld) {
		toSerialize["udld"] = o.Udld
	}
	if !isNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !isNil(o.AccessPolicyNumber) {
		toSerialize["accessPolicyNumber"] = o.AccessPolicyNumber
	}
	if !isNil(o.MacAllowList) {
		toSerialize["macAllowList"] = o.MacAllowList
	}
	if !isNil(o.StickyMacAllowList) {
		toSerialize["stickyMacAllowList"] = o.StickyMacAllowList
	}
	if !isNil(o.StickyMacAllowListLimit) {
		toSerialize["stickyMacAllowListLimit"] = o.StickyMacAllowListLimit
	}
	if !isNil(o.StormControlEnabled) {
		toSerialize["stormControlEnabled"] = o.StormControlEnabled
	}
	if !isNil(o.FlexibleStackingEnabled) {
		toSerialize["flexibleStackingEnabled"] = o.FlexibleStackingEnabled
	}
	if !isNil(o.DaiTrusted) {
		toSerialize["daiTrusted"] = o.DaiTrusted
	}
	if !isNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !isNil(o.Dot3az) {
		toSerialize["dot3az"] = o.Dot3az
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject232 struct {
	value *InlineObject232
	isSet bool
}

func (v NullableInlineObject232) Get() *InlineObject232 {
	return v.value
}

func (v *NullableInlineObject232) Set(val *InlineObject232) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject232) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject232) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject232(val *InlineObject232) *NullableInlineObject232 {
	return &NullableInlineObject232{value: val, isSet: true}
}

func (v NullableInlineObject232) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject232) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


