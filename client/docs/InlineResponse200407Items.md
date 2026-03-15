# InlineResponse200407Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200407Network**](InlineResponse200407Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200407Readings**](InlineResponse200407Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200407Items

`func NewInlineResponse200407Items() *InlineResponse200407Items`

NewInlineResponse200407Items instantiates a new InlineResponse200407Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200407ItemsWithDefaults

`func NewInlineResponse200407ItemsWithDefaults() *InlineResponse200407Items`

NewInlineResponse200407ItemsWithDefaults instantiates a new InlineResponse200407Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200407Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200407Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200407Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200407Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200407Items) GetNetwork() InlineResponse200407Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200407Items) GetNetworkOk() (*InlineResponse200407Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200407Items) SetNetwork(v InlineResponse200407Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200407Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200407Items) GetReadings() []InlineResponse200407Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200407Items) GetReadingsOk() (*[]InlineResponse200407Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200407Items) SetReadings(v []InlineResponse200407Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200407Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


