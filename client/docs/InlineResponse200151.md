# InlineResponse200151

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the access policy | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesRadiusServers**](NetworksNetworkIdSwitchAccessPoliciesRadiusServers.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | [optional] 
**Radius** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius**](NetworksNetworkIdSwitchAccessPoliciesRadius.md) |  | [optional] 
**EnforceRadiusMonitoring** | Pointer to **bool** | This is a readonly flag, indicating whether the access policy was under has_guest_port_bouncing NFO | [optional] 
**GuestPortBouncing** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusTestingEnabled** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusCoaSupportEnabled** | Pointer to **bool** | Change of authentication for RADIUS re-authentication and disconnection | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | [optional] 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers**](NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | Pointer to **string** | Choose the Host Mode for the access policy. | [optional] 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **int32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesDot1x**](NetworksNetworkIdSwitchAccessPoliciesDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | Pointer to **bool** | Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication | [optional] 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 
**Counts** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesCounts**](NetworksNetworkIdSwitchAccessPoliciesCounts.md) |  | [optional] 

## Methods

### NewInlineResponse200151

`func NewInlineResponse200151() *InlineResponse200151`

NewInlineResponse200151 instantiates a new InlineResponse200151 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200151WithDefaults

`func NewInlineResponse200151WithDefaults() *InlineResponse200151`

NewInlineResponse200151WithDefaults instantiates a new InlineResponse200151 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200151) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200151) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200151) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200151) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineResponse200151) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineResponse200151) GetRadiusServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineResponse200151) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineResponse200151) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadius

`func (o *InlineResponse200151) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineResponse200151) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineResponse200151) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineResponse200151) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetEnforceRadiusMonitoring

`func (o *InlineResponse200151) GetEnforceRadiusMonitoring() bool`

GetEnforceRadiusMonitoring returns the EnforceRadiusMonitoring field if non-nil, zero value otherwise.

### GetEnforceRadiusMonitoringOk

`func (o *InlineResponse200151) GetEnforceRadiusMonitoringOk() (*bool, bool)`

GetEnforceRadiusMonitoringOk returns a tuple with the EnforceRadiusMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceRadiusMonitoring

`func (o *InlineResponse200151) SetEnforceRadiusMonitoring(v bool)`

SetEnforceRadiusMonitoring sets EnforceRadiusMonitoring field to given value.

### HasEnforceRadiusMonitoring

`func (o *InlineResponse200151) HasEnforceRadiusMonitoring() bool`

HasEnforceRadiusMonitoring returns a boolean if a field has been set.

### GetGuestPortBouncing

`func (o *InlineResponse200151) GetGuestPortBouncing() bool`

GetGuestPortBouncing returns the GuestPortBouncing field if non-nil, zero value otherwise.

### GetGuestPortBouncingOk

`func (o *InlineResponse200151) GetGuestPortBouncingOk() (*bool, bool)`

GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPortBouncing

`func (o *InlineResponse200151) SetGuestPortBouncing(v bool)`

SetGuestPortBouncing sets GuestPortBouncing field to given value.

### HasGuestPortBouncing

`func (o *InlineResponse200151) HasGuestPortBouncing() bool`

HasGuestPortBouncing returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineResponse200151) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineResponse200151) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineResponse200151) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.

### HasRadiusTestingEnabled

`func (o *InlineResponse200151) HasRadiusTestingEnabled() bool`

HasRadiusTestingEnabled returns a boolean if a field has been set.

### GetRadiusCoaSupportEnabled

`func (o *InlineResponse200151) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *InlineResponse200151) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *InlineResponse200151) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.

### HasRadiusCoaSupportEnabled

`func (o *InlineResponse200151) HasRadiusCoaSupportEnabled() bool`

HasRadiusCoaSupportEnabled returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *InlineResponse200151) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineResponse200151) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineResponse200151) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *InlineResponse200151) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *InlineResponse200151) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineResponse200151) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineResponse200151) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineResponse200151) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *InlineResponse200151) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *InlineResponse200151) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *InlineResponse200151) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *InlineResponse200151) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *InlineResponse200151) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *InlineResponse200151) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *InlineResponse200151) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.

### HasHostMode

`func (o *InlineResponse200151) HasHostMode() bool`

HasHostMode returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineResponse200151) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineResponse200151) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineResponse200151) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineResponse200151) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *InlineResponse200151) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *InlineResponse200151) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *InlineResponse200151) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *InlineResponse200151) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *InlineResponse200151) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *InlineResponse200151) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *InlineResponse200151) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *InlineResponse200151) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### GetDot1x

`func (o *InlineResponse200151) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *InlineResponse200151) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *InlineResponse200151) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *InlineResponse200151) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *InlineResponse200151) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *InlineResponse200151) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *InlineResponse200151) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *InlineResponse200151) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *InlineResponse200151) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *InlineResponse200151) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *InlineResponse200151) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.

### HasUrlRedirectWalledGardenEnabled

`func (o *InlineResponse200151) HasUrlRedirectWalledGardenEnabled() bool`

HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenRanges

`func (o *InlineResponse200151) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *InlineResponse200151) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *InlineResponse200151) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *InlineResponse200151) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200151) GetCounts() NetworksNetworkIdSwitchAccessPoliciesCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200151) GetCountsOk() (*NetworksNetworkIdSwitchAccessPoliciesCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200151) SetCounts(v NetworksNetworkIdSwitchAccessPoliciesCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200151) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


