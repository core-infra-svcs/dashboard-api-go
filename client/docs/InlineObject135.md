# InlineObject135

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the access policy | 
**RadiusServers** | [**[]NetworksNetworkIdSwitchAccessPoliciesRadiusServers1**](NetworksNetworkIdSwitchAccessPoliciesRadiusServers1.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | 
**Radius** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius**](NetworksNetworkIdSwitchAccessPoliciesRadius.md) |  | [optional] 
**GuestPortBouncing** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusTestingEnabled** | **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | 
**RadiusCoaSupportEnabled** | **bool** | Change of authentication for RADIUS re-authentication and disconnection | 
**RadiusAccountingEnabled** | **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1**](NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | **string** | Choose the Host Mode for the access policy. | 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **int32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesDot1x**](NetworksNetworkIdSwitchAccessPoliciesDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | **bool** | Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 

## Methods

### NewInlineObject135

`func NewInlineObject135(name string, radiusServers []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1, radiusTestingEnabled bool, radiusCoaSupportEnabled bool, radiusAccountingEnabled bool, hostMode string, urlRedirectWalledGardenEnabled bool, ) *InlineObject135`

NewInlineObject135 instantiates a new InlineObject135 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject135WithDefaults

`func NewInlineObject135WithDefaults() *InlineObject135`

NewInlineObject135WithDefaults instantiates a new InlineObject135 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject135) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject135) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject135) SetName(v string)`

SetName sets Name field to given value.


### GetRadiusServers

`func (o *InlineObject135) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineObject135) GetRadiusServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesRadiusServers1, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineObject135) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1)`

SetRadiusServers sets RadiusServers field to given value.


### GetRadius

`func (o *InlineObject135) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineObject135) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineObject135) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineObject135) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetGuestPortBouncing

`func (o *InlineObject135) GetGuestPortBouncing() bool`

GetGuestPortBouncing returns the GuestPortBouncing field if non-nil, zero value otherwise.

### GetGuestPortBouncingOk

`func (o *InlineObject135) GetGuestPortBouncingOk() (*bool, bool)`

GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPortBouncing

`func (o *InlineObject135) SetGuestPortBouncing(v bool)`

SetGuestPortBouncing sets GuestPortBouncing field to given value.

### HasGuestPortBouncing

`func (o *InlineObject135) HasGuestPortBouncing() bool`

HasGuestPortBouncing returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineObject135) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineObject135) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineObject135) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.


### GetRadiusCoaSupportEnabled

`func (o *InlineObject135) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *InlineObject135) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *InlineObject135) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.


### GetRadiusAccountingEnabled

`func (o *InlineObject135) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineObject135) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineObject135) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.


### GetRadiusAccountingServers

`func (o *InlineObject135) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineObject135) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineObject135) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineObject135) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *InlineObject135) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *InlineObject135) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *InlineObject135) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *InlineObject135) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *InlineObject135) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *InlineObject135) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *InlineObject135) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.


### GetAccessPolicyType

`func (o *InlineObject135) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject135) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject135) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject135) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *InlineObject135) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *InlineObject135) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *InlineObject135) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *InlineObject135) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *InlineObject135) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *InlineObject135) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *InlineObject135) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *InlineObject135) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### GetDot1x

`func (o *InlineObject135) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *InlineObject135) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *InlineObject135) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *InlineObject135) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *InlineObject135) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *InlineObject135) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *InlineObject135) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *InlineObject135) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *InlineObject135) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *InlineObject135) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *InlineObject135) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.


### GetUrlRedirectWalledGardenRanges

`func (o *InlineObject135) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *InlineObject135) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *InlineObject135) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *InlineObject135) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


