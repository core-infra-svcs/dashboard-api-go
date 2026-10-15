# InlineResponse200533Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Readings** | Pointer to [**[]InlineResponse200533Readings**](InlineResponse200533Readings.md) | The status of packets counter on the interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200533Interfaces

`func NewInlineResponse200533Interfaces() *InlineResponse200533Interfaces`

NewInlineResponse200533Interfaces instantiates a new InlineResponse200533Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200533InterfacesWithDefaults

`func NewInlineResponse200533InterfacesWithDefaults() *InlineResponse200533Interfaces`

NewInlineResponse200533InterfacesWithDefaults instantiates a new InlineResponse200533Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200533Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200533Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200533Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200533Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200533Interfaces) GetReadings() []InlineResponse200533Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200533Interfaces) GetReadingsOk() (*[]InlineResponse200533Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200533Interfaces) SetReadings(v []InlineResponse200533Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200533Interfaces) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


