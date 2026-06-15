# InlineResponse20042NextHops

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The priority of the next hop starting from 0. Lower numbers indicate higher priority. | [optional] 
**Address** | Pointer to **string** | The IP address of the next hop | [optional] 
**Vpn** | Pointer to [**InlineResponse20042Vpn**](InlineResponse20042Vpn.md) |  | [optional] 

## Methods

### NewInlineResponse20042NextHops

`func NewInlineResponse20042NextHops() *InlineResponse20042NextHops`

NewInlineResponse20042NextHops instantiates a new InlineResponse20042NextHops object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20042NextHopsWithDefaults

`func NewInlineResponse20042NextHopsWithDefaults() *InlineResponse20042NextHops`

NewInlineResponse20042NextHopsWithDefaults instantiates a new InlineResponse20042NextHops object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse20042NextHops) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse20042NextHops) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse20042NextHops) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse20042NextHops) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse20042NextHops) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20042NextHops) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20042NextHops) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse20042NextHops) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse20042NextHops) GetVpn() InlineResponse20042Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse20042NextHops) GetVpnOk() (*InlineResponse20042Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse20042NextHops) SetVpn(v InlineResponse20042Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse20042NextHops) HasVpn() bool`

HasVpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


