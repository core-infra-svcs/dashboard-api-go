# InlineResponse20075UplinkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Interface name (wan1|wan2|wanN) | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20075UplinkAddresses**](InlineResponse20075UplinkAddresses.md) | Addresses for this interface | [optional] 

## Methods

### NewInlineResponse20075UplinkInterfaces

`func NewInlineResponse20075UplinkInterfaces() *InlineResponse20075UplinkInterfaces`

NewInlineResponse20075UplinkInterfaces instantiates a new InlineResponse20075UplinkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20075UplinkInterfacesWithDefaults

`func NewInlineResponse20075UplinkInterfacesWithDefaults() *InlineResponse20075UplinkInterfaces`

NewInlineResponse20075UplinkInterfacesWithDefaults instantiates a new InlineResponse20075UplinkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20075UplinkInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20075UplinkInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20075UplinkInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20075UplinkInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20075UplinkInterfaces) GetAddresses() []InlineResponse20075UplinkAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20075UplinkInterfaces) GetAddressesOk() (*[]InlineResponse20075UplinkAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20075UplinkInterfaces) SetAddresses(v []InlineResponse20075UplinkAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20075UplinkInterfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


