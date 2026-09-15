# InlineResponse200247

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | Unique identifier of the SSID | 
**Name** | **string** | The name of the SSID | 
**Enabled** | **bool** | Whether or not the SSID is enabled | 
**WlanIdentifier** | Pointer to **int32** | WLAN identifier for the SSID. Only present for Meraki admin users. | [optional] 
**SplashPage** | **string** | The type of splash page for the SSID | 
**SsidAdminAccessible** | **bool** | SSID Administrator access status | 
**LocalAuth** | Pointer to **bool** | Extended local auth flag for Enterprise NAC. Only present when authMode is &#39;8021x-nac&#39;. | [optional] 
**AuthMode** | **string** | The association control method for the SSID | 
**Psk** | Pointer to **string** | The passkey for the SSID. This param is only valid if the authMode is &#39;psk&#39;. Read-only admins receive a masked value. Only present when authMode is &#39;psk&#39;. | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID. Only present when authMode requires encryption. | [optional] 
**WpaEncryptionMode** | Pointer to **string** | The types of WPA encryption. Only present when authMode requires encryption. | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsRadiusServers**](NetworksNetworkIdWirelessSsidsRadiusServers.md) | List of RADIUS 802.1X servers to be used for authentication. Only present when the SSID is configured to use RADIUS. | [optional] 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsRadiusAccountingServers**](NetworksNetworkIdWirelessSsidsRadiusAccountingServers.md) | List of RADIUS accounting 802.1X servers to be used for authentication. Only present when RADIUS accounting is enabled. | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Whether or not RADIUS accounting is enabled. Only present when the SSID is configured to use RADIUS. | [optional] 
**RadiusEnabled** | Pointer to **bool** | Whether RADIUS authentication is enabled. Only present when the SSID is configured to use RADIUS. | [optional] 
**RadiusAttributeForGroupPolicies** | Pointer to **string** | RADIUS attribute used to look up group policies. Only present when the SSID is configured to use RADIUS and the network supports group policies. | [optional] 
**RadiusFailoverPolicy** | Pointer to **string** | Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable. Only present when the SSID uses a RADIUS-based splash page. | [optional] 
**RadiusLoadBalancingPolicy** | Pointer to **string** | Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts. Only present when the SSID uses a RADIUS-based splash page. | [optional] 
**IpAssignmentMode** | **string** | The client IP assignment mode | 
**AdminSplashUrl** | Pointer to **string** | URL for the admin splash page. Only present when a splash page is configured. | [optional] 
**SplashTimeout** | Pointer to **string** | Splash page timeout. Only present when a splash page is configured. | [optional] 
**WalledGardenEnabled** | Pointer to **bool** | Allow users to access a configurable list of IP ranges prior to sign-on. Only present when a splash page is configured. | [optional] 
**WalledGardenRanges** | Pointer to **[]string** | Domain names and IP address ranges available in Walled Garden mode. Only present when walledGardenEnabled is true and ranges are configured. | [optional] 
**MinBitrate** | **float32** | The minimum bitrate in Mbps of this SSID in the default indoor RF profile | 
**BandSelection** | **string** | The client-serving radio frequencies of this SSID in the default indoor RF profile | 
**PerClientBandwidthLimitUp** | **int32** | The upload bandwidth limit in Kbps. (0 represents no limit.) | 
**PerClientBandwidthLimitDown** | **int32** | The download bandwidth limit in Kbps. (0 represents no limit.) | 
**Visible** | **bool** | Whether the SSID is advertised or hidden by the AP | 
**AvailableOnAllAps** | **bool** | Whether all APs broadcast the SSID or if it&#39;s restricted to APs matching any availability tags | 
**AvailabilityTags** | **[]string** | List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list | 
**PerSsidBandwidthLimitUp** | **int32** | The total upload bandwidth limit in Kbps (0 represents no limit) | 
**PerSsidBandwidthLimitDown** | **int32** | The total download bandwidth limit in Kbps (0 represents no limit) | 
**AdaptivePolicyGroupId** | Pointer to **string** | Adaptive policy group ID this SSID is assigned to. | [optional] 
**MandatoryDhcpEnabled** | **bool** | Whether clients connecting to this SSID must use the IP address assigned by the DHCP server | 
**Dot11w** | Pointer to [**NetworksNetworkIdWirelessSsidsDot11w**](NetworksNetworkIdWirelessSsidsDot11w.md) |  | [optional] 
**Dot11r** | Pointer to [**NetworksNetworkIdWirelessSsidsDot11r**](NetworksNetworkIdWirelessSsidsDot11r.md) |  | [optional] 
**EnterpriseAdminAccess** | Pointer to **string** | Whether or not an SSID is accessible by &#39;enterprise&#39; administrators. Only present when the organization supports enterprise admins. | [optional] 
**RadiusTestingEnabled** | Pointer to **bool** | If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity &#39;meraki_8021x_test&#39; to ensure that the RADIUS servers are reachable. Only present when the SSID is configured to use RADIUS. | [optional] 
**RadiusCalledStationId** | Pointer to **string** | The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$). Only present when configured and the SSID uses RADIUS-based auth. | [optional] 
**RadiusAuthenticationNasId** | Pointer to **string** | The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$). Only present when configured and the SSID uses RADIUS-based auth. | [optional] 
**RadiusServerTimeout** | Pointer to **int32** | The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds). Only present when the SSID uses RADIUS-based auth. | [optional] 
**RadiusServerAttemptsLimit** | Pointer to **int32** | The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5). Only present when the SSID uses RADIUS-based auth. | [optional] 
**RadiusFallbackEnabled** | Pointer to **bool** | Whether or not higher priority RADIUS servers should be retried after 60 seconds. Only present when the SSID uses RADIUS-based auth. | [optional] 
**RadiusProxyEnabled** | Pointer to **bool** | If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers. Only present when the SSID uses RADIUS-based auth. | [optional] 
**RadiusCoaEnabled** | Pointer to **bool** | If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server. Only present when the SSID uses RADIUS-based auth. | [optional] 
**RadiusOverride** | Pointer to **bool** | If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is &#39;NAT mode&#39;. Only present when the SSID uses RADIUS-based auth and the network supports group policies. | [optional] 
**UseVlanTagging** | Pointer to **bool** | Whether or not traffic should be directed to use specific VLANs. Only present when the ipAssignmentMode is &#39;Bridge mode&#39; or &#39;Layer 3 roaming&#39;. | [optional] 
**DefaultVlanId** | Pointer to **int32** | The default VLAN ID used for &#39;all other APs&#39;. Only present when the ipAssignmentMode is &#39;Bridge mode&#39; or &#39;Layer 3 roaming&#39; and VLAN tagging is enabled. | [optional] 
**AdultContentFilteringEnabled** | Pointer to **bool** | Boolean indicating whether or not adult content will be blocked. Only present when ipAssignmentMode is &#39;NAT mode&#39;. | [optional] 
**DnsRewrite** | Pointer to [**NetworksNetworkIdWirelessSsidsDnsRewrite**](NetworksNetworkIdWirelessSsidsDnsRewrite.md) |  | [optional] 
**Gre** | Pointer to [**NetworksNetworkIdWirelessSsidsGre**](NetworksNetworkIdWirelessSsidsGre.md) |  | [optional] 
**CampusGateway** | Pointer to [**NetworksNetworkIdWirelessSsidsCampusGateway**](NetworksNetworkIdWirelessSsidsCampusGateway.md) |  | [optional] 
**LanIsolationEnabled** | Pointer to **bool** | Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is &#39;Bridge mode&#39;. | [optional] 
**Oauth** | Pointer to [**NetworksNetworkIdWirelessSsidsOauth**](NetworksNetworkIdWirelessSsidsOauth.md) |  | [optional] 
**LocalAuthFallback** | Pointer to [**NetworksNetworkIdWirelessSsidsLocalAuthFallback**](NetworksNetworkIdWirelessSsidsLocalAuthFallback.md) |  | [optional] 
**NamedVlans** | Pointer to [**NetworksNetworkIdWirelessSsidsNamedVlans**](NetworksNetworkIdWirelessSsidsNamedVlans.md) |  | [optional] 
**SpeedBurst** | [**NetworksNetworkIdWirelessSsidsSpeedBurst**](NetworksNetworkIdWirelessSsidsSpeedBurst.md) |  | 
**WifiPersonalNetworkEnabled** | Pointer to **bool** | Boolean indicating whether or not Wi-Fi Personal Network is enabled for this SSID. Only present when the network supports Wi-Fi Personal Network. | [optional] 
**Security** | Pointer to [**NetworksNetworkIdWirelessSsidsSecurity**](NetworksNetworkIdWirelessSsidsSecurity.md) |  | [optional] 

## Methods

### NewInlineResponse200247

`func NewInlineResponse200247(number int32, name string, enabled bool, splashPage string, ssidAdminAccessible bool, authMode string, ipAssignmentMode string, minBitrate float32, bandSelection string, perClientBandwidthLimitUp int32, perClientBandwidthLimitDown int32, visible bool, availableOnAllAps bool, availabilityTags []string, perSsidBandwidthLimitUp int32, perSsidBandwidthLimitDown int32, mandatoryDhcpEnabled bool, speedBurst NetworksNetworkIdWirelessSsidsSpeedBurst, ) *InlineResponse200247`

NewInlineResponse200247 instantiates a new InlineResponse200247 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200247WithDefaults

`func NewInlineResponse200247WithDefaults() *InlineResponse200247`

NewInlineResponse200247WithDefaults instantiates a new InlineResponse200247 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse200247) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200247) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200247) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetName

`func (o *InlineResponse200247) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200247) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200247) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *InlineResponse200247) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200247) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200247) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWlanIdentifier

`func (o *InlineResponse200247) GetWlanIdentifier() int32`

GetWlanIdentifier returns the WlanIdentifier field if non-nil, zero value otherwise.

### GetWlanIdentifierOk

`func (o *InlineResponse200247) GetWlanIdentifierOk() (*int32, bool)`

GetWlanIdentifierOk returns a tuple with the WlanIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanIdentifier

`func (o *InlineResponse200247) SetWlanIdentifier(v int32)`

SetWlanIdentifier sets WlanIdentifier field to given value.

### HasWlanIdentifier

`func (o *InlineResponse200247) HasWlanIdentifier() bool`

HasWlanIdentifier returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineResponse200247) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineResponse200247) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineResponse200247) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.


### GetSsidAdminAccessible

`func (o *InlineResponse200247) GetSsidAdminAccessible() bool`

GetSsidAdminAccessible returns the SsidAdminAccessible field if non-nil, zero value otherwise.

### GetSsidAdminAccessibleOk

`func (o *InlineResponse200247) GetSsidAdminAccessibleOk() (*bool, bool)`

GetSsidAdminAccessibleOk returns a tuple with the SsidAdminAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidAdminAccessible

`func (o *InlineResponse200247) SetSsidAdminAccessible(v bool)`

SetSsidAdminAccessible sets SsidAdminAccessible field to given value.


### GetLocalAuth

`func (o *InlineResponse200247) GetLocalAuth() bool`

GetLocalAuth returns the LocalAuth field if non-nil, zero value otherwise.

### GetLocalAuthOk

`func (o *InlineResponse200247) GetLocalAuthOk() (*bool, bool)`

GetLocalAuthOk returns a tuple with the LocalAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuth

`func (o *InlineResponse200247) SetLocalAuth(v bool)`

SetLocalAuth sets LocalAuth field to given value.

### HasLocalAuth

`func (o *InlineResponse200247) HasLocalAuth() bool`

HasLocalAuth returns a boolean if a field has been set.

### GetAuthMode

`func (o *InlineResponse200247) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InlineResponse200247) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InlineResponse200247) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.


### GetPsk

`func (o *InlineResponse200247) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *InlineResponse200247) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *InlineResponse200247) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *InlineResponse200247) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *InlineResponse200247) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *InlineResponse200247) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *InlineResponse200247) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *InlineResponse200247) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *InlineResponse200247) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *InlineResponse200247) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *InlineResponse200247) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *InlineResponse200247) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineResponse200247) GetRadiusServers() []NetworksNetworkIdWirelessSsidsRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineResponse200247) GetRadiusServersOk() (*[]NetworksNetworkIdWirelessSsidsRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineResponse200247) SetRadiusServers(v []NetworksNetworkIdWirelessSsidsRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineResponse200247) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *InlineResponse200247) GetRadiusAccountingServers() []NetworksNetworkIdWirelessSsidsRadiusAccountingServers`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineResponse200247) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdWirelessSsidsRadiusAccountingServers, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineResponse200247) SetRadiusAccountingServers(v []NetworksNetworkIdWirelessSsidsRadiusAccountingServers)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineResponse200247) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *InlineResponse200247) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineResponse200247) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineResponse200247) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *InlineResponse200247) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusEnabled

`func (o *InlineResponse200247) GetRadiusEnabled() bool`

GetRadiusEnabled returns the RadiusEnabled field if non-nil, zero value otherwise.

### GetRadiusEnabledOk

`func (o *InlineResponse200247) GetRadiusEnabledOk() (*bool, bool)`

GetRadiusEnabledOk returns a tuple with the RadiusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusEnabled

`func (o *InlineResponse200247) SetRadiusEnabled(v bool)`

SetRadiusEnabled sets RadiusEnabled field to given value.

### HasRadiusEnabled

`func (o *InlineResponse200247) HasRadiusEnabled() bool`

HasRadiusEnabled returns a boolean if a field has been set.

### GetRadiusAttributeForGroupPolicies

`func (o *InlineResponse200247) GetRadiusAttributeForGroupPolicies() string`

GetRadiusAttributeForGroupPolicies returns the RadiusAttributeForGroupPolicies field if non-nil, zero value otherwise.

### GetRadiusAttributeForGroupPoliciesOk

`func (o *InlineResponse200247) GetRadiusAttributeForGroupPoliciesOk() (*string, bool)`

GetRadiusAttributeForGroupPoliciesOk returns a tuple with the RadiusAttributeForGroupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttributeForGroupPolicies

`func (o *InlineResponse200247) SetRadiusAttributeForGroupPolicies(v string)`

SetRadiusAttributeForGroupPolicies sets RadiusAttributeForGroupPolicies field to given value.

### HasRadiusAttributeForGroupPolicies

`func (o *InlineResponse200247) HasRadiusAttributeForGroupPolicies() bool`

HasRadiusAttributeForGroupPolicies returns a boolean if a field has been set.

### GetRadiusFailoverPolicy

`func (o *InlineResponse200247) GetRadiusFailoverPolicy() string`

GetRadiusFailoverPolicy returns the RadiusFailoverPolicy field if non-nil, zero value otherwise.

### GetRadiusFailoverPolicyOk

`func (o *InlineResponse200247) GetRadiusFailoverPolicyOk() (*string, bool)`

GetRadiusFailoverPolicyOk returns a tuple with the RadiusFailoverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusFailoverPolicy

`func (o *InlineResponse200247) SetRadiusFailoverPolicy(v string)`

SetRadiusFailoverPolicy sets RadiusFailoverPolicy field to given value.

### HasRadiusFailoverPolicy

`func (o *InlineResponse200247) HasRadiusFailoverPolicy() bool`

HasRadiusFailoverPolicy returns a boolean if a field has been set.

### GetRadiusLoadBalancingPolicy

`func (o *InlineResponse200247) GetRadiusLoadBalancingPolicy() string`

GetRadiusLoadBalancingPolicy returns the RadiusLoadBalancingPolicy field if non-nil, zero value otherwise.

### GetRadiusLoadBalancingPolicyOk

`func (o *InlineResponse200247) GetRadiusLoadBalancingPolicyOk() (*string, bool)`

GetRadiusLoadBalancingPolicyOk returns a tuple with the RadiusLoadBalancingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusLoadBalancingPolicy

`func (o *InlineResponse200247) SetRadiusLoadBalancingPolicy(v string)`

SetRadiusLoadBalancingPolicy sets RadiusLoadBalancingPolicy field to given value.

### HasRadiusLoadBalancingPolicy

`func (o *InlineResponse200247) HasRadiusLoadBalancingPolicy() bool`

HasRadiusLoadBalancingPolicy returns a boolean if a field has been set.

### GetIpAssignmentMode

`func (o *InlineResponse200247) GetIpAssignmentMode() string`

GetIpAssignmentMode returns the IpAssignmentMode field if non-nil, zero value otherwise.

### GetIpAssignmentModeOk

`func (o *InlineResponse200247) GetIpAssignmentModeOk() (*string, bool)`

GetIpAssignmentModeOk returns a tuple with the IpAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignmentMode

`func (o *InlineResponse200247) SetIpAssignmentMode(v string)`

SetIpAssignmentMode sets IpAssignmentMode field to given value.


### GetAdminSplashUrl

`func (o *InlineResponse200247) GetAdminSplashUrl() string`

GetAdminSplashUrl returns the AdminSplashUrl field if non-nil, zero value otherwise.

### GetAdminSplashUrlOk

`func (o *InlineResponse200247) GetAdminSplashUrlOk() (*string, bool)`

GetAdminSplashUrlOk returns a tuple with the AdminSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSplashUrl

`func (o *InlineResponse200247) SetAdminSplashUrl(v string)`

SetAdminSplashUrl sets AdminSplashUrl field to given value.

### HasAdminSplashUrl

`func (o *InlineResponse200247) HasAdminSplashUrl() bool`

HasAdminSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineResponse200247) GetSplashTimeout() string`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineResponse200247) GetSplashTimeoutOk() (*string, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineResponse200247) SetSplashTimeout(v string)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineResponse200247) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetWalledGardenEnabled

`func (o *InlineResponse200247) GetWalledGardenEnabled() bool`

GetWalledGardenEnabled returns the WalledGardenEnabled field if non-nil, zero value otherwise.

### GetWalledGardenEnabledOk

`func (o *InlineResponse200247) GetWalledGardenEnabledOk() (*bool, bool)`

GetWalledGardenEnabledOk returns a tuple with the WalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenEnabled

`func (o *InlineResponse200247) SetWalledGardenEnabled(v bool)`

SetWalledGardenEnabled sets WalledGardenEnabled field to given value.

### HasWalledGardenEnabled

`func (o *InlineResponse200247) HasWalledGardenEnabled() bool`

HasWalledGardenEnabled returns a boolean if a field has been set.

### GetWalledGardenRanges

`func (o *InlineResponse200247) GetWalledGardenRanges() []string`

GetWalledGardenRanges returns the WalledGardenRanges field if non-nil, zero value otherwise.

### GetWalledGardenRangesOk

`func (o *InlineResponse200247) GetWalledGardenRangesOk() (*[]string, bool)`

GetWalledGardenRangesOk returns a tuple with the WalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenRanges

`func (o *InlineResponse200247) SetWalledGardenRanges(v []string)`

SetWalledGardenRanges sets WalledGardenRanges field to given value.

### HasWalledGardenRanges

`func (o *InlineResponse200247) HasWalledGardenRanges() bool`

HasWalledGardenRanges returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse200247) GetMinBitrate() float32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse200247) GetMinBitrateOk() (*float32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse200247) SetMinBitrate(v float32)`

SetMinBitrate sets MinBitrate field to given value.


### GetBandSelection

`func (o *InlineResponse200247) GetBandSelection() string`

GetBandSelection returns the BandSelection field if non-nil, zero value otherwise.

### GetBandSelectionOk

`func (o *InlineResponse200247) GetBandSelectionOk() (*string, bool)`

GetBandSelectionOk returns a tuple with the BandSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelection

`func (o *InlineResponse200247) SetBandSelection(v string)`

SetBandSelection sets BandSelection field to given value.


### GetPerClientBandwidthLimitUp

`func (o *InlineResponse200247) GetPerClientBandwidthLimitUp() int32`

GetPerClientBandwidthLimitUp returns the PerClientBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitUpOk

`func (o *InlineResponse200247) GetPerClientBandwidthLimitUpOk() (*int32, bool)`

GetPerClientBandwidthLimitUpOk returns a tuple with the PerClientBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitUp

`func (o *InlineResponse200247) SetPerClientBandwidthLimitUp(v int32)`

SetPerClientBandwidthLimitUp sets PerClientBandwidthLimitUp field to given value.


### GetPerClientBandwidthLimitDown

`func (o *InlineResponse200247) GetPerClientBandwidthLimitDown() int32`

GetPerClientBandwidthLimitDown returns the PerClientBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitDownOk

`func (o *InlineResponse200247) GetPerClientBandwidthLimitDownOk() (*int32, bool)`

GetPerClientBandwidthLimitDownOk returns a tuple with the PerClientBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitDown

`func (o *InlineResponse200247) SetPerClientBandwidthLimitDown(v int32)`

SetPerClientBandwidthLimitDown sets PerClientBandwidthLimitDown field to given value.


### GetVisible

`func (o *InlineResponse200247) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineResponse200247) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineResponse200247) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetAvailableOnAllAps

`func (o *InlineResponse200247) GetAvailableOnAllAps() bool`

GetAvailableOnAllAps returns the AvailableOnAllAps field if non-nil, zero value otherwise.

### GetAvailableOnAllApsOk

`func (o *InlineResponse200247) GetAvailableOnAllApsOk() (*bool, bool)`

GetAvailableOnAllApsOk returns a tuple with the AvailableOnAllAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOnAllAps

`func (o *InlineResponse200247) SetAvailableOnAllAps(v bool)`

SetAvailableOnAllAps sets AvailableOnAllAps field to given value.


### GetAvailabilityTags

`func (o *InlineResponse200247) GetAvailabilityTags() []string`

GetAvailabilityTags returns the AvailabilityTags field if non-nil, zero value otherwise.

### GetAvailabilityTagsOk

`func (o *InlineResponse200247) GetAvailabilityTagsOk() (*[]string, bool)`

GetAvailabilityTagsOk returns a tuple with the AvailabilityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTags

`func (o *InlineResponse200247) SetAvailabilityTags(v []string)`

SetAvailabilityTags sets AvailabilityTags field to given value.


### GetPerSsidBandwidthLimitUp

`func (o *InlineResponse200247) GetPerSsidBandwidthLimitUp() int32`

GetPerSsidBandwidthLimitUp returns the PerSsidBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitUpOk

`func (o *InlineResponse200247) GetPerSsidBandwidthLimitUpOk() (*int32, bool)`

GetPerSsidBandwidthLimitUpOk returns a tuple with the PerSsidBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitUp

`func (o *InlineResponse200247) SetPerSsidBandwidthLimitUp(v int32)`

SetPerSsidBandwidthLimitUp sets PerSsidBandwidthLimitUp field to given value.


### GetPerSsidBandwidthLimitDown

`func (o *InlineResponse200247) GetPerSsidBandwidthLimitDown() int32`

GetPerSsidBandwidthLimitDown returns the PerSsidBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitDownOk

`func (o *InlineResponse200247) GetPerSsidBandwidthLimitDownOk() (*int32, bool)`

GetPerSsidBandwidthLimitDownOk returns a tuple with the PerSsidBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitDown

`func (o *InlineResponse200247) SetPerSsidBandwidthLimitDown(v int32)`

SetPerSsidBandwidthLimitDown sets PerSsidBandwidthLimitDown field to given value.


### GetAdaptivePolicyGroupId

`func (o *InlineResponse200247) GetAdaptivePolicyGroupId() string`

GetAdaptivePolicyGroupId returns the AdaptivePolicyGroupId field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupIdOk

`func (o *InlineResponse200247) GetAdaptivePolicyGroupIdOk() (*string, bool)`

GetAdaptivePolicyGroupIdOk returns a tuple with the AdaptivePolicyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroupId

`func (o *InlineResponse200247) SetAdaptivePolicyGroupId(v string)`

SetAdaptivePolicyGroupId sets AdaptivePolicyGroupId field to given value.

### HasAdaptivePolicyGroupId

`func (o *InlineResponse200247) HasAdaptivePolicyGroupId() bool`

HasAdaptivePolicyGroupId returns a boolean if a field has been set.

### GetMandatoryDhcpEnabled

`func (o *InlineResponse200247) GetMandatoryDhcpEnabled() bool`

GetMandatoryDhcpEnabled returns the MandatoryDhcpEnabled field if non-nil, zero value otherwise.

### GetMandatoryDhcpEnabledOk

`func (o *InlineResponse200247) GetMandatoryDhcpEnabledOk() (*bool, bool)`

GetMandatoryDhcpEnabledOk returns a tuple with the MandatoryDhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcpEnabled

`func (o *InlineResponse200247) SetMandatoryDhcpEnabled(v bool)`

SetMandatoryDhcpEnabled sets MandatoryDhcpEnabled field to given value.


### GetDot11w

`func (o *InlineResponse200247) GetDot11w() NetworksNetworkIdWirelessSsidsDot11w`

GetDot11w returns the Dot11w field if non-nil, zero value otherwise.

### GetDot11wOk

`func (o *InlineResponse200247) GetDot11wOk() (*NetworksNetworkIdWirelessSsidsDot11w, bool)`

GetDot11wOk returns a tuple with the Dot11w field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11w

`func (o *InlineResponse200247) SetDot11w(v NetworksNetworkIdWirelessSsidsDot11w)`

SetDot11w sets Dot11w field to given value.

### HasDot11w

`func (o *InlineResponse200247) HasDot11w() bool`

HasDot11w returns a boolean if a field has been set.

### GetDot11r

`func (o *InlineResponse200247) GetDot11r() NetworksNetworkIdWirelessSsidsDot11r`

GetDot11r returns the Dot11r field if non-nil, zero value otherwise.

### GetDot11rOk

`func (o *InlineResponse200247) GetDot11rOk() (*NetworksNetworkIdWirelessSsidsDot11r, bool)`

GetDot11rOk returns a tuple with the Dot11r field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11r

`func (o *InlineResponse200247) SetDot11r(v NetworksNetworkIdWirelessSsidsDot11r)`

SetDot11r sets Dot11r field to given value.

### HasDot11r

`func (o *InlineResponse200247) HasDot11r() bool`

HasDot11r returns a boolean if a field has been set.

### GetEnterpriseAdminAccess

`func (o *InlineResponse200247) GetEnterpriseAdminAccess() string`

GetEnterpriseAdminAccess returns the EnterpriseAdminAccess field if non-nil, zero value otherwise.

### GetEnterpriseAdminAccessOk

`func (o *InlineResponse200247) GetEnterpriseAdminAccessOk() (*string, bool)`

GetEnterpriseAdminAccessOk returns a tuple with the EnterpriseAdminAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseAdminAccess

`func (o *InlineResponse200247) SetEnterpriseAdminAccess(v string)`

SetEnterpriseAdminAccess sets EnterpriseAdminAccess field to given value.

### HasEnterpriseAdminAccess

`func (o *InlineResponse200247) HasEnterpriseAdminAccess() bool`

HasEnterpriseAdminAccess returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineResponse200247) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineResponse200247) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineResponse200247) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.

### HasRadiusTestingEnabled

`func (o *InlineResponse200247) HasRadiusTestingEnabled() bool`

HasRadiusTestingEnabled returns a boolean if a field has been set.

### GetRadiusCalledStationId

`func (o *InlineResponse200247) GetRadiusCalledStationId() string`

GetRadiusCalledStationId returns the RadiusCalledStationId field if non-nil, zero value otherwise.

### GetRadiusCalledStationIdOk

`func (o *InlineResponse200247) GetRadiusCalledStationIdOk() (*string, bool)`

GetRadiusCalledStationIdOk returns a tuple with the RadiusCalledStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCalledStationId

`func (o *InlineResponse200247) SetRadiusCalledStationId(v string)`

SetRadiusCalledStationId sets RadiusCalledStationId field to given value.

### HasRadiusCalledStationId

`func (o *InlineResponse200247) HasRadiusCalledStationId() bool`

HasRadiusCalledStationId returns a boolean if a field has been set.

### GetRadiusAuthenticationNasId

`func (o *InlineResponse200247) GetRadiusAuthenticationNasId() string`

GetRadiusAuthenticationNasId returns the RadiusAuthenticationNasId field if non-nil, zero value otherwise.

### GetRadiusAuthenticationNasIdOk

`func (o *InlineResponse200247) GetRadiusAuthenticationNasIdOk() (*string, bool)`

GetRadiusAuthenticationNasIdOk returns a tuple with the RadiusAuthenticationNasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAuthenticationNasId

`func (o *InlineResponse200247) SetRadiusAuthenticationNasId(v string)`

SetRadiusAuthenticationNasId sets RadiusAuthenticationNasId field to given value.

### HasRadiusAuthenticationNasId

`func (o *InlineResponse200247) HasRadiusAuthenticationNasId() bool`

HasRadiusAuthenticationNasId returns a boolean if a field has been set.

### GetRadiusServerTimeout

`func (o *InlineResponse200247) GetRadiusServerTimeout() int32`

GetRadiusServerTimeout returns the RadiusServerTimeout field if non-nil, zero value otherwise.

### GetRadiusServerTimeoutOk

`func (o *InlineResponse200247) GetRadiusServerTimeoutOk() (*int32, bool)`

GetRadiusServerTimeoutOk returns a tuple with the RadiusServerTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServerTimeout

`func (o *InlineResponse200247) SetRadiusServerTimeout(v int32)`

SetRadiusServerTimeout sets RadiusServerTimeout field to given value.

### HasRadiusServerTimeout

`func (o *InlineResponse200247) HasRadiusServerTimeout() bool`

HasRadiusServerTimeout returns a boolean if a field has been set.

### GetRadiusServerAttemptsLimit

`func (o *InlineResponse200247) GetRadiusServerAttemptsLimit() int32`

GetRadiusServerAttemptsLimit returns the RadiusServerAttemptsLimit field if non-nil, zero value otherwise.

### GetRadiusServerAttemptsLimitOk

`func (o *InlineResponse200247) GetRadiusServerAttemptsLimitOk() (*int32, bool)`

GetRadiusServerAttemptsLimitOk returns a tuple with the RadiusServerAttemptsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServerAttemptsLimit

`func (o *InlineResponse200247) SetRadiusServerAttemptsLimit(v int32)`

SetRadiusServerAttemptsLimit sets RadiusServerAttemptsLimit field to given value.

### HasRadiusServerAttemptsLimit

`func (o *InlineResponse200247) HasRadiusServerAttemptsLimit() bool`

HasRadiusServerAttemptsLimit returns a boolean if a field has been set.

### GetRadiusFallbackEnabled

`func (o *InlineResponse200247) GetRadiusFallbackEnabled() bool`

GetRadiusFallbackEnabled returns the RadiusFallbackEnabled field if non-nil, zero value otherwise.

### GetRadiusFallbackEnabledOk

`func (o *InlineResponse200247) GetRadiusFallbackEnabledOk() (*bool, bool)`

GetRadiusFallbackEnabledOk returns a tuple with the RadiusFallbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusFallbackEnabled

`func (o *InlineResponse200247) SetRadiusFallbackEnabled(v bool)`

SetRadiusFallbackEnabled sets RadiusFallbackEnabled field to given value.

### HasRadiusFallbackEnabled

`func (o *InlineResponse200247) HasRadiusFallbackEnabled() bool`

HasRadiusFallbackEnabled returns a boolean if a field has been set.

### GetRadiusProxyEnabled

`func (o *InlineResponse200247) GetRadiusProxyEnabled() bool`

GetRadiusProxyEnabled returns the RadiusProxyEnabled field if non-nil, zero value otherwise.

### GetRadiusProxyEnabledOk

`func (o *InlineResponse200247) GetRadiusProxyEnabledOk() (*bool, bool)`

GetRadiusProxyEnabledOk returns a tuple with the RadiusProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProxyEnabled

`func (o *InlineResponse200247) SetRadiusProxyEnabled(v bool)`

SetRadiusProxyEnabled sets RadiusProxyEnabled field to given value.

### HasRadiusProxyEnabled

`func (o *InlineResponse200247) HasRadiusProxyEnabled() bool`

HasRadiusProxyEnabled returns a boolean if a field has been set.

### GetRadiusCoaEnabled

`func (o *InlineResponse200247) GetRadiusCoaEnabled() bool`

GetRadiusCoaEnabled returns the RadiusCoaEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaEnabledOk

`func (o *InlineResponse200247) GetRadiusCoaEnabledOk() (*bool, bool)`

GetRadiusCoaEnabledOk returns a tuple with the RadiusCoaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaEnabled

`func (o *InlineResponse200247) SetRadiusCoaEnabled(v bool)`

SetRadiusCoaEnabled sets RadiusCoaEnabled field to given value.

### HasRadiusCoaEnabled

`func (o *InlineResponse200247) HasRadiusCoaEnabled() bool`

HasRadiusCoaEnabled returns a boolean if a field has been set.

### GetRadiusOverride

`func (o *InlineResponse200247) GetRadiusOverride() bool`

GetRadiusOverride returns the RadiusOverride field if non-nil, zero value otherwise.

### GetRadiusOverrideOk

`func (o *InlineResponse200247) GetRadiusOverrideOk() (*bool, bool)`

GetRadiusOverrideOk returns a tuple with the RadiusOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusOverride

`func (o *InlineResponse200247) SetRadiusOverride(v bool)`

SetRadiusOverride sets RadiusOverride field to given value.

### HasRadiusOverride

`func (o *InlineResponse200247) HasRadiusOverride() bool`

HasRadiusOverride returns a boolean if a field has been set.

### GetUseVlanTagging

`func (o *InlineResponse200247) GetUseVlanTagging() bool`

GetUseVlanTagging returns the UseVlanTagging field if non-nil, zero value otherwise.

### GetUseVlanTaggingOk

`func (o *InlineResponse200247) GetUseVlanTaggingOk() (*bool, bool)`

GetUseVlanTaggingOk returns a tuple with the UseVlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVlanTagging

`func (o *InlineResponse200247) SetUseVlanTagging(v bool)`

SetUseVlanTagging sets UseVlanTagging field to given value.

### HasUseVlanTagging

`func (o *InlineResponse200247) HasUseVlanTagging() bool`

HasUseVlanTagging returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *InlineResponse200247) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *InlineResponse200247) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *InlineResponse200247) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *InlineResponse200247) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetAdultContentFilteringEnabled

`func (o *InlineResponse200247) GetAdultContentFilteringEnabled() bool`

GetAdultContentFilteringEnabled returns the AdultContentFilteringEnabled field if non-nil, zero value otherwise.

### GetAdultContentFilteringEnabledOk

`func (o *InlineResponse200247) GetAdultContentFilteringEnabledOk() (*bool, bool)`

GetAdultContentFilteringEnabledOk returns a tuple with the AdultContentFilteringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContentFilteringEnabled

`func (o *InlineResponse200247) SetAdultContentFilteringEnabled(v bool)`

SetAdultContentFilteringEnabled sets AdultContentFilteringEnabled field to given value.

### HasAdultContentFilteringEnabled

`func (o *InlineResponse200247) HasAdultContentFilteringEnabled() bool`

HasAdultContentFilteringEnabled returns a boolean if a field has been set.

### GetDnsRewrite

`func (o *InlineResponse200247) GetDnsRewrite() NetworksNetworkIdWirelessSsidsDnsRewrite`

GetDnsRewrite returns the DnsRewrite field if non-nil, zero value otherwise.

### GetDnsRewriteOk

`func (o *InlineResponse200247) GetDnsRewriteOk() (*NetworksNetworkIdWirelessSsidsDnsRewrite, bool)`

GetDnsRewriteOk returns a tuple with the DnsRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRewrite

`func (o *InlineResponse200247) SetDnsRewrite(v NetworksNetworkIdWirelessSsidsDnsRewrite)`

SetDnsRewrite sets DnsRewrite field to given value.

### HasDnsRewrite

`func (o *InlineResponse200247) HasDnsRewrite() bool`

HasDnsRewrite returns a boolean if a field has been set.

### GetGre

`func (o *InlineResponse200247) GetGre() NetworksNetworkIdWirelessSsidsGre`

GetGre returns the Gre field if non-nil, zero value otherwise.

### GetGreOk

`func (o *InlineResponse200247) GetGreOk() (*NetworksNetworkIdWirelessSsidsGre, bool)`

GetGreOk returns a tuple with the Gre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGre

`func (o *InlineResponse200247) SetGre(v NetworksNetworkIdWirelessSsidsGre)`

SetGre sets Gre field to given value.

### HasGre

`func (o *InlineResponse200247) HasGre() bool`

HasGre returns a boolean if a field has been set.

### GetCampusGateway

`func (o *InlineResponse200247) GetCampusGateway() NetworksNetworkIdWirelessSsidsCampusGateway`

GetCampusGateway returns the CampusGateway field if non-nil, zero value otherwise.

### GetCampusGatewayOk

`func (o *InlineResponse200247) GetCampusGatewayOk() (*NetworksNetworkIdWirelessSsidsCampusGateway, bool)`

GetCampusGatewayOk returns a tuple with the CampusGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampusGateway

`func (o *InlineResponse200247) SetCampusGateway(v NetworksNetworkIdWirelessSsidsCampusGateway)`

SetCampusGateway sets CampusGateway field to given value.

### HasCampusGateway

`func (o *InlineResponse200247) HasCampusGateway() bool`

HasCampusGateway returns a boolean if a field has been set.

### GetLanIsolationEnabled

`func (o *InlineResponse200247) GetLanIsolationEnabled() bool`

GetLanIsolationEnabled returns the LanIsolationEnabled field if non-nil, zero value otherwise.

### GetLanIsolationEnabledOk

`func (o *InlineResponse200247) GetLanIsolationEnabledOk() (*bool, bool)`

GetLanIsolationEnabledOk returns a tuple with the LanIsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIsolationEnabled

`func (o *InlineResponse200247) SetLanIsolationEnabled(v bool)`

SetLanIsolationEnabled sets LanIsolationEnabled field to given value.

### HasLanIsolationEnabled

`func (o *InlineResponse200247) HasLanIsolationEnabled() bool`

HasLanIsolationEnabled returns a boolean if a field has been set.

### GetOauth

`func (o *InlineResponse200247) GetOauth() NetworksNetworkIdWirelessSsidsOauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *InlineResponse200247) GetOauthOk() (*NetworksNetworkIdWirelessSsidsOauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *InlineResponse200247) SetOauth(v NetworksNetworkIdWirelessSsidsOauth)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *InlineResponse200247) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetLocalAuthFallback

`func (o *InlineResponse200247) GetLocalAuthFallback() NetworksNetworkIdWirelessSsidsLocalAuthFallback`

GetLocalAuthFallback returns the LocalAuthFallback field if non-nil, zero value otherwise.

### GetLocalAuthFallbackOk

`func (o *InlineResponse200247) GetLocalAuthFallbackOk() (*NetworksNetworkIdWirelessSsidsLocalAuthFallback, bool)`

GetLocalAuthFallbackOk returns a tuple with the LocalAuthFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuthFallback

`func (o *InlineResponse200247) SetLocalAuthFallback(v NetworksNetworkIdWirelessSsidsLocalAuthFallback)`

SetLocalAuthFallback sets LocalAuthFallback field to given value.

### HasLocalAuthFallback

`func (o *InlineResponse200247) HasLocalAuthFallback() bool`

HasLocalAuthFallback returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineResponse200247) GetNamedVlans() NetworksNetworkIdWirelessSsidsNamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineResponse200247) GetNamedVlansOk() (*NetworksNetworkIdWirelessSsidsNamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineResponse200247) SetNamedVlans(v NetworksNetworkIdWirelessSsidsNamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineResponse200247) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.

### GetSpeedBurst

`func (o *InlineResponse200247) GetSpeedBurst() NetworksNetworkIdWirelessSsidsSpeedBurst`

GetSpeedBurst returns the SpeedBurst field if non-nil, zero value otherwise.

### GetSpeedBurstOk

`func (o *InlineResponse200247) GetSpeedBurstOk() (*NetworksNetworkIdWirelessSsidsSpeedBurst, bool)`

GetSpeedBurstOk returns a tuple with the SpeedBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedBurst

`func (o *InlineResponse200247) SetSpeedBurst(v NetworksNetworkIdWirelessSsidsSpeedBurst)`

SetSpeedBurst sets SpeedBurst field to given value.


### GetWifiPersonalNetworkEnabled

`func (o *InlineResponse200247) GetWifiPersonalNetworkEnabled() bool`

GetWifiPersonalNetworkEnabled returns the WifiPersonalNetworkEnabled field if non-nil, zero value otherwise.

### GetWifiPersonalNetworkEnabledOk

`func (o *InlineResponse200247) GetWifiPersonalNetworkEnabledOk() (*bool, bool)`

GetWifiPersonalNetworkEnabledOk returns a tuple with the WifiPersonalNetworkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiPersonalNetworkEnabled

`func (o *InlineResponse200247) SetWifiPersonalNetworkEnabled(v bool)`

SetWifiPersonalNetworkEnabled sets WifiPersonalNetworkEnabled field to given value.

### HasWifiPersonalNetworkEnabled

`func (o *InlineResponse200247) HasWifiPersonalNetworkEnabled() bool`

HasWifiPersonalNetworkEnabled returns a boolean if a field has been set.

### GetSecurity

`func (o *InlineResponse200247) GetSecurity() NetworksNetworkIdWirelessSsidsSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineResponse200247) GetSecurityOk() (*NetworksNetworkIdWirelessSsidsSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineResponse200247) SetSecurity(v NetworksNetworkIdWirelessSsidsSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineResponse200247) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


