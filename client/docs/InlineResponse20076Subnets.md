# InlineResponse20076Subnets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSubnet** | Pointer to **string** | The CIDR notation subnet used within the VPN | [optional] 
**UseVpn** | Pointer to **bool** | Indicates the presence of the subnet in the VPN | [optional] 
**Nat** | Pointer to [**InlineResponse20076Nat**](InlineResponse20076Nat.md) |  | [optional] 

## Methods

### NewInlineResponse20076Subnets

`func NewInlineResponse20076Subnets() *InlineResponse20076Subnets`

NewInlineResponse20076Subnets instantiates a new InlineResponse20076Subnets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20076SubnetsWithDefaults

`func NewInlineResponse20076SubnetsWithDefaults() *InlineResponse20076Subnets`

NewInlineResponse20076SubnetsWithDefaults instantiates a new InlineResponse20076Subnets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSubnet

`func (o *InlineResponse20076Subnets) GetLocalSubnet() string`

GetLocalSubnet returns the LocalSubnet field if non-nil, zero value otherwise.

### GetLocalSubnetOk

`func (o *InlineResponse20076Subnets) GetLocalSubnetOk() (*string, bool)`

GetLocalSubnetOk returns a tuple with the LocalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnet

`func (o *InlineResponse20076Subnets) SetLocalSubnet(v string)`

SetLocalSubnet sets LocalSubnet field to given value.

### HasLocalSubnet

`func (o *InlineResponse20076Subnets) HasLocalSubnet() bool`

HasLocalSubnet returns a boolean if a field has been set.

### GetUseVpn

`func (o *InlineResponse20076Subnets) GetUseVpn() bool`

GetUseVpn returns the UseVpn field if non-nil, zero value otherwise.

### GetUseVpnOk

`func (o *InlineResponse20076Subnets) GetUseVpnOk() (*bool, bool)`

GetUseVpnOk returns a tuple with the UseVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVpn

`func (o *InlineResponse20076Subnets) SetUseVpn(v bool)`

SetUseVpn sets UseVpn field to given value.

### HasUseVpn

`func (o *InlineResponse20076Subnets) HasUseVpn() bool`

HasUseVpn returns a boolean if a field has been set.

### GetNat

`func (o *InlineResponse20076Subnets) GetNat() InlineResponse20076Nat`

GetNat returns the Nat field if non-nil, zero value otherwise.

### GetNatOk

`func (o *InlineResponse20076Subnets) GetNatOk() (*InlineResponse20076Nat, bool)`

GetNatOk returns a tuple with the Nat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNat

`func (o *InlineResponse20076Subnets) SetNat(v InlineResponse20076Nat)`

SetNat sets Nat field to given value.

### HasNat

`func (o *InlineResponse20076Subnets) HasNat() bool`

HasNat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


