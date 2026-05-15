# InlineResponse20113Uplinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20113Addresses**](InlineResponse20113Addresses.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse20113Uplinks

`func NewInlineResponse20113Uplinks() *InlineResponse20113Uplinks`

NewInlineResponse20113Uplinks instantiates a new InlineResponse20113Uplinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20113UplinksWithDefaults

`func NewInlineResponse20113UplinksWithDefaults() *InlineResponse20113Uplinks`

NewInlineResponse20113UplinksWithDefaults instantiates a new InlineResponse20113Uplinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *InlineResponse20113Uplinks) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20113Uplinks) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20113Uplinks) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20113Uplinks) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20113Uplinks) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20113Uplinks) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20113Uplinks) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20113Uplinks) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20113Uplinks) GetAddresses() []InlineResponse20113Addresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20113Uplinks) GetAddressesOk() (*[]InlineResponse20113Addresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20113Uplinks) SetAddresses(v []InlineResponse20113Addresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20113Uplinks) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


