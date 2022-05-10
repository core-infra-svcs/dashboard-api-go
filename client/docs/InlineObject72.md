# InlineObject72

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

### NewInlineObject72

`func NewInlineObject72(name string, ) *InlineObject72`

NewInlineObject72 instantiates a new InlineObject72 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject72WithDefaults

`func NewInlineObject72WithDefaults() *InlineObject72`

NewInlineObject72WithDefaults instantiates a new InlineObject72 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject72) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject72) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject72) SetName(v string)`

SetName sets Name field to given value.


### GetScheduling

`func (o *InlineObject72) GetScheduling() NetworksNetworkIdGroupPoliciesScheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *InlineObject72) GetSchedulingOk() (*NetworksNetworkIdGroupPoliciesScheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *InlineObject72) SetScheduling(v NetworksNetworkIdGroupPoliciesScheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *InlineObject72) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetBandwidth

`func (o *InlineObject72) GetBandwidth() NetworksNetworkIdGroupPoliciesBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InlineObject72) GetBandwidthOk() (*NetworksNetworkIdGroupPoliciesBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InlineObject72) SetBandwidth(v NetworksNetworkIdGroupPoliciesBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InlineObject72) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetFirewallAndTrafficShaping

`func (o *InlineObject72) GetFirewallAndTrafficShaping() NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping`

GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field if non-nil, zero value otherwise.

### GetFirewallAndTrafficShapingOk

`func (o *InlineObject72) GetFirewallAndTrafficShapingOk() (*NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping, bool)`

GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallAndTrafficShaping

`func (o *InlineObject72) SetFirewallAndTrafficShaping(v NetworksNetworkIdGroupPoliciesFirewallAndTrafficShaping)`

SetFirewallAndTrafficShaping sets FirewallAndTrafficShaping field to given value.

### HasFirewallAndTrafficShaping

`func (o *InlineObject72) HasFirewallAndTrafficShaping() bool`

HasFirewallAndTrafficShaping returns a boolean if a field has been set.

### GetContentFiltering

`func (o *InlineObject72) GetContentFiltering() NetworksNetworkIdGroupPoliciesContentFiltering`

GetContentFiltering returns the ContentFiltering field if non-nil, zero value otherwise.

### GetContentFilteringOk

`func (o *InlineObject72) GetContentFilteringOk() (*NetworksNetworkIdGroupPoliciesContentFiltering, bool)`

GetContentFilteringOk returns a tuple with the ContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFiltering

`func (o *InlineObject72) SetContentFiltering(v NetworksNetworkIdGroupPoliciesContentFiltering)`

SetContentFiltering sets ContentFiltering field to given value.

### HasContentFiltering

`func (o *InlineObject72) HasContentFiltering() bool`

HasContentFiltering returns a boolean if a field has been set.

### GetSplashAuthSettings

`func (o *InlineObject72) GetSplashAuthSettings() string`

GetSplashAuthSettings returns the SplashAuthSettings field if non-nil, zero value otherwise.

### GetSplashAuthSettingsOk

`func (o *InlineObject72) GetSplashAuthSettingsOk() (*string, bool)`

GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashAuthSettings

`func (o *InlineObject72) SetSplashAuthSettings(v string)`

SetSplashAuthSettings sets SplashAuthSettings field to given value.

### HasSplashAuthSettings

`func (o *InlineObject72) HasSplashAuthSettings() bool`

HasSplashAuthSettings returns a boolean if a field has been set.

### GetVlanTagging

`func (o *InlineObject72) GetVlanTagging() NetworksNetworkIdGroupPoliciesVlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *InlineObject72) GetVlanTaggingOk() (*NetworksNetworkIdGroupPoliciesVlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *InlineObject72) SetVlanTagging(v NetworksNetworkIdGroupPoliciesVlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *InlineObject72) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineObject72) GetBonjourForwarding() NetworksNetworkIdGroupPoliciesBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineObject72) GetBonjourForwardingOk() (*NetworksNetworkIdGroupPoliciesBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineObject72) SetBonjourForwarding(v NetworksNetworkIdGroupPoliciesBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineObject72) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


