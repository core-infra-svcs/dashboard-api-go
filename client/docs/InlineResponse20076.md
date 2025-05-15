# InlineResponse20076

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20076Hubs**](InlineResponse20076Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20076Subnets**](InlineResponse20076Subnets.md) | The list of subnets and their VPN presence. | [optional] 
**Subnet** | Pointer to [**InlineResponse20076Subnet**](InlineResponse20076Subnet.md) |  | [optional] 

## Methods

### NewInlineResponse20076

`func NewInlineResponse20076() *InlineResponse20076`

NewInlineResponse20076 instantiates a new InlineResponse20076 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20076WithDefaults

`func NewInlineResponse20076WithDefaults() *InlineResponse20076`

NewInlineResponse20076WithDefaults instantiates a new InlineResponse20076 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20076) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20076) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20076) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20076) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20076) GetHubs() []InlineResponse20076Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20076) GetHubsOk() (*[]InlineResponse20076Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20076) SetHubs(v []InlineResponse20076Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20076) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20076) GetSubnets() []InlineResponse20076Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20076) GetSubnetsOk() (*[]InlineResponse20076Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20076) SetSubnets(v []InlineResponse20076Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20076) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20076) GetSubnet() InlineResponse20076Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20076) GetSubnetOk() (*InlineResponse20076Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20076) SetSubnet(v InlineResponse20076Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20076) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


