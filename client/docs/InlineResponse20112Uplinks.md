# InlineResponse20112Uplinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20112Addresses**](InlineResponse20112Addresses.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse20112Uplinks

`func NewInlineResponse20112Uplinks() *InlineResponse20112Uplinks`

NewInlineResponse20112Uplinks instantiates a new InlineResponse20112Uplinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20112UplinksWithDefaults

`func NewInlineResponse20112UplinksWithDefaults() *InlineResponse20112Uplinks`

NewInlineResponse20112UplinksWithDefaults instantiates a new InlineResponse20112Uplinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *InlineResponse20112Uplinks) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20112Uplinks) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20112Uplinks) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20112Uplinks) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20112Uplinks) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20112Uplinks) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20112Uplinks) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20112Uplinks) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20112Uplinks) GetAddresses() []InlineResponse20112Addresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20112Uplinks) GetAddressesOk() (*[]InlineResponse20112Addresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20112Uplinks) SetAddresses(v []InlineResponse20112Addresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20112Uplinks) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


