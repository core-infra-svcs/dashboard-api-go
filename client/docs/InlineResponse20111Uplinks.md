# InlineResponse20111Uplinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20111Addresses**](InlineResponse20111Addresses.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse20111Uplinks

`func NewInlineResponse20111Uplinks() *InlineResponse20111Uplinks`

NewInlineResponse20111Uplinks instantiates a new InlineResponse20111Uplinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20111UplinksWithDefaults

`func NewInlineResponse20111UplinksWithDefaults() *InlineResponse20111Uplinks`

NewInlineResponse20111UplinksWithDefaults instantiates a new InlineResponse20111Uplinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *InlineResponse20111Uplinks) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20111Uplinks) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20111Uplinks) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20111Uplinks) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20111Uplinks) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20111Uplinks) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20111Uplinks) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20111Uplinks) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20111Uplinks) GetAddresses() []InlineResponse20111Addresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20111Uplinks) GetAddressesOk() (*[]InlineResponse20111Addresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20111Uplinks) SetAddresses(v []InlineResponse20111Addresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20111Uplinks) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


