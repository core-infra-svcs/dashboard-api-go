# InlineResponse200291

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse200106Uplinks**](InlineResponse200106Uplinks.md) | Per-uplink NAT exception configuration on the network. | [optional] 

## Methods

### NewInlineResponse200291

`func NewInlineResponse200291() *InlineResponse200291`

NewInlineResponse200291 instantiates a new InlineResponse200291 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200291WithDefaults

`func NewInlineResponse200291WithDefaults() *InlineResponse200291`

NewInlineResponse200291WithDefaults instantiates a new InlineResponse200291 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200291) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200291) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200291) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200291) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200291) GetUplinks() []InlineResponse200106Uplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200291) GetUplinksOk() (*[]InlineResponse200106Uplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200291) SetUplinks(v []InlineResponse200106Uplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200291) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


