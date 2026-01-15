# InlineResponse200399Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Readings** | Pointer to [**[]InlineResponse200399Readings**](InlineResponse200399Readings.md) | The status of packets counter on the interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200399Interfaces

`func NewInlineResponse200399Interfaces() *InlineResponse200399Interfaces`

NewInlineResponse200399Interfaces instantiates a new InlineResponse200399Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200399InterfacesWithDefaults

`func NewInlineResponse200399InterfacesWithDefaults() *InlineResponse200399Interfaces`

NewInlineResponse200399InterfacesWithDefaults instantiates a new InlineResponse200399Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200399Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200399Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200399Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200399Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200399Interfaces) GetReadings() []InlineResponse200399Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200399Interfaces) GetReadingsOk() (*[]InlineResponse200399Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200399Interfaces) SetReadings(v []InlineResponse200399Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200399Interfaces) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


