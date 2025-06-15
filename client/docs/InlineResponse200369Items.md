# InlineResponse200369Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200369Network**](InlineResponse200369Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200369Readings**](InlineResponse200369Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200369Items

`func NewInlineResponse200369Items() *InlineResponse200369Items`

NewInlineResponse200369Items instantiates a new InlineResponse200369Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200369ItemsWithDefaults

`func NewInlineResponse200369ItemsWithDefaults() *InlineResponse200369Items`

NewInlineResponse200369ItemsWithDefaults instantiates a new InlineResponse200369Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200369Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200369Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200369Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200369Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200369Items) GetNetwork() InlineResponse200369Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200369Items) GetNetworkOk() (*InlineResponse200369Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200369Items) SetNetwork(v InlineResponse200369Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200369Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200369Items) GetReadings() []InlineResponse200369Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200369Items) GetReadingsOk() (*[]InlineResponse200369Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200369Items) SetReadings(v []InlineResponse200369Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200369Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


