# InlineResponse20065

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveActiveAutoVpnEnabled** | Pointer to **bool** | Whether active-active AutoVPN is enabled | [optional] 
**DefaultUplink** | Pointer to **string** | The default uplink. Must be one of: &#39;wan1&#39; or &#39;wan2&#39; | [optional] 
**LoadBalancingEnabled** | Pointer to **bool** | Whether load balancing is enabled | [optional] 
**FailoverAndFailback** | Pointer to [**InlineResponse20065FailoverAndFailback**](InlineResponse20065FailoverAndFailback.md) |  | [optional] 
**WanTrafficUplinkPreferences** | Pointer to [**[]InlineResponse20065WanTrafficUplinkPreferences**](InlineResponse20065WanTrafficUplinkPreferences.md) | Uplink preference rules for WAN traffic | [optional] 
**VpnTrafficUplinkPreferences** | Pointer to [**[]InlineResponse20065VpnTrafficUplinkPreferences**](InlineResponse20065VpnTrafficUplinkPreferences.md) | Uplink preference rules for VPN traffic | [optional] 

## Methods

### NewInlineResponse20065

`func NewInlineResponse20065() *InlineResponse20065`

NewInlineResponse20065 instantiates a new InlineResponse20065 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20065WithDefaults

`func NewInlineResponse20065WithDefaults() *InlineResponse20065`

NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveActiveAutoVpnEnabled

`func (o *InlineResponse20065) GetActiveActiveAutoVpnEnabled() bool`

GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field if non-nil, zero value otherwise.

### GetActiveActiveAutoVpnEnabledOk

`func (o *InlineResponse20065) GetActiveActiveAutoVpnEnabledOk() (*bool, bool)`

GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveAutoVpnEnabled

`func (o *InlineResponse20065) SetActiveActiveAutoVpnEnabled(v bool)`

SetActiveActiveAutoVpnEnabled sets ActiveActiveAutoVpnEnabled field to given value.

### HasActiveActiveAutoVpnEnabled

`func (o *InlineResponse20065) HasActiveActiveAutoVpnEnabled() bool`

HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.

### GetDefaultUplink

`func (o *InlineResponse20065) GetDefaultUplink() string`

GetDefaultUplink returns the DefaultUplink field if non-nil, zero value otherwise.

### GetDefaultUplinkOk

`func (o *InlineResponse20065) GetDefaultUplinkOk() (*string, bool)`

GetDefaultUplinkOk returns a tuple with the DefaultUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUplink

`func (o *InlineResponse20065) SetDefaultUplink(v string)`

SetDefaultUplink sets DefaultUplink field to given value.

### HasDefaultUplink

`func (o *InlineResponse20065) HasDefaultUplink() bool`

HasDefaultUplink returns a boolean if a field has been set.

### GetLoadBalancingEnabled

`func (o *InlineResponse20065) GetLoadBalancingEnabled() bool`

GetLoadBalancingEnabled returns the LoadBalancingEnabled field if non-nil, zero value otherwise.

### GetLoadBalancingEnabledOk

`func (o *InlineResponse20065) GetLoadBalancingEnabledOk() (*bool, bool)`

GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingEnabled

`func (o *InlineResponse20065) SetLoadBalancingEnabled(v bool)`

SetLoadBalancingEnabled sets LoadBalancingEnabled field to given value.

### HasLoadBalancingEnabled

`func (o *InlineResponse20065) HasLoadBalancingEnabled() bool`

HasLoadBalancingEnabled returns a boolean if a field has been set.

### GetFailoverAndFailback

`func (o *InlineResponse20065) GetFailoverAndFailback() InlineResponse20065FailoverAndFailback`

GetFailoverAndFailback returns the FailoverAndFailback field if non-nil, zero value otherwise.

### GetFailoverAndFailbackOk

`func (o *InlineResponse20065) GetFailoverAndFailbackOk() (*InlineResponse20065FailoverAndFailback, bool)`

GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAndFailback

`func (o *InlineResponse20065) SetFailoverAndFailback(v InlineResponse20065FailoverAndFailback)`

SetFailoverAndFailback sets FailoverAndFailback field to given value.

### HasFailoverAndFailback

`func (o *InlineResponse20065) HasFailoverAndFailback() bool`

HasFailoverAndFailback returns a boolean if a field has been set.

### GetWanTrafficUplinkPreferences

`func (o *InlineResponse20065) GetWanTrafficUplinkPreferences() []InlineResponse20065WanTrafficUplinkPreferences`

GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetWanTrafficUplinkPreferencesOk

`func (o *InlineResponse20065) GetWanTrafficUplinkPreferencesOk() (*[]InlineResponse20065WanTrafficUplinkPreferences, bool)`

GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTrafficUplinkPreferences

`func (o *InlineResponse20065) SetWanTrafficUplinkPreferences(v []InlineResponse20065WanTrafficUplinkPreferences)`

SetWanTrafficUplinkPreferences sets WanTrafficUplinkPreferences field to given value.

### HasWanTrafficUplinkPreferences

`func (o *InlineResponse20065) HasWanTrafficUplinkPreferences() bool`

HasWanTrafficUplinkPreferences returns a boolean if a field has been set.

### GetVpnTrafficUplinkPreferences

`func (o *InlineResponse20065) GetVpnTrafficUplinkPreferences() []InlineResponse20065VpnTrafficUplinkPreferences`

GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field if non-nil, zero value otherwise.

### GetVpnTrafficUplinkPreferencesOk

`func (o *InlineResponse20065) GetVpnTrafficUplinkPreferencesOk() (*[]InlineResponse20065VpnTrafficUplinkPreferences, bool)`

GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTrafficUplinkPreferences

`func (o *InlineResponse20065) SetVpnTrafficUplinkPreferences(v []InlineResponse20065VpnTrafficUplinkPreferences)`

SetVpnTrafficUplinkPreferences sets VpnTrafficUplinkPreferences field to given value.

### HasVpnTrafficUplinkPreferences

`func (o *InlineResponse20065) HasVpnTrafficUplinkPreferences() bool`

HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


