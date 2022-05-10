/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject93 struct for InlineObject93
type InlineObject93 struct {
	// Name of the access policy
	Name *string `json:"name,omitempty"`
	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []NetworksNetworkIdSwitchAccessPoliciesRadiusServers `json:"radiusServers,omitempty"`
	Radius *NetworksNetworkIdSwitchAccessPoliciesRadius `json:"radius,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled *bool `json:"radiusTestingEnabled,omitempty"`
	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled *bool `json:"radiusCoaSupportEnabled,omitempty"`
	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled *bool `json:"radiusAccountingEnabled,omitempty"`
	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`
	// Can be either `\"\"`, which means `None` on Dashboard, or `\"11\"`, which means `Filter-Id` on Dashboard and will use Group Policy ACLs when supported (firmware 14+)
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
	// Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenEnabled *bool `json:"urlRedirectWalledGardenEnabled,omitempty"`
	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges,omitempty"`
}

// NewInlineObject93 instantiates a new InlineObject93 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject93() *InlineObject93 {
	this := InlineObject93{}
	return &this
}

// NewInlineObject93WithDefaults instantiates a new InlineObject93 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject93WithDefaults() *InlineObject93 {
	this := InlineObject93{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject93) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject93) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject93) SetName(v string) {
	o.Name = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *InlineObject93) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusServers {
	if o == nil || o.RadiusServers == nil {
		var ret []NetworksNetworkIdSwitchAccessPoliciesRadiusServers
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusServersOk() ([]NetworksNetworkIdSwitchAccessPoliciesRadiusServers, bool) {
	if o == nil || o.RadiusServers == nil {
		return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *InlineObject93) HasRadiusServers() bool {
	if o != nil && o.RadiusServers != nil {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []NetworksNetworkIdSwitchAccessPoliciesRadiusServers and assigns it to the RadiusServers field.
func (o *InlineObject93) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusServers) {
	o.RadiusServers = v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *InlineObject93) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius {
	if o == nil || o.Radius == nil {
		var ret NetworksNetworkIdSwitchAccessPoliciesRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *InlineObject93) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesRadius and assigns it to the Radius field.
func (o *InlineObject93) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius) {
	o.Radius = &v
}

// GetRadiusTestingEnabled returns the RadiusTestingEnabled field value if set, zero value otherwise.
func (o *InlineObject93) GetRadiusTestingEnabled() bool {
	if o == nil || o.RadiusTestingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RadiusTestingEnabled
}

// GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusTestingEnabledOk() (*bool, bool) {
	if o == nil || o.RadiusTestingEnabled == nil {
		return nil, false
	}
	return o.RadiusTestingEnabled, true
}

// HasRadiusTestingEnabled returns a boolean if a field has been set.
func (o *InlineObject93) HasRadiusTestingEnabled() bool {
	if o != nil && o.RadiusTestingEnabled != nil {
		return true
	}

	return false
}

// SetRadiusTestingEnabled gets a reference to the given bool and assigns it to the RadiusTestingEnabled field.
func (o *InlineObject93) SetRadiusTestingEnabled(v bool) {
	o.RadiusTestingEnabled = &v
}

// GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field value if set, zero value otherwise.
func (o *InlineObject93) GetRadiusCoaSupportEnabled() bool {
	if o == nil || o.RadiusCoaSupportEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RadiusCoaSupportEnabled
}

// GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusCoaSupportEnabledOk() (*bool, bool) {
	if o == nil || o.RadiusCoaSupportEnabled == nil {
		return nil, false
	}
	return o.RadiusCoaSupportEnabled, true
}

// HasRadiusCoaSupportEnabled returns a boolean if a field has been set.
func (o *InlineObject93) HasRadiusCoaSupportEnabled() bool {
	if o != nil && o.RadiusCoaSupportEnabled != nil {
		return true
	}

	return false
}

// SetRadiusCoaSupportEnabled gets a reference to the given bool and assigns it to the RadiusCoaSupportEnabled field.
func (o *InlineObject93) SetRadiusCoaSupportEnabled(v bool) {
	o.RadiusCoaSupportEnabled = &v
}

// GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field value if set, zero value otherwise.
func (o *InlineObject93) GetRadiusAccountingEnabled() bool {
	if o == nil || o.RadiusAccountingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RadiusAccountingEnabled
}

// GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusAccountingEnabledOk() (*bool, bool) {
	if o == nil || o.RadiusAccountingEnabled == nil {
		return nil, false
	}
	return o.RadiusAccountingEnabled, true
}

// HasRadiusAccountingEnabled returns a boolean if a field has been set.
func (o *InlineObject93) HasRadiusAccountingEnabled() bool {
	if o != nil && o.RadiusAccountingEnabled != nil {
		return true
	}

	return false
}

// SetRadiusAccountingEnabled gets a reference to the given bool and assigns it to the RadiusAccountingEnabled field.
func (o *InlineObject93) SetRadiusAccountingEnabled(v bool) {
	o.RadiusAccountingEnabled = &v
}

// GetRadiusAccountingServers returns the RadiusAccountingServers field value if set, zero value otherwise.
func (o *InlineObject93) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers {
	if o == nil || o.RadiusAccountingServers == nil {
		var ret []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers
		return ret
	}
	return o.RadiusAccountingServers
}

// GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusAccountingServersOk() ([]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers, bool) {
	if o == nil || o.RadiusAccountingServers == nil {
		return nil, false
	}
	return o.RadiusAccountingServers, true
}

// HasRadiusAccountingServers returns a boolean if a field has been set.
func (o *InlineObject93) HasRadiusAccountingServers() bool {
	if o != nil && o.RadiusAccountingServers != nil {
		return true
	}

	return false
}

// SetRadiusAccountingServers gets a reference to the given []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers and assigns it to the RadiusAccountingServers field.
func (o *InlineObject93) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers) {
	o.RadiusAccountingServers = v
}

// GetRadiusGroupAttribute returns the RadiusGroupAttribute field value if set, zero value otherwise.
func (o *InlineObject93) GetRadiusGroupAttribute() string {
	if o == nil || o.RadiusGroupAttribute == nil {
		var ret string
		return ret
	}
	return *o.RadiusGroupAttribute
}

// GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetRadiusGroupAttributeOk() (*string, bool) {
	if o == nil || o.RadiusGroupAttribute == nil {
		return nil, false
	}
	return o.RadiusGroupAttribute, true
}

// HasRadiusGroupAttribute returns a boolean if a field has been set.
func (o *InlineObject93) HasRadiusGroupAttribute() bool {
	if o != nil && o.RadiusGroupAttribute != nil {
		return true
	}

	return false
}

// SetRadiusGroupAttribute gets a reference to the given string and assigns it to the RadiusGroupAttribute field.
func (o *InlineObject93) SetRadiusGroupAttribute(v string) {
	o.RadiusGroupAttribute = &v
}

// GetHostMode returns the HostMode field value if set, zero value otherwise.
func (o *InlineObject93) GetHostMode() string {
	if o == nil || o.HostMode == nil {
		var ret string
		return ret
	}
	return *o.HostMode
}

// GetHostModeOk returns a tuple with the HostMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetHostModeOk() (*string, bool) {
	if o == nil || o.HostMode == nil {
		return nil, false
	}
	return o.HostMode, true
}

// HasHostMode returns a boolean if a field has been set.
func (o *InlineObject93) HasHostMode() bool {
	if o != nil && o.HostMode != nil {
		return true
	}

	return false
}

// SetHostMode gets a reference to the given string and assigns it to the HostMode field.
func (o *InlineObject93) SetHostMode(v string) {
	o.HostMode = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *InlineObject93) GetAccessPolicyType() string {
	if o == nil || o.AccessPolicyType == nil {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || o.AccessPolicyType == nil {
		return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *InlineObject93) HasAccessPolicyType() bool {
	if o != nil && o.AccessPolicyType != nil {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *InlineObject93) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field value if set, zero value otherwise.
func (o *InlineObject93) GetIncreaseAccessSpeed() bool {
	if o == nil || o.IncreaseAccessSpeed == nil {
		var ret bool
		return ret
	}
	return *o.IncreaseAccessSpeed
}

// GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetIncreaseAccessSpeedOk() (*bool, bool) {
	if o == nil || o.IncreaseAccessSpeed == nil {
		return nil, false
	}
	return o.IncreaseAccessSpeed, true
}

// HasIncreaseAccessSpeed returns a boolean if a field has been set.
func (o *InlineObject93) HasIncreaseAccessSpeed() bool {
	if o != nil && o.IncreaseAccessSpeed != nil {
		return true
	}

	return false
}

// SetIncreaseAccessSpeed gets a reference to the given bool and assigns it to the IncreaseAccessSpeed field.
func (o *InlineObject93) SetIncreaseAccessSpeed(v bool) {
	o.IncreaseAccessSpeed = &v
}

// GetGuestVlanId returns the GuestVlanId field value if set, zero value otherwise.
func (o *InlineObject93) GetGuestVlanId() int32 {
	if o == nil || o.GuestVlanId == nil {
		var ret int32
		return ret
	}
	return *o.GuestVlanId
}

// GetGuestVlanIdOk returns a tuple with the GuestVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetGuestVlanIdOk() (*int32, bool) {
	if o == nil || o.GuestVlanId == nil {
		return nil, false
	}
	return o.GuestVlanId, true
}

// HasGuestVlanId returns a boolean if a field has been set.
func (o *InlineObject93) HasGuestVlanId() bool {
	if o != nil && o.GuestVlanId != nil {
		return true
	}

	return false
}

// SetGuestVlanId gets a reference to the given int32 and assigns it to the GuestVlanId field.
func (o *InlineObject93) SetGuestVlanId(v int32) {
	o.GuestVlanId = &v
}

// GetDot1x returns the Dot1x field value if set, zero value otherwise.
func (o *InlineObject93) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x {
	if o == nil || o.Dot1x == nil {
		var ret NetworksNetworkIdSwitchAccessPoliciesDot1x
		return ret
	}
	return *o.Dot1x
}

// GetDot1xOk returns a tuple with the Dot1x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool) {
	if o == nil || o.Dot1x == nil {
		return nil, false
	}
	return o.Dot1x, true
}

// HasDot1x returns a boolean if a field has been set.
func (o *InlineObject93) HasDot1x() bool {
	if o != nil && o.Dot1x != nil {
		return true
	}

	return false
}

// SetDot1x gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesDot1x and assigns it to the Dot1x field.
func (o *InlineObject93) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x) {
	o.Dot1x = &v
}

// GetVoiceVlanClients returns the VoiceVlanClients field value if set, zero value otherwise.
func (o *InlineObject93) GetVoiceVlanClients() bool {
	if o == nil || o.VoiceVlanClients == nil {
		var ret bool
		return ret
	}
	return *o.VoiceVlanClients
}

// GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetVoiceVlanClientsOk() (*bool, bool) {
	if o == nil || o.VoiceVlanClients == nil {
		return nil, false
	}
	return o.VoiceVlanClients, true
}

// HasVoiceVlanClients returns a boolean if a field has been set.
func (o *InlineObject93) HasVoiceVlanClients() bool {
	if o != nil && o.VoiceVlanClients != nil {
		return true
	}

	return false
}

// SetVoiceVlanClients gets a reference to the given bool and assigns it to the VoiceVlanClients field.
func (o *InlineObject93) SetVoiceVlanClients(v bool) {
	o.VoiceVlanClients = &v
}

// GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field value if set, zero value otherwise.
func (o *InlineObject93) GetUrlRedirectWalledGardenEnabled() bool {
	if o == nil || o.UrlRedirectWalledGardenEnabled == nil {
		var ret bool
		return ret
	}
	return *o.UrlRedirectWalledGardenEnabled
}

// GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool) {
	if o == nil || o.UrlRedirectWalledGardenEnabled == nil {
		return nil, false
	}
	return o.UrlRedirectWalledGardenEnabled, true
}

// HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.
func (o *InlineObject93) HasUrlRedirectWalledGardenEnabled() bool {
	if o != nil && o.UrlRedirectWalledGardenEnabled != nil {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenEnabled gets a reference to the given bool and assigns it to the UrlRedirectWalledGardenEnabled field.
func (o *InlineObject93) SetUrlRedirectWalledGardenEnabled(v bool) {
	o.UrlRedirectWalledGardenEnabled = &v
}

// GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field value if set, zero value otherwise.
func (o *InlineObject93) GetUrlRedirectWalledGardenRanges() []string {
	if o == nil || o.UrlRedirectWalledGardenRanges == nil {
		var ret []string
		return ret
	}
	return o.UrlRedirectWalledGardenRanges
}

// GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject93) GetUrlRedirectWalledGardenRangesOk() ([]string, bool) {
	if o == nil || o.UrlRedirectWalledGardenRanges == nil {
		return nil, false
	}
	return o.UrlRedirectWalledGardenRanges, true
}

// HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.
func (o *InlineObject93) HasUrlRedirectWalledGardenRanges() bool {
	if o != nil && o.UrlRedirectWalledGardenRanges != nil {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenRanges gets a reference to the given []string and assigns it to the UrlRedirectWalledGardenRanges field.
func (o *InlineObject93) SetUrlRedirectWalledGardenRanges(v []string) {
	o.UrlRedirectWalledGardenRanges = v
}

func (o InlineObject93) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RadiusServers != nil {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	if o.RadiusTestingEnabled != nil {
		toSerialize["radiusTestingEnabled"] = o.RadiusTestingEnabled
	}
	if o.RadiusCoaSupportEnabled != nil {
		toSerialize["radiusCoaSupportEnabled"] = o.RadiusCoaSupportEnabled
	}
	if o.RadiusAccountingEnabled != nil {
		toSerialize["radiusAccountingEnabled"] = o.RadiusAccountingEnabled
	}
	if o.RadiusAccountingServers != nil {
		toSerialize["radiusAccountingServers"] = o.RadiusAccountingServers
	}
	if o.RadiusGroupAttribute != nil {
		toSerialize["radiusGroupAttribute"] = o.RadiusGroupAttribute
	}
	if o.HostMode != nil {
		toSerialize["hostMode"] = o.HostMode
	}
	if o.AccessPolicyType != nil {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if o.IncreaseAccessSpeed != nil {
		toSerialize["increaseAccessSpeed"] = o.IncreaseAccessSpeed
	}
	if o.GuestVlanId != nil {
		toSerialize["guestVlanId"] = o.GuestVlanId
	}
	if o.Dot1x != nil {
		toSerialize["dot1x"] = o.Dot1x
	}
	if o.VoiceVlanClients != nil {
		toSerialize["voiceVlanClients"] = o.VoiceVlanClients
	}
	if o.UrlRedirectWalledGardenEnabled != nil {
		toSerialize["urlRedirectWalledGardenEnabled"] = o.UrlRedirectWalledGardenEnabled
	}
	if o.UrlRedirectWalledGardenRanges != nil {
		toSerialize["urlRedirectWalledGardenRanges"] = o.UrlRedirectWalledGardenRanges
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject93 struct {
	value *InlineObject93
	isSet bool
}

func (v NullableInlineObject93) Get() *InlineObject93 {
	return v.value
}

func (v *NullableInlineObject93) Set(val *InlineObject93) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject93) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject93) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject93(val *InlineObject93) *NullableInlineObject93 {
	return &NullableInlineObject93{value: val, isSet: true}
}

func (v NullableInlineObject93) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject93) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


