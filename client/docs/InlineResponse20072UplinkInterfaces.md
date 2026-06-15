# InlineResponse20072UplinkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Interface name (wan1|wan2|wanN) | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20072UplinkAddresses**](InlineResponse20072UplinkAddresses.md) | Addresses for this interface | [optional] 

## Methods

### NewInlineResponse20072UplinkInterfaces

`func NewInlineResponse20072UplinkInterfaces() *InlineResponse20072UplinkInterfaces`

NewInlineResponse20072UplinkInterfaces instantiates a new InlineResponse20072UplinkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20072UplinkInterfacesWithDefaults

`func NewInlineResponse20072UplinkInterfacesWithDefaults() *InlineResponse20072UplinkInterfaces`

NewInlineResponse20072UplinkInterfacesWithDefaults instantiates a new InlineResponse20072UplinkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20072UplinkInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20072UplinkInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20072UplinkInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20072UplinkInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20072UplinkInterfaces) GetAddresses() []InlineResponse20072UplinkAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20072UplinkInterfaces) GetAddressesOk() (*[]InlineResponse20072UplinkAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20072UplinkInterfaces) SetAddresses(v []InlineResponse20072UplinkAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20072UplinkInterfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


