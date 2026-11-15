# InlineResponse200114

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The site-to-site VPN mode. | [optional] 
**Hubs** | Pointer to [**[]InlineResponse200114Hubs**](InlineResponse200114Hubs.md) | The list of VPN hubs, in order of preference. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse200114Subnets**](InlineResponse200114Subnets.md) | The list of local VPN subnets and their VPN presence. VLAN entries include vlanId and subnetting where applicable. When the corresponding NFO is enabled, NAT settings are returned for every VLAN; non-VLAN entries omit vlanId, subnetting, and nat.legacy. | [optional] 
**Sgt** | Pointer to [**InlineResponse200114Sgt**](InlineResponse200114Sgt.md) |  | [optional] 
**Subnet** | Pointer to [**InlineResponse200114Subnet**](InlineResponse200114Subnet.md) |  | [optional] 
**HostTranslations** | Pointer to [**[]InlineResponse200114HostTranslations**](InlineResponse200114HostTranslations.md) | The list of VPN host translations. Host translations are supported starting from MX firmware version 26.1.2 | [optional] 
**NatTraversal** | Pointer to [**InlineResponse200114NatTraversal**](InlineResponse200114NatTraversal.md) |  | [optional] 

## Methods

### NewInlineResponse200114

`func NewInlineResponse200114() *InlineResponse200114`

NewInlineResponse200114 instantiates a new InlineResponse200114 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200114WithDefaults

`func NewInlineResponse200114WithDefaults() *InlineResponse200114`

NewInlineResponse200114WithDefaults instantiates a new InlineResponse200114 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200114) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200114) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200114) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200114) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHubs

`func (o *InlineResponse200114) GetHubs() []InlineResponse200114Hubs`

GetHubs returns the Hubs field if non-nil, zero value otherwise.

### GetHubsOk

`func (o *InlineResponse200114) GetHubsOk() (*[]InlineResponse200114Hubs, bool)`

GetHubsOk returns a tuple with the Hubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubs

`func (o *InlineResponse200114) SetHubs(v []InlineResponse200114Hubs)`

SetHubs sets Hubs field to given value.

### HasHubs

`func (o *InlineResponse200114) HasHubs() bool`

HasHubs returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse200114) GetSubnets() []InlineResponse200114Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse200114) GetSubnetsOk() (*[]InlineResponse200114Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse200114) SetSubnets(v []InlineResponse200114Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse200114) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse200114) GetSgt() InlineResponse200114Sgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse200114) GetSgtOk() (*InlineResponse200114Sgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse200114) SetSgt(v InlineResponse200114Sgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse200114) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse200114) GetSubnet() InlineResponse200114Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse200114) GetSubnetOk() (*InlineResponse200114Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse200114) SetSubnet(v InlineResponse200114Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse200114) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetHostTranslations

`func (o *InlineResponse200114) GetHostTranslations() []InlineResponse200114HostTranslations`

GetHostTranslations returns the HostTranslations field if non-nil, zero value otherwise.

### GetHostTranslationsOk

`func (o *InlineResponse200114) GetHostTranslationsOk() (*[]InlineResponse200114HostTranslations, bool)`

GetHostTranslationsOk returns a tuple with the HostTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostTranslations

`func (o *InlineResponse200114) SetHostTranslations(v []InlineResponse200114HostTranslations)`

SetHostTranslations sets HostTranslations field to given value.

### HasHostTranslations

`func (o *InlineResponse200114) HasHostTranslations() bool`

HasHostTranslations returns a boolean if a field has been set.

### GetNatTraversal

`func (o *InlineResponse200114) GetNatTraversal() InlineResponse200114NatTraversal`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *InlineResponse200114) GetNatTraversalOk() (*InlineResponse200114NatTraversal, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *InlineResponse200114) SetNatTraversal(v InlineResponse200114NatTraversal)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *InlineResponse200114) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


