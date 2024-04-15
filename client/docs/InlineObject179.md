# InlineObject179

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the SSID | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID (&#39;open&#39;, &#39;open-enhanced&#39;, &#39;psk&#39;, &#39;open-with-radius&#39;, &#39;open-with-nac&#39;, &#39;8021x-meraki&#39;, &#39;8021x-nac&#39;, &#39;8021x-radius&#39;, &#39;8021x-google&#39;, &#39;8021x-localradius&#39;, &#39;ipsk-with-radius&#39;, &#39;ipsk-without-radius&#39; or &#39;ipsk-with-nac&#39;) | [optional] 
**EnterpriseAdminAccess** | Pointer to **string** | Whether or not an SSID is accessible by &#39;enterprise&#39; administrators (&#39;access disabled&#39; or &#39;access enabled&#39;) | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID (&#39;wep&#39; or &#39;wpa&#39;). This param is only valid if the authMode is &#39;psk&#39; | [optional] 
**Psk** | Pointer to **string** | The passkey for the SSID. This param is only valid if the authMode is &#39;psk&#39; | [optional] 
**WpaEncryptionMode** | Pointer to **string** | The types of WPA encryption. (&#39;WPA1 only&#39;, &#39;WPA1 and WPA2&#39;, &#39;WPA2 only&#39;, &#39;WPA3 Transition Mode&#39;, &#39;WPA3 only&#39; or &#39;WPA3 192-bit Security&#39;) | [optional] 
**Dot11w** | Pointer to [**NetworksNetworkIdApplianceSsidsNumberDot11w**](NetworksNetworkIdApplianceSsidsNumberDot11w.md) |  | [optional] 
**Dot11r** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberDot11r**](NetworksNetworkIdWirelessSsidsNumberDot11r.md) |  | [optional] 
**SplashPage** | Pointer to **string** | The type of splash page for the SSID (&#39;None&#39;, &#39;Click-through splash page&#39;, &#39;Billing&#39;, &#39;Password-protected with Meraki RADIUS&#39;, &#39;Password-protected with custom RADIUS&#39;, &#39;Password-protected with Active Directory&#39;, &#39;Password-protected with LDAP&#39;, &#39;SMS authentication&#39;, &#39;Systems Manager Sentry&#39;, &#39;Facebook Wi-Fi&#39;, &#39;Google OAuth&#39;, &#39;Sponsored guest&#39;, &#39;Cisco ISE&#39; or &#39;Google Apps domain&#39;). This attribute is not supported for template children. | [optional] 
**SplashGuestSponsorDomains** | Pointer to **[]string** | Array of valid sponsor email domains for sponsored guest splash type. | [optional] 
**Oauth** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberOauth**](NetworksNetworkIdWirelessSsidsNumberOauth.md) |  | [optional] 
**LocalRadius** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberLocalRadius**](NetworksNetworkIdWirelessSsidsNumberLocalRadius.md) |  | [optional] 
**Ldap** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberLdap**](NetworksNetworkIdWirelessSsidsNumberLdap.md) |  | [optional] 
**ActiveDirectory** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberActiveDirectory**](NetworksNetworkIdWirelessSsidsNumberActiveDirectory.md) |  | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberRadiusServers**](NetworksNetworkIdWirelessSsidsNumberRadiusServers.md) | The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is &#39;open-with-radius&#39;, &#39;8021x-radius&#39; or &#39;ipsk-with-radius&#39; | [optional] 
**RadiusProxyEnabled** | Pointer to **bool** | If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers. | [optional] 
**RadiusTestingEnabled** | Pointer to **bool** | If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity &#39;meraki_8021x_test&#39; to ensure that the RADIUS servers are reachable. | [optional] 
**RadiusCalledStationId** | Pointer to **string** | The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$). | [optional] 
**RadiusAuthenticationNasId** | Pointer to **string** | The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$). | [optional] 
**RadiusServerTimeout** | Pointer to **int32** | The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds). | [optional] 
**RadiusServerAttemptsLimit** | Pointer to **int32** | The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5). | [optional] 
**RadiusFallbackEnabled** | Pointer to **bool** | Whether or not higher priority RADIUS servers should be retried after 60 seconds. | [optional] 
**RadiusCoaEnabled** | Pointer to **bool** | If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server. | [optional] 
**RadiusFailoverPolicy** | Pointer to **string** | This policy determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable (&#39;Deny access&#39; or &#39;Allow access&#39;) | [optional] 
**RadiusLoadBalancingPolicy** | Pointer to **string** | This policy determines which RADIUS server will be contacted first in an authentication attempt and the ordering of any necessary retry attempts (&#39;Strict priority order&#39; or &#39;Round robin&#39;) | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Whether or not RADIUS accounting is enabled. This param is only valid if the authMode is &#39;open-with-radius&#39;, &#39;8021x-radius&#39; or &#39;ipsk-with-radius&#39; | [optional] 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberRadiusAccountingServers**](NetworksNetworkIdWirelessSsidsNumberRadiusAccountingServers.md) | The RADIUS accounting 802.1X servers to be used for authentication. This param is only valid if the authMode is &#39;open-with-radius&#39;, &#39;8021x-radius&#39; or &#39;ipsk-with-radius&#39; and radiusAccountingEnabled is &#39;true&#39; | [optional] 
**RadiusAccountingInterimInterval** | Pointer to **int32** | The interval (in seconds) in which accounting information is updated and sent to the RADIUS accounting server. | [optional] 
**RadiusAttributeForGroupPolicies** | Pointer to **string** | Specify the RADIUS attribute used to look up group policies (&#39;Filter-Id&#39;, &#39;Reply-Message&#39;, &#39;Airespace-ACL-Name&#39; or &#39;Aruba-User-Role&#39;). Access points must receive this attribute in the RADIUS Access-Accept message | [optional] 
**IpAssignmentMode** | Pointer to **string** | The client IP assignment mode (&#39;NAT mode&#39;, &#39;Bridge mode&#39;, &#39;Layer 3 roaming&#39;, &#39;Ethernet over GRE&#39;, &#39;Layer 3 roaming with a concentrator&#39; or &#39;VPN&#39;) | [optional] 
**UseVlanTagging** | Pointer to **bool** | Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is &#39;Bridge mode&#39; or &#39;Layer 3 roaming&#39; | [optional] 
**ConcentratorNetworkId** | Pointer to **string** | The concentrator to use when the ipAssignmentMode is &#39;Layer 3 roaming with a concentrator&#39; or &#39;VPN&#39;. | [optional] 
**SecondaryConcentratorNetworkId** | Pointer to **string** | The secondary concentrator to use when the ipAssignmentMode is &#39;VPN&#39;. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. (&#39;disabled&#39; represents no secondary concentrator.) | [optional] 
**DisassociateClientsOnVpnFailover** | Pointer to **bool** | Disassociate clients when &#39;VPN&#39; concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is &#39;VPN&#39;. | [optional] 
**VlanId** | Pointer to **int32** | The VLAN ID used for VLAN tagging. This param is only valid when the ipAssignmentMode is &#39;Layer 3 roaming with a concentrator&#39; or &#39;VPN&#39; | [optional] 
**DefaultVlanId** | Pointer to **int32** | The default VLAN ID used for &#39;all other APs&#39;. This param is only valid when the ipAssignmentMode is &#39;Bridge mode&#39; or &#39;Layer 3 roaming&#39; | [optional] 
**ApTagsAndVlanIds** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds**](NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds.md) | The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is &#39;Bridge mode&#39; or &#39;Layer 3 roaming&#39; | [optional] 
**WalledGardenEnabled** | Pointer to **bool** | Allow access to a configurable list of IP ranges, which users may access prior to sign-on. | [optional] 
**WalledGardenRanges** | Pointer to **[]string** | Specify your walled garden by entering an array of addresses, ranges using CIDR notation, domain names, and domain wildcards (e.g. &#39;192.168.1.1/24&#39;, &#39;192.168.37.10/32&#39;, &#39;www.yahoo.com&#39;, &#39;*.google.com&#39;]). Meraki&#39;s splash page is automatically included in your walled garden. | [optional] 
**Gre** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberGre**](NetworksNetworkIdWirelessSsidsNumberGre.md) |  | [optional] 
**RadiusOverride** | Pointer to **bool** | If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is &#39;NAT mode&#39;. | [optional] 
**RadiusGuestVlanEnabled** | Pointer to **bool** | Whether or not RADIUS Guest VLAN is enabled. This param is only valid if the authMode is &#39;open-with-radius&#39; and addressing mode is not set to &#39;isolated&#39; or &#39;nat&#39; mode | [optional] 
**RadiusGuestVlanId** | Pointer to **int32** | VLAN ID of the RADIUS Guest VLAN. This param is only valid if the authMode is &#39;open-with-radius&#39; and addressing mode is not set to &#39;isolated&#39; or &#39;nat&#39; mode | [optional] 
**MinBitrate** | Pointer to **float32** | The minimum bitrate in Mbps of this SSID in the default indoor RF profile. (&#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;) | [optional] 
**BandSelection** | Pointer to **string** | The client-serving radio frequencies of this SSID in the default indoor RF profile. (&#39;Dual band operation&#39;, &#39;5 GHz band only&#39; or &#39;Dual band operation with Band Steering&#39;) | [optional] 
**PerClientBandwidthLimitUp** | Pointer to **int32** | The upload bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**PerClientBandwidthLimitDown** | Pointer to **int32** | The download bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**PerSsidBandwidthLimitUp** | Pointer to **int32** | The total upload bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**PerSsidBandwidthLimitDown** | Pointer to **int32** | The total download bandwidth limit in Kbps. (0 represents no limit.) | [optional] 
**LanIsolationEnabled** | Pointer to **bool** | Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is &#39;Bridge mode&#39;. | [optional] 
**Visible** | Pointer to **bool** | Boolean indicating whether APs should advertise or hide this SSID. APs will only broadcast this SSID if set to true | [optional] 
**AvailableOnAllAps** | Pointer to **bool** | Boolean indicating whether all APs should broadcast the SSID or if it should be restricted to APs matching any availability tags. Can only be false if the SSID has availability tags. | [optional] 
**AvailabilityTags** | Pointer to **[]string** | Accepts a list of tags for this SSID. If availableOnAllAps is false, then the SSID will only be broadcast by APs with tags matching any of the tags in this list. | [optional] 
**MandatoryDhcpEnabled** | Pointer to **bool** | If true, Mandatory DHCP will enforce that clients connecting to this SSID must use the IP address assigned by the DHCP server. Clients who use a static IP address won&#39;t be able to associate. | [optional] 
**AdultContentFilteringEnabled** | Pointer to **bool** | Boolean indicating whether or not adult content will be blocked | [optional] 
**DnsRewrite** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberDnsRewrite**](NetworksNetworkIdWirelessSsidsNumberDnsRewrite.md) |  | [optional] 
**SpeedBurst** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSpeedBurst**](NetworksNetworkIdWirelessSsidsNumberSpeedBurst.md) |  | [optional] 
**NamedVlans** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberNamedVlans**](NetworksNetworkIdWirelessSsidsNumberNamedVlans.md) |  | [optional] 

## Methods

### NewInlineObject179

`func NewInlineObject179() *InlineObject179`

NewInlineObject179 instantiates a new InlineObject179 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject179WithDefaults

`func NewInlineObject179WithDefaults() *InlineObject179`

NewInlineObject179WithDefaults instantiates a new InlineObject179 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject179) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject179) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject179) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject179) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject179) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject179) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject179) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject179) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthMode

`func (o *InlineObject179) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InlineObject179) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InlineObject179) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *InlineObject179) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetEnterpriseAdminAccess

`func (o *InlineObject179) GetEnterpriseAdminAccess() string`

GetEnterpriseAdminAccess returns the EnterpriseAdminAccess field if non-nil, zero value otherwise.

### GetEnterpriseAdminAccessOk

`func (o *InlineObject179) GetEnterpriseAdminAccessOk() (*string, bool)`

GetEnterpriseAdminAccessOk returns a tuple with the EnterpriseAdminAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseAdminAccess

`func (o *InlineObject179) SetEnterpriseAdminAccess(v string)`

SetEnterpriseAdminAccess sets EnterpriseAdminAccess field to given value.

### HasEnterpriseAdminAccess

`func (o *InlineObject179) HasEnterpriseAdminAccess() bool`

HasEnterpriseAdminAccess returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *InlineObject179) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *InlineObject179) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *InlineObject179) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *InlineObject179) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetPsk

`func (o *InlineObject179) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *InlineObject179) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *InlineObject179) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *InlineObject179) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *InlineObject179) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *InlineObject179) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *InlineObject179) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *InlineObject179) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetDot11w

`func (o *InlineObject179) GetDot11w() NetworksNetworkIdApplianceSsidsNumberDot11w`

GetDot11w returns the Dot11w field if non-nil, zero value otherwise.

### GetDot11wOk

`func (o *InlineObject179) GetDot11wOk() (*NetworksNetworkIdApplianceSsidsNumberDot11w, bool)`

GetDot11wOk returns a tuple with the Dot11w field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11w

`func (o *InlineObject179) SetDot11w(v NetworksNetworkIdApplianceSsidsNumberDot11w)`

SetDot11w sets Dot11w field to given value.

### HasDot11w

`func (o *InlineObject179) HasDot11w() bool`

HasDot11w returns a boolean if a field has been set.

### GetDot11r

`func (o *InlineObject179) GetDot11r() NetworksNetworkIdWirelessSsidsNumberDot11r`

GetDot11r returns the Dot11r field if non-nil, zero value otherwise.

### GetDot11rOk

`func (o *InlineObject179) GetDot11rOk() (*NetworksNetworkIdWirelessSsidsNumberDot11r, bool)`

GetDot11rOk returns a tuple with the Dot11r field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11r

`func (o *InlineObject179) SetDot11r(v NetworksNetworkIdWirelessSsidsNumberDot11r)`

SetDot11r sets Dot11r field to given value.

### HasDot11r

`func (o *InlineObject179) HasDot11r() bool`

HasDot11r returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineObject179) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineObject179) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineObject179) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *InlineObject179) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetSplashGuestSponsorDomains

`func (o *InlineObject179) GetSplashGuestSponsorDomains() []string`

GetSplashGuestSponsorDomains returns the SplashGuestSponsorDomains field if non-nil, zero value otherwise.

### GetSplashGuestSponsorDomainsOk

`func (o *InlineObject179) GetSplashGuestSponsorDomainsOk() (*[]string, bool)`

GetSplashGuestSponsorDomainsOk returns a tuple with the SplashGuestSponsorDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashGuestSponsorDomains

`func (o *InlineObject179) SetSplashGuestSponsorDomains(v []string)`

SetSplashGuestSponsorDomains sets SplashGuestSponsorDomains field to given value.

### HasSplashGuestSponsorDomains

`func (o *InlineObject179) HasSplashGuestSponsorDomains() bool`

HasSplashGuestSponsorDomains returns a boolean if a field has been set.

### GetOauth

`func (o *InlineObject179) GetOauth() NetworksNetworkIdWirelessSsidsNumberOauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *InlineObject179) GetOauthOk() (*NetworksNetworkIdWirelessSsidsNumberOauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *InlineObject179) SetOauth(v NetworksNetworkIdWirelessSsidsNumberOauth)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *InlineObject179) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetLocalRadius

`func (o *InlineObject179) GetLocalRadius() NetworksNetworkIdWirelessSsidsNumberLocalRadius`

GetLocalRadius returns the LocalRadius field if non-nil, zero value otherwise.

### GetLocalRadiusOk

`func (o *InlineObject179) GetLocalRadiusOk() (*NetworksNetworkIdWirelessSsidsNumberLocalRadius, bool)`

GetLocalRadiusOk returns a tuple with the LocalRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRadius

`func (o *InlineObject179) SetLocalRadius(v NetworksNetworkIdWirelessSsidsNumberLocalRadius)`

SetLocalRadius sets LocalRadius field to given value.

### HasLocalRadius

`func (o *InlineObject179) HasLocalRadius() bool`

HasLocalRadius returns a boolean if a field has been set.

### GetLdap

`func (o *InlineObject179) GetLdap() NetworksNetworkIdWirelessSsidsNumberLdap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *InlineObject179) GetLdapOk() (*NetworksNetworkIdWirelessSsidsNumberLdap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *InlineObject179) SetLdap(v NetworksNetworkIdWirelessSsidsNumberLdap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *InlineObject179) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetActiveDirectory

`func (o *InlineObject179) GetActiveDirectory() NetworksNetworkIdWirelessSsidsNumberActiveDirectory`

GetActiveDirectory returns the ActiveDirectory field if non-nil, zero value otherwise.

### GetActiveDirectoryOk

`func (o *InlineObject179) GetActiveDirectoryOk() (*NetworksNetworkIdWirelessSsidsNumberActiveDirectory, bool)`

GetActiveDirectoryOk returns a tuple with the ActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectory

`func (o *InlineObject179) SetActiveDirectory(v NetworksNetworkIdWirelessSsidsNumberActiveDirectory)`

SetActiveDirectory sets ActiveDirectory field to given value.

### HasActiveDirectory

`func (o *InlineObject179) HasActiveDirectory() bool`

HasActiveDirectory returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineObject179) GetRadiusServers() []NetworksNetworkIdWirelessSsidsNumberRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineObject179) GetRadiusServersOk() (*[]NetworksNetworkIdWirelessSsidsNumberRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineObject179) SetRadiusServers(v []NetworksNetworkIdWirelessSsidsNumberRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineObject179) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadiusProxyEnabled

`func (o *InlineObject179) GetRadiusProxyEnabled() bool`

GetRadiusProxyEnabled returns the RadiusProxyEnabled field if non-nil, zero value otherwise.

### GetRadiusProxyEnabledOk

`func (o *InlineObject179) GetRadiusProxyEnabledOk() (*bool, bool)`

GetRadiusProxyEnabledOk returns a tuple with the RadiusProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProxyEnabled

`func (o *InlineObject179) SetRadiusProxyEnabled(v bool)`

SetRadiusProxyEnabled sets RadiusProxyEnabled field to given value.

### HasRadiusProxyEnabled

`func (o *InlineObject179) HasRadiusProxyEnabled() bool`

HasRadiusProxyEnabled returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineObject179) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineObject179) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineObject179) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.

### HasRadiusTestingEnabled

`func (o *InlineObject179) HasRadiusTestingEnabled() bool`

HasRadiusTestingEnabled returns a boolean if a field has been set.

### GetRadiusCalledStationId

`func (o *InlineObject179) GetRadiusCalledStationId() string`

GetRadiusCalledStationId returns the RadiusCalledStationId field if non-nil, zero value otherwise.

### GetRadiusCalledStationIdOk

`func (o *InlineObject179) GetRadiusCalledStationIdOk() (*string, bool)`

GetRadiusCalledStationIdOk returns a tuple with the RadiusCalledStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCalledStationId

`func (o *InlineObject179) SetRadiusCalledStationId(v string)`

SetRadiusCalledStationId sets RadiusCalledStationId field to given value.

### HasRadiusCalledStationId

`func (o *InlineObject179) HasRadiusCalledStationId() bool`

HasRadiusCalledStationId returns a boolean if a field has been set.

### GetRadiusAuthenticationNasId

`func (o *InlineObject179) GetRadiusAuthenticationNasId() string`

GetRadiusAuthenticationNasId returns the RadiusAuthenticationNasId field if non-nil, zero value otherwise.

### GetRadiusAuthenticationNasIdOk

`func (o *InlineObject179) GetRadiusAuthenticationNasIdOk() (*string, bool)`

GetRadiusAuthenticationNasIdOk returns a tuple with the RadiusAuthenticationNasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAuthenticationNasId

`func (o *InlineObject179) SetRadiusAuthenticationNasId(v string)`

SetRadiusAuthenticationNasId sets RadiusAuthenticationNasId field to given value.

### HasRadiusAuthenticationNasId

`func (o *InlineObject179) HasRadiusAuthenticationNasId() bool`

HasRadiusAuthenticationNasId returns a boolean if a field has been set.

### GetRadiusServerTimeout

`func (o *InlineObject179) GetRadiusServerTimeout() int32`

GetRadiusServerTimeout returns the RadiusServerTimeout field if non-nil, zero value otherwise.

### GetRadiusServerTimeoutOk

`func (o *InlineObject179) GetRadiusServerTimeoutOk() (*int32, bool)`

GetRadiusServerTimeoutOk returns a tuple with the RadiusServerTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServerTimeout

`func (o *InlineObject179) SetRadiusServerTimeout(v int32)`

SetRadiusServerTimeout sets RadiusServerTimeout field to given value.

### HasRadiusServerTimeout

`func (o *InlineObject179) HasRadiusServerTimeout() bool`

HasRadiusServerTimeout returns a boolean if a field has been set.

### GetRadiusServerAttemptsLimit

`func (o *InlineObject179) GetRadiusServerAttemptsLimit() int32`

GetRadiusServerAttemptsLimit returns the RadiusServerAttemptsLimit field if non-nil, zero value otherwise.

### GetRadiusServerAttemptsLimitOk

`func (o *InlineObject179) GetRadiusServerAttemptsLimitOk() (*int32, bool)`

GetRadiusServerAttemptsLimitOk returns a tuple with the RadiusServerAttemptsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServerAttemptsLimit

`func (o *InlineObject179) SetRadiusServerAttemptsLimit(v int32)`

SetRadiusServerAttemptsLimit sets RadiusServerAttemptsLimit field to given value.

### HasRadiusServerAttemptsLimit

`func (o *InlineObject179) HasRadiusServerAttemptsLimit() bool`

HasRadiusServerAttemptsLimit returns a boolean if a field has been set.

### GetRadiusFallbackEnabled

`func (o *InlineObject179) GetRadiusFallbackEnabled() bool`

GetRadiusFallbackEnabled returns the RadiusFallbackEnabled field if non-nil, zero value otherwise.

### GetRadiusFallbackEnabledOk

`func (o *InlineObject179) GetRadiusFallbackEnabledOk() (*bool, bool)`

GetRadiusFallbackEnabledOk returns a tuple with the RadiusFallbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusFallbackEnabled

`func (o *InlineObject179) SetRadiusFallbackEnabled(v bool)`

SetRadiusFallbackEnabled sets RadiusFallbackEnabled field to given value.

### HasRadiusFallbackEnabled

`func (o *InlineObject179) HasRadiusFallbackEnabled() bool`

HasRadiusFallbackEnabled returns a boolean if a field has been set.

### GetRadiusCoaEnabled

`func (o *InlineObject179) GetRadiusCoaEnabled() bool`

GetRadiusCoaEnabled returns the RadiusCoaEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaEnabledOk

`func (o *InlineObject179) GetRadiusCoaEnabledOk() (*bool, bool)`

GetRadiusCoaEnabledOk returns a tuple with the RadiusCoaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaEnabled

`func (o *InlineObject179) SetRadiusCoaEnabled(v bool)`

SetRadiusCoaEnabled sets RadiusCoaEnabled field to given value.

### HasRadiusCoaEnabled

`func (o *InlineObject179) HasRadiusCoaEnabled() bool`

HasRadiusCoaEnabled returns a boolean if a field has been set.

### GetRadiusFailoverPolicy

`func (o *InlineObject179) GetRadiusFailoverPolicy() string`

GetRadiusFailoverPolicy returns the RadiusFailoverPolicy field if non-nil, zero value otherwise.

### GetRadiusFailoverPolicyOk

`func (o *InlineObject179) GetRadiusFailoverPolicyOk() (*string, bool)`

GetRadiusFailoverPolicyOk returns a tuple with the RadiusFailoverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusFailoverPolicy

`func (o *InlineObject179) SetRadiusFailoverPolicy(v string)`

SetRadiusFailoverPolicy sets RadiusFailoverPolicy field to given value.

### HasRadiusFailoverPolicy

`func (o *InlineObject179) HasRadiusFailoverPolicy() bool`

HasRadiusFailoverPolicy returns a boolean if a field has been set.

### GetRadiusLoadBalancingPolicy

`func (o *InlineObject179) GetRadiusLoadBalancingPolicy() string`

GetRadiusLoadBalancingPolicy returns the RadiusLoadBalancingPolicy field if non-nil, zero value otherwise.

### GetRadiusLoadBalancingPolicyOk

`func (o *InlineObject179) GetRadiusLoadBalancingPolicyOk() (*string, bool)`

GetRadiusLoadBalancingPolicyOk returns a tuple with the RadiusLoadBalancingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusLoadBalancingPolicy

`func (o *InlineObject179) SetRadiusLoadBalancingPolicy(v string)`

SetRadiusLoadBalancingPolicy sets RadiusLoadBalancingPolicy field to given value.

### HasRadiusLoadBalancingPolicy

`func (o *InlineObject179) HasRadiusLoadBalancingPolicy() bool`

HasRadiusLoadBalancingPolicy returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *InlineObject179) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineObject179) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineObject179) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *InlineObject179) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *InlineObject179) GetRadiusAccountingServers() []NetworksNetworkIdWirelessSsidsNumberRadiusAccountingServers`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineObject179) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdWirelessSsidsNumberRadiusAccountingServers, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineObject179) SetRadiusAccountingServers(v []NetworksNetworkIdWirelessSsidsNumberRadiusAccountingServers)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineObject179) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusAccountingInterimInterval

`func (o *InlineObject179) GetRadiusAccountingInterimInterval() int32`

GetRadiusAccountingInterimInterval returns the RadiusAccountingInterimInterval field if non-nil, zero value otherwise.

### GetRadiusAccountingInterimIntervalOk

`func (o *InlineObject179) GetRadiusAccountingInterimIntervalOk() (*int32, bool)`

GetRadiusAccountingInterimIntervalOk returns a tuple with the RadiusAccountingInterimInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingInterimInterval

`func (o *InlineObject179) SetRadiusAccountingInterimInterval(v int32)`

SetRadiusAccountingInterimInterval sets RadiusAccountingInterimInterval field to given value.

### HasRadiusAccountingInterimInterval

`func (o *InlineObject179) HasRadiusAccountingInterimInterval() bool`

HasRadiusAccountingInterimInterval returns a boolean if a field has been set.

### GetRadiusAttributeForGroupPolicies

`func (o *InlineObject179) GetRadiusAttributeForGroupPolicies() string`

GetRadiusAttributeForGroupPolicies returns the RadiusAttributeForGroupPolicies field if non-nil, zero value otherwise.

### GetRadiusAttributeForGroupPoliciesOk

`func (o *InlineObject179) GetRadiusAttributeForGroupPoliciesOk() (*string, bool)`

GetRadiusAttributeForGroupPoliciesOk returns a tuple with the RadiusAttributeForGroupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttributeForGroupPolicies

`func (o *InlineObject179) SetRadiusAttributeForGroupPolicies(v string)`

SetRadiusAttributeForGroupPolicies sets RadiusAttributeForGroupPolicies field to given value.

### HasRadiusAttributeForGroupPolicies

`func (o *InlineObject179) HasRadiusAttributeForGroupPolicies() bool`

HasRadiusAttributeForGroupPolicies returns a boolean if a field has been set.

### GetIpAssignmentMode

`func (o *InlineObject179) GetIpAssignmentMode() string`

GetIpAssignmentMode returns the IpAssignmentMode field if non-nil, zero value otherwise.

### GetIpAssignmentModeOk

`func (o *InlineObject179) GetIpAssignmentModeOk() (*string, bool)`

GetIpAssignmentModeOk returns a tuple with the IpAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignmentMode

`func (o *InlineObject179) SetIpAssignmentMode(v string)`

SetIpAssignmentMode sets IpAssignmentMode field to given value.

### HasIpAssignmentMode

`func (o *InlineObject179) HasIpAssignmentMode() bool`

HasIpAssignmentMode returns a boolean if a field has been set.

### GetUseVlanTagging

`func (o *InlineObject179) GetUseVlanTagging() bool`

GetUseVlanTagging returns the UseVlanTagging field if non-nil, zero value otherwise.

### GetUseVlanTaggingOk

`func (o *InlineObject179) GetUseVlanTaggingOk() (*bool, bool)`

GetUseVlanTaggingOk returns a tuple with the UseVlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVlanTagging

`func (o *InlineObject179) SetUseVlanTagging(v bool)`

SetUseVlanTagging sets UseVlanTagging field to given value.

### HasUseVlanTagging

`func (o *InlineObject179) HasUseVlanTagging() bool`

HasUseVlanTagging returns a boolean if a field has been set.

### GetConcentratorNetworkId

`func (o *InlineObject179) GetConcentratorNetworkId() string`

GetConcentratorNetworkId returns the ConcentratorNetworkId field if non-nil, zero value otherwise.

### GetConcentratorNetworkIdOk

`func (o *InlineObject179) GetConcentratorNetworkIdOk() (*string, bool)`

GetConcentratorNetworkIdOk returns a tuple with the ConcentratorNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentratorNetworkId

`func (o *InlineObject179) SetConcentratorNetworkId(v string)`

SetConcentratorNetworkId sets ConcentratorNetworkId field to given value.

### HasConcentratorNetworkId

`func (o *InlineObject179) HasConcentratorNetworkId() bool`

HasConcentratorNetworkId returns a boolean if a field has been set.

### GetSecondaryConcentratorNetworkId

`func (o *InlineObject179) GetSecondaryConcentratorNetworkId() string`

GetSecondaryConcentratorNetworkId returns the SecondaryConcentratorNetworkId field if non-nil, zero value otherwise.

### GetSecondaryConcentratorNetworkIdOk

`func (o *InlineObject179) GetSecondaryConcentratorNetworkIdOk() (*string, bool)`

GetSecondaryConcentratorNetworkIdOk returns a tuple with the SecondaryConcentratorNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryConcentratorNetworkId

`func (o *InlineObject179) SetSecondaryConcentratorNetworkId(v string)`

SetSecondaryConcentratorNetworkId sets SecondaryConcentratorNetworkId field to given value.

### HasSecondaryConcentratorNetworkId

`func (o *InlineObject179) HasSecondaryConcentratorNetworkId() bool`

HasSecondaryConcentratorNetworkId returns a boolean if a field has been set.

### GetDisassociateClientsOnVpnFailover

`func (o *InlineObject179) GetDisassociateClientsOnVpnFailover() bool`

GetDisassociateClientsOnVpnFailover returns the DisassociateClientsOnVpnFailover field if non-nil, zero value otherwise.

### GetDisassociateClientsOnVpnFailoverOk

`func (o *InlineObject179) GetDisassociateClientsOnVpnFailoverOk() (*bool, bool)`

GetDisassociateClientsOnVpnFailoverOk returns a tuple with the DisassociateClientsOnVpnFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisassociateClientsOnVpnFailover

`func (o *InlineObject179) SetDisassociateClientsOnVpnFailover(v bool)`

SetDisassociateClientsOnVpnFailover sets DisassociateClientsOnVpnFailover field to given value.

### HasDisassociateClientsOnVpnFailover

`func (o *InlineObject179) HasDisassociateClientsOnVpnFailover() bool`

HasDisassociateClientsOnVpnFailover returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineObject179) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineObject179) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineObject179) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineObject179) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *InlineObject179) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *InlineObject179) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *InlineObject179) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *InlineObject179) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetApTagsAndVlanIds

`func (o *InlineObject179) GetApTagsAndVlanIds() []NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds`

GetApTagsAndVlanIds returns the ApTagsAndVlanIds field if non-nil, zero value otherwise.

### GetApTagsAndVlanIdsOk

`func (o *InlineObject179) GetApTagsAndVlanIdsOk() (*[]NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds, bool)`

GetApTagsAndVlanIdsOk returns a tuple with the ApTagsAndVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApTagsAndVlanIds

`func (o *InlineObject179) SetApTagsAndVlanIds(v []NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds)`

SetApTagsAndVlanIds sets ApTagsAndVlanIds field to given value.

### HasApTagsAndVlanIds

`func (o *InlineObject179) HasApTagsAndVlanIds() bool`

HasApTagsAndVlanIds returns a boolean if a field has been set.

### GetWalledGardenEnabled

`func (o *InlineObject179) GetWalledGardenEnabled() bool`

GetWalledGardenEnabled returns the WalledGardenEnabled field if non-nil, zero value otherwise.

### GetWalledGardenEnabledOk

`func (o *InlineObject179) GetWalledGardenEnabledOk() (*bool, bool)`

GetWalledGardenEnabledOk returns a tuple with the WalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenEnabled

`func (o *InlineObject179) SetWalledGardenEnabled(v bool)`

SetWalledGardenEnabled sets WalledGardenEnabled field to given value.

### HasWalledGardenEnabled

`func (o *InlineObject179) HasWalledGardenEnabled() bool`

HasWalledGardenEnabled returns a boolean if a field has been set.

### GetWalledGardenRanges

`func (o *InlineObject179) GetWalledGardenRanges() []string`

GetWalledGardenRanges returns the WalledGardenRanges field if non-nil, zero value otherwise.

### GetWalledGardenRangesOk

`func (o *InlineObject179) GetWalledGardenRangesOk() (*[]string, bool)`

GetWalledGardenRangesOk returns a tuple with the WalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalledGardenRanges

`func (o *InlineObject179) SetWalledGardenRanges(v []string)`

SetWalledGardenRanges sets WalledGardenRanges field to given value.

### HasWalledGardenRanges

`func (o *InlineObject179) HasWalledGardenRanges() bool`

HasWalledGardenRanges returns a boolean if a field has been set.

### GetGre

`func (o *InlineObject179) GetGre() NetworksNetworkIdWirelessSsidsNumberGre`

GetGre returns the Gre field if non-nil, zero value otherwise.

### GetGreOk

`func (o *InlineObject179) GetGreOk() (*NetworksNetworkIdWirelessSsidsNumberGre, bool)`

GetGreOk returns a tuple with the Gre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGre

`func (o *InlineObject179) SetGre(v NetworksNetworkIdWirelessSsidsNumberGre)`

SetGre sets Gre field to given value.

### HasGre

`func (o *InlineObject179) HasGre() bool`

HasGre returns a boolean if a field has been set.

### GetRadiusOverride

`func (o *InlineObject179) GetRadiusOverride() bool`

GetRadiusOverride returns the RadiusOverride field if non-nil, zero value otherwise.

### GetRadiusOverrideOk

`func (o *InlineObject179) GetRadiusOverrideOk() (*bool, bool)`

GetRadiusOverrideOk returns a tuple with the RadiusOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusOverride

`func (o *InlineObject179) SetRadiusOverride(v bool)`

SetRadiusOverride sets RadiusOverride field to given value.

### HasRadiusOverride

`func (o *InlineObject179) HasRadiusOverride() bool`

HasRadiusOverride returns a boolean if a field has been set.

### GetRadiusGuestVlanEnabled

`func (o *InlineObject179) GetRadiusGuestVlanEnabled() bool`

GetRadiusGuestVlanEnabled returns the RadiusGuestVlanEnabled field if non-nil, zero value otherwise.

### GetRadiusGuestVlanEnabledOk

`func (o *InlineObject179) GetRadiusGuestVlanEnabledOk() (*bool, bool)`

GetRadiusGuestVlanEnabledOk returns a tuple with the RadiusGuestVlanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGuestVlanEnabled

`func (o *InlineObject179) SetRadiusGuestVlanEnabled(v bool)`

SetRadiusGuestVlanEnabled sets RadiusGuestVlanEnabled field to given value.

### HasRadiusGuestVlanEnabled

`func (o *InlineObject179) HasRadiusGuestVlanEnabled() bool`

HasRadiusGuestVlanEnabled returns a boolean if a field has been set.

### GetRadiusGuestVlanId

`func (o *InlineObject179) GetRadiusGuestVlanId() int32`

GetRadiusGuestVlanId returns the RadiusGuestVlanId field if non-nil, zero value otherwise.

### GetRadiusGuestVlanIdOk

`func (o *InlineObject179) GetRadiusGuestVlanIdOk() (*int32, bool)`

GetRadiusGuestVlanIdOk returns a tuple with the RadiusGuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGuestVlanId

`func (o *InlineObject179) SetRadiusGuestVlanId(v int32)`

SetRadiusGuestVlanId sets RadiusGuestVlanId field to given value.

### HasRadiusGuestVlanId

`func (o *InlineObject179) HasRadiusGuestVlanId() bool`

HasRadiusGuestVlanId returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineObject179) GetMinBitrate() float32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineObject179) GetMinBitrateOk() (*float32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineObject179) SetMinBitrate(v float32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineObject179) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetBandSelection

`func (o *InlineObject179) GetBandSelection() string`

GetBandSelection returns the BandSelection field if non-nil, zero value otherwise.

### GetBandSelectionOk

`func (o *InlineObject179) GetBandSelectionOk() (*string, bool)`

GetBandSelectionOk returns a tuple with the BandSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelection

`func (o *InlineObject179) SetBandSelection(v string)`

SetBandSelection sets BandSelection field to given value.

### HasBandSelection

`func (o *InlineObject179) HasBandSelection() bool`

HasBandSelection returns a boolean if a field has been set.

### GetPerClientBandwidthLimitUp

`func (o *InlineObject179) GetPerClientBandwidthLimitUp() int32`

GetPerClientBandwidthLimitUp returns the PerClientBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitUpOk

`func (o *InlineObject179) GetPerClientBandwidthLimitUpOk() (*int32, bool)`

GetPerClientBandwidthLimitUpOk returns a tuple with the PerClientBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitUp

`func (o *InlineObject179) SetPerClientBandwidthLimitUp(v int32)`

SetPerClientBandwidthLimitUp sets PerClientBandwidthLimitUp field to given value.

### HasPerClientBandwidthLimitUp

`func (o *InlineObject179) HasPerClientBandwidthLimitUp() bool`

HasPerClientBandwidthLimitUp returns a boolean if a field has been set.

### GetPerClientBandwidthLimitDown

`func (o *InlineObject179) GetPerClientBandwidthLimitDown() int32`

GetPerClientBandwidthLimitDown returns the PerClientBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitDownOk

`func (o *InlineObject179) GetPerClientBandwidthLimitDownOk() (*int32, bool)`

GetPerClientBandwidthLimitDownOk returns a tuple with the PerClientBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimitDown

`func (o *InlineObject179) SetPerClientBandwidthLimitDown(v int32)`

SetPerClientBandwidthLimitDown sets PerClientBandwidthLimitDown field to given value.

### HasPerClientBandwidthLimitDown

`func (o *InlineObject179) HasPerClientBandwidthLimitDown() bool`

HasPerClientBandwidthLimitDown returns a boolean if a field has been set.

### GetPerSsidBandwidthLimitUp

`func (o *InlineObject179) GetPerSsidBandwidthLimitUp() int32`

GetPerSsidBandwidthLimitUp returns the PerSsidBandwidthLimitUp field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitUpOk

`func (o *InlineObject179) GetPerSsidBandwidthLimitUpOk() (*int32, bool)`

GetPerSsidBandwidthLimitUpOk returns a tuple with the PerSsidBandwidthLimitUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitUp

`func (o *InlineObject179) SetPerSsidBandwidthLimitUp(v int32)`

SetPerSsidBandwidthLimitUp sets PerSsidBandwidthLimitUp field to given value.

### HasPerSsidBandwidthLimitUp

`func (o *InlineObject179) HasPerSsidBandwidthLimitUp() bool`

HasPerSsidBandwidthLimitUp returns a boolean if a field has been set.

### GetPerSsidBandwidthLimitDown

`func (o *InlineObject179) GetPerSsidBandwidthLimitDown() int32`

GetPerSsidBandwidthLimitDown returns the PerSsidBandwidthLimitDown field if non-nil, zero value otherwise.

### GetPerSsidBandwidthLimitDownOk

`func (o *InlineObject179) GetPerSsidBandwidthLimitDownOk() (*int32, bool)`

GetPerSsidBandwidthLimitDownOk returns a tuple with the PerSsidBandwidthLimitDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidBandwidthLimitDown

`func (o *InlineObject179) SetPerSsidBandwidthLimitDown(v int32)`

SetPerSsidBandwidthLimitDown sets PerSsidBandwidthLimitDown field to given value.

### HasPerSsidBandwidthLimitDown

`func (o *InlineObject179) HasPerSsidBandwidthLimitDown() bool`

HasPerSsidBandwidthLimitDown returns a boolean if a field has been set.

### GetLanIsolationEnabled

`func (o *InlineObject179) GetLanIsolationEnabled() bool`

GetLanIsolationEnabled returns the LanIsolationEnabled field if non-nil, zero value otherwise.

### GetLanIsolationEnabledOk

`func (o *InlineObject179) GetLanIsolationEnabledOk() (*bool, bool)`

GetLanIsolationEnabledOk returns a tuple with the LanIsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIsolationEnabled

`func (o *InlineObject179) SetLanIsolationEnabled(v bool)`

SetLanIsolationEnabled sets LanIsolationEnabled field to given value.

### HasLanIsolationEnabled

`func (o *InlineObject179) HasLanIsolationEnabled() bool`

HasLanIsolationEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *InlineObject179) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineObject179) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineObject179) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InlineObject179) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAvailableOnAllAps

`func (o *InlineObject179) GetAvailableOnAllAps() bool`

GetAvailableOnAllAps returns the AvailableOnAllAps field if non-nil, zero value otherwise.

### GetAvailableOnAllApsOk

`func (o *InlineObject179) GetAvailableOnAllApsOk() (*bool, bool)`

GetAvailableOnAllApsOk returns a tuple with the AvailableOnAllAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOnAllAps

`func (o *InlineObject179) SetAvailableOnAllAps(v bool)`

SetAvailableOnAllAps sets AvailableOnAllAps field to given value.

### HasAvailableOnAllAps

`func (o *InlineObject179) HasAvailableOnAllAps() bool`

HasAvailableOnAllAps returns a boolean if a field has been set.

### GetAvailabilityTags

`func (o *InlineObject179) GetAvailabilityTags() []string`

GetAvailabilityTags returns the AvailabilityTags field if non-nil, zero value otherwise.

### GetAvailabilityTagsOk

`func (o *InlineObject179) GetAvailabilityTagsOk() (*[]string, bool)`

GetAvailabilityTagsOk returns a tuple with the AvailabilityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTags

`func (o *InlineObject179) SetAvailabilityTags(v []string)`

SetAvailabilityTags sets AvailabilityTags field to given value.

### HasAvailabilityTags

`func (o *InlineObject179) HasAvailabilityTags() bool`

HasAvailabilityTags returns a boolean if a field has been set.

### GetMandatoryDhcpEnabled

`func (o *InlineObject179) GetMandatoryDhcpEnabled() bool`

GetMandatoryDhcpEnabled returns the MandatoryDhcpEnabled field if non-nil, zero value otherwise.

### GetMandatoryDhcpEnabledOk

`func (o *InlineObject179) GetMandatoryDhcpEnabledOk() (*bool, bool)`

GetMandatoryDhcpEnabledOk returns a tuple with the MandatoryDhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryDhcpEnabled

`func (o *InlineObject179) SetMandatoryDhcpEnabled(v bool)`

SetMandatoryDhcpEnabled sets MandatoryDhcpEnabled field to given value.

### HasMandatoryDhcpEnabled

`func (o *InlineObject179) HasMandatoryDhcpEnabled() bool`

HasMandatoryDhcpEnabled returns a boolean if a field has been set.

### GetAdultContentFilteringEnabled

`func (o *InlineObject179) GetAdultContentFilteringEnabled() bool`

GetAdultContentFilteringEnabled returns the AdultContentFilteringEnabled field if non-nil, zero value otherwise.

### GetAdultContentFilteringEnabledOk

`func (o *InlineObject179) GetAdultContentFilteringEnabledOk() (*bool, bool)`

GetAdultContentFilteringEnabledOk returns a tuple with the AdultContentFilteringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContentFilteringEnabled

`func (o *InlineObject179) SetAdultContentFilteringEnabled(v bool)`

SetAdultContentFilteringEnabled sets AdultContentFilteringEnabled field to given value.

### HasAdultContentFilteringEnabled

`func (o *InlineObject179) HasAdultContentFilteringEnabled() bool`

HasAdultContentFilteringEnabled returns a boolean if a field has been set.

### GetDnsRewrite

`func (o *InlineObject179) GetDnsRewrite() NetworksNetworkIdWirelessSsidsNumberDnsRewrite`

GetDnsRewrite returns the DnsRewrite field if non-nil, zero value otherwise.

### GetDnsRewriteOk

`func (o *InlineObject179) GetDnsRewriteOk() (*NetworksNetworkIdWirelessSsidsNumberDnsRewrite, bool)`

GetDnsRewriteOk returns a tuple with the DnsRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRewrite

`func (o *InlineObject179) SetDnsRewrite(v NetworksNetworkIdWirelessSsidsNumberDnsRewrite)`

SetDnsRewrite sets DnsRewrite field to given value.

### HasDnsRewrite

`func (o *InlineObject179) HasDnsRewrite() bool`

HasDnsRewrite returns a boolean if a field has been set.

### GetSpeedBurst

`func (o *InlineObject179) GetSpeedBurst() NetworksNetworkIdWirelessSsidsNumberSpeedBurst`

GetSpeedBurst returns the SpeedBurst field if non-nil, zero value otherwise.

### GetSpeedBurstOk

`func (o *InlineObject179) GetSpeedBurstOk() (*NetworksNetworkIdWirelessSsidsNumberSpeedBurst, bool)`

GetSpeedBurstOk returns a tuple with the SpeedBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedBurst

`func (o *InlineObject179) SetSpeedBurst(v NetworksNetworkIdWirelessSsidsNumberSpeedBurst)`

SetSpeedBurst sets SpeedBurst field to given value.

### HasSpeedBurst

`func (o *InlineObject179) HasSpeedBurst() bool`

HasSpeedBurst returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineObject179) GetNamedVlans() NetworksNetworkIdWirelessSsidsNumberNamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineObject179) GetNamedVlansOk() (*NetworksNetworkIdWirelessSsidsNumberNamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineObject179) SetNamedVlans(v NetworksNetworkIdWirelessSsidsNumberNamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineObject179) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


