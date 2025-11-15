# InlineResponse200382Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200382Network**](InlineResponse200382Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200382Readings**](InlineResponse200382Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200382Items

`func NewInlineResponse200382Items() *InlineResponse200382Items`

NewInlineResponse200382Items instantiates a new InlineResponse200382Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200382ItemsWithDefaults

`func NewInlineResponse200382ItemsWithDefaults() *InlineResponse200382Items`

NewInlineResponse200382ItemsWithDefaults instantiates a new InlineResponse200382Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200382Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200382Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200382Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200382Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200382Items) GetNetwork() InlineResponse200382Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200382Items) GetNetworkOk() (*InlineResponse200382Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200382Items) SetNetwork(v InlineResponse200382Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200382Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200382Items) GetReadings() []InlineResponse200382Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200382Items) GetReadingsOk() (*[]InlineResponse200382Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200382Items) SetReadings(v []InlineResponse200382Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200382Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


