# InlineResponse200102Subnets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSubnet** | Pointer to **string** | The CIDR notation subnet used within the VPN | [optional] 
**UseVpn** | Pointer to **bool** | Indicates the presence of the subnet in the VPN | [optional] 
**Nat** | Pointer to [**InlineResponse200102Nat**](InlineResponse200102Nat.md) |  | [optional] 

## Methods

### NewInlineResponse200102Subnets

`func NewInlineResponse200102Subnets() *InlineResponse200102Subnets`

NewInlineResponse200102Subnets instantiates a new InlineResponse200102Subnets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102SubnetsWithDefaults

`func NewInlineResponse200102SubnetsWithDefaults() *InlineResponse200102Subnets`

NewInlineResponse200102SubnetsWithDefaults instantiates a new InlineResponse200102Subnets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSubnet

`func (o *InlineResponse200102Subnets) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *InlineResponse200102Subnets) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *InlineResponse200102Subnets) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.

### HasLocalSubnet

`func (o *InlineResponse200102Subnets) HasLocalSubnet() bool`

HasLocalSubnet returns a boolean if a field has been set.

### GetUseVpn

`func (o *InlineResponse200102Subnets) GetUseVpn() bool`

GetUseVpn returns the UseVpn field if non-nil, zero value otherwise.

### GetUseVpnOk

`func (o *InlineResponse200102Subnets) GetUseVpnOk() (*bool, bool)`

GetUseVpnOk returns a tuple with the UseVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVpn

`func (o *InlineResponse200102Subnets) SetUseVpn(v bool)`

SetUseVpn sets UseVpn field to given value.

### HasUseVpn

`func (o *InlineResponse200102Subnets) HasUseVpn() bool`

HasUseVpn returns a boolean if a field has been set.

### GetNat

`func (o *InlineResponse200102Subnets) GetNat() InlineResponse200102Nat`

GetNat returns the Nat field if non-nil, zero value otherwise.

### GetNatOk

`func (o *InlineResponse200102Subnets) GetNatOk() (*InlineResponse200102Nat, bool)`

GetNatOk returns a tuple with the Nat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat

`func (o *InlineResponse200102Subnets) SetNat(v InlineResponse200102Nat)`

SetNat sets Nat field to given value.

### HasNat

`func (o *InlineResponse200102Subnets) HasNat() bool`

HasNat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


