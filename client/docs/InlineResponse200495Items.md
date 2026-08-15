# InlineResponse200495Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200495Network**](InlineResponse200495Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200495Readings**](InlineResponse200495Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200495Items

`func NewInlineResponse200495Items() *InlineResponse200495Items`

NewInlineResponse200495Items instantiates a new InlineResponse200495Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200495ItemsWithDefaults

`func NewInlineResponse200495ItemsWithDefaults() *InlineResponse200495Items`

NewInlineResponse200495ItemsWithDefaults instantiates a new InlineResponse200495Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200495Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200495Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200495Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200495Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200495Items) GetNetwork() InlineResponse200495Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200495Items) GetNetworkOk() (*InlineResponse200495Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200495Items) SetNetwork(v InlineResponse200495Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200495Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200495Items) GetReadings() []InlineResponse200495Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200495Items) GetReadingsOk() (*[]InlineResponse200495Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200495Items) SetReadings(v []InlineResponse200495Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200495Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


