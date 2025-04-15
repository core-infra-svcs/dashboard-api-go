# InlineResponse200358Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse200358Interfaces**](InlineResponse200358Interfaces.md) | layer 3 interfaces belongs to the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200358Items

`func NewInlineResponse200358Items() *InlineResponse200358Items`

NewInlineResponse200358Items instantiates a new InlineResponse200358Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200358ItemsWithDefaults

`func NewInlineResponse200358ItemsWithDefaults() *InlineResponse200358Items`

NewInlineResponse200358ItemsWithDefaults instantiates a new InlineResponse200358Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200358Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200358Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200358Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200358Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse200358Items) GetInterfaces() []InlineResponse200358Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200358Items) GetInterfacesOk() (*[]InlineResponse200358Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200358Items) SetInterfaces(v []InlineResponse200358Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse200358Items) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


