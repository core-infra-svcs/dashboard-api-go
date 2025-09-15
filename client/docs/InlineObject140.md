# InlineObject140

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the access policy(max length 255) | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers**](NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | [optional] 
**Radius** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius1**](NetworksNetworkIdSwitchAccessPoliciesRadius1.md) |  | [optional] 
**GuestPortBouncing** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusTestingEnabled** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusCoaSupportEnabled** | Pointer to **bool** | Change of authentication for RADIUS re-authentication and disconnection | [optional] 
**RadiusAccountingEnabled** | Pointer to **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | [optional] 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers**](NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | Pointer to **string** | Choose the Host Mode for the access policy. | [optional] 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **NullableInt32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesDot1x**](NetworksNetworkIdSwitchAccessPoliciesDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | Pointer to **bool** | Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 
**GuestGroupPolicyId** | Pointer to **NullableString** | Group policy Number for guest group policy | [optional] 
**GuestSgtId** | Pointer to **NullableInt32** | Security Group Tag ID for guest group policy | [optional] 

## Methods

### NewInlineObject140

`func NewInlineObject140() *InlineObject140`

NewInlineObject140 instantiates a new InlineObject140 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject140WithDefaults

`func NewInlineObject140WithDefaults() *InlineObject140`

NewInlineObject140WithDefaults instantiates a new InlineObject140 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject140) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject140) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject140) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject140) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineObject140) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineObject140) GetRadiusServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineObject140) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineObject140) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetRadius

`func (o *InlineObject140) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius1`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineObject140) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius1, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineObject140) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius1)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineObject140) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetGuestPortBouncing

`func (o *InlineObject140) GetGuestPortBouncing() bool`

GetGuestPortBouncing returns the GuestPortBouncing field if non-nil, zero value otherwise.

### GetGuestPortBouncingOk

`func (o *InlineObject140) GetGuestPortBouncingOk() (*bool, bool)`

GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPortBouncing

`func (o *InlineObject140) SetGuestPortBouncing(v bool)`

SetGuestPortBouncing sets GuestPortBouncing field to given value.

### HasGuestPortBouncing

`func (o *InlineObject140) HasGuestPortBouncing() bool`

HasGuestPortBouncing returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineObject140) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineObject140) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineObject140) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.

### HasRadiusTestingEnabled

`func (o *InlineObject140) HasRadiusTestingEnabled() bool`

HasRadiusTestingEnabled returns a boolean if a field has been set.

### GetRadiusCoaSupportEnabled

`func (o *InlineObject140) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *InlineObject140) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *InlineObject140) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.

### HasRadiusCoaSupportEnabled

`func (o *InlineObject140) HasRadiusCoaSupportEnabled() bool`

HasRadiusCoaSupportEnabled returns a boolean if a field has been set.

### GetRadiusAccountingEnabled

`func (o *InlineObject140) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineObject140) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineObject140) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.

### HasRadiusAccountingEnabled

`func (o *InlineObject140) HasRadiusAccountingEnabled() bool`

HasRadiusAccountingEnabled returns a boolean if a field has been set.

### GetRadiusAccountingServers

`func (o *InlineObject140) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineObject140) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineObject140) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineObject140) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *InlineObject140) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *InlineObject140) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *InlineObject140) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *InlineObject140) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *InlineObject140) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *InlineObject140) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *InlineObject140) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.

### HasHostMode

`func (o *InlineObject140) HasHostMode() bool`

HasHostMode returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineObject140) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject140) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject140) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject140) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *InlineObject140) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *InlineObject140) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *InlineObject140) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *InlineObject140) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *InlineObject140) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *InlineObject140) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *InlineObject140) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *InlineObject140) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### SetGuestVlanIdNil

`func (o *InlineObject140) SetGuestVlanIdNil(b bool)`

 SetGuestVlanIdNil sets the value for GuestVlanId to be an explicit nil

### UnsetGuestVlanId
`func (o *InlineObject140) UnsetGuestVlanId()`

UnsetGuestVlanId ensures that no value is present for GuestVlanId, not even an explicit nil
### GetDot1x

`func (o *InlineObject140) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *InlineObject140) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *InlineObject140) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *InlineObject140) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *InlineObject140) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *InlineObject140) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *InlineObject140) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *InlineObject140) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *InlineObject140) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *InlineObject140) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *InlineObject140) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.

### HasUrlRedirectWalledGardenEnabled

`func (o *InlineObject140) HasUrlRedirectWalledGardenEnabled() bool`

HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenRanges

`func (o *InlineObject140) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *InlineObject140) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *InlineObject140) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *InlineObject140) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.

### GetGuestGroupPolicyId

`func (o *InlineObject140) GetGuestGroupPolicyId() string`

GetGuestGroupPolicyId returns the GuestGroupPolicyId field if non-nil, zero value otherwise.

### GetGuestGroupPolicyIdOk

`func (o *InlineObject140) GetGuestGroupPolicyIdOk() (*string, bool)`

GetGuestGroupPolicyIdOk returns a tuple with the GuestGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestGroupPolicyId

`func (o *InlineObject140) SetGuestGroupPolicyId(v string)`

SetGuestGroupPolicyId sets GuestGroupPolicyId field to given value.

### HasGuestGroupPolicyId

`func (o *InlineObject140) HasGuestGroupPolicyId() bool`

HasGuestGroupPolicyId returns a boolean if a field has been set.

### SetGuestGroupPolicyIdNil

`func (o *InlineObject140) SetGuestGroupPolicyIdNil(b bool)`

 SetGuestGroupPolicyIdNil sets the value for GuestGroupPolicyId to be an explicit nil

### UnsetGuestGroupPolicyId
`func (o *InlineObject140) UnsetGuestGroupPolicyId()`

UnsetGuestGroupPolicyId ensures that no value is present for GuestGroupPolicyId, not even an explicit nil
### GetGuestSgtId

`func (o *InlineObject140) GetGuestSgtId() int32`

GetGuestSgtId returns the GuestSgtId field if non-nil, zero value otherwise.

### GetGuestSgtIdOk

`func (o *InlineObject140) GetGuestSgtIdOk() (*int32, bool)`

GetGuestSgtIdOk returns a tuple with the GuestSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSgtId

`func (o *InlineObject140) SetGuestSgtId(v int32)`

SetGuestSgtId sets GuestSgtId field to given value.

### HasGuestSgtId

`func (o *InlineObject140) HasGuestSgtId() bool`

HasGuestSgtId returns a boolean if a field has been set.

### SetGuestSgtIdNil

`func (o *InlineObject140) SetGuestSgtIdNil(b bool)`

 SetGuestSgtIdNil sets the value for GuestSgtId to be an explicit nil

### UnsetGuestSgtId
`func (o *InlineObject140) UnsetGuestSgtId()`

UnsetGuestSgtId ensures that no value is present for GuestSgtId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


