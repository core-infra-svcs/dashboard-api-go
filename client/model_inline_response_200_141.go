/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200141 struct for InlineResponse200141
type InlineResponse200141 struct {
	// Name of the access policy
	Name *string `json:"name,omitempty"`
	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []NetworksNetworkIdSwitchAccessPoliciesRadiusServers `json:"radiusServers,omitempty"`
	Radius *NetworksNetworkIdSwitchAccessPoliciesRadius `json:"radius,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestPortBouncing *bool `json:"guestPortBouncing,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled *bool `json:"radiusTestingEnabled,omitempty"`
	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled *bool `json:"radiusCoaSupportEnabled,omitempty"`
	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled *bool `json:"radiusAccountingEnabled,omitempty"`
	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`
	// Acceptable values are `\"\"` for None, or `\"11\"` for Group Policies ACL
	RadiusGroupAttribute *string `json:"radiusGroupAttribute,omitempty"`
	// Choose the Host Mode for the access policy.
	HostMode *string `json:"hostMode,omitempty"`
	// Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	IncreaseAccessSpeed *bool `json:"increaseAccessSpeed,omitempty"`
	// ID for the guest VLAN allow unauthorized devices access to limited network resources
	GuestVlanId *int32 `json:"guestVlanId,omitempty"`
	Dot1x *NetworksNetworkIdSwitchAccessPoliciesDot1x `json:"dot1x,omitempty"`
	// CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
	VoiceVlanClients *bool `json:"voiceVlanClients,omitempty"`
	// Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenEnabled *bool `json:"urlRedirectWalledGardenEnabled,omitempty"`
	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges,omitempty"`
	Counts *NetworksNetworkIdSwitchAccessPoliciesCounts `json:"counts,omitempty"`
}

// NewInlineResponse200141 instantiates a new InlineResponse200141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200141() *InlineResponse200141 {
	this := InlineResponse200141{}
	return &this
}

// NewInlineResponse200141WithDefaults instantiates a new InlineResponse200141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200141WithDefaults() *InlineResponse200141 {
	this := InlineResponse200141{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200141) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200141) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200141) SetName(v string) {
	o.Name = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusServers {
	if o == nil || isNil(o.RadiusServers) {
		var ret []NetworksNetworkIdSwitchAccessPoliciesRadiusServers
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusServersOk() ([]NetworksNetworkIdSwitchAccessPoliciesRadiusServers, bool) {
	if o == nil || isNil(o.RadiusServers) {
    return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadiusServers() bool {
	if o != nil && !isNil(o.RadiusServers) {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []NetworksNetworkIdSwitchAccessPoliciesRadiusServers and assigns it to the RadiusServers field.
func (o *InlineResponse200141) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusServers) {
	o.RadiusServers = v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius {
	if o == nil || isNil(o.Radius) {
		var ret NetworksNetworkIdSwitchAccessPoliciesRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius, bool) {
	if o == nil || isNil(o.Radius) {
    return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadius() bool {
	if o != nil && !isNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesRadius and assigns it to the Radius field.
func (o *InlineResponse200141) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius) {
	o.Radius = &v
}

// GetGuestPortBouncing returns the GuestPortBouncing field value if set, zero value otherwise.
func (o *InlineResponse200141) GetGuestPortBouncing() bool {
	if o == nil || isNil(o.GuestPortBouncing) {
		var ret bool
		return ret
	}
	return *o.GuestPortBouncing
}

// GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetGuestPortBouncingOk() (*bool, bool) {
	if o == nil || isNil(o.GuestPortBouncing) {
    return nil, false
	}
	return o.GuestPortBouncing, true
}

// HasGuestPortBouncing returns a boolean if a field has been set.
func (o *InlineResponse200141) HasGuestPortBouncing() bool {
	if o != nil && !isNil(o.GuestPortBouncing) {
		return true
	}

	return false
}

// SetGuestPortBouncing gets a reference to the given bool and assigns it to the GuestPortBouncing field.
func (o *InlineResponse200141) SetGuestPortBouncing(v bool) {
	o.GuestPortBouncing = &v
}

// GetRadiusTestingEnabled returns the RadiusTestingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadiusTestingEnabled() bool {
	if o == nil || isNil(o.RadiusTestingEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusTestingEnabled
}

// GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusTestingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RadiusTestingEnabled) {
    return nil, false
	}
	return o.RadiusTestingEnabled, true
}

// HasRadiusTestingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadiusTestingEnabled() bool {
	if o != nil && !isNil(o.RadiusTestingEnabled) {
		return true
	}

	return false
}

// SetRadiusTestingEnabled gets a reference to the given bool and assigns it to the RadiusTestingEnabled field.
func (o *InlineResponse200141) SetRadiusTestingEnabled(v bool) {
	o.RadiusTestingEnabled = &v
}

// GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadiusCoaSupportEnabled() bool {
	if o == nil || isNil(o.RadiusCoaSupportEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusCoaSupportEnabled
}

// GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusCoaSupportEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RadiusCoaSupportEnabled) {
    return nil, false
	}
	return o.RadiusCoaSupportEnabled, true
}

// HasRadiusCoaSupportEnabled returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadiusCoaSupportEnabled() bool {
	if o != nil && !isNil(o.RadiusCoaSupportEnabled) {
		return true
	}

	return false
}

// SetRadiusCoaSupportEnabled gets a reference to the given bool and assigns it to the RadiusCoaSupportEnabled field.
func (o *InlineResponse200141) SetRadiusCoaSupportEnabled(v bool) {
	o.RadiusCoaSupportEnabled = &v
}

// GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadiusAccountingEnabled() bool {
	if o == nil || isNil(o.RadiusAccountingEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusAccountingEnabled
}

// GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusAccountingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RadiusAccountingEnabled) {
    return nil, false
	}
	return o.RadiusAccountingEnabled, true
}

// HasRadiusAccountingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadiusAccountingEnabled() bool {
	if o != nil && !isNil(o.RadiusAccountingEnabled) {
		return true
	}

	return false
}

// SetRadiusAccountingEnabled gets a reference to the given bool and assigns it to the RadiusAccountingEnabled field.
func (o *InlineResponse200141) SetRadiusAccountingEnabled(v bool) {
	o.RadiusAccountingEnabled = &v
}

// GetRadiusAccountingServers returns the RadiusAccountingServers field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers {
	if o == nil || isNil(o.RadiusAccountingServers) {
		var ret []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers
		return ret
	}
	return o.RadiusAccountingServers
}

// GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusAccountingServersOk() ([]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers, bool) {
	if o == nil || isNil(o.RadiusAccountingServers) {
    return nil, false
	}
	return o.RadiusAccountingServers, true
}

// HasRadiusAccountingServers returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadiusAccountingServers() bool {
	if o != nil && !isNil(o.RadiusAccountingServers) {
		return true
	}

	return false
}

// SetRadiusAccountingServers gets a reference to the given []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers and assigns it to the RadiusAccountingServers field.
func (o *InlineResponse200141) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers) {
	o.RadiusAccountingServers = v
}

// GetRadiusGroupAttribute returns the RadiusGroupAttribute field value if set, zero value otherwise.
func (o *InlineResponse200141) GetRadiusGroupAttribute() string {
	if o == nil || isNil(o.RadiusGroupAttribute) {
		var ret string
		return ret
	}
	return *o.RadiusGroupAttribute
}

// GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetRadiusGroupAttributeOk() (*string, bool) {
	if o == nil || isNil(o.RadiusGroupAttribute) {
    return nil, false
	}
	return o.RadiusGroupAttribute, true
}

// HasRadiusGroupAttribute returns a boolean if a field has been set.
func (o *InlineResponse200141) HasRadiusGroupAttribute() bool {
	if o != nil && !isNil(o.RadiusGroupAttribute) {
		return true
	}

	return false
}

// SetRadiusGroupAttribute gets a reference to the given string and assigns it to the RadiusGroupAttribute field.
func (o *InlineResponse200141) SetRadiusGroupAttribute(v string) {
	o.RadiusGroupAttribute = &v
}

// GetHostMode returns the HostMode field value if set, zero value otherwise.
func (o *InlineResponse200141) GetHostMode() string {
	if o == nil || isNil(o.HostMode) {
		var ret string
		return ret
	}
	return *o.HostMode
}

// GetHostModeOk returns a tuple with the HostMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetHostModeOk() (*string, bool) {
	if o == nil || isNil(o.HostMode) {
    return nil, false
	}
	return o.HostMode, true
}

// HasHostMode returns a boolean if a field has been set.
func (o *InlineResponse200141) HasHostMode() bool {
	if o != nil && !isNil(o.HostMode) {
		return true
	}

	return false
}

// SetHostMode gets a reference to the given string and assigns it to the HostMode field.
func (o *InlineResponse200141) SetHostMode(v string) {
	o.HostMode = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *InlineResponse200141) GetAccessPolicyType() string {
	if o == nil || isNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessPolicyType) {
    return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *InlineResponse200141) HasAccessPolicyType() bool {
	if o != nil && !isNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *InlineResponse200141) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field value if set, zero value otherwise.
func (o *InlineResponse200141) GetIncreaseAccessSpeed() bool {
	if o == nil || isNil(o.IncreaseAccessSpeed) {
		var ret bool
		return ret
	}
	return *o.IncreaseAccessSpeed
}

// GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetIncreaseAccessSpeedOk() (*bool, bool) {
	if o == nil || isNil(o.IncreaseAccessSpeed) {
    return nil, false
	}
	return o.IncreaseAccessSpeed, true
}

// HasIncreaseAccessSpeed returns a boolean if a field has been set.
func (o *InlineResponse200141) HasIncreaseAccessSpeed() bool {
	if o != nil && !isNil(o.IncreaseAccessSpeed) {
		return true
	}

	return false
}

// SetIncreaseAccessSpeed gets a reference to the given bool and assigns it to the IncreaseAccessSpeed field.
func (o *InlineResponse200141) SetIncreaseAccessSpeed(v bool) {
	o.IncreaseAccessSpeed = &v
}

// GetGuestVlanId returns the GuestVlanId field value if set, zero value otherwise.
func (o *InlineResponse200141) GetGuestVlanId() int32 {
	if o == nil || isNil(o.GuestVlanId) {
		var ret int32
		return ret
	}
	return *o.GuestVlanId
}

// GetGuestVlanIdOk returns a tuple with the GuestVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetGuestVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.GuestVlanId) {
    return nil, false
	}
	return o.GuestVlanId, true
}

// HasGuestVlanId returns a boolean if a field has been set.
func (o *InlineResponse200141) HasGuestVlanId() bool {
	if o != nil && !isNil(o.GuestVlanId) {
		return true
	}

	return false
}

// SetGuestVlanId gets a reference to the given int32 and assigns it to the GuestVlanId field.
func (o *InlineResponse200141) SetGuestVlanId(v int32) {
	o.GuestVlanId = &v
}

// GetDot1x returns the Dot1x field value if set, zero value otherwise.
func (o *InlineResponse200141) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x {
	if o == nil || isNil(o.Dot1x) {
		var ret NetworksNetworkIdSwitchAccessPoliciesDot1x
		return ret
	}
	return *o.Dot1x
}

// GetDot1xOk returns a tuple with the Dot1x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool) {
	if o == nil || isNil(o.Dot1x) {
    return nil, false
	}
	return o.Dot1x, true
}

// HasDot1x returns a boolean if a field has been set.
func (o *InlineResponse200141) HasDot1x() bool {
	if o != nil && !isNil(o.Dot1x) {
		return true
	}

	return false
}

// SetDot1x gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesDot1x and assigns it to the Dot1x field.
func (o *InlineResponse200141) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x) {
	o.Dot1x = &v
}

// GetVoiceVlanClients returns the VoiceVlanClients field value if set, zero value otherwise.
func (o *InlineResponse200141) GetVoiceVlanClients() bool {
	if o == nil || isNil(o.VoiceVlanClients) {
		var ret bool
		return ret
	}
	return *o.VoiceVlanClients
}

// GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetVoiceVlanClientsOk() (*bool, bool) {
	if o == nil || isNil(o.VoiceVlanClients) {
    return nil, false
	}
	return o.VoiceVlanClients, true
}

// HasVoiceVlanClients returns a boolean if a field has been set.
func (o *InlineResponse200141) HasVoiceVlanClients() bool {
	if o != nil && !isNil(o.VoiceVlanClients) {
		return true
	}

	return false
}

// SetVoiceVlanClients gets a reference to the given bool and assigns it to the VoiceVlanClients field.
func (o *InlineResponse200141) SetVoiceVlanClients(v bool) {
	o.VoiceVlanClients = &v
}

// GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field value if set, zero value otherwise.
func (o *InlineResponse200141) GetUrlRedirectWalledGardenEnabled() bool {
	if o == nil || isNil(o.UrlRedirectWalledGardenEnabled) {
		var ret bool
		return ret
	}
	return *o.UrlRedirectWalledGardenEnabled
}

// GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.UrlRedirectWalledGardenEnabled) {
    return nil, false
	}
	return o.UrlRedirectWalledGardenEnabled, true
}

// HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.
func (o *InlineResponse200141) HasUrlRedirectWalledGardenEnabled() bool {
	if o != nil && !isNil(o.UrlRedirectWalledGardenEnabled) {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenEnabled gets a reference to the given bool and assigns it to the UrlRedirectWalledGardenEnabled field.
func (o *InlineResponse200141) SetUrlRedirectWalledGardenEnabled(v bool) {
	o.UrlRedirectWalledGardenEnabled = &v
}

// GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field value if set, zero value otherwise.
func (o *InlineResponse200141) GetUrlRedirectWalledGardenRanges() []string {
	if o == nil || isNil(o.UrlRedirectWalledGardenRanges) {
		var ret []string
		return ret
	}
	return o.UrlRedirectWalledGardenRanges
}

// GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetUrlRedirectWalledGardenRangesOk() ([]string, bool) {
	if o == nil || isNil(o.UrlRedirectWalledGardenRanges) {
    return nil, false
	}
	return o.UrlRedirectWalledGardenRanges, true
}

// HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.
func (o *InlineResponse200141) HasUrlRedirectWalledGardenRanges() bool {
	if o != nil && !isNil(o.UrlRedirectWalledGardenRanges) {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenRanges gets a reference to the given []string and assigns it to the UrlRedirectWalledGardenRanges field.
func (o *InlineResponse200141) SetUrlRedirectWalledGardenRanges(v []string) {
	o.UrlRedirectWalledGardenRanges = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200141) GetCounts() NetworksNetworkIdSwitchAccessPoliciesCounts {
	if o == nil || isNil(o.Counts) {
		var ret NetworksNetworkIdSwitchAccessPoliciesCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetCountsOk() (*NetworksNetworkIdSwitchAccessPoliciesCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200141) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesCounts and assigns it to the Counts field.
func (o *InlineResponse200141) SetCounts(v NetworksNetworkIdSwitchAccessPoliciesCounts) {
	o.Counts = &v
}

func (o InlineResponse200141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.RadiusServers) {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !isNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !isNil(o.GuestPortBouncing) {
		toSerialize["guestPortBouncing"] = o.GuestPortBouncing
	}
	if !isNil(o.RadiusTestingEnabled) {
		toSerialize["radiusTestingEnabled"] = o.RadiusTestingEnabled
	}
	if !isNil(o.RadiusCoaSupportEnabled) {
		toSerialize["radiusCoaSupportEnabled"] = o.RadiusCoaSupportEnabled
	}
	if !isNil(o.RadiusAccountingEnabled) {
		toSerialize["radiusAccountingEnabled"] = o.RadiusAccountingEnabled
	}
	if !isNil(o.RadiusAccountingServers) {
		toSerialize["radiusAccountingServers"] = o.RadiusAccountingServers
	}
	if !isNil(o.RadiusGroupAttribute) {
		toSerialize["radiusGroupAttribute"] = o.RadiusGroupAttribute
	}
	if !isNil(o.HostMode) {
		toSerialize["hostMode"] = o.HostMode
	}
	if !isNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !isNil(o.IncreaseAccessSpeed) {
		toSerialize["increaseAccessSpeed"] = o.IncreaseAccessSpeed
	}
	if !isNil(o.GuestVlanId) {
		toSerialize["guestVlanId"] = o.GuestVlanId
	}
	if !isNil(o.Dot1x) {
		toSerialize["dot1x"] = o.Dot1x
	}
	if !isNil(o.VoiceVlanClients) {
		toSerialize["voiceVlanClients"] = o.VoiceVlanClients
	}
	if !isNil(o.UrlRedirectWalledGardenEnabled) {
		toSerialize["urlRedirectWalledGardenEnabled"] = o.UrlRedirectWalledGardenEnabled
	}
	if !isNil(o.UrlRedirectWalledGardenRanges) {
		toSerialize["urlRedirectWalledGardenRanges"] = o.UrlRedirectWalledGardenRanges
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200141 struct {
	value *InlineResponse200141
	isSet bool
}

func (v NullableInlineResponse200141) Get() *InlineResponse200141 {
	return v.value
}

func (v *NullableInlineResponse200141) Set(val *InlineResponse200141) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200141) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200141(val *InlineResponse200141) *NullableInlineResponse200141 {
	return &NullableInlineResponse200141{value: val, isSet: true}
}

func (v NullableInlineResponse200141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


