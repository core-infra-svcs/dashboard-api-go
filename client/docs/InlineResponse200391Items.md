# InlineResponse200391Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200391Network**](InlineResponse200391Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200391Readings**](InlineResponse200391Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200391Items

`func NewInlineResponse200391Items() *InlineResponse200391Items`

NewInlineResponse200391Items instantiates a new InlineResponse200391Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200391ItemsWithDefaults

`func NewInlineResponse200391ItemsWithDefaults() *InlineResponse200391Items`

NewInlineResponse200391ItemsWithDefaults instantiates a new InlineResponse200391Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200391Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200391Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200391Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200391Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200391Items) GetNetwork() InlineResponse200391Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200391Items) GetNetworkOk() (*InlineResponse200391Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200391Items) SetNetwork(v InlineResponse200391Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200391Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200391Items) GetReadings() []InlineResponse200391Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200391Items) GetReadingsOk() (*[]InlineResponse200391Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200391Items) SetReadings(v []InlineResponse200391Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200391Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


