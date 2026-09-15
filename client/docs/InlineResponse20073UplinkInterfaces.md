# InlineResponse20073UplinkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Interface name (wan1|wan2|wanN) | [optional] 
**Addresses** | Pointer to [**[]InlineResponse20073UplinkAddresses**](InlineResponse20073UplinkAddresses.md) | Addresses for this interface | [optional] 

## Methods

### NewInlineResponse20073UplinkInterfaces

`func NewInlineResponse20073UplinkInterfaces() *InlineResponse20073UplinkInterfaces`

NewInlineResponse20073UplinkInterfaces instantiates a new InlineResponse20073UplinkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20073UplinkInterfacesWithDefaults

`func NewInlineResponse20073UplinkInterfacesWithDefaults() *InlineResponse20073UplinkInterfaces`

NewInlineResponse20073UplinkInterfacesWithDefaults instantiates a new InlineResponse20073UplinkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20073UplinkInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20073UplinkInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20073UplinkInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20073UplinkInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddresses

`func (o *InlineResponse20073UplinkInterfaces) GetAddresses() []InlineResponse20073UplinkAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InlineResponse20073UplinkInterfaces) GetAddressesOk() (*[]InlineResponse20073UplinkAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InlineResponse20073UplinkInterfaces) SetAddresses(v []InlineResponse20073UplinkAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InlineResponse20073UplinkInterfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


