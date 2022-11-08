# InlineObject16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the switch port. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the switch port. | [optional] 
**Enabled** | Pointer to **bool** | The status of the switch port. | [optional] 
**PoeEnabled** | Pointer to **bool** | The PoE status of the switch port. | [optional] 
**Type** | Pointer to **string** | The type of the switch port (&#39;trunk&#39; or &#39;access&#39;). | [optional] 
**Vlan** | Pointer to **int32** | The VLAN of the switch port. A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **int32** | The voice VLAN of the switch port. Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the switch port. Only applicable to trunk ports. | [optional] 
**IsolationEnabled** | Pointer to **bool** | The isolation status of the switch port. | [optional] 
**RstpEnabled** | Pointer to **bool** | The rapid spanning tree protocol status. | [optional] 
**StpGuard** | Pointer to **string** | The state of the STP guard (&#39;disabled&#39;, &#39;root guard&#39;, &#39;bpdu guard&#39; or &#39;loop guard&#39;). | [optional] 
**LinkNegotiation** | Pointer to **string** | The link speed for the switch port. | [optional] 
**PortScheduleId** | Pointer to **string** | The ID of the port schedule. A value of null will clear the port schedule. | [optional] 
**Udld** | Pointer to **string** | The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only. | [optional] 
**AccessPolicyType** | Pointer to **string** | The type of the access policy of the switch port. Only applicable to access ports. Can be one of &#39;Open&#39;, &#39;Custom access policy&#39;, &#39;MAC allow list&#39; or &#39;Sticky MAC allow list&#39;. | [optional] 
**AccessPolicyNumber** | Pointer to **int32** | The number of a custom access policy to configure on the switch port. Only applicable when &#39;accessPolicyType&#39; is &#39;Custom access policy&#39;. | [optional] 
**MacAllowList** | Pointer to **[]string** | Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;. | [optional] 
**StickyMacAllowList** | Pointer to **[]string** | The initial list of MAC addresses for sticky Mac allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowListLimit** | Pointer to **int32** | The maximum number of MAC addresses for sticky MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StormControlEnabled** | Pointer to **bool** | The storm control status of the switch port. | [optional] 
**AdaptivePolicyGroupId** | Pointer to **string** | The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile. | [optional] 
**PeerSgtCapable** | Pointer to **bool** | If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile. | [optional] 
**FlexibleStackingEnabled** | Pointer to **bool** | For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled. | [optional] 

## Methods

### NewInlineObject16

`func NewInlineObject16() *InlineObject16`

NewInlineObject16 instantiates a new InlineObject16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject16WithDefaults

`func NewInlineObject16WithDefaults() *InlineObject16`

NewInlineObject16WithDefaults instantiates a new InlineObject16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject16) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject16) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject16) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject16) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject16) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject16) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject16) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject16) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject16) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject16) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject16) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject16) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *InlineObject16) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *InlineObject16) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *InlineObject16) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *InlineObject16) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineObject16) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject16) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject16) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject16) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *InlineObject16) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject16) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject16) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject16) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVoiceVlan

`func (o *InlineObject16) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *InlineObject16) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *InlineObject16) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *InlineObject16) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *InlineObject16) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *InlineObject16) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *InlineObject16) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *InlineObject16) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetIsolationEnabled

`func (o *InlineObject16) GetIsolationEnabled() bool`

GetIsolationEnabled returns the IsolationEnabled field if non-nil, zero value otherwise.

### GetIsolationEnabledOk

`func (o *InlineObject16) GetIsolationEnabledOk() (*bool, bool)`

GetIsolationEnabledOk returns a tuple with the IsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationEnabled

`func (o *InlineObject16) SetIsolationEnabled(v bool)`

SetIsolationEnabled sets IsolationEnabled field to given value.

### HasIsolationEnabled

`func (o *InlineObject16) HasIsolationEnabled() bool`

HasIsolationEnabled returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *InlineObject16) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineObject16) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineObject16) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineObject16) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *InlineObject16) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *InlineObject16) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *InlineObject16) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *InlineObject16) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineObject16) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineObject16) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineObject16) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineObject16) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetPortScheduleId

`func (o *InlineObject16) GetPortScheduleId() string`

GetPortScheduleId returns the PortScheduleId field if non-nil, zero value otherwise.

### GetPortScheduleIdOk

`func (o *InlineObject16) GetPortScheduleIdOk() (*string, bool)`

GetPortScheduleIdOk returns a tuple with the PortScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScheduleId

`func (o *InlineObject16) SetPortScheduleId(v string)`

SetPortScheduleId sets PortScheduleId field to given value.

### HasPortScheduleId

`func (o *InlineObject16) HasPortScheduleId() bool`

HasPortScheduleId returns a boolean if a field has been set.

### GetUdld

`func (o *InlineObject16) GetUdld() string`

GetUdld returns the Udld field if non-nil, zero value otherwise.

### GetUdldOk

`func (o *InlineObject16) GetUdldOk() (*string, bool)`

GetUdldOk returns a tuple with the Udld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdld

`func (o *InlineObject16) SetUdld(v string)`

SetUdld sets Udld field to given value.

### HasUdld

`func (o *InlineObject16) HasUdld() bool`

HasUdld returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineObject16) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject16) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject16) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject16) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetAccessPolicyNumber

`func (o *InlineObject16) GetAccessPolicyNumber() int32`

GetAccessPolicyNumber returns the AccessPolicyNumber field if non-nil, zero value otherwise.

### GetAccessPolicyNumberOk

`func (o *InlineObject16) GetAccessPolicyNumberOk() (*int32, bool)`

GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyNumber

`func (o *InlineObject16) SetAccessPolicyNumber(v int32)`

SetAccessPolicyNumber sets AccessPolicyNumber field to given value.

### HasAccessPolicyNumber

`func (o *InlineObject16) HasAccessPolicyNumber() bool`

HasAccessPolicyNumber returns a boolean if a field has been set.

### GetMacAllowList

`func (o *InlineObject16) GetMacAllowList() []string`

GetMacAllowList returns the MacAllowList field if non-nil, zero value otherwise.

### GetMacAllowListOk

`func (o *InlineObject16) GetMacAllowListOk() (*[]string, bool)`

GetMacAllowListOk returns a tuple with the MacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAllowList

`func (o *InlineObject16) SetMacAllowList(v []string)`

SetMacAllowList sets MacAllowList field to given value.

### HasMacAllowList

`func (o *InlineObject16) HasMacAllowList() bool`

HasMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *InlineObject16) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *InlineObject16) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *InlineObject16) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *InlineObject16) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *InlineObject16) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *InlineObject16) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *InlineObject16) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *InlineObject16) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.

### GetStormControlEnabled

`func (o *InlineObject16) GetStormControlEnabled() bool`

GetStormControlEnabled returns the StormControlEnabled field if non-nil, zero value otherwise.

### GetStormControlEnabledOk

`func (o *InlineObject16) GetStormControlEnabledOk() (*bool, bool)`

GetStormControlEnabledOk returns a tuple with the StormControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControlEnabled

`func (o *InlineObject16) SetStormControlEnabled(v bool)`

SetStormControlEnabled sets StormControlEnabled field to given value.

### HasStormControlEnabled

`func (o *InlineObject16) HasStormControlEnabled() bool`

HasStormControlEnabled returns a boolean if a field has been set.

### GetAdaptivePolicyGroupId

`func (o *InlineObject16) GetAdaptivePolicyGroupId() string`

GetAdaptivePolicyGroupId returns the AdaptivePolicyGroupId field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupIdOk

`func (o *InlineObject16) GetAdaptivePolicyGroupIdOk() (*string, bool)`

GetAdaptivePolicyGroupIdOk returns a tuple with the AdaptivePolicyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroupId

`func (o *InlineObject16) SetAdaptivePolicyGroupId(v string)`

SetAdaptivePolicyGroupId sets AdaptivePolicyGroupId field to given value.

### HasAdaptivePolicyGroupId

`func (o *InlineObject16) HasAdaptivePolicyGroupId() bool`

HasAdaptivePolicyGroupId returns a boolean if a field has been set.

### GetPeerSgtCapable

`func (o *InlineObject16) GetPeerSgtCapable() bool`

GetPeerSgtCapable returns the PeerSgtCapable field if non-nil, zero value otherwise.

### GetPeerSgtCapableOk

`func (o *InlineObject16) GetPeerSgtCapableOk() (*bool, bool)`

GetPeerSgtCapableOk returns a tuple with the PeerSgtCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSgtCapable

`func (o *InlineObject16) SetPeerSgtCapable(v bool)`

SetPeerSgtCapable sets PeerSgtCapable field to given value.

### HasPeerSgtCapable

`func (o *InlineObject16) HasPeerSgtCapable() bool`

HasPeerSgtCapable returns a boolean if a field has been set.

### GetFlexibleStackingEnabled

`func (o *InlineObject16) GetFlexibleStackingEnabled() bool`

GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field if non-nil, zero value otherwise.

### GetFlexibleStackingEnabledOk

`func (o *InlineObject16) GetFlexibleStackingEnabledOk() (*bool, bool)`

GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleStackingEnabled

`func (o *InlineObject16) SetFlexibleStackingEnabled(v bool)`

SetFlexibleStackingEnabled sets FlexibleStackingEnabled field to given value.

### HasFlexibleStackingEnabled

`func (o *InlineObject16) HasFlexibleStackingEnabled() bool`

HasFlexibleStackingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


