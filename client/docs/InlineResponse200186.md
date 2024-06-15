# InlineResponse200186

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Unique identifier of the SSID | [optional] 
**Name** | Pointer to **string** | The name of the SSID | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled | [optional] 
**SplashPage** | Pointer to **string** | The type of splash page for the SSID | [optional] 
**SsidAdminAccessible** | Pointer to **bool** | SSID Administrator access status | [optional] 
**LocalAuth** | Pointer to **bool** | Extended local auth flag for Enterprise NAC | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID | [optional] 
**WpaEncryptionMode** | Pointer to **string** | The types of WPA encryption | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsRadiusServers**](NetworksNetworkIdWirelessSsidsRadiusServers.md) | List of RADIUS 802.1X servers to be used for authentication | [optional] 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsRadiusAccountingServers**](NetworksNetworkIdWirelessSsidsRadiusAccountingServers.md) | List of RADIUS accounting 802.1X servers to be used for authentication | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Whether or not RADIUS accounting is enabled | [optional] 
**RadiusEnabled** | Pointer to **bool** | Whether RADIUS authentication is enabled | [optional] 
**RadiusAttributeForGroupPolicies** | Pointer to **string** | RADIUS attribute used to look up group policies | [optional] 
**RadiusFailoverPolicy** | Pointer to **string** | Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable | [optional] 
**RadiusLoadBalancingPolicy** | Pointer to **string** | Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts | [optional] 
**IpAssignmentMode** | Pointer to **string** | The client IP assignment mode | [optional] 
**AdminSplashUrl** | Pointer to **string** | URL for the admin splash page | [optional] 
**SplashTimeout** | Pointer to **string** | Splash page timeout | [optional] 
**WalledGardenEnabled** | Pointer to **bool** | Allow users to access a configurable list of IP ranges prior to sign-on | [optional] 
**WalledGardenRanges** | Pointer to **[]string** | Domain names and IP address ranges available in Walled Garden mode | [optional] 
**MinBitrate** | Pointer to **int32** | The minimum bitrate in Mbps of this SSID in the default indoor RF profile | [optional] 
**BandSelection** | Pointer to **string** | The client-serving radio frequencies of this SSID in the default indoor RF profile | [optional] 
**PerClientBandwidthLimitUp** | Pointer to **int32** | The upload bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**PerClientBandwidthLimitDown** | Pointer to **int32** | The download bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**Visible** | Pointer to **bool** | Whether the SSID is advertised or hidden by the AP | [optional] 
**AvailableOnAllAps** | Pointer to **bool** | Whether all APs broadcast the SSID or if it&#39;s restricted to APs matching any availability tags | [optional] 
**AvailabilityTags** | Pointer to **[]string** | List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list | [optional] 
**PerSsidBandwidthLimitUp** | Pointer to **int32** | The total upload bandwidth limit in Kbps (0 represents no limit) | [optional] 
**PerSsidBandwidthLimitDown** | Pointer to **int32** | The total download bandwidth limit in Kbps (0 represents no limit) | [optional] 
**MandatoryDhcpEnabled** | Pointer to **bool** | Whether clients connecting to this SSID must use the IP address assigned by the DHCP server | [optional] 

## Methods

### NewInlineResponse200186

`func NewInlineResponse200186() *InlineResponse200186`

NewInlineResponse200186 instantiates a new InlineResponse200186 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200186WithDefaults

`func NewInlineResponse200186WithDefaults() *InlineResponse200186`

NewInlineResponse200186WithDefaults instantiates a new InlineResponse200186 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse200186) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200186) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200186) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200186) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200186) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200186) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200186) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200186) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200186) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200186) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200186) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200186) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineResponse200186) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineResponse200186) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineResponse200186) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *InlineResponse200186) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetSsidAdminAccessible

`func (o *InlineResponse200186) GetSsidAdminAccessible() bool`

GetSsidAdminAccessible returns the SsidAdminAccessible field if non-nil, zero value otherwise.

### GetSsidAdminAccessibleOk

`func (o *InlineResponse200186) GetSsidAdminAccessibleOk() (*bool, bool)`

GetSsidAdminAccessibleOk returns a tuple with the SsidAdminAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidAdminAccessible

`func (o *InlineResponse200186) SetSsidAdminAccessible(v bool)`

SetSsidAdminAccessible sets SsidAdminAccessible field to given value.

### HasSsidAdminAccessible

`func (o *InlineResponse200186) HasSsidAdminAccessible() bool`

HasSsidAdminAccessible returns a boolean if a field has been set.

### GetLocalAuth

`func (o *InlineResponse200186) GetLocalAuth() bool`

GetLocalAuth returns the LocalAuth field if non-nil, zero value otherwise.

### GetLocalAuthOk

`func (o *InlineResponse200186) GetLocalAuthOk() (*bool, bool)`

GetLocalAuthOk returns a tuple with the LocalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuth

`func (o *InlineResponse200186) SetLocalAuth(v bool)`

SetLocalAuth sets LocalAuth field to given value.

### HasLocalAuth

`func (o *InlineResponse200186) HasLocalAuth() bool`

HasLocalAuth returns a boolean if a field has been set.

### GetAuthMode

`func (o *InlineResponse200186) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InlineResponse200186) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InlineResponse200186) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *InlineResponse200186) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *InlineResponse200186) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *InlineResponse200186) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *InlineResponse200186) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *InlineResponse200186) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *InlineResponse200186) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *InlineResponse200186) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *InlineResponse200186) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *InlineResponse200186) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineResponse200186) GetRadiusServers() []NetworksNetworkIdWirelessSsidsRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineResponse200186) GetRadiusServersOk() (*[]NetworksNetworkIdWirelessSsidsRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineResponse200186) SetRadiusServers(v []NetworksNetworkIdWirelessSsidsRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineResponse200186) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *InlineResponse200186) GetRadiusAccountingServers() []NetworksNetworkIdWirelessSsidsRadiusAccountingServers`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineResponse200186) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdWirelessSsidsRadiusAccountingServers, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineResponse200186) SetRadiusAccountingServers(v []NetworksNetworkIdWirelessSsidsRadiusAccountingServers)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineResponse200186) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *InlineResponse200186) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineResponse200186) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineResponse200186) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *InlineResponse200186) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusEnabled

`func (o *InlineResponse200186) GetRadiusEnabled() bool`

GetRadiusEnabled returns the RadiusEnabled field if non-nil, zero value otherwise.

### GetRadiusEnabledOk

`func (o *InlineResponse200186) GetRadiusEnabledOk() (*bool, bool)`

GetRadiusEnabledOk returns a tuple with the RadiusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusEnabled

`func (o *InlineResponse200186) SetRadiusEnabled(v bool)`

SetRadiusEnabled sets RadiusEnabled field to given value.

### HasRadiusEnabled

`func (o *InlineResponse200186) HasRadiusEnabled() bool`

HasRadiusEnabled returns a boolean if a field has been set.

### GetRadiusAttributeForGroupPolicies

`func (o *InlineResponse200186) GetRadiusAttributeForGroupPolicies() string`

GetRadiusAttributeForGroupPolicies returns the RadiusAttributeForGroupPolicies field if non-nil, zero value otherwise.

### GetRadiusAttributeForGroupPoliciesOk

`func (o *InlineResponse200186) GetRadiusAttributeForGroupPoliciesOk() (*string, bool)`

GetRadiusAttributeForGroupPoliciesOk returns a tuple with the RadiusAttributeForGroupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttributeForGroupPolicies

`func (o *InlineResponse200186) SetRadiusAttributeForGroupPolicies(v string)`

SetRadiusAttributeForGroupPolicies sets RadiusAttributeForGroupPolicies field to given value.

### HasRadiusAttributeForGroupPolicies

`func (o *InlineResponse200186) HasRadiusAttributeForGroupPolicies() bool`

HasRadiusAttributeForGroupPolicies returns a boolean if a field has been set.

### GetRadiusFailoverPolicy

`func (o *InlineResponse200186) GetRadiusFailoverPolicy() string`

GetRadiusFailoverPolicy returns the RadiusFailoverPolicy field if non-nil, zero value otherwise.

### GetRadiusFailoverPolicyOk

`func (o *InlineResponse200186) GetRadiusFailoverPolicyOk() (*string, bool)`

GetRadiusFailoverPolicyOk returns a tuple with the RadiusFailoverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusFailoverPolicy

`func (o *InlineResponse200186) SetRadiusFailoverPolicy(v string)`

SetRadiusFailoverPolicy sets RadiusFailoverPolicy field to given value.

### HasRadiusFailoverPolicy

`func (o *InlineResponse200186) HasRadiusFailoverPolicy() bool`

HasRadiusFailoverPolicy returns a boolean if a field has been set.

### GetRadiusLoadBalancingPolicy

`func (o *InlineResponse200186) GetRadiusLoadBalancingPolicy() string`

GetRadiusLoadBalancingPolicy returns the RadiusLoadBalancingPolicy field if non-nil, zero value otherwise.

### GetRadiusLoadBalancingPolicyOk

`func (o *InlineResponse200186) GetRadiusLoadBalancingPolicyOk() (*string, bool)`

GetRadiusLoadBalancingPolicyOk returns a tuple with the RadiusLoadBalancingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusLoadBalancingPolicy

`func (o *InlineResponse200186) SetRadiusLoadBalancingPolicy(v string)`

SetRadiusLoadBalancingPolicy sets RadiusLoadBalancingPolicy field to given value.

### HasRadiusLoadBalancingPolicy

`func (o *InlineResponse200186) HasRadiusLoadBalancingPolicy() bool`

HasRadiusLoadBalancingPolicy returns a boolean if a field has been set.

### GetIpAssignmentMode

`func (o *InlineResponse200186) GetIpAssignmentMode() string`

GetIpAssignmentMode returns the IpAssignmentMode field if non-nil, zero value otherwise.

### GetIpAssignmentModeOk

`func (o *InlineResponse200186) GetIpAssignmentModeOk() (*string, bool)`

GetIpAssignmentModeOk returns a tuple with the IpAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignmentMode

`func (o *InlineResponse200186) SetIpAssignmentMode(v string)`

SetIpAssignmentMode sets IpAssignmentMode field to given value.

### HasIpAssignmentMode

`func (o *InlineResponse200186) HasIpAssignmentMode() bool`

HasIpAssignmentMode returns a boolean if a field has been set.

### GetAdminSplashUrl

`func (o *InlineResponse200186) GetAdminSplashUrl() string`

GetAdminSplashUrl returns the AdminSplashUrl field if non-nil, zero value otherwise.

### GetAdminSplashUrlOk

`func (o *InlineResponse200186) GetAdminSplashUrlOk() (*string, bool)`

GetAdminSplashUrlOk returns a tuple with the AdminSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSplashUrl

`func (o *InlineResponse200186) SetAdminSplashUrl(v string)`

SetAdminSplashUrl sets AdminSplashUrl field to given value.

### HasAdminSplashUrl

`func (o *InlineResponse200186) HasAdminSplashUrl() bool`

HasAdminSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineResponse200186) GetSplashTimeout() string`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineResponse200186) GetSplashTimeoutOk() (*string, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineResponse200186) SetSplashTimeout(v string)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineResponse200186) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetWalledGardenEnabled

`func (o *InlineResponse200186) GetWalledGardenEnabled() bool`

GetWalledGardenEnabled returns the WalledGardenEnabled field if non-nil, zero value otherwise.

### GetWalledGardenEnabledOk

`func (o *InlineResponse200186) GetWalledGardenEnabledOk() (*bool, bool)`

GetWalledGardenEnabledOk returns a tuple with the WalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenEnabled

`func (o *InlineResponse200186) SetWalledGardenEnabled(v bool)`

SetWalledGardenEnabled sets WalledGardenEnabled field to given value.

### HasWalledGardenEnabled

`func (o *InlineResponse200186) HasWalledGardenEnabled() bool`

HasWalledGardenEnabled returns a boolean if a field has been set.

### GetWalledGardenRanges

`func (o *InlineResponse200186) GetWalledGardenRanges() []string`

GetWalledGardenRanges returns the WalledGardenRanges field if non-nil, zero value otherwise.

### GetWalledGardenRangesOk

`func (o *InlineResponse200186) GetWalledGardenRangesOk() (*[]string, bool)`

GetWalledGardenRangesOk returns a tuple with the WalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenRanges

`func (o *InlineResponse200186) SetWalledGardenRanges(v []string)`

SetWalledGardenRanges sets WalledGardenRanges field to given value.

### HasWalledGardenRanges

`func (o *InlineResponse200186) HasWalledGardenRanges() bool`

HasWalledGardenRanges returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse200186) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse200186) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse200186) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineResponse200186) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetBandSelection

`func (o *InlineResponse200186) GetBandSelection() string`

GetBandSelection returns the BandSelection field if non-nil, zero value otherwise.

### GetBandSelectionOk

`func (o *InlineResponse200186) GetBandSelectionOk() (*string, bool)`

GetBandSelectionOk returns a tuple with the BandSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelection

`func (o *InlineResponse200186) SetBandSelection(v string)`

SetBandSelection sets BandSelection field to given value.

### HasBandSelection

`func (o *InlineResponse200186) HasBandSelection() bool`

HasBandSelection returns a boolean if a field has been set.

### GetPerClientBandwidthLimitUp

`func (o *InlineResponse200186) GetPerClientBandwidthLimitUp() int32`

GetPerClientBandwidthLimitUp returns the PerClientBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitUpOk

`func (o *InlineResponse200186) GetPerClientBandwidthLimitUpOk() (*int32, bool)`

GetPerClientBandwidthLimitUpOk returns a tuple with the PerClientBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitUp

`func (o *InlineResponse200186) SetPerClientBandwidthLimitUp(v int32)`

SetPerClientBandwidthLimitUp sets PerClientBandwidthLimitUp field to given value.

### HasPerClientBandwidthLimitUp

`func (o *InlineResponse200186) HasPerClientBandwidthLimitUp() bool`

HasPerClientBandwidthLimitUp returns a boolean if a field has been set.

### GetPerClientBandwidthLimitDown

`func (o *InlineResponse200186) GetPerClientBandwidthLimitDown() int32`

GetPerClientBandwidthLimitDown returns the PerClientBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitDownOk

`func (o *InlineResponse200186) GetPerClientBandwidthLimitDownOk() (*int32, bool)`

GetPerClientBandwidthLimitDownOk returns a tuple with the PerClientBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitDown

`func (o *InlineResponse200186) SetPerClientBandwidthLimitDown(v int32)`

SetPerClientBandwidthLimitDown sets PerClientBandwidthLimitDown field to given value.

### HasPerClientBandwidthLimitDown

`func (o *InlineResponse200186) HasPerClientBandwidthLimitDown() bool`

HasPerClientBandwidthLimitDown returns a boolean if a field has been set.

### GetVisible

`func (o *InlineResponse200186) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineResponse200186) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineResponse200186) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InlineResponse200186) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAvailableOnAllAps

`func (o *InlineResponse200186) GetAvailableOnAllAps() bool`

GetAvailableOnAllAps returns the AvailableOnAllAps field if non-nil, zero value otherwise.

### GetAvailableOnAllApsOk

`func (o *InlineResponse200186) GetAvailableOnAllApsOk() (*bool, bool)`

GetAvailableOnAllApsOk returns a tuple with the AvailableOnAllAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOnAllAps

`func (o *InlineResponse200186) SetAvailableOnAllAps(v bool)`

SetAvailableOnAllAps sets AvailableOnAllAps field to given value.

### HasAvailableOnAllAps

`func (o *InlineResponse200186) HasAvailableOnAllAps() bool`

HasAvailableOnAllAps returns a boolean if a field has been set.

### GetAvailabilityTags

`func (o *InlineResponse200186) GetAvailabilityTags() []string`

GetAvailabilityTags returns the AvailabilityTags field if non-nil, zero value otherwise.

### GetAvailabilityTagsOk

`func (o *InlineResponse200186) GetAvailabilityTagsOk() (*[]string, bool)`

GetAvailabilityTagsOk returns a tuple with the AvailabilityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTags

`func (o *InlineResponse200186) SetAvailabilityTags(v []string)`

SetAvailabilityTags sets AvailabilityTags field to given value.

### HasAvailabilityTags

`func (o *InlineResponse200186) HasAvailabilityTags() bool`

HasAvailabilityTags returns a boolean if a field has been set.

### GetPerSsidBandwidthLimitUp

`func (o *InlineResponse200186) GetPerSsidBandwidthLimitUp() int32`

GetPerSsidBandwidthLimitUp returns the PerSsidBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitUpOk

`func (o *InlineResponse200186) GetPerSsidBandwidthLimitUpOk() (*int32, bool)`

GetPerSsidBandwidthLimitUpOk returns a tuple with the PerSsidBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitUp

`func (o *InlineResponse200186) SetPerSsidBandwidthLimitUp(v int32)`

SetPerSsidBandwidthLimitUp sets PerSsidBandwidthLimitUp field to given value.

### HasPerSsidBandwidthLimitUp

`func (o *InlineResponse200186) HasPerSsidBandwidthLimitUp() bool`

HasPerSsidBandwidthLimitUp returns a boolean if a field has been set.

### GetPerSsidBandwidthLimitDown

`func (o *InlineResponse200186) GetPerSsidBandwidthLimitDown() int32`

GetPerSsidBandwidthLimitDown returns the PerSsidBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitDownOk

`func (o *InlineResponse200186) GetPerSsidBandwidthLimitDownOk() (*int32, bool)`

GetPerSsidBandwidthLimitDownOk returns a tuple with the PerSsidBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitDown

`func (o *InlineResponse200186) SetPerSsidBandwidthLimitDown(v int32)`

SetPerSsidBandwidthLimitDown sets PerSsidBandwidthLimitDown field to given value.

### HasPerSsidBandwidthLimitDown

`func (o *InlineResponse200186) HasPerSsidBandwidthLimitDown() bool`

HasPerSsidBandwidthLimitDown returns a boolean if a field has been set.

### GetMandatoryDhcpEnabled

`func (o *InlineResponse200186) GetMandatoryDhcpEnabled() bool`

GetMandatoryDhcpEnabled returns the MandatoryDhcpEnabled field if non-nil, zero value otherwise.

### GetMandatoryDhcpEnabledOk

`func (o *InlineResponse200186) GetMandatoryDhcpEnabledOk() (*bool, bool)`

GetMandatoryDhcpEnabledOk returns a tuple with the MandatoryDhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcpEnabled

`func (o *InlineResponse200186) SetMandatoryDhcpEnabled(v bool)`

SetMandatoryDhcpEnabled sets MandatoryDhcpEnabled field to given value.

### HasMandatoryDhcpEnabled

`func (o *InlineResponse200186) HasMandatoryDhcpEnabled() bool`

HasMandatoryDhcpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


