# InlineObject28

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the switch port. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the switch port. | [optional] 
**Enabled** | Pointer to **bool** | The status of the switch port. | [optional] 
**PoeEnabled** | Pointer to **bool** | The PoE status of the switch port. | [optional] 
**Type** | Pointer to **string** | The type of the switch port (&#39;access&#39;, &#39;trunk&#39;, &#39;stack&#39;, &#39;routed&#39;, &#39;svl&#39; or &#39;dad&#39;). | [optional] 
**Vlan** | Pointer to **NullableInt32** | The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **NullableInt32** | The voice VLAN of the switch port. Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the switch port. Only applicable to trunk ports. | [optional] 
**IsolationEnabled** | Pointer to **bool** | The isolation status of the switch port. | [optional] 
**RstpEnabled** | Pointer to **bool** | The rapid spanning tree protocol status. | [optional] 
**StpGuard** | Pointer to **string** | The state of the STP guard (&#39;disabled&#39;, &#39;root guard&#39;, &#39;bpdu guard&#39; or &#39;loop guard&#39;). | [optional] 
**LinkNegotiation** | Pointer to **string** | The link speed for the switch port. | [optional] 
**PortScheduleId** | Pointer to **NullableString** | The ID of the port schedule. A value of null will clear the port schedule. | [optional] 
**Udld** | Pointer to **string** | The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only. | [optional] 
**AccessPolicyType** | Pointer to **string** | The type of the access policy of the switch port. Only applicable to access ports. Can be one of &#39;Open&#39;, &#39;Custom access policy&#39;, &#39;MAC allow list&#39; or &#39;Sticky MAC allow list&#39;. | [optional] 
**AccessPolicyNumber** | Pointer to **int32** | The number of a custom access policy to configure on the switch port. Only applicable when &#39;accessPolicyType&#39; is &#39;Custom access policy&#39;. | [optional] 
**MacAllowList** | Pointer to **[]string** | Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;. | [optional] 
**MacWhitelistLimit** | Pointer to **int32** | The maximum number of MAC addresses for regular MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;.           Note: Config only supported on verions greater than ms18 only for classic switches. | [optional] 
**StickyMacAllowList** | Pointer to **[]string** | The initial list of MAC addresses for sticky Mac allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowListLimit** | Pointer to **int32** | The maximum number of MAC addresses for sticky MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StormControlEnabled** | Pointer to **bool** | The storm control status of the switch port. | [optional] 
**AdaptivePolicyGroupId** | Pointer to **NullableString** | The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile. | [optional] 
**PeerSgtCapable** | Pointer to **bool** | If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile. | [optional] 
**FlexibleStackingEnabled** | Pointer to **bool** | For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled. | [optional] 
**DaiTrusted** | Pointer to **bool** | If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic. | [optional] 
**Profile** | Pointer to [**DevicesSerialSwitchPortsProfile**](DevicesSerialSwitchPortsProfile.md) |  | [optional] 
**Dot3az** | Pointer to [**DevicesSerialSwitchPortsDot3az**](DevicesSerialSwitchPortsDot3az.md) |  | [optional] 

## Methods

### NewInlineObject28

`func NewInlineObject28() *InlineObject28`

NewInlineObject28 instantiates a new InlineObject28 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject28WithDefaults

`func NewInlineObject28WithDefaults() *InlineObject28`

NewInlineObject28WithDefaults instantiates a new InlineObject28 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject28) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject28) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject28) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject28) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InlineObject28) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InlineObject28) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTags

`func (o *InlineObject28) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject28) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject28) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject28) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject28) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject28) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject28) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject28) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *InlineObject28) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *InlineObject28) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *InlineObject28) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *InlineObject28) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineObject28) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject28) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject28) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject28) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *InlineObject28) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject28) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject28) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject28) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *InlineObject28) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *InlineObject28) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetVoiceVlan

`func (o *InlineObject28) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *InlineObject28) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *InlineObject28) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *InlineObject28) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### SetVoiceVlanNil

`func (o *InlineObject28) SetVoiceVlanNil(b bool)`

 SetVoiceVlanNil sets the value for VoiceVlan to be an explicit nil

### UnsetVoiceVlan
`func (o *InlineObject28) UnsetVoiceVlan()`

UnsetVoiceVlan ensures that no value is present for VoiceVlan, not even an explicit nil
### GetAllowedVlans

`func (o *InlineObject28) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *InlineObject28) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *InlineObject28) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *InlineObject28) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetIsolationEnabled

`func (o *InlineObject28) GetIsolationEnabled() bool`

GetIsolationEnabled returns the IsolationEnabled field if non-nil, zero value otherwise.

### GetIsolationEnabledOk

`func (o *InlineObject28) GetIsolationEnabledOk() (*bool, bool)`

GetIsolationEnabledOk returns a tuple with the IsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationEnabled

`func (o *InlineObject28) SetIsolationEnabled(v bool)`

SetIsolationEnabled sets IsolationEnabled field to given value.

### HasIsolationEnabled

`func (o *InlineObject28) HasIsolationEnabled() bool`

HasIsolationEnabled returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *InlineObject28) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineObject28) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineObject28) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineObject28) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *InlineObject28) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *InlineObject28) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *InlineObject28) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *InlineObject28) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineObject28) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineObject28) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineObject28) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineObject28) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetPortScheduleId

`func (o *InlineObject28) GetPortScheduleId() string`

GetPortScheduleId returns the PortScheduleId field if non-nil, zero value otherwise.

### GetPortScheduleIdOk

`func (o *InlineObject28) GetPortScheduleIdOk() (*string, bool)`

GetPortScheduleIdOk returns a tuple with the PortScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScheduleId

`func (o *InlineObject28) SetPortScheduleId(v string)`

SetPortScheduleId sets PortScheduleId field to given value.

### HasPortScheduleId

`func (o *InlineObject28) HasPortScheduleId() bool`

HasPortScheduleId returns a boolean if a field has been set.

### SetPortScheduleIdNil

`func (o *InlineObject28) SetPortScheduleIdNil(b bool)`

 SetPortScheduleIdNil sets the value for PortScheduleId to be an explicit nil

### UnsetPortScheduleId
`func (o *InlineObject28) UnsetPortScheduleId()`

UnsetPortScheduleId ensures that no value is present for PortScheduleId, not even an explicit nil
### GetUdld

`func (o *InlineObject28) GetUdld() string`

GetUdld returns the Udld field if non-nil, zero value otherwise.

### GetUdldOk

`func (o *InlineObject28) GetUdldOk() (*string, bool)`

GetUdldOk returns a tuple with the Udld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdld

`func (o *InlineObject28) SetUdld(v string)`

SetUdld sets Udld field to given value.

### HasUdld

`func (o *InlineObject28) HasUdld() bool`

HasUdld returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineObject28) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject28) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject28) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject28) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetAccessPolicyNumber

`func (o *InlineObject28) GetAccessPolicyNumber() int32`

GetAccessPolicyNumber returns the AccessPolicyNumber field if non-nil, zero value otherwise.

### GetAccessPolicyNumberOk

`func (o *InlineObject28) GetAccessPolicyNumberOk() (*int32, bool)`

GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyNumber

`func (o *InlineObject28) SetAccessPolicyNumber(v int32)`

SetAccessPolicyNumber sets AccessPolicyNumber field to given value.

### HasAccessPolicyNumber

`func (o *InlineObject28) HasAccessPolicyNumber() bool`

HasAccessPolicyNumber returns a boolean if a field has been set.

### GetMacAllowList

`func (o *InlineObject28) GetMacAllowList() []string`

GetMacAllowList returns the MacAllowList field if non-nil, zero value otherwise.

### GetMacAllowListOk

`func (o *InlineObject28) GetMacAllowListOk() (*[]string, bool)`

GetMacAllowListOk returns a tuple with the MacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAllowList

`func (o *InlineObject28) SetMacAllowList(v []string)`

SetMacAllowList sets MacAllowList field to given value.

### HasMacAllowList

`func (o *InlineObject28) HasMacAllowList() bool`

HasMacAllowList returns a boolean if a field has been set.

### GetMacWhitelistLimit

`func (o *InlineObject28) GetMacWhitelistLimit() int32`

GetMacWhitelistLimit returns the MacWhitelistLimit field if non-nil, zero value otherwise.

### GetMacWhitelistLimitOk

`func (o *InlineObject28) GetMacWhitelistLimitOk() (*int32, bool)`

GetMacWhitelistLimitOk returns a tuple with the MacWhitelistLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacWhitelistLimit

`func (o *InlineObject28) SetMacWhitelistLimit(v int32)`

SetMacWhitelistLimit sets MacWhitelistLimit field to given value.

### HasMacWhitelistLimit

`func (o *InlineObject28) HasMacWhitelistLimit() bool`

HasMacWhitelistLimit returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *InlineObject28) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *InlineObject28) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *InlineObject28) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *InlineObject28) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *InlineObject28) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *InlineObject28) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *InlineObject28) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *InlineObject28) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.

### GetStormControlEnabled

`func (o *InlineObject28) GetStormControlEnabled() bool`

GetStormControlEnabled returns the StormControlEnabled field if non-nil, zero value otherwise.

### GetStormControlEnabledOk

`func (o *InlineObject28) GetStormControlEnabledOk() (*bool, bool)`

GetStormControlEnabledOk returns a tuple with the StormControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControlEnabled

`func (o *InlineObject28) SetStormControlEnabled(v bool)`

SetStormControlEnabled sets StormControlEnabled field to given value.

### HasStormControlEnabled

`func (o *InlineObject28) HasStormControlEnabled() bool`

HasStormControlEnabled returns a boolean if a field has been set.

### GetAdaptivePolicyGroupId

`func (o *InlineObject28) GetAdaptivePolicyGroupId() string`

GetAdaptivePolicyGroupId returns the AdaptivePolicyGroupId field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupIdOk

`func (o *InlineObject28) GetAdaptivePolicyGroupIdOk() (*string, bool)`

GetAdaptivePolicyGroupIdOk returns a tuple with the AdaptivePolicyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroupId

`func (o *InlineObject28) SetAdaptivePolicyGroupId(v string)`

SetAdaptivePolicyGroupId sets AdaptivePolicyGroupId field to given value.

### HasAdaptivePolicyGroupId

`func (o *InlineObject28) HasAdaptivePolicyGroupId() bool`

HasAdaptivePolicyGroupId returns a boolean if a field has been set.

### SetAdaptivePolicyGroupIdNil

`func (o *InlineObject28) SetAdaptivePolicyGroupIdNil(b bool)`

 SetAdaptivePolicyGroupIdNil sets the value for AdaptivePolicyGroupId to be an explicit nil

### UnsetAdaptivePolicyGroupId
`func (o *InlineObject28) UnsetAdaptivePolicyGroupId()`

UnsetAdaptivePolicyGroupId ensures that no value is present for AdaptivePolicyGroupId, not even an explicit nil
### GetPeerSgtCapable

`func (o *InlineObject28) GetPeerSgtCapable() bool`

GetPeerSgtCapable returns the PeerSgtCapable field if non-nil, zero value otherwise.

### GetPeerSgtCapableOk

`func (o *InlineObject28) GetPeerSgtCapableOk() (*bool, bool)`

GetPeerSgtCapableOk returns a tuple with the PeerSgtCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSgtCapable

`func (o *InlineObject28) SetPeerSgtCapable(v bool)`

SetPeerSgtCapable sets PeerSgtCapable field to given value.

### HasPeerSgtCapable

`func (o *InlineObject28) HasPeerSgtCapable() bool`

HasPeerSgtCapable returns a boolean if a field has been set.

### GetFlexibleStackingEnabled

`func (o *InlineObject28) GetFlexibleStackingEnabled() bool`

GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field if non-nil, zero value otherwise.

### GetFlexibleStackingEnabledOk

`func (o *InlineObject28) GetFlexibleStackingEnabledOk() (*bool, bool)`

GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleStackingEnabled

`func (o *InlineObject28) SetFlexibleStackingEnabled(v bool)`

SetFlexibleStackingEnabled sets FlexibleStackingEnabled field to given value.

### HasFlexibleStackingEnabled

`func (o *InlineObject28) HasFlexibleStackingEnabled() bool`

HasFlexibleStackingEnabled returns a boolean if a field has been set.

### GetDaiTrusted

`func (o *InlineObject28) GetDaiTrusted() bool`

GetDaiTrusted returns the DaiTrusted field if non-nil, zero value otherwise.

### GetDaiTrustedOk

`func (o *InlineObject28) GetDaiTrustedOk() (*bool, bool)`

GetDaiTrustedOk returns a tuple with the DaiTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaiTrusted

`func (o *InlineObject28) SetDaiTrusted(v bool)`

SetDaiTrusted sets DaiTrusted field to given value.

### HasDaiTrusted

`func (o *InlineObject28) HasDaiTrusted() bool`

HasDaiTrusted returns a boolean if a field has been set.

### GetProfile

`func (o *InlineObject28) GetProfile() DevicesSerialSwitchPortsProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineObject28) GetProfileOk() (*DevicesSerialSwitchPortsProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineObject28) SetProfile(v DevicesSerialSwitchPortsProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineObject28) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetDot3az

`func (o *InlineObject28) GetDot3az() DevicesSerialSwitchPortsDot3az`

GetDot3az returns the Dot3az field if non-nil, zero value otherwise.

### GetDot3azOk

`func (o *InlineObject28) GetDot3azOk() (*DevicesSerialSwitchPortsDot3az, bool)`

GetDot3azOk returns a tuple with the Dot3az field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot3az

`func (o *InlineObject28) SetDot3az(v DevicesSerialSwitchPortsDot3az)`

SetDot3az sets Dot3az field to given value.

### HasDot3az

`func (o *InlineObject28) HasDot3az() bool`

HasDot3az returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


