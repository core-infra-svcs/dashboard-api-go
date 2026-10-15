# InlineResponse200311EbgpNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeighborId** | Pointer to **int32** | ID of ebgp neighbor | [optional] 
**NeighborIp** | Pointer to **string** | IPv4/IPv6 address of the neighbor | [optional] 
**IpVersion** | Pointer to **int32** | The IP version of the neighbor | [optional] 
**RemoteAsNumber** | Pointer to **int32** | Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295. | [optional] 
**EbgpHoldTimer** | Pointer to **int32** | The eBGP hold timer in seconds for each neighbor. The eBGP hold timer must be an integer between 12 and 240. | [optional] 
**EbgpMultihop** | Pointer to **int32** | Configure this if the neighbor is not adjacent. The eBGP multi-hop must be an integer between 1 and 255. | [optional] 
**SourceIp** | Pointer to **string** | Source IP of eBGP neighbor | [optional] 
**TtlSecurity** | Pointer to [**InlineResponse200311EbgpNeighborTtlSecurity**](InlineResponse200311EbgpNeighborTtlSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**InlineResponse200311EbgpNeighborAuthentication**](InlineResponse200311EbgpNeighborAuthentication.md) |  | [optional] 
**AllowTransit** | Pointer to **bool** | Whether the appliance will advertise routes learned from other Autonomous Systems. Supported only for MX 26.2 and above | [optional] 
**ReceiveLimit** | Pointer to **int32** | Maximum number of prefixes accepted from the remote peer. Must be an integer between 0 and 2147483647. Supported only for MX 26.2 and above | [optional] 
**ExplicitOutboundRouteAdvertisement** | Pointer to [**NullableInlineResponse200311EbgpNeighborExplicitOutboundRouteAdvertisement**](InlineResponse200311EbgpNeighborExplicitOutboundRouteAdvertisement.md) |  | [optional] 
**Filter** | Pointer to [**InlineResponse200311EbgpNeighborFilter**](InlineResponse200311EbgpNeighborFilter.md) |  | [optional] 
**FilterIn** | Pointer to **[]string** | Filters routes received from an eBGP neighbor. Each entry must be an IPv4 CIDR string. Supported only for MX 26.2 and above | [optional] 
**CommunityOut** | Pointer to **[]string** | List of BGP communities tagged to the routes advertised to an eBGP neighbor. Supported only for MX 26.2 and above | [optional] 
**PathPrepend** | Pointer to **[]int32** | Prepends the AS_PATH BGP Attribute associated with routes received from the remote peer. Configurable value of ASNs to prepend. Length of the array may not exceed 10, and each ASN in the array must be an integer between 1 and 4294967295. AS_PATH is 4th in the decision tree when identical routes from multiple peers exist. | [optional] 
**MultiExitDiscriminator** | Pointer to **int32** | Configures the local metric associated with routes received from the remote peer. Routes from peers with lower metrics are will be preferred. Must be an integer between 0 and 4294967295. MED is 6th in the decision tree when identical routes from multiple peers exist. | [optional] 
**Weight** | Pointer to **int32** | Configures the local metric associated with routes received from the remote peer. Routes from peers with lower metrics are will be preferred. Must be an integer between 0 and 4294967295. MED is 6th in the decision tree when identical routes from multiple peers exist. | [optional] 

## Methods

### NewInlineResponse200311EbgpNeighbor

`func NewInlineResponse200311EbgpNeighbor() *InlineResponse200311EbgpNeighbor`

NewInlineResponse200311EbgpNeighbor instantiates a new InlineResponse200311EbgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311EbgpNeighborWithDefaults

`func NewInlineResponse200311EbgpNeighborWithDefaults() *InlineResponse200311EbgpNeighbor`

NewInlineResponse200311EbgpNeighborWithDefaults instantiates a new InlineResponse200311EbgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeighborId

`func (o *InlineResponse200311EbgpNeighbor) GetNeighborId() int32`

GetNeighborId returns the NeighborId field if non-nil, zero value otherwise.

### GetNeighborIdOk

`func (o *InlineResponse200311EbgpNeighbor) GetNeighborIdOk() (*int32, bool)`

GetNeighborIdOk returns a tuple with the NeighborId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborId

`func (o *InlineResponse200311EbgpNeighbor) SetNeighborId(v int32)`

SetNeighborId sets NeighborId field to given value.

### HasNeighborId

`func (o *InlineResponse200311EbgpNeighbor) HasNeighborId() bool`

HasNeighborId returns a boolean if a field has been set.

### GetNeighborIp

`func (o *InlineResponse200311EbgpNeighbor) GetNeighborIp() string`

GetNeighborIp returns the NeighborIp field if non-nil, zero value otherwise.

### GetNeighborIpOk

`func (o *InlineResponse200311EbgpNeighbor) GetNeighborIpOk() (*string, bool)`

GetNeighborIpOk returns a tuple with the NeighborIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborIp

`func (o *InlineResponse200311EbgpNeighbor) SetNeighborIp(v string)`

SetNeighborIp sets NeighborIp field to given value.

### HasNeighborIp

`func (o *InlineResponse200311EbgpNeighbor) HasNeighborIp() bool`

HasNeighborIp returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse200311EbgpNeighbor) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse200311EbgpNeighbor) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse200311EbgpNeighbor) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse200311EbgpNeighbor) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *InlineResponse200311EbgpNeighbor) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *InlineResponse200311EbgpNeighbor) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *InlineResponse200311EbgpNeighbor) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.

### HasRemoteAsNumber

`func (o *InlineResponse200311EbgpNeighbor) HasRemoteAsNumber() bool`

HasRemoteAsNumber returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *InlineResponse200311EbgpNeighbor) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *InlineResponse200311EbgpNeighbor) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *InlineResponse200311EbgpNeighbor) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.

### HasEbgpHoldTimer

`func (o *InlineResponse200311EbgpNeighbor) HasEbgpHoldTimer() bool`

HasEbgpHoldTimer returns a boolean if a field has been set.

### GetEbgpMultihop

`func (o *InlineResponse200311EbgpNeighbor) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *InlineResponse200311EbgpNeighbor) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *InlineResponse200311EbgpNeighbor) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.

### HasEbgpMultihop

`func (o *InlineResponse200311EbgpNeighbor) HasEbgpMultihop() bool`

HasEbgpMultihop returns a boolean if a field has been set.

### GetSourceIp

`func (o *InlineResponse200311EbgpNeighbor) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *InlineResponse200311EbgpNeighbor) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *InlineResponse200311EbgpNeighbor) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *InlineResponse200311EbgpNeighbor) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetTtlSecurity

`func (o *InlineResponse200311EbgpNeighbor) GetTtlSecurity() InlineResponse200311EbgpNeighborTtlSecurity`

GetTtlSecurity returns the TtlSecurity field if non-nil, zero value otherwise.

### GetTtlSecurityOk

`func (o *InlineResponse200311EbgpNeighbor) GetTtlSecurityOk() (*InlineResponse200311EbgpNeighborTtlSecurity, bool)`

GetTtlSecurityOk returns a tuple with the TtlSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecurity

`func (o *InlineResponse200311EbgpNeighbor) SetTtlSecurity(v InlineResponse200311EbgpNeighborTtlSecurity)`

SetTtlSecurity sets TtlSecurity field to given value.

### HasTtlSecurity

`func (o *InlineResponse200311EbgpNeighbor) HasTtlSecurity() bool`

HasTtlSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200311EbgpNeighbor) GetAuthentication() InlineResponse200311EbgpNeighborAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200311EbgpNeighbor) GetAuthenticationOk() (*InlineResponse200311EbgpNeighborAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200311EbgpNeighbor) SetAuthentication(v InlineResponse200311EbgpNeighborAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200311EbgpNeighbor) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAllowTransit

`func (o *InlineResponse200311EbgpNeighbor) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *InlineResponse200311EbgpNeighbor) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *InlineResponse200311EbgpNeighbor) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *InlineResponse200311EbgpNeighbor) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetReceiveLimit

`func (o *InlineResponse200311EbgpNeighbor) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *InlineResponse200311EbgpNeighbor) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *InlineResponse200311EbgpNeighbor) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *InlineResponse200311EbgpNeighbor) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetExplicitOutboundRouteAdvertisement

`func (o *InlineResponse200311EbgpNeighbor) GetExplicitOutboundRouteAdvertisement() InlineResponse200311EbgpNeighborExplicitOutboundRouteAdvertisement`

GetExplicitOutboundRouteAdvertisement returns the ExplicitOutboundRouteAdvertisement field if non-nil, zero value otherwise.

### GetExplicitOutboundRouteAdvertisementOk

`func (o *InlineResponse200311EbgpNeighbor) GetExplicitOutboundRouteAdvertisementOk() (*InlineResponse200311EbgpNeighborExplicitOutboundRouteAdvertisement, bool)`

GetExplicitOutboundRouteAdvertisementOk returns a tuple with the ExplicitOutboundRouteAdvertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitOutboundRouteAdvertisement

`func (o *InlineResponse200311EbgpNeighbor) SetExplicitOutboundRouteAdvertisement(v InlineResponse200311EbgpNeighborExplicitOutboundRouteAdvertisement)`

SetExplicitOutboundRouteAdvertisement sets ExplicitOutboundRouteAdvertisement field to given value.

### HasExplicitOutboundRouteAdvertisement

`func (o *InlineResponse200311EbgpNeighbor) HasExplicitOutboundRouteAdvertisement() bool`

HasExplicitOutboundRouteAdvertisement returns a boolean if a field has been set.

### SetExplicitOutboundRouteAdvertisementNil

`func (o *InlineResponse200311EbgpNeighbor) SetExplicitOutboundRouteAdvertisementNil(b bool)`

 SetExplicitOutboundRouteAdvertisementNil sets the value for ExplicitOutboundRouteAdvertisement to be an explicit nil

### UnsetExplicitOutboundRouteAdvertisement
`func (o *InlineResponse200311EbgpNeighbor) UnsetExplicitOutboundRouteAdvertisement()`

UnsetExplicitOutboundRouteAdvertisement ensures that no value is present for ExplicitOutboundRouteAdvertisement, not even an explicit nil
### GetFilter

`func (o *InlineResponse200311EbgpNeighbor) GetFilter() InlineResponse200311EbgpNeighborFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InlineResponse200311EbgpNeighbor) GetFilterOk() (*InlineResponse200311EbgpNeighborFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InlineResponse200311EbgpNeighbor) SetFilter(v InlineResponse200311EbgpNeighborFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InlineResponse200311EbgpNeighbor) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterIn

`func (o *InlineResponse200311EbgpNeighbor) GetFilterIn() []string`

GetFilterIn returns the FilterIn field if non-nil, zero value otherwise.

### GetFilterInOk

`func (o *InlineResponse200311EbgpNeighbor) GetFilterInOk() (*[]string, bool)`

GetFilterInOk returns a tuple with the FilterIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIn

`func (o *InlineResponse200311EbgpNeighbor) SetFilterIn(v []string)`

SetFilterIn sets FilterIn field to given value.

### HasFilterIn

`func (o *InlineResponse200311EbgpNeighbor) HasFilterIn() bool`

HasFilterIn returns a boolean if a field has been set.

### GetCommunityOut

`func (o *InlineResponse200311EbgpNeighbor) GetCommunityOut() []string`

GetCommunityOut returns the CommunityOut field if non-nil, zero value otherwise.

### GetCommunityOutOk

`func (o *InlineResponse200311EbgpNeighbor) GetCommunityOutOk() (*[]string, bool)`

GetCommunityOutOk returns a tuple with the CommunityOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityOut

`func (o *InlineResponse200311EbgpNeighbor) SetCommunityOut(v []string)`

SetCommunityOut sets CommunityOut field to given value.

### HasCommunityOut

`func (o *InlineResponse200311EbgpNeighbor) HasCommunityOut() bool`

HasCommunityOut returns a boolean if a field has been set.

### GetPathPrepend

`func (o *InlineResponse200311EbgpNeighbor) GetPathPrepend() []int32`

GetPathPrepend returns the PathPrepend field if non-nil, zero value otherwise.

### GetPathPrependOk

`func (o *InlineResponse200311EbgpNeighbor) GetPathPrependOk() (*[]int32, bool)`

GetPathPrependOk returns a tuple with the PathPrepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrepend

`func (o *InlineResponse200311EbgpNeighbor) SetPathPrepend(v []int32)`

SetPathPrepend sets PathPrepend field to given value.

### HasPathPrepend

`func (o *InlineResponse200311EbgpNeighbor) HasPathPrepend() bool`

HasPathPrepend returns a boolean if a field has been set.

### GetMultiExitDiscriminator

`func (o *InlineResponse200311EbgpNeighbor) GetMultiExitDiscriminator() int32`

GetMultiExitDiscriminator returns the MultiExitDiscriminator field if non-nil, zero value otherwise.

### GetMultiExitDiscriminatorOk

`func (o *InlineResponse200311EbgpNeighbor) GetMultiExitDiscriminatorOk() (*int32, bool)`

GetMultiExitDiscriminatorOk returns a tuple with the MultiExitDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiExitDiscriminator

`func (o *InlineResponse200311EbgpNeighbor) SetMultiExitDiscriminator(v int32)`

SetMultiExitDiscriminator sets MultiExitDiscriminator field to given value.

### HasMultiExitDiscriminator

`func (o *InlineResponse200311EbgpNeighbor) HasMultiExitDiscriminator() bool`

HasMultiExitDiscriminator returns a boolean if a field has been set.

### GetWeight

`func (o *InlineResponse200311EbgpNeighbor) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse200311EbgpNeighbor) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse200311EbgpNeighbor) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse200311EbgpNeighbor) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


