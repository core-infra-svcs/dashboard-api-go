# InlineObject111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for your group policy. Required. | 
**Scheduling** | Pointer to [**NetworksNetworkIdGroupPoliciesScheduling**](NetworksNetworkIdGroupPoliciesScheduling.md) |  | [optional] 
**Bandwidth** | Pointer to [**NetworksNetworkIdGroupPoliciesBandwidth**](NetworksNetworkIdGroupPoliciesBandwidth.md) |  | [optional] 
**FirewallAndTrafficShaping** | Pointer to [**NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping**](NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping.md) |  | [optional] 
**ContentFiltering** | Pointer to [**NetworksNetworkIdGroupPoliciesContentFiltering**](NetworksNetworkIdGroupPoliciesContentFiltering.md) |  | [optional] 
**SplashAuthSettings** | Pointer to **string** | Whether clients bound to your policy will bypass splash authorization or behave according to the network&#39;s rules. Can be one of &#39;network default&#39; or &#39;bypass&#39;. Only available if your network has a wireless configuration. | [optional] 
**VlanTagging** | Pointer to [**NetworksNetworkIdGroupPoliciesVlanTagging**](NetworksNetworkIdGroupPoliciesVlanTagging.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**NetworksNetworkIdGroupPoliciesBonjourForwarding**](NetworksNetworkIdGroupPoliciesBonjourForwarding.md) |  | [optional] 

## Methods

### NewInlineObject111

`func NewInlineObject111(name string, ) *InlineObject111`

NewInlineObject111 instantiates a new InlineObject111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject111WithDefaults

`func NewInlineObject111WithDefaults() *InlineObject111`

NewInlineObject111WithDefaults instantiates a new InlineObject111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject111) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject111) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject111) SetName(v string)`

SetName sets Name field to given value.


### GetScheduling

`func (o *InlineObject111) GetScheduling() NetworksNetworkIdGroupPoliciesScheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *InlineObject111) GetSchedulingOk() (*NetworksNetworkIdGroupPoliciesScheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *InlineObject111) SetScheduling(v NetworksNetworkIdGroupPoliciesScheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *InlineObject111) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetBandwidth

`func (o *InlineObject111) GetBandwidth() NetworksNetworkIdGroupPoliciesBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InlineObject111) GetBandwidthOk() (*NetworksNetworkIdGroupPoliciesBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InlineObject111) SetBandwidth(v NetworksNetworkIdGroupPoliciesBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InlineObject111) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetFirewallAndTrafficShaping

`func (o *InlineObject111) GetFirewallAndTrafficShaping() NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping`

GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field if non-nil, zero value otherwise.

### GetFirewallAndTrafficShapingOk

`func (o *InlineObject111) GetFirewallAndTrafficShapingOk() (*NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping, bool)`

GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallAndTrafficShaping

`func (o *InlineObject111) SetFirewallAndTrafficShaping(v NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping)`

SetFirewallAndTrafficShaping sets FirewallAndTrafficShaping field to given value.

### HasFirewallAndTrafficShaping

`func (o *InlineObject111) HasFirewallAndTrafficShaping() bool`

HasFirewallAndTrafficShaping returns a boolean if a field has been set.

### GetContentFiltering

`func (o *InlineObject111) GetContentFiltering() NetworksNetworkIdGroupPoliciesContentFiltering`

GetContentFiltering returns the ContentFiltering field if non-nil, zero value otherwise.

### GetContentFilteringOk

`func (o *InlineObject111) GetContentFilteringOk() (*NetworksNetworkIdGroupPoliciesContentFiltering, bool)`

GetContentFilteringOk returns a tuple with the ContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFiltering

`func (o *InlineObject111) SetContentFiltering(v NetworksNetworkIdGroupPoliciesContentFiltering)`

SetContentFiltering sets ContentFiltering field to given value.

### HasContentFiltering

`func (o *InlineObject111) HasContentFiltering() bool`

HasContentFiltering returns a boolean if a field has been set.

### GetSplashAuthSettings

`func (o *InlineObject111) GetSplashAuthSettings() string`

GetSplashAuthSettings returns the SplashAuthSettings field if non-nil, zero value otherwise.

### GetSplashAuthSettingsOk

`func (o *InlineObject111) GetSplashAuthSettingsOk() (*string, bool)`

GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashAuthSettings

`func (o *InlineObject111) SetSplashAuthSettings(v string)`

SetSplashAuthSettings sets SplashAuthSettings field to given value.

### HasSplashAuthSettings

`func (o *InlineObject111) HasSplashAuthSettings() bool`

HasSplashAuthSettings returns a boolean if a field has been set.

### GetVlanTagging

`func (o *InlineObject111) GetVlanTagging() NetworksNetworkIdGroupPoliciesVlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *InlineObject111) GetVlanTaggingOk() (*NetworksNetworkIdGroupPoliciesVlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *InlineObject111) SetVlanTagging(v NetworksNetworkIdGroupPoliciesVlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *InlineObject111) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineObject111) GetBonjourForwarding() NetworksNetworkIdGroupPoliciesBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineObject111) GetBonjourForwardingOk() (*NetworksNetworkIdGroupPoliciesBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineObject111) SetBonjourForwarding(v NetworksNetworkIdGroupPoliciesBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineObject111) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


