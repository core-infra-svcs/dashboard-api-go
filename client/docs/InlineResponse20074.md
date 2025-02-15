# InlineResponse20074

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse20074Hubs**](InlineResponse20074Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20074Subnets**](InlineResponse20074Subnets.md) | The list of subnets and their VPN presence. | [optional] 

## Methods

### NewInlineResponse20074

`func NewInlineResponse20074() *InlineResponse20074`

NewInlineResponse20074 instantiates a new InlineResponse20074 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074WithDefaults

`func NewInlineResponse20074WithDefaults() *InlineResponse20074`

NewInlineResponse20074WithDefaults instantiates a new InlineResponse20074 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20074) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20074) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20074) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20074) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse20074) GetHubs() []InlineResponse20074Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse20074) GetHubsOk() (*[]InlineResponse20074Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse20074) SetHubs(v []InlineResponse20074Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse20074) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20074) GetSubnets() []InlineResponse20074Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20074) GetSubnetsOk() (*[]InlineResponse20074Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20074) SetSubnets(v []InlineResponse20074Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20074) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


