# InlineResponse200113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether BGP is enabled on the appliance | [optional] 
**AsNumber** | Pointer to **int32** | Organization-specific number of the Autonomous System to which the appliance belongs | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The iBGP hold time in seconds | [optional] 
**Ipv6** | Pointer to [**InlineResponse200113Ipv6**](InlineResponse200113Ipv6.md) |  | [optional] 
**TunnelDownTermination** | Pointer to [**InlineResponse200113TunnelDownTermination**](InlineResponse200113TunnelDownTermination.md) |  | [optional] 
**LocalAsNumber** | Pointer to **NullableInt32** | Network-specific local Autonomous System Number (ASN) to which the appliance belongs. | [optional] 
**PriorityRoute** | Pointer to **NullableString** | Sets the priority route between eBGP and Auto VPN. | [optional] 
**RouterId** | Pointer to **string** | The router ID of the appliance | [optional] 
**Neighbors** | Pointer to [**[]InlineResponse200113Neighbors**](InlineResponse200113Neighbors.md) | List of eBGP neighbor configurations | [optional] 

## Methods

### NewInlineResponse200113

`func NewInlineResponse200113() *InlineResponse200113`

NewInlineResponse200113 instantiates a new InlineResponse200113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200113WithDefaults

`func NewInlineResponse200113WithDefaults() *InlineResponse200113`

NewInlineResponse200113WithDefaults instantiates a new InlineResponse200113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200113) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200113) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200113) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200113) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAsNumber

`func (o *InlineResponse200113) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *InlineResponse200113) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *InlineResponse200113) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *InlineResponse200113) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *InlineResponse200113) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *InlineResponse200113) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *InlineResponse200113) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *InlineResponse200113) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse200113) GetIpv6() InlineResponse200113Ipv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse200113) GetIpv6Ok() (*InlineResponse200113Ipv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse200113) SetIpv6(v InlineResponse200113Ipv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse200113) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetTunnelDownTermination

`func (o *InlineResponse200113) GetTunnelDownTermination() InlineResponse200113TunnelDownTermination`

GetTunnelDownTermination returns the TunnelDownTermination field if non-nil, zero value otherwise.

### GetTunnelDownTerminationOk

`func (o *InlineResponse200113) GetTunnelDownTerminationOk() (*InlineResponse200113TunnelDownTermination, bool)`

GetTunnelDownTerminationOk returns a tuple with the TunnelDownTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelDownTermination

`func (o *InlineResponse200113) SetTunnelDownTermination(v InlineResponse200113TunnelDownTermination)`

SetTunnelDownTermination sets TunnelDownTermination field to given value.

### HasTunnelDownTermination

`func (o *InlineResponse200113) HasTunnelDownTermination() bool`

HasTunnelDownTermination returns a boolean if a field has been set.

### GetLocalAsNumber

`func (o *InlineResponse200113) GetLocalAsNumber() int32`

GetLocalAsNumber returns the LocalAsNumber field if non-nil, zero value otherwise.

### GetLocalAsNumberOk

`func (o *InlineResponse200113) GetLocalAsNumberOk() (*int32, bool)`

GetLocalAsNumberOk returns a tuple with the LocalAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsNumber

`func (o *InlineResponse200113) SetLocalAsNumber(v int32)`

SetLocalAsNumber sets LocalAsNumber field to given value.

### HasLocalAsNumber

`func (o *InlineResponse200113) HasLocalAsNumber() bool`

HasLocalAsNumber returns a boolean if a field has been set.

### SetLocalAsNumberNil

`func (o *InlineResponse200113) SetLocalAsNumberNil(b bool)`

 SetLocalAsNumberNil sets the value for LocalAsNumber to be an explicit nil

### UnsetLocalAsNumber
`func (o *InlineResponse200113) UnsetLocalAsNumber()`

UnsetLocalAsNumber ensures that no value is present for LocalAsNumber, not even an explicit nil
### GetPriorityRoute

`func (o *InlineResponse200113) GetPriorityRoute() string`

GetPriorityRoute returns the PriorityRoute field if non-nil, zero value otherwise.

### GetPriorityRouteOk

`func (o *InlineResponse200113) GetPriorityRouteOk() (*string, bool)`

GetPriorityRouteOk returns a tuple with the PriorityRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityRoute

`func (o *InlineResponse200113) SetPriorityRoute(v string)`

SetPriorityRoute sets PriorityRoute field to given value.

### HasPriorityRoute

`func (o *InlineResponse200113) HasPriorityRoute() bool`

HasPriorityRoute returns a boolean if a field has been set.

### SetPriorityRouteNil

`func (o *InlineResponse200113) SetPriorityRouteNil(b bool)`

 SetPriorityRouteNil sets the value for PriorityRoute to be an explicit nil

### UnsetPriorityRoute
`func (o *InlineResponse200113) UnsetPriorityRoute()`

UnsetPriorityRoute ensures that no value is present for PriorityRoute, not even an explicit nil
### GetRouterId

`func (o *InlineResponse200113) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *InlineResponse200113) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *InlineResponse200113) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *InlineResponse200113) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse200113) GetNeighbors() []InlineResponse200113Neighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse200113) GetNeighborsOk() (*[]InlineResponse200113Neighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse200113) SetNeighbors(v []InlineResponse200113Neighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse200113) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


