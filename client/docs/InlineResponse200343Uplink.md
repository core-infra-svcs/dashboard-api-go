# InlineResponse200343Uplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Name of the interface | [optional] 
**Vlan** | Pointer to **int32** | VLAN of the interface | [optional] 
**Addresses** | Pointer to [**[]InlineResponse200343UplinkAddresses**](InlineResponse200343UplinkAddresses.md) | Addresses of the interface | [optional] 

## Methods

### NewInlineResponse200343Uplink

`func NewInlineResponse200343Uplink() *InlineResponse200343Uplink`

NewInlineResponse200343Uplink instantiates a new InlineResponse200343Uplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200343UplinkWithDefaults

`func NewInlineResponse200343UplinkWithDefaults() *InlineResponse200343Uplink`

NewInlineResponse200343UplinkWithDefaults instantiates a new InlineResponse200343Uplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *InlineResponse200343Uplink) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse200343Uplink) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse200343Uplink) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse200343Uplink) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200343Uplink) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200343Uplink) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200343Uplink) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200343Uplink) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse200343Uplink) GetAddresses() []InlineResponse200343UplinkAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse200343Uplink) GetAddressesOk() (*[]InlineResponse200343UplinkAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse200343Uplink) SetAddresses(v []InlineResponse200343UplinkAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse200343Uplink) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


