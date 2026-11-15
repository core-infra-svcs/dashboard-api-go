# InlineResponse200537Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200537Network**](InlineResponse200537Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200537Readings**](InlineResponse200537Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200537Items

`func NewInlineResponse200537Items() *InlineResponse200537Items`

NewInlineResponse200537Items instantiates a new InlineResponse200537Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200537ItemsWithDefaults

`func NewInlineResponse200537ItemsWithDefaults() *InlineResponse200537Items`

NewInlineResponse200537ItemsWithDefaults instantiates a new InlineResponse200537Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200537Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200537Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200537Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200537Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200537Items) GetNetwork() InlineResponse200537Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200537Items) GetNetworkOk() (*InlineResponse200537Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200537Items) SetNetwork(v InlineResponse200537Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200537Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200537Items) GetReadings() []InlineResponse200537Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200537Items) GetReadingsOk() (*[]InlineResponse200537Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200537Items) SetReadings(v []InlineResponse200537Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200537Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


