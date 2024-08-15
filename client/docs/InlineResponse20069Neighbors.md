# InlineResponse20069Neighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IPv4 address of the neighbor | [optional] 
**Ipv6** | Pointer to [**InlineResponse20069Ipv6**](InlineResponse20069Ipv6.md) |  | [optional] 
**RemoteAsNumber** | Pointer to **int32** | Remote AS number of the neighbor | [optional] 
**ReceiveLimit** | Pointer to **int32** | The maximum number of routes that the appliance can receive from the neighbor | [optional] 
**AllowTransit** | Pointer to **bool** | Whether the appliance will advertise routes learned from other Autonomous Systems | [optional] 
**EbgpHoldTimer** | Pointer to **int32** | The eBGP hold time in seconds for the neighbor | [optional] 
**EbgpMultihop** | Pointer to **int32** | The number of hops the appliance must traverse to establish a peering relationship with the neighbor | [optional] 
**SourceInterface** | Pointer to **string** | The output interface the appliance uses to establish a peering relationship with the neighbor | [optional] 
**NextHopIp** | Pointer to **string** | The IPv4 address of the neighbor that will establish a TCP session with the appliance | [optional] 
**TtlSecurity** | Pointer to [**InlineResponse20069TtlSecurity**](InlineResponse20069TtlSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**InlineResponse20069Authentication**](InlineResponse20069Authentication.md) |  | [optional] 

## Methods

### NewInlineResponse20069Neighbors

`func NewInlineResponse20069Neighbors() *InlineResponse20069Neighbors`

NewInlineResponse20069Neighbors instantiates a new InlineResponse20069Neighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20069NeighborsWithDefaults

`func NewInlineResponse20069NeighborsWithDefaults() *InlineResponse20069Neighbors`

NewInlineResponse20069NeighborsWithDefaults instantiates a new InlineResponse20069Neighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse20069Neighbors) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20069Neighbors) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20069Neighbors) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20069Neighbors) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20069Neighbors) GetIpv6() InlineResponse20069Ipv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20069Neighbors) GetIpv6Ok() (*InlineResponse20069Ipv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20069Neighbors) SetIpv6(v InlineResponse20069Ipv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20069Neighbors) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *InlineResponse20069Neighbors) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *InlineResponse20069Neighbors) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *InlineResponse20069Neighbors) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.

### HasRemoteAsNumber

`func (o *InlineResponse20069Neighbors) HasRemoteAsNumber() bool`

HasRemoteAsNumber returns a boolean if a field has been set.

### GetReceiveLimit

`func (o *InlineResponse20069Neighbors) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *InlineResponse20069Neighbors) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *InlineResponse20069Neighbors) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *InlineResponse20069Neighbors) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetAllowTransit

`func (o *InlineResponse20069Neighbors) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *InlineResponse20069Neighbors) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *InlineResponse20069Neighbors) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *InlineResponse20069Neighbors) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *InlineResponse20069Neighbors) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *InlineResponse20069Neighbors) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *InlineResponse20069Neighbors) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.

### HasEbgpHoldTimer

`func (o *InlineResponse20069Neighbors) HasEbgpHoldTimer() bool`

HasEbgpHoldTimer returns a boolean if a field has been set.

### GetEbgpMultihop

`func (o *InlineResponse20069Neighbors) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *InlineResponse20069Neighbors) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *InlineResponse20069Neighbors) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.

### HasEbgpMultihop

`func (o *InlineResponse20069Neighbors) HasEbgpMultihop() bool`

HasEbgpMultihop returns a boolean if a field has been set.

### GetSourceInterface

`func (o *InlineResponse20069Neighbors) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *InlineResponse20069Neighbors) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *InlineResponse20069Neighbors) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *InlineResponse20069Neighbors) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetNextHopIp

`func (o *InlineResponse20069Neighbors) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineResponse20069Neighbors) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineResponse20069Neighbors) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *InlineResponse20069Neighbors) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetTtlSecurity

`func (o *InlineResponse20069Neighbors) GetTtlSecurity() InlineResponse20069TtlSecurity`

GetTtlSecurity returns the TtlSecurity field if non-nil, zero value otherwise.

### GetTtlSecurityOk

`func (o *InlineResponse20069Neighbors) GetTtlSecurityOk() (*InlineResponse20069TtlSecurity, bool)`

GetTtlSecurityOk returns a tuple with the TtlSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecurity

`func (o *InlineResponse20069Neighbors) SetTtlSecurity(v InlineResponse20069TtlSecurity)`

SetTtlSecurity sets TtlSecurity field to given value.

### HasTtlSecurity

`func (o *InlineResponse20069Neighbors) HasTtlSecurity() bool`

HasTtlSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse20069Neighbors) GetAuthentication() InlineResponse20069Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse20069Neighbors) GetAuthenticationOk() (*InlineResponse20069Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse20069Neighbors) SetAuthentication(v InlineResponse20069Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse20069Neighbors) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


