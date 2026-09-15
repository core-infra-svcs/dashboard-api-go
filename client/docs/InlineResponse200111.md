# InlineResponse200111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse200111Hubs**](InlineResponse200111Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse200111Subnets**](InlineResponse200111Subnets.md) | The list of subnets and their VPN presence. | [optional] 
**Sgt** | Pointer to [**InlineResponse200111Sgt**](InlineResponse200111Sgt.md) |  | [optional] 
**Subnet** | Pointer to [**InlineResponse200111Subnet**](InlineResponse200111Subnet.md) |  | [optional] 
**HostTranslations** | Pointer to [**[]InlineResponse200111HostTranslations**](InlineResponse200111HostTranslations.md) | The list of VPN host translations. Host translations are supported starting from MX firmware version 26.1.2 | [optional] 

## Methods

### NewInlineResponse200111

`func NewInlineResponse200111() *InlineResponse200111`

NewInlineResponse200111 instantiates a new InlineResponse200111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111WithDefaults

`func NewInlineResponse200111WithDefaults() *InlineResponse200111`

NewInlineResponse200111WithDefaults instantiates a new InlineResponse200111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200111) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200111) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200111) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200111) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse200111) GetHubs() []InlineResponse200111Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse200111) GetHubsOk() (*[]InlineResponse200111Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse200111) SetHubs(v []InlineResponse200111Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse200111) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse200111) GetSubnets() []InlineResponse200111Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse200111) GetSubnetsOk() (*[]InlineResponse200111Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse200111) SetSubnets(v []InlineResponse200111Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse200111) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse200111) GetSgt() InlineResponse200111Sgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse200111) GetSgtOk() (*InlineResponse200111Sgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse200111) SetSgt(v InlineResponse200111Sgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse200111) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse200111) GetSubnet() InlineResponse200111Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse200111) GetSubnetOk() (*InlineResponse200111Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse200111) SetSubnet(v InlineResponse200111Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse200111) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetHostTranslations

`func (o *InlineResponse200111) GetHostTranslations() []InlineResponse200111HostTranslations`

GetHostTranslations returns the HostTranslations field if non-nil, zero value otherwise.

### GetHostTranslationsOk

`func (o *InlineResponse200111) GetHostTranslationsOk() (*[]InlineResponse200111HostTranslations, bool)`

GetHostTranslationsOk returns a tuple with the HostTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTranslations

`func (o *InlineResponse200111) SetHostTranslations(v []InlineResponse200111HostTranslations)`

SetHostTranslations sets HostTranslations field to given value.

### HasHostTranslations

`func (o *InlineResponse200111) HasHostTranslations() bool`

HasHostTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


