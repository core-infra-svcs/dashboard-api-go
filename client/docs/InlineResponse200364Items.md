# InlineResponse200364Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The cloud ID of the wireless LAN controller | [optional] 
**Readings** | Pointer to [**[]InlineResponse200364Readings**](InlineResponse200364Readings.md) | The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes | [optional] 

## Methods

### NewInlineResponse200364Items

`func NewInlineResponse200364Items() *InlineResponse200364Items`

NewInlineResponse200364Items instantiates a new InlineResponse200364Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200364ItemsWithDefaults

`func NewInlineResponse200364ItemsWithDefaults() *InlineResponse200364Items`

NewInlineResponse200364ItemsWithDefaults instantiates a new InlineResponse200364Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200364Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200364Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200364Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200364Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200364Items) GetReadings() []InlineResponse200364Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200364Items) GetReadingsOk() (*[]InlineResponse200364Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200364Items) SetReadings(v []InlineResponse200364Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200364Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


