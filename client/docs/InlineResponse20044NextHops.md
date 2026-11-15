# InlineResponse20044NextHops

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The priority of the next hop starting from 0. Lower numbers indicate higher priority. | [optional] 
**Address** | Pointer to **string** | The IP address of the next hop | [optional] 
**Vlan** | Pointer to [**InlineResponse20044Vlan**](InlineResponse20044Vlan.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse20044Vpn**](InlineResponse20044Vpn.md) |  | [optional] 

## Methods

### NewInlineResponse20044NextHops

`func NewInlineResponse20044NextHops() *InlineResponse20044NextHops`

NewInlineResponse20044NextHops instantiates a new InlineResponse20044NextHops object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20044NextHopsWithDefaults

`func NewInlineResponse20044NextHopsWithDefaults() *InlineResponse20044NextHops`

NewInlineResponse20044NextHopsWithDefaults instantiates a new InlineResponse20044NextHops object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse20044NextHops) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse20044NextHops) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse20044NextHops) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse20044NextHops) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse20044NextHops) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20044NextHops) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20044NextHops) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse20044NextHops) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20044NextHops) GetVlan() InlineResponse20044Vlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20044NextHops) GetVlanOk() (*InlineResponse20044Vlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20044NextHops) SetVlan(v InlineResponse20044Vlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20044NextHops) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse20044NextHops) GetVpn() InlineResponse20044Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse20044NextHops) GetVpnOk() (*InlineResponse20044Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse20044NextHops) SetVpn(v InlineResponse20044Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse20044NextHops) HasVpn() bool`

HasVpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


