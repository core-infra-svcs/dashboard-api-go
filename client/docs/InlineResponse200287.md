# InlineResponse200287

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200105Uplinks**](InlineResponse200105Uplinks.md) | Per-uplink NAT exception configuration on the network. | [optional] 

## Methods

### NewInlineResponse200287

`func NewInlineResponse200287() *InlineResponse200287`

NewInlineResponse200287 instantiates a new InlineResponse200287 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200287WithDefaults

`func NewInlineResponse200287WithDefaults() *InlineResponse200287`

NewInlineResponse200287WithDefaults instantiates a new InlineResponse200287 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200287) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200287) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200287) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200287) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200287) GetUplinks() []InlineResponse200105Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200287) GetUplinksOk() (*[]InlineResponse200105Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200287) SetUplinks(v []InlineResponse200105Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200287) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


