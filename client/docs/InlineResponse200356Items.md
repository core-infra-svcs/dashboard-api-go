# InlineResponse200356Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Readings** | Pointer to [**[]InlineResponse200356Readings**](InlineResponse200356Readings.md) | The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes | [optional] 

## Methods

### NewInlineResponse200356Items

`func NewInlineResponse200356Items() *InlineResponse200356Items`

NewInlineResponse200356Items instantiates a new InlineResponse200356Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200356ItemsWithDefaults

`func NewInlineResponse200356ItemsWithDefaults() *InlineResponse200356Items`

NewInlineResponse200356ItemsWithDefaults instantiates a new InlineResponse200356Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200356Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200356Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200356Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200356Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200356Items) GetReadings() []InlineResponse200356Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200356Items) GetReadingsOk() (*[]InlineResponse200356Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200356Items) SetReadings(v []InlineResponse200356Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200356Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


