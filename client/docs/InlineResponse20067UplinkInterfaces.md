# InlineResponse20067UplinkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Interface name (wan1|wan2|wanN) | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20067UplinkAddresses**](InlineResponse20067UplinkAddresses.md) | Addresses for this interface | [optional] 

## Methods

### NewInlineResponse20067UplinkInterfaces

`func NewInlineResponse20067UplinkInterfaces() *InlineResponse20067UplinkInterfaces`

NewInlineResponse20067UplinkInterfaces instantiates a new InlineResponse20067UplinkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20067UplinkInterfacesWithDefaults

`func NewInlineResponse20067UplinkInterfacesWithDefaults() *InlineResponse20067UplinkInterfaces`

NewInlineResponse20067UplinkInterfacesWithDefaults instantiates a new InlineResponse20067UplinkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20067UplinkInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20067UplinkInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20067UplinkInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20067UplinkInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20067UplinkInterfaces) GetAddresses() []InlineResponse20067UplinkAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20067UplinkInterfaces) GetAddressesOk() (*[]InlineResponse20067UplinkAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20067UplinkInterfaces) SetAddresses(v []InlineResponse20067UplinkAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20067UplinkInterfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


