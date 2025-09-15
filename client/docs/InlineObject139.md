# InlineObject139

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the access policy(max length 255) | 
**RadiusServers** | [**[]NetworksNetworkIdSwitchAccessPoliciesRadiusServers1**](NetworksNetworkIdSwitchAccessPoliciesRadiusServers1.md) | List of RADIUS servers to require connecting devices to authenticate against before granting network access | 
**Radius** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadius1**](NetworksNetworkIdSwitchAccessPoliciesRadius1.md) |  | [optional] 
**GuestPortBouncing** | Pointer to **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | [optional] 
**RadiusTestingEnabled** | **bool** | If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers | 
**RadiusCoaSupportEnabled** | **bool** | Change of authentication for RADIUS re-authentication and disconnection | 
**RadiusAccountingEnabled** | **bool** | Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients | 
**RadiusAccountingServers** | Pointer to [**[]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1**](NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1.md) | List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access | [optional] 
**RadiusGroupAttribute** | Pointer to **string** | Acceptable values are &#x60;\&quot;\&quot;&#x60; for None, or &#x60;\&quot;11\&quot;&#x60; for Group Policies ACL | [optional] 
**HostMode** | **string** | Choose the Host Mode for the access policy. | 
**AccessPolicyType** | Pointer to **string** | Access Type of the policy. Automatically &#39;Hybrid authentication&#39; when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**IncreaseAccessSpeed** | Pointer to **bool** | Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is &#39;Hybrid Authentication. | [optional] 
**GuestVlanId** | Pointer to **NullableInt32** | ID for the guest VLAN allow unauthorized devices access to limited network resources | [optional] 
**Dot1x** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesDot1x**](NetworksNetworkIdSwitchAccessPoliciesDot1x.md) |  | [optional] 
**VoiceVlanClients** | Pointer to **bool** | CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is &#39;Multi-Domain&#39;. | [optional] 
**UrlRedirectWalledGardenEnabled** | **bool** | Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | 
**UrlRedirectWalledGardenRanges** | Pointer to **[]string** | IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication | [optional] 
**GuestGroupPolicyId** | Pointer to **NullableString** | Group policy Number for guest group policy | [optional] 
**GuestSgtId** | Pointer to **NullableInt32** | Security Group Tag ID for guest group policy | [optional] 

## Methods

### NewInlineObject139

`func NewInlineObject139(name string, radiusServers []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1, radiusTestingEnabled bool, radiusCoaSupportEnabled bool, radiusAccountingEnabled bool, hostMode string, urlRedirectWalledGardenEnabled bool, ) *InlineObject139`

NewInlineObject139 instantiates a new InlineObject139 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject139WithDefaults

`func NewInlineObject139WithDefaults() *InlineObject139`

NewInlineObject139WithDefaults instantiates a new InlineObject139 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject139) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject139) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject139) SetName(v string)`

SetName sets Name field to given value.


### GetRadiusServers

`func (o *InlineObject139) GetRadiusServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineObject139) GetRadiusServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesRadiusServers1, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineObject139) SetRadiusServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusServers1)`

SetRadiusServers sets RadiusServers field to given value.


### GetRadius

`func (o *InlineObject139) GetRadius() NetworksNetworkIdSwitchAccessPoliciesRadius1`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineObject139) GetRadiusOk() (*NetworksNetworkIdSwitchAccessPoliciesRadius1, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineObject139) SetRadius(v NetworksNetworkIdSwitchAccessPoliciesRadius1)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineObject139) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetGuestPortBouncing

`func (o *InlineObject139) GetGuestPortBouncing() bool`

GetGuestPortBouncing returns the GuestPortBouncing field if non-nil, zero value otherwise.

### GetGuestPortBouncingOk

`func (o *InlineObject139) GetGuestPortBouncingOk() (*bool, bool)`

GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPortBouncing

`func (o *InlineObject139) SetGuestPortBouncing(v bool)`

SetGuestPortBouncing sets GuestPortBouncing field to given value.

### HasGuestPortBouncing

`func (o *InlineObject139) HasGuestPortBouncing() bool`

HasGuestPortBouncing returns a boolean if a field has been set.

### GetRadiusTestingEnabled

`func (o *InlineObject139) GetRadiusTestingEnabled() bool`

GetRadiusTestingEnabled returns the RadiusTestingEnabled field if non-nil, zero value otherwise.

### GetRadiusTestingEnabledOk

`func (o *InlineObject139) GetRadiusTestingEnabledOk() (*bool, bool)`

GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusTestingEnabled

`func (o *InlineObject139) SetRadiusTestingEnabled(v bool)`

SetRadiusTestingEnabled sets RadiusTestingEnabled field to given value.


### GetRadiusCoaSupportEnabled

`func (o *InlineObject139) GetRadiusCoaSupportEnabled() bool`

GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field if non-nil, zero value otherwise.

### GetRadiusCoaSupportEnabledOk

`func (o *InlineObject139) GetRadiusCoaSupportEnabledOk() (*bool, bool)`

GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusCoaSupportEnabled

`func (o *InlineObject139) SetRadiusCoaSupportEnabled(v bool)`

SetRadiusCoaSupportEnabled sets RadiusCoaSupportEnabled field to given value.


### GetRadiusAccountingEnabled

`func (o *InlineObject139) GetRadiusAccountingEnabled() bool`

GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field if non-nil, zero value otherwise.

### GetRadiusAccountingEnabledOk

`func (o *InlineObject139) GetRadiusAccountingEnabledOk() (*bool, bool)`

GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnabled

`func (o *InlineObject139) SetRadiusAccountingEnabled(v bool)`

SetRadiusAccountingEnabled sets RadiusAccountingEnabled field to given value.


### GetRadiusAccountingServers

`func (o *InlineObject139) GetRadiusAccountingServers() []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1`

GetRadiusAccountingServers returns the RadiusAccountingServers field if non-nil, zero value otherwise.

### GetRadiusAccountingServersOk

`func (o *InlineObject139) GetRadiusAccountingServersOk() (*[]NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1, bool)`

GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingServers

`func (o *InlineObject139) SetRadiusAccountingServers(v []NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1)`

SetRadiusAccountingServers sets RadiusAccountingServers field to given value.

### HasRadiusAccountingServers

`func (o *InlineObject139) HasRadiusAccountingServers() bool`

HasRadiusAccountingServers returns a boolean if a field has been set.

### GetRadiusGroupAttribute

`func (o *InlineObject139) GetRadiusGroupAttribute() string`

GetRadiusGroupAttribute returns the RadiusGroupAttribute field if non-nil, zero value otherwise.

### GetRadiusGroupAttributeOk

`func (o *InlineObject139) GetRadiusGroupAttributeOk() (*string, bool)`

GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroupAttribute

`func (o *InlineObject139) SetRadiusGroupAttribute(v string)`

SetRadiusGroupAttribute sets RadiusGroupAttribute field to given value.

### HasRadiusGroupAttribute

`func (o *InlineObject139) HasRadiusGroupAttribute() bool`

HasRadiusGroupAttribute returns a boolean if a field has been set.

### GetHostMode

`func (o *InlineObject139) GetHostMode() string`

GetHostMode returns the HostMode field if non-nil, zero value otherwise.

### GetHostModeOk

`func (o *InlineObject139) GetHostModeOk() (*string, bool)`

GetHostModeOk returns a tuple with the HostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMode

`func (o *InlineObject139) SetHostMode(v string)`

SetHostMode sets HostMode field to given value.


### GetAccessPolicyType

`func (o *InlineObject139) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject139) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject139) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject139) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetIncreaseAccessSpeed

`func (o *InlineObject139) GetIncreaseAccessSpeed() bool`

GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field if non-nil, zero value otherwise.

### GetIncreaseAccessSpeedOk

`func (o *InlineObject139) GetIncreaseAccessSpeedOk() (*bool, bool)`

GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseAccessSpeed

`func (o *InlineObject139) SetIncreaseAccessSpeed(v bool)`

SetIncreaseAccessSpeed sets IncreaseAccessSpeed field to given value.

### HasIncreaseAccessSpeed

`func (o *InlineObject139) HasIncreaseAccessSpeed() bool`

HasIncreaseAccessSpeed returns a boolean if a field has been set.

### GetGuestVlanId

`func (o *InlineObject139) GetGuestVlanId() int32`

GetGuestVlanId returns the GuestVlanId field if non-nil, zero value otherwise.

### GetGuestVlanIdOk

`func (o *InlineObject139) GetGuestVlanIdOk() (*int32, bool)`

GetGuestVlanIdOk returns a tuple with the GuestVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanId

`func (o *InlineObject139) SetGuestVlanId(v int32)`

SetGuestVlanId sets GuestVlanId field to given value.

### HasGuestVlanId

`func (o *InlineObject139) HasGuestVlanId() bool`

HasGuestVlanId returns a boolean if a field has been set.

### SetGuestVlanIdNil

`func (o *InlineObject139) SetGuestVlanIdNil(b bool)`

 SetGuestVlanIdNil sets the value for GuestVlanId to be an explicit nil

### UnsetGuestVlanId
`func (o *InlineObject139) UnsetGuestVlanId()`

UnsetGuestVlanId ensures that no value is present for GuestVlanId, not even an explicit nil
### GetDot1x

`func (o *InlineObject139) GetDot1x() NetworksNetworkIdSwitchAccessPoliciesDot1x`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *InlineObject139) GetDot1xOk() (*NetworksNetworkIdSwitchAccessPoliciesDot1x, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *InlineObject139) SetDot1x(v NetworksNetworkIdSwitchAccessPoliciesDot1x)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *InlineObject139) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetVoiceVlanClients

`func (o *InlineObject139) GetVoiceVlanClients() bool`

GetVoiceVlanClients returns the VoiceVlanClients field if non-nil, zero value otherwise.

### GetVoiceVlanClientsOk

`func (o *InlineObject139) GetVoiceVlanClientsOk() (*bool, bool)`

GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlanClients

`func (o *InlineObject139) SetVoiceVlanClients(v bool)`

SetVoiceVlanClients sets VoiceVlanClients field to given value.

### HasVoiceVlanClients

`func (o *InlineObject139) HasVoiceVlanClients() bool`

HasVoiceVlanClients returns a boolean if a field has been set.

### GetUrlRedirectWalledGardenEnabled

`func (o *InlineObject139) GetUrlRedirectWalledGardenEnabled() bool`

GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenEnabledOk

`func (o *InlineObject139) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool)`

GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenEnabled

`func (o *InlineObject139) SetUrlRedirectWalledGardenEnabled(v bool)`

SetUrlRedirectWalledGardenEnabled sets UrlRedirectWalledGardenEnabled field to given value.


### GetUrlRedirectWalledGardenRanges

`func (o *InlineObject139) GetUrlRedirectWalledGardenRanges() []string`

GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field if non-nil, zero value otherwise.

### GetUrlRedirectWalledGardenRangesOk

`func (o *InlineObject139) GetUrlRedirectWalledGardenRangesOk() (*[]string, bool)`

GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirectWalledGardenRanges

`func (o *InlineObject139) SetUrlRedirectWalledGardenRanges(v []string)`

SetUrlRedirectWalledGardenRanges sets UrlRedirectWalledGardenRanges field to given value.

### HasUrlRedirectWalledGardenRanges

`func (o *InlineObject139) HasUrlRedirectWalledGardenRanges() bool`

HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.

### GetGuestGroupPolicyId

`func (o *InlineObject139) GetGuestGroupPolicyId() string`

GetGuestGroupPolicyId returns the GuestGroupPolicyId field if non-nil, zero value otherwise.

### GetGuestGroupPolicyIdOk

`func (o *InlineObject139) GetGuestGroupPolicyIdOk() (*string, bool)`

GetGuestGroupPolicyIdOk returns a tuple with the GuestGroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestGroupPolicyId

`func (o *InlineObject139) SetGuestGroupPolicyId(v string)`

SetGuestGroupPolicyId sets GuestGroupPolicyId field to given value.

### HasGuestGroupPolicyId

`func (o *InlineObject139) HasGuestGroupPolicyId() bool`

HasGuestGroupPolicyId returns a boolean if a field has been set.

### SetGuestGroupPolicyIdNil

`func (o *InlineObject139) SetGuestGroupPolicyIdNil(b bool)`

 SetGuestGroupPolicyIdNil sets the value for GuestGroupPolicyId to be an explicit nil

### UnsetGuestGroupPolicyId
`func (o *InlineObject139) UnsetGuestGroupPolicyId()`

UnsetGuestGroupPolicyId ensures that no value is present for GuestGroupPolicyId, not even an explicit nil
### GetGuestSgtId

`func (o *InlineObject139) GetGuestSgtId() int32`

GetGuestSgtId returns the GuestSgtId field if non-nil, zero value otherwise.

### GetGuestSgtIdOk

`func (o *InlineObject139) GetGuestSgtIdOk() (*int32, bool)`

GetGuestSgtIdOk returns a tuple with the GuestSgtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSgtId

`func (o *InlineObject139) SetGuestSgtId(v int32)`

SetGuestSgtId sets GuestSgtId field to given value.

### HasGuestSgtId

`func (o *InlineObject139) HasGuestSgtId() bool`

HasGuestSgtId returns a boolean if a field has been set.

### SetGuestSgtIdNil

`func (o *InlineObject139) SetGuestSgtIdNil(b bool)`

 SetGuestSgtIdNil sets the value for GuestSgtId to be an explicit nil

### UnsetGuestSgtId
`func (o *InlineObject139) UnsetGuestSgtId()`

UnsetGuestSgtId ensures that no value is present for GuestSgtId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


