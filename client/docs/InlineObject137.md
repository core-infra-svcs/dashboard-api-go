# InlineObject137

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the access policy(max length 255) | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers**](NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | [optional] 
**Radius** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius**](NetworksNetworkIdSwitchAccessPoliciesRadius.md) |  | [optional] 
**GuestPortBouncing** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusTestingEnabled** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusCoaSupportEnabled** | Pointer to **bool** | Change of authentication for RADIUS re-authentication and disconnection | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | [optional] 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers**](NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | Pointer to **string** | Choose the Host Mode for the access policy. | [optional] 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **int32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesDot1x**](NetworksNetworkIdSwitchAccessPoliciesDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | Pointer to **bool** | Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 

## Methods

### NewInlineObject137

`func NewInlineObject137() *InlineObject137`

NewInlineObject137 instantiates a new InlineObject137 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject137WithDefaults

`func NewInlineObject137WithDefaults() *InlineObject137`

NewInlineObject137WithDefaults instantiates a new InlineObject137 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject137) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject137) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject137) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject137) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineObject137) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineObject137) GetRadiusServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineObject137) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineObject137) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadius

`func (o *InlineObject137) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineObject137) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineObject137) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineObject137) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetGuestPortBouncing

`func (o *InlineObject137) GetGuestPortBouncing() bool`

GetGuestPortBouncing returns the GuestPortBouncing field if non-nil, zero value otherwise.

### GetGuestPortBouncingOk

`func (o *InlineObject137) GetGuestPortBouncingOk() (*bool, bool)`

GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPortBouncing

`func (o *InlineObject137) SetGuestPortBouncing(v bool)`

SetGuestPortBouncing sets GuestPortBouncing field to given value.

### HasGuestPortBouncing

`func (o *InlineObject137) HasGuestPortBouncing() bool`

HasGuestPortBouncing returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineObject137) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineObject137) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineObject137) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.

### HasRadiusTestingEnabled

`func (o *InlineObject137) HasRadiusTestingEnabled() bool`

HasRadiusTestingEnabled returns a boolean if a field has been set.

### GetRadiusCoaSupportEnabled

`func (o *InlineObject137) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *InlineObject137) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *InlineObject137) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.

### HasRadiusCoaSupportEnabled

`func (o *InlineObject137) HasRadiusCoaSupportEnabled() bool`

HasRadiusCoaSupportEnabled returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *InlineObject137) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineObject137) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineObject137) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *InlineObject137) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *InlineObject137) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineObject137) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineObject137) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineObject137) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *InlineObject137) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *InlineObject137) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *InlineObject137) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *InlineObject137) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *InlineObject137) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *InlineObject137) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *InlineObject137) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.

### HasHostMode

`func (o *InlineObject137) HasHostMode() bool`

HasHostMode returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineObject137) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject137) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject137) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject137) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *InlineObject137) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *InlineObject137) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *InlineObject137) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *InlineObject137) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *InlineObject137) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *InlineObject137) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *InlineObject137) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *InlineObject137) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### GetDot1x

`func (o *InlineObject137) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *InlineObject137) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *InlineObject137) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *InlineObject137) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *InlineObject137) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *InlineObject137) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *InlineObject137) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *InlineObject137) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *InlineObject137) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *InlineObject137) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *InlineObject137) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.

### HasUrlRedirectWalledGardenEnabled

`func (o *InlineObject137) HasUrlRedirectWalledGardenEnabled() bool`

HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenRanges

`func (o *InlineObject137) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *InlineObject137) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *InlineObject137) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *InlineObject137) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


