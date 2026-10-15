# InlineResponse200110Neighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IPv4 address of the neighbor | [optional] 
**Ipv6** | Pointer to [**InlineResponse200110Ipv61**](InlineResponse200110Ipv61.md) |  | [optional] 
**RemoteAsNumber** | Pointer to **int32** | Remote AS number of the neighbor | [optional] 
**ReceiveLimit** | Pointer to **int32** | The maximum number of routes that the appliance can receive from the neighbor | [optional] 
**AllowTransit** | Pointer to **bool** | Whether the appliance will advertise routes learned from other Autonomous Systems | [optional] 
**EbgpHoldTimer** | Pointer to **int32** | The eBGP hold time in seconds for the neighbor | [optional] 
**EbgpMultihop** | Pointer to **int32** | The number of hops the appliance must traverse to establish a peering relationship with the neighbor | [optional] 
**SourceInterface** | Pointer to **string** | The output interface the appliance uses to establish a peering relationship with the neighbor | [optional] 
**NextHopIp** | Pointer to **string** | The IPv4 address of the neighbor that will establish a TCP session with the appliance | [optional] 
**TtlSecurity** | Pointer to [**InlineResponse200110TtlSecurity**](InlineResponse200110TtlSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**InlineResponse200110Authentication**](InlineResponse200110Authentication.md) |  | [optional] 
**MultiExitDiscriminator** | Pointer to **int32** | Multi-exit Discriminator (MED) metric associated with routes received from the neighbor | [optional] 
**PathPrepend** | Pointer to **[]int32** | Values to prepend to the AS_PATH BGP Attribute associated with routes received from the neighbor | [optional] 
**Weight** | Pointer to **int32** | Local weight for routes received from the neighbor | [optional] 
**FilterIn** | Pointer to **[]string** | Filters routes received from an eBGP neighbor. Each entry must be an IPv4 CIDR string (e.g., &#39;10.0.0.0/8&#39; or &#39;10.0.0.5/32&#39;). | [optional] 
**CommunityOut** | Pointer to **[]string** | List of BGP communities tagged to the routes advertised to an eBGP neighbor. | [optional] 
**Filter** | Pointer to [**InlineResponse200110Filter**](InlineResponse200110Filter.md) |  | [optional] 
**ExplicitOutboundRouteAdvertisement** | Pointer to [**NullableInlineResponse200110ExplicitOutboundRouteAdvertisement**](InlineResponse200110ExplicitOutboundRouteAdvertisement.md) |  | [optional] 
**Vrf** | Pointer to [**InlineResponse200110Vrf**](InlineResponse200110Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse200110Neighbors

`func NewInlineResponse200110Neighbors() *InlineResponse200110Neighbors`

NewInlineResponse200110Neighbors instantiates a new InlineResponse200110Neighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110NeighborsWithDefaults

`func NewInlineResponse200110NeighborsWithDefaults() *InlineResponse200110Neighbors`

NewInlineResponse200110NeighborsWithDefaults instantiates a new InlineResponse200110Neighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse200110Neighbors) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse200110Neighbors) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse200110Neighbors) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse200110Neighbors) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse200110Neighbors) GetIpv6() InlineResponse200110Ipv61`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse200110Neighbors) GetIpv6Ok() (*InlineResponse200110Ipv61, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse200110Neighbors) SetIpv6(v InlineResponse200110Ipv61)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse200110Neighbors) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *InlineResponse200110Neighbors) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *InlineResponse200110Neighbors) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *InlineResponse200110Neighbors) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.

### HasRemoteAsNumber

`func (o *InlineResponse200110Neighbors) HasRemoteAsNumber() bool`

HasRemoteAsNumber returns a boolean if a field has been set.

### GetReceiveLimit

`func (o *InlineResponse200110Neighbors) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *InlineResponse200110Neighbors) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *InlineResponse200110Neighbors) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *InlineResponse200110Neighbors) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetAllowTransit

`func (o *InlineResponse200110Neighbors) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *InlineResponse200110Neighbors) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *InlineResponse200110Neighbors) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *InlineResponse200110Neighbors) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *InlineResponse200110Neighbors) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *InlineResponse200110Neighbors) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *InlineResponse200110Neighbors) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.

### HasEbgpHoldTimer

`func (o *InlineResponse200110Neighbors) HasEbgpHoldTimer() bool`

HasEbgpHoldTimer returns a boolean if a field has been set.

### GetEbgpMultihop

`func (o *InlineResponse200110Neighbors) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *InlineResponse200110Neighbors) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *InlineResponse200110Neighbors) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.

### HasEbgpMultihop

`func (o *InlineResponse200110Neighbors) HasEbgpMultihop() bool`

HasEbgpMultihop returns a boolean if a field has been set.

### GetSourceInterface

`func (o *InlineResponse200110Neighbors) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *InlineResponse200110Neighbors) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *InlineResponse200110Neighbors) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *InlineResponse200110Neighbors) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetNextHopIp

`func (o *InlineResponse200110Neighbors) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineResponse200110Neighbors) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineResponse200110Neighbors) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *InlineResponse200110Neighbors) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetTtlSecurity

`func (o *InlineResponse200110Neighbors) GetTtlSecurity() InlineResponse200110TtlSecurity`

GetTtlSecurity returns the TtlSecurity field if non-nil, zero value otherwise.

### GetTtlSecurityOk

`func (o *InlineResponse200110Neighbors) GetTtlSecurityOk() (*InlineResponse200110TtlSecurity, bool)`

GetTtlSecurityOk returns a tuple with the TtlSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecurity

`func (o *InlineResponse200110Neighbors) SetTtlSecurity(v InlineResponse200110TtlSecurity)`

SetTtlSecurity sets TtlSecurity field to given value.

### HasTtlSecurity

`func (o *InlineResponse200110Neighbors) HasTtlSecurity() bool`

HasTtlSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200110Neighbors) GetAuthentication() InlineResponse200110Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200110Neighbors) GetAuthenticationOk() (*InlineResponse200110Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200110Neighbors) SetAuthentication(v InlineResponse200110Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200110Neighbors) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetMultiExitDiscriminator

`func (o *InlineResponse200110Neighbors) GetMultiExitDiscriminator() int32`

GetMultiExitDiscriminator returns the MultiExitDiscriminator field if non-nil, zero value otherwise.

### GetMultiExitDiscriminatorOk

`func (o *InlineResponse200110Neighbors) GetMultiExitDiscriminatorOk() (*int32, bool)`

GetMultiExitDiscriminatorOk returns a tuple with the MultiExitDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiExitDiscriminator

`func (o *InlineResponse200110Neighbors) SetMultiExitDiscriminator(v int32)`

SetMultiExitDiscriminator sets MultiExitDiscriminator field to given value.

### HasMultiExitDiscriminator

`func (o *InlineResponse200110Neighbors) HasMultiExitDiscriminator() bool`

HasMultiExitDiscriminator returns a boolean if a field has been set.

### GetPathPrepend

`func (o *InlineResponse200110Neighbors) GetPathPrepend() []int32`

GetPathPrepend returns the PathPrepend field if non-nil, zero value otherwise.

### GetPathPrependOk

`func (o *InlineResponse200110Neighbors) GetPathPrependOk() (*[]int32, bool)`

GetPathPrependOk returns a tuple with the PathPrepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrepend

`func (o *InlineResponse200110Neighbors) SetPathPrepend(v []int32)`

SetPathPrepend sets PathPrepend field to given value.

### HasPathPrepend

`func (o *InlineResponse200110Neighbors) HasPathPrepend() bool`

HasPathPrepend returns a boolean if a field has been set.

### GetWeight

`func (o *InlineResponse200110Neighbors) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse200110Neighbors) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse200110Neighbors) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse200110Neighbors) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetFilterIn

`func (o *InlineResponse200110Neighbors) GetFilterIn() []string`

GetFilterIn returns the FilterIn field if non-nil, zero value otherwise.

### GetFilterInOk

`func (o *InlineResponse200110Neighbors) GetFilterInOk() (*[]string, bool)`

GetFilterInOk returns a tuple with the FilterIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIn

`func (o *InlineResponse200110Neighbors) SetFilterIn(v []string)`

SetFilterIn sets FilterIn field to given value.

### HasFilterIn

`func (o *InlineResponse200110Neighbors) HasFilterIn() bool`

HasFilterIn returns a boolean if a field has been set.

### GetCommunityOut

`func (o *InlineResponse200110Neighbors) GetCommunityOut() []string`

GetCommunityOut returns the CommunityOut field if non-nil, zero value otherwise.

### GetCommunityOutOk

`func (o *InlineResponse200110Neighbors) GetCommunityOutOk() (*[]string, bool)`

GetCommunityOutOk returns a tuple with the CommunityOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityOut

`func (o *InlineResponse200110Neighbors) SetCommunityOut(v []string)`

SetCommunityOut sets CommunityOut field to given value.

### HasCommunityOut

`func (o *InlineResponse200110Neighbors) HasCommunityOut() bool`

HasCommunityOut returns a boolean if a field has been set.

### GetFilter

`func (o *InlineResponse200110Neighbors) GetFilter() InlineResponse200110Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InlineResponse200110Neighbors) GetFilterOk() (*InlineResponse200110Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InlineResponse200110Neighbors) SetFilter(v InlineResponse200110Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InlineResponse200110Neighbors) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetExplicitOutboundRouteAdvertisement

`func (o *InlineResponse200110Neighbors) GetExplicitOutboundRouteAdvertisement() InlineResponse200110ExplicitOutboundRouteAdvertisement`

GetExplicitOutboundRouteAdvertisement returns the ExplicitOutboundRouteAdvertisement field if non-nil, zero value otherwise.

### GetExplicitOutboundRouteAdvertisementOk

`func (o *InlineResponse200110Neighbors) GetExplicitOutboundRouteAdvertisementOk() (*InlineResponse200110ExplicitOutboundRouteAdvertisement, bool)`

GetExplicitOutboundRouteAdvertisementOk returns a tuple with the ExplicitOutboundRouteAdvertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitOutboundRouteAdvertisement

`func (o *InlineResponse200110Neighbors) SetExplicitOutboundRouteAdvertisement(v InlineResponse200110ExplicitOutboundRouteAdvertisement)`

SetExplicitOutboundRouteAdvertisement sets ExplicitOutboundRouteAdvertisement field to given value.

### HasExplicitOutboundRouteAdvertisement

`func (o *InlineResponse200110Neighbors) HasExplicitOutboundRouteAdvertisement() bool`

HasExplicitOutboundRouteAdvertisement returns a boolean if a field has been set.

### SetExplicitOutboundRouteAdvertisementNil

`func (o *InlineResponse200110Neighbors) SetExplicitOutboundRouteAdvertisementNil(b bool)`

 SetExplicitOutboundRouteAdvertisementNil sets the value for ExplicitOutboundRouteAdvertisement to be an explicit nil

### UnsetExplicitOutboundRouteAdvertisement
`func (o *InlineResponse200110Neighbors) UnsetExplicitOutboundRouteAdvertisement()`

UnsetExplicitOutboundRouteAdvertisement ensures that no value is present for ExplicitOutboundRouteAdvertisement, not even an explicit nil
### GetVrf

`func (o *InlineResponse200110Neighbors) GetVrf() InlineResponse200110Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200110Neighbors) GetVrfOk() (*InlineResponse200110Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200110Neighbors) SetVrf(v InlineResponse200110Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200110Neighbors) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


