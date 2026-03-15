# InlineResponse200255EcmpUplinkConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the ECMP uplink configuration | [optional] 
**Wan** | Pointer to **string** | The WAN uplink associated with this ECMP configuration. | [optional] 
**PrivateSubnets** | Pointer to **[]string** | The private subnets associated with this ECMP uplink configuration. | [optional] 
**EbgpNeighbor** | Pointer to [**InlineResponse200255EbgpNeighbor1**](InlineResponse200255EbgpNeighbor1.md) |  | [optional] 

## Methods

### NewInlineResponse200255EcmpUplinkConfigs

`func NewInlineResponse200255EcmpUplinkConfigs() *InlineResponse200255EcmpUplinkConfigs`

NewInlineResponse200255EcmpUplinkConfigs instantiates a new InlineResponse200255EcmpUplinkConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200255EcmpUplinkConfigsWithDefaults

`func NewInlineResponse200255EcmpUplinkConfigsWithDefaults() *InlineResponse200255EcmpUplinkConfigs`

NewInlineResponse200255EcmpUplinkConfigsWithDefaults instantiates a new InlineResponse200255EcmpUplinkConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200255EcmpUplinkConfigs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200255EcmpUplinkConfigs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200255EcmpUplinkConfigs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200255EcmpUplinkConfigs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWan

`func (o *InlineResponse200255EcmpUplinkConfigs) GetWan() string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *InlineResponse200255EcmpUplinkConfigs) GetWanOk() (*string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *InlineResponse200255EcmpUplinkConfigs) SetWan(v string)`

SetWan sets Wan field to given value.

### HasWan

`func (o *InlineResponse200255EcmpUplinkConfigs) HasWan() bool`

HasWan returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *InlineResponse200255EcmpUplinkConfigs) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *InlineResponse200255EcmpUplinkConfigs) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *InlineResponse200255EcmpUplinkConfigs) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *InlineResponse200255EcmpUplinkConfigs) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetEbgpNeighbor

`func (o *InlineResponse200255EcmpUplinkConfigs) GetEbgpNeighbor() InlineResponse200255EbgpNeighbor1`

GetEbgpNeighbor returns the EbgpNeighbor field if non-nil, zero value otherwise.

### GetEbgpNeighborOk

`func (o *InlineResponse200255EcmpUplinkConfigs) GetEbgpNeighborOk() (*InlineResponse200255EbgpNeighbor1, bool)`

GetEbgpNeighborOk returns a tuple with the EbgpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpNeighbor

`func (o *InlineResponse200255EcmpUplinkConfigs) SetEbgpNeighbor(v InlineResponse200255EbgpNeighbor1)`

SetEbgpNeighbor sets EbgpNeighbor field to given value.

### HasEbgpNeighbor

`func (o *InlineResponse200255EcmpUplinkConfigs) HasEbgpNeighbor() bool`

HasEbgpNeighbor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


