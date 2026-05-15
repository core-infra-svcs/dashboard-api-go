# InlineResponse200102

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse200102Hubs**](InlineResponse200102Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse200102Subnets**](InlineResponse200102Subnets.md) | The list of subnets and their VPN presence. | [optional] 
**Subnet** | Pointer to [**InlineResponse200102Subnet**](InlineResponse200102Subnet.md) |  | [optional] 
**HostTranslations** | Pointer to [**[]InlineResponse200102HostTranslations**](InlineResponse200102HostTranslations.md) | The list of VPN host translations. Host translations are supported starting from MX firmware version 26.1.2 | [optional] 

## Methods

### NewInlineResponse200102

`func NewInlineResponse200102() *InlineResponse200102`

NewInlineResponse200102 instantiates a new InlineResponse200102 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102WithDefaults

`func NewInlineResponse200102WithDefaults() *InlineResponse200102`

NewInlineResponse200102WithDefaults instantiates a new InlineResponse200102 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200102) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200102) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200102) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200102) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse200102) GetHubs() []InlineResponse200102Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse200102) GetHubsOk() (*[]InlineResponse200102Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse200102) SetHubs(v []InlineResponse200102Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse200102) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse200102) GetSubnets() []InlineResponse200102Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse200102) GetSubnetsOk() (*[]InlineResponse200102Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse200102) SetSubnets(v []InlineResponse200102Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse200102) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse200102) GetSubnet() InlineResponse200102Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse200102) GetSubnetOk() (*InlineResponse200102Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse200102) SetSubnet(v InlineResponse200102Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse200102) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetHostTranslations

`func (o *InlineResponse200102) GetHostTranslations() []InlineResponse200102HostTranslations`

GetHostTranslations returns the HostTranslations field if non-nil, zero value otherwise.

### GetHostTranslationsOk

`func (o *InlineResponse200102) GetHostTranslationsOk() (*[]InlineResponse200102HostTranslations, bool)`

GetHostTranslationsOk returns a tuple with the HostTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTranslations

`func (o *InlineResponse200102) SetHostTranslations(v []InlineResponse200102HostTranslations)`

SetHostTranslations sets HostTranslations field to given value.

### HasHostTranslations

`func (o *InlineResponse200102) HasHostTranslations() bool`

HasHostTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


