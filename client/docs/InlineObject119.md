# InlineObject119

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for your group policy. | [optional] 
**Scheduling** | Pointer to [**NetworksNetworkIdGroupPoliciesScheduling**](NetworksNetworkIdGroupPoliciesScheduling.md) |  | [optional] 
**Bandwidth** | Pointer to [**NetworksNetworkIdGroupPoliciesBandwidth**](NetworksNetworkIdGroupPoliciesBandwidth.md) |  | [optional] 
**FirewallAndTrafficShaping** | Pointer to [**NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping**](NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping.md) |  | [optional] 
**ContentFiltering** | Pointer to [**NetworksNetworkIdGroupPoliciesContentFiltering**](NetworksNetworkIdGroupPoliciesContentFiltering.md) |  | [optional] 
**SplashAuthSettings** | Pointer to **string** | Whether clients bound to your policy will bypass splash authorization or behave according to the network&#39;s rules. Can be one of &#39;network default&#39; or &#39;bypass&#39;. Only available if your network has a wireless configuration. | [optional] 
**VlanTagging** | Pointer to [**NetworksNetworkIdGroupPoliciesVlanTagging**](NetworksNetworkIdGroupPoliciesVlanTagging.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**NetworksNetworkIdGroupPoliciesBonjourForwarding**](NetworksNetworkIdGroupPoliciesBonjourForwarding.md) |  | [optional] 

## Methods

### NewInlineObject119

`func NewInlineObject119() *InlineObject119`

NewInlineObject119 instantiates a new InlineObject119 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject119WithDefaults

`func NewInlineObject119WithDefaults() *InlineObject119`

NewInlineObject119WithDefaults instantiates a new InlineObject119 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject119) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject119) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject119) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject119) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduling

`func (o *InlineObject119) GetScheduling() NetworksNetworkIdGroupPoliciesScheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *InlineObject119) GetSchedulingOk() (*NetworksNetworkIdGroupPoliciesScheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *InlineObject119) SetScheduling(v NetworksNetworkIdGroupPoliciesScheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *InlineObject119) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetBandwidth

`func (o *InlineObject119) GetBandwidth() NetworksNetworkIdGroupPoliciesBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InlineObject119) GetBandwidthOk() (*NetworksNetworkIdGroupPoliciesBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InlineObject119) SetBandwidth(v NetworksNetworkIdGroupPoliciesBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InlineObject119) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetFirewallAndTrafficShaping

`func (o *InlineObject119) GetFirewallAndTrafficShaping() NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping`

GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field if non-nil, zero value otherwise.

### GetFirewallAndTrafficShapingOk

`func (o *InlineObject119) GetFirewallAndTrafficShapingOk() (*NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping, bool)`

GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallAndTrafficShaping

`func (o *InlineObject119) SetFirewallAndTrafficShaping(v NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping)`

SetFirewallAndTrafficShaping sets FirewallAndTrafficShaping field to given value.

### HasFirewallAndTrafficShaping

`func (o *InlineObject119) HasFirewallAndTrafficShaping() bool`

HasFirewallAndTrafficShaping returns a boolean if a field has been set.

### GetContentFiltering

`func (o *InlineObject119) GetContentFiltering() NetworksNetworkIdGroupPoliciesContentFiltering`

GetContentFiltering returns the ContentFiltering field if non-nil, zero value otherwise.

### GetContentFilteringOk

`func (o *InlineObject119) GetContentFilteringOk() (*NetworksNetworkIdGroupPoliciesContentFiltering, bool)`

GetContentFilteringOk returns a tuple with the ContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFiltering

`func (o *InlineObject119) SetContentFiltering(v NetworksNetworkIdGroupPoliciesContentFiltering)`

SetContentFiltering sets ContentFiltering field to given value.

### HasContentFiltering

`func (o *InlineObject119) HasContentFiltering() bool`

HasContentFiltering returns a boolean if a field has been set.

### GetSplashAuthSettings

`func (o *InlineObject119) GetSplashAuthSettings() string`

GetSplashAuthSettings returns the SplashAuthSettings field if non-nil, zero value otherwise.

### GetSplashAuthSettingsOk

`func (o *InlineObject119) GetSplashAuthSettingsOk() (*string, bool)`

GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashAuthSettings

`func (o *InlineObject119) SetSplashAuthSettings(v string)`

SetSplashAuthSettings sets SplashAuthSettings field to given value.

### HasSplashAuthSettings

`func (o *InlineObject119) HasSplashAuthSettings() bool`

HasSplashAuthSettings returns a boolean if a field has been set.

### GetVlanTagging

`func (o *InlineObject119) GetVlanTagging() NetworksNetworkIdGroupPoliciesVlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *InlineObject119) GetVlanTaggingOk() (*NetworksNetworkIdGroupPoliciesVlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *InlineObject119) SetVlanTagging(v NetworksNetworkIdGroupPoliciesVlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *InlineObject119) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineObject119) GetBonjourForwarding() NetworksNetworkIdGroupPoliciesBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineObject119) GetBonjourForwardingOk() (*NetworksNetworkIdGroupPoliciesBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineObject119) SetBonjourForwarding(v NetworksNetworkIdGroupPoliciesBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineObject119) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


