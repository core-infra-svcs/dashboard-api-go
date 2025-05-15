# InlineResponse200363Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse200363Interfaces**](InlineResponse200363Interfaces.md) | layer 2 interfaces belongs to the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200363Items

`func NewInlineResponse200363Items() *InlineResponse200363Items`

NewInlineResponse200363Items instantiates a new InlineResponse200363Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200363ItemsWithDefaults

`func NewInlineResponse200363ItemsWithDefaults() *InlineResponse200363Items`

NewInlineResponse200363ItemsWithDefaults instantiates a new InlineResponse200363Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200363Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200363Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200363Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200363Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse200363Items) GetInterfaces() []InlineResponse200363Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200363Items) GetInterfacesOk() (*[]InlineResponse200363Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200363Items) SetInterfaces(v []InlineResponse200363Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse200363Items) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


