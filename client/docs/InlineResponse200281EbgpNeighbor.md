# InlineResponse200281EbgpNeighbor

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
**ReceiveLimit** | Pointer to **int32** | Maximum number of prefixes accepted from the remote peer. Must be an integer between 0 and 2147483647. Supported only for MX 26.2 and above | [optional] 
**PathPrepend** | Pointer to **[]int32** | Prepends the AS_PATH BGP Attribute associated with routes received from the remote peer. Configurable value of ASNs to prepend. Length of the array may not exceed 10, and each ASN in the array must be an integer between 1 and 4294967295. AS_PATH is 4th in the decision tree when identical routes from multiple peers exist. | [optional] 
**MultiExitDiscriminator** | Pointer to **int32** | Configures the local metric associated with routes received from the remote peer. Routes from peers with lower metrics are will be preferred. Must be an integer between 0 and 4294967295. MED is 6th in the decision tree when identical routes from multiple peers exist. | [optional] 
**Weight** | Pointer to **int32** | Configures the local metric associated with routes received from the remote peer. Routes from peers with lower metrics are will be preferred. Must be an integer between 0 and 4294967295. MED is 6th in the decision tree when identical routes from multiple peers exist. | [optional] 

## Methods

### NewInlineResponse200281EbgpNeighbor

`func NewInlineResponse200281EbgpNeighbor() *InlineResponse200281EbgpNeighbor`

NewInlineResponse200281EbgpNeighbor instantiates a new InlineResponse200281EbgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200281EbgpNeighborWithDefaults

`func NewInlineResponse200281EbgpNeighborWithDefaults() *InlineResponse200281EbgpNeighbor`

NewInlineResponse200281EbgpNeighborWithDefaults instantiates a new InlineResponse200281EbgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeighborId

`func (o *InlineResponse200281EbgpNeighbor) GetNeighborId() int32`

GetNeighborId returns the NeighborId field if non-nil, zero value otherwise.

### GetNeighborIdOk

`func (o *InlineResponse200281EbgpNeighbor) GetNeighborIdOk() (*int32, bool)`

GetNeighborIdOk returns a tuple with the NeighborId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborId

`func (o *InlineResponse200281EbgpNeighbor) SetNeighborId(v int32)`

SetNeighborId sets NeighborId field to given value.

### HasNeighborId

`func (o *InlineResponse200281EbgpNeighbor) HasNeighborId() bool`

HasNeighborId returns a boolean if a field has been set.

### GetNeighborIp

`func (o *InlineResponse200281EbgpNeighbor) GetNeighborIp() string`

GetNeighborIp returns the NeighborIp field if non-nil, zero value otherwise.

### GetNeighborIpOk

`func (o *InlineResponse200281EbgpNeighbor) GetNeighborIpOk() (*string, bool)`

GetNeighborIpOk returns a tuple with the NeighborIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborIp

`func (o *InlineResponse200281EbgpNeighbor) SetNeighborIp(v string)`

SetNeighborIp sets NeighborIp field to given value.

### HasNeighborIp

`func (o *InlineResponse200281EbgpNeighbor) HasNeighborIp() bool`

HasNeighborIp returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse200281EbgpNeighbor) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse200281EbgpNeighbor) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse200281EbgpNeighbor) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse200281EbgpNeighbor) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *InlineResponse200281EbgpNeighbor) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *InlineResponse200281EbgpNeighbor) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *InlineResponse200281EbgpNeighbor) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.

### HasRemoteAsNumber

`func (o *InlineResponse200281EbgpNeighbor) HasRemoteAsNumber() bool`

HasRemoteAsNumber returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *InlineResponse200281EbgpNeighbor) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *InlineResponse200281EbgpNeighbor) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *InlineResponse200281EbgpNeighbor) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.

### HasEbgpHoldTimer

`func (o *InlineResponse200281EbgpNeighbor) HasEbgpHoldTimer() bool`

HasEbgpHoldTimer returns a boolean if a field has been set.

### GetEbgpMultihop

`func (o *InlineResponse200281EbgpNeighbor) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *InlineResponse200281EbgpNeighbor) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *InlineResponse200281EbgpNeighbor) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.

### HasEbgpMultihop

`func (o *InlineResponse200281EbgpNeighbor) HasEbgpMultihop() bool`

HasEbgpMultihop returns a boolean if a field has been set.

### GetSourceIp

`func (o *InlineResponse200281EbgpNeighbor) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *InlineResponse200281EbgpNeighbor) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *InlineResponse200281EbgpNeighbor) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *InlineResponse200281EbgpNeighbor) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetReceiveLimit

`func (o *InlineResponse200281EbgpNeighbor) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *InlineResponse200281EbgpNeighbor) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *InlineResponse200281EbgpNeighbor) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *InlineResponse200281EbgpNeighbor) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetPathPrepend

`func (o *InlineResponse200281EbgpNeighbor) GetPathPrepend() []int32`

GetPathPrepend returns the PathPrepend field if non-nil, zero value otherwise.

### GetPathPrependOk

`func (o *InlineResponse200281EbgpNeighbor) GetPathPrependOk() (*[]int32, bool)`

GetPathPrependOk returns a tuple with the PathPrepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrepend

`func (o *InlineResponse200281EbgpNeighbor) SetPathPrepend(v []int32)`

SetPathPrepend sets PathPrepend field to given value.

### HasPathPrepend

`func (o *InlineResponse200281EbgpNeighbor) HasPathPrepend() bool`

HasPathPrepend returns a boolean if a field has been set.

### GetMultiExitDiscriminator

`func (o *InlineResponse200281EbgpNeighbor) GetMultiExitDiscriminator() int32`

GetMultiExitDiscriminator returns the MultiExitDiscriminator field if non-nil, zero value otherwise.

### GetMultiExitDiscriminatorOk

`func (o *InlineResponse200281EbgpNeighbor) GetMultiExitDiscriminatorOk() (*int32, bool)`

GetMultiExitDiscriminatorOk returns a tuple with the MultiExitDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiExitDiscriminator

`func (o *InlineResponse200281EbgpNeighbor) SetMultiExitDiscriminator(v int32)`

SetMultiExitDiscriminator sets MultiExitDiscriminator field to given value.

### HasMultiExitDiscriminator

`func (o *InlineResponse200281EbgpNeighbor) HasMultiExitDiscriminator() bool`

HasMultiExitDiscriminator returns a boolean if a field has been set.

### GetWeight

`func (o *InlineResponse200281EbgpNeighbor) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse200281EbgpNeighbor) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse200281EbgpNeighbor) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse200281EbgpNeighbor) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


