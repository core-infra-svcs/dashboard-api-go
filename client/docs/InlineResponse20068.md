# InlineResponse20068

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20068Hubs**](InlineResponse20068Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20068Subnets**](InlineResponse20068Subnets.md) | The list of subnets and their VPN presence. | [optional] 

## Methods

### NewInlineResponse20068

`func NewInlineResponse20068() *InlineResponse20068`

NewInlineResponse20068 instantiates a new InlineResponse20068 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20068WithDefaults

`func NewInlineResponse20068WithDefaults() *InlineResponse20068`

NewInlineResponse20068WithDefaults instantiates a new InlineResponse20068 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20068) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20068) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20068) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20068) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20068) GetHubs() []InlineResponse20068Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20068) GetHubsOk() (*[]InlineResponse20068Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20068) SetHubs(v []InlineResponse20068Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20068) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20068) GetSubnets() []InlineResponse20068Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20068) GetSubnetsOk() (*[]InlineResponse20068Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20068) SetSubnets(v []InlineResponse20068Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20068) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


