# InlineResponse20080

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20080Hubs**](InlineResponse20080Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20080Subnets**](InlineResponse20080Subnets.md) | The list of subnets and their VPN presence. | [optional] 
**Subnet** | Pointer to [**InlineResponse20080Subnet**](InlineResponse20080Subnet.md) |  | [optional] 

## Methods

### NewInlineResponse20080

`func NewInlineResponse20080() *InlineResponse20080`

NewInlineResponse20080 instantiates a new InlineResponse20080 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20080WithDefaults

`func NewInlineResponse20080WithDefaults() *InlineResponse20080`

NewInlineResponse20080WithDefaults instantiates a new InlineResponse20080 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20080) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20080) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20080) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20080) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20080) GetHubs() []InlineResponse20080Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20080) GetHubsOk() (*[]InlineResponse20080Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20080) SetHubs(v []InlineResponse20080Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20080) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20080) GetSubnets() []InlineResponse20080Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20080) GetSubnetsOk() (*[]InlineResponse20080Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20080) SetSubnets(v []InlineResponse20080Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20080) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20080) GetSubnet() InlineResponse20080Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20080) GetSubnetOk() (*InlineResponse20080Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20080) SetSubnet(v InlineResponse20080Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20080) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


