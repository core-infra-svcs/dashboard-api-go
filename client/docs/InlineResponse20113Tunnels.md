# InlineResponse20113Tunnels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uplink** | Pointer to [**InlineResponse20113Uplink**](InlineResponse20113Uplink.md) |  | [optional] 
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20113Addresses1**](InlineResponse20113Addresses1.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse20113Tunnels

`func NewInlineResponse20113Tunnels() *InlineResponse20113Tunnels`

NewInlineResponse20113Tunnels instantiates a new InlineResponse20113Tunnels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20113TunnelsWithDefaults

`func NewInlineResponse20113TunnelsWithDefaults() *InlineResponse20113Tunnels`

NewInlineResponse20113TunnelsWithDefaults instantiates a new InlineResponse20113Tunnels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplink

`func (o *InlineResponse20113Tunnels) GetUplink() InlineResponse20113Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20113Tunnels) GetUplinkOk() (*InlineResponse20113Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20113Tunnels) SetUplink(v InlineResponse20113Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20113Tunnels) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse20113Tunnels) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20113Tunnels) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20113Tunnels) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20113Tunnels) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20113Tunnels) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20113Tunnels) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20113Tunnels) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20113Tunnels) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20113Tunnels) GetAddresses() []InlineResponse20113Addresses1`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20113Tunnels) GetAddressesOk() (*[]InlineResponse20113Addresses1, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20113Tunnels) SetAddresses(v []InlineResponse20113Addresses1)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20113Tunnels) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


