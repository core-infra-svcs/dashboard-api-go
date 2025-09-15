# InlineResponse20079

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20079Hubs**](InlineResponse20079Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20079Subnets**](InlineResponse20079Subnets.md) | The list of subnets and their VPN presence. | [optional] 
**Subnet** | Pointer to [**InlineResponse20079Subnet**](InlineResponse20079Subnet.md) |  | [optional] 

## Methods

### NewInlineResponse20079

`func NewInlineResponse20079() *InlineResponse20079`

NewInlineResponse20079 instantiates a new InlineResponse20079 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079WithDefaults

`func NewInlineResponse20079WithDefaults() *InlineResponse20079`

NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20079) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20079) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20079) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20079) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20079) GetHubs() []InlineResponse20079Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20079) GetHubsOk() (*[]InlineResponse20079Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20079) SetHubs(v []InlineResponse20079Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20079) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20079) GetSubnets() []InlineResponse20079Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20079) GetSubnetsOk() (*[]InlineResponse20079Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20079) SetSubnets(v []InlineResponse20079Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20079) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20079) GetSubnet() InlineResponse20079Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20079) GetSubnetOk() (*InlineResponse20079Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20079) SetSubnet(v InlineResponse20079Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20079) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


