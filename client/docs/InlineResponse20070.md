# InlineResponse20070

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20070Hubs**](InlineResponse20070Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20070Subnets**](InlineResponse20070Subnets.md) | The list of subnets and their VPN presence. | [optional] 

## Methods

### NewInlineResponse20070

`func NewInlineResponse20070() *InlineResponse20070`

NewInlineResponse20070 instantiates a new InlineResponse20070 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20070WithDefaults

`func NewInlineResponse20070WithDefaults() *InlineResponse20070`

NewInlineResponse20070WithDefaults instantiates a new InlineResponse20070 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20070) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20070) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20070) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20070) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20070) GetHubs() []InlineResponse20070Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20070) GetHubsOk() (*[]InlineResponse20070Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20070) SetHubs(v []InlineResponse20070Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20070) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20070) GetSubnets() []InlineResponse20070Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20070) GetSubnetsOk() (*[]InlineResponse20070Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20070) SetSubnets(v []InlineResponse20070Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20070) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


