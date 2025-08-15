# InlineObject254

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the switch template port. | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of the switch template port. | [optional] 
**Enabled** | Pointer to **bool** | The status of the switch template port. | [optional] 
**PoeEnabled** | Pointer to **bool** | The PoE status of the switch template port. | [optional] 
**Type** | Pointer to **string** | The type of the switch template port (&#39;access&#39;, &#39;trunk&#39;, &#39;stack&#39;, &#39;routed&#39;, &#39;svl&#39; or &#39;dad&#39;). | [optional] 
**Vlan** | Pointer to **NullableInt32** | The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports. | [optional] 
**VoiceVlan** | Pointer to **NullableInt32** | The voice VLAN of the switch template port. Only applicable to access ports. | [optional] 
**AllowedVlans** | Pointer to **string** | The VLANs allowed on the switch template port. Only applicable to trunk ports. | [optional] 
**IsolationEnabled** | Pointer to **bool** | The isolation status of the switch template port. | [optional] 
**RstpEnabled** | Pointer to **bool** | The rapid spanning tree protocol status. | [optional] 
**StpGuard** | Pointer to **string** | The state of the STP guard (&#39;disabled&#39;, &#39;root guard&#39;, &#39;bpdu guard&#39; or &#39;loop guard&#39;). | [optional] 
**LinkNegotiation** | Pointer to **string** | The link speed for the switch template port. | [optional] 
**PortScheduleId** | Pointer to **NullableString** | The ID of the port schedule. A value of null will clear the port schedule. | [optional] 
**Udld** | Pointer to **string** | The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only. | [optional] 
**AccessPolicyType** | Pointer to **string** | The type of the access policy of the switch template port. Only applicable to access ports. Can be one of &#39;Open&#39;, &#39;Custom access policy&#39;, &#39;MAC allow list&#39; or &#39;Sticky MAC allow list&#39;. | [optional] 
**AccessPolicyNumber** | Pointer to **int32** | The number of a custom access policy to configure on the switch template port. Only applicable when &#39;accessPolicyType&#39; is &#39;Custom access policy&#39;. | [optional] 
**MacAllowList** | Pointer to **[]string** | Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;. | [optional] 
**MacWhitelistLimit** | Pointer to **int32** | The maximum number of MAC addresses for regular MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;MAC allow list&#39;.           Note: Config only supported on verions greater than ms18 only for classic switches. | [optional] 
**StickyMacAllowList** | Pointer to **[]string** | The initial list of MAC addresses for sticky Mac allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StickyMacAllowListLimit** | Pointer to **int32** | The maximum number of MAC addresses for sticky MAC allow list. Only applicable when &#39;accessPolicyType&#39; is &#39;Sticky MAC allow list&#39;. | [optional] 
**StormControlEnabled** | Pointer to **bool** | The storm control status of the switch template port. | [optional] 
**FlexibleStackingEnabled** | Pointer to **bool** | For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled. | [optional] 
**DaiTrusted** | Pointer to **bool** | If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic. | [optional] 
**Profile** | Pointer to [**DevicesSerialSwitchPortsProfile**](DevicesSerialSwitchPortsProfile.md) |  | [optional] 
**Dot3az** | Pointer to [**OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az**](OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az.md) |  | [optional] 

## Methods

### NewInlineObject254

`func NewInlineObject254() *InlineObject254`

NewInlineObject254 instantiates a new InlineObject254 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject254WithDefaults

`func NewInlineObject254WithDefaults() *InlineObject254`

NewInlineObject254WithDefaults instantiates a new InlineObject254 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject254) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject254) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject254) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject254) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InlineObject254) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InlineObject254) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTags

`func (o *InlineObject254) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject254) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject254) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject254) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject254) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject254) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject254) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject254) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoeEnabled

`func (o *InlineObject254) GetPoeEnabled() bool`

GetPoeEnabled returns the PoeEnabled field if non-nil, zero value otherwise.

### GetPoeEnabledOk

`func (o *InlineObject254) GetPoeEnabledOk() (*bool, bool)`

GetPoeEnabledOk returns a tuple with the PoeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeEnabled

`func (o *InlineObject254) SetPoeEnabled(v bool)`

SetPoeEnabled sets PoeEnabled field to given value.

### HasPoeEnabled

`func (o *InlineObject254) HasPoeEnabled() bool`

HasPoeEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineObject254) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject254) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject254) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject254) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlan

`func (o *InlineObject254) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineObject254) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineObject254) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineObject254) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *InlineObject254) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *InlineObject254) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetVoiceVlan

`func (o *InlineObject254) GetVoiceVlan() int32`

GetVoiceVlan returns the VoiceVlan field if non-nil, zero value otherwise.

### GetVoiceVlanOk

`func (o *InlineObject254) GetVoiceVlanOk() (*int32, bool)`

GetVoiceVlanOk returns a tuple with the VoiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceVlan

`func (o *InlineObject254) SetVoiceVlan(v int32)`

SetVoiceVlan sets VoiceVlan field to given value.

### HasVoiceVlan

`func (o *InlineObject254) HasVoiceVlan() bool`

HasVoiceVlan returns a boolean if a field has been set.

### SetVoiceVlanNil

`func (o *InlineObject254) SetVoiceVlanNil(b bool)`

 SetVoiceVlanNil sets the value for VoiceVlan to be an explicit nil

### UnsetVoiceVlan
`func (o *InlineObject254) UnsetVoiceVlan()`

UnsetVoiceVlan ensures that no value is present for VoiceVlan, not even an explicit nil
### GetAllowedVlans

`func (o *InlineObject254) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *InlineObject254) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *InlineObject254) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *InlineObject254) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetIsolationEnabled

`func (o *InlineObject254) GetIsolationEnabled() bool`

GetIsolationEnabled returns the IsolationEnabled field if non-nil, zero value otherwise.

### GetIsolationEnabledOk

`func (o *InlineObject254) GetIsolationEnabledOk() (*bool, bool)`

GetIsolationEnabledOk returns a tuple with the IsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationEnabled

`func (o *InlineObject254) SetIsolationEnabled(v bool)`

SetIsolationEnabled sets IsolationEnabled field to given value.

### HasIsolationEnabled

`func (o *InlineObject254) HasIsolationEnabled() bool`

HasIsolationEnabled returns a boolean if a field has been set.

### GetRstpEnabled

`func (o *InlineObject254) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineObject254) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineObject254) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineObject254) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpGuard

`func (o *InlineObject254) GetStpGuard() string`

GetStpGuard returns the StpGuard field if non-nil, zero value otherwise.

### GetStpGuardOk

`func (o *InlineObject254) GetStpGuardOk() (*string, bool)`

GetStpGuardOk returns a tuple with the StpGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpGuard

`func (o *InlineObject254) SetStpGuard(v string)`

SetStpGuard sets StpGuard field to given value.

### HasStpGuard

`func (o *InlineObject254) HasStpGuard() bool`

HasStpGuard returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *InlineObject254) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *InlineObject254) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *InlineObject254) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *InlineObject254) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetPortScheduleId

`func (o *InlineObject254) GetPortScheduleId() string`

GetPortScheduleId returns the PortScheduleId field if non-nil, zero value otherwise.

### GetPortScheduleIdOk

`func (o *InlineObject254) GetPortScheduleIdOk() (*string, bool)`

GetPortScheduleIdOk returns a tuple with the PortScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScheduleId

`func (o *InlineObject254) SetPortScheduleId(v string)`

SetPortScheduleId sets PortScheduleId field to given value.

### HasPortScheduleId

`func (o *InlineObject254) HasPortScheduleId() bool`

HasPortScheduleId returns a boolean if a field has been set.

### SetPortScheduleIdNil

`func (o *InlineObject254) SetPortScheduleIdNil(b bool)`

 SetPortScheduleIdNil sets the value for PortScheduleId to be an explicit nil

### UnsetPortScheduleId
`func (o *InlineObject254) UnsetPortScheduleId()`

UnsetPortScheduleId ensures that no value is present for PortScheduleId, not even an explicit nil
### GetUdld

`func (o *InlineObject254) GetUdld() string`

GetUdld returns the Udld field if non-nil, zero value otherwise.

### GetUdldOk

`func (o *InlineObject254) GetUdldOk() (*string, bool)`

GetUdldOk returns a tuple with the Udld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdld

`func (o *InlineObject254) SetUdld(v string)`

SetUdld sets Udld field to given value.

### HasUdld

`func (o *InlineObject254) HasUdld() bool`

HasUdld returns a boolean if a field has been set.

### GetAccessPolicyType

`func (o *InlineObject254) GetAccessPolicyType() string`

GetAccessPolicyType returns the AccessPolicyType field if non-nil, zero value otherwise.

### GetAccessPolicyTypeOk

`func (o *InlineObject254) GetAccessPolicyTypeOk() (*string, bool)`

GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyType

`func (o *InlineObject254) SetAccessPolicyType(v string)`

SetAccessPolicyType sets AccessPolicyType field to given value.

### HasAccessPolicyType

`func (o *InlineObject254) HasAccessPolicyType() bool`

HasAccessPolicyType returns a boolean if a field has been set.

### GetAccessPolicyNumber

`func (o *InlineObject254) GetAccessPolicyNumber() int32`

GetAccessPolicyNumber returns the AccessPolicyNumber field if non-nil, zero value otherwise.

### GetAccessPolicyNumberOk

`func (o *InlineObject254) GetAccessPolicyNumberOk() (*int32, bool)`

GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyNumber

`func (o *InlineObject254) SetAccessPolicyNumber(v int32)`

SetAccessPolicyNumber sets AccessPolicyNumber field to given value.

### HasAccessPolicyNumber

`func (o *InlineObject254) HasAccessPolicyNumber() bool`

HasAccessPolicyNumber returns a boolean if a field has been set.

### GetMacAllowList

`func (o *InlineObject254) GetMacAllowList() []string`

GetMacAllowList returns the MacAllowList field if non-nil, zero value otherwise.

### GetMacAllowListOk

`func (o *InlineObject254) GetMacAllowListOk() (*[]string, bool)`

GetMacAllowListOk returns a tuple with the MacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAllowList

`func (o *InlineObject254) SetMacAllowList(v []string)`

SetMacAllowList sets MacAllowList field to given value.

### HasMacAllowList

`func (o *InlineObject254) HasMacAllowList() bool`

HasMacAllowList returns a boolean if a field has been set.

### GetMacWhitelistLimit

`func (o *InlineObject254) GetMacWhitelistLimit() int32`

GetMacWhitelistLimit returns the MacWhitelistLimit field if non-nil, zero value otherwise.

### GetMacWhitelistLimitOk

`func (o *InlineObject254) GetMacWhitelistLimitOk() (*int32, bool)`

GetMacWhitelistLimitOk returns a tuple with the MacWhitelistLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacWhitelistLimit

`func (o *InlineObject254) SetMacWhitelistLimit(v int32)`

SetMacWhitelistLimit sets MacWhitelistLimit field to given value.

### HasMacWhitelistLimit

`func (o *InlineObject254) HasMacWhitelistLimit() bool`

HasMacWhitelistLimit returns a boolean if a field has been set.

### GetStickyMacAllowList

`func (o *InlineObject254) GetStickyMacAllowList() []string`

GetStickyMacAllowList returns the StickyMacAllowList field if non-nil, zero value otherwise.

### GetStickyMacAllowListOk

`func (o *InlineObject254) GetStickyMacAllowListOk() (*[]string, bool)`

GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowList

`func (o *InlineObject254) SetStickyMacAllowList(v []string)`

SetStickyMacAllowList sets StickyMacAllowList field to given value.

### HasStickyMacAllowList

`func (o *InlineObject254) HasStickyMacAllowList() bool`

HasStickyMacAllowList returns a boolean if a field has been set.

### GetStickyMacAllowListLimit

`func (o *InlineObject254) GetStickyMacAllowListLimit() int32`

GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field if non-nil, zero value otherwise.

### GetStickyMacAllowListLimitOk

`func (o *InlineObject254) GetStickyMacAllowListLimitOk() (*int32, bool)`

GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickyMacAllowListLimit

`func (o *InlineObject254) SetStickyMacAllowListLimit(v int32)`

SetStickyMacAllowListLimit sets StickyMacAllowListLimit field to given value.

### HasStickyMacAllowListLimit

`func (o *InlineObject254) HasStickyMacAllowListLimit() bool`

HasStickyMacAllowListLimit returns a boolean if a field has been set.

### GetStormControlEnabled

`func (o *InlineObject254) GetStormControlEnabled() bool`

GetStormControlEnabled returns the StormControlEnabled field if non-nil, zero value otherwise.

### GetStormControlEnabledOk

`func (o *InlineObject254) GetStormControlEnabledOk() (*bool, bool)`

GetStormControlEnabledOk returns a tuple with the StormControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStormControlEnabled

`func (o *InlineObject254) SetStormControlEnabled(v bool)`

SetStormControlEnabled sets StormControlEnabled field to given value.

### HasStormControlEnabled

`func (o *InlineObject254) HasStormControlEnabled() bool`

HasStormControlEnabled returns a boolean if a field has been set.

### GetFlexibleStackingEnabled

`func (o *InlineObject254) GetFlexibleStackingEnabled() bool`

GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field if non-nil, zero value otherwise.

### GetFlexibleStackingEnabledOk

`func (o *InlineObject254) GetFlexibleStackingEnabledOk() (*bool, bool)`

GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexibleStackingEnabled

`func (o *InlineObject254) SetFlexibleStackingEnabled(v bool)`

SetFlexibleStackingEnabled sets FlexibleStackingEnabled field to given value.

### HasFlexibleStackingEnabled

`func (o *InlineObject254) HasFlexibleStackingEnabled() bool`

HasFlexibleStackingEnabled returns a boolean if a field has been set.

### GetDaiTrusted

`func (o *InlineObject254) GetDaiTrusted() bool`

GetDaiTrusted returns the DaiTrusted field if non-nil, zero value otherwise.

### GetDaiTrustedOk

`func (o *InlineObject254) GetDaiTrustedOk() (*bool, bool)`

GetDaiTrustedOk returns a tuple with the DaiTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaiTrusted

`func (o *InlineObject254) SetDaiTrusted(v bool)`

SetDaiTrusted sets DaiTrusted field to given value.

### HasDaiTrusted

`func (o *InlineObject254) HasDaiTrusted() bool`

HasDaiTrusted returns a boolean if a field has been set.

### GetProfile

`func (o *InlineObject254) GetProfile() DevicesSerialSwitchPortsProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineObject254) GetProfileOk() (*DevicesSerialSwitchPortsProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineObject254) SetProfile(v DevicesSerialSwitchPortsProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineObject254) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetDot3az

`func (o *InlineObject254) GetDot3az() OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az`

GetDot3az returns the Dot3az field if non-nil, zero value otherwise.

### GetDot3azOk

`func (o *InlineObject254) GetDot3azOk() (*OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az, bool)`

GetDot3azOk returns a tuple with the Dot3az field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot3az

`func (o *InlineObject254) SetDot3az(v OrganizationsOrganizationIdConfigTemplatesConfigTemplateIdSwitchProfilesProfileIdPortsDot3az)`

SetDot3az sets Dot3az field to given value.

### HasDot3az

`func (o *InlineObject254) HasDot3az() bool`

HasDot3az returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


