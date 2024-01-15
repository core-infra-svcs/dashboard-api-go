/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject120 struct for InlineObject120
type InlineObject120 struct {
	// Name of the access policy
	Name string `json:"name"`
	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 `json:"radiusServers"`
	Radius *NetworksNetworkIdSwitchAccessPoliciesRadius `json:"radius,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestPortBouncing *bool `json:"guestPortBouncing,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled bool `json:"radiusTestingEnabled"`
	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled bool `json:"radiusCoaSupportEnabled"`
	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled"`
	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 `json:"radiusAccountingServers,omitempty"`
	// Acceptable values are `\"\"` for None, or `\"11\"` for Group Policies ACL
	RadiusGroupAttribute *string `json:"radiusGroupAttribute,omitempty"`
	// Choose the Host Mode for the access policy.
	HostMode string `json:"hostMode"`
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
	UrlRedirectWalledGardenEnabled bool `json:"urlRedirectWalledGardenEnabled"`
	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges,omitempty"`
}

// NewInlineObject120 instantiates a new InlineObject120 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject120(name string, radiusServers []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1, radiusTestingEnabled bool, radiusCoaSupportEnabled bool, radiusAccountingEnabled bool, hostMode string, urlRedirectWalledGardenEnabled bool) *InlineObject120 {
	this := InlineObject120{}
	this.Name = name
	this.RadiusServers = radiusServers
	this.RadiusTestingEnabled = radiusTestingEnabled
	this.RadiusCoaSupportEnabled = radiusCoaSupportEnabled
	this.RadiusAccountingEnabled = radiusAccountingEnabled
	this.HostMode = hostMode
	this.UrlRedirectWalledGardenEnabled = urlRedirectWalledGardenEnabled
	return &this
}

// NewInlineObject120WithDefaults instantiates a new InlineObject120 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject120WithDefaults() *InlineObject120 {
	this := InlineObject120{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject120) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject120) SetName(v string) {
	o.Name = v
}

// GetRadiusServers returns the RadiusServers field value
func (o *InlineObject120) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1 {
	if o == nil {
		var ret []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1
		return ret
	}

	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusServersOk() ([]NetworksNetworkIdSwitchAccessPoliciesRadiusServers1, bool) {
	if o == nil {
    return nil, false
	}
	return o.RadiusServers, true
}

// SetRadiusServers sets field value
func (o *InlineObject120) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1) {
	o.RadiusServers = v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *InlineObject120) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius {
	if o == nil || isNil(o.Radius) {
		var ret NetworksNetworkIdSwitchAccessPoliciesRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius, bool) {
	if o == nil || isNil(o.Radius) {
    return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *InlineObject120) HasRadius() bool {
	if o != nil && !isNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesRadius and assigns it to the Radius field.
func (o *InlineObject120) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius) {
	o.Radius = &v
}

// GetGuestPortBouncing returns the GuestPortBouncing field value if set, zero value otherwise.
func (o *InlineObject120) GetGuestPortBouncing() bool {
	if o == nil || isNil(o.GuestPortBouncing) {
		var ret bool
		return ret
	}
	return *o.GuestPortBouncing
}

// GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetGuestPortBouncingOk() (*bool, bool) {
	if o == nil || isNil(o.GuestPortBouncing) {
    return nil, false
	}
	return o.GuestPortBouncing, true
}

// HasGuestPortBouncing returns a boolean if a field has been set.
func (o *InlineObject120) HasGuestPortBouncing() bool {
	if o != nil && !isNil(o.GuestPortBouncing) {
		return true
	}

	return false
}

// SetGuestPortBouncing gets a reference to the given bool and assigns it to the GuestPortBouncing field.
func (o *InlineObject120) SetGuestPortBouncing(v bool) {
	o.GuestPortBouncing = &v
}

// GetRadiusTestingEnabled returns the RadiusTestingEnabled field value
func (o *InlineObject120) GetRadiusTestingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RadiusTestingEnabled
}

// GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusTestingEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RadiusTestingEnabled, true
}

// SetRadiusTestingEnabled sets field value
func (o *InlineObject120) SetRadiusTestingEnabled(v bool) {
	o.RadiusTestingEnabled = v
}

// GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field value
func (o *InlineObject120) GetRadiusCoaSupportEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RadiusCoaSupportEnabled
}

// GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusCoaSupportEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RadiusCoaSupportEnabled, true
}

// SetRadiusCoaSupportEnabled sets field value
func (o *InlineObject120) SetRadiusCoaSupportEnabled(v bool) {
	o.RadiusCoaSupportEnabled = v
}

// GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field value
func (o *InlineObject120) GetRadiusAccountingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RadiusAccountingEnabled
}

// GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusAccountingEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RadiusAccountingEnabled, true
}

// SetRadiusAccountingEnabled sets field value
func (o *InlineObject120) SetRadiusAccountingEnabled(v bool) {
	o.RadiusAccountingEnabled = v
}

// GetRadiusAccountingServers returns the RadiusAccountingServers field value if set, zero value otherwise.
func (o *InlineObject120) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	if o == nil || isNil(o.RadiusAccountingServers) {
		var ret []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1
		return ret
	}
	return o.RadiusAccountingServers
}

// GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusAccountingServersOk() ([]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1, bool) {
	if o == nil || isNil(o.RadiusAccountingServers) {
    return nil, false
	}
	return o.RadiusAccountingServers, true
}

// HasRadiusAccountingServers returns a boolean if a field has been set.
func (o *InlineObject120) HasRadiusAccountingServers() bool {
	if o != nil && !isNil(o.RadiusAccountingServers) {
		return true
	}

	return false
}

// SetRadiusAccountingServers gets a reference to the given []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 and assigns it to the RadiusAccountingServers field.
func (o *InlineObject120) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) {
	o.RadiusAccountingServers = v
}

// GetRadiusGroupAttribute returns the RadiusGroupAttribute field value if set, zero value otherwise.
func (o *InlineObject120) GetRadiusGroupAttribute() string {
	if o == nil || isNil(o.RadiusGroupAttribute) {
		var ret string
		return ret
	}
	return *o.RadiusGroupAttribute
}

// GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetRadiusGroupAttributeOk() (*string, bool) {
	if o == nil || isNil(o.RadiusGroupAttribute) {
    return nil, false
	}
	return o.RadiusGroupAttribute, true
}

// HasRadiusGroupAttribute returns a boolean if a field has been set.
func (o *InlineObject120) HasRadiusGroupAttribute() bool {
	if o != nil && !isNil(o.RadiusGroupAttribute) {
		return true
	}

	return false
}

// SetRadiusGroupAttribute gets a reference to the given string and assigns it to the RadiusGroupAttribute field.
func (o *InlineObject120) SetRadiusGroupAttribute(v string) {
	o.RadiusGroupAttribute = &v
}

// GetHostMode returns the HostMode field value
func (o *InlineObject120) GetHostMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostMode
}

// GetHostModeOk returns a tuple with the HostMode field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetHostModeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HostMode, true
}

// SetHostMode sets field value
func (o *InlineObject120) SetHostMode(v string) {
	o.HostMode = v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *InlineObject120) GetAccessPolicyType() string {
	if o == nil || isNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccessPolicyType) {
    return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *InlineObject120) HasAccessPolicyType() bool {
	if o != nil && !isNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *InlineObject120) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field value if set, zero value otherwise.
func (o *InlineObject120) GetIncreaseAccessSpeed() bool {
	if o == nil || isNil(o.IncreaseAccessSpeed) {
		var ret bool
		return ret
	}
	return *o.IncreaseAccessSpeed
}

// GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetIncreaseAccessSpeedOk() (*bool, bool) {
	if o == nil || isNil(o.IncreaseAccessSpeed) {
    return nil, false
	}
	return o.IncreaseAccessSpeed, true
}

// HasIncreaseAccessSpeed returns a boolean if a field has been set.
func (o *InlineObject120) HasIncreaseAccessSpeed() bool {
	if o != nil && !isNil(o.IncreaseAccessSpeed) {
		return true
	}

	return false
}

// SetIncreaseAccessSpeed gets a reference to the given bool and assigns it to the IncreaseAccessSpeed field.
func (o *InlineObject120) SetIncreaseAccessSpeed(v bool) {
	o.IncreaseAccessSpeed = &v
}

// GetGuestVlanId returns the GuestVlanId field value if set, zero value otherwise.
func (o *InlineObject120) GetGuestVlanId() int32 {
	if o == nil || isNil(o.GuestVlanId) {
		var ret int32
		return ret
	}
	return *o.GuestVlanId
}

// GetGuestVlanIdOk returns a tuple with the GuestVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetGuestVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.GuestVlanId) {
    return nil, false
	}
	return o.GuestVlanId, true
}

// HasGuestVlanId returns a boolean if a field has been set.
func (o *InlineObject120) HasGuestVlanId() bool {
	if o != nil && !isNil(o.GuestVlanId) {
		return true
	}

	return false
}

// SetGuestVlanId gets a reference to the given int32 and assigns it to the GuestVlanId field.
func (o *InlineObject120) SetGuestVlanId(v int32) {
	o.GuestVlanId = &v
}

// GetDot1x returns the Dot1x field value if set, zero value otherwise.
func (o *InlineObject120) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x {
	if o == nil || isNil(o.Dot1x) {
		var ret NetworksNetworkIdSwitchAccessPoliciesDot1x
		return ret
	}
	return *o.Dot1x
}

// GetDot1xOk returns a tuple with the Dot1x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool) {
	if o == nil || isNil(o.Dot1x) {
    return nil, false
	}
	return o.Dot1x, true
}

// HasDot1x returns a boolean if a field has been set.
func (o *InlineObject120) HasDot1x() bool {
	if o != nil && !isNil(o.Dot1x) {
		return true
	}

	return false
}

// SetDot1x gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesDot1x and assigns it to the Dot1x field.
func (o *InlineObject120) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x) {
	o.Dot1x = &v
}

// GetVoiceVlanClients returns the VoiceVlanClients field value if set, zero value otherwise.
func (o *InlineObject120) GetVoiceVlanClients() bool {
	if o == nil || isNil(o.VoiceVlanClients) {
		var ret bool
		return ret
	}
	return *o.VoiceVlanClients
}

// GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetVoiceVlanClientsOk() (*bool, bool) {
	if o == nil || isNil(o.VoiceVlanClients) {
    return nil, false
	}
	return o.VoiceVlanClients, true
}

// HasVoiceVlanClients returns a boolean if a field has been set.
func (o *InlineObject120) HasVoiceVlanClients() bool {
	if o != nil && !isNil(o.VoiceVlanClients) {
		return true
	}

	return false
}

// SetVoiceVlanClients gets a reference to the given bool and assigns it to the VoiceVlanClients field.
func (o *InlineObject120) SetVoiceVlanClients(v bool) {
	o.VoiceVlanClients = &v
}

// GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field value
func (o *InlineObject120) GetUrlRedirectWalledGardenEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UrlRedirectWalledGardenEnabled
}

// GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UrlRedirectWalledGardenEnabled, true
}

// SetUrlRedirectWalledGardenEnabled sets field value
func (o *InlineObject120) SetUrlRedirectWalledGardenEnabled(v bool) {
	o.UrlRedirectWalledGardenEnabled = v
}

// GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field value if set, zero value otherwise.
func (o *InlineObject120) GetUrlRedirectWalledGardenRanges() []string {
	if o == nil || isNil(o.UrlRedirectWalledGardenRanges) {
		var ret []string
		return ret
	}
	return o.UrlRedirectWalledGardenRanges
}

// GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetUrlRedirectWalledGardenRangesOk() ([]string, bool) {
	if o == nil || isNil(o.UrlRedirectWalledGardenRanges) {
    return nil, false
	}
	return o.UrlRedirectWalledGardenRanges, true
}

// HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.
func (o *InlineObject120) HasUrlRedirectWalledGardenRanges() bool {
	if o != nil && !isNil(o.UrlRedirectWalledGardenRanges) {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenRanges gets a reference to the given []string and assigns it to the UrlRedirectWalledGardenRanges field.
func (o *InlineObject120) SetUrlRedirectWalledGardenRanges(v []string) {
	o.UrlRedirectWalledGardenRanges = v
}

func (o InlineObject120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !isNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !isNil(o.GuestPortBouncing) {
		toSerialize["guestPortBouncing"] = o.GuestPortBouncing
	}
	if true {
		toSerialize["radiusTestingEnabled"] = o.RadiusTestingEnabled
	}
	if true {
		toSerialize["radiusCoaSupportEnabled"] = o.RadiusCoaSupportEnabled
	}
	if true {
		toSerialize["radiusAccountingEnabled"] = o.RadiusAccountingEnabled
	}
	if !isNil(o.RadiusAccountingServers) {
		toSerialize["radiusAccountingServers"] = o.RadiusAccountingServers
	}
	if !isNil(o.RadiusGroupAttribute) {
		toSerialize["radiusGroupAttribute"] = o.RadiusGroupAttribute
	}
	if true {
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
	if true {
		toSerialize["urlRedirectWalledGardenEnabled"] = o.UrlRedirectWalledGardenEnabled
	}
	if !isNil(o.UrlRedirectWalledGardenRanges) {
		toSerialize["urlRedirectWalledGardenRanges"] = o.UrlRedirectWalledGardenRanges
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject120 struct {
	value *InlineObject120
	isSet bool
}

func (v NullableInlineObject120) Get() *InlineObject120 {
	return v.value
}

func (v *NullableInlineObject120) Set(val *InlineObject120) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject120) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject120) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject120(val *InlineObject120) *NullableInlineObject120 {
	return &NullableInlineObject120{value: val, isSet: true}
}

func (v NullableInlineObject120) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject120) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


