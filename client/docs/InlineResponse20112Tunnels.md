# InlineResponse20112Tunnels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uplink** | Pointer to [**InlineResponse20112Uplink**](InlineResponse20112Uplink.md) |  | [optional] 
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20112Addresses1**](InlineResponse20112Addresses1.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse20112Tunnels

`func NewInlineResponse20112Tunnels() *InlineResponse20112Tunnels`

NewInlineResponse20112Tunnels instantiates a new InlineResponse20112Tunnels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20112TunnelsWithDefaults

`func NewInlineResponse20112TunnelsWithDefaults() *InlineResponse20112Tunnels`

NewInlineResponse20112TunnelsWithDefaults instantiates a new InlineResponse20112Tunnels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplink

`func (o *InlineResponse20112Tunnels) GetUplink() InlineResponse20112Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20112Tunnels) GetUplinkOk() (*InlineResponse20112Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20112Tunnels) SetUplink(v InlineResponse20112Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20112Tunnels) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse20112Tunnels) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20112Tunnels) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20112Tunnels) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20112Tunnels) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20112Tunnels) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20112Tunnels) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20112Tunnels) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20112Tunnels) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20112Tunnels) GetAddresses() []InlineResponse20112Addresses1`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20112Tunnels) GetAddressesOk() (*[]InlineResponse20112Addresses1, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20112Tunnels) SetAddresses(v []InlineResponse20112Addresses1)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20112Tunnels) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


