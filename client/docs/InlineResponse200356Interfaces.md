# InlineResponse200356Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Readings** | Pointer to [**[]InlineResponse200356Readings**](InlineResponse200356Readings.md) | The status of packets counter on the interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200356Interfaces

`func NewInlineResponse200356Interfaces() *InlineResponse200356Interfaces`

NewInlineResponse200356Interfaces instantiates a new InlineResponse200356Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200356InterfacesWithDefaults

`func NewInlineResponse200356InterfacesWithDefaults() *InlineResponse200356Interfaces`

NewInlineResponse200356InterfacesWithDefaults instantiates a new InlineResponse200356Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200356Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200356Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200356Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200356Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200356Interfaces) GetReadings() []InlineResponse200356Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200356Interfaces) GetReadingsOk() (*[]InlineResponse200356Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200356Interfaces) SetReadings(v []InlineResponse200356Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200356Interfaces) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


