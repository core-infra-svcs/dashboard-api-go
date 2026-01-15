# InlineResponse20079Neighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IPv4 address of the neighbor | [optional] 
**Ipv6** | Pointer to [**InlineResponse20079Ipv6**](InlineResponse20079Ipv6.md) |  | [optional] 
**RemoteAsNumber** | Pointer to **int32** | Remote AS number of the neighbor | [optional] 
**ReceiveLimit** | Pointer to **int32** | The maximum number of routes that the appliance can receive from the neighbor | [optional] 
**AllowTransit** | Pointer to **bool** | Whether the appliance will advertise routes learned from other Autonomous Systems | [optional] 
**EbgpHoldTimer** | Pointer to **int32** | The eBGP hold time in seconds for the neighbor | [optional] 
**EbgpMultihop** | Pointer to **int32** | The number of hops the appliance must traverse to establish a peering relationship with the neighbor | [optional] 
**SourceInterface** | Pointer to **string** | The output interface the appliance uses to establish a peering relationship with the neighbor | [optional] 
**NextHopIp** | Pointer to **string** | The IPv4 address of the neighbor that will establish a TCP session with the appliance | [optional] 
**TtlSecurity** | Pointer to [**InlineResponse20079TtlSecurity**](InlineResponse20079TtlSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**InlineResponse20079Authentication**](InlineResponse20079Authentication.md) |  | [optional] 
**MultiExitDiscriminator** | Pointer to **int32** | Multi-exit Discriminator (MED) metric associated with routes received from the neighbor | [optional] 
**PathPrepend** | Pointer to **[]int32** | Values to prepend to the AS_PATH BGP Attribute associated with routes received from the neighbor | [optional] 
**Weight** | Pointer to **int32** | Local weight for routes received from the neighbor | [optional] 

## Methods

### NewInlineResponse20079Neighbors

`func NewInlineResponse20079Neighbors() *InlineResponse20079Neighbors`

NewInlineResponse20079Neighbors instantiates a new InlineResponse20079Neighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079NeighborsWithDefaults

`func NewInlineResponse20079NeighborsWithDefaults() *InlineResponse20079Neighbors`

NewInlineResponse20079NeighborsWithDefaults instantiates a new InlineResponse20079Neighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse20079Neighbors) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20079Neighbors) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20079Neighbors) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20079Neighbors) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20079Neighbors) GetIpv6() InlineResponse20079Ipv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20079Neighbors) GetIpv6Ok() (*InlineResponse20079Ipv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20079Neighbors) SetIpv6(v InlineResponse20079Ipv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20079Neighbors) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *InlineResponse20079Neighbors) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *InlineResponse20079Neighbors) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *InlineResponse20079Neighbors) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.

### HasRemoteAsNumber

`func (o *InlineResponse20079Neighbors) HasRemoteAsNumber() bool`

HasRemoteAsNumber returns a boolean if a field has been set.

### GetReceiveLimit

`func (o *InlineResponse20079Neighbors) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *InlineResponse20079Neighbors) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *InlineResponse20079Neighbors) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *InlineResponse20079Neighbors) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetAllowTransit

`func (o *InlineResponse20079Neighbors) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *InlineResponse20079Neighbors) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *InlineResponse20079Neighbors) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *InlineResponse20079Neighbors) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *InlineResponse20079Neighbors) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *InlineResponse20079Neighbors) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *InlineResponse20079Neighbors) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.

### HasEbgpHoldTimer

`func (o *InlineResponse20079Neighbors) HasEbgpHoldTimer() bool`

HasEbgpHoldTimer returns a boolean if a field has been set.

### GetEbgpMultihop

`func (o *InlineResponse20079Neighbors) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *InlineResponse20079Neighbors) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *InlineResponse20079Neighbors) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.

### HasEbgpMultihop

`func (o *InlineResponse20079Neighbors) HasEbgpMultihop() bool`

HasEbgpMultihop returns a boolean if a field has been set.

### GetSourceInterface

`func (o *InlineResponse20079Neighbors) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *InlineResponse20079Neighbors) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *InlineResponse20079Neighbors) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *InlineResponse20079Neighbors) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetNextHopIp

`func (o *InlineResponse20079Neighbors) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineResponse20079Neighbors) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineResponse20079Neighbors) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *InlineResponse20079Neighbors) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetTtlSecurity

`func (o *InlineResponse20079Neighbors) GetTtlSecurity() InlineResponse20079TtlSecurity`

GetTtlSecurity returns the TtlSecurity field if non-nil, zero value otherwise.

### GetTtlSecurityOk

`func (o *InlineResponse20079Neighbors) GetTtlSecurityOk() (*InlineResponse20079TtlSecurity, bool)`

GetTtlSecurityOk returns a tuple with the TtlSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecurity

`func (o *InlineResponse20079Neighbors) SetTtlSecurity(v InlineResponse20079TtlSecurity)`

SetTtlSecurity sets TtlSecurity field to given value.

### HasTtlSecurity

`func (o *InlineResponse20079Neighbors) HasTtlSecurity() bool`

HasTtlSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse20079Neighbors) GetAuthentication() InlineResponse20079Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse20079Neighbors) GetAuthenticationOk() (*InlineResponse20079Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse20079Neighbors) SetAuthentication(v InlineResponse20079Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse20079Neighbors) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetMultiExitDiscriminator

`func (o *InlineResponse20079Neighbors) GetMultiExitDiscriminator() int32`

GetMultiExitDiscriminator returns the MultiExitDiscriminator field if non-nil, zero value otherwise.

### GetMultiExitDiscriminatorOk

`func (o *InlineResponse20079Neighbors) GetMultiExitDiscriminatorOk() (*int32, bool)`

GetMultiExitDiscriminatorOk returns a tuple with the MultiExitDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiExitDiscriminator

`func (o *InlineResponse20079Neighbors) SetMultiExitDiscriminator(v int32)`

SetMultiExitDiscriminator sets MultiExitDiscriminator field to given value.

### HasMultiExitDiscriminator

`func (o *InlineResponse20079Neighbors) HasMultiExitDiscriminator() bool`

HasMultiExitDiscriminator returns a boolean if a field has been set.

### GetPathPrepend

`func (o *InlineResponse20079Neighbors) GetPathPrepend() []int32`

GetPathPrepend returns the PathPrepend field if non-nil, zero value otherwise.

### GetPathPrependOk

`func (o *InlineResponse20079Neighbors) GetPathPrependOk() (*[]int32, bool)`

GetPathPrependOk returns a tuple with the PathPrepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrepend

`func (o *InlineResponse20079Neighbors) SetPathPrepend(v []int32)`

SetPathPrepend sets PathPrepend field to given value.

### HasPathPrepend

`func (o *InlineResponse20079Neighbors) HasPathPrepend() bool`

HasPathPrepend returns a boolean if a field has been set.

### GetWeight

`func (o *InlineResponse20079Neighbors) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse20079Neighbors) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse20079Neighbors) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse20079Neighbors) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


