# InlineResponse200255Uplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse200255UplinkAddresses**](InlineResponse200255UplinkAddresses.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse200255Uplink

`func NewInlineResponse200255Uplink() *InlineResponse200255Uplink`

NewInlineResponse200255Uplink instantiates a new InlineResponse200255Uplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200255UplinkWithDefaults

`func NewInlineResponse200255UplinkWithDefaults() *InlineResponse200255Uplink`

NewInlineResponse200255UplinkWithDefaults instantiates a new InlineResponse200255Uplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *InlineResponse200255Uplink) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse200255Uplink) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse200255Uplink) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse200255Uplink) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200255Uplink) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200255Uplink) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200255Uplink) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200255Uplink) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse200255Uplink) GetAddresses() []InlineResponse200255UplinkAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse200255Uplink) GetAddressesOk() (*[]InlineResponse200255UplinkAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse200255Uplink) SetAddresses(v []InlineResponse200255UplinkAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse200255Uplink) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


