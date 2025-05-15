# InlineObject72

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveActiveAutoVpnEnabled** | Pointer to **bool** | Toggle for enabling or disabling active-active AutoVPN | [optional] 
**DefaultUplink** | Pointer to **string** | The default uplink. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | [optional] 
**LoadBalancingEnabled** | Pointer to **bool** | Toggle for enabling or disabling load balancing | [optional] 
**FailoverAndFailback** | Pointer to [**NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback.md) |  | [optional] 
**WanTrafficUplinkPreferences** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences.md) | Array of uplink preference rules for WAN traffic | [optional] 
**VpnTrafficUplinkPreferences** | Pointer to [**[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences**](NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences.md) | Array of uplink preference rules for VPN traffic | [optional] 

## Methods

### NewInlineObject72

`func NewInlineObject72() *InlineObject72`

NewInlineObject72 instantiates a new InlineObject72 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject72WithDefaults

`func NewInlineObject72WithDefaults() *InlineObject72`

NewInlineObject72WithDefaults instantiates a new InlineObject72 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveActiveAutoVpnEnabled

`func (o *InlineObject72) GetActiveActiveAutoVpnEnabled() bool`

GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field if non-nil, zero value otherwise.

### GetActiveActiveAutoVpnEnabledOk

`func (o *InlineObject72) GetActiveActiveAutoVpnEnabledOk() (*bool, bool)`

GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveAutoVpnEnabled

`func (o *InlineObject72) SetActiveActiveAutoVpnEnabled(v bool)`

SetActiveActiveAutoVpnEnabled sets ActiveActiveAutoVpnEnabled field to given value.

### HasActiveActiveAutoVpnEnabled

`func (o *InlineObject72) HasActiveActiveAutoVpnEnabled() bool`

HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.

### GetDefaultUplink

`func (o *InlineObject72) GetDefaultUplink() string`

GetDefaultUplink returns the DefaultUplink field if non-nil, zero value otherwise.

### GetDefaultUplinkOk

`func (o *InlineObject72) GetDefaultUplinkOk() (*string, bool)`

GetDefaultUplinkOk returns a tuple with the DefaultUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUplink

`func (o *InlineObject72) SetDefaultUplink(v string)`

SetDefaultUplink sets DefaultUplink field to given value.

### HasDefaultUplink

`func (o *InlineObject72) HasDefaultUplink() bool`

HasDefaultUplink returns a boolean if a field has been set.

### GetLoadBalancingEnabled

`func (o *InlineObject72) GetLoadBalancingEnabled() bool`

GetLoadBalancingEnabled returns the LoadBalancingEnabled field if non-nil, zero value otherwise.

### GetLoadBalancingEnabledOk

`func (o *InlineObject72) GetLoadBalancingEnabledOk() (*bool, bool)`

GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingEnabled

`func (o *InlineObject72) SetLoadBalancingEnabled(v bool)`

SetLoadBalancingEnabled sets LoadBalancingEnabled field to given value.

### HasLoadBalancingEnabled

`func (o *InlineObject72) HasLoadBalancingEnabled() bool`

HasLoadBalancingEnabled returns a boolean if a field has been set.

### GetFailoverAndFailback

`func (o *InlineObject72) GetFailoverAndFailback() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback`

GetFailoverAndFailback returns the FailoverAndFailback field if non-nil, zero value otherwise.

### GetFailoverAndFailbackOk

`func (o *InlineObject72) GetFailoverAndFailbackOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback, bool)`

GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAndFailback

`func (o *InlineObject72) SetFailoverAndFailback(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback)`

SetFailoverAndFailback sets FailoverAndFailback field to given value.

### HasFailoverAndFailback

`func (o *InlineObject72) HasFailoverAndFailback() bool`

HasFailoverAndFailback returns a boolean if a field has been set.

### GetWanTrafficUplinkPreferences

`func (o *InlineObject72) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences`

GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetWanTrafficUplinkPreferencesOk

`func (o *InlineObject72) GetWanTrafficUplinkPreferencesOk() (*[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, bool)`

GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTrafficUplinkPreferences

`func (o *InlineObject72) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences)`

SetWanTrafficUplinkPreferences sets WanTrafficUplinkPreferences field to given value.

### HasWanTrafficUplinkPreferences

`func (o *InlineObject72) HasWanTrafficUplinkPreferences() bool`

HasWanTrafficUplinkPreferences returns a boolean if a field has been set.

### GetVpnTrafficUplinkPreferences

`func (o *InlineObject72) GetVpnTrafficUplinkPreferences() []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences`

GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetVpnTrafficUplinkPreferencesOk

`func (o *InlineObject72) GetVpnTrafficUplinkPreferencesOk() (*[]NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, bool)`

GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTrafficUplinkPreferences

`func (o *InlineObject72) SetVpnTrafficUplinkPreferences(v []NetworksNetworkIdApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences)`

SetVpnTrafficUplinkPreferences sets VpnTrafficUplinkPreferences field to given value.

### HasVpnTrafficUplinkPreferences

`func (o *InlineObject72) HasVpnTrafficUplinkPreferences() bool`

HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


