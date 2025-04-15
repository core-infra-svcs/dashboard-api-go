# InlineResponse200355Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse200355Interfaces**](InlineResponse200355Interfaces.md) | layer 2 interfaces belongs to the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200355Items

`func NewInlineResponse200355Items() *InlineResponse200355Items`

NewInlineResponse200355Items instantiates a new InlineResponse200355Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200355ItemsWithDefaults

`func NewInlineResponse200355ItemsWithDefaults() *InlineResponse200355Items`

NewInlineResponse200355ItemsWithDefaults instantiates a new InlineResponse200355Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200355Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200355Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200355Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200355Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse200355Items) GetInterfaces() []InlineResponse200355Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200355Items) GetInterfacesOk() (*[]InlineResponse200355Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200355Items) SetInterfaces(v []InlineResponse200355Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse200355Items) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


