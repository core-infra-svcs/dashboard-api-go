# InlineResponse200110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether BGP is enabled on the appliance | [optional] 
**AsNumber** | Pointer to **int32** | Organization-specific number of the Autonomous System to which the appliance belongs | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The iBGP hold time in seconds | [optional] 
**Ipv6** | Pointer to [**InlineResponse200110Ipv6**](InlineResponse200110Ipv6.md) |  | [optional] 
**TunnelDownTermination** | Pointer to [**InlineResponse200110TunnelDownTermination**](InlineResponse200110TunnelDownTermination.md) |  | [optional] 
**LocalAsNumber** | Pointer to **NullableInt32** | Network-specific local Autonomous System Number (ASN) to which the appliance belongs. | [optional] 
**PriorityRoute** | Pointer to **NullableString** | Sets the priority route between eBGP and Auto VPN. | [optional] 
**RouterId** | Pointer to **string** | The router ID of the appliance | [optional] 
**Neighbors** | Pointer to [**[]InlineResponse200110Neighbors**](InlineResponse200110Neighbors.md) | List of eBGP neighbor configurations | [optional] 

## Methods

### NewInlineResponse200110

`func NewInlineResponse200110() *InlineResponse200110`

NewInlineResponse200110 instantiates a new InlineResponse200110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110WithDefaults

`func NewInlineResponse200110WithDefaults() *InlineResponse200110`

NewInlineResponse200110WithDefaults instantiates a new InlineResponse200110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200110) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200110) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200110) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200110) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAsNumber

`func (o *InlineResponse200110) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *InlineResponse200110) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *InlineResponse200110) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *InlineResponse200110) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *InlineResponse200110) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *InlineResponse200110) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *InlineResponse200110) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *InlineResponse200110) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse200110) GetIpv6() InlineResponse200110Ipv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse200110) GetIpv6Ok() (*InlineResponse200110Ipv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse200110) SetIpv6(v InlineResponse200110Ipv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse200110) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetTunnelDownTermination

`func (o *InlineResponse200110) GetTunnelDownTermination() InlineResponse200110TunnelDownTermination`

GetTunnelDownTermination returns the TunnelDownTermination field if non-nil, zero value otherwise.

### GetTunnelDownTerminationOk

`func (o *InlineResponse200110) GetTunnelDownTerminationOk() (*InlineResponse200110TunnelDownTermination, bool)`

GetTunnelDownTerminationOk returns a tuple with the TunnelDownTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelDownTermination

`func (o *InlineResponse200110) SetTunnelDownTermination(v InlineResponse200110TunnelDownTermination)`

SetTunnelDownTermination sets TunnelDownTermination field to given value.

### HasTunnelDownTermination

`func (o *InlineResponse200110) HasTunnelDownTermination() bool`

HasTunnelDownTermination returns a boolean if a field has been set.

### GetLocalAsNumber

`func (o *InlineResponse200110) GetLocalAsNumber() int32`

GetLocalAsNumber returns the LocalAsNumber field if non-nil, zero value otherwise.

### GetLocalAsNumberOk

`func (o *InlineResponse200110) GetLocalAsNumberOk() (*int32, bool)`

GetLocalAsNumberOk returns a tuple with the LocalAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsNumber

`func (o *InlineResponse200110) SetLocalAsNumber(v int32)`

SetLocalAsNumber sets LocalAsNumber field to given value.

### HasLocalAsNumber

`func (o *InlineResponse200110) HasLocalAsNumber() bool`

HasLocalAsNumber returns a boolean if a field has been set.

### SetLocalAsNumberNil

`func (o *InlineResponse200110) SetLocalAsNumberNil(b bool)`

 SetLocalAsNumberNil sets the value for LocalAsNumber to be an explicit nil

### UnsetLocalAsNumber
`func (o *InlineResponse200110) UnsetLocalAsNumber()`

UnsetLocalAsNumber ensures that no value is present for LocalAsNumber, not even an explicit nil
### GetPriorityRoute

`func (o *InlineResponse200110) GetPriorityRoute() string`

GetPriorityRoute returns the PriorityRoute field if non-nil, zero value otherwise.

### GetPriorityRouteOk

`func (o *InlineResponse200110) GetPriorityRouteOk() (*string, bool)`

GetPriorityRouteOk returns a tuple with the PriorityRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityRoute

`func (o *InlineResponse200110) SetPriorityRoute(v string)`

SetPriorityRoute sets PriorityRoute field to given value.

### HasPriorityRoute

`func (o *InlineResponse200110) HasPriorityRoute() bool`

HasPriorityRoute returns a boolean if a field has been set.

### SetPriorityRouteNil

`func (o *InlineResponse200110) SetPriorityRouteNil(b bool)`

 SetPriorityRouteNil sets the value for PriorityRoute to be an explicit nil

### UnsetPriorityRoute
`func (o *InlineResponse200110) UnsetPriorityRoute()`

UnsetPriorityRoute ensures that no value is present for PriorityRoute, not even an explicit nil
### GetRouterId

`func (o *InlineResponse200110) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *InlineResponse200110) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *InlineResponse200110) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *InlineResponse200110) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse200110) GetNeighbors() []InlineResponse200110Neighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse200110) GetNeighborsOk() (*[]InlineResponse200110Neighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse200110) SetNeighbors(v []InlineResponse200110Neighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse200110) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


