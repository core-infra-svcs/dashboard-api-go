/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20098 struct for InlineResponse20098
type InlineResponse20098 struct {
	// Unique identifier of the SSID
	Number *int32 `json:"number,omitempty"`
	// The name of the SSID
	Name *string `json:"name,omitempty"`
	// Whether or not the SSID is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The type of splash page for the SSID
	SplashPage *string `json:"splashPage,omitempty"`
	// SSID Administrator access status
	SsidAdminAccessible *bool `json:"ssidAdminAccessible,omitempty"`
	// The association control method for the SSID
	AuthMode *string `json:"authMode,omitempty"`
	// The psk encryption mode for the SSID
	EncryptionMode *string `json:"encryptionMode,omitempty"`
	// The types of WPA encryption
	WpaEncryptionMode *string `json:"wpaEncryptionMode,omitempty"`
	// List of RADIUS 802.1X servers to be used for authentication
	RadiusServers []NetworksNetworkIdWirelessSsidsRadiusServers `json:"radiusServers,omitempty"`
	// List of RADIUS accounting 802.1X servers to be used for authentication
	RadiusAccountingServers []NetworksNetworkIdWirelessSsidsRadiusAccountingServers `json:"radiusAccountingServers,omitempty"`
	// Whether or not RADIUS accounting is enabled
	RadiusAccountingEnabled *bool `json:"radiusAccountingEnabled,omitempty"`
	// Whether RADIUS authentication is enabled
	RadiusEnabled *bool `json:"radiusEnabled,omitempty"`
	// RADIUS attribute used to look up group policies
	RadiusAttributeForGroupPolicies *string `json:"radiusAttributeForGroupPolicies,omitempty"`
	// Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable
	RadiusFailoverPolicy *string `json:"radiusFailoverPolicy,omitempty"`
	// Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts
	RadiusLoadBalancingPolicy *string `json:"radiusLoadBalancingPolicy,omitempty"`
	// The client IP assignment mode
	IpAssignmentMode *string `json:"ipAssignmentMode,omitempty"`
	// URL for the admin splash page
	AdminSplashUrl *string `json:"adminSplashUrl,omitempty"`
	// Splash page timeout
	SplashTimeout *string `json:"splashTimeout,omitempty"`
	// Allow users to access a configurable list of IP ranges prior to sign-on
	WalledGardenEnabled *bool `json:"walledGardenEnabled,omitempty"`
	// Domain names and IP address ranges available in Walled Garden mode
	WalledGardenRanges []string `json:"walledGardenRanges,omitempty"`
	// The minimum bitrate in Mbps of this SSID in the default indoor RF profile
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// The client-serving radio frequencies of this SSID in the default indoor RF profile
	BandSelection *string `json:"bandSelection,omitempty"`
	// The upload bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitUp *int32 `json:"perClientBandwidthLimitUp,omitempty"`
	// The download bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitDown *int32 `json:"perClientBandwidthLimitDown,omitempty"`
	// Whether the SSID is advertised or hidden by the AP
	Visible *bool `json:"visible,omitempty"`
	// Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags
	AvailableOnAllAps *bool `json:"availableOnAllAps,omitempty"`
	// List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list
	AvailabilityTags []string `json:"availabilityTags,omitempty"`
	// The total upload bandwidth limit in Kbps (0 represents no limit)
	PerSsidBandwidthLimitUp *int32 `json:"perSsidBandwidthLimitUp,omitempty"`
	// The total download bandwidth limit in Kbps (0 represents no limit)
	PerSsidBandwidthLimitDown *int32 `json:"perSsidBandwidthLimitDown,omitempty"`
	// Whether clients connecting to this SSID must use the IP address assigned by the DHCP server
	MandatoryDhcpEnabled *bool `json:"mandatoryDhcpEnabled,omitempty"`
}

// NewInlineResponse20098 instantiates a new InlineResponse20098 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20098() *InlineResponse20098 {
	this := InlineResponse20098{}
	return &this
}

// NewInlineResponse20098WithDefaults instantiates a new InlineResponse20098 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20098WithDefaults() *InlineResponse20098 {
	this := InlineResponse20098{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InlineResponse20098) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineResponse20098) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *InlineResponse20098) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20098) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20098) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20098) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20098) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20098) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20098) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSplashPage returns the SplashPage field value if set, zero value otherwise.
func (o *InlineResponse20098) GetSplashPage() string {
	if o == nil || isNil(o.SplashPage) {
		var ret string
		return ret
	}
	return *o.SplashPage
}

// GetSplashPageOk returns a tuple with the SplashPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetSplashPageOk() (*string, bool) {
	if o == nil || isNil(o.SplashPage) {
    return nil, false
	}
	return o.SplashPage, true
}

// HasSplashPage returns a boolean if a field has been set.
func (o *InlineResponse20098) HasSplashPage() bool {
	if o != nil && !isNil(o.SplashPage) {
		return true
	}

	return false
}

// SetSplashPage gets a reference to the given string and assigns it to the SplashPage field.
func (o *InlineResponse20098) SetSplashPage(v string) {
	o.SplashPage = &v
}

// GetSsidAdminAccessible returns the SsidAdminAccessible field value if set, zero value otherwise.
func (o *InlineResponse20098) GetSsidAdminAccessible() bool {
	if o == nil || isNil(o.SsidAdminAccessible) {
		var ret bool
		return ret
	}
	return *o.SsidAdminAccessible
}

// GetSsidAdminAccessibleOk returns a tuple with the SsidAdminAccessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetSsidAdminAccessibleOk() (*bool, bool) {
	if o == nil || isNil(o.SsidAdminAccessible) {
    return nil, false
	}
	return o.SsidAdminAccessible, true
}

// HasSsidAdminAccessible returns a boolean if a field has been set.
func (o *InlineResponse20098) HasSsidAdminAccessible() bool {
	if o != nil && !isNil(o.SsidAdminAccessible) {
		return true
	}

	return false
}

// SetSsidAdminAccessible gets a reference to the given bool and assigns it to the SsidAdminAccessible field.
func (o *InlineResponse20098) SetSsidAdminAccessible(v bool) {
	o.SsidAdminAccessible = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *InlineResponse20098) GetAuthMode() string {
	if o == nil || isNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetAuthModeOk() (*string, bool) {
	if o == nil || isNil(o.AuthMode) {
    return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *InlineResponse20098) HasAuthMode() bool {
	if o != nil && !isNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *InlineResponse20098) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetEncryptionMode returns the EncryptionMode field value if set, zero value otherwise.
func (o *InlineResponse20098) GetEncryptionMode() string {
	if o == nil || isNil(o.EncryptionMode) {
		var ret string
		return ret
	}
	return *o.EncryptionMode
}

// GetEncryptionModeOk returns a tuple with the EncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionMode) {
    return nil, false
	}
	return o.EncryptionMode, true
}

// HasEncryptionMode returns a boolean if a field has been set.
func (o *InlineResponse20098) HasEncryptionMode() bool {
	if o != nil && !isNil(o.EncryptionMode) {
		return true
	}

	return false
}

// SetEncryptionMode gets a reference to the given string and assigns it to the EncryptionMode field.
func (o *InlineResponse20098) SetEncryptionMode(v string) {
	o.EncryptionMode = &v
}

// GetWpaEncryptionMode returns the WpaEncryptionMode field value if set, zero value otherwise.
func (o *InlineResponse20098) GetWpaEncryptionMode() string {
	if o == nil || isNil(o.WpaEncryptionMode) {
		var ret string
		return ret
	}
	return *o.WpaEncryptionMode
}

// GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetWpaEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.WpaEncryptionMode) {
    return nil, false
	}
	return o.WpaEncryptionMode, true
}

// HasWpaEncryptionMode returns a boolean if a field has been set.
func (o *InlineResponse20098) HasWpaEncryptionMode() bool {
	if o != nil && !isNil(o.WpaEncryptionMode) {
		return true
	}

	return false
}

// SetWpaEncryptionMode gets a reference to the given string and assigns it to the WpaEncryptionMode field.
func (o *InlineResponse20098) SetWpaEncryptionMode(v string) {
	o.WpaEncryptionMode = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusServers() []NetworksNetworkIdWirelessSsidsRadiusServers {
	if o == nil || isNil(o.RadiusServers) {
		var ret []NetworksNetworkIdWirelessSsidsRadiusServers
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusServersOk() ([]NetworksNetworkIdWirelessSsidsRadiusServers, bool) {
	if o == nil || isNil(o.RadiusServers) {
    return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusServers() bool {
	if o != nil && !isNil(o.RadiusServers) {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []NetworksNetworkIdWirelessSsidsRadiusServers and assigns it to the RadiusServers field.
func (o *InlineResponse20098) SetRadiusServers(v []NetworksNetworkIdWirelessSsidsRadiusServers) {
	o.RadiusServers = v
}

// GetRadiusAccountingServers returns the RadiusAccountingServers field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusAccountingServers() []NetworksNetworkIdWirelessSsidsRadiusAccountingServers {
	if o == nil || isNil(o.RadiusAccountingServers) {
		var ret []NetworksNetworkIdWirelessSsidsRadiusAccountingServers
		return ret
	}
	return o.RadiusAccountingServers
}

// GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusAccountingServersOk() ([]NetworksNetworkIdWirelessSsidsRadiusAccountingServers, bool) {
	if o == nil || isNil(o.RadiusAccountingServers) {
    return nil, false
	}
	return o.RadiusAccountingServers, true
}

// HasRadiusAccountingServers returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusAccountingServers() bool {
	if o != nil && !isNil(o.RadiusAccountingServers) {
		return true
	}

	return false
}

// SetRadiusAccountingServers gets a reference to the given []NetworksNetworkIdWirelessSsidsRadiusAccountingServers and assigns it to the RadiusAccountingServers field.
func (o *InlineResponse20098) SetRadiusAccountingServers(v []NetworksNetworkIdWirelessSsidsRadiusAccountingServers) {
	o.RadiusAccountingServers = v
}

// GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusAccountingEnabled() bool {
	if o == nil || isNil(o.RadiusAccountingEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusAccountingEnabled
}

// GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusAccountingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RadiusAccountingEnabled) {
    return nil, false
	}
	return o.RadiusAccountingEnabled, true
}

// HasRadiusAccountingEnabled returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusAccountingEnabled() bool {
	if o != nil && !isNil(o.RadiusAccountingEnabled) {
		return true
	}

	return false
}

// SetRadiusAccountingEnabled gets a reference to the given bool and assigns it to the RadiusAccountingEnabled field.
func (o *InlineResponse20098) SetRadiusAccountingEnabled(v bool) {
	o.RadiusAccountingEnabled = &v
}

// GetRadiusEnabled returns the RadiusEnabled field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusEnabled() bool {
	if o == nil || isNil(o.RadiusEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusEnabled
}

// GetRadiusEnabledOk returns a tuple with the RadiusEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RadiusEnabled) {
    return nil, false
	}
	return o.RadiusEnabled, true
}

// HasRadiusEnabled returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusEnabled() bool {
	if o != nil && !isNil(o.RadiusEnabled) {
		return true
	}

	return false
}

// SetRadiusEnabled gets a reference to the given bool and assigns it to the RadiusEnabled field.
func (o *InlineResponse20098) SetRadiusEnabled(v bool) {
	o.RadiusEnabled = &v
}

// GetRadiusAttributeForGroupPolicies returns the RadiusAttributeForGroupPolicies field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusAttributeForGroupPolicies() string {
	if o == nil || isNil(o.RadiusAttributeForGroupPolicies) {
		var ret string
		return ret
	}
	return *o.RadiusAttributeForGroupPolicies
}

// GetRadiusAttributeForGroupPoliciesOk returns a tuple with the RadiusAttributeForGroupPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusAttributeForGroupPoliciesOk() (*string, bool) {
	if o == nil || isNil(o.RadiusAttributeForGroupPolicies) {
    return nil, false
	}
	return o.RadiusAttributeForGroupPolicies, true
}

// HasRadiusAttributeForGroupPolicies returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusAttributeForGroupPolicies() bool {
	if o != nil && !isNil(o.RadiusAttributeForGroupPolicies) {
		return true
	}

	return false
}

// SetRadiusAttributeForGroupPolicies gets a reference to the given string and assigns it to the RadiusAttributeForGroupPolicies field.
func (o *InlineResponse20098) SetRadiusAttributeForGroupPolicies(v string) {
	o.RadiusAttributeForGroupPolicies = &v
}

// GetRadiusFailoverPolicy returns the RadiusFailoverPolicy field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusFailoverPolicy() string {
	if o == nil || isNil(o.RadiusFailoverPolicy) {
		var ret string
		return ret
	}
	return *o.RadiusFailoverPolicy
}

// GetRadiusFailoverPolicyOk returns a tuple with the RadiusFailoverPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusFailoverPolicyOk() (*string, bool) {
	if o == nil || isNil(o.RadiusFailoverPolicy) {
    return nil, false
	}
	return o.RadiusFailoverPolicy, true
}

// HasRadiusFailoverPolicy returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusFailoverPolicy() bool {
	if o != nil && !isNil(o.RadiusFailoverPolicy) {
		return true
	}

	return false
}

// SetRadiusFailoverPolicy gets a reference to the given string and assigns it to the RadiusFailoverPolicy field.
func (o *InlineResponse20098) SetRadiusFailoverPolicy(v string) {
	o.RadiusFailoverPolicy = &v
}

// GetRadiusLoadBalancingPolicy returns the RadiusLoadBalancingPolicy field value if set, zero value otherwise.
func (o *InlineResponse20098) GetRadiusLoadBalancingPolicy() string {
	if o == nil || isNil(o.RadiusLoadBalancingPolicy) {
		var ret string
		return ret
	}
	return *o.RadiusLoadBalancingPolicy
}

// GetRadiusLoadBalancingPolicyOk returns a tuple with the RadiusLoadBalancingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetRadiusLoadBalancingPolicyOk() (*string, bool) {
	if o == nil || isNil(o.RadiusLoadBalancingPolicy) {
    return nil, false
	}
	return o.RadiusLoadBalancingPolicy, true
}

// HasRadiusLoadBalancingPolicy returns a boolean if a field has been set.
func (o *InlineResponse20098) HasRadiusLoadBalancingPolicy() bool {
	if o != nil && !isNil(o.RadiusLoadBalancingPolicy) {
		return true
	}

	return false
}

// SetRadiusLoadBalancingPolicy gets a reference to the given string and assigns it to the RadiusLoadBalancingPolicy field.
func (o *InlineResponse20098) SetRadiusLoadBalancingPolicy(v string) {
	o.RadiusLoadBalancingPolicy = &v
}

// GetIpAssignmentMode returns the IpAssignmentMode field value if set, zero value otherwise.
func (o *InlineResponse20098) GetIpAssignmentMode() string {
	if o == nil || isNil(o.IpAssignmentMode) {
		var ret string
		return ret
	}
	return *o.IpAssignmentMode
}

// GetIpAssignmentModeOk returns a tuple with the IpAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetIpAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.IpAssignmentMode) {
    return nil, false
	}
	return o.IpAssignmentMode, true
}

// HasIpAssignmentMode returns a boolean if a field has been set.
func (o *InlineResponse20098) HasIpAssignmentMode() bool {
	if o != nil && !isNil(o.IpAssignmentMode) {
		return true
	}

	return false
}

// SetIpAssignmentMode gets a reference to the given string and assigns it to the IpAssignmentMode field.
func (o *InlineResponse20098) SetIpAssignmentMode(v string) {
	o.IpAssignmentMode = &v
}

// GetAdminSplashUrl returns the AdminSplashUrl field value if set, zero value otherwise.
func (o *InlineResponse20098) GetAdminSplashUrl() string {
	if o == nil || isNil(o.AdminSplashUrl) {
		var ret string
		return ret
	}
	return *o.AdminSplashUrl
}

// GetAdminSplashUrlOk returns a tuple with the AdminSplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetAdminSplashUrlOk() (*string, bool) {
	if o == nil || isNil(o.AdminSplashUrl) {
    return nil, false
	}
	return o.AdminSplashUrl, true
}

// HasAdminSplashUrl returns a boolean if a field has been set.
func (o *InlineResponse20098) HasAdminSplashUrl() bool {
	if o != nil && !isNil(o.AdminSplashUrl) {
		return true
	}

	return false
}

// SetAdminSplashUrl gets a reference to the given string and assigns it to the AdminSplashUrl field.
func (o *InlineResponse20098) SetAdminSplashUrl(v string) {
	o.AdminSplashUrl = &v
}

// GetSplashTimeout returns the SplashTimeout field value if set, zero value otherwise.
func (o *InlineResponse20098) GetSplashTimeout() string {
	if o == nil || isNil(o.SplashTimeout) {
		var ret string
		return ret
	}
	return *o.SplashTimeout
}

// GetSplashTimeoutOk returns a tuple with the SplashTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetSplashTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.SplashTimeout) {
    return nil, false
	}
	return o.SplashTimeout, true
}

// HasSplashTimeout returns a boolean if a field has been set.
func (o *InlineResponse20098) HasSplashTimeout() bool {
	if o != nil && !isNil(o.SplashTimeout) {
		return true
	}

	return false
}

// SetSplashTimeout gets a reference to the given string and assigns it to the SplashTimeout field.
func (o *InlineResponse20098) SetSplashTimeout(v string) {
	o.SplashTimeout = &v
}

// GetWalledGardenEnabled returns the WalledGardenEnabled field value if set, zero value otherwise.
func (o *InlineResponse20098) GetWalledGardenEnabled() bool {
	if o == nil || isNil(o.WalledGardenEnabled) {
		var ret bool
		return ret
	}
	return *o.WalledGardenEnabled
}

// GetWalledGardenEnabledOk returns a tuple with the WalledGardenEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetWalledGardenEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.WalledGardenEnabled) {
    return nil, false
	}
	return o.WalledGardenEnabled, true
}

// HasWalledGardenEnabled returns a boolean if a field has been set.
func (o *InlineResponse20098) HasWalledGardenEnabled() bool {
	if o != nil && !isNil(o.WalledGardenEnabled) {
		return true
	}

	return false
}

// SetWalledGardenEnabled gets a reference to the given bool and assigns it to the WalledGardenEnabled field.
func (o *InlineResponse20098) SetWalledGardenEnabled(v bool) {
	o.WalledGardenEnabled = &v
}

// GetWalledGardenRanges returns the WalledGardenRanges field value if set, zero value otherwise.
func (o *InlineResponse20098) GetWalledGardenRanges() []string {
	if o == nil || isNil(o.WalledGardenRanges) {
		var ret []string
		return ret
	}
	return o.WalledGardenRanges
}

// GetWalledGardenRangesOk returns a tuple with the WalledGardenRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetWalledGardenRangesOk() ([]string, bool) {
	if o == nil || isNil(o.WalledGardenRanges) {
    return nil, false
	}
	return o.WalledGardenRanges, true
}

// HasWalledGardenRanges returns a boolean if a field has been set.
func (o *InlineResponse20098) HasWalledGardenRanges() bool {
	if o != nil && !isNil(o.WalledGardenRanges) {
		return true
	}

	return false
}

// SetWalledGardenRanges gets a reference to the given []string and assigns it to the WalledGardenRanges field.
func (o *InlineResponse20098) SetWalledGardenRanges(v []string) {
	o.WalledGardenRanges = v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *InlineResponse20098) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *InlineResponse20098) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *InlineResponse20098) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetBandSelection returns the BandSelection field value if set, zero value otherwise.
func (o *InlineResponse20098) GetBandSelection() string {
	if o == nil || isNil(o.BandSelection) {
		var ret string
		return ret
	}
	return *o.BandSelection
}

// GetBandSelectionOk returns a tuple with the BandSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetBandSelectionOk() (*string, bool) {
	if o == nil || isNil(o.BandSelection) {
    return nil, false
	}
	return o.BandSelection, true
}

// HasBandSelection returns a boolean if a field has been set.
func (o *InlineResponse20098) HasBandSelection() bool {
	if o != nil && !isNil(o.BandSelection) {
		return true
	}

	return false
}

// SetBandSelection gets a reference to the given string and assigns it to the BandSelection field.
func (o *InlineResponse20098) SetBandSelection(v string) {
	o.BandSelection = &v
}

// GetPerClientBandwidthLimitUp returns the PerClientBandwidthLimitUp field value if set, zero value otherwise.
func (o *InlineResponse20098) GetPerClientBandwidthLimitUp() int32 {
	if o == nil || isNil(o.PerClientBandwidthLimitUp) {
		var ret int32
		return ret
	}
	return *o.PerClientBandwidthLimitUp
}

// GetPerClientBandwidthLimitUpOk returns a tuple with the PerClientBandwidthLimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetPerClientBandwidthLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.PerClientBandwidthLimitUp) {
    return nil, false
	}
	return o.PerClientBandwidthLimitUp, true
}

// HasPerClientBandwidthLimitUp returns a boolean if a field has been set.
func (o *InlineResponse20098) HasPerClientBandwidthLimitUp() bool {
	if o != nil && !isNil(o.PerClientBandwidthLimitUp) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimitUp gets a reference to the given int32 and assigns it to the PerClientBandwidthLimitUp field.
func (o *InlineResponse20098) SetPerClientBandwidthLimitUp(v int32) {
	o.PerClientBandwidthLimitUp = &v
}

// GetPerClientBandwidthLimitDown returns the PerClientBandwidthLimitDown field value if set, zero value otherwise.
func (o *InlineResponse20098) GetPerClientBandwidthLimitDown() int32 {
	if o == nil || isNil(o.PerClientBandwidthLimitDown) {
		var ret int32
		return ret
	}
	return *o.PerClientBandwidthLimitDown
}

// GetPerClientBandwidthLimitDownOk returns a tuple with the PerClientBandwidthLimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetPerClientBandwidthLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.PerClientBandwidthLimitDown) {
    return nil, false
	}
	return o.PerClientBandwidthLimitDown, true
}

// HasPerClientBandwidthLimitDown returns a boolean if a field has been set.
func (o *InlineResponse20098) HasPerClientBandwidthLimitDown() bool {
	if o != nil && !isNil(o.PerClientBandwidthLimitDown) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimitDown gets a reference to the given int32 and assigns it to the PerClientBandwidthLimitDown field.
func (o *InlineResponse20098) SetPerClientBandwidthLimitDown(v int32) {
	o.PerClientBandwidthLimitDown = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *InlineResponse20098) GetVisible() bool {
	if o == nil || isNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetVisibleOk() (*bool, bool) {
	if o == nil || isNil(o.Visible) {
    return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *InlineResponse20098) HasVisible() bool {
	if o != nil && !isNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *InlineResponse20098) SetVisible(v bool) {
	o.Visible = &v
}

// GetAvailableOnAllAps returns the AvailableOnAllAps field value if set, zero value otherwise.
func (o *InlineResponse20098) GetAvailableOnAllAps() bool {
	if o == nil || isNil(o.AvailableOnAllAps) {
		var ret bool
		return ret
	}
	return *o.AvailableOnAllAps
}

// GetAvailableOnAllApsOk returns a tuple with the AvailableOnAllAps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetAvailableOnAllApsOk() (*bool, bool) {
	if o == nil || isNil(o.AvailableOnAllAps) {
    return nil, false
	}
	return o.AvailableOnAllAps, true
}

// HasAvailableOnAllAps returns a boolean if a field has been set.
func (o *InlineResponse20098) HasAvailableOnAllAps() bool {
	if o != nil && !isNil(o.AvailableOnAllAps) {
		return true
	}

	return false
}

// SetAvailableOnAllAps gets a reference to the given bool and assigns it to the AvailableOnAllAps field.
func (o *InlineResponse20098) SetAvailableOnAllAps(v bool) {
	o.AvailableOnAllAps = &v
}

// GetAvailabilityTags returns the AvailabilityTags field value if set, zero value otherwise.
func (o *InlineResponse20098) GetAvailabilityTags() []string {
	if o == nil || isNil(o.AvailabilityTags) {
		var ret []string
		return ret
	}
	return o.AvailabilityTags
}

// GetAvailabilityTagsOk returns a tuple with the AvailabilityTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetAvailabilityTagsOk() ([]string, bool) {
	if o == nil || isNil(o.AvailabilityTags) {
    return nil, false
	}
	return o.AvailabilityTags, true
}

// HasAvailabilityTags returns a boolean if a field has been set.
func (o *InlineResponse20098) HasAvailabilityTags() bool {
	if o != nil && !isNil(o.AvailabilityTags) {
		return true
	}

	return false
}

// SetAvailabilityTags gets a reference to the given []string and assigns it to the AvailabilityTags field.
func (o *InlineResponse20098) SetAvailabilityTags(v []string) {
	o.AvailabilityTags = v
}

// GetPerSsidBandwidthLimitUp returns the PerSsidBandwidthLimitUp field value if set, zero value otherwise.
func (o *InlineResponse20098) GetPerSsidBandwidthLimitUp() int32 {
	if o == nil || isNil(o.PerSsidBandwidthLimitUp) {
		var ret int32
		return ret
	}
	return *o.PerSsidBandwidthLimitUp
}

// GetPerSsidBandwidthLimitUpOk returns a tuple with the PerSsidBandwidthLimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetPerSsidBandwidthLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.PerSsidBandwidthLimitUp) {
    return nil, false
	}
	return o.PerSsidBandwidthLimitUp, true
}

// HasPerSsidBandwidthLimitUp returns a boolean if a field has been set.
func (o *InlineResponse20098) HasPerSsidBandwidthLimitUp() bool {
	if o != nil && !isNil(o.PerSsidBandwidthLimitUp) {
		return true
	}

	return false
}

// SetPerSsidBandwidthLimitUp gets a reference to the given int32 and assigns it to the PerSsidBandwidthLimitUp field.
func (o *InlineResponse20098) SetPerSsidBandwidthLimitUp(v int32) {
	o.PerSsidBandwidthLimitUp = &v
}

// GetPerSsidBandwidthLimitDown returns the PerSsidBandwidthLimitDown field value if set, zero value otherwise.
func (o *InlineResponse20098) GetPerSsidBandwidthLimitDown() int32 {
	if o == nil || isNil(o.PerSsidBandwidthLimitDown) {
		var ret int32
		return ret
	}
	return *o.PerSsidBandwidthLimitDown
}

// GetPerSsidBandwidthLimitDownOk returns a tuple with the PerSsidBandwidthLimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetPerSsidBandwidthLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.PerSsidBandwidthLimitDown) {
    return nil, false
	}
	return o.PerSsidBandwidthLimitDown, true
}

// HasPerSsidBandwidthLimitDown returns a boolean if a field has been set.
func (o *InlineResponse20098) HasPerSsidBandwidthLimitDown() bool {
	if o != nil && !isNil(o.PerSsidBandwidthLimitDown) {
		return true
	}

	return false
}

// SetPerSsidBandwidthLimitDown gets a reference to the given int32 and assigns it to the PerSsidBandwidthLimitDown field.
func (o *InlineResponse20098) SetPerSsidBandwidthLimitDown(v int32) {
	o.PerSsidBandwidthLimitDown = &v
}

// GetMandatoryDhcpEnabled returns the MandatoryDhcpEnabled field value if set, zero value otherwise.
func (o *InlineResponse20098) GetMandatoryDhcpEnabled() bool {
	if o == nil || isNil(o.MandatoryDhcpEnabled) {
		var ret bool
		return ret
	}
	return *o.MandatoryDhcpEnabled
}

// GetMandatoryDhcpEnabledOk returns a tuple with the MandatoryDhcpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetMandatoryDhcpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MandatoryDhcpEnabled) {
    return nil, false
	}
	return o.MandatoryDhcpEnabled, true
}

// HasMandatoryDhcpEnabled returns a boolean if a field has been set.
func (o *InlineResponse20098) HasMandatoryDhcpEnabled() bool {
	if o != nil && !isNil(o.MandatoryDhcpEnabled) {
		return true
	}

	return false
}

// SetMandatoryDhcpEnabled gets a reference to the given bool and assigns it to the MandatoryDhcpEnabled field.
func (o *InlineResponse20098) SetMandatoryDhcpEnabled(v bool) {
	o.MandatoryDhcpEnabled = &v
}

func (o InlineResponse20098) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SplashPage) {
		toSerialize["splashPage"] = o.SplashPage
	}
	if !isNil(o.SsidAdminAccessible) {
		toSerialize["ssidAdminAccessible"] = o.SsidAdminAccessible
	}
	if !isNil(o.AuthMode) {
		toSerialize["authMode"] = o.AuthMode
	}
	if !isNil(o.EncryptionMode) {
		toSerialize["encryptionMode"] = o.EncryptionMode
	}
	if !isNil(o.WpaEncryptionMode) {
		toSerialize["wpaEncryptionMode"] = o.WpaEncryptionMode
	}
	if !isNil(o.RadiusServers) {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !isNil(o.RadiusAccountingServers) {
		toSerialize["radiusAccountingServers"] = o.RadiusAccountingServers
	}
	if !isNil(o.RadiusAccountingEnabled) {
		toSerialize["radiusAccountingEnabled"] = o.RadiusAccountingEnabled
	}
	if !isNil(o.RadiusEnabled) {
		toSerialize["radiusEnabled"] = o.RadiusEnabled
	}
	if !isNil(o.RadiusAttributeForGroupPolicies) {
		toSerialize["radiusAttributeForGroupPolicies"] = o.RadiusAttributeForGroupPolicies
	}
	if !isNil(o.RadiusFailoverPolicy) {
		toSerialize["radiusFailoverPolicy"] = o.RadiusFailoverPolicy
	}
	if !isNil(o.RadiusLoadBalancingPolicy) {
		toSerialize["radiusLoadBalancingPolicy"] = o.RadiusLoadBalancingPolicy
	}
	if !isNil(o.IpAssignmentMode) {
		toSerialize["ipAssignmentMode"] = o.IpAssignmentMode
	}
	if !isNil(o.AdminSplashUrl) {
		toSerialize["adminSplashUrl"] = o.AdminSplashUrl
	}
	if !isNil(o.SplashTimeout) {
		toSerialize["splashTimeout"] = o.SplashTimeout
	}
	if !isNil(o.WalledGardenEnabled) {
		toSerialize["walledGardenEnabled"] = o.WalledGardenEnabled
	}
	if !isNil(o.WalledGardenRanges) {
		toSerialize["walledGardenRanges"] = o.WalledGardenRanges
	}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.BandSelection) {
		toSerialize["bandSelection"] = o.BandSelection
	}
	if !isNil(o.PerClientBandwidthLimitUp) {
		toSerialize["perClientBandwidthLimitUp"] = o.PerClientBandwidthLimitUp
	}
	if !isNil(o.PerClientBandwidthLimitDown) {
		toSerialize["perClientBandwidthLimitDown"] = o.PerClientBandwidthLimitDown
	}
	if !isNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !isNil(o.AvailableOnAllAps) {
		toSerialize["availableOnAllAps"] = o.AvailableOnAllAps
	}
	if !isNil(o.AvailabilityTags) {
		toSerialize["availabilityTags"] = o.AvailabilityTags
	}
	if !isNil(o.PerSsidBandwidthLimitUp) {
		toSerialize["perSsidBandwidthLimitUp"] = o.PerSsidBandwidthLimitUp
	}
	if !isNil(o.PerSsidBandwidthLimitDown) {
		toSerialize["perSsidBandwidthLimitDown"] = o.PerSsidBandwidthLimitDown
	}
	if !isNil(o.MandatoryDhcpEnabled) {
		toSerialize["mandatoryDhcpEnabled"] = o.MandatoryDhcpEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20098 struct {
	value *InlineResponse20098
	isSet bool
}

func (v NullableInlineResponse20098) Get() *InlineResponse20098 {
	return v.value
}

func (v *NullableInlineResponse20098) Set(val *InlineResponse20098) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20098) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20098) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20098(val *InlineResponse20098) *NullableInlineResponse20098 {
	return &NullableInlineResponse20098{value: val, isSet: true}
}

func (v NullableInlineResponse20098) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20098) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


