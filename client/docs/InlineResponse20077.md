# InlineResponse20077

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20077Hubs**](InlineResponse20077Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20077Subnets**](InlineResponse20077Subnets.md) | The list of subnets and their VPN presence. | [optional] 
**Subnet** | Pointer to [**InlineResponse20077Subnet**](InlineResponse20077Subnet.md) |  | [optional] 

## Methods

### NewInlineResponse20077

`func NewInlineResponse20077() *InlineResponse20077`

NewInlineResponse20077 instantiates a new InlineResponse20077 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20077WithDefaults

`func NewInlineResponse20077WithDefaults() *InlineResponse20077`

NewInlineResponse20077WithDefaults instantiates a new InlineResponse20077 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20077) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20077) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20077) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20077) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20077) GetHubs() []InlineResponse20077Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20077) GetHubsOk() (*[]InlineResponse20077Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20077) SetHubs(v []InlineResponse20077Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20077) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20077) GetSubnets() []InlineResponse20077Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20077) GetSubnetsOk() (*[]InlineResponse20077Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20077) SetSubnets(v []InlineResponse20077Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20077) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20077) GetSubnet() InlineResponse20077Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20077) GetSubnetOk() (*InlineResponse20077Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20077) SetSubnet(v InlineResponse20077Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20077) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


