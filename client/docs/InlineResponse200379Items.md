# InlineResponse200379Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200379Network**](InlineResponse200379Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200379Readings**](InlineResponse200379Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200379Items

`func NewInlineResponse200379Items() *InlineResponse200379Items`

NewInlineResponse200379Items instantiates a new InlineResponse200379Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200379ItemsWithDefaults

`func NewInlineResponse200379ItemsWithDefaults() *InlineResponse200379Items`

NewInlineResponse200379ItemsWithDefaults instantiates a new InlineResponse200379Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200379Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200379Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200379Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200379Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200379Items) GetNetwork() InlineResponse200379Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200379Items) GetNetworkOk() (*InlineResponse200379Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200379Items) SetNetwork(v InlineResponse200379Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200379Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200379Items) GetReadings() []InlineResponse200379Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200379Items) GetReadingsOk() (*[]InlineResponse200379Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200379Items) SetReadings(v []InlineResponse200379Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200379Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


